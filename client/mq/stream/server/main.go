package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	sdklog "cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/server/grpc/gogoreflection"
	"github.com/google/uuid"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

func main() {
	logger := sdklog.NewLogger(os.Stderr)
	if err := startMQStream(context.Background(), logger); err != nil && !errors.Is(err, context.Canceled) {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startMQStream(ctx context.Context, logger sdklog.Logger) error {
	config, err := parseConfig()
	if err != nil {
		return err
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(config.KafkaBrokers...),
		kgo.ConsumerGroup(uuid.New().String()),
		kgo.ConsumeTopics(config.Topic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return err
	}

	defer client.Close()

	signalCtx, stopSignals := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stopSignals()

	cctx, cancel := context.WithCancelCause(signalCtx)
	defer cancel(nil)

	var grpcServerOptions []grpc.ServerOption
	if config.EnforceKeepalive {
		grpcServerOptions = []grpc.ServerOption{
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
				MinTime: config.MinClientPingInterval,
			}),
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle: config.MaxConnectionIdle,
				Time:              config.ServerPingInterval,
				Timeout:           config.ServerPingResponseTimeout,
			}),
		}
	}

	srv := NewStreamServer(logger, client)
	grpcServer := grpc.NewServer(grpcServerOptions...)
	types.RegisterEventStreamServer(grpcServer, srv)
	gogoreflection.Register(grpcServer)

	// start subscription
	go func() {
		if err := srv.Poll(cctx); err != nil {
			logger.Error(err.Error())
		}
	}()

	listener, err := net.Listen("tcp", strings.TrimPrefix(config.ListenAddress, "tcp://"))
	if err != nil {
		return err
	}

	defer func() {
		_ = listener.Close()
	}()

	serveErrCh := make(chan error, 1)
	go func() {
		serveErr := grpcServer.Serve(listener)
		if errors.Is(serveErr, grpc.ErrServerStopped) {
			serveErr = nil
		}

		serveErrCh <- serveErr
	}()

	logger.Info("event stream server started", "address", config.ListenAddress)

	select {
	case <-ctx.Done():
		logger.Info("stopping event stream server")
		grpcServer.Stop()
		if err := context.Cause(ctx); err != nil && !errors.Is(err, context.Canceled) {
			return err
		}

		return nil

	case err := <-serveErrCh:
		cancel(err)
		return err
	}
}

type StreamServer struct {
	logger        sdklog.Logger
	kafka         *kgo.Client
	mtx           sync.RWMutex
	subscriptions map[string]chan *types.EventStreamResponse
}

func NewStreamServer(l sdklog.Logger, k *kgo.Client) *StreamServer {
	return &StreamServer{
		logger:        l,
		kafka:         k,
		subscriptions: make(map[string]chan *types.EventStreamResponse),
	}
}

func (srv *StreamServer) Poll(ctx context.Context) error {
	notifyAll := func(msg *types.EventStreamResponse) {
		srv.mtx.RLock()
		defer srv.mtx.RUnlock()

		for _, sub := range srv.subscriptions {
			select {
			case sub <- msg:
			default: // subscriber not consuming
				srv.logger.Warn("consumer not reading the channel, skipping")
			}
		}
	}

	for {
		fetches := srv.kafka.PollFetches(ctx)
		if err := fetches.Err(); err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) {
				return nil
			}

			err = fmt.Errorf("poll kafka fetches: %w", err)
			srv.logger.Error("error polling fetches, stopping consumer loop", "err", err.Error())
			return err
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			var msg types.EventStreamResponse
			if err := msg.Unmarshal(record.Value); err != nil {
				srv.logger.Error("error decoding event message", "err", err.Error())
				continue
			}

			notifyAll(&msg)
		}
	}
}

func (srv *StreamServer) EventStream(_ *types.EventStreamRequest, server types.EventStream_EventStreamServer) error {
	id := uuid.New().String()
	ch := make(chan *types.EventStreamResponse, 100)

	srv.mtx.Lock()
	srv.subscriptions[id] = ch
	srv.mtx.Unlock()

	defer func() {
		srv.mtx.Lock()
		delete(srv.subscriptions, id)
		srv.mtx.RUnlock()
	}()

	for {
		select {
		case <-server.Context().Done():
			return server.Context().Err()
		case message, ok := <-ch:
			if !ok {
				return errors.New("stream closed")
			}

			if err := server.Send(message); err != nil {
				return fmt.Errorf("error sending message to client: %w", err)
			}
		}
	}
}

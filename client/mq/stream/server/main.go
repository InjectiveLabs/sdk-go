package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	sdklog "cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/server/grpc/gogoreflection"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	"github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

func main() {
	var (
		listenAddress             = flag.String(flagMQStreamListenAddress, "0.0.0.0:9988", "Listen address")
		kafkaBrokers              = flag.String(flagMQStreamKafkaBrokers, "", "Comma-separated Kafka broker addresses")
		enforceKeepalive          = flag.Bool(flagMQStreamEnforceKeepalive, false, "Define if Keepalive configuration params should be applied to MQ event stream gRPC server")
		minClientPingInterval     = flag.Duration(flagMQStreamMinClientPingInterval, defaultMQStreamMinClientPingInterval, "Duration a client should wait before sending a keepalive ping")
		maxConnectionIdle         = flag.Duration(flagMQStreamMaxConnectionIdle, defaultMQStreamMaxConnectionIdle, "Duration a connection is allowed to stay idle before forcing the disconnection")
		serverPingInterval        = flag.Duration(flagMQStreamServerPingInterval, defaultMQStreamServerPingInterval, "Duration after which the server will send a keepalive ping to the client on an idle connection")
		serverPingResponseTimeout = flag.Duration(flagMQStreamServerPingResponseTimeout, defaultMQStreamServerPingResponseTimeout, "Duration the server waits for the client to respond to a ping message before forcing a disconnection")
	)

	flag.Parse()

	cfg := mqStreamConfig{
		ListenAddress:             *listenAddress,
		EnforceKeepalive:          *enforceKeepalive,
		MinClientPingInterval:     *minClientPingInterval,
		MaxConnectionIdle:         *maxConnectionIdle,
		ServerPingInterval:        *serverPingInterval,
		ServerPingResponseTimeout: *serverPingResponseTimeout,
	}

	if *kafkaBrokers != "" {
		cfg.KafkaBrokers = strings.Split(*kafkaBrokers, ",")
	}

	if err := startMQStream(context.Background(), cfg); err != nil && !errors.Is(err, context.Canceled) {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startMQStream(ctx context.Context, cfg mqStreamConfig) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	signalCtx, stopSignals := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stopSignals()

	ctx, cancel := context.WithCancelCause(signalCtx)
	defer cancel(nil)

	var grpcServerOptions []grpc.ServerOption
	if cfg.EnforceKeepalive {
		grpcServerOptions = []grpc.ServerOption{
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
				MinTime: cfg.MinClientPingInterval,
			}),
			grpc.KeepaliveParams(keepalive.ServerParameters{
				MaxConnectionIdle: cfg.MaxConnectionIdle,
				Time:              cfg.ServerPingInterval,
				Timeout:           cfg.ServerPingResponseTimeout,
			}),
		}
	}

	logger := sdklog.NewLogger(os.Stderr)
	srv := NewStreamServer(logger, cfg)
	grpcServer := grpc.NewServer(grpcServerOptions...)
	types.RegisterEventStreamServer(grpcServer, srv)
	gogoreflection.Register(grpcServer)

	listener, err := net.Listen("tcp", strings.TrimPrefix(cfg.ListenAddress, "tcp://"))
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

	logger.Info("event stream server started", "address", cfg.ListenAddress)

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
	logger       sdklog.Logger
	kafkaBrokers []string
}

func NewStreamServer(l sdklog.Logger, cfg mqStreamConfig) *StreamServer {
	return &StreamServer{
		logger:       l,
		kafkaBrokers: cfg.KafkaBrokers,
	}
}

func (srv *StreamServer) EventStream(req *types.EventStreamRequest, server types.EventStream_EventStreamServer) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "event stream request cannot be nil")
	}

	consumerID := strings.TrimSpace(req.ConsumerId)
	if consumerID == "" {
		return status.Error(codes.InvalidArgument, "consumer id cannot be empty")
	}

	topic := strings.TrimSpace(req.Topic)
	if topic == "" {
		return status.Error(codes.InvalidArgument, "topic cannot be empty")
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(srv.kafkaBrokers...),
		kgo.ConsumerGroup(consumerID),
		kgo.ConsumeTopics(topic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return err
	}

	defer client.Close()

	srv.logger.Info("event stream consumer started", "consumer_id", consumerID, "topic", topic)
	defer srv.logger.Info("event stream consumer stopped", "consumer_id", consumerID, "topic", topic)

	for {
		fetches := client.PollFetches(server.Context())
		if err := fetches.Err(); err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(server.Context().Err(), context.Canceled) {
				return server.Context().Err()
			}

			return fmt.Errorf("poll kafka fetches: %w", err)
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			var msg types.EventStreamResponse
			if err := msg.Unmarshal(record.Value); err != nil {
				srv.logger.Error("error decoding event message", "err", err.Error())
				continue
			}

			if err := server.Send(&msg); err != nil {
				return fmt.Errorf("error sending message to client: %w", err)
			}
		}
	}
}

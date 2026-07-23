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
	"github.com/spf13/cobra"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

func main() {
	logger := sdklog.NewLogger(os.Stderr)
	cmd := &cobra.Command{
		Use:          "stream",
		Short:        "Run the MQ stream server",
		SilenceUsage: true,
		Long: `Run the gRPC stream for MQ Consumer. It runs as MQ stream-only mode.
For '--listen-address' option, it should be a valid address like '0.0.0.0:9090'.
For '--kafka-brokers', provide Kafka-compatible broker addresses.
For '--topic', provide the topic name that the MQ publisher node has already filled with.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return startMQStream(cmd, logger)
		},
	}

	cmd.Flags().String(flagMQStreamListenAddress, "0.0.0.0:9988", "Listen address")
	cmd.Flags().StringSlice(flagMQStreamKafkaBrokers, []string{}, "Kafka broker addresses")
	cmd.Flags().String(flagMQStreamTopicName, "", "Topic to consume")
	cmd.Flags().Bool(flagMQStreamEnforceKeepalive, false,
		"Define if Keepalive configuration params should be applied to MQ event stream gRPC server",
	)
	cmd.Flags().Duration(flagMQStreamMinClientPingInterval, defaultMQStreamMinClientPingInterval,
		"Duration a client should wait before sending a keepalive ping",
	)
	cmd.Flags().Duration(flagMQStreamMaxConnectionIdle, defaultMQStreamMaxConnectionIdle,
		"Duration a connection is allowed to stay idle before forcing the disconnection",
	)
	cmd.Flags().Duration(flagMQStreamServerPingInterval, defaultMQStreamServerPingInterval,
		"Duration after which the server will send a keepalive ping to the client on an idle connection",
	)
	cmd.Flags().Duration(flagMQStreamServerPingResponseTimeout, defaultMQStreamServerPingResponseTimeout,
		"Duration the server waits for the client to respond to a ping message before forcing a disconnection",
	)

	if err := cmd.Execute(); err != nil && !errors.Is(err, context.Canceled) {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startMQStream(cmd *cobra.Command, logger sdklog.Logger) error {
	config, err := parseConfig(cmd)
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

	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var mtx sync.RWMutex
	subscriptions := make(map[string]chan *types.EventStreamResponse)

	topicCh := make(chan *types.EventStreamResponse, 1)
	go func() {
		defer func() {
			stop() // kill the ctx for stream as well
			close(topicCh)
		}()

		for {
			fetches := client.PollFetches(ctx)
			if err := fetches.Err(); err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) {
					return
				}

				logger.Error("error polling fetches, stopping consumer loop", "err", err.Error())
				return
			}

			iter := fetches.RecordIter()
			for !iter.Done() {
				record := iter.Next()
				var msg types.EventStreamResponse
				if err := msg.Unmarshal(record.Value); err != nil {
					logger.Error("error decoding event message", "err", err.Error())
					continue
				}

				mtx.RLock()
				for _, sub := range subscriptions {
					select {
					case sub <- &msg:
					default: // subscriber not consuming
					}
				}
				mtx.RUnlock()
			}
		}
	}()

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

	subscribe := func() (<-chan *types.EventStreamResponse, func()) {
		ch := make(chan *types.EventStreamResponse, 100)
		id := uuid.New().String()

		mtx.Lock()
		defer mtx.Unlock()

		subscriptions[id] = ch

		cancel := func() {
			mtx.Lock()
			defer mtx.Unlock()

			delete(subscriptions, id)
		}

		return ch, cancel
	}

	grpcServer := grpc.NewServer(grpcServerOptions...)
	types.RegisterEventStreamServer(grpcServer, NewStreamServer(subscribe))
	gogoreflection.Register(grpcServer)

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
		return nil

	case err := <-serveErrCh:
		stop()
		return err
	}
}

type StreamServer struct {
	subscribe func() (<-chan *types.EventStreamResponse, func())
}

func NewStreamServer(sub func() (<-chan *types.EventStreamResponse, func())) StreamServer {
	return StreamServer{subscribe: sub}
}

func (srv StreamServer) EventStream(_ *types.EventStreamRequest, server types.EventStream_EventStreamServer) error {
	sub, cancel := srv.subscribe()
	defer cancel()

	for {
		select {
		case <-server.Context().Done():
			return nil
		case message, ok := <-sub:
			if !ok {
				return errors.New("stream closed")
			}

			if err := server.Send(message); err != nil {
				return fmt.Errorf("error sending message to client: %w", err)
			}
		}
	}
}

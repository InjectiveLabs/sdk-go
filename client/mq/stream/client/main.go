package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	mqtypes "github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

var kacp = keepalive.ClientParameters{
	Time:                30 * time.Second, // send pings every 30 seconds if there is no activity
	Timeout:             5 * time.Second,  // wait 5 second for ping ack before considering the connection dead
	PermitWithoutStream: false,            // do not send pings without active streams
}

func main() {
	var (
		address        = flag.String(flagMQStreamClientAddress, "localhost:9988", "MQ gRPC stream server address")
		consumerID     = flag.String(flagMQStreamClientConsumerID, "", "Stable Kafka consumer id")
		consumerIDFile = flag.String(flagMQStreamClientConsumerIDFile, defaultConsumerIDFile, "File used to persist a generated Kafka consumer id")
		topic          = flag.String(flagMQStreamClientTopic, "", "Topic to consume")
		format         = flag.String(flagMQStreamClientFormat, "verbose", "Output format: verbose or minimal")
		eventsDir      = flag.String(flagMQStreamClientEventsDir, "", "Directory to write one JSON file per streamed block")
	)

	flag.Parse()

	cfg := mqStreamClientConfig{
		Address:        *address,
		ConsumerID:     *consumerID,
		ConsumerIDFile: *consumerIDFile,
		Topic:          *topic,
		Format:         *format,
		EventsDir:      *eventsDir,
	}

	if err := startMQStreamClient(context.Background(), cfg); err != nil && !errors.Is(err, context.Canceled) {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startMQStreamClient(ctx context.Context, cfg mqStreamClientConfig) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	resolvedConsumerID, err := resolveConsumerID(cfg.ConsumerID, cfg.ConsumerIDFile)
	if err != nil {
		return err
	}

	cdc, err := newPublishEventDecoder()
	if err != nil {
		return fmt.Errorf("failed to initialize publish event decoder: %w", err)
	}

	cc, err := grpc.NewClient(
		cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
	)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}

	defer func() {
		_ = cc.Close()
	}()

	stream, err := mqtypes.NewEventStreamClient(cc).EventStream(ctx, &mqtypes.EventStreamRequest{
		ConsumerId: resolvedConsumerID,
		Topic:      cfg.Topic,
	})
	if err != nil {
		return fmt.Errorf("failed to start event stream: %w", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if isExpectedStreamClose(err) {
				return nil
			}

			return fmt.Errorf("event stream failed: %w", err)
		}

		if err := writeEventsFile(cfg.EventsDir, res, cdc); err != nil {
			return fmt.Errorf("failed to write events file: %w", err)
		}

		switch cfg.Format {
		case "minimal":
			printMinimal(res)
		case "verbose":
			printVerbose(res, cdc)
		default:
			return fmt.Errorf("unsupported format %q", cfg.Format)
		}
	}
}

func isExpectedStreamClose(err error) bool {
	if errors.Is(err, context.Canceled) || errors.Is(err, io.EOF) {
		return true
	}

	st, ok := status.FromError(err)
	if !ok {
		return false
	}

	switch st.Code() {
	case codes.Canceled:
		return true
	case codes.Unavailable:
		msg := st.Message()
		return strings.Contains(msg, "EOF") ||
			strings.Contains(msg, "transport is closing") ||
			strings.Contains(msg, "client connection is closing")
	default:
		return false
	}
}

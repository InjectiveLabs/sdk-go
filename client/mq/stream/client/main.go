package main

import (
	"context"
	"errors"
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
	if err := startMQStreamClient(context.Background()); err != nil && !errors.Is(err, context.Canceled) {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startMQStreamClient(ctx context.Context) error {
	config, err := parseConfig()
	if err != nil {
		return err
	}

	publishEventDecoder, err := newPublishEventDecoder()
	if err != nil {
		return fmt.Errorf("failed to initialize publish event decoder: %w", err)
	}

	cc, err := grpc.NewClient(
		config.Address,
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
		ConsumerId: config.ConsumerID,
		Topic:      config.Topic,
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

		if err := writeEventsFile(config.EventsDir, res, publishEventDecoder); err != nil {
			return fmt.Errorf("failed to write events file: %w", err)
		}

		switch config.Format {
		case "minimal":
			printMinimal(res)
		case "verbose":
			printVerbose(res, publishEventDecoder)
		default:
			return fmt.Errorf("unsupported format %q", config.Format)
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

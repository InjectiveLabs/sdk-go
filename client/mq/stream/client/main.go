package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
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
	address := flag.String("address", "localhost:9988", "MQ gRPC stream server address")
	format := flag.String("format", "verbose", "Output format: verbose or minimal")
	eventsDir := flag.String("events-dir", "", "Directory to write one JSON file per streamed block")
	flag.Parse()

	if err := run(*address, *format, *eventsDir); err != nil {
		log.Fatal(err)
	}
}

func run(address, format, eventsDir string) error {
	cc, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(kacp),
	)
	// nolint:staticcheck //ignored on purpose
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer cc.Close()

	client := mqtypes.NewEventStreamClient(cc)

	ctx := context.Background()
	stream, err := client.EventStream(ctx, &mqtypes.EventStreamRequest{})
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

		if err := writeEventsFile(eventsDir, res); err != nil {
			return fmt.Errorf("failed to write events file: %w", err)
		}

		switch format {
		case "minimal":
			printMinimal(res)
		case "verbose":
			printVerbose(res)
		default:
			return fmt.Errorf("unsupported format %q", format)
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

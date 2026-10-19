package main

import (
	"context"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"

	mqtypes "github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

type largeBlockServer struct {
	mqtypes.UnimplementedEventStreamServer
}

func (*largeBlockServer) EventStream(_ *mqtypes.EventStreamRequest, stream mqtypes.EventStream_EventStreamServer) error {
	return stream.Send(&mqtypes.EventStreamResponse{
		BlockHeight:           1,
		AdditionalDataEntries: []*mqtypes.AdditionalDataEntry{{Data: make([]byte, 5<<20)}},
	})
}

func TestClientReceivesBlocksLargerThanDefaultGRPCLimit(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	server := grpc.NewServer()
	mqtypes.RegisterEventStreamServer(server, &largeBlockServer{})
	go func() { _ = server.Serve(listener) }()
	defer server.Stop()

	for _, tc := range []struct {
		name      string
		limit     int
		wantError bool
	}{
		{name: "compatible default"},
		{name: "explicit small limit", limit: 4 << 20, wantError: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			err := startMQStreamClient(ctx, mqStreamClientConfig{
				Address: listener.Addr().String(), ConsumerID: "large-block-test", Topic: "latest", Format: "minimal",
				MaxRecvMsgBytes: tc.limit,
			})
			if (err != nil) != tc.wantError {
				t.Fatalf("err=%v, wantError=%v", err, tc.wantError)
			}
		})
	}
}

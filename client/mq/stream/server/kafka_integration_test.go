package main

import (
	"context"
	"errors"
	"net"
	"os"
	"strings"
	"testing"
	"time"

	sdklog "cosmossdk.io/log"
	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"github.com/InjectiveLabs/sdk-go/chain/mq/types"
	"github.com/InjectiveLabs/sdk-go/client/mq/kafkaconfig"
)

// TestKafkaStreamTLSAndACLs requires a disposable Kafka fixture; it writes a
// record and advances the allowed consumer group. See the MQ README for setup.
func TestKafkaStreamTLSAndACLs(t *testing.T) {
	brokers := os.Getenv("MQ_TEST_KAFKA_BROKERS")
	if brokers == "" {
		t.Skip("set MQ_TEST_KAFKA_BROKERS to test a disposable TLS/SCRAM Kafka fixture")
	}
	for _, name := range []string{"MQ_KAFKA_USERNAME", "MQ_KAFKA_PASSWORD", "MQ_KAFKA_CA_FILE", "MQ_TEST_WRITER_USERNAME", "MQ_TEST_WRITER_PASSWORD"} {
		if os.Getenv(name) == "" {
			t.Fatalf("%s is required for the Kafka integration test", name)
		}
	}
	readerOptions, err := kafkaconfig.OptionsFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	writerOptions := kafkaOptionsForCredentials(t, os.Getenv("MQ_TEST_WRITER_USERNAME"), os.Getenv("MQ_TEST_WRITER_PASSWORD"))
	writerOptions = append(writerOptions, kgo.SeedBrokers(strings.Split(brokers, ",")...))
	writer, err := kgo.NewClient(writerOptions...)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(writer.Close)
	height := time.Now().UnixNano()
	value, err := (&types.EventStreamResponse{BlockHeight: height, AppHash: []byte("tls-scram-acl-test")}).Marshal()
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := writer.ProduceSync(ctx, &kgo.Record{Topic: "latest", Value: value}).FirstErr(); err != nil {
		t.Fatalf("seed allowed topic: %v", err)
	}

	for _, tc := range []struct {
		name          string
		topic         string
		group         string
		wrongPassword bool
		wantError     string
	}{
		{name: "allowed", topic: "latest", group: "mm-rfq-inj1jvrnhahd"},
		{name: "denied-topic", topic: "forbidden", group: "mm-rfq-inj1jvrnhahd", wantError: "TOPIC_AUTHORIZATION_FAILED"},
		{name: "denied-group", topic: "latest", group: "not-authorized", wantError: "GROUP_AUTHORIZATION_FAILED"},
		{name: "wrong-password", topic: "latest", group: "mm-rfq-inj1jvrnhahd", wrongPassword: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			opts := append([]kgo.Opt{}, readerOptions...)
			if tc.wrongPassword {
				opts = kafkaOptionsForCredentials(t, os.Getenv("MQ_KAFKA_USERNAME"), os.Getenv("MQ_KAFKA_PASSWORD")+"-incorrect")
				probeOptions := append([]kgo.Opt{}, opts...)
				probeOptions = append(probeOptions, kgo.SeedBrokers(strings.Split(brokers, ",")...))
				probe, err := kgo.NewClient(probeOptions...)
				if err != nil {
					t.Fatal(err)
				}
				defer probe.Close()
				authCtx, cancelAuth := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancelAuth()
				if err := probe.Ping(authCtx); !errors.Is(err, kerr.SaslAuthenticationFailed) {
					t.Fatalf("expected broker SCRAM rejection, received %v", err)
				}
			}
			client := newKafkaTestStreamClient(t, strings.Split(brokers, ","), opts)
			timeout := 30 * time.Second
			if tc.wrongPassword {
				// The consumer retries failed authentication; it must deliver no
				// record while the broker rejects the same credentials above.
				timeout = 2 * time.Second
			}
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			stream, err := client.EventStream(ctx, &types.EventStreamRequest{Topic: tc.topic, ConsumerId: tc.group})
			if err != nil {
				t.Fatal(err)
			}
			for {
				response, err := stream.Recv()
				if tc.wrongPassword {
					if response != nil || status.Code(err) != codes.DeadlineExceeded {
						t.Fatalf("invalid credentials should deliver no stream record: %v", err)
					}
					return
				}
				if tc.wantError != "" {
					if err == nil || !strings.Contains(err.Error(), tc.wantError) {
						t.Fatalf("expected broker rejection %s, received %v", tc.wantError, err)
					}
					return
				}
				if err != nil {
					t.Fatalf("receive authorized Kafka record: %v", err)
				}
				if response.BlockHeight == height {
					if string(response.AppHash) != "tls-scram-acl-test" {
						t.Fatal("stream response changed the record payload")
					}
					return
				}
			}
		})
	}
}

func kafkaOptionsForCredentials(t *testing.T, username, password string) []kgo.Opt {
	t.Helper()
	originalUser, originalPassword := os.Getenv("MQ_KAFKA_USERNAME"), os.Getenv("MQ_KAFKA_PASSWORD")
	t.Setenv("MQ_KAFKA_USERNAME", username)
	t.Setenv("MQ_KAFKA_PASSWORD", password)
	defer t.Setenv("MQ_KAFKA_USERNAME", originalUser)
	defer t.Setenv("MQ_KAFKA_PASSWORD", originalPassword)
	opts, err := kafkaconfig.OptionsFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	return opts
}

func newKafkaTestStreamClient(t *testing.T, brokers []string, opts []kgo.Opt) types.EventStreamClient {
	t.Helper()
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	types.RegisterEventStreamServer(server, NewStreamServer(sdklog.NewNopLogger(), mqStreamConfig{KafkaBrokers: brokers}, opts...))
	done := make(chan error, 1)
	go func() { done <- server.Serve(listener) }()
	t.Cleanup(func() {
		server.Stop()
		_ = listener.Close()
		if err := <-done; err != nil {
			t.Errorf("stop test gRPC server: %v", err)
		}
	})
	conn, err := grpc.NewClient("passthrough:///bufnet",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return listener.DialContext(ctx)
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	return types.NewEventStreamClient(conn)
}

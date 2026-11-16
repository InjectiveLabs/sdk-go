package main

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kmsg"

	"github.com/InjectiveLabs/sdk-go/chain/mq/types"
	"github.com/InjectiveLabs/sdk-go/client/mq/kafkaconfig"
)

func TestKafkaDetectorTLSWatermark(t *testing.T) {
	brokers := os.Getenv("MQ_TEST_KAFKA_BROKERS")
	if brokers == "" {
		t.Skip("set MQ_TEST_KAFKA_BROKERS to test a disposable TLS/SCRAM Kafka fixture")
	}
	for _, name := range []string{"MQ_KAFKA_USERNAME", "MQ_KAFKA_PASSWORD", "MQ_KAFKA_CA_FILE", "MQ_TEST_WRITER_USERNAME", "MQ_TEST_WRITER_PASSWORD"} {
		if os.Getenv(name) == "" {
			t.Fatalf("%s is required for the Kafka integration test", name)
		}
	}
	opts, err := kafkaconfig.OptionsFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	observerOptions := append([]kgo.Opt{}, opts...)
	observerOptions = append(observerOptions, kgo.SeedBrokers(strings.Split(brokers, ",")...))
	observer, err := kgo.NewClient(observerOptions...)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(observer.Close)
	opts = kafkaOptionsForCredentials(t, os.Getenv("MQ_TEST_WRITER_USERNAME"), os.Getenv("MQ_TEST_WRITER_PASSWORD"))
	opts = append(opts, kgo.SeedBrokers(strings.Split(brokers, ",")...))
	writer, err := kgo.NewClient(opts...)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(writer.Close)
	height := time.Now().UnixNano()
	value, err := (&types.EventStreamResponse{BlockHeight: height}).Marshal()
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	records := []*kgo.Record{
		{Topic: "latest", Value: value},
		{Topic: "unrefined_block_results", Value: value},
	}
	if err := writer.ProduceSync(ctx, records...).FirstErr(); err != nil {
		t.Fatalf("seed raw and latest records: %v", err)
	}
	cfg := mqDetectorConfig{
		KafkaBrokers: strings.Split(brokers, ","), ConsumerID: "quants-gap-integration",
		RawTopic: "unrefined_block_results", LatestTopic: "latest",
		FullNodes: []string{"http://127.0.0.1:1"}, RequestTimeout: time.Second,
		MessageTimeout: time.Hour, RetryBackoff: time.Second,
		StateFile: filepath.Join(t.TempDir(), "detector.json"),
	}
	done := make(chan error, 1)
	go func() { done <- startMQDetector(ctx, cfg) }()
	t.Cleanup(func() {
		cancel()
		select {
		case err := <-done:
			if !errors.Is(err, context.Canceled) && !errors.Is(err, context.DeadlineExceeded) {
				t.Errorf("detector stopped unexpectedly: %v", err)
			}
		case <-time.After(5 * time.Second):
			t.Error("detector did not stop")
		}
	})
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			t.Fatal("authenticated detector did not persist the latest watermark")
		case <-ticker.C:
			state, err := loadDetectorState(cfg)
			if err != nil {
				t.Fatal(err)
			}
			if state.LatestHeight >= height && detectorOffsetsCommitted(t, ctx, observer, cfg.ConsumerID, records) {
				return
			}
		}
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

func detectorOffsetsCommitted(t *testing.T, ctx context.Context, observer *kgo.Client, group string, records []*kgo.Record) bool {
	t.Helper()
	req := kmsg.NewPtrOffsetFetchRequest()
	reqGroup := kmsg.NewOffsetFetchRequestGroup()
	reqGroup.Group = group
	for _, record := range records {
		reqGroup.Topics = append(reqGroup.Topics, kmsg.OffsetFetchRequestGroupTopic{Topic: record.Topic, Partitions: []int32{record.Partition}})
	}
	req.Groups = []kmsg.OffsetFetchRequestGroup{reqGroup}
	response, err := req.RequestWith(ctx, observer)
	if err != nil {
		t.Fatalf("read detector committed offsets: %v", err)
	}
	committed := 0
	for _, result := range response.Groups {
		if result.ErrorCode != 0 {
			t.Fatalf("read detector group offsets: Kafka error code %d", result.ErrorCode)
		}
		for _, topic := range result.Topics {
			for _, partition := range topic.Partitions {
				for _, record := range records {
					if topic.Topic == record.Topic && partition.Partition == record.Partition && partition.ErrorCode == 0 && partition.Offset > record.Offset {
						committed++
					}
				}
			}
		}
	}
	return committed == len(records)
}

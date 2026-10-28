package main

import (
	"testing"
	"time"
)

func TestValidateRequiresConsumerID(t *testing.T) {
	cfg := mqDetectorConfig{
		KafkaBrokers:   []string{"broker-1:9092"},
		ConsumerID:     " ",
		RawTopic:       "raw",
		LatestTopic:    "latest",
		FullNodes:      []string{"http://node-1:9999"},
		RequestTimeout: time.Second,
		MessageTimeout: time.Second,
	}

	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected Validate to reject empty consumer id")
	}

	if err.Error() != "invalid MQ detector config: consumer id cannot be empty" {
		t.Fatalf("unexpected error: %v", err)
	}
}

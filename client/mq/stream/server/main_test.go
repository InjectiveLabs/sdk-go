package main

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestStartMQStreamRejectsPartialAuthentication(t *testing.T) {
	t.Setenv("MQ_KAFKA_USERNAME", "reader")
	t.Setenv("MQ_KAFKA_PASSWORD", "")
	t.Setenv("MQ_KAFKA_CA_FILE", "")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := startMQStream(ctx, mqStreamConfig{
		ListenAddress: "127.0.0.1:0", KafkaBrokers: []string{"unused:9093"},
		MinClientPingInterval: time.Second, MaxConnectionIdle: time.Second,
		ServerPingInterval: time.Second, ServerPingResponseTimeout: time.Second,
	})
	if err == nil || !strings.Contains(err.Error(), "must be set together") {
		t.Fatalf("stream startup accepted incomplete Kafka authentication: %v", err)
	}
}

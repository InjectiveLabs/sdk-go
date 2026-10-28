package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	flagMQDetectorKafkaBrokers   = "kafka-brokers"
	flagMQDetectorConsumerID     = "consumer-id"
	flagMQDetectorRawTopic       = "raw-topic"
	flagMQDetectorLatestTopic    = "latest-topic"
	flagMQDetectorFullNodes      = "full-nodes"
	flagMQDetectorControlToken   = "control-token"
	flagMQDetectorRequestTimeout = "request-timeout"
	flagMQDetectorMessageTimeout = "message-timeout"
	flagMQDetectorStateFile      = "state-file"
	flagMQDetectorRetryBackoff   = "retry-backoff"
)

type mqDetectorConfig struct {
	KafkaBrokers   []string
	ConsumerID     string
	RawTopic       string
	LatestTopic    string
	FullNodes      []string
	ControlToken   string
	RequestTimeout time.Duration
	MessageTimeout time.Duration
	StateFile      string
	RetryBackoff   time.Duration
}

func (cfg mqDetectorConfig) Validate() error {
	if len(cfg.KafkaBrokers) == 0 {
		return errors.New("invalid MQ detector config: no Kafka brokers specified")
	}

	if strings.TrimSpace(cfg.ConsumerID) == "" {
		return errors.New("invalid MQ detector config: consumer id cannot be empty")
	}

	for i, broker := range cfg.KafkaBrokers {
		if strings.TrimSpace(broker) == "" {
			return fmt.Errorf("invalid MQ detector config: Kafka broker #%d is empty", i+1)
		}
	}

	if strings.TrimSpace(cfg.LatestTopic) == "" {
		return errors.New("invalid MQ detector config: latest topic cannot be empty")
	}

	if strings.TrimSpace(cfg.RawTopic) == "" {
		return errors.New("invalid MQ detector config: raw topic cannot be empty")
	}

	if cfg.RawTopic == cfg.LatestTopic {
		return errors.New("invalid MQ detector config: raw topic and latest topic must be different")
	}

	if len(cfg.FullNodes) == 0 {
		return errors.New("invalid MQ detector config: no full nodes specified")
	}

	for i, node := range cfg.FullNodes {
		if strings.TrimSpace(node) == "" {
			return fmt.Errorf("invalid MQ detector config: full node URL #%d is empty", i+1)
		}
	}

	if cfg.RequestTimeout <= 0 {
		return errors.New("invalid MQ detector config: request timeout must be positive")
	}

	if cfg.MessageTimeout <= 0 {
		return errors.New("invalid MQ detector config: message timeout must be positive")
	}
	if strings.TrimSpace(cfg.StateFile) == "" {
		return errors.New("invalid MQ detector config: state file cannot be empty")
	}
	if cfg.RetryBackoff <= 0 {
		return errors.New("invalid MQ detector config: retry backoff must be positive")
	}

	return nil
}

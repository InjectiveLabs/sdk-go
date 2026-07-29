package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

const (
	flagMQDetectorKafkaBrokers   = "kafka-brokers"
	flagMQDetectorRawTopic       = "raw-topic"
	flagMQDetectorLatestTopic    = "latest-topic"
	flagMQDetectorFullNodes      = "full-nodes"
	flagMQDetectorControlToken   = "control-token"
	flagMQDetectorRequestTimeout = "request-timeout"
	flagMQDetectorMessageTimeout = "message-timeout"
)

type mqDetectorConfig struct {
	KafkaBrokers   []string
	RawTopic       string
	LatestTopic    string
	FullNodes      []string
	ControlToken   string
	RequestTimeout time.Duration
	MessageTimeout time.Duration
}

func parseConfig() (mqDetectorConfig, error) {
	kafkaBrokers := flag.String(flagMQDetectorKafkaBrokers, "", "Comma-separated Kafka broker addresses")
	rawTopic := flag.String(flagMQDetectorRawTopic, "", "Topic name for raw messages")
	latestTopic := flag.String(flagMQDetectorLatestTopic, "", "Topic name for latest messages")
	fullNodes := flag.String(flagMQDetectorFullNodes, "", "Comma-separated full node control plane URLs")
	controlToken := flag.String(flagMQDetectorControlToken, "", "Bearer token for full node control plane requests")
	requestTimeout := flag.Duration(flagMQDetectorRequestTimeout, 10*time.Second, "Timeout for block requests")
	messageTimeout := flag.Duration(flagMQDetectorMessageTimeout, 30*time.Second, "Message waiting timeout duration")
	flag.Parse()

	cfg := mqDetectorConfig{
		RawTopic:       *rawTopic,
		LatestTopic:    *latestTopic,
		ControlToken:   *controlToken,
		RequestTimeout: *requestTimeout,
		MessageTimeout: *messageTimeout,
	}

	if *kafkaBrokers != "" {
		cfg.KafkaBrokers = strings.Split(*kafkaBrokers, ",")
	}

	if *fullNodes != "" {
		cfg.FullNodes = strings.Split(*fullNodes, ",")
	}

	if err := cfg.Validate(); err != nil {
		return mqDetectorConfig{}, err
	}

	return cfg, nil
}

func (cfg mqDetectorConfig) Validate() error {
	if len(cfg.KafkaBrokers) == 0 {
		return errors.New("invalid MQ detector config: no Kafka brokers specified")
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

	return nil
}

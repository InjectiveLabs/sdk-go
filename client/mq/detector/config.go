package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	flagMQDetectorKafkaBrokers   = "kafka-brokers"
	flagMQDetectorRawTopic       = "raw-topic"
	flagMQDetectorLatestTopic    = "latest-topic"
	flagMQDetectorFullNodes      = "full-nodes"
	flagMQDetectorRequestTimeout = "request-timeout"
	flagMQDetectorMessageTimeout = "message-timeout"
)

type mqDetectorConfig struct {
	KafkaBrokers   []string
	RawTopic       string
	LatestTopic    string
	FullNodes      []string
	RequestTimeout time.Duration
	MessageTimeout time.Duration
}

func parseConfig(cmd *cobra.Command) (mqDetectorConfig, error) {
	kafkaBrokers, err := cmd.Flags().GetStringSlice(flagMQDetectorKafkaBrokers)
	if err != nil {
		return mqDetectorConfig{}, err
	}

	rawTopic, err := cmd.Flags().GetString(flagMQDetectorRawTopic)
	if err != nil {
		return mqDetectorConfig{}, err
	}

	latestTopic, err := cmd.Flags().GetString(flagMQDetectorLatestTopic)
	if err != nil {
		return mqDetectorConfig{}, err
	}

	fullNodes, err := cmd.Flags().GetStringSlice(flagMQDetectorFullNodes)
	if err != nil {
		return mqDetectorConfig{}, err
	}

	requestTimeout, err := cmd.Flags().GetDuration(flagMQDetectorRequestTimeout)
	if err != nil {
		return mqDetectorConfig{}, err
	}

	messageTimeout, err := cmd.Flags().GetDuration(flagMQDetectorMessageTimeout)
	if err != nil {
		return mqDetectorConfig{}, err
	}

	cfg := mqDetectorConfig{
		KafkaBrokers:   kafkaBrokers,
		RawTopic:       rawTopic,
		LatestTopic:    latestTopic,
		FullNodes:      fullNodes,
		RequestTimeout: requestTimeout,
		MessageTimeout: messageTimeout,
	}

	if len(cfg.KafkaBrokers) == 0 {
		return mqDetectorConfig{}, errors.New("invalid MQ detector config: no Kafka brokers specified")
	}

	for i, broker := range cfg.KafkaBrokers {
		if strings.TrimSpace(broker) == "" {
			return mqDetectorConfig{}, fmt.Errorf("invalid MQ detector config: Kafka broker #%d is empty", i+1)
		}
	}

	if strings.TrimSpace(cfg.LatestTopic) == "" {
		return mqDetectorConfig{}, errors.New("invalid MQ detector config: latest topic cannot be empty")
	}

	if strings.TrimSpace(cfg.RawTopic) == "" {
		return mqDetectorConfig{}, errors.New("invalid MQ detector config: raw topic cannot be empty")
	}

	if cfg.RawTopic == cfg.LatestTopic {
		return mqDetectorConfig{}, errors.New("invalid MQ detector config: raw topic and latest topic must be different")
	}

	if len(cfg.FullNodes) == 0 {
		return mqDetectorConfig{}, errors.New("invalid MQ detector config: no full nodes specified")
	}

	for i, node := range cfg.FullNodes {
		if strings.TrimSpace(node) == "" {
			return mqDetectorConfig{}, fmt.Errorf("invalid MQ detector config: full node URL #%d is empty", i+1)
		}
	}

	if cfg.RequestTimeout <= 0 {
		return mqDetectorConfig{}, errors.New("invalid MQ detector config: request timeout must be positive")
	}

	if cfg.MessageTimeout <= 0 {
		return mqDetectorConfig{}, errors.New("invalid MQ detector config: message timeout must be positive")
	}

	return cfg, nil
}

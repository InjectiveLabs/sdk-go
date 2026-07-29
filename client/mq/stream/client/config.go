package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const (
	flagMQStreamClientAddress        = "address"
	flagMQStreamClientConsumerID     = "consumer-id"
	flagMQStreamClientConsumerIDFile = "consumer-id-file"
	flagMQStreamClientTopic          = "topic"
	flagMQStreamClientFormat         = "format"
	flagMQStreamClientEventsDir      = "events-dir"

	defaultConsumerIDFile = ".mq-stream-consumer-id"
)

type mqStreamClientConfig struct {
	Address    string
	ConsumerID string
	Topic      string
	Format     string
	EventsDir  string
}

func parseConfig() (mqStreamClientConfig, error) {
	address := flag.String(flagMQStreamClientAddress, "localhost:9988", "MQ gRPC stream server address")
	consumerID := flag.String(flagMQStreamClientConsumerID, "", "Stable Kafka consumer id")
	consumerIDFile := flag.String(flagMQStreamClientConsumerIDFile, defaultConsumerIDFile, "File used to persist a generated Kafka consumer id")
	topic := flag.String(flagMQStreamClientTopic, "", "Topic to consume")
	format := flag.String(flagMQStreamClientFormat, "verbose", "Output format: verbose or minimal")
	eventsDir := flag.String(flagMQStreamClientEventsDir, "", "Directory to write one JSON file per streamed block")
	flag.Parse()

	resolvedConsumerID, err := resolveConsumerID(*consumerID, *consumerIDFile)
	if err != nil {
		return mqStreamClientConfig{}, err
	}

	cfg := mqStreamClientConfig{
		Address:    *address,
		ConsumerID: resolvedConsumerID,
		Topic:      *topic,
		Format:     *format,
		EventsDir:  *eventsDir,
	}

	if err := cfg.Validate(); err != nil {
		return mqStreamClientConfig{}, err
	}

	return cfg, nil
}

func (cfg mqStreamClientConfig) Validate() error {
	if strings.TrimSpace(cfg.Address) == "" {
		return errors.New("invalid MQ stream client config: address cannot be empty")
	}

	if strings.TrimSpace(cfg.ConsumerID) == "" {
		return errors.New("invalid MQ stream client config: consumer id cannot be empty")
	}

	if strings.TrimSpace(cfg.Topic) == "" {
		return errors.New("invalid MQ stream client config: topic cannot be empty")
	}

	switch cfg.Format {
	case "minimal", "verbose":
		return nil
	default:
		return fmt.Errorf("invalid MQ stream client config: unsupported format %q", cfg.Format)
	}
}

func resolveConsumerID(flagValue, filePath string) (string, error) {
	if consumerID := strings.TrimSpace(flagValue); consumerID != "" {
		return consumerID, nil
	}

	filePath = strings.TrimSpace(filePath)
	if filePath == "" {
		return uuid.NewString(), nil
	}

	bz, err := os.ReadFile(filePath)
	if err == nil {
		consumerID := strings.TrimSpace(string(bz))
		if consumerID == "" {
			return "", fmt.Errorf("consumer id file %q is empty", filePath)
		}

		return consumerID, nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("read consumer id file: %w", err)
	}

	consumerID := uuid.NewString()
	if err := writeConsumerIDFile(filePath, consumerID); err != nil {
		return "", err
	}

	return consumerID, nil
}

func writeConsumerIDFile(filePath, consumerID string) error {
	dir := filepath.Dir(filePath)
	if dir != "." {
		if err := os.MkdirAll(dir, 0o750); err != nil {
			return fmt.Errorf("create consumer id file dir: %w", err)
		}
	}

	if err := os.WriteFile(filePath, []byte(consumerID+"\n"), 0o600); err != nil {
		return fmt.Errorf("write consumer id file: %w", err)
	}

	return nil
}

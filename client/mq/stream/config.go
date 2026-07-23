package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	flagMQStreamListenAddress             = "listen-address"
	flagMQStreamKafkaBrokers              = "kafka-brokers"
	flagMQStreamTopicName                 = "topic"
	flagMQStreamEnforceKeepalive          = "mq-enforce-keepalive"
	flagMQStreamMinClientPingInterval     = "mq-min-client-ping-interval"
	flagMQStreamMaxConnectionIdle         = "mq-max-connection-idle"
	flagMQStreamServerPingInterval        = "mq-server-ping-interval"
	flagMQStreamServerPingResponseTimeout = "mq-server-ping-response-timeout"

	defaultMQStreamMinClientPingInterval     = 30 * time.Second
	defaultMQStreamMaxConnectionIdle         = 180 * time.Second
	defaultMQStreamServerPingInterval        = 60 * time.Second
	defaultMQStreamServerPingResponseTimeout = 40 * time.Second
)

type mqStreamConfig struct {
	ListenAddress             string
	KafkaBrokers              []string
	Topic                     string
	EnforceKeepalive          bool
	MinClientPingInterval     time.Duration
	MaxConnectionIdle         time.Duration
	ServerPingInterval        time.Duration
	ServerPingResponseTimeout time.Duration
}

func parseConfig(cmd *cobra.Command) (mqStreamConfig, error) {
	listenAddress, err := cmd.Flags().GetString(flagMQStreamListenAddress)
	if err != nil {
		return mqStreamConfig{}, err
	}

	kafkaBrokers, err := cmd.Flags().GetStringSlice(flagMQStreamKafkaBrokers)
	if err != nil {
		return mqStreamConfig{}, err
	}

	topic, err := cmd.Flags().GetString(flagMQStreamTopicName)
	if err != nil {
		return mqStreamConfig{}, err
	}

	enforceKeepalive, err := cmd.Flags().GetBool(flagMQStreamEnforceKeepalive)
	if err != nil {
		return mqStreamConfig{}, err
	}

	minClientPingInterval, err := cmd.Flags().GetDuration(flagMQStreamMinClientPingInterval)
	if err != nil {
		return mqStreamConfig{}, err
	}

	maxConnectionIdle, err := cmd.Flags().GetDuration(flagMQStreamMaxConnectionIdle)
	if err != nil {
		return mqStreamConfig{}, err
	}

	serverPingInterval, err := cmd.Flags().GetDuration(flagMQStreamServerPingInterval)
	if err != nil {
		return mqStreamConfig{}, err
	}

	serverPingResponseTimeout, err := cmd.Flags().GetDuration(flagMQStreamServerPingResponseTimeout)
	if err != nil {
		return mqStreamConfig{}, err
	}

	cfg := mqStreamConfig{
		ListenAddress:             listenAddress,
		KafkaBrokers:              kafkaBrokers,
		Topic:                     topic,
		EnforceKeepalive:          enforceKeepalive,
		MinClientPingInterval:     minClientPingInterval,
		MaxConnectionIdle:         maxConnectionIdle,
		ServerPingInterval:        serverPingInterval,
		ServerPingResponseTimeout: serverPingResponseTimeout,
	}

	if strings.TrimSpace(cfg.ListenAddress) == "" {
		return mqStreamConfig{}, errors.New("invalid MQ stream config: listen address cannot be empty")
	}

	if len(cfg.KafkaBrokers) == 0 {
		return mqStreamConfig{}, errors.New("invalid MQ stream config: no Kafka brokers specified")
	}

	for i, broker := range cfg.KafkaBrokers {
		if strings.TrimSpace(broker) == "" {
			return mqStreamConfig{}, fmt.Errorf("invalid MQ stream config: Kafka broker #%d is empty", i+1)
		}
	}

	if strings.TrimSpace(cfg.Topic) == "" {
		return mqStreamConfig{}, errors.New("invalid MQ stream config: topic cannot be empty")
	}

	if cfg.MinClientPingInterval <= 0 {
		return mqStreamConfig{}, errors.New("invalid MQ stream config: min client ping interval must be positive")
	}

	if cfg.MaxConnectionIdle <= 0 {
		return mqStreamConfig{}, errors.New("invalid MQ stream config: max connection idle must be positive")
	}

	if cfg.ServerPingInterval <= 0 {
		return mqStreamConfig{}, errors.New("invalid MQ stream config: server ping interval must be positive")
	}

	if cfg.ServerPingResponseTimeout <= 0 {
		return mqStreamConfig{}, errors.New("invalid MQ stream config: server ping response timeout must be positive")
	}

	return cfg, nil
}

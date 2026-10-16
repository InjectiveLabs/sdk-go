package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	flagMQStreamListenAddress             = "listen-address"
	flagMQStreamKafkaBrokers              = "kafka-brokers"
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
	EnforceKeepalive          bool
	MinClientPingInterval     time.Duration
	MaxConnectionIdle         time.Duration
	ServerPingInterval        time.Duration
	ServerPingResponseTimeout time.Duration
}

func (cfg mqStreamConfig) Validate() error {
	if strings.TrimSpace(cfg.ListenAddress) == "" {
		return errors.New("invalid MQ stream config: listen address cannot be empty")
	}

	if len(cfg.KafkaBrokers) == 0 {
		return errors.New("invalid MQ stream config: no Kafka brokers specified")
	}

	for i, broker := range cfg.KafkaBrokers {
		if strings.TrimSpace(broker) == "" {
			return fmt.Errorf("invalid MQ stream config: Kafka broker #%d is empty", i+1)
		}
	}

	if cfg.MinClientPingInterval <= 0 {
		return errors.New("invalid MQ stream config: min client ping interval must be positive")
	}

	if cfg.MaxConnectionIdle <= 0 {
		return errors.New("invalid MQ stream config: max connection idle must be positive")
	}

	if cfg.ServerPingInterval <= 0 {
		return errors.New("invalid MQ stream config: server ping interval must be positive")
	}

	if cfg.ServerPingResponseTimeout <= 0 {
		return errors.New("invalid MQ stream config: server ping response timeout must be positive")
	}

	return nil
}

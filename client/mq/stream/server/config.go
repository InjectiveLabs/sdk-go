package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
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

func parseConfig() (mqStreamConfig, error) {
	listenAddress := flag.String(flagMQStreamListenAddress, "0.0.0.0:9988", "Listen address")
	kafkaBrokers := flag.String(flagMQStreamKafkaBrokers, "", "Comma-separated Kafka broker addresses")
	topic := flag.String(flagMQStreamTopicName, "", "Topic to consume")
	enforceKeepalive := flag.Bool(flagMQStreamEnforceKeepalive, false,
		"Define if Keepalive configuration params should be applied to MQ event stream gRPC server",
	)
	minClientPingInterval := flag.Duration(flagMQStreamMinClientPingInterval, defaultMQStreamMinClientPingInterval,
		"Duration a client should wait before sending a keepalive ping",
	)
	maxConnectionIdle := flag.Duration(flagMQStreamMaxConnectionIdle, defaultMQStreamMaxConnectionIdle,
		"Duration a connection is allowed to stay idle before forcing the disconnection",
	)
	serverPingInterval := flag.Duration(flagMQStreamServerPingInterval, defaultMQStreamServerPingInterval,
		"Duration after which the server will send a keepalive ping to the client on an idle connection",
	)
	serverPingResponseTimeout := flag.Duration(flagMQStreamServerPingResponseTimeout, defaultMQStreamServerPingResponseTimeout,
		"Duration the server waits for the client to respond to a ping message before forcing a disconnection",
	)

	flag.Parse()

	cfg := mqStreamConfig{
		ListenAddress:             *listenAddress,
		Topic:                     *topic,
		EnforceKeepalive:          *enforceKeepalive,
		MinClientPingInterval:     *minClientPingInterval,
		MaxConnectionIdle:         *maxConnectionIdle,
		ServerPingInterval:        *serverPingInterval,
		ServerPingResponseTimeout: *serverPingResponseTimeout,
	}

	if *kafkaBrokers != "" {
		cfg.KafkaBrokers = strings.Split(*kafkaBrokers, ",")
	}

	if err := cfg.Validate(); err != nil {
		return mqStreamConfig{}, err
	}

	return cfg, nil
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

	if strings.TrimSpace(cfg.Topic) == "" {
		return errors.New("invalid MQ stream config: topic cannot be empty")
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

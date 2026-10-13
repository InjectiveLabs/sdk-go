package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	sdklog "cosmossdk.io/log"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	var (
		kafkaBrokers   = flag.String(flagMQDetectorKafkaBrokers, "", "Comma-separated Kafka broker addresses")
		consumerID     = flag.String(flagMQDetectorConsumerID, "", "Stable, unique Kafka consumer id (required)")
		rawTopic       = flag.String(flagMQDetectorRawTopic, "", "Topic name for raw messages")
		latestTopic    = flag.String(flagMQDetectorLatestTopic, "", "Topic name for latest messages")
		fullNodes      = flag.String(flagMQDetectorFullNodes, "", "Comma-separated full node control plane URLs")
		controlToken   = flag.String(flagMQDetectorControlToken, "", "Bearer token for full node control plane requests")
		requestTimeout = flag.Duration(flagMQDetectorRequestTimeout, 10*time.Second, "Timeout for block requests")
		messageTimeout = flag.Duration(flagMQDetectorMessageTimeout, 30*time.Second, "Latest progress timeout before replay")
		stateFile      = flag.String(flagMQDetectorStateFile, ".mq-detector-state.json", "Persistent detector watermark file")
		retryBackoff   = flag.Duration(flagMQDetectorRetryBackoff, 5*time.Second, "Minimum delay between Kafka or replay retries")
	)

	flag.Parse()

	cfg := mqDetectorConfig{
		ConsumerID:     strings.TrimSpace(*consumerID),
		RawTopic:       *rawTopic,
		LatestTopic:    *latestTopic,
		ControlToken:   *controlToken,
		RequestTimeout: *requestTimeout,
		MessageTimeout: *messageTimeout,
		StateFile:      *stateFile,
		RetryBackoff:   *retryBackoff,
	}

	if *kafkaBrokers != "" {
		cfg.KafkaBrokers = strings.Split(*kafkaBrokers, ",")
	}

	if *fullNodes != "" {
		cfg.FullNodes = strings.Split(*fullNodes, ",")
	}

	if err := startMQDetector(context.Background(), cfg); err != nil && !errors.Is(err, context.Canceled) {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startMQDetector(ctx context.Context, cfg mqDetectorConfig) error {
	if err := cfg.Validate(); err != nil {
		return err
	}

	state, err := loadDetectorState(cfg)
	if err != nil {
		return err
	}
	logger := sdklog.NewLogger(os.Stderr)
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.KafkaBrokers...),
		kgo.ConsumerGroup(cfg.ConsumerID),
		kgo.ConsumeTopics(cfg.RawTopic, cfg.LatestTopic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
		// The detector saves its watermark before marking a latest record. A
		// restart can therefore never resume beyond the last durable watermark.
		kgo.AutoCommitMarks(),
	)
	if err != nil {
		return err
	}
	defer client.Close()

	signalCtx, stopSignals := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stopSignals()
	return runDetector(signalCtx, cfg, state, client, logger)
}

func requestControlPlaneBlocks(
	ctx context.Context,
	client *http.Client,
	startHeight uint64,
	nodeURL,
	controlToken string,
) error {
	requestURL := fmt.Sprintf("%s/request?from_height=%d", nodeURL, startHeight)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, http.NoBody)
	if err != nil {
		return err
	}

	if controlToken != "" {
		req.Header.Set("Authorization", "Bearer "+controlToken)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

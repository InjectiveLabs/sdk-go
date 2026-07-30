package main

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"syscall"
	"time"

	sdklog "cosmossdk.io/log"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

func main() {
	if err := startMQDetector(context.Background()); err != nil && !errors.Is(err, context.Canceled) {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startMQDetector(ctx context.Context) error {
	logger := sdklog.NewLogger(os.Stderr)
	config, err := parseConfig()
	if err != nil {
		return err
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(config.KafkaBrokers...),
		kgo.ConsumerGroup(config.ConsumerID),
		kgo.ConsumeTopics(config.RawTopic, config.LatestTopic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return err
	}

	defer client.Close()

	signalCtx, stopSignals := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stopSignals()

	ctx, cancel := context.WithCancelCause(signalCtx)
	defer cancel(nil)

	var (
		rawTopicCh    = make(chan *types.EventStreamResponse, 1)
		latestTopicCh = make(chan *types.EventStreamResponse, 1)
	)

	poll := func(ctx context.Context, raw, latest chan<- *types.EventStreamResponse) error {
		for {
			fetches := client.PollFetches(ctx)
			if err := fetches.Err(); err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) {
					return err
				}

				err = fmt.Errorf("poll kafka fetches: %w", err)
				logger.Error("error polling fetches, stopping consumer loop", "err", err.Error())
				return err
			}

			iter := fetches.RecordIter()
			for !iter.Done() {
				record := iter.Next()
				var msg types.EventStreamResponse
				if err := msg.Unmarshal(record.Value); err != nil {
					logger.Error("error decoding event message", "err", err.Error())
					continue
				}

				switch record.Topic {
				case config.LatestTopic:
					latest <- &msg
				case config.RawTopic:
					raw <- &msg
				default:
					err := fmt.Errorf("unknown kafka topic %q", record.Topic)
					logger.Error("unknown topic", "topic", record.Topic)
					return err
				}
			}
		}
	}

	controlPlaneClient := &http.Client{Timeout: config.RequestTimeout}
	callControlPlane := func(startHeight uint64) {
		logger.Info("calling control plane", "startHeight", startHeight)
		for _, node := range config.FullNodes {
			if err := requestControlPlaneBlocks(ctx, controlPlaneClient, startHeight, node, config.ControlToken); err != nil {
				logger.Error("error requesting blocks from node", "node", node, "err", err.Error())
			}

			logger.Debug("requested blocks from node", "node", node)
		}
	}

	// start the kafka polling loop
	go func() {
		defer func() {
			close(rawTopicCh)
			close(latestTopicCh)
		}()

		if err := poll(ctx, rawTopicCh, latestTopicCh); err != nil {
			if !errors.Is(err, context.Canceled) && !errors.Is(ctx.Err(), context.Canceled) {
				cancel(err)
				return
			}
		}

		cancel(nil)
	}()

	var (
		latestHeight   = int64(0)
		rawHeights     = make(map[int64]struct{})
		lastSeenLatest time.Time
	)

	for {
		if noLatestForAWhile := !lastSeenLatest.IsZero() && time.Since(lastSeenLatest) > config.MessageTimeout; noLatestForAWhile {
			go callControlPlane(uint64(latestHeight) + 1)
		}

		select {
		case <-ctx.Done():
			if err := context.Cause(ctx); err != nil {
				return err
			}

			return ctx.Err()
		case msg, ok := <-rawTopicCh:
			if msg == nil || !ok {
				if err := context.Cause(ctx); err != nil && !errors.Is(err, context.Canceled) {
					return err
				}

				return nil
			}

			if msg.BlockHeight <= latestHeight {
				continue // no need to store previous heights
			}

			if _, ok := rawHeights[msg.BlockHeight]; ok {
				continue // already here
			}

			rawHeights[msg.BlockHeight] = struct{}{}
			logger.Info("received raw event",
				"height", msg.BlockHeight,
				"app_hash", hex.EncodeToString(msg.AppHash),
				"last_app_hash", hex.EncodeToString(msg.LastAppHash),
			)
		case msg, ok := <-latestTopicCh:
			if msg == nil || !ok {
				if err := context.Cause(ctx); err != nil && !errors.Is(err, context.Canceled) {
					return err
				}

				return nil
			}

			latestHeight = msg.BlockHeight
			lastSeenLatest = time.Now()
			logger.Info("received latest event",
				"height", msg.BlockHeight,
				"app_hash", hex.EncodeToString(msg.AppHash),
				"last_app_hash", hex.EncodeToString(msg.LastAppHash),
			)

			for h := range rawHeights {
				if h <= latestHeight {
					delete(rawHeights, h)
				}
			}
		}

		if nothingYet := latestHeight == 0 && len(rawHeights) == 0; nothingYet {
			continue
		}

		if needLatestHeightFirst := latestHeight == 0 && len(rawHeights) != 0; needLatestHeightFirst {
			continue
		}

		if needSomeRawHeightAlso := latestHeight != 0 && len(rawHeights) == 0; needSomeRawHeightAlso {
			continue
		}

		sortedHeights := make([]int64, 0, len(rawHeights))
		for h := range rawHeights {
			sortedHeights = append(sortedHeights, h)
		}

		slices.Sort(sortedHeights)

		if thereIsAGap := latestHeight+1 < sortedHeights[0]; thereIsAGap {
			logger.Warn("height gap detected", "want", latestHeight+1, "got", sortedHeights[0])
			go callControlPlane(uint64(latestHeight) + 1)
		}
	}
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
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}

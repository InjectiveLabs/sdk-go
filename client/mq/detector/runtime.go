package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	sdklog "cosmossdk.io/log"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

type detectorConsumer interface {
	PollFetches(context.Context) kgo.Fetches
	MarkCommitRecords(...*kgo.Record)
}

type detectorEvent struct {
	height int64
	ack    chan struct{}
}

func runDetector(ctx context.Context, cfg mqDetectorConfig, state detectorState, consumer detectorConsumer, logger sdklog.Logger) error {
	ctx, cancel := context.WithCancel(ctx)
	var workers sync.WaitGroup
	defer func() {
		cancel()
		workers.Wait()
	}()

	events := make(chan detectorEvent)
	workers.Go(func() { pollDetector(ctx, cfg, consumer, events, logger) })
	replayDone := make(chan error, 1)
	client := &http.Client{Timeout: cfg.RequestTimeout}
	lastProgress := time.Now()
	nextAttempt := lastProgress.Add(cfg.MessageTimeout)
	replayInFlight := false
	// A real timer must wake the detector even when neither topic has records.
	timer := time.NewTimer(cfg.MessageTimeout)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event := <-events:
			if event.height > state.LatestHeight {
				state.LatestHeight = event.height
				if err := saveDetectorState(cfg.StateFile, state); err != nil {
					return fmt.Errorf("persist detector watermark: %w", err)
				}
				lastProgress = time.Now()
				logger.Debug("latest watermark advanced", "height", state.LatestHeight)
			}
			close(event.ack)
		case err := <-replayDone:
			replayInFlight = false
			nextAttempt = time.Now().Add(cfg.RetryBackoff)
			if err != nil && ctx.Err() == nil {
				logger.Error("replay request failed; will retry", "error", err)
			}
		case <-timer.C:
		}

		now := time.Now()
		if !replayInFlight && !now.Before(nextAttempt) && !now.Before(lastProgress.Add(cfg.MessageTimeout)) {
			replayInFlight = true
			// Zero means no baseline, not a request for height zero. The worker
			// discovers the earliest retained block using the node control API.
			fromHeight := int64(0)
			if state.LatestHeight > 0 {
				fromHeight = state.LatestHeight + 1
			}
			workers.Go(func() { replayDone <- requestReplay(ctx, client, cfg, fromHeight) })
		}

		wakeAt := maxTime(nextAttempt, lastProgress.Add(cfg.MessageTimeout))
		if replayInFlight {
			wakeAt = now.Add(cfg.RetryBackoff)
		}
		// Stopping and draining also supports Go timer behavior before 1.23.
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(max(time.Until(wakeAt), time.Millisecond))
	}
}

func maxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}

// pollDetector never abandons the group on a broker outage. Successful records
// in a partial fetch are handled even when other partitions report errors.
func pollDetector(ctx context.Context, cfg mqDetectorConfig, consumer detectorConsumer, events chan<- detectorEvent, logger sdklog.Logger) {
	for ctx.Err() == nil {
		fetches := consumer.PollFetches(ctx)
		for _, fetchErr := range fetches.Errors() {
			if ctx.Err() != nil {
				return
			}
			logger.Error("Kafka fetch failed; will retry", "topic", fetchErr.Topic, "error", fetchErr.Err)
		}
		for _, record := range fetches.Records() {
			if record.Topic != cfg.LatestTopic {
				// Raw heights need no map: lack of latest progress is enough to
				// request the entire missing suffix. Memory remains bounded.
				consumer.MarkCommitRecords(record)
				continue
			}
			var message types.EventStreamResponse
			if err := message.Unmarshal(record.Value); err != nil || message.BlockHeight <= 0 {
				logger.Error("invalid latest event", "offset", record.Offset, "error", err)
				consumer.MarkCommitRecords(record)
				continue
			}
			event := detectorEvent{height: message.BlockHeight, ack: make(chan struct{})}
			select {
			case <-ctx.Done():
				return
			case events <- event:
			}
			select {
			case <-ctx.Done():
				return
			case <-event.ack:
				consumer.MarkCommitRecords(record)
			}
		}
		if len(fetches.Errors()) > 0 && !waitRetry(ctx, cfg.RetryBackoff) {
			return
		}
	}
}

func waitRetry(ctx context.Context, delay time.Duration) bool {
	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-timer.C:
		return true
	}
}

func requestReplay(ctx context.Context, client *http.Client, cfg mqDetectorConfig, fromHeight int64) error {
	var failures []error
	for _, node := range cfg.FullNodes {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		first, last, err := retainedHeights(ctx, client, node, cfg.ControlToken)
		if err != nil {
			failures = append(failures, fmt.Errorf("%s: %w", node, err))
			continue
		}
		start := fromHeight
		if start == 0 {
			start = first
		}
		if start < first || start > last {
			failures = append(failures, fmt.Errorf("%s: requested height %d outside retained range [%d, %d]", node, start, first, last))
			continue
		}
		if err := requestControlPlaneBlocks(ctx, client, uint64(start), node, cfg.ControlToken); err != nil {
			failures = append(failures, fmt.Errorf("%s: %w", node, err))
		}
	}
	return errors.Join(failures...)
}

func retainedHeights(ctx context.Context, client *http.Client, node, token string) (earliest, latest int64, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, node+"/status", http.NoBody)
	if err != nil {
		return 0, 0, err
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("status endpoint returned HTTP %d", resp.StatusCode)
	}
	var status struct {
		EarliestHeight *int64 `json:"earliest_height"`
		LastHeight     int64  `json:"last_height"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, 4096)).Decode(&status); err != nil {
		return 0, 0, err
	}
	if status.EarliestHeight == nil {
		return 0, 0, errors.New("control API lacks earliest_height; upgrade the MQ publisher before enabling recovery")
	}
	if *status.EarliestHeight <= 0 || status.LastHeight < *status.EarliestHeight {
		return 0, 0, errors.New("publisher has no retained blocks yet")
	}
	return *status.EarliestHeight, status.LastHeight, nil
}

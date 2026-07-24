package main

import (
	"context"
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
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

func main() {
	logger := sdklog.NewLogger(os.Stderr)
	cmd := &cobra.Command{
		Use:          "detector",
		Short:        "Starts the MQ detector service",
		SilenceUsage: true,
		Long: `Run the MQ detector service. It consumes raw and latest MQ messages and requests missing blocks
from configured full node control planes when raw/latest state indicates a gap.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return startMQDetector(cmd, logger)
		},
	}

	cmd.Flags().StringSlice(flagMQDetectorKafkaBrokers, []string{}, "Kafka broker addresses")
	cmd.Flags().String(flagMQDetectorRawTopic, "", "Topic name for raw messages")
	cmd.Flags().String(flagMQDetectorLatestTopic, "", "Topic name for latest messages")
	cmd.Flags().StringSlice(flagMQDetectorFullNodes, []string{}, "Full node control plane URLs")
	cmd.Flags().Duration(flagMQDetectorRequestTimeout, 10*time.Second, "Timeout for block requests")
	cmd.Flags().Duration(flagMQDetectorMessageTimeout, 30*time.Second, "Message waiting timeout duration")

	if err := cmd.Execute(); err != nil && !errors.Is(err, context.Canceled) {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startMQDetector(cmd *cobra.Command, logger sdklog.Logger) error {
	config, err := parseConfig(cmd)
	if err != nil {
		return err
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(config.KafkaBrokers...),
		kgo.ConsumerGroup(uuid.New().String()),
		kgo.ConsumeTopics(config.RawTopic, config.LatestTopic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return err
	}

	defer client.Close()

	signalCtx, stopSignals := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
	defer stopSignals()

	ctx, cancel := context.WithCancelCause(signalCtx)
	defer cancel(nil)

	var (
		rawTopicCh    = make(chan *types.EventStreamResponse, 1)
		latestTopicCh = make(chan *types.EventStreamResponse, 1)
	)

	// start the kafka polling loop
	go func() {
		defer func() {
			close(rawTopicCh)
			close(latestTopicCh)
		}()

		for {
			fetches := client.PollFetches(ctx)
			if err := fetches.Err(); err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) {
					cancel(nil)
					return
				}

				err = fmt.Errorf("poll kafka fetches: %w", err)
				logger.Error("error polling fetches, stopping consumer loop", "err", err.Error())
				cancel(err)
				return
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
					latestTopicCh <- &msg
				case config.RawTopic:
					rawTopicCh <- &msg
				default:
					err := fmt.Errorf("unknown kafka topic %q", record.Topic)
					logger.Error("unknown topic", "topic", record.Topic)
					cancel(err)
					return
				}
			}
		}
	}()

	controlPlaneClient := &http.Client{Timeout: config.RequestTimeout}
	requestBlocksFromNode := func(nodeURL string, startHeight uint64) error {
		requestURL := fmt.Sprintf("%s/request?from_height=%d", nodeURL, startHeight)
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, http.NoBody)
		if err != nil {
			return err
		}

		resp, err := controlPlaneClient.Do(req)
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

	callControlPlane := func(startHeight uint64) {
		for _, node := range config.FullNodes {
			if err := requestBlocksFromNode(node, startHeight); err != nil {
				logger.Error("error requesting blocks from node", "node", node, "err", err.Error())
			}

			logger.Info("requested blocks from node", "node", node, "start_height", startHeight)
		}
	}

	var (
		latestHeight  = int64(0)
		rawHeights    = make(map[int64]struct{})
		latestTimeout <-chan time.Time
		latestTimer   *time.Timer
	)

	resetLatestTimer := func() {
		if latestTimer == nil {
			latestTimer = time.NewTimer(config.MessageTimeout)
		} else {
			if !latestTimer.Stop() {
				select {
				case <-latestTimer.C:
				default:
				}
			}
			latestTimer.Reset(config.MessageTimeout)
		}

		latestTimeout = latestTimer.C
	}

	defer func() {
		if latestTimer != nil {
			latestTimer.Stop()
		}
	}()

	for {
		select {
		case <-ctx.Done():
			if err := context.Cause(ctx); err != nil {
				return err
			}

			return ctx.Err()
		case <-latestTimeout:
			go callControlPlane(uint64(latestHeight) + 1)
			resetLatestTimer()
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
		case msg, ok := <-latestTopicCh:
			if msg == nil || !ok {
				if err := context.Cause(ctx); err != nil && !errors.Is(err, context.Canceled) {
					return err
				}

				return nil
			}

			latestHeight = msg.BlockHeight
			resetLatestTimer()

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
			go callControlPlane(uint64(latestHeight) + 1)
		}
	}
}

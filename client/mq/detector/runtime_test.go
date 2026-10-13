package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	sdklog "cosmossdk.io/log"
	"github.com/twmb/franz-go/pkg/kgo"

	"github.com/InjectiveLabs/sdk-go/chain/mq/types"
)

type fakeConsumer struct {
	fetches chan kgo.Fetches
	marked  chan *kgo.Record
	onMark  func()
}

func (consumer *fakeConsumer) PollFetches(ctx context.Context) kgo.Fetches {
	select {
	case <-ctx.Done():
		return nil
	case fetch := <-consumer.fetches:
		return fetch
	}
}

func (consumer *fakeConsumer) MarkCommitRecords(records ...*kgo.Record) {
	if consumer.onMark != nil {
		consumer.onMark()
	}
	for _, record := range records {
		consumer.marked <- record
	}
}

func testConfig(t *testing.T) mqDetectorConfig {
	t.Helper()
	return mqDetectorConfig{
		KafkaBrokers: []string{"unused:9092"}, ConsumerID: "publisher-0-detector", RawTopic: "raw", LatestTopic: "latest",
		StateFile: filepath.Join(t.TempDir(), "detector.json"), MessageTimeout: 30 * time.Millisecond,
		RetryBackoff: 25 * time.Millisecond, RequestTimeout: time.Second,
	}
}

func startTestDetector(t *testing.T, cfg mqDetectorConfig, consumer *fakeConsumer) context.CancelFunc {
	t.Helper()
	state, err := loadDetectorState(cfg)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error, 1)
	go func() { done <- runDetector(ctx, cfg, state, consumer, sdklog.NewNopLogger()) }()
	t.Cleanup(func() {
		cancel()
		select {
		case err := <-done:
			if !errors.Is(err, context.Canceled) {
				t.Errorf("detector stopped unexpectedly: %v", err)
			}
		case <-time.After(5 * time.Second):
			t.Error("detector did not shut down")
		}
	})
	return cancel
}

func newFakeConsumer() *fakeConsumer {
	return &fakeConsumer{fetches: make(chan kgo.Fetches, 32), marked: make(chan *kgo.Record, 128)}
}

func receive[T any](t *testing.T, ch <-chan T) T {
	t.Helper()
	select {
	case value := <-ch:
		return value
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for detector")
		var zero T
		return zero
	}
}

func TestRestartDuringGapReplaysWithoutNewKafkaRecords(t *testing.T) {
	cfg := testConfig(t)
	state := detectorState{ConsumerID: cfg.ConsumerID, RawTopic: cfg.RawTopic, LatestTopic: cfg.LatestTopic, LatestHeight: 100}
	if err := saveDetectorState(cfg.StateFile, state); err != nil {
		t.Fatal(err)
	}
	requests := make(chan string, 8)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/status" {
			_, _ = fmt.Fprint(w, `{"earliest_height":90,"last_height":110}`)
			return
		}
		requests <- r.URL.Query().Get("from_height")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()
	cfg.FullNodes = []string{server.URL}
	cancel := startTestDetector(t, cfg, newFakeConsumer())
	if got := receive(t, requests); got != "101" {
		t.Fatalf("replay starts at %s, want persisted watermark + 1", got)
	}
	cancel()
}

func TestBootstrapUsesEarliestRetainedHeight(t *testing.T) {
	cfg := testConfig(t)
	requests := make(chan string, 8)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/status" {
			_, _ = fmt.Fprint(w, `{"earliest_height":200000000,"last_height":200007200}`)
			return
		}
		requests <- r.URL.Query().Get("from_height")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()
	cfg.FullNodes = []string{server.URL}
	cancel := startTestDetector(t, cfg, newFakeConsumer())
	if got := receive(t, requests); got != "200000000" {
		t.Fatalf("bootstrap requested %s instead of earliest retained height", got)
	}
	cancel()
}

func TestKafkaErrorRetriesAndPersistsBeforeMarking(t *testing.T) {
	cfg := testConfig(t)
	cfg.MessageTimeout = time.Hour
	consumer := newFakeConsumer()
	consumer.onMark = func() {
		state, err := loadDetectorState(cfg)
		if err != nil || state.LatestHeight != 42 {
			t.Errorf("Kafka offset marked before watermark persisted: %+v, %v", state, err)
		}
	}
	consumer.fetches <- kgo.Fetches{{Topics: []kgo.FetchTopic{{Topic: cfg.LatestTopic, Partitions: []kgo.FetchPartition{{Err: errors.New("broker unavailable")}}}}}}
	for _, height := range []int64{42, 41} {
		data, err := (&types.EventStreamResponse{BlockHeight: height}).Marshal()
		if err != nil {
			t.Fatal(err)
		}
		record := &kgo.Record{Topic: cfg.LatestTopic, Value: data}
		consumer.fetches <- kgo.Fetches{{Topics: []kgo.FetchTopic{{Topic: cfg.LatestTopic, Partitions: []kgo.FetchPartition{{Records: []*kgo.Record{record}}}}}}}
	}
	cancel := startTestDetector(t, cfg, consumer)
	receive(t, consumer.marked)
	receive(t, consumer.marked)
	cancel()
}

func TestReplayRequestsAreSingleFlightAndRetryAfterControlFailure(t *testing.T) {
	cfg := testConfig(t)
	var active, maxActive, statusCalls, replayCalls atomic.Int32
	requests := make(chan string, 16)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		concurrent := active.Add(1)
		defer active.Add(-1)
		for previous := maxActive.Load(); concurrent > previous && !maxActive.CompareAndSwap(previous, concurrent); previous = maxActive.Load() {
		}
		if r.URL.Path == "/status" {
			if statusCalls.Add(1) <= 2 {
				w.WriteHeader(http.StatusServiceUnavailable)
				return
			}
			_, _ = fmt.Fprint(w, `{"earliest_height":100,"last_height":110}`)
			return
		}
		replayCalls.Add(1)
		select {
		case <-time.After(100 * time.Millisecond):
		case <-r.Context().Done():
			return
		}
		requests <- r.URL.Query().Get("from_height")
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()
	cfg.FullNodes = []string{server.URL}
	consumer := newFakeConsumer()
	cancel := startTestDetector(t, cfg, consumer)
	// Raw traffic used to spawn one replay goroutine per message during a stall.
	for range 20 {
		record := &kgo.Record{Topic: cfg.RawTopic}
		consumer.fetches <- kgo.Fetches{{Topics: []kgo.FetchTopic{{Topic: cfg.RawTopic, Partitions: []kgo.FetchPartition{{Records: []*kgo.Record{record}}}}}}}
	}
	if got := receive(t, requests); got != "100" {
		t.Fatalf("unexpected replay start %s", got)
	}
	cancel()
	if maxActive.Load() != 1 || replayCalls.Load() != 1 || statusCalls.Load() != 3 {
		t.Fatalf("unbounded requests: active=%d replays=%d statuses=%d", maxActive.Load(), replayCalls.Load(), statusCalls.Load())
	}
}

func TestReplayRejectsUnsupportedAndExpiredHistory(t *testing.T) {
	for _, response := range []string{`{"last_height":110}`, `{"earliest_height":105,"last_height":110}`} {
		t.Run(response, func(t *testing.T) {
			var replayCalled atomic.Bool
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/request" {
					replayCalled.Store(true)
				}
				_, _ = fmt.Fprint(w, response)
			}))
			defer server.Close()
			cfg := testConfig(t)
			cfg.FullNodes = []string{server.URL}
			if err := requestReplay(context.Background(), server.Client(), cfg, 101); err == nil || replayCalled.Load() {
				t.Fatalf("expected explicit error without unsafe replay; err=%v requested=%v", err, replayCalled.Load())
			}
		})
	}
}

func TestStateCannotBeReusedForAnotherConsumer(t *testing.T) {
	cfg := testConfig(t)
	state := detectorState{ConsumerID: "different", RawTopic: cfg.RawTopic, LatestTopic: cfg.LatestTopic, LatestHeight: 100}
	if err := saveDetectorState(cfg.StateFile, state); err != nil {
		t.Fatal(err)
	}
	if _, err := loadDetectorState(cfg); err == nil {
		t.Fatal("expected mismatched state to be rejected")
	}
}

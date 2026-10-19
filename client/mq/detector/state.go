package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// detectorState is tied to one consumer and topic pair. Persisting it before
// marking a Kafka record prevents group offsets from outrunning local state.
type detectorState struct {
	ConsumerID   string `json:"consumer_id"`
	RawTopic     string `json:"raw_topic"`
	LatestTopic  string `json:"latest_topic"`
	LatestHeight int64  `json:"latest_height"`
}

func loadDetectorState(cfg mqDetectorConfig) (detectorState, error) {
	state := detectorState{ConsumerID: cfg.ConsumerID, RawTopic: cfg.RawTopic, LatestTopic: cfg.LatestTopic}
	data, err := os.ReadFile(cfg.StateFile)
	if errors.Is(err, os.ErrNotExist) {
		return state, nil
	}
	if err != nil {
		return state, fmt.Errorf("read detector state: %w", err)
	}
	var saved detectorState
	if err := json.Unmarshal(data, &saved); err != nil {
		return state, fmt.Errorf("decode detector state: %w", err)
	}
	if saved.ConsumerID != state.ConsumerID || saved.RawTopic != state.RawTopic || saved.LatestTopic != state.LatestTopic {
		return state, errors.New("detector state belongs to a different consumer or topic pair")
	}
	if saved.LatestHeight < 0 {
		return state, errors.New("detector state has a negative height")
	}
	return saved, nil
}

func saveDetectorState(path string, state detectorState) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return err
	}
	file, err := os.CreateTemp(dir, ".mq-detector-state-*")
	if err != nil {
		return err
	}
	name := file.Name()
	defer func() { _ = os.Remove(name) }()
	if _, err := file.Write(data); err != nil {
		_ = file.Close()
		return err
	}
	if err := file.Sync(); err != nil {
		_ = file.Close()
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := os.Rename(name, path); err != nil {
		return err
	}
	parent, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer func() { _ = parent.Close() }()
	return parent.Sync()
}

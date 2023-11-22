package event_fetcher

import (
	"context"
)

type EventFetcher interface {
	Fetch(ctx context.Context, height int64) ([]Event, error)
	LastHeight(ctx context.Context) (int64, error)
}

type FilterMap map[string]Event

type Event interface {
	Type() string
	Unmarshal(version string, bz []byte) error
}

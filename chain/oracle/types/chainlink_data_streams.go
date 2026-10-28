package types

import (
	"cosmossdk.io/math"
)

const (
	ChainlinkDataStreamsSchemaV3 uint32 = 3
	ChainlinkDataStreamsSchemaV8 uint32 = 8

	ChainlinkDataStreamsMarketStatusUnknown uint32 = 0
	ChainlinkDataStreamsMarketStatusClosed  uint32 = 1
	ChainlinkDataStreamsMarketStatusOpen    uint32 = 2
)

// NewChainlinkDataStreamsPriceState creates a new ChainlinkDataStreamsPriceState instance.
func NewChainlinkDataStreamsPriceState(
	feedID string,
	reportPrice math.Int,
	validFromTimestamp uint64,
	observationsTimestamp uint64,
	expiresAt uint64,
	reportSchemaVersion uint32,
	marketStatus uint32,
	lastUpdateTimestamp uint64,
	price math.LegacyDec,
	blockTime int64,
) *ChainlinkDataStreamsPriceState {
	return &ChainlinkDataStreamsPriceState{
		FeedId:                feedID,
		ReportPrice:           reportPrice,
		ValidFromTimestamp:    validFromTimestamp,
		ObservationsTimestamp: observationsTimestamp,
		ExpiresAt:             expiresAt,
		ReportSchemaVersion:   reportSchemaVersion,
		MarketStatus:          marketStatus,
		LastUpdateTimestamp:   lastUpdateTimestamp,
		PriceState:            *NewPriceState(price, blockTime),
	}
}

// Update updates the ChainlinkDataStreamsPriceState with new values.
func (c *ChainlinkDataStreamsPriceState) Update(
	reportPrice math.Int,
	validFromTimestamp uint64,
	observationsTimestamp uint64,
	expiresAt uint64,
	reportSchemaVersion uint32,
	marketStatus uint32,
	lastUpdateTimestamp uint64,
	price math.LegacyDec,
	blockTime int64,
) {
	c.ReportPrice = reportPrice
	c.ValidFromTimestamp = validFromTimestamp
	c.ObservationsTimestamp = observationsTimestamp
	c.ExpiresAt = expiresAt
	c.ReportSchemaVersion = reportSchemaVersion
	c.MarketStatus = marketStatus
	c.LastUpdateTimestamp = lastUpdateTimestamp
	c.PriceState.UpdatePrice(price, blockTime)
}

package types

import (
	"fmt"
	"strconv"

	"cosmossdk.io/math"
)

// ValidateCanonicalPythProFeedID ensures s is the decimal form of a uint32 feed ID
// with no leading zeros or stray characters, matching strconv.FormatUint(id, 10).
// This keeps exchange market IDs (which hash the oracle strings) aligned with
// on-chain oracle keys parsed from the same strings.
func ValidateCanonicalPythProFeedID(s string) error {
	id, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid pyth pro feed id %q: %w", s, err)
	}
	if strconv.FormatUint(id, 10) != s {
		return fmt.Errorf("pyth pro feed id %q is not canonical (use %d without leading zeros)", s, id)
	}
	return nil
}

// NewPythProPriceState creates a PythProPriceState for a feed seen for the first time.
func NewPythProPriceState(feedID uint32, price math.LegacyDec, timestamp uint64, blockTime int64) *PythProPriceState {
	return &PythProPriceState{
		FeedId:     feedID,
		Timestamp:  timestamp,
		PriceState: *NewPriceState(price, blockTime),
	}
}

// Update applies a new price to an existing PythProPriceState.
// The caller is responsible for staleness and threshold checks; Update
// unconditionally replaces price and timestamps.
func (p *PythProPriceState) Update(price math.LegacyDec, timestamp uint64, blockTime int64) {
	if timestamp <= p.Timestamp {
		return
	}
	p.Timestamp = timestamp
	p.PriceState.UpdatePrice(price, blockTime)
}

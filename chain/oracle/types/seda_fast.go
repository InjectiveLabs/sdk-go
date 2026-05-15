package types

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"cosmossdk.io/math"
)

// MaxSedaFastUpdateSize is the maximum allowed size in bytes for a single
// MsgRelaySedaFastPrices update entry.
const MaxSedaFastUpdateSize = 8 * 1024 // 8 KiB

// MaxSedaFastUpdatesPerMsg is the maximum number of updates allowed per
// MsgRelaySedaFastPrices message.
const MaxSedaFastUpdatesPerMsg = 64

// MaxSedaFastExponent is the maximum absolute value allowed for the price
// exponent in JSON-format results. It matches LegacyDec's 18-decimal precision
// and comfortably exceeds every real-world Pyth/SEDA exponent (~5–12), while
// preventing an attacker from triggering unbounded big.Int exponentiation via
// LegacyDec.Power.
const MaxSedaFastExponent = 18

// MaxSedaFastMantissaBits is the maximum bit length of the mantissa in
// JSON-format results. A mantissa exceeding this limit would, after
// multiplication by 10^MaxSedaFastExponent, produce a LegacyDec large enough
// to cause excessive memory use in exchange arithmetic. 192 bits leaves ample
// headroom for the ~60-bit multiplier while keeping prices in a sane range.
const MaxSedaFastMantissaBits = 192

// ValidateCanonicalSedaFastFeedID ensures s is the canonical hex form used as
// the on-chain store key for SEDA Fast feeds: lowercase hex, no 0x prefix,
// non-empty, even length. This keeps exchange market IDs (which the exchange
// hashes to look up prices) aligned with the keys the relay path writes.
func ValidateCanonicalSedaFastFeedID(s string) error {
	if s == "" {
		return errors.New("seda fast feed id must not be empty")
	}
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return fmt.Errorf("seda fast feed id %q must not have a 0x prefix", s)
	}
	if strings.ToLower(s) != s {
		return fmt.Errorf("seda fast feed id %q must be lowercase hex", s)
	}
	if _, err := hex.DecodeString(s); err != nil {
		return fmt.Errorf("seda fast feed id %q: %w", s, err)
	}
	return nil
}

// NewSedaFastPriceState creates a SedaFastPriceState for a feed seen for the
// first time. timestamp is dataResult.blockTimestamp in milliseconds;
// blockTime is ctx.BlockTime().Unix().
func NewSedaFastPriceState(feedID string, price math.LegacyDec, timestamp uint64, blockTime int64) *SedaFastPriceState {
	return &SedaFastPriceState{
		FeedId:     feedID,
		Timestamp:  timestamp,
		PriceState: *NewPriceState(price, blockTime),
	}
}

// Update applies a new price to an existing SedaFastPriceState. The caller is
// responsible for the monotonic-timestamp guard and threshold check before
// calling Update.
func (s *SedaFastPriceState) Update(price math.LegacyDec, timestamp uint64, blockTime int64) {
	if timestamp <= s.Timestamp {
		return
	}
	s.Timestamp = timestamp
	s.PriceState.UpdatePrice(price, blockTime)
}

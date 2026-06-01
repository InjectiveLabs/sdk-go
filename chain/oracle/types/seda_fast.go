package types

import (
	"encoding/hex"
	"fmt"
	"strings"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/crypto"
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

// DecodeSedaFastExecInputs maps the SEDA Fast wire representation into the
// bytes used by SEDA's own drId preimage and Injective's composite feed ID.
func DecodeSedaFastExecInputs(raw string) ([]byte, error) {
	stripped := strings.TrimPrefix(strings.TrimPrefix(raw, "0x"), "0X")
	if raw != stripped {
		if stripped == "" {
			return []byte{}, nil
		}
		b, err := hex.DecodeString(stripped)
		if err != nil {
			return nil, fmt.Errorf("seda fast execInputs: invalid hex %q: %w", stripped, err)
		}
		return b, nil
	}
	if stripped == "" {
		return []byte{}, nil
	}
	if isLowercaseHex(stripped) && len(stripped)%2 == 0 {
		b, err := hex.DecodeString(stripped)
		if err != nil {
			return nil, fmt.Errorf("seda fast execInputs: invalid hex %q: %w", stripped, err)
		}
		return b, nil
	}
	return []byte(raw), nil
}

// ComputeSedaFastFeedID derives the stable on-chain feed identifier for a SEDA
// Fast result from the Oracle Program ID and the raw execution input bytes.
func ComputeSedaFastFeedID(execProgramIDHex string, execInputs []byte) (string, error) {
	if err := validateLowercaseHex32("seda fast execProgramId", execProgramIDHex); err != nil {
		return "", err
	}

	programIDBytes, err := hex.DecodeString(execProgramIDHex)
	if err != nil {
		return "", fmt.Errorf("seda fast execProgramId %q: %w", execProgramIDHex, err)
	}

	inputHash := crypto.Keccak256(execInputs)
	preimage := make([]byte, 0, len(programIDBytes)+len(inputHash))
	preimage = append(preimage, programIDBytes...)
	preimage = append(preimage, inputHash...)

	return hex.EncodeToString(crypto.Keccak256(preimage)), nil
}

// ValidateCanonicalSedaFastFeedID ensures s is the canonical composite feed ID
// written by the SEDA Fast relay path: exactly 32 bytes encoded as lowercase hex.
func ValidateCanonicalSedaFastFeedID(s string) error {
	return validateLowercaseHex32("seda fast feed id", s)
}

func validateLowercaseHex32(name, s string) error {
	if s == "" {
		return fmt.Errorf("%s must not be empty", name)
	}
	if len(s) != 64 {
		return fmt.Errorf("%s %q must be exactly 64 characters", name, s)
	}
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return fmt.Errorf("%s %q must not have a 0x prefix", name, s)
	}
	if strings.ToLower(s) != s {
		return fmt.Errorf("%s %q must be lowercase hex", name, s)
	}
	if _, err := hex.DecodeString(s); err != nil {
		return fmt.Errorf("%s %q: %w", name, s, err)
	}
	return nil
}

func isLowercaseHex(s string) bool {
	for _, r := range s {
		if (r < '0' || r > '9') && (r < 'a' || r > 'f') {
			return false
		}
	}
	return true
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

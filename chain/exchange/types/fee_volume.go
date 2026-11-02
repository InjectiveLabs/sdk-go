package types

import (
	"fmt"
	"math"
)

// MaxFeeDiscountBuckets bounds legacy point lookups and per-account history
// work independently of governance input. It accommodates the 28/30-day schedules.
const MaxFeeDiscountBuckets uint64 = 64

func ValidateFeeDiscountBucketConfig(count uint64, duration int64) error {
	if count < 2 || count > MaxFeeDiscountBuckets {
		return fmt.Errorf("fee discount bucket count must be between 2 and %d", MaxFeeDiscountBuckets)
	}
	if duration < 10 || duration > math.MaxInt64/int64(count) {
		return fmt.Errorf("invalid fee discount bucket duration %d", duration)
	}
	return nil
}

package v2

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

const MaxFeeDiscountBuckets = types.MaxFeeDiscountBuckets

// ValidateFeeDiscountVolumes enforces the alignment required by bounded point
// lookups. Empty bucket records carry the current timestamp and grace anchor
// through the existing genesis schema, without exporting derived running totals.
func (gs GenesisState) ValidateFeeDiscountVolumes() error {
	if gs.FeeDiscountSchedule == nil {
		if len(gs.FeeDiscountBucketVolumeAccounts) != 0 || len(gs.FeeDiscountAccountTierTtl) != 0 {
			return errors.New("fee discount state requires a schedule")
		}
		return nil
	}
	schedule := gs.FeeDiscountSchedule
	if err := types.ValidateFeeDiscountBucketConfig(schedule.BucketCount, schedule.BucketDuration); err != nil {
		return err
	}
	var current int64
	for _, bucket := range gs.FeeDiscountBucketVolumeAccounts {
		if bucket == nil || bucket.BucketStartTimestamp <= 0 {
			return errors.New("invalid fee discount bucket timestamp")
		}
		current = max(current, bucket.BucketStartTimestamp)
	}
	seen := make(map[int64]bool)
	for _, bucket := range gs.FeeDiscountBucketVolumeAccounts {
		timestamp := bucket.BucketStartTimestamp
		if seen[timestamp] || (current-timestamp)%schedule.BucketDuration != 0 {
			return fmt.Errorf("duplicate or misaligned fee discount bucket %d", timestamp)
		}
		seen[timestamp] = true
		accounts := make(map[string]bool)
		for _, entry := range bucket.AccountVolume {
			if entry == nil || entry.Volume.IsNil() || entry.Volume.IsNegative() || !entry.Volume.IsInValidRange() {
				return errors.New("invalid fee discount volume")
			}
			account, err := sdk.AccAddressFromBech32(entry.Account)
			if err != nil || accounts[string(account)] {
				return fmt.Errorf("invalid or duplicate fee discount account %s", entry.Account)
			}
			accounts[string(account)] = true
		}
	}
	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *FeeDiscountSchedule) CalculateFeeDiscountTier(
	stakedAmount sdk.Int,
	tradingFeePaid sdk.Dec,
) (
	tierInfo *FeeDiscountTierInfo,
	tierLevel uint64,
) {
	highestTierLevel := 0
	// O(N) but probably the most efficient way nonetheless since we just have 10 tiers and most will be on low tiers
	for idx, tier := range s.TierInfos {
		// both tier conditions must be satisfied to reach a tier level
		if stakedAmount.LT(tier.StakedAmount) || tradingFeePaid.LT(tier.FeePaidAmount) {
			break
		}
		highestTierLevel = idx
	}

	return s.TierInfos[highestTierLevel], uint64(highestTierLevel)
}

func (s *FeeDiscountSchedule) TierOneRequirements() (
	minStakedAmount sdk.Int,
	minTradingFeePaid sdk.Dec,
) {
	return s.TierInfos[1].StakedAmount, s.TierInfos[1].FeePaidAmount
}

type TierLevelMap map[uint64]*FeeDiscountTierInfo

func (s *FeeDiscountSchedule) GetTierLevelMap() TierLevelMap {
	tierLevelMap := make(TierLevelMap, len(s.TierInfos))
	for idx, tierInfo := range s.TierInfos {
		tierLevelMap[uint64(idx)] = tierInfo
	}
	return tierLevelMap
}

func NewFeeDiscountTierTTL(tier uint64, ttlTimestamp int64) *FeeDiscountTierTTL {
	return &FeeDiscountTierTTL{
		Tier:         tier,
		TtlTimestamp: ttlTimestamp,
	}
}

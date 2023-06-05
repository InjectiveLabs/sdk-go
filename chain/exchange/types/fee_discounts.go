package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type FeeDiscountRates struct {
	MakerDiscountRate sdk.Dec
	TakerDiscountRate sdk.Dec
}

func (s *FeeDiscountSchedule) CalculateFeeDiscountTier(
	stakedAmount sdkmath.Int,
	tradingVolume sdk.Dec,
) (
	feeDiscountRates *FeeDiscountRates,
	tierLevel uint64,
) {
	highestTierLevel := 0
	// O(N) but probably the most efficient way nonetheless since we just have 10 tiers and most will be on low tiers
	for idx, tier := range s.TierInfos {
		// both tier conditions must be satisfied to reach a tier level
		if stakedAmount.LT(tier.StakedAmount) || tradingVolume.LT(tier.Volume) {
			break
		}
		highestTierLevel = idx + 1
	}

	tierLevel = uint64(highestTierLevel)

	if tierLevel == 0 {
		feeDiscountRates = &FeeDiscountRates{
			MakerDiscountRate: sdk.ZeroDec(),
			TakerDiscountRate: sdk.ZeroDec(),
		}
	} else {
		feeDiscountRates = &FeeDiscountRates{
			MakerDiscountRate: s.TierInfos[highestTierLevel-1].MakerDiscountRate,
			TakerDiscountRate: s.TierInfos[highestTierLevel-1].TakerDiscountRate,
		}
	}

	return feeDiscountRates, tierLevel
}

func (s *FeeDiscountSchedule) TierOneRequirements() (
	minStakedAmount sdkmath.Int,
	minTradingFeePaid sdk.Dec,
) {
	// s.TierInfos[0] is tier one, since tier 0 is implicit
	return s.TierInfos[0].StakedAmount, s.TierInfos[0].Volume
}

type FeeDiscountRatesMap map[uint64]*FeeDiscountRates

func (s *FeeDiscountSchedule) GetFeeDiscountRatesMap() FeeDiscountRatesMap {
	if s == nil {
		return make(FeeDiscountRatesMap)
	}

	feeDiscountRatesMap := make(FeeDiscountRatesMap, len(s.TierInfos))
	feeDiscountRatesMap[0] = &FeeDiscountRates{
		MakerDiscountRate: sdk.ZeroDec(),
		TakerDiscountRate: sdk.ZeroDec(),
	}

	for idx, tierInfo := range s.TierInfos {
		feeDiscountRatesMap[uint64(idx+1)] = &FeeDiscountRates{
			MakerDiscountRate: tierInfo.MakerDiscountRate,
			TakerDiscountRate: tierInfo.TakerDiscountRate,
		}
	}
	return feeDiscountRatesMap
}

func NewFeeDiscountTierTTL(tier uint64, ttlTimestamp int64) *FeeDiscountTierTTL {
	return &FeeDiscountTierTTL{
		Tier:         tier,
		TtlTimestamp: ttlTimestamp,
	}
}

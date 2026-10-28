package v2

import (
	"cosmossdk.io/math"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

func ValidateMakerWithTakerFeeAndDiscounts(
	makerFeeRate,
	takerFeeRate,
	relayerFeeShareRate,
	minimalProtocolFeeRate math.LegacyDec,
	discountSchedule *FeeDiscountSchedule,
) error {
	if err := types.ValidateMakerWithTakerFee(makerFeeRate, takerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate); err != nil {
		return err
	}

	smallestMakerFeeRate := makerFeeRate
	smallestTakerFeeRate := takerFeeRate
	if discountSchedule != nil && len(discountSchedule.TierInfos) > 0 {
		maxDiscount := discountSchedule.TierInfos[len(discountSchedule.TierInfos)-1]
		smallestTakerFeeRate = takerFeeRate.Mul(math.LegacyOneDec().Sub(maxDiscount.TakerDiscountRate))
		// Negative maker rebates are never discounted at execution.
		if !makerFeeRate.IsNegative() {
			smallestMakerFeeRate = makerFeeRate.Mul(math.LegacyOneDec().Sub(maxDiscount.MakerDiscountRate))
		}
	}

	if err := types.ValidateProtocolFeeRate(smallestMakerFeeRate, smallestTakerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate); err != nil {
		return err
	}

	// Two new orders matched in the same auction both pay taker fees. Discounts
	// can make the effective taker rate lower than the effective maker rate.
	return types.ValidateProtocolFeeRate(smallestTakerFeeRate, smallestTakerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate)
}

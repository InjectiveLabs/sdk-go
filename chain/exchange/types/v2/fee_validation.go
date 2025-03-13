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
	smallestTakerFeeRate := takerFeeRate
	if makerFeeRate.IsNegative() && discountSchedule != nil && len(discountSchedule.TierInfos) > 0 {
		maxTakerDiscount := discountSchedule.TierInfos[len(discountSchedule.TierInfos)-1].TakerDiscountRate
		smallestTakerFeeRate = smallestTakerFeeRate.Mul(math.LegacyOneDec().Sub(maxTakerDiscount))
	}

	return types.ValidateMakerWithTakerFee(makerFeeRate, smallestTakerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate)
}

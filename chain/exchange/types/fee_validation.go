package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ValidateMakerWithTakerFee(makerFeeRate, takerFeeRate, relayerFeeShareRate sdk.Dec) error {
	if makerFeeRate.GT(takerFeeRate) {
		return ErrFeeRatesRelation
	}

	if !makerFeeRate.IsNegative() {
		return nil
	}

	// if makerFeeRate is negative, takerFeeRate * (1 - relayerFeeShareRate) + makerFeeRate < 0
	if takerFeeRate.Mul(sdk.OneDec().Sub(relayerFeeShareRate)).Add(makerFeeRate).IsNegative() {
		return ErrFeeRatesRelation
	}

	return nil
}

func ValidateMakerWithTakerFeeAndDiscounts(makerFeeRate, takerFeeRate, relayerFeeShareRate sdk.Dec, discountSchedule *FeeDiscountSchedule) error {
	smallestTakerFeeRate := takerFeeRate

	if makerFeeRate.IsNegative() && discountSchedule != nil && len(discountSchedule.TierInfos) > 0 {
		maxTakerDiscount := discountSchedule.TierInfos[len(discountSchedule.TierInfos)-1].TakerDiscountRate
		smallestTakerFeeRate = smallestTakerFeeRate.Mul(sdk.OneDec().Sub(maxTakerDiscount))
	}

	return ValidateMakerWithTakerFee(makerFeeRate, smallestTakerFeeRate, relayerFeeShareRate)
}

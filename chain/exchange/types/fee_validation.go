package types

import (
	"fmt"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
)

func ValidateMakerWithTakerFee(makerFeeRate, takerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate math.LegacyDec) error {
	if makerFeeRate.GT(takerFeeRate) {
		return ErrFeeRatesRelation
	}

	if !makerFeeRate.IsNegative() {
		return nil
	}

	// if makerFeeRate is negative, must hold: takerFeeRate * (1 - relayerFeeShareRate) + makerFeeRate > minimalProtocolFeeRate
	if takerFeeRate.Mul(math.LegacyOneDec().Sub(relayerFeeShareRate)).Add(makerFeeRate).LT(minimalProtocolFeeRate) {
		errMsg := fmt.Sprintf("if makerFeeRate (%v) is negative, (takerFeeRate = %v) * (1 - relayerFeeShareRate = %v) + makerFeeRate < %v", makerFeeRate.String(), takerFeeRate.String(), relayerFeeShareRate.String(), minimalProtocolFeeRate.String())
		return errors.Wrap(ErrFeeRatesRelation, errMsg)
	}

	return nil
}

func ValidateMakerWithTakerFeeAndDiscounts(makerFeeRate, takerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate math.LegacyDec, discountSchedule *FeeDiscountSchedule) error {
	smallestTakerFeeRate := takerFeeRate

	if makerFeeRate.IsNegative() && discountSchedule != nil && len(discountSchedule.TierInfos) > 0 {
		maxTakerDiscount := discountSchedule.TierInfos[len(discountSchedule.TierInfos)-1].TakerDiscountRate
		smallestTakerFeeRate = smallestTakerFeeRate.Mul(math.LegacyOneDec().Sub(maxTakerDiscount))
	}

	return ValidateMakerWithTakerFee(makerFeeRate, smallestTakerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate)
}

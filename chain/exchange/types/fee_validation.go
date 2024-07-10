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

	if takerFeeRate.Mul(math.LegacyOneDec().Sub(relayerFeeShareRate)).Add(makerFeeRate).LT(minimalProtocolFeeRate) {
		// if makerFeeRate is negative then takerFeeRate >= (minimalProtocolFeeRate - relayerFeeShareRate)/(1 - makerFeeRate)
		numerator := minimalProtocolFeeRate.Sub(relayerFeeShareRate)
		denominator := math.LegacyOneDec().Sub(makerFeeRate)
		return errors.Wrap(ErrFeeRatesRelation, fmt.Sprintf("if maker_fee_rate is negative (%v), taker_fee_rate must be GTE than %v [ taker_fee_rate >= (minimum_protocol_fee_rate - maker_fee_rate)/(1 - relayer_fee_share_rate) ]", makerFeeRate.String(), numerator.Quo(denominator).String()))
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

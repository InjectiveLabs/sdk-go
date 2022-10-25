package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func ValidateMakerWithTakerFee(makerFeeRate, takerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate sdk.Dec) error {
	if makerFeeRate.GT(takerFeeRate) {
		return ErrFeeRatesRelation
	}

	if !makerFeeRate.IsNegative() {
		return nil
	}

	// if makerFeeRate is negative, must hold: takerFeeRate * (1 - relayerFeeShareRate) + makerFeeRate > minimalProtocolFeeRate
	if takerFeeRate.Mul(sdk.OneDec().Sub(relayerFeeShareRate)).Add(makerFeeRate).LT(minimalProtocolFeeRate) {
		errMsg := fmt.Sprintf("if makerFeeRate (%v) is negative, (takerFeeRate = %v) * (1 - relayerFeeShareRate = %v) + makerFeeRate < %v", makerFeeRate.String(), takerFeeRate.String(), relayerFeeShareRate.String(), minimalProtocolFeeRate.String())
		return sdkerrors.Wrap(ErrFeeRatesRelation, errMsg)
	}

	return nil
}

func ValidateMakerWithTakerFeeAndDiscounts(makerFeeRate, takerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate sdk.Dec, discountSchedule *FeeDiscountSchedule) error {
	smallestTakerFeeRate := takerFeeRate

	if makerFeeRate.IsNegative() && discountSchedule != nil && len(discountSchedule.TierInfos) > 0 {
		maxTakerDiscount := discountSchedule.TierInfos[len(discountSchedule.TierInfos)-1].TakerDiscountRate
		smallestTakerFeeRate = smallestTakerFeeRate.Mul(sdk.OneDec().Sub(maxTakerDiscount))
	}

	return ValidateMakerWithTakerFee(makerFeeRate, smallestTakerFeeRate, relayerFeeShareRate, minimalProtocolFeeRate)
}

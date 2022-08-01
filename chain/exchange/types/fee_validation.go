package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func ValidateMakerWithTakerFee(makerFeeRate, takerFeeRate, relayerFeeShareRate sdk.Dec) error {
	if makerFeeRate.GT(takerFeeRate) {
		return ErrFeeRatesRelation
	}

	if !makerFeeRate.IsNegative() {
		return nil
	}

	// if makerFeeRate is negative, should: takerFeeRate * (1 - relayerFeeShareRate)  > | makerFeeRate |
	if takerFeeRate.Mul(sdk.OneDec().Sub(relayerFeeShareRate)).Add(makerFeeRate).IsNegative() {
		errMsg := fmt.Sprintf("if makerFeeRate (%v) is negative, (takerFeeRate = %v) * (1 - relayerFeeShareRate = %v) + makerFeeRate < 0", makerFeeRate.String(), takerFeeRate.String(), relayerFeeShareRate.String())
		return sdkerrors.Wrap(ErrFeeRatesRelation, errMsg)
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

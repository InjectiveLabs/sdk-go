package v2

import "cosmossdk.io/math"

type TradeFeeData struct {
	TotalTradeFee math.LegacyDec
	TraderFee     math.LegacyDec

	FeeRecipientReward     math.LegacyDec
	AuctionFeeReward       math.LegacyDec
	DiscountedTradeFeeRate math.LegacyDec
}

func NewEmptyTradeFeeData(discountedTradeFeeRate math.LegacyDec) *TradeFeeData {
	return &TradeFeeData{
		TotalTradeFee:          math.LegacyZeroDec(),
		TraderFee:              math.LegacyZeroDec(),
		FeeRecipientReward:     math.LegacyZeroDec(),
		AuctionFeeReward:       math.LegacyZeroDec(),
		DiscountedTradeFeeRate: discountedTradeFeeRate,
	}
}

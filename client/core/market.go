package core

import (
	sdkmath "cosmossdk.io/math"
	"github.com/shopspring/decimal"
)

type SpotMarket interface {
	QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec
	PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec
	NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec
	QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal
	PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal
	NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal
}

type DerivativeMarket interface {
	QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec
	PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec
	MarginToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec
	CalculateMarginInChainFormat(humanReadableQuantity, humanReadablePrice, leverage decimal.Decimal) sdkmath.LegacyDec
	NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec
	QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal
	PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal
	MarginFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal
	NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal
}

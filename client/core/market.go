package core

import (
	sdkmath "cosmossdk.io/math"
	"github.com/shopspring/decimal"

	"github.com/InjectiveLabs/sdk-go/client/common"
)

const AdditionalChainFormatDecimals = 18

type SpotMarket struct {
	Id                  string
	Status              string
	Ticker              string
	BaseToken           Token
	QuoteToken          Token
	MakerFeeRate        decimal.Decimal
	TakerFeeRate        decimal.Decimal
	ServiceProviderFee  decimal.Decimal
	MinPriceTickSize    decimal.Decimal
	MinQuantityTickSize decimal.Decimal
	MinNotional         decimal.Decimal
}

func (spotMarket SpotMarket) QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(spotMarket.MinQuantityTickSize, 0).Mul(spotMarket.MinQuantityTickSize)
	chainFormattedValue := quantizedValue.Mul(decimal.New(1, spotMarket.BaseToken.Decimals))
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarket) PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(spotMarket.MinPriceTickSize, 0).Mul(spotMarket.MinPriceTickSize)
	decimals := spotMarket.QuoteToken.Decimals - spotMarket.BaseToken.Decimals
	chainFormattedValue := quantizedValue.Mul(decimal.New(1, decimals))
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarket) NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := spotMarket.QuoteToken.Decimals
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.BigInt()
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarket) QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String()).Div(decimal.New(1, spotMarket.BaseToken.Decimals))
}

func (spotMarket SpotMarket) PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := spotMarket.BaseToken.Decimals - spotMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (spotMarket SpotMarket) NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -spotMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (spotMarket SpotMarket) QuantityFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(spotMarket.QuantityFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (spotMarket SpotMarket) PriceFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(spotMarket.PriceFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (spotMarket SpotMarket) NotionalFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(spotMarket.NotionalFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

type DerivativeMarket struct {
	Id                     string
	Status                 string
	Ticker                 string
	OracleBase             string
	OracleQuote            string
	OracleType             string
	OracleScaleFactor      uint32
	InitialMarginRatio     decimal.Decimal
	MaintenanceMarginRatio decimal.Decimal
	QuoteToken             Token
	MakerFeeRate           decimal.Decimal
	TakerFeeRate           decimal.Decimal
	ServiceProviderFee     decimal.Decimal
	MinPriceTickSize       decimal.Decimal
	MinQuantityTickSize    decimal.Decimal
	MinNotional            decimal.Decimal
}

func (derivativeMarket DerivativeMarket) QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	chainFormattedValue := quantizedValue
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(derivativeMarket.MinPriceTickSize, 0).Mul(derivativeMarket.MinPriceTickSize)
	decimals := derivativeMarket.QuoteToken.Decimals
	chainFormattedValue := quantizedValue.Mul(decimal.New(1, decimals))
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) MarginToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	return derivativeMarket.NotionalToChainFormat(humanReadableValue)
}

func (derivativeMarket DerivativeMarket) CalculateMarginInChainFormat(humanReadableQuantity, humanReadablePrice, leverage decimal.Decimal) sdkmath.LegacyDec {
	margin := humanReadableQuantity.Mul(humanReadablePrice).Div(leverage)
	// We are using the min_quantity_tick_size to quantize the margin because that is the way margin is validated
	// in the chain (it might be changed to a min_notional in the future)
	quantizedMargin := margin.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)

	return derivativeMarket.NotionalToChainFormat(quantizedMargin)
}

func (derivativeMarket DerivativeMarket) NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := derivativeMarket.QuoteToken.Decimals
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.BigInt()
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (DerivativeMarket) QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String())
}

func (derivativeMarket DerivativeMarket) PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -derivativeMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (derivativeMarket DerivativeMarket) MarginFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return derivativeMarket.NotionalFromChainFormat(chainValue)
}

func (derivativeMarket DerivativeMarket) NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -derivativeMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (derivativeMarket DerivativeMarket) QuantityFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(derivativeMarket.QuantityFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (derivativeMarket DerivativeMarket) PriceFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(derivativeMarket.PriceFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (derivativeMarket DerivativeMarket) MarginFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(derivativeMarket.MarginFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (derivativeMarket DerivativeMarket) NotionalFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(derivativeMarket.NotionalFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

type BinaryOptionMarket struct {
	Id                  string
	Status              string
	Ticker              string
	OracleSymbol        string
	OracleProvider      string
	OracleType          string
	OracleScaleFactor   uint32
	ExpirationTimestamp int64
	SettlementTimestamp int64
	QuoteToken          Token
	MakerFeeRate        decimal.Decimal
	TakerFeeRate        decimal.Decimal
	ServiceProviderFee  decimal.Decimal
	MinPriceTickSize    decimal.Decimal
	MinQuantityTickSize decimal.Decimal
	MinNotional         decimal.Decimal
	SettlementPrice     *decimal.Decimal
}

func (market BinaryOptionMarket) QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(market.MinQuantityTickSize, 0).Mul(market.MinQuantityTickSize)
	chainFormattedValue := quantizedValue
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (market BinaryOptionMarket) PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(market.MinPriceTickSize, 0).Mul(market.MinPriceTickSize)
	decimals := market.QuoteToken.Decimals
	chainFormattedValue := quantizedValue.Mul(decimal.New(1, decimals))
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (market BinaryOptionMarket) MarginToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	return market.NotionalToChainFormat(humanReadableValue)
}

func (market BinaryOptionMarket) CalculateMarginInChainFormat(humanReadableQuantity, humanReadablePrice, leverage decimal.Decimal) sdkmath.LegacyDec {
	margin := humanReadableQuantity.Mul(humanReadablePrice).Div(leverage)
	// We are using the min_quantity_tick_size to quantize the margin because that is the way margin is validated
	// in the chain (it might be changed to a min_notional in the future)
	quantizedMargin := margin.DivRound(market.MinQuantityTickSize, 0).Mul(market.MinQuantityTickSize)

	return market.NotionalToChainFormat(quantizedMargin)
}

func (market BinaryOptionMarket) NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := market.QuoteToken.Decimals
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.BigInt()
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (BinaryOptionMarket) QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String())
}

func (market BinaryOptionMarket) PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -market.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (market BinaryOptionMarket) MarginFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return market.NotionalFromChainFormat(chainValue)
}

func (market BinaryOptionMarket) NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -market.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (market BinaryOptionMarket) QuantityFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(market.QuantityFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (market BinaryOptionMarket) PriceFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(market.PriceFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (market BinaryOptionMarket) MarginFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(market.MarginFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

func (market BinaryOptionMarket) NotionalFromExtendedChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return common.RemoveExtraDecimals(market.NotionalFromChainFormat(chainValue), AdditionalChainFormatDecimals)
}

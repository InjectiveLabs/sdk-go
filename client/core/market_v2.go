package core

import (
	sdkmath "cosmossdk.io/math"
	"github.com/shopspring/decimal"
)

type SpotMarketV2 struct {
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
	BaseDecimals        uint32
	QuoteDecimals       uint32
}

var _ SpotMarket = &SpotMarketV2{}

func (spotMarket SpotMarketV2) QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(spotMarket.MinQuantityTickSize, 0).Mul(spotMarket.MinQuantityTickSize)
	chainFormattedValue := quantizedValue.Mul(decimal.New(1, int32(spotMarket.BaseDecimals)))
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarketV2) PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(spotMarket.MinPriceTickSize, 0).Mul(spotMarket.MinPriceTickSize)
	decimals := int32(spotMarket.QuoteDecimals - spotMarket.BaseDecimals)
	chainFormattedValue := quantizedValue.Mul(decimal.New(1, decimals))
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarketV2) NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := int32(spotMarket.QuoteDecimals)
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.BigInt()
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarketV2) QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String()).Div(decimal.New(1, int32(spotMarket.BaseDecimals)))
}

func (spotMarket SpotMarketV2) PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := int32(spotMarket.BaseDecimals - spotMarket.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (spotMarket SpotMarketV2) NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(spotMarket.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

type DerivativeMarketV2 struct {
	Id                     string
	Status                 string
	Ticker                 string
	OracleBase             string
	OracleQuote            string
	OracleType             string
	OracleScaleFactor      uint32
	InitialMarginRatio     decimal.Decimal
	MaintenanceMarginRatio decimal.Decimal
	ReduceMarginRatio      decimal.Decimal
	QuoteToken             Token
	MakerFeeRate           decimal.Decimal
	TakerFeeRate           decimal.Decimal
	ServiceProviderFee     decimal.Decimal
	MinPriceTickSize       decimal.Decimal
	MinQuantityTickSize    decimal.Decimal
	MinNotional            decimal.Decimal
	QuoteDecimals          uint32
}

var _ DerivativeMarket = &DerivativeMarketV2{}

func (derivativeMarket DerivativeMarketV2) QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	chainFormattedValue := quantizedValue
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarketV2) PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(derivativeMarket.MinPriceTickSize, 0).Mul(derivativeMarket.MinPriceTickSize)
	decimals := int32(derivativeMarket.QuoteDecimals)
	chainFormattedValue := quantizedValue.Mul(decimal.New(1, decimals))
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarketV2) MarginToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	return derivativeMarket.NotionalToChainFormat(humanReadableValue)
}

func (derivativeMarket DerivativeMarketV2) CalculateMarginInChainFormat(humanReadableQuantity, humanReadablePrice, leverage decimal.Decimal) sdkmath.LegacyDec {
	margin := humanReadableQuantity.Mul(humanReadablePrice).Div(leverage)
	// We are using the min_quantity_tick_size to quantize the margin because that is the way margin is validated
	// in the chain (it might be changed to a min_notional in the future)
	quantizedMargin := margin.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)

	return derivativeMarket.NotionalToChainFormat(quantizedMargin)
}

func (derivativeMarket DerivativeMarketV2) NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := int32(derivativeMarket.QuoteDecimals)
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.BigInt()
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (DerivativeMarketV2) QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String())
}

func (derivativeMarket DerivativeMarketV2) PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(derivativeMarket.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (derivativeMarket DerivativeMarketV2) MarginFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return derivativeMarket.NotionalFromChainFormat(chainValue)
}

func (derivativeMarket DerivativeMarketV2) NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(derivativeMarket.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

type BinaryOptionMarketV2 struct {
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
	QuoteDecimals       uint32
}

var _ DerivativeMarket = &BinaryOptionMarketV2{}

func (market BinaryOptionMarketV2) QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(market.MinQuantityTickSize, 0).Mul(market.MinQuantityTickSize)
	chainFormattedValue := quantizedValue
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (market BinaryOptionMarketV2) PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	quantizedValue := humanReadableValue.DivRound(market.MinPriceTickSize, 0).Mul(market.MinPriceTickSize)
	decimals := int32(market.QuoteDecimals)
	chainFormattedValue := quantizedValue.Mul(decimal.New(1, decimals))
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(chainFormattedValue.String())

	return valueInChainFormat
}

func (market BinaryOptionMarketV2) MarginToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	return market.NotionalToChainFormat(humanReadableValue)
}

func (market BinaryOptionMarketV2) CalculateMarginInChainFormat(humanReadableQuantity, humanReadablePrice, leverage decimal.Decimal) sdkmath.LegacyDec {
	margin := humanReadableQuantity.Mul(humanReadablePrice).Div(leverage)
	// We are using the min_quantity_tick_size to quantize the margin because that is the way margin is validated
	// in the chain (it might be changed to a min_notional in the future)
	quantizedMargin := margin.DivRound(market.MinQuantityTickSize, 0).Mul(market.MinQuantityTickSize)

	return market.NotionalToChainFormat(quantizedMargin)
}

func (market BinaryOptionMarketV2) NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := int32(market.QuoteDecimals)
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.BigInt()
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (BinaryOptionMarketV2) QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String())
}

func (market BinaryOptionMarketV2) PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(market.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (market BinaryOptionMarketV2) MarginFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return market.NotionalFromChainFormat(chainValue)
}

func (market BinaryOptionMarketV2) NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(market.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

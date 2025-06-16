package core

import (
	sdkmath "cosmossdk.io/math"
	"github.com/shopspring/decimal"
)

type SpotMarketV1 struct {
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

var _ SpotMarket = &SpotMarketV1{}

func (spotMarket SpotMarketV1) QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, int32(spotMarket.BaseDecimals)))
	quantizedValue := chainFormattedValue.DivRound(spotMarket.MinQuantityTickSize, 0).Mul(spotMarket.MinQuantityTickSize)
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarketV1) PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := int32(spotMarket.QuoteDecimals - spotMarket.BaseDecimals)
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.DivRound(spotMarket.MinPriceTickSize, 0).Mul(spotMarket.MinPriceTickSize)
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarketV1) NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := int32(spotMarket.QuoteDecimals)
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.Ceil()
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarketV1) QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String()).Div(decimal.New(1, int32(spotMarket.BaseDecimals)))
}

func (spotMarket SpotMarketV1) PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := int32(spotMarket.BaseDecimals - spotMarket.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (spotMarket SpotMarketV1) NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(spotMarket.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

type DerivativeMarketV1 struct {
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
	QuoteDecimals          uint32
}

var _ DerivativeMarket = &DerivativeMarketV1{}

func (derivativeMarket DerivativeMarketV1) QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	chainFormattedValue := humanReadableValue
	quantizedValue := chainFormattedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarketV1) PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := int32(derivativeMarket.QuoteDecimals)
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.DivRound(derivativeMarket.MinPriceTickSize, 0).Mul(derivativeMarket.MinPriceTickSize)
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarketV1) MarginToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	return derivativeMarket.NotionalToChainFormat(humanReadableValue)
}

func (derivativeMarket DerivativeMarketV1) CalculateMarginInChainFormat(humanReadableQuantity, humanReadablePrice, leverage decimal.Decimal) sdkmath.LegacyDec {
	chainFormattedQuantity := humanReadableQuantity
	chainFormattedPrice := humanReadablePrice.Mul(decimal.New(1, int32(derivativeMarket.QuoteDecimals)))

	margin := chainFormattedQuantity.Mul(chainFormattedPrice).Div(leverage)
	// We are using the min_quantity_tick_size to quantize the margin because that is the way margin is validated
	// in the chain (it might be changed to a min_notional in the future)
	quantizedMargin := margin.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedMargin.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarketV1) NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := int32(derivativeMarket.QuoteDecimals)
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.Ceil()
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarketV1) QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String())
}

func (derivativeMarket DerivativeMarketV1) PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(derivativeMarket.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (derivativeMarket DerivativeMarketV1) MarginFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return derivativeMarket.NotionalFromChainFormat(chainValue)
}

func (derivativeMarket DerivativeMarketV1) NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(derivativeMarket.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

type BinaryOptionMarketV1 struct {
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

var _ DerivativeMarket = &BinaryOptionMarketV1{}

func (market BinaryOptionMarketV1) QuantityToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	chainFormattedValue := humanReadableValue
	quantizedValue := chainFormattedValue.DivRound(market.MinQuantityTickSize, 0).Mul(market.MinQuantityTickSize)
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (market BinaryOptionMarketV1) PriceToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := int32(market.QuoteDecimals)
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.DivRound(market.MinPriceTickSize, 0).Mul(market.MinPriceTickSize)
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (market BinaryOptionMarketV1) MarginToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	return market.NotionalToChainFormat(humanReadableValue)
}

func (market BinaryOptionMarketV1) CalculateMarginInChainFormat(humanReadableQuantity, humanReadablePrice, leverage decimal.Decimal) sdkmath.LegacyDec {
	chainFormattedQuantity := humanReadableQuantity
	chainFormattedPrice := humanReadablePrice.Mul(decimal.New(1, int32(market.QuoteDecimals)))

	margin := chainFormattedQuantity.Mul(chainFormattedPrice).Div(leverage)
	// We are using the min_quantity_tick_size to quantize the margin because that is the way margin is validated
	// in the chain (it might be changed to a min_notional in the future)
	quantizedMargin := margin.DivRound(market.MinQuantityTickSize, 0).Mul(market.MinQuantityTickSize)
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedMargin.String())

	return valueInChainFormat
}

func (market BinaryOptionMarketV1) NotionalToChainFormat(humanReadableValue decimal.Decimal) sdkmath.LegacyDec {
	decimals := int32(market.QuoteDecimals)
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.Ceil()
	valueInChainFormat, _ := sdkmath.LegacyNewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (market BinaryOptionMarketV1) QuantityFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String())
}

func (market BinaryOptionMarketV1) PriceFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(market.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (market BinaryOptionMarketV1) MarginFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	return market.NotionalFromChainFormat(chainValue)
}

func (market BinaryOptionMarketV1) NotionalFromChainFormat(chainValue sdkmath.LegacyDec) decimal.Decimal {
	decimals := -int32(market.QuoteDecimals)
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

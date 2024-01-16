package core

import (
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/shopspring/decimal"
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
}

func (spotMarket SpotMarket) QuantityToChainFormat(humanReadableValue decimal.Decimal) cosmtypes.Dec {
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, spotMarket.BaseToken.Decimals))
	quantizedValue := chainFormattedValue.DivRound(spotMarket.MinQuantityTickSize, 0).Mul(spotMarket.MinQuantityTickSize)
	valueInChainFormat, _ := cosmtypes.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarket) PriceToChainFormat(humanReadableValue decimal.Decimal) cosmtypes.Dec {
	decimals := spotMarket.QuoteToken.Decimals - spotMarket.BaseToken.Decimals
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.DivRound(spotMarket.MinPriceTickSize, 0).Mul(spotMarket.MinPriceTickSize)
	valueInChainFormat, _ := cosmtypes.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (spotMarket SpotMarket) QuantityFromChainFormat(chainValue cosmtypes.Dec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String()).Div(decimal.New(1, spotMarket.BaseToken.Decimals))
}

func (spotMarket SpotMarket) PriceFromChainFormat(chainValue cosmtypes.Dec) decimal.Decimal {
	decimals := spotMarket.BaseToken.Decimals - spotMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
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
}

func (derivativeMarket DerivativeMarket) QuantityToChainFormat(humanReadableValue decimal.Decimal) cosmtypes.Dec {
	chainFormattedValue := humanReadableValue
	quantizedValue := chainFormattedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	valueInChainFormat, _ := cosmtypes.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) PriceToChainFormat(humanReadableValue decimal.Decimal) cosmtypes.Dec {
	decimals := derivativeMarket.QuoteToken.Decimals
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.DivRound(derivativeMarket.MinPriceTickSize, 0).Mul(derivativeMarket.MinPriceTickSize)
	valueInChainFormat, _ := cosmtypes.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) MarginToChainFormat(humanReadableValue decimal.Decimal) cosmtypes.Dec {
	decimals := derivativeMarket.QuoteToken.Decimals
	chainFormattedValue := humanReadableValue.Mul(decimal.New(1, decimals))
	quantizedValue := chainFormattedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	valueInChainFormat, _ := cosmtypes.NewDecFromStr(quantizedValue.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) CalculateMarginInChainFormat(humanReadableQuantity decimal.Decimal, humanReadablePrice decimal.Decimal, leverage decimal.Decimal) cosmtypes.Dec {
	chainFormattedQuantity := humanReadableQuantity
	chainFormattedPrice := humanReadablePrice.Mul(decimal.New(1, derivativeMarket.QuoteToken.Decimals))

	margin := chainFormattedQuantity.Mul(chainFormattedPrice).Div(leverage)
	// We are using the min_quantity_tick_size to quantize the margin because that is the way margin is validated
	// in the chain (it might be changed to a min_notional in the future)
	quantizedMargin := margin.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	valueInChainFormat, _ := cosmtypes.NewDecFromStr(quantizedMargin.String())

	return valueInChainFormat
}

func (derivativeMarket DerivativeMarket) QuantityFromChainFormat(chainValue cosmtypes.Dec) decimal.Decimal {
	return decimal.RequireFromString(chainValue.String())
}

func (derivativeMarket DerivativeMarket) PriceFromChainFormat(chainValue cosmtypes.Dec) decimal.Decimal {
	decimals := -derivativeMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

func (derivativeMarket DerivativeMarket) MarginFromChainFormat(chainValue cosmtypes.Dec) decimal.Decimal {
	decimals := -derivativeMarket.QuoteToken.Decimals
	return decimal.RequireFromString(chainValue.String()).Mul(decimal.New(1, decimals))
}

package core

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/huandu/go-assert"
	"github.com/shopspring/decimal"
	"testing"
)

func createINJUSDTSpotMarket() SpotMarket {
	injToken := createINJToken()
	usdtToken := createUSDTToken()

	makerFeeRate := decimal.RequireFromString("-0.0001")
	takerFeeRate := decimal.RequireFromString("0.001")
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("0.000000000000001")
	minQuantityTickSize := decimal.RequireFromString("1000000000000000")

	market := SpotMarket{
		Id:                  "0x7a57e705bb4e09c88aecfc295569481dbf2fe1d5efe364651fbe72385938e9b0",
		Status:              "active",
		Ticker:              "INJ/USDT",
		BaseToken:           injToken,
		QuoteToken:          usdtToken,
		MakerFeeRate:        makerFeeRate,
		TakerFeeRate:        takerFeeRate,
		ServiceProviderFee:  serviceProviderFee,
		MinPriceTickSize:    minPriceTickSize,
		MinQuantityTickSize: minQuantityTickSize,
	}
	return market
}

func createBTCUSDTPerpMarket() DerivativeMarket {
	usdtToken := createUSDTToken()

	initialMarginRatio := decimal.RequireFromString("0.095")
	maintenanceMarginRatio := decimal.RequireFromString("0.025")
	makerFeeRate := decimal.RequireFromString("-0.0001")
	takerFeeRate := decimal.RequireFromString("0.001")
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("1000000")
	minQuantityTickSize := decimal.RequireFromString("0.0001")

	market := DerivativeMarket{
		Id:                     "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce",
		Status:                 "active",
		Ticker:                 "BTC/USDT PERP",
		OracleBase:             "BTC",
		OracleQuote:            usdtToken.Symbol,
		OracleType:             "bandibc",
		OracleScaleFactor:      6,
		InitialMarginRatio:     initialMarginRatio,
		MaintenanceMarginRatio: maintenanceMarginRatio,
		QuoteToken:             usdtToken,
		MakerFeeRate:           makerFeeRate,
		TakerFeeRate:           takerFeeRate,
		ServiceProviderFee:     serviceProviderFee,
		MinPriceTickSize:       minPriceTickSize,
		MinQuantityTickSize:    minQuantityTickSize,
	}
	return market
}

// Spot market tests

func TestConvertQuantityToChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.QuantityToChainFormat(originalQuantity)
	expectedValue := originalQuantity.Mul(decimal.New(1, spotMarket.BaseToken.Decimals))
	quantizedValue := expectedValue.DivRound(spotMarket.MinQuantityTickSize, 0).Mul(spotMarket.MinQuantityTickSize)
	quantizedChainFormatValue := types.MustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.PriceToChainFormat(originalPrice)
	priceDecimals := spotMarket.QuoteToken.Decimals - spotMarket.BaseToken.Decimals
	expectedValue := originalPrice.Mul(decimal.New(1, priceDecimals))
	quantizedValue := expectedValue.DivRound(spotMarket.MinPriceTickSize, 0).Mul(spotMarket.MinPriceTickSize)
	quantizedChainFormatValue := types.MustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity.Mul(decimal.New(1, spotMarket.BaseToken.Decimals))
	humanReadableQuantity := spotMarket.QuantityFromChainFormat(types.MustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := spotMarket.QuoteToken.Decimals - spotMarket.BaseToken.Decimals
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := spotMarket.PriceFromChainFormat(types.MustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

//Derivative markets tests

func TestConvertQuantityToChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.QuantityToChainFormat(originalQuantity)
	quantizedValue := originalQuantity.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	quantizedChainFormatValue := types.MustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.PriceToChainFormat(originalPrice)
	priceDecimals := derivativeMarket.QuoteToken.Decimals
	expectedValue := originalPrice.Mul(decimal.New(1, priceDecimals))
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinPriceTickSize, 0).Mul(derivativeMarket.MinPriceTickSize)
	quantizedChainFormatValue := types.MustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertMarginToChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.MarginToChainFormat(originalPrice)
	marginDecimals := derivativeMarket.QuoteToken.Decimals
	expectedValue := originalPrice.Mul(decimal.New(1, marginDecimals))
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	quantizedChainFormatValue := types.MustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestCalculateMarginInChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	originalQuantity := decimal.RequireFromString("10")
	originalPrice := decimal.RequireFromString("123.456789")
	originalLeverage := decimal.RequireFromString("2.5")

	chainValue := derivativeMarket.CalculateMarginInChainFormat(originalQuantity, originalPrice, originalLeverage)
	decimals := derivativeMarket.QuoteToken.Decimals
	expectedValue := originalQuantity.Mul(originalPrice).Div(originalLeverage).Mul(decimal.New(1, decimals))
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	legacyDecimalQuantizedValue := types.MustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, chainValue.Equal(legacyDecimalQuantizedValue))
}

func TestConvertQuantityFromChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity
	humanReadableQuantity := derivativeMarket.QuantityFromChainFormat(types.MustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := derivativeMarket.QuoteToken.Decimals
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := derivativeMarket.PriceFromChainFormat(types.MustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertMarginFromChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedMargin := decimal.RequireFromString("123.456")

	marginDecimals := derivativeMarket.QuoteToken.Decimals
	chainFormatMargin := expectedMargin.Mul(decimal.New(1, marginDecimals))
	humanReadablePrice := derivativeMarket.MarginFromChainFormat(types.MustNewDecFromStr(chainFormatMargin.String()))

	assert.Assert(t, expectedMargin.Equal(humanReadablePrice))
}

package core

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/huandu/go-assert"
	"github.com/shopspring/decimal"
)

func createINJUSDTSpotMarketV2() SpotMarketV2 {
	injToken := createINJToken()
	usdtToken := createUSDTToken()

	makerFeeRate := decimal.RequireFromString("-0.0001")
	takerFeeRate := decimal.RequireFromString("0.001")
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("0.01")
	minQuantityTickSize := decimal.RequireFromString("0.000001")
	minNotional := decimal.RequireFromString("1")

	market := SpotMarketV2{
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
		MinNotional:         minNotional,
		BaseDecimals:        18,
		QuoteDecimals:       6,
	}
	return market
}

func createBTCUSDTPerpMarketV2() DerivativeMarketV2 {
	usdtToken := createUSDTToken()

	initialMarginRatio := decimal.RequireFromString("0.095")
	maintenanceMarginRatio := decimal.RequireFromString("0.025")
	reduceMarginRatio := decimal.RequireFromString("0.01")
	makerFeeRate := decimal.RequireFromString("-0.0001")
	takerFeeRate := decimal.RequireFromString("0.001")
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("0.01")
	minQuantityTickSize := decimal.RequireFromString("0.0001")
	minNotional := decimal.RequireFromString("1")

	market := DerivativeMarketV2{
		Id:                     "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce",
		Status:                 "active",
		Ticker:                 "BTC/USDT PERP",
		OracleBase:             "BTC",
		OracleQuote:            usdtToken.Symbol,
		OracleType:             "bandibc",
		OracleScaleFactor:      0,
		InitialMarginRatio:     initialMarginRatio,
		MaintenanceMarginRatio: maintenanceMarginRatio,
		ReduceMarginRatio:      reduceMarginRatio,
		QuoteToken:             usdtToken,
		MakerFeeRate:           makerFeeRate,
		TakerFeeRate:           takerFeeRate,
		ServiceProviderFee:     serviceProviderFee,
		MinPriceTickSize:       minPriceTickSize,
		MinQuantityTickSize:    minQuantityTickSize,
		MinNotional:            minNotional,
		QuoteDecimals:          6,
	}
	return market
}

func createBetBinaryOptionMarketV2() BinaryOptionMarketV2 {
	usdtToken := createUSDTToken()

	makerFeeRate := decimal.Zero
	takerFeeRate := decimal.Zero
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("0.001")
	minQuantityTickSize := decimal.RequireFromString("1")
	minNotional := decimal.RequireFromString("0.00001")

	market := BinaryOptionMarketV2{
		Id:                  "0x230dcce315364ff6360097838701b14713e2f4007d704df20ed3d81d09eec957",
		Status:              "active",
		Ticker:              "5fdbe0b1-1707800399-WAS",
		OracleSymbol:        "Frontrunner",
		OracleProvider:      "Frontrunner",
		OracleType:          "provider",
		OracleScaleFactor:   0,
		ExpirationTimestamp: 1707800399,
		SettlementTimestamp: 1707843599,
		QuoteToken:          usdtToken,
		MakerFeeRate:        makerFeeRate,
		TakerFeeRate:        takerFeeRate,
		ServiceProviderFee:  serviceProviderFee,
		MinPriceTickSize:    minPriceTickSize,
		MinQuantityTickSize: minQuantityTickSize,
		MinNotional:         minNotional,
		QuoteDecimals:       6,
	}
	return market
}

// Spot market tests

func TestConvertQuantityToChainFormatForSpotMarketV2(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV2()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.QuantityToChainFormat(originalQuantity)
	expectedValue := originalQuantity.Mul(decimal.New(1, int32(spotMarket.BaseDecimals)))
	quantizedValue := expectedValue.DivRound(spotMarket.MinQuantityTickSize, 0).Mul(spotMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForSpotMarketV2(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV2()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.PriceToChainFormat(originalPrice)
	quantizedValue := originalPrice.DivRound(spotMarket.MinPriceTickSize, 0).Mul(spotMarket.MinPriceTickSize)
	priceDecimals := int32(spotMarket.QuoteDecimals - spotMarket.BaseDecimals)
	expectedValue := quantizedValue.Mul(decimal.New(1, priceDecimals))
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertNotionalToChainFormatForSpotMarketV2(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV2()
	originalNotional := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.NotionalToChainFormat(originalNotional)
	notionalDecimals := int32(spotMarket.QuoteDecimals)
	expectedValue := originalNotional.Mul(decimal.New(1, notionalDecimals))
	chainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, chainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForSpotMarketV2(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV2()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity.Mul(decimal.New(1, int32(spotMarket.BaseDecimals)))
	humanReadableQuantity := spotMarket.QuantityFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForSpotMarketV2(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV2()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := int32(spotMarket.QuoteDecimals - spotMarket.BaseDecimals)
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := spotMarket.PriceFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertNotionalFromChainFormatForSpotMarketV2(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV2()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := int32(spotMarket.QuoteDecimals)
	chainFormatPrice := expectedNotional.Mul(decimal.New(1, notionalDecimals))
	humanReadableNotional := spotMarket.NotionalFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

// Derivative markets tests

func TestConvertQuantityToChainFormatForDerivativeMarketV2(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV2()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.QuantityToChainFormat(originalQuantity)
	quantizedValue := originalQuantity.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForDerivativeMarketV2(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV2()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.PriceToChainFormat(originalPrice)
	priceDecimals := int32(derivativeMarket.QuoteDecimals)
	quantizedValue := originalPrice.DivRound(derivativeMarket.MinPriceTickSize, 0).Mul(derivativeMarket.MinPriceTickSize)
	expectedValue := quantizedValue.Mul(decimal.New(1, priceDecimals))
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertMarginToChainFormatForDerivativeMarketV2(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV2()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.MarginToChainFormat(originalPrice)
	marginDecimals := int32(derivativeMarket.QuoteDecimals)
	expectedValue := originalPrice.Mul(decimal.New(1, marginDecimals))
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestCalculateMarginInChainFormatForDerivativeMarketV2(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV2()
	originalQuantity := decimal.RequireFromString("10")
	originalPrice := decimal.RequireFromString("123.456789")
	originalLeverage := decimal.RequireFromString("2.5")

	chainValue := derivativeMarket.CalculateMarginInChainFormat(originalQuantity, originalPrice, originalLeverage)
	decimals := int32(derivativeMarket.QuoteDecimals)
	expectedValue := originalQuantity.Mul(originalPrice).Div(originalLeverage)
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	chainFormatValue := quantizedValue.Mul(decimal.New(1, decimals)).BigInt()
	legacyDecimalQuantizedValue := sdkmath.LegacyMustNewDecFromStr(chainFormatValue.String())

	assert.Assert(t, chainValue.Equal(legacyDecimalQuantizedValue))
}

func TestConvertNotionalToChainFormatForDerivativeMarketV2(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV2()
	originalNotional := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.NotionalToChainFormat(originalNotional)
	notionalDecimals := int32(derivativeMarket.QuoteDecimals)
	expectedValue := originalNotional.Mul(decimal.New(1, notionalDecimals))
	expectedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, expectedChainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForDerivativeMarketV2(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV2()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity
	humanReadableQuantity := derivativeMarket.QuantityFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForDerivativeMarketV2(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV2()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := int32(derivativeMarket.QuoteDecimals)
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := derivativeMarket.PriceFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertMarginFromChainFormatForDerivativeMarketV2(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV2()
	expectedMargin := decimal.RequireFromString("123.456")

	marginDecimals := int32(derivativeMarket.QuoteDecimals)
	chainFormatMargin := expectedMargin.Mul(decimal.New(1, marginDecimals))
	humanReadablePrice := derivativeMarket.MarginFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatMargin.String()))

	assert.Assert(t, expectedMargin.Equal(humanReadablePrice))
}

func TestConvertNotionalFromChainFormatForDerivativeMarketV2(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV2()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := int32(derivativeMarket.QuoteDecimals)
	chainFormatPrice := expectedNotional.Mul(decimal.New(1, notionalDecimals))
	humanReadableNotional := derivativeMarket.NotionalFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

// Binary Option markets tests

func TestConvertQuantityToChainFormatForBinaryOptionMarketV2(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV2()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.QuantityToChainFormat(originalQuantity)
	quantizedValue := originalQuantity.DivRound(binaryOptionMarket.MinQuantityTickSize, 0).Mul(binaryOptionMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForBinaryOptionMarketV2(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV2()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.PriceToChainFormat(originalPrice)
	priceDecimals := int32(binaryOptionMarket.QuoteDecimals)
	quantizedValue := originalPrice.DivRound(binaryOptionMarket.MinPriceTickSize, 0).Mul(binaryOptionMarket.MinPriceTickSize)
	expectedValue := quantizedValue.Mul(decimal.New(1, priceDecimals))
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertMarginToChainFormatForBinaryOptionMarketV2(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV2()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.MarginToChainFormat(originalPrice)
	marginDecimals := int32(binaryOptionMarket.QuoteDecimals)
	expectedValue := originalPrice.Mul(decimal.New(1, marginDecimals))
	quantizedValue := expectedValue.DivRound(binaryOptionMarket.MinQuantityTickSize, 0).Mul(binaryOptionMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestCalculateMarginInChainFormatForBinaryOptionMarketV2(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV2()
	originalQuantity := decimal.RequireFromString("10")
	originalPrice := decimal.RequireFromString("123.456789")
	originalLeverage := decimal.RequireFromString("2.5")

	chainValue := binaryOptionMarket.CalculateMarginInChainFormat(originalQuantity, originalPrice, originalLeverage)
	decimals := int32(binaryOptionMarket.QuoteDecimals)
	expectedValue := originalQuantity.Mul(originalPrice).Div(originalLeverage)
	quantizedValue := expectedValue.DivRound(binaryOptionMarket.MinQuantityTickSize, 0).Mul(binaryOptionMarket.MinQuantityTickSize)
	chainFormatValue := quantizedValue.Mul(decimal.New(1, decimals)).BigInt()
	legacyDecimalQuantizedValue := sdkmath.LegacyMustNewDecFromStr(chainFormatValue.String())

	assert.Assert(t, chainValue.Equal(legacyDecimalQuantizedValue))
}

func TestConvertNotionalToChainFormatForBinaryOptionMarketV2(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV2()
	originalNotional := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.NotionalToChainFormat(originalNotional)
	notionalDecimals := int32(binaryOptionMarket.QuoteDecimals)
	expectedValue := originalNotional.Mul(decimal.New(1, notionalDecimals))
	expectedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, expectedChainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForBinaryOptionMarketV2(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV2()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity
	humanReadableQuantity := binaryOptionMarket.QuantityFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForBinaryOptionMarketV2(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV2()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := int32(binaryOptionMarket.QuoteDecimals)
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := binaryOptionMarket.PriceFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertMarginFromChainFormatForBinaryOptionMarketV2(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV2()
	expectedMargin := decimal.RequireFromString("123.456")

	marginDecimals := int32(binaryOptionMarket.QuoteDecimals)
	chainFormatMargin := expectedMargin.Mul(decimal.New(1, marginDecimals))
	humanReadablePrice := binaryOptionMarket.MarginFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatMargin.String()))

	assert.Assert(t, expectedMargin.Equal(humanReadablePrice))
}

func TestConvertNotionalFromChainFormatForBinaryOptionMarketV2(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV2()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := int32(binaryOptionMarket.QuoteDecimals)
	chainFormatPrice := expectedNotional.Mul(decimal.New(1, notionalDecimals))
	humanReadableNotional := binaryOptionMarket.NotionalFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

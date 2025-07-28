package core

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	"github.com/huandu/go-assert"
	"github.com/shopspring/decimal"
)

func createINJUSDTSpotMarketV1() SpotMarketV1 {
	injToken := createINJToken()
	usdtToken := createUSDTToken()

	makerFeeRate := decimal.RequireFromString("-0.0001")
	takerFeeRate := decimal.RequireFromString("0.001")
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("0.000000000000001")
	minQuantityTickSize := decimal.RequireFromString("1000000000000000")
	minNotional := decimal.RequireFromString("1000000")

	market := SpotMarketV1{
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

func createBTCUSDTPerpMarketV1() DerivativeMarketV1 {
	usdtToken := createUSDTToken()

	initialMarginRatio := decimal.RequireFromString("0.095")
	maintenanceMarginRatio := decimal.RequireFromString("0.025")
	makerFeeRate := decimal.RequireFromString("-0.0001")
	takerFeeRate := decimal.RequireFromString("0.001")
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("1000000")
	minQuantityTickSize := decimal.RequireFromString("0.0001")
	minNotional := decimal.RequireFromString("1000000")

	market := DerivativeMarketV1{
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
		MinNotional:            minNotional,
		QuoteDecimals:          6,
	}
	return market
}

func createBetBinaryOptionMarketV1() BinaryOptionMarketV1 {
	usdtToken := createUSDTToken()

	makerFeeRate := decimal.Zero
	takerFeeRate := decimal.Zero
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("10000")
	minQuantityTickSize := decimal.RequireFromString("1")
	minNotional := decimal.RequireFromString("0.00001")

	market := BinaryOptionMarketV1{
		Id:                  "0x230dcce315364ff6360097838701b14713e2f4007d704df20ed3d81d09eec957",
		Status:              "active",
		Ticker:              "5fdbe0b1-1707800399-WAS",
		OracleSymbol:        "Frontrunner",
		OracleProvider:      "Frontrunner",
		OracleType:          "provider",
		OracleScaleFactor:   6,
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

func TestConvertQuantityToChainFormatForSpotMarketV1(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV1()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.QuantityToChainFormat(originalQuantity)
	expectedValue := originalQuantity.Mul(decimal.New(1, int32(spotMarket.BaseDecimals)))
	quantizedValue := expectedValue.DivRound(spotMarket.MinQuantityTickSize, 0).Mul(spotMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.AssertEqual(t, quantizedChainFormatValue, chainValue)
}

func TestConvertPriceToChainFormatForSpotMarketV1(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV1()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.PriceToChainFormat(originalPrice)
	priceDecimals := int32(spotMarket.QuoteDecimals - spotMarket.BaseDecimals)
	expectedValue := originalPrice.Mul(decimal.New(1, priceDecimals))
	quantizedValue := expectedValue.DivRound(spotMarket.MinPriceTickSize, 0).Mul(spotMarket.MinPriceTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertNotionalToChainFormatForSpotMarketV1(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV1()
	originalNotional := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.NotionalToChainFormat(originalNotional)
	notionalDecimals := spotMarket.QuoteDecimals
	expectedValue := originalNotional.Mul(decimal.New(1, int32(notionalDecimals)))
	chainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, chainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForSpotMarketV1(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV1()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity.Mul(decimal.New(1, int32(spotMarket.BaseDecimals)))
	humanReadableQuantity := spotMarket.QuantityFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForSpotMarketV1(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV1()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := int32(spotMarket.QuoteDecimals - spotMarket.BaseDecimals)
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := spotMarket.PriceFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertNotionalFromChainFormatForSpotMarketV1(t *testing.T) {
	spotMarket := createINJUSDTSpotMarketV1()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := int32(spotMarket.QuoteDecimals)
	chainFormatPrice := expectedNotional.Mul(decimal.New(1, notionalDecimals))
	humanReadableNotional := spotMarket.NotionalFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

// Derivative markets tests

func TestConvertQuantityToChainFormatForDerivativeMarketV1(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV1()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.QuantityToChainFormat(originalQuantity)
	quantizedValue := originalQuantity.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForDerivativeMarketV1(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV1()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.PriceToChainFormat(originalPrice)
	priceDecimals := int32(derivativeMarket.QuoteDecimals)
	expectedValue := originalPrice.Mul(decimal.New(1, priceDecimals))
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinPriceTickSize, 0).Mul(derivativeMarket.MinPriceTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertMarginToChainFormatForDerivativeMarketV1(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV1()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.MarginToChainFormat(originalPrice)
	marginDecimals := int32(derivativeMarket.QuoteDecimals)
	expectedValue := originalPrice.Mul(decimal.New(1, marginDecimals))
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestCalculateMarginInChainFormatForDerivativeMarketV1(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV1()
	originalQuantity := decimal.RequireFromString("10")
	originalPrice := decimal.RequireFromString("123.456789")
	originalLeverage := decimal.RequireFromString("2.5")

	chainValue := derivativeMarket.CalculateMarginInChainFormat(originalQuantity, originalPrice, originalLeverage)
	decimals := int32(derivativeMarket.QuoteDecimals)
	expectedValue := originalQuantity.Mul(originalPrice).Div(originalLeverage).Mul(decimal.New(1, decimals))
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	legacyDecimalQuantizedValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, chainValue.Equal(legacyDecimalQuantizedValue))
}

func TestConvertNotionalToChainFormatForDerivativeMarketV1(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV1()
	originalNotional := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.NotionalToChainFormat(originalNotional)
	notionalDecimals := int32(derivativeMarket.QuoteDecimals)
	expectedValue := originalNotional.Mul(decimal.New(1, notionalDecimals))
	expectedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, expectedChainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForDerivativeMarketV1(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV1()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity
	humanReadableQuantity := derivativeMarket.QuantityFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForDerivativeMarketV1(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV1()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := int32(derivativeMarket.QuoteDecimals)
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := derivativeMarket.PriceFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertMarginFromChainFormatForDerivativeMarketV1(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV1()
	expectedMargin := decimal.RequireFromString("123.456")

	marginDecimals := int32(derivativeMarket.QuoteDecimals)
	chainFormatMargin := expectedMargin.Mul(decimal.New(1, marginDecimals))
	humanReadablePrice := derivativeMarket.MarginFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatMargin.String()))

	assert.Assert(t, expectedMargin.Equal(humanReadablePrice))
}

func TestConvertNotionalFromChainFormatForDerivativeMarketV1(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarketV1()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := int32(derivativeMarket.QuoteDecimals)
	chainFormatPrice := expectedNotional.Mul(decimal.New(1, notionalDecimals))
	humanReadableNotional := derivativeMarket.NotionalFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

// Binary Option markets tests

func TestConvertQuantityToChainFormatForBinaryOptionMarketV1(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV1()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.QuantityToChainFormat(originalQuantity)
	quantizedValue := originalQuantity.DivRound(binaryOptionMarket.MinQuantityTickSize, 0).Mul(binaryOptionMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForBinaryOptionMarketV1(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV1()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.PriceToChainFormat(originalPrice)
	priceDecimals := int32(binaryOptionMarket.QuoteDecimals)
	expectedValue := originalPrice.Mul(decimal.New(1, priceDecimals))
	quantizedValue := expectedValue.DivRound(binaryOptionMarket.MinPriceTickSize, 0).Mul(binaryOptionMarket.MinPriceTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertMarginToChainFormatForBinaryOptionMarketV1(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV1()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.MarginToChainFormat(originalPrice)
	marginDecimals := int32(binaryOptionMarket.QuoteDecimals)
	expectedValue := originalPrice.Mul(decimal.New(1, marginDecimals))
	quantizedValue := expectedValue.DivRound(binaryOptionMarket.MinQuantityTickSize, 0).Mul(binaryOptionMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestCalculateMarginInChainFormatForBinaryOptionMarketV1(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV1()
	originalQuantity := decimal.RequireFromString("10")
	originalPrice := decimal.RequireFromString("123.456789")
	originalLeverage := decimal.RequireFromString("2.5")

	chainValue := binaryOptionMarket.CalculateMarginInChainFormat(originalQuantity, originalPrice, originalLeverage)
	decimals := int32(binaryOptionMarket.QuoteDecimals)
	expectedValue := originalQuantity.Mul(originalPrice).Div(originalLeverage).Mul(decimal.New(1, decimals))
	quantizedValue := expectedValue.DivRound(binaryOptionMarket.MinQuantityTickSize, 0).Mul(binaryOptionMarket.MinQuantityTickSize)
	legacyDecimalQuantizedValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, chainValue.Equal(legacyDecimalQuantizedValue))
}

func TestConvertNotionalToChainFormatForBinaryOptionMarketV1(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV1()
	originalNotional := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.NotionalToChainFormat(originalNotional)
	notionalDecimals := int32(binaryOptionMarket.QuoteDecimals)
	expectedValue := originalNotional.Mul(decimal.New(1, notionalDecimals))
	expectedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, expectedChainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForBinaryOptionMarketV1(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV1()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity
	humanReadableQuantity := binaryOptionMarket.QuantityFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForBinaryOptionMarketV1(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV1()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := int32(binaryOptionMarket.QuoteDecimals)
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := binaryOptionMarket.PriceFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertMarginFromChainFormatForBinaryOptionMarketV1(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV1()
	expectedMargin := decimal.RequireFromString("123.456")

	marginDecimals := int32(binaryOptionMarket.QuoteDecimals)
	chainFormatMargin := expectedMargin.Mul(decimal.New(1, marginDecimals))
	humanReadablePrice := binaryOptionMarket.MarginFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatMargin.String()))

	assert.Assert(t, expectedMargin.Equal(humanReadablePrice))
}

func TestConvertNotionalFromChainFormatForBinaryOptionMarketV1(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarketV1()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := int32(binaryOptionMarket.QuoteDecimals)
	chainFormatPrice := expectedNotional.Mul(decimal.New(1, notionalDecimals))
	humanReadableNotional := binaryOptionMarket.NotionalFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

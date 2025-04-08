package core

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	"github.com/huandu/go-assert"
	"github.com/shopspring/decimal"
)

func createINJUSDTSpotMarket() SpotMarket {
	injToken := createINJToken()
	usdtToken := createUSDTToken()

	makerFeeRate := decimal.RequireFromString("-0.0001")
	takerFeeRate := decimal.RequireFromString("0.001")
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("0.000000000000001")
	minQuantityTickSize := decimal.RequireFromString("1000000000000000")
	minNotional := decimal.RequireFromString("1000000")

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
		MinNotional:         minNotional,
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
	minNotional := decimal.RequireFromString("1000000")

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
		MinNotional:            minNotional,
	}
	return market
}

func createBetBinaryOptionMarket() BinaryOptionMarket {
	usdtToken := createUSDTToken()

	makerFeeRate := decimal.Zero
	takerFeeRate := decimal.Zero
	serviceProviderFee := decimal.RequireFromString("0.4")
	minPriceTickSize := decimal.RequireFromString("10000")
	minQuantityTickSize := decimal.RequireFromString("1")
	minNotional := decimal.RequireFromString("0.00001")

	market := BinaryOptionMarket{
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
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.PriceToChainFormat(originalPrice)
	priceDecimals := spotMarket.QuoteToken.Decimals - spotMarket.BaseToken.Decimals
	expectedValue := originalPrice.Mul(decimal.New(1, priceDecimals))
	quantizedValue := expectedValue.DivRound(spotMarket.MinPriceTickSize, 0).Mul(spotMarket.MinPriceTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertNotionalToChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	originalNotional := decimal.RequireFromString("123.456789")

	chainValue := spotMarket.NotionalToChainFormat(originalNotional)
	notionalDecimals := spotMarket.QuoteToken.Decimals
	expectedValue := originalNotional.Mul(decimal.New(1, notionalDecimals))
	chainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, chainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity.Mul(decimal.New(1, spotMarket.BaseToken.Decimals))
	humanReadableQuantity := spotMarket.QuantityFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := spotMarket.QuoteToken.Decimals - spotMarket.BaseToken.Decimals
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := spotMarket.PriceFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertNotionalFromChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := spotMarket.QuoteToken.Decimals
	chainFormatPrice := expectedNotional.Mul(decimal.New(1, notionalDecimals))
	humanReadableNotional := spotMarket.NotionalFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

func TestConvertQuantityFromExtendedChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity.Mul(decimal.New(1, spotMarket.BaseToken.Decimals)).Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadableQuantity := spotMarket.QuantityFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromExtendedChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := spotMarket.QuoteToken.Decimals - spotMarket.BaseToken.Decimals
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals)).Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadablePrice := spotMarket.PriceFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertNotionalFromExtendedChainFormatForSpotMarket(t *testing.T) {
	spotMarket := createINJUSDTSpotMarket()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := spotMarket.QuoteToken.Decimals
	chainFormatNotional := expectedNotional.Mul(decimal.New(1, notionalDecimals)).Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadableNotional := spotMarket.NotionalFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatNotional.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

// Derivative markets tests

func TestConvertQuantityToChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.QuantityToChainFormat(originalQuantity)
	quantizedValue := originalQuantity.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.PriceToChainFormat(originalPrice)
	priceDecimals := derivativeMarket.QuoteToken.Decimals
	expectedValue := originalPrice.Mul(decimal.New(1, priceDecimals))
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinPriceTickSize, 0).Mul(derivativeMarket.MinPriceTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertMarginToChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.MarginToChainFormat(originalPrice)
	marginDecimals := derivativeMarket.QuoteToken.Decimals
	expectedValue := originalPrice.Mul(decimal.New(1, marginDecimals))
	quantizedValue := expectedValue.DivRound(derivativeMarket.MinQuantityTickSize, 0).Mul(derivativeMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

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
	legacyDecimalQuantizedValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, chainValue.Equal(legacyDecimalQuantizedValue))
}

func TestConvertNotionalToChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	originalNotional := decimal.RequireFromString("123.456789")

	chainValue := derivativeMarket.NotionalToChainFormat(originalNotional)
	notionalDecimals := derivativeMarket.QuoteToken.Decimals
	expectedValue := originalNotional.Mul(decimal.New(1, notionalDecimals))
	expectedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, expectedChainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity
	humanReadableQuantity := derivativeMarket.QuantityFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := derivativeMarket.QuoteToken.Decimals
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := derivativeMarket.PriceFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertMarginFromChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedMargin := decimal.RequireFromString("123.456")

	marginDecimals := derivativeMarket.QuoteToken.Decimals
	chainFormatMargin := expectedMargin.Mul(decimal.New(1, marginDecimals))
	humanReadablePrice := derivativeMarket.MarginFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatMargin.String()))

	assert.Assert(t, expectedMargin.Equal(humanReadablePrice))
}

func TestConvertNotionalFromChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := derivativeMarket.QuoteToken.Decimals
	chainFormatPrice := expectedNotional.Mul(decimal.New(1, notionalDecimals))
	humanReadableNotional := derivativeMarket.NotionalFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

func TestConvertQuantityFromExtendedChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity.Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadableQuantity := derivativeMarket.QuantityFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromExtendedChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := derivativeMarket.QuoteToken.Decimals
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals)).Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadablePrice := derivativeMarket.PriceFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertMarginFromExtendedChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedMargin := decimal.RequireFromString("123.456")

	marginDecimals := derivativeMarket.QuoteToken.Decimals
	chainFormatMargin := expectedMargin.Mul(decimal.New(1, marginDecimals)).Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadablePrice := derivativeMarket.MarginFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatMargin.String()))

	assert.Assert(t, expectedMargin.Equal(humanReadablePrice))
}

func TestConvertNotionalFromExtendedChainFormatForDerivativeMarket(t *testing.T) {
	derivativeMarket := createBTCUSDTPerpMarket()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := derivativeMarket.QuoteToken.Decimals
	chainFormatNotional := expectedNotional.Mul(decimal.New(1, notionalDecimals)).Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadableNotional := derivativeMarket.NotionalFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatNotional.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

// Binary Option markets tests

func TestConvertQuantityToChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	originalQuantity := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.QuantityToChainFormat(originalQuantity)
	quantizedValue := originalQuantity.DivRound(binaryOptionMarket.MinQuantityTickSize, 0).Mul(binaryOptionMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertPriceToChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.PriceToChainFormat(originalPrice)
	priceDecimals := binaryOptionMarket.QuoteToken.Decimals
	expectedValue := originalPrice.Mul(decimal.New(1, priceDecimals))
	quantizedValue := expectedValue.DivRound(binaryOptionMarket.MinPriceTickSize, 0).Mul(binaryOptionMarket.MinPriceTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestConvertMarginToChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	originalPrice := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.MarginToChainFormat(originalPrice)
	marginDecimals := binaryOptionMarket.QuoteToken.Decimals
	expectedValue := originalPrice.Mul(decimal.New(1, marginDecimals))
	quantizedValue := expectedValue.DivRound(binaryOptionMarket.MinQuantityTickSize, 0).Mul(binaryOptionMarket.MinQuantityTickSize)
	quantizedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, quantizedChainFormatValue.Equal(chainValue))
}

func TestCalculateMarginInChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	originalQuantity := decimal.RequireFromString("10")
	originalPrice := decimal.RequireFromString("123.456789")
	originalLeverage := decimal.RequireFromString("2.5")

	chainValue := binaryOptionMarket.CalculateMarginInChainFormat(originalQuantity, originalPrice, originalLeverage)
	decimals := binaryOptionMarket.QuoteToken.Decimals
	expectedValue := originalQuantity.Mul(originalPrice).Div(originalLeverage).Mul(decimal.New(1, decimals))
	quantizedValue := expectedValue.DivRound(binaryOptionMarket.MinQuantityTickSize, 0).Mul(binaryOptionMarket.MinQuantityTickSize)
	legacyDecimalQuantizedValue := sdkmath.LegacyMustNewDecFromStr(quantizedValue.String())

	assert.Assert(t, chainValue.Equal(legacyDecimalQuantizedValue))
}

func TestConvertNotionalToChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	originalNotional := decimal.RequireFromString("123.456789")

	chainValue := binaryOptionMarket.NotionalToChainFormat(originalNotional)
	notionalDecimals := binaryOptionMarket.QuoteToken.Decimals
	expectedValue := originalNotional.Mul(decimal.New(1, notionalDecimals))
	expectedChainFormatValue := sdkmath.LegacyMustNewDecFromStr(expectedValue.String())

	assert.Assert(t, expectedChainFormatValue.Equal(chainValue))
}

func TestConvertQuantityFromChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity
	humanReadableQuantity := binaryOptionMarket.QuantityFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := binaryOptionMarket.QuoteToken.Decimals
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals))
	humanReadablePrice := binaryOptionMarket.PriceFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertMarginFromChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	expectedMargin := decimal.RequireFromString("123.456")

	marginDecimals := binaryOptionMarket.QuoteToken.Decimals
	chainFormatMargin := expectedMargin.Mul(decimal.New(1, marginDecimals))
	humanReadablePrice := binaryOptionMarket.MarginFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatMargin.String()))

	assert.Assert(t, expectedMargin.Equal(humanReadablePrice))
}

func TestConvertNotionalFromChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := binaryOptionMarket.QuoteToken.Decimals
	chainFormatPrice := expectedNotional.Mul(decimal.New(1, notionalDecimals))
	humanReadableNotional := binaryOptionMarket.NotionalFromChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

func TestConvertQuantityFromExtendedChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	expectedQuantity := decimal.RequireFromString("123.456")

	chainFormatQuantity := expectedQuantity.Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadableQuantity := binaryOptionMarket.QuantityFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatQuantity.String()))

	assert.Assert(t, expectedQuantity.Equal(humanReadableQuantity))
}

func TestConvertPriceFromExtendedChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	expectedPrice := decimal.RequireFromString("123.456")

	priceDecimals := binaryOptionMarket.QuoteToken.Decimals
	chainFormatPrice := expectedPrice.Mul(decimal.New(1, priceDecimals)).Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadablePrice := binaryOptionMarket.PriceFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatPrice.String()))

	assert.Assert(t, expectedPrice.Equal(humanReadablePrice))
}

func TestConvertMarginFromExtendedChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	expectedMargin := decimal.RequireFromString("123.456")

	marginDecimals := binaryOptionMarket.QuoteToken.Decimals
	chainFormatMargin := expectedMargin.Mul(decimal.New(1, marginDecimals)).Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadablePrice := binaryOptionMarket.MarginFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatMargin.String()))

	assert.Assert(t, expectedMargin.Equal(humanReadablePrice))
}

func TestConvertNotionalFromExtendedChainFormatForBinaryOptionMarket(t *testing.T) {
	binaryOptionMarket := createBetBinaryOptionMarket()
	expectedNotional := decimal.RequireFromString("123.456")

	notionalDecimals := binaryOptionMarket.QuoteToken.Decimals
	chainFormatNotional := expectedNotional.Mul(decimal.New(1, notionalDecimals)).Mul(decimal.New(1, AdditionalChainFormatDecimals))
	humanReadableNotional := binaryOptionMarket.NotionalFromExtendedChainFormat(sdkmath.LegacyMustNewDecFromStr(chainFormatNotional.String()))

	assert.Assert(t, expectedNotional.Equal(humanReadableNotional))
}

package chain

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	exchangev1types "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	exchangev2types "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
)

func createSmartDenomMetadata() types.Metadata {
	firstDenomUnit := types.DenomUnit{
		Denom:    "factory/inj105ujajd95znwjvcy3hwcz80pgy8tc6v77spur0/SMART",
		Exponent: 0,
		Aliases:  []string{"microSMART"},
	}
	secondDenomUnit := types.DenomUnit{
		Denom:    "SMART",
		Exponent: 6,
		Aliases:  []string{"SMART"},
	}
	metadata := types.Metadata{
		Description: "SMART",
		DenomUnits:  []*types.DenomUnit{&firstDenomUnit, &secondDenomUnit},
		Base:        "factory/inj105ujajd95znwjvcy3hwcz80pgy8tc6v77spur0/SMART",
		Display:     "SMART",
		Name:        "SMART",
		Symbol:      "SMART",
		URI:         "https://upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Flag_of_the_People%27s_Republic_of_China.svg/2560px-Flag_of_the_People%27s_Republic_of_China.svg.png",
		URIHash:     "",
	}

	return metadata
}

func createINJUSDTChainSpotMarketV1() *exchangev1types.SpotMarket {
	marketInfo := exchangev1types.SpotMarket{
		Ticker:              "INJ/USDT",
		BaseDenom:           "inj",
		QuoteDenom:          "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5",
		MakerFeeRate:        math.LegacyMustNewDecFromStr("-0.0001"),
		TakerFeeRate:        math.LegacyMustNewDecFromStr("0.001"),
		RelayerFeeShareRate: math.LegacyMustNewDecFromStr("0.4"),
		MarketId:            "0x7a57e705bb4e09c88aecfc295569481dbf2fe1d5efe364651fbe72385938e9b0",
		Status:              exchangev1types.MarketStatus_Active,
		MinPriceTickSize:    math.LegacyMustNewDecFromStr("0.000000000000001"),
		MinQuantityTickSize: math.LegacyMustNewDecFromStr("1000000000000000"),
		MinNotional:         math.LegacyMustNewDecFromStr("1000000"),
		BaseDecimals:        18,
		QuoteDecimals:       6,
	}

	return &marketInfo
}

func createAPEUSDTChainSpotMarketV1() *exchangev1types.SpotMarket {
	marketInfo := exchangev1types.SpotMarket{
		Ticker:              "APE/USDT",
		BaseDenom:           "peggy0x44C21afAaF20c270EBbF5914Cfc3b5022173FEB7",
		QuoteDenom:          "factory/inj10vkkttgxdeqcgeppu20x9qtyvuaxxev8qh0awq/usdt",
		MakerFeeRate:        math.LegacyMustNewDecFromStr("-0.0001"),
		TakerFeeRate:        math.LegacyMustNewDecFromStr("0.001"),
		RelayerFeeShareRate: math.LegacyMustNewDecFromStr("0.4"),
		MarketId:            "0x8b67e705bb4e09c88aecfc295569481dbf2fe1d5efe364651fbe72385938e000",
		Status:              exchangev1types.MarketStatus_Active,
		MinPriceTickSize:    math.LegacyMustNewDecFromStr("0.000000000000001"),
		MinQuantityTickSize: math.LegacyMustNewDecFromStr("1000000000000000"),
		MinNotional:         math.LegacyMustNewDecFromStr("1000000"),
		BaseDecimals:        6,
		QuoteDecimals:       6,
	}

	return &marketInfo
}

func createBTCUSDTChainDerivativeMarketV1() *exchangev1types.DerivativeMarket {
	marketInfo := exchangev1types.DerivativeMarket{
		Ticker:                 "BTC/USDT PERP",
		OracleBase:             "BTC",
		OracleQuote:            "USDT",
		OracleType:             oracletypes.OracleType_Band,
		OracleScaleFactor:      6,
		QuoteDenom:             "peggy0xdAC17F958D2ee523a2206206994597C13D831ec7",
		MarketId:               "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce",
		InitialMarginRatio:     math.LegacyMustNewDecFromStr("0.095"),
		MaintenanceMarginRatio: math.LegacyMustNewDecFromStr("0.025"),
		ReduceMarginRatio:      math.LegacyMustNewDecFromStr("0.01"),
		MakerFeeRate:           math.LegacyMustNewDecFromStr("-0.0001"),
		TakerFeeRate:           math.LegacyMustNewDecFromStr("0.001"),
		RelayerFeeShareRate:    math.LegacyMustNewDecFromStr("0.4"),
		IsPerpetual:            true,
		Status:                 exchangev1types.MarketStatus_Active,
		MinPriceTickSize:       math.LegacyMustNewDecFromStr("1000000"),
		MinQuantityTickSize:    math.LegacyMustNewDecFromStr("0.0001"),
		MinNotional:            math.LegacyMustNewDecFromStr("1000000"),
		QuoteDecimals:          6,
	}

	return &marketInfo
}

func createINJUSDTChainSpotMarketV2() *exchangev2types.SpotMarket {
	marketInfo := exchangev2types.SpotMarket{
		Ticker:              "INJ/USDT",
		BaseDenom:           "inj",
		QuoteDenom:          "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5",
		MakerFeeRate:        math.LegacyMustNewDecFromStr("-0.0001"),
		TakerFeeRate:        math.LegacyMustNewDecFromStr("0.001"),
		RelayerFeeShareRate: math.LegacyMustNewDecFromStr("0.4"),
		MarketId:            "0x7a57e705bb4e09c88aecfc295569481dbf2fe1d5efe364651fbe72385938e9b0",
		Status:              exchangev2types.MarketStatus_Active,
		MinPriceTickSize:    math.LegacyMustNewDecFromStr("0.000000000000001"),
		MinQuantityTickSize: math.LegacyMustNewDecFromStr("1000000000000000"),
		MinNotional:         math.LegacyMustNewDecFromStr("1000000"),
		BaseDecimals:        18,
		QuoteDecimals:       6,
	}

	return &marketInfo
}

func createAPEUSDTChainSpotMarketV2() *exchangev2types.SpotMarket {
	marketInfo := exchangev2types.SpotMarket{
		Ticker:              "APE/USDT",
		BaseDenom:           "peggy0x44C21afAaF20c270EBbF5914Cfc3b5022173FEB7",
		QuoteDenom:          "factory/inj10vkkttgxdeqcgeppu20x9qtyvuaxxev8qh0awq/usdt",
		MakerFeeRate:        math.LegacyMustNewDecFromStr("-0.0001"),
		TakerFeeRate:        math.LegacyMustNewDecFromStr("0.001"),
		RelayerFeeShareRate: math.LegacyMustNewDecFromStr("0.4"),
		MarketId:            "0x8b67e705bb4e09c88aecfc295569481dbf2fe1d5efe364651fbe72385938e000",
		Status:              exchangev2types.MarketStatus_Active,
		MinPriceTickSize:    math.LegacyMustNewDecFromStr("0.000000000000001"),
		MinQuantityTickSize: math.LegacyMustNewDecFromStr("1000000000000000"),
		MinNotional:         math.LegacyMustNewDecFromStr("1000000"),
		BaseDecimals:        6,
		QuoteDecimals:       6,
	}

	return &marketInfo
}

func createBTCUSDTChainDerivativeMarketV2() *exchangev2types.DerivativeMarket {
	marketInfo := exchangev2types.DerivativeMarket{
		Ticker:                 "BTC/USDT PERP",
		OracleBase:             "BTC",
		OracleQuote:            "USDT",
		OracleType:             oracletypes.OracleType_Band,
		OracleScaleFactor:      6,
		QuoteDenom:             "peggy0xdAC17F958D2ee523a2206206994597C13D831ec7",
		MarketId:               "0x4ca0f92fc28be0c9761326016b5a1a2177dd6375558365116b5bdda9abc229ce",
		InitialMarginRatio:     math.LegacyMustNewDecFromStr("0.095"),
		MaintenanceMarginRatio: math.LegacyMustNewDecFromStr("0.025"),
		ReduceMarginRatio:      math.LegacyMustNewDecFromStr("0.01"),
		MakerFeeRate:           math.LegacyMustNewDecFromStr("-0.0001"),
		TakerFeeRate:           math.LegacyMustNewDecFromStr("0.001"),
		RelayerFeeShareRate:    math.LegacyMustNewDecFromStr("0.4"),
		IsPerpetual:            true,
		Status:                 exchangev2types.MarketStatus_Active,
		MinPriceTickSize:       math.LegacyMustNewDecFromStr("1000000"),
		MinQuantityTickSize:    math.LegacyMustNewDecFromStr("0.0001"),
		MinNotional:            math.LegacyMustNewDecFromStr("1000000"),
		QuoteDecimals:          6,
		OpenNotionalCap: exchangev2types.OpenNotionalCap{
			Cap: &exchangev2types.OpenNotionalCap_Uncapped{},
		},
	}

	return &marketInfo
}

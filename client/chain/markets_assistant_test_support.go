package chain

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

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

func createINJUSDTChainSpotMarket() *exchangev2types.SpotMarket {
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
	}

	return &marketInfo
}

func createAPEUSDTChainSpotMarket() *exchangev2types.SpotMarket {
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
	}

	return &marketInfo
}

func createBTCUSDTChainDerivativeMarket() *exchangev2types.DerivativeMarket {
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
		MakerFeeRate:           math.LegacyMustNewDecFromStr("-0.0001"),
		TakerFeeRate:           math.LegacyMustNewDecFromStr("0.001"),
		RelayerFeeShareRate:    math.LegacyMustNewDecFromStr("0.4"),
		IsPerpetual:            true,
		Status:                 exchangev2types.MarketStatus_Active,
		MinPriceTickSize:       math.LegacyMustNewDecFromStr("1000000"),
		MinQuantityTickSize:    math.LegacyMustNewDecFromStr("0.0001"),
		MinNotional:            math.LegacyMustNewDecFromStr("1000000"),
	}

	return &marketInfo
}

func createFirstMatchBetBinaryOptionsMarket() *exchangev2types.BinaryOptionsMarket {
	market := exchangev2types.BinaryOptionsMarket{
		Ticker:              "5fdbe0b1-1707800399-WAS",
		OracleSymbol:        "Frontrunner",
		OracleProvider:      "Frontrunner",
		OracleType:          oracletypes.OracleType_Provider,
		OracleScaleFactor:   6,
		ExpirationTimestamp: 1708099200,
		SettlementTimestamp: 1707099200,
		Admin:               "inj1zlh5sqevkfphtwnu9cul8p89vseme2eqt0snn9",
		QuoteDenom:          "peggy0x87aB3B4C8661e07D6372361211B96ed4Dc36B1B5",
		MarketId:            "0x230dcce315364ff6360097838701b14713e2f4007d704df20ed3d81d09eec957",
		MakerFeeRate:        math.LegacyMustNewDecFromStr("0"),
		TakerFeeRate:        math.LegacyMustNewDecFromStr("0"),
		RelayerFeeShareRate: math.LegacyMustNewDecFromStr("0.4"),
		Status:              exchangev2types.MarketStatus_Active,
		MinPriceTickSize:    math.LegacyMustNewDecFromStr("0.01"),
		MinQuantityTickSize: math.LegacyMustNewDecFromStr("1"),
		MinNotional:         math.LegacyMustNewDecFromStr("1"),
		AdminPermissions:    1,
	}

	return &market
}

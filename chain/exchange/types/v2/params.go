package v2

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

var _ paramtypes.ParamSet = &Params{}

// Parameter keys
var (
	KeySpotMarketInstantListingFee                 = []byte("SpotMarketInstantListingFee")
	KeyDerivativeMarketInstantListingFee           = []byte("DerivativeMarketInstantListingFee")
	KeyDefaultSpotMakerFeeRate                     = []byte("DefaultSpotMakerFeeRate")
	KeyDefaultSpotTakerFeeRate                     = []byte("DefaultSpotTakerFeeRate")
	KeyDefaultDerivativeMakerFeeRate               = []byte("DefaultDerivativeMakerFeeRate")
	KeyDefaultDerivativeTakerFeeRate               = []byte("DefaultDerivativeTakerFeeRate")
	KeyDefaultInitialMarginRatio                   = []byte("DefaultInitialMarginRatio")
	KeyDefaultMaintenanceMarginRatio               = []byte("DefaultMaintenanceMarginRatio")
	KeyDefaultFundingInterval                      = []byte("DefaultFundingInterval")
	KeyFundingMultiple                             = []byte("FundingMultiple")
	KeyRelayerFeeShareRate                         = []byte("RelayerFeeShareRate")
	KeyDefaultHourlyFundingRateCap                 = []byte("DefaultHourlyFundingRateCap")
	KeyDefaultHourlyInterestRate                   = []byte("DefaultHourlyInterestRate")
	KeyMaxDerivativeOrderSideCount                 = []byte("MaxDerivativeOrderSideCount")
	KeyInjRewardStakedRequirementThreshold         = []byte("KeyInjRewardStakedRequirementThreshold")
	KeyTradingRewardsVestingDuration               = []byte("TradingRewardsVestingDuration")
	KeyLiquidatorRewardShareRate                   = []byte("LiquidatorRewardShareRate")
	KeyBinaryOptionsMarketInstantListingFee        = []byte("BinaryOptionsMarketInstantListingFee")
	KeyAtomicMarketOrderAccessLevel                = []byte("AtomicMarketOrderAccessLevel")
	KeySpotAtomicMarketOrderFeeMultiplier          = []byte("SpotAtomicMarketOrderFeeMultiplier")
	KeyDerivativeAtomicMarketOrderFeeMultiplier    = []byte("DerivativeAtomicMarketOrderFeeMultiplier")
	KeyBinaryOptionsAtomicMarketOrderFeeMultiplier = []byte("BinaryOptionsAtomicMarketOrderFeeMultiplier")
	KeyMinimalProtocolFeeRate                      = []byte("MinimalProtocolFeeRate")
	KeyIsInstantDerivativeMarketLaunchEnabled      = []byte("IsInstantDerivativeMarketLaunchEnabled")
	KeyPostOnlyModeHeightThreshold                 = []byte("PostOnlyModeHeightThreshold")
)

func NewV1ExchangeParamsFromV2(params Params) types.Params {
	return types.Params{
		SpotMarketInstantListingFee:                  params.SpotMarketInstantListingFee,
		DerivativeMarketInstantListingFee:            params.DerivativeMarketInstantListingFee,
		DefaultSpotMakerFeeRate:                      params.DefaultSpotMakerFeeRate,
		DefaultSpotTakerFeeRate:                      params.DefaultSpotTakerFeeRate,
		DefaultDerivativeMakerFeeRate:                params.DefaultDerivativeMakerFeeRate,
		DefaultDerivativeTakerFeeRate:                params.DefaultDerivativeTakerFeeRate,
		DefaultInitialMarginRatio:                    params.DefaultInitialMarginRatio,
		DefaultMaintenanceMarginRatio:                params.DefaultMaintenanceMarginRatio,
		DefaultFundingInterval:                       params.DefaultFundingInterval,
		FundingMultiple:                              params.FundingMultiple,
		RelayerFeeShareRate:                          params.RelayerFeeShareRate,
		DefaultHourlyFundingRateCap:                  params.DefaultHourlyFundingRateCap,
		DefaultHourlyInterestRate:                    params.DefaultHourlyInterestRate,
		MaxDerivativeOrderSideCount:                  params.MaxDerivativeOrderSideCount,
		InjRewardStakedRequirementThreshold:          params.InjRewardStakedRequirementThreshold,
		TradingRewardsVestingDuration:                params.TradingRewardsVestingDuration,
		LiquidatorRewardShareRate:                    params.LiquidatorRewardShareRate,
		BinaryOptionsMarketInstantListingFee:         params.BinaryOptionsMarketInstantListingFee,
		AtomicMarketOrderAccessLevel:                 types.AtomicMarketOrderAccessLevel(params.AtomicMarketOrderAccessLevel),
		SpotAtomicMarketOrderFeeMultiplier:           params.SpotAtomicMarketOrderFeeMultiplier,
		DerivativeAtomicMarketOrderFeeMultiplier:     params.DerivativeAtomicMarketOrderFeeMultiplier,
		BinaryOptionsAtomicMarketOrderFeeMultiplier:  params.BinaryOptionsAtomicMarketOrderFeeMultiplier,
		MinimalProtocolFeeRate:                       params.MinimalProtocolFeeRate,
		IsInstantDerivativeMarketLaunchEnabled:       params.IsInstantDerivativeMarketLaunchEnabled,
		PostOnlyModeHeightThreshold:                  params.PostOnlyModeHeightThreshold,
		MarginDecreasePriceTimestampThresholdSeconds: params.MarginDecreasePriceTimestampThresholdSeconds,
		ExchangeAdmins:                               params.ExchangeAdmins,
		InjAuctionMaxCap:                             params.InjAuctionMaxCap,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySpotMarketInstantListingFee, &p.SpotMarketInstantListingFee, types.ValidateSpotMarketInstantListingFee),
		paramtypes.NewParamSetPair(KeyDerivativeMarketInstantListingFee, &p.DerivativeMarketInstantListingFee, types.ValidateDerivativeMarketInstantListingFee),
		paramtypes.NewParamSetPair(KeyDefaultSpotMakerFeeRate, &p.DefaultSpotMakerFeeRate, types.ValidateMakerFee),
		paramtypes.NewParamSetPair(KeyDefaultSpotTakerFeeRate, &p.DefaultSpotTakerFeeRate, types.ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultDerivativeMakerFeeRate, &p.DefaultDerivativeMakerFeeRate, types.ValidateMakerFee),
		paramtypes.NewParamSetPair(KeyDefaultDerivativeTakerFeeRate, &p.DefaultDerivativeTakerFeeRate, types.ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultInitialMarginRatio, &p.DefaultInitialMarginRatio, types.ValidateMarginRatio),
		paramtypes.NewParamSetPair(KeyDefaultMaintenanceMarginRatio, &p.DefaultMaintenanceMarginRatio, types.ValidateMarginRatio),
		paramtypes.NewParamSetPair(KeyDefaultFundingInterval, &p.DefaultFundingInterval, types.ValidateFundingInterval),
		paramtypes.NewParamSetPair(KeyFundingMultiple, &p.FundingMultiple, types.ValidateFundingMultiple),
		paramtypes.NewParamSetPair(KeyRelayerFeeShareRate, &p.RelayerFeeShareRate, types.ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultHourlyFundingRateCap, &p.DefaultHourlyFundingRateCap, types.ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultHourlyInterestRate, &p.DefaultHourlyInterestRate, types.ValidateFee),
		paramtypes.NewParamSetPair(KeyMaxDerivativeOrderSideCount, &p.MaxDerivativeOrderSideCount, types.ValidateDerivativeOrderSideCount),
		paramtypes.NewParamSetPair(KeyInjRewardStakedRequirementThreshold, &p.InjRewardStakedRequirementThreshold, types.ValidateInjRewardStakedRequirementThreshold),
		paramtypes.NewParamSetPair(KeyTradingRewardsVestingDuration, &p.TradingRewardsVestingDuration, types.ValidateTradingRewardsVestingDuration),
		paramtypes.NewParamSetPair(KeyLiquidatorRewardShareRate, &p.LiquidatorRewardShareRate, types.ValidateLiquidatorRewardShareRate),
		paramtypes.NewParamSetPair(KeyBinaryOptionsMarketInstantListingFee, &p.BinaryOptionsMarketInstantListingFee, types.ValidateBinaryOptionsMarketInstantListingFee),
		paramtypes.NewParamSetPair(KeyAtomicMarketOrderAccessLevel, &p.AtomicMarketOrderAccessLevel, ValidateAtomicMarketOrderAccessLevel),
		paramtypes.NewParamSetPair(KeySpotAtomicMarketOrderFeeMultiplier, &p.SpotAtomicMarketOrderFeeMultiplier, types.ValidateAtomicMarketOrderFeeMultiplier),
		paramtypes.NewParamSetPair(KeyDerivativeAtomicMarketOrderFeeMultiplier, &p.DerivativeAtomicMarketOrderFeeMultiplier, types.ValidateAtomicMarketOrderFeeMultiplier),
		paramtypes.NewParamSetPair(KeyBinaryOptionsAtomicMarketOrderFeeMultiplier, &p.BinaryOptionsAtomicMarketOrderFeeMultiplier, types.ValidateAtomicMarketOrderFeeMultiplier),
		paramtypes.NewParamSetPair(KeyMinimalProtocolFeeRate, &p.MinimalProtocolFeeRate, types.ValidateFee),
		paramtypes.NewParamSetPair(KeyIsInstantDerivativeMarketLaunchEnabled, &p.IsInstantDerivativeMarketLaunchEnabled, types.ValidateBool),
		paramtypes.NewParamSetPair(KeyPostOnlyModeHeightThreshold, &p.PostOnlyModeHeightThreshold, types.ValidatePostOnlyModeHeightThreshold),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		SpotMarketInstantListingFee:                  sdk.NewCoin("inj", math.NewIntWithDecimal(types.SpotMarketInstantListingFee, 18)),
		DerivativeMarketInstantListingFee:            sdk.NewCoin("inj", math.NewIntWithDecimal(types.DerivativeMarketInstantListingFee, 18)),
		DefaultSpotMakerFeeRate:                      math.LegacyNewDecWithPrec(-1, 4), // default -0.01% maker fees
		DefaultSpotTakerFeeRate:                      math.LegacyNewDecWithPrec(1, 3),  // default 0.1% taker fees
		DefaultDerivativeMakerFeeRate:                math.LegacyNewDecWithPrec(-1, 4), // default -0.01% maker fees
		DefaultDerivativeTakerFeeRate:                math.LegacyNewDecWithPrec(1, 3),  // default 0.1% taker fees
		DefaultInitialMarginRatio:                    math.LegacyNewDecWithPrec(5, 2),  // default 5% initial margin ratio
		DefaultMaintenanceMarginRatio:                math.LegacyNewDecWithPrec(2, 2),  // default 2% maintenance margin ratio
		DefaultFundingInterval:                       types.DefaultFundingIntervalSeconds,
		FundingMultiple:                              types.DefaultFundingMultipleSeconds,
		RelayerFeeShareRate:                          math.LegacyNewDecWithPrec(40, 2),      // default 40% relayer fee share
		DefaultHourlyFundingRateCap:                  math.LegacyNewDecWithPrec(625, 6),     // default 0.0625% max hourly funding rate
		DefaultHourlyInterestRate:                    math.LegacyNewDecWithPrec(416666, 11), // 0.01% daily interest rate = 0.0001 / 24 = 0.00000416666
		MaxDerivativeOrderSideCount:                  types.MaxDerivativeOrderSideCount,
		InjRewardStakedRequirementThreshold:          math.NewIntWithDecimal(100, 18), // 100 INJ
		TradingRewardsVestingDuration:                604800,                          // 7 days
		LiquidatorRewardShareRate:                    math.LegacyNewDecWithPrec(5, 2), // 5% liquidator reward
		BinaryOptionsMarketInstantListingFee:         sdk.NewCoin("inj", math.NewIntWithDecimal(types.BinaryOptionsMarketInstantListingFee, 18)),
		AtomicMarketOrderAccessLevel:                 AtomicMarketOrderAccessLevel_SmartContractsOnly,
		SpotAtomicMarketOrderFeeMultiplier:           math.LegacyNewDecWithPrec(25, 1),        // default 2.5 multiplier
		DerivativeAtomicMarketOrderFeeMultiplier:     math.LegacyNewDecWithPrec(25, 1),        // default 2.5 multiplier
		BinaryOptionsAtomicMarketOrderFeeMultiplier:  math.LegacyNewDecWithPrec(25, 1),        // default 2.5 multiplier
		MinimalProtocolFeeRate:                       math.LegacyMustNewDecFromStr("0.00005"), // default 0.005% minimal fee rate
		IsInstantDerivativeMarketLaunchEnabled:       false,
		PostOnlyModeHeightThreshold:                  0,
		MarginDecreasePriceTimestampThresholdSeconds: 60,
		ExchangeAdmins:                               []string{},
		InjAuctionMaxCap:                             types.DefaultInjAuctionMaxCap,
	}
}

// Validate performs basic validation on exchange parameters.
func (p Params) Validate() error {
	if err := types.ValidateSpotMarketInstantListingFee(p.SpotMarketInstantListingFee); err != nil {
		return fmt.Errorf("spot_market_instant_listing_fee is incorrect: %w", err)
	}
	if err := types.ValidateDerivativeMarketInstantListingFee(p.DerivativeMarketInstantListingFee); err != nil {
		return fmt.Errorf("derivative_market_instant_listing_fee is incorrect: %w", err)
	}
	if err := types.ValidateMakerFee(p.DefaultSpotMakerFeeRate); err != nil {
		return fmt.Errorf("default_spot_maker_fee_rate is incorrect: %w", err)
	}
	if err := types.ValidateFee(p.DefaultSpotTakerFeeRate); err != nil {
		return fmt.Errorf("default_spot_taker_fee_rate is incorrect: %w", err)
	}
	if err := types.ValidateMakerFee(p.DefaultDerivativeMakerFeeRate); err != nil {
		return fmt.Errorf("default_derivative_maker_fee_rate is incorrect: %w", err)
	}
	if err := types.ValidateFee(p.DefaultDerivativeTakerFeeRate); err != nil {
		return fmt.Errorf("default_derivative_taker_fee_rate is incorrect: %w", err)
	}
	if err := types.ValidateMarginRatio(p.DefaultInitialMarginRatio); err != nil {
		return fmt.Errorf("default_initial_margin_ratio is incorrect: %w", err)
	}
	if err := types.ValidateMarginRatio(p.DefaultMaintenanceMarginRatio); err != nil {
		return fmt.Errorf("default_maintenance_margin_ratio is incorrect: %w", err)
	}
	if err := types.ValidateFundingInterval(p.DefaultFundingInterval); err != nil {
		return fmt.Errorf("default_funding_interval is incorrect: %w", err)
	}
	if err := types.ValidateFundingMultiple(p.FundingMultiple); err != nil {
		return fmt.Errorf("funding_multiple is incorrect: %w", err)
	}
	if err := types.ValidateFee(p.RelayerFeeShareRate); err != nil {
		return fmt.Errorf("relayer_fee_share_rate is incorrect: %w", err)
	}
	if err := types.ValidateFee(p.DefaultHourlyFundingRateCap); err != nil {
		return fmt.Errorf("default_hourly_funding_rate_cap is incorrect: %w", err)
	}
	if err := types.ValidateFee(p.DefaultHourlyInterestRate); err != nil {
		return fmt.Errorf("default_hourly_interest_rate is incorrect: %w", err)
	}
	if err := types.ValidateDerivativeOrderSideCount(p.MaxDerivativeOrderSideCount); err != nil {
		return fmt.Errorf("max_derivative_order_side_count is incorrect: %w", err)
	}
	if err := types.ValidateInjRewardStakedRequirementThreshold(p.InjRewardStakedRequirementThreshold); err != nil {
		return fmt.Errorf("inj_reward_staked_requirement_threshold is incorrect: %w", err)
	}
	if err := types.ValidateLiquidatorRewardShareRate(p.LiquidatorRewardShareRate); err != nil {
		return fmt.Errorf("liquidator_reward_share_rate is incorrect: %w", err)
	}
	if err := types.ValidateBinaryOptionsMarketInstantListingFee(p.BinaryOptionsMarketInstantListingFee); err != nil {
		return fmt.Errorf("binary_options_market_instant_listing_fee is incorrect: %w", err)
	}
	if err := ValidateAtomicMarketOrderAccessLevel(p.AtomicMarketOrderAccessLevel); err != nil {
		return fmt.Errorf("atomic_market_order_access_level is incorrect: %w", err)
	}
	if err := types.ValidateAtomicMarketOrderFeeMultiplier(p.SpotAtomicMarketOrderFeeMultiplier); err != nil {
		return fmt.Errorf("spot_atomic_market_order_fee_multiplier is incorrect: %w", err)
	}
	if err := types.ValidateAtomicMarketOrderFeeMultiplier(p.DerivativeAtomicMarketOrderFeeMultiplier); err != nil {
		return fmt.Errorf("derivative_atomic_market_order_fee_multiplier is incorrect: %w", err)
	}
	if err := types.ValidateAtomicMarketOrderFeeMultiplier(p.BinaryOptionsAtomicMarketOrderFeeMultiplier); err != nil {
		return fmt.Errorf("binary_options_atomic_market_order_fee_multiplier is incorrect: %w", err)
	}
	if err := types.ValidateFee(p.MinimalProtocolFeeRate); err != nil {
		return fmt.Errorf("minimal_protocol_fee_rate is incorrect: %w", err)
	}
	if err := types.ValidatePostOnlyModeHeightThreshold(p.PostOnlyModeHeightThreshold); err != nil {
		return fmt.Errorf("post_only_mode_height_threshold is incorrect: %w", err)
	}
	if err := validateAdmins(p.ExchangeAdmins); err != nil {
		return fmt.Errorf("ExchangeAdmins is incorrect: %w", err)
	}
	return nil
}

func validateAdmins(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	admins := make(map[string]struct{})

	for _, admin := range v {
		adminAddr, err := sdk.AccAddressFromBech32(admin)
		if err != nil {
			return fmt.Errorf("invalid admin address: %s", admin)
		}

		if _, found := admins[adminAddr.String()]; found {
			return fmt.Errorf("duplicate admin: %s", admin)
		}
		admins[adminAddr.String()] = struct{}{}
	}

	return nil
}

func ValidateAtomicMarketOrderAccessLevel(i interface{}) error {
	v, ok := i.(AtomicMarketOrderAccessLevel)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !v.IsValid() {
		return fmt.Errorf("invalid AtomicMarketOrderAccessLevel value: %v", v)
	}
	return nil
}

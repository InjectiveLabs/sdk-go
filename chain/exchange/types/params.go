package types

import (
	"errors"
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// Exchange params default values
const (
	// DefaultFundingIntervalSeconds is 3600. This represents the number of seconds in one hour which is the frequency at which
	// funding is applied by default on derivative markets.
	DefaultFundingIntervalSeconds int64 = 3600

	// DefaultFundingMultipleSeconds is 3600. This represents the number of seconds in one hour which is multiple of the
	// unix time seconds timestamp that each perpetual market's funding timestamp should be. This ensures that
	// funding is consistently applied on the hour for all perpetual markets.
	DefaultFundingMultipleSeconds int64 = 3600

	// SpotMarketInstantListingFee is 20 INJ
	SpotMarketInstantListingFee int64 = 20

	// DerivativeMarketInstantListingFee is 20 INJ
	DerivativeMarketInstantListingFee int64 = 20

	// BinaryOptionsMarketInstantListingFee is 100 INJ
	BinaryOptionsMarketInstantListingFee int64 = 100

	MaxOracleScaleFactor uint32 = 18

	MaxDecimals uint32 = 18

	MaxTickerLength int = 40

	MaxOracleSymbolLength int = 256

	MaxOracleProviderLength int = 256

	MaxMarketLaunchDenomLength int = 256

	// MaxHistoricalTradeRecordAge is the maximum age of trade records to track.
	MaxHistoricalTradeRecordAge = 60 * 5

	// MaxSubaccountNonceLength restricts the size of a subaccount number from 0 to 999
	MaxSubaccountNonceLength = 3

	// MaxGranterDelegations is the maximum number of delegations that are checked for stake granter
	MaxGranterDelegations = 25

	// MaxTickSizeDecimalPlaces defines the maximum number of decimal places allowed for tick size
	// The real max for LegacyDec is 18, but we allow 15 to ensure that we are absolutely safe of any rounding issues
	MaxTickSizeDecimalPlaces = 15

	// MaxWhiteKnightLiquidators defines the maximum number of white knight liquidators.
	MaxWhiteKnightLiquidators = 1024
)

var MaxBinaryOptionsOrderPrice = math.LegacyOneDec()

// MaxOrderPrice equals 10^32
var MaxOrderPrice = math.LegacyMustNewDecFromStr("100000000000000000000000000000000")

// MaxOrderMargin equals 10^32
var MaxOrderMargin = math.LegacyMustNewDecFromStr("100000000000000000000000000000000")

// MaxTokenInt equals 100,000,000 * 10^18
var MaxTokenInt, _ = math.NewIntFromString("100000000000000000000000000")

var MaxOrderQuantity = math.LegacyMustNewDecFromStr("100000000000000000000000000000000")
var MaxFeeMultiplier = math.LegacyMustNewDecFromStr("100")

var MinMarginRatio = math.LegacyNewDecWithPrec(5, 3)

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

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySpotMarketInstantListingFee, &p.SpotMarketInstantListingFee, ValidateSpotMarketInstantListingFee),
		paramtypes.NewParamSetPair(
			KeyDerivativeMarketInstantListingFee, &p.DerivativeMarketInstantListingFee, ValidateDerivativeMarketInstantListingFee,
		),
		paramtypes.NewParamSetPair(KeyDefaultSpotMakerFeeRate, &p.DefaultSpotMakerFeeRate, ValidateMakerFee),
		paramtypes.NewParamSetPair(KeyDefaultSpotTakerFeeRate, &p.DefaultSpotTakerFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultDerivativeMakerFeeRate, &p.DefaultDerivativeMakerFeeRate, ValidateMakerFee),
		paramtypes.NewParamSetPair(KeyDefaultDerivativeTakerFeeRate, &p.DefaultDerivativeTakerFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultInitialMarginRatio, &p.DefaultInitialMarginRatio, ValidateMarginRatio),
		paramtypes.NewParamSetPair(KeyDefaultMaintenanceMarginRatio, &p.DefaultMaintenanceMarginRatio, ValidateMarginRatio),
		paramtypes.NewParamSetPair(KeyDefaultFundingInterval, &p.DefaultFundingInterval, ValidateFundingInterval),
		paramtypes.NewParamSetPair(KeyFundingMultiple, &p.FundingMultiple, ValidateFundingMultiple),
		paramtypes.NewParamSetPair(KeyRelayerFeeShareRate, &p.RelayerFeeShareRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultHourlyFundingRateCap, &p.DefaultHourlyFundingRateCap, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultHourlyInterestRate, &p.DefaultHourlyInterestRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyMaxDerivativeOrderSideCount, &p.MaxDerivativeOrderSideCount, ValidateDerivativeOrderSideCount),
		paramtypes.NewParamSetPair(
			KeyInjRewardStakedRequirementThreshold, &p.InjRewardStakedRequirementThreshold, ValidateInjRewardStakedRequirementThreshold,
		),
		paramtypes.NewParamSetPair(
			KeyTradingRewardsVestingDuration, &p.TradingRewardsVestingDuration, ValidateTradingRewardsVestingDuration),
		paramtypes.NewParamSetPair(KeyLiquidatorRewardShareRate, &p.LiquidatorRewardShareRate, ValidateLiquidatorRewardShareRate),
		paramtypes.NewParamSetPair(
			KeyBinaryOptionsMarketInstantListingFee, &p.BinaryOptionsMarketInstantListingFee, ValidateBinaryOptionsMarketInstantListingFee,
		),
		paramtypes.NewParamSetPair(
			KeyAtomicMarketOrderAccessLevel, &p.AtomicMarketOrderAccessLevel, ValidateAtomicMarketOrderAccessLevel,
		),
		paramtypes.NewParamSetPair(
			KeySpotAtomicMarketOrderFeeMultiplier, &p.SpotAtomicMarketOrderFeeMultiplier, ValidateAtomicMarketOrderFeeMultiplier,
		),
		paramtypes.NewParamSetPair(
			KeyDerivativeAtomicMarketOrderFeeMultiplier,
			&p.DerivativeAtomicMarketOrderFeeMultiplier,
			ValidateAtomicMarketOrderFeeMultiplier,
		),
		paramtypes.NewParamSetPair(
			KeyBinaryOptionsAtomicMarketOrderFeeMultiplier,
			&p.BinaryOptionsAtomicMarketOrderFeeMultiplier,
			ValidateAtomicMarketOrderFeeMultiplier,
		),
		paramtypes.NewParamSetPair(KeyMinimalProtocolFeeRate, &p.MinimalProtocolFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyIsInstantDerivativeMarketLaunchEnabled, &p.IsInstantDerivativeMarketLaunchEnabled, ValidateBool),
		paramtypes.NewParamSetPair(KeyPostOnlyModeHeightThreshold, &p.PostOnlyModeHeightThreshold, ValidatePostOnlyModeHeightThreshold),
	}
}

// Validate performs basic validation on exchange parameters.
func (p Params) Validate() error {
	if err := ValidateSpotMarketInstantListingFee(p.SpotMarketInstantListingFee); err != nil {
		return fmt.Errorf("spot_market_instant_listing_fee is incorrect: %w", err)
	}
	if err := ValidateDerivativeMarketInstantListingFee(p.DerivativeMarketInstantListingFee); err != nil {
		return fmt.Errorf("derivative_market_instant_listing_fee is incorrect: %w", err)
	}
	if err := ValidateMakerFee(p.DefaultSpotMakerFeeRate); err != nil {
		return fmt.Errorf("default_spot_maker_fee_rate is incorrect: %w", err)
	}
	if err := ValidateFee(p.DefaultSpotTakerFeeRate); err != nil {
		return fmt.Errorf("default_spot_taker_fee_rate is incorrect: %w", err)
	}
	if err := ValidateMakerFee(p.DefaultDerivativeMakerFeeRate); err != nil {
		return fmt.Errorf("default_derivative_maker_fee_rate is incorrect: %w", err)
	}
	if err := ValidateFee(p.DefaultDerivativeTakerFeeRate); err != nil {
		return fmt.Errorf("default_derivative_taker_fee_rate is incorrect: %w", err)
	}
	if err := ValidateMarginRatio(p.DefaultInitialMarginRatio); err != nil {
		return fmt.Errorf("default_initial_margin_ratio is incorrect: %w", err)
	}
	if err := ValidateMarginRatio(p.DefaultMaintenanceMarginRatio); err != nil {
		return fmt.Errorf("default_maintenance_margin_ratio is incorrect: %w", err)
	}
	if err := ValidateFundingInterval(p.DefaultFundingInterval); err != nil {
		return fmt.Errorf("default_funding_interval is incorrect: %w", err)
	}
	if err := ValidateFundingMultiple(p.FundingMultiple); err != nil {
		return fmt.Errorf("funding_multiple is incorrect: %w", err)
	}
	if err := ValidateFee(p.RelayerFeeShareRate); err != nil {
		return fmt.Errorf("relayer_fee_share_rate is incorrect: %w", err)
	}
	if err := ValidateFee(p.DefaultHourlyFundingRateCap); err != nil {
		return fmt.Errorf("default_hourly_funding_rate_cap is incorrect: %w", err)
	}
	if err := ValidateFee(p.DefaultHourlyInterestRate); err != nil {
		return fmt.Errorf("default_hourly_interest_rate is incorrect: %w", err)
	}
	if err := ValidateDerivativeOrderSideCount(p.MaxDerivativeOrderSideCount); err != nil {
		return fmt.Errorf("max_derivative_order_side_count is incorrect: %w", err)
	}
	if err := ValidateInjRewardStakedRequirementThreshold(p.InjRewardStakedRequirementThreshold); err != nil {
		return fmt.Errorf("inj_reward_staked_requirement_threshold is incorrect: %w", err)
	}
	if err := ValidateLiquidatorRewardShareRate(p.LiquidatorRewardShareRate); err != nil {
		return fmt.Errorf("liquidator_reward_share_rate is incorrect: %w", err)
	}
	if err := ValidateBinaryOptionsMarketInstantListingFee(p.BinaryOptionsMarketInstantListingFee); err != nil {
		return fmt.Errorf("binary_options_market_instant_listing_fee is incorrect: %w", err)
	}
	if err := ValidateAtomicMarketOrderAccessLevel(p.AtomicMarketOrderAccessLevel); err != nil {
		return fmt.Errorf("atomic_market_order_access_level is incorrect: %w", err)
	}
	if err := ValidateAtomicMarketOrderFeeMultiplier(p.SpotAtomicMarketOrderFeeMultiplier); err != nil {
		return fmt.Errorf("spot_atomic_market_order_fee_multiplier is incorrect: %w", err)
	}
	if err := ValidateAtomicMarketOrderFeeMultiplier(p.DerivativeAtomicMarketOrderFeeMultiplier); err != nil {
		return fmt.Errorf("derivative_atomic_market_order_fee_multiplier is incorrect: %w", err)
	}
	if err := ValidateAtomicMarketOrderFeeMultiplier(p.BinaryOptionsAtomicMarketOrderFeeMultiplier); err != nil {
		return fmt.Errorf("binary_options_atomic_market_order_fee_multiplier is incorrect: %w", err)
	}
	if err := ValidateFee(p.MinimalProtocolFeeRate); err != nil {
		return fmt.Errorf("minimal_protocol_fee_rate is incorrect: %w", err)
	}
	if err := ValidatePostOnlyModeHeightThreshold(p.PostOnlyModeHeightThreshold); err != nil {
		return fmt.Errorf("post_only_mode_height_threshold is incorrect: %w", err)
	}
	if err := ValidateAdmins(p.ExchangeAdmins); err != nil {
		return fmt.Errorf("ExchangeAdmins is incorrect: %w", err)
	}

	if err := ValidateFixedGasFlag(p.FixedGasEnabled); err != nil {
		return fmt.Errorf("fixed_gas_enabled is incorrect: %w", err)
	}

	return nil
}

func ValidateFixedGasFlag(enabled any) error {
	if _, ok := enabled.(bool); !ok {
		return fmt.Errorf("invalid parameter type: %T", enabled)
	}

	return nil
}

func ValidateSpotMarketInstantListingFee(i any) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() || !v.Amount.IsPositive() {
		return fmt.Errorf("invalid SpotMarketInstantListingFee: %T", i)
	}

	return nil
}

func ValidateDerivativeMarketInstantListingFee(i any) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() || !v.Amount.IsPositive() {
		return fmt.Errorf("invalid DerivativeMarketInstantListingFee: %T", i)
	}

	return nil
}

func ValidateBinaryOptionsMarketInstantListingFee(i any) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() || !v.Amount.IsPositive() {
		return fmt.Errorf("invalid BinaryOptionsMarketInstantListingFee: %T", i)
	}

	return nil
}

func ValidateFee(i any) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("exchange fee cannot be nil: %s", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("exchange fee cannot be negative: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("exchange fee cannot be greater than 1: %s", v)
	}

	return nil
}

func ValidateNonNegativeDec(i any) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("value cannot be nil: %s", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("value cannot be negative: %s", v)
	}

	return nil
}

func ValidateMakerFee(i any) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("exchange fee cannot be nil: %s", v)
	}

	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("exchange fee cannot be greater than 1: %s", v)
	}

	if v.LT(math.LegacyOneDec().Neg()) {
		return fmt.Errorf("exchange fee cannot be less than -1: %s", v)
	}

	return nil
}

func ValidateHourlyFundingRateCap(i any) error {
	v, ok := i.(math.LegacyDec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("hourly funding rate cap cannot be nil: %s", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("hourly funding rate cap cannot be negative: %s", v)
	}

	if v.IsZero() {
		return fmt.Errorf("hourly funding rate cap cannot be zero: %s", v)
	}

	if v.GT(math.LegacyNewDecWithPrec(3, 2)) {
		return fmt.Errorf("hourly funding rate cap cannot be larger than 3 percent: %s", v)
	}

	return nil
}

func ValidateHourlyInterestRate(i any) error {
	v, ok := i.(math.LegacyDec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("hourly interest rate cannot be nil: %s", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("hourly interest rate cannot be negative: %s", v)
	}

	if v.GT(math.LegacyNewDecWithPrec(1, 2)) {
		return fmt.Errorf("hourly interest rate cannot be larger than 1 percent: %s", v)
	}

	return nil
}

func ValidateTickSize(i any) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("tick size cannot be nil: %s", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("tick size cannot be negative: %s", v)
	}

	if v.IsZero() {
		return fmt.Errorf("tick size cannot be zero: %s", v)
	}

	if v.GT(MaxOrderPrice) {
		return fmt.Errorf("unsupported tick size amount")
	}

	// Use 10^MaxTickSizeDecimalPlaces as scaleFactor to naturally enforce decimal places limit
	// Any tick size with more than MaxTickSizeDecimalPlaces decimal places will result in
	// a scaled value < 1, which cannot be a power of 10
	scaleFactor := math.LegacyNewDec(10).Power(MaxTickSizeDecimalPlaces)
	// v can be a decimal (e.g. 1e-15) so we scale by 10^15
	scaledValue := v.Mul(scaleFactor)

	if !isPowerOf10(scaledValue) {
		return errors.New("unsupported tick size")
	}

	return nil
}

// isPowerOf10 checks whether the given decimal is a positive power of 10 (including 10^0 = 1).
func isPowerOf10(v math.LegacyDec) bool {
	if v.LTE(math.LegacyZeroDec()) {
		return false
	}

	if v.Equal(math.LegacyOneDec()) {
		return true
	}

	temp := v
	ten := math.LegacyNewDec(10)

	// Keep dividing by 10 while the result is >= 10
	for temp.GTE(ten) {
		quotient := temp.Quo(ten)
		// Check if the division was exact (no remainder)
		if !quotient.Mul(ten).Equal(temp) {
			return false
		}
		temp = quotient
	}

	// After all divisions, we should have exactly 1
	return temp.Equal(math.LegacyOneDec())
}

func ValidateMinNotional(i any) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("min notional cannot be nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("min notional cannot be negative: %s", v)
	}

	return nil
}

func ValidateMarginRatio(i any) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("margin ratio cannot be nil: %s", v)
	}
	if v.LT(MinMarginRatio) {
		return fmt.Errorf("margin ratio cannot be less than minimum: %s", v)
	}
	if v.GTE(math.LegacyOneDec()) {
		return fmt.Errorf("margin ratio cannot be greater than or equal to 1: %s", v)
	}

	return nil
}

func ValidateFundingInterval(i any) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("fundingInterval must be positive: %d", v)
	}

	return nil
}

func ValidatePostOnlyModeHeightThreshold(i any) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("postOnlyModeHeightThreshold must be non-negative: %d", v)
	}

	return nil
}

func ValidateAdmins(i any) error {
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

func ValidateWhiteKnightLiquidators(i any) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) > MaxWhiteKnightLiquidators {
		return fmt.Errorf("number of white knight liquidators cannot exceed %d: %d", MaxWhiteKnightLiquidators, len(v))
	}

	liquidators := make(map[string]struct{})

	for _, liquidator := range v {
		liquidatorAddr, err := sdk.AccAddressFromBech32(liquidator)
		if err != nil {
			return fmt.Errorf("invalid white knight liquidator address: %s", liquidator)
		}

		if _, found := liquidators[liquidatorAddr.String()]; found {
			return fmt.Errorf("duplicate white knight liquidator: %s", liquidator)
		}
		liquidators[liquidatorAddr.String()] = struct{}{}
	}

	return nil
}

func ValidateFundingMultiple(i any) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("fundingMultiple must be positive: %d", v)
	}

	return nil
}

func ValidateDerivativeOrderSideCount(i any) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("DerivativeOrderSideCount must be positive: %d", v)
	}

	const maxDerivativeOrderSideCount = 1000
	if v > maxDerivativeOrderSideCount {
		return fmt.Errorf("DerivativeOrderSideCount must not exceed %d: %d", maxDerivativeOrderSideCount, v)
	}

	return nil
}

func ValidateInjRewardStakedRequirementThreshold(i any) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsZero() {
		return fmt.Errorf("InjRewardStakedRequirementThreshold cannot be zero: %d", v)
	}

	if v.IsNegative() {
		return fmt.Errorf("InjRewardStakedRequirementThreshold cannot be negative: %d", v)
	}

	return nil
}

func ValidateTradingRewardsVestingDuration(i any) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("trading rewards vesting duration must be non-negative: %d", v)
	}

	return nil
}

func ValidateLiquidatorRewardShareRate(i any) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("reward ratio cannot be nil: %s", v)
	}
	if v.IsNegative() {
		return fmt.Errorf("reward ratio cannot be negative: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("reward ratio cannot be greater than 1: %s", v)
	}

	return nil
}

func ValidateWhiteKnightLiquidatorRewardShareRate(i any) error {
	return ValidateLiquidatorRewardShareRate(i)
}

func ValidateAtomicMarketOrderAccessLevel(i any) error {
	v, ok := i.(AtomicMarketOrderAccessLevel)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !v.IsValid() {
		return fmt.Errorf("invalid AtomicMarketOrderAccessLevel value: %v", v)
	}
	return nil
}

func ValidateAtomicMarketOrderFeeMultiplier(i any) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("atomicMarketOrderFeeMultiplier cannot be nil: %s", v)
	}
	if v.LT(math.LegacyOneDec()) {
		return fmt.Errorf("atomicMarketOrderFeeMultiplier cannot be less than one: %s", v)
	}
	if v.GT(MaxFeeMultiplier) {
		return fmt.Errorf("atomicMarketOrderFeeMultiplier cannot be bigger than %v: %v", v, MaxFeeMultiplier)
	}
	return nil
}

func ValidateBool(i any) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

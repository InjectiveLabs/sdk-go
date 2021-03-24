package types

import (
	"fmt"

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

	// SpotMarketInstantListingFee is 1000 INJ
	SpotMarketInstantListingFee int64 = 1000

	// DerivativeMarketInstantListingFee is 1000 INJ
	DerivativeMarketInstantListingFee int64 = 1000
)

// Parameter keys
var (
	KeySpotMarketInstantListingFee       = []byte("SpotMarketInstantListingFee")
	KeyDerivativeMarketInstantListingFee = []byte("DerivativeMarketInstantListingFee")
	KeyDefaultSpotMakerFeeRate           = []byte("DefaultSpotMakerFeeRate")
	KeyDefaultSpotTakerFeeRate           = []byte("DefaultSpotTakerFeeRate")
	KeyDefaultDerivativeMakerFeeRate     = []byte("DefaultDerivativeMakerFeeRate")
	KeyDefaultDerivativeTakerFeeRate     = []byte("DefaultDerivativeTakerFeeRate")
	KeyDefaultInitialMarginRatio         = []byte("DefaultInitialMarginRatio")
	KeyDefaultMaintenanceMarginRatio     = []byte("DefaultMaintenanceMarginRatio")
	KeyDefaultFundingInterval            = []byte("DefaultFundingInterval")
	KeyFundingMultiple                   = []byte("FundingMultiple")
	KeyRelayerFeeShareRate               = []byte("RelayerFeeShareRate")
	KeyDefaultHourlyFundingRateCap       = []byte("DefaultHourlyFundingRateCap")
	KeyDefaultHourlyInterestRate         = []byte("DefaultHourlyInterestRate")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	SpotMarketInstantListingFee sdk.Coin,
	derivativeMarketInstantListingFee sdk.Coin,
	defaultSpotMakerFee sdk.Dec,
	defaultSpotTakerFee sdk.Dec,
	defaultDerivativeMakerFee sdk.Dec,
	defaultDerivativeTakerFee sdk.Dec,
	defaultInitialMarginRatio sdk.Dec,
	defaultMaintenanceMarginRatio sdk.Dec,
	defaultFundingInterval int64,
	fundingMultiple int64,
	relayerFeeShare sdk.Dec,
	defaultHourlyFundingRateCap sdk.Dec,
	defaultHourlyInterestRate sdk.Dec,
) Params {
	return Params{
		SpotMarketInstantListingFee:       SpotMarketInstantListingFee,
		DerivativeMarketInstantListingFee: derivativeMarketInstantListingFee,
		DefaultSpotMakerFeeRate:           defaultSpotMakerFee,
		DefaultSpotTakerFeeRate:           defaultSpotTakerFee,
		DefaultDerivativeMakerFeeRate:     defaultDerivativeMakerFee,
		DefaultDerivativeTakerFeeRate:     defaultDerivativeTakerFee,
		DefaultInitialMarginRatio:         defaultInitialMarginRatio,
		DefaultMaintenanceMarginRatio:     defaultMaintenanceMarginRatio,
		DefaultFundingInterval:            defaultFundingInterval,
		FundingMultiple:                   fundingMultiple,
		RelayerFeeShareRate:               relayerFeeShare,
		DefaultHourlyFundingRateCap:       defaultHourlyFundingRateCap,
		DefaultHourlyInterestRate:         defaultHourlyInterestRate,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	// TODO: @albert, add the rest of the parameters
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySpotMarketInstantListingFee, &p.SpotMarketInstantListingFee, validateSpotMarketInstantListingFee),
		paramtypes.NewParamSetPair(KeyDerivativeMarketInstantListingFee, &p.DerivativeMarketInstantListingFee, validateDerivativeMarketInstantListingFee),
		paramtypes.NewParamSetPair(KeyDefaultSpotMakerFeeRate, &p.DefaultSpotMakerFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultSpotTakerFeeRate, &p.DefaultSpotTakerFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultDerivativeMakerFeeRate, &p.DefaultDerivativeMakerFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultDerivativeTakerFeeRate, &p.DefaultDerivativeTakerFeeRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultInitialMarginRatio, &p.DefaultInitialMarginRatio, ValidateMarginRatio),
		paramtypes.NewParamSetPair(KeyDefaultMaintenanceMarginRatio, &p.DefaultMaintenanceMarginRatio, ValidateMarginRatio),
		paramtypes.NewParamSetPair(KeyDefaultFundingInterval, &p.DefaultFundingInterval, validateFundingInterval),
		paramtypes.NewParamSetPair(KeyFundingMultiple, &p.FundingMultiple, validateFundingMultiple),
		paramtypes.NewParamSetPair(KeyRelayerFeeShareRate, &p.RelayerFeeShareRate, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultHourlyFundingRateCap, &p.DefaultHourlyFundingRateCap, ValidateFee),
		paramtypes.NewParamSetPair(KeyDefaultHourlyInterestRate, &p.DefaultHourlyInterestRate, ValidateFee),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {

	return Params{
		SpotMarketInstantListingFee:       sdk.NewCoin("inj", sdk.NewIntWithDecimal(SpotMarketInstantListingFee, 18)),
		DerivativeMarketInstantListingFee: sdk.NewCoin("inj", sdk.NewIntWithDecimal(DerivativeMarketInstantListingFee, 18)),
		DefaultSpotMakerFeeRate:           sdk.NewDecWithPrec(1, 3), // default 0.1% maker fees
		DefaultSpotTakerFeeRate:           sdk.NewDecWithPrec(2, 3), // default 0.2% taker fees
		DefaultDerivativeMakerFeeRate:     sdk.NewDecWithPrec(1, 3), // default 0.1% maker fees
		DefaultDerivativeTakerFeeRate:     sdk.NewDecWithPrec(2, 3), // default 0.2% taker fees
		DefaultInitialMarginRatio:         sdk.NewDecWithPrec(5, 2), // default 5% initial margin ratio
		DefaultMaintenanceMarginRatio:     sdk.NewDecWithPrec(2, 2), // default 2% maintenance margin ratio
		DefaultFundingInterval:            DefaultFundingIntervalSeconds,
		FundingMultiple:                   DefaultFundingMultipleSeconds,
		RelayerFeeShareRate:               sdk.NewDecWithPrec(40, 2),      // default 40% relayer fee share
		DefaultHourlyFundingRateCap:       sdk.NewDecWithPrec(625, 6),     // default 0.0625% max hourly funding rate
		DefaultHourlyInterestRate:         sdk.NewDecWithPrec(416666, 11), // 0.01% daily interest rate = 0.0001 / 24 = 0.00000416666
	}
}

// Validate performs basic validation on exchange parameters.
func (p Params) Validate() error {
	if err := validateSpotMarketInstantListingFee(p.SpotMarketInstantListingFee); err != nil {
		return err
	}
	if err := validateDerivativeMarketInstantListingFee(p.DerivativeMarketInstantListingFee); err != nil {
		return err
	}
	if err := ValidateFee(p.DefaultSpotMakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(p.DefaultSpotTakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(p.DefaultDerivativeMakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(p.DefaultDerivativeTakerFeeRate); err != nil {
		return err
	}
	if err := ValidateMarginRatio(p.DefaultInitialMarginRatio); err != nil {
		return err
	}
	if err := ValidateMarginRatio(p.DefaultMaintenanceMarginRatio); err != nil {
		return err
	}
	if err := validateFundingInterval(p.DefaultFundingInterval); err != nil {
		return err
	}
	if err := validateFundingMultiple(p.FundingMultiple); err != nil {
		return err
	}
	if err := ValidateFee(p.RelayerFeeShareRate); err != nil {
		return err
	}
	if err := ValidateFee(p.DefaultHourlyFundingRateCap); err != nil {
		return err
	}
	if err := ValidateFee(p.DefaultHourlyInterestRate); err != nil {
		return err
	}

	return nil
}

func validateSpotMarketInstantListingFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() || !v.Amount.IsPositive() {
		return fmt.Errorf("invalid SpotMarketInstantListingFee: %T", i)
	}

	return nil
}

func validateDerivativeMarketInstantListingFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsValid() || !v.Amount.IsPositive() {
		return fmt.Errorf("invalid DerivativeMarketInstantListingFee: %T", i)
	}

	return nil
}

func ValidateFee(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("exchange fee cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("exchange fee cannot be greater than 1: %s", v)
	}

	return nil
}

func ValidateMarginRatio(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("margin ratio cannot be negative: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("margin ratio cannot be greater than 1: %s", v)
	}

	return nil
}

func validateFundingInterval(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("fundingInterval must be positive: %d", v)
	}

	return nil
}

func validateFundingMultiple(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("fundingMultiple must be positive: %d", v)
	}

	return nil
}

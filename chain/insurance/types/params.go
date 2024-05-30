package types

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// insurance params default values
const (
	// DefaultInsurancePeriodDurationSeconds represents the number of seconds in two weeks
	DefaultInsurancePeriod              = time.Hour * 24 * 14
	DefaultBinaryOptionsInsurancePeriod = time.Minute
)

// MaxUnderwritingAmount equals 1 trillion * 1e18
var MaxUnderwritingAmount, _ = math.NewIntFromString("1000000000000000000000000000")
var PerpetualExpiryFlag int64 = -1
var BinaryOptionsExpiryFlag int64 = -2

// Parameter keys
var (
	KeyDefaultRedemptionNoticePeriodDuration = []byte("defaultRedemptionNoticePeriodDuration")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	defaultRedemptionNoticePeriodDuration time.Duration,
) Params {
	return Params{
		DefaultRedemptionNoticePeriodDuration: defaultRedemptionNoticePeriodDuration,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDefaultRedemptionNoticePeriodDuration, &p.DefaultRedemptionNoticePeriodDuration, validateNoticePeriodDuration),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		DefaultRedemptionNoticePeriodDuration: DefaultInsurancePeriod,
	}
}

// Validate performs basic validation on insurance parameters.
func (p Params) Validate() error {
	if err := validateNoticePeriodDuration(p.DefaultRedemptionNoticePeriodDuration); err != nil {
		return err
	}

	return nil
}

func validateNoticePeriodDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("DefaultRedemptionNoticePeriodDuration must be positive: %d", v)
	}

	return nil
}

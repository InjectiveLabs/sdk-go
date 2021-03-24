package types

import (
	"fmt"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// Auction params default values
const (
	// DefaultOraclePeriodDurationSeconds represents the number of seconds in two weeks
	DefaultOraclePeriod int64 = 60 * 60 * 24 * 14
)

// Parameter keys
var (
	KeyOraclePeriod = []byte("OraclePeriod")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	oraclePeriod int64,
) Params {
	return Params{
		OraclePeriod: oraclePeriod,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	// TODO: @albert, add the rest of the parameters
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyOraclePeriod, &p.OraclePeriod, validateOraclePeriodDuration),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		OraclePeriod: DefaultOraclePeriod,
	}
}

// Validate performs basic validation on oracle parameters.
func (p Params) Validate() error {
	if err := validateOraclePeriodDuration(p.OraclePeriod); err != nil {
		return err
	}

	return nil
}

func validateOraclePeriodDuration(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("OraclePeriodDuration must be positive: %d", v)
	}

	return nil
}

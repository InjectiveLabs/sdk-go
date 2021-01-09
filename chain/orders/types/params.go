package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// Parameter keys
var (
	ParamStoreKeyFuturesEnabled = []byte("FuturesEnabled")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(futuresAddresses FuturesEnabledParams) Params {
	return Params{
		FuturesEnabled: futuresAddresses,
	}
}

// DefaultParams returns default evm parameters
func DefaultParams() Params {
	return Params{
		FuturesEnabled: FuturesEnabledParams{},
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyFuturesEnabled, &p.FuturesEnabled, validateFuturesEnabled),
	}
}

// Validate performs basic validation on orders parameters.
func (p Params) Validate() error {
	return validateFuturesEnabled(p.FuturesEnabled)
}

func validateFuturesEnabled(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// FuturesEnabledParams is a collection of parameters indicating if a futures address is enabled
type FuturesEnabledParams []string

package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// Parameter keys
var (
	ParamStoreKeyExchangeParams = []byte("ExchangeParams")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(ep ExchangeParams) Params {
	return Params{
		ExchangeParams: ep,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (m *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyExchangeParams, &m.ExchangeParams, validateExchangeParams),
	}
}

// Validate performs basic validation on orders parameters.
func (m Params) Validate() error {
	return validateExchangeParams(m.ExchangeParams)
}

func validateExchangeParams(i interface{}) error {
	_, ok := i.(ExchangeParams)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

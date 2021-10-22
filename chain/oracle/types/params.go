package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

var (
	LargestDecPrice sdk.Dec = sdk.MustNewDecFromStr("10000000")
)

const (
	// Each value below is the default value for each parameter when generating the default
	// genesis file.
	DefaultBandIBCEnabled         = false
	DefaultBandIbcRequestInterval = int64(7) //every 7 blocks
	DefaultBandIBCVersion         = "bandchain-1"
	DefaultBandIBCPortID          = "oracle"
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{}
}

// DefaultBandIBCParams returns a default set of band ibc parameters.
func DefaultBandIBCParams() BandIBCParams {
	return BandIBCParams{
		BandIbcEnabled:     DefaultBandIBCEnabled,
		IbcRequestInterval: DefaultBandIbcRequestInterval,
		IbcVersion:         DefaultBandIBCVersion,
		IbcPortId:          DefaultBandIBCPortID,
	}
}

// Validate performs basic validation on auction parameters.
func (p Params) Validate() error {
	return nil
}

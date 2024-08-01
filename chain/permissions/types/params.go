package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ParamTable
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(wasmHookQueryMaxGas uint64) Params {
	return Params{
		WasmHookQueryMaxGas: wasmHookQueryMaxGas,
	}
}

// default module parameters.
func DefaultParams() Params {
	return Params{
		WasmHookQueryMaxGas: 200_000,
	}
}

// validate params.
func (p Params) Validate() error {
	return nil
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

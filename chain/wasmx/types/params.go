package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

const (
	GasForFeeDeduction     uint64 = 13419
	GasForFeeRefund        uint64 = 13419
	DefaultGasContractCall uint64 = 63558
	MinExecutionGasLimit          = GasForFeeDeduction + GasForFeeRefund + DefaultGasContractCall
)

// Wasmx params default values
var (
	DefaultIsExecutionEnabled           = false
	DefaultMaxBeginBlockTotalGas uint64 = 42_000_000                        // 42M
	DefaultMaxContractGasLimit   uint64 = DefaultMaxBeginBlockTotalGas / 12 // 3.5M
	DefaultMinGasPrice           uint64 = 1_000_000_000                     // 1B
)

// Parameter keys
var (
	KeyIsExecutionEnabled    = []byte("IsExecutionEnabled")
	KeyRegistryContract      = []byte("RegistryContract")
	KeyMaxBeginBlockTotalGas = []byte("MaxBeginBlockTotalGas")
	KeyMaxContractGasLimit   = []byte("MaxContractGasLimit")
	KeyMinGasPrice           = []byte("MinGasPrice")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	isExecutionEnabled bool,
	registryContract string,
	maxBeginBlockTotalGas uint64,
	maxContractGasLimit uint64,
	minGasPrice uint64,
) Params {
	return Params{
		IsExecutionEnabled:    isExecutionEnabled,
		RegistryContract:      registryContract,
		MaxBeginBlockTotalGas: maxBeginBlockTotalGas,
		MaxContractGasLimit:   maxContractGasLimit,
		MinGasPrice:           minGasPrice,
	}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyRegistryContract, &p.RegistryContract, validateRegistryContract),
		paramtypes.NewParamSetPair(KeyMinGasPrice, &p.MinGasPrice, validateMinGasPrice),
		paramtypes.NewParamSetPair(KeyIsExecutionEnabled, &p.IsExecutionEnabled, validateIsExecutionEnabled),
		paramtypes.NewParamSetPair(KeyMaxBeginBlockTotalGas, &p.MaxBeginBlockTotalGas, validateMaxBeginBlockTotalGas),
		paramtypes.NewParamSetPair(KeyMaxContractGasLimit, &p.MaxContractGasLimit, validateMaxContractGasLimit),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		IsExecutionEnabled:    DefaultIsExecutionEnabled,
		RegistryContract:      "",
		MaxBeginBlockTotalGas: DefaultMaxBeginBlockTotalGas,
		MaxContractGasLimit:   DefaultMaxContractGasLimit,
		MinGasPrice:           DefaultMinGasPrice,
	}
}

// Validate performs basic validation on wasmx parameters.
func (p Params) Validate() error {
	if err := validateIsExecutionEnabled(p.IsExecutionEnabled); err != nil {
		return err
	}

	if err := validateRegistryContract(p.RegistryContract); err != nil {
		return err
	}

	if err := validateMaxBeginBlockTotalGas(p.MaxBeginBlockTotalGas); err != nil {
		return err
	}

	if err := validateMaxContractGasLimit(p.MaxContractGasLimit); err != nil {
		return err
	}

	if err := validateMinGasPrice(p.MinGasPrice); err != nil {
		return err
	}

	return nil
}

func validateMaxBeginBlockTotalGas(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("MaxBeginBlockTotalGas must be positive: %d", v)
	}

	return nil
}

func validateMaxContractGasLimit(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < MinExecutionGasLimit {
		return fmt.Errorf("MaxContractGasLimit %d must be greater than the MinExecutionGasLimit: %d", v, MinExecutionGasLimit)
	}
	return nil
}

func validateIsExecutionEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateMinGasPrice(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("MinGasPrice must be positive: %d", v)
	}
	return nil
}

func validateRegistryContract(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == "" {
		return nil
	}

	if _, err := sdk.AccAddressFromBech32(v); err != nil {
		return err
	}

	return nil
}

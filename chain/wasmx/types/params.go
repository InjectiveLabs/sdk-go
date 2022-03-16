package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

// Wasmx params default values
var (
	DefaultMinGasPriceBeginBlock = "0.002inj"
	DefaultIsExecutionEnabled    = false
	DefaultMaxGasBeginBlock      = int64(8987600)
	DefaultMinNumContracts       = int64(20)
)

// Parameter keys
var (
	KeyRegistryContractAddress = []byte("RegistryContractAddress")
	KeyMinGasPriceBeginBlock   = []byte("MinGasPriceBeginBlock")
	KeyIsExecutionEnabled      = []byte("IsExecutionEnabled")
	KeyMaxGasBeginBlock        = []byte("MaxGasBeginBlock")
	KeyMinNumOfContracts       = []byte("MinNumOfContracts")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	wasmxPeriod int64,
	minNextBidIncrementRate sdk.Dec,
) Params {
	return Params{}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyRegistryContractAddress, &p.RegistryContractAddress, validateRegistryContractAddress),
		paramtypes.NewParamSetPair(KeyMinGasPriceBeginBlock, &p.MinGasPriceBeginBlock, validateMinGasPriceBeginBlock),
		paramtypes.NewParamSetPair(KeyIsExecutionEnabled, &p.IsExecutionEnabled, validateIsExecutionEnabled),
		paramtypes.NewParamSetPair(KeyMaxGasBeginBlock, &p.MaxGasBeginBlock, validateMaxGasBeginBlock),
		paramtypes.NewParamSetPair(KeyMinNumOfContracts, &p.MinNumOfContracts, validateMinNumOfContracts),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		MinGasPriceBeginBlock: DefaultMinGasPriceBeginBlock,
		IsExecutionEnabled:    DefaultIsExecutionEnabled,
		MaxGasBeginBlock:      DefaultMaxGasBeginBlock,
		MinNumOfContracts:     DefaultMinNumContracts,
	}
}

// Validate performs basic validation on wasmx parameters.
func (p Params) Validate() error {
	if err := validateMaxGasBeginBlock(p.MaxGasBeginBlock); err != nil {
		return err
	}

	if err := validateMinNumOfContracts(p.MinNumOfContracts); err != nil {
		return err
	}

	if err := validateIsExecutionEnabled(p.IsExecutionEnabled); err != nil {
		return err
	}

	if err := validateMinGasPriceBeginBlock(p.MinGasPriceBeginBlock); err != nil {
		return err
	}

	if err := validateRegistryContractAddress(p.RegistryContractAddress); err != nil {
		return err
	}

	return nil
}

func validateMaxGasBeginBlock(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("MaxGasBeginBlock must be positive: %d", v)
	}

	return nil
}

func validateMinNumOfContracts(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("MinNumOfContracts must be positive: %d", v)
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

func validateMinGasPriceBeginBlock(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == "" {
		return nil
	}

	if _, err := sdk.ParseDecCoins(v); err != nil {
		return err
	}

	return nil
}

func validateRegistryContractAddress(i interface{}) error {
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

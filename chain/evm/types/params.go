package types

import (
	"bytes"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

const (
	// DefaultEVMDenom defines the default EVM denomination on Injective
	DefaultEVMDenom = "inj"
	// DefaultAllowUnprotectedTxs rejects all unprotected txs (true = allow, false = reject)
	DefaultAllowUnprotectedTxs = true
	// DefaultEnableCreate enables contract creation (i.e true)
	DefaultEnableCreate = true
	// DefaultEnableCall enables contract calls (i.e true)
	DefaultEnableCall = true
	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 160000000
	// DefaultRPCGasLimit is default gas limit for RPC call operations
	DefaultRPCGasLimit = 80000000
)

// NewParams creates a new Params instance
func NewParams(evmDenom string, allowUnprotectedTxs, enableCreate, enableCall bool, config ChainConfig, extraEIPs []int64) Params {
	return Params{
		EvmDenom:            evmDenom,
		AllowUnprotectedTxs: allowUnprotectedTxs,
		EnableCreate:        enableCreate,
		EnableCall:          enableCall,
		ExtraEIPs:           extraEIPs,
		ChainConfig:         config,
	}
}

// DefaultParams returns default evm parameters
// ExtraEIPs is empty to prevent overriding the latest hard fork instruction set
func DefaultParams() Params {
	config := DefaultChainConfig()
	return Params{
		EvmDenom:            DefaultEVMDenom,
		EnableCreate:        DefaultEnableCreate,
		EnableCall:          DefaultEnableCall,
		ChainConfig:         config,
		AllowUnprotectedTxs: DefaultAllowUnprotectedTxs,
	}
}

// Validate performs basic validation on evm parameters.
func (p Params) Validate() error {
	if err := ValidateEVMDenom(p.EvmDenom); err != nil {
		return err
	}

	if err := validateEIPs(p.ExtraEIPs); err != nil {
		return err
	}

	if err := ValidateBool(p.EnableCall); err != nil {
		return err
	}

	if err := ValidateBool(p.EnableCreate); err != nil {
		return err
	}

	if err := ValidateBool(p.AllowUnprotectedTxs); err != nil {
		return err
	}

	if err := ValidateBool(p.Permissioned); err != nil {
		return err
	}

	if err := validateAuthorizedDeployers(p.AuthorizedDeployers); err != nil {
		return err
	}

	return ValidateChainConfig(p.ChainConfig)
}

// EIPs returns the ExtraEIPS as a int slice
func (p Params) EIPs() []int {
	eips := make([]int, len(p.ExtraEIPs))
	for i, eip := range p.ExtraEIPs {
		eips[i] = int(eip)
	}
	return eips
}

func (p Params) WithPermissioned(permissioned bool) Params {
	p.Permissioned = permissioned
	return p
}

func (p Params) WithAuthorizedDeployers(authorizedDeployers []string) Params {
	p.AuthorizedDeployers = authorizedDeployers
	return p
}

func (p Params) IsAuthorisedDeployer(addr ethcommon.Address) bool {
	for _, addrStr := range p.AuthorizedDeployers {
		if bytes.Equal(ethcommon.HexToAddress(addrStr).Bytes(), addr.Bytes()) {
			return true
		}
	}

	return false
}

func ValidateEVMDenom(i interface{}) error {
	denom, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter EVM denom type: %T", i)
	}

	return sdk.ValidateDenom(denom)
}

func ValidateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateEIPs(i interface{}) error {
	eips, ok := i.([]int64)
	if !ok {
		return fmt.Errorf("invalid EIP slice type: %T", i)
	}

	for _, eip := range eips {
		if !vm.ValidEip(int(eip)) {
			return fmt.Errorf("EIP %d is not activateable, valid EIPS are: %s", eip, vm.ActivateableEips())
		}
	}
	return nil
}

func ValidateChainConfig(i interface{}) error {
	cfg, ok := i.(ChainConfig)
	if !ok {
		return fmt.Errorf("invalid chain config type: %T", i)
	}
	return cfg.Validate()
}

// IsLondon returns if london hardfork is enabled.
func IsLondon(ethConfig *params.ChainConfig, height int64) bool {
	return ethConfig.IsLondon(big.NewInt(height))
}

func validateAuthorizedDeployers(hexAddresses []string) error {
	authorisedAddresses := make(map[string]struct{})

	for _, addrStr := range hexAddresses {
		if !ethcommon.IsHexAddress(addrStr) {
			return fmt.Errorf("invalid address: %s", addrStr)
		}

		addr := ethcommon.HexToAddress(addrStr)

		if _, found := authorisedAddresses[string(addr.Bytes())]; found {
			return fmt.Errorf("duplicate authorised address: %s", addrStr)
		}

		authorisedAddresses[string(addr.Bytes())] = struct{}{}
	}

	return nil
}

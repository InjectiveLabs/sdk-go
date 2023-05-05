package types

import "cosmossdk.io/errors"

var (
	ErrInvalidGasLimit        = errors.Register(ModuleName, 1, "invalid gas limit")
	ErrInvalidGasPrice        = errors.Register(ModuleName, 2, "invalid gas price")
	ErrInvalidContractAddress = errors.Register(ModuleName, 3, "invalid contract address")
	ErrAlreadyRegistered      = errors.Register(ModuleName, 4, "contract already registered")
	ErrDuplicateContract      = errors.Register(ModuleName, 5, "duplicate contract")
	ErrNoContractAddresses    = errors.Register(ModuleName, 6, "no contract addresses found")
	ErrInvalidCodeId          = errors.Register(ModuleName, 7, "invalid code id")
	ErrDeductingGasFees       = errors.Register(ModuleName, 8, "not possible to deduct gas fees")
)

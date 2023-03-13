package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidGasLimit        = sdkerrors.Register(ModuleName, 1, "invalid gas limit")
	ErrInvalidGasPrice        = sdkerrors.Register(ModuleName, 2, "invalid gas price")
	ErrInvalidContractAddress = sdkerrors.Register(ModuleName, 3, "invalid contract address")
	ErrAlreadyRegistered      = sdkerrors.Register(ModuleName, 4, "contract already registered")
	ErrDuplicateContract      = sdkerrors.Register(ModuleName, 5, "duplicate contract")
	ErrNoContractAddresses    = sdkerrors.Register(ModuleName, 6, "no contract addresses found")
	ErrInvalidCodeId          = sdkerrors.Register(ModuleName, 7, "invalid code id")
	ErrDeductingGasFees       = sdkerrors.Register(ModuleName, 8, "not possible to deduct gas fees")
)

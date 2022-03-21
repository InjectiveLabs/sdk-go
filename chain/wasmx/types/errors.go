package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidGasLimit        = sdkerrors.Register(ModuleName, 1, "invalid gas limit")
	ErrInvalidGasPrice        = sdkerrors.Register(ModuleName, 2, "invalid gas price")
	ErrInvalidContractAddress = sdkerrors.Register(ModuleName, 3, "invalid contract address")
	ErrAlreadyRegistered      = sdkerrors.Register(ModuleName, 4, "contract already registered")
)

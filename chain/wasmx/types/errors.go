package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrBidInvalid             = sdkerrors.Register(ModuleName, 1, "invalid bid denom")
	ErrBidRound               = sdkerrors.Register(ModuleName, 2, "invalid bid round")
	ErrInvalidGasLimit        = sdkerrors.Register(ModuleName, 3, "invalid gas limit")
	ErrInvalidContractAddress = sdkerrors.Register(ModuleName, 4, "invalid contract address")
	ErrAlreadyRegistered      = sdkerrors.Register(ModuleName, 5, "contract already registered")
)

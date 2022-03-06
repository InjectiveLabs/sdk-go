package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrBidInvalid = sdkerrors.Register(ModuleName, 1, "invalid bid denom")
	ErrBidRound   = sdkerrors.Register(ModuleName, 2, "invalid bid round")
)

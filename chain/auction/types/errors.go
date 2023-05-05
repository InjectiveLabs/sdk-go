package types

import "cosmossdk.io/errors"

var (
	ErrBidInvalid = errors.Register(ModuleName, 1, "invalid bid denom")
	ErrBidRound   = errors.Register(ModuleName, 2, "invalid bid round")
)

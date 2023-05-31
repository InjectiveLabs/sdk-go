package types

import "cosmossdk.io/errors"

const (
	// RootCodespace is the codespace for all errors defined in this package
	RootCodespace = "injective"
)

// NOTE: We can't use 1 since that error code is reserved for internal errors.

var (
	// ErrInvalidChainID returns an error resulting from an invalid chain ID.
	ErrInvalidChainID = errors.Register(RootCodespace, 3, "invalid chain ID")
)

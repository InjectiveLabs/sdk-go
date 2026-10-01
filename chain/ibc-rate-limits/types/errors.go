//nolint:revive // types is a perfectly fine package name
package types

import "cosmossdk.io/errors"

var (
	ErrRateLimitExceeded      = errors.Register(ModuleName, 1, "exceeded configured rate limit")
	ErrInvalidRateLimit       = errors.Register(ModuleName, 2, "invalid rate limit object")
	ErrOracleUnavailable      = errors.Register(ModuleName, 3, "oracle price unavailable")
	ErrRateLimitAlreadyExists = errors.Register(ModuleName, 4, "rate limit already exists")
	ErrRateLimitNotFound      = errors.Register(ModuleName, 5, "rate limit not found")
)

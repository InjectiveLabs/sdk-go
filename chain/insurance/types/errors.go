package types

import (
	"cosmossdk.io/errors"
)

var (
	ErrInsuranceFundAlreadyExists = errors.Register(ModuleName, 1, "insurance fund already exists")
	ErrInsuranceFundNotFound      = errors.Register(ModuleName, 2, "insurance fund not found")
	ErrRedemptionAlreadyExists    = errors.Register(ModuleName, 3, "redemption already exists")
	ErrInvalidDepositAmount       = errors.Register(ModuleName, 4, "invalid deposit amount")
	ErrInvalidDepositDenom        = errors.Register(ModuleName, 5, "invalid deposit denom")
	ErrPayoutTooLarge             = errors.Register(ModuleName, 6, "insurance payout exceeds deposits")
	ErrInvalidTicker              = errors.Register(ModuleName, 7, "invalid ticker")
	ErrInvalidQuoteDenom          = errors.Register(ModuleName, 8, "invalid quote denom")
	ErrInvalidOracle              = errors.Register(ModuleName, 9, "invalid oracle")
	ErrInvalidExpirationTime      = errors.Register(ModuleName, 10, "invalid expiration time")
	ErrInvalidMarketID            = errors.Register(ModuleName, 11, "invalid marketID")
	ErrInvalidShareDenom          = errors.Register(ModuleName, 12, "invalid share denom")
)

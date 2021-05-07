package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInsuranceFundAlreadyExists = sdkerrors.Register(ModuleName, 1, "insurance fund already exists")
	ErrInsuranceFundNotFound      = sdkerrors.Register(ModuleName, 2, "insurance fund not found")
	ErrRedemptionAlreadyExists    = sdkerrors.Register(ModuleName, 3, "redemption already exists")
	ErrInvalidDepositAmount       = sdkerrors.Register(ModuleName, 4, "invalid deposit amount")
	ErrInvalidDepositDenom        = sdkerrors.Register(ModuleName, 5, "invalid deposit denom")
	ErrPayoutTooLarge             = sdkerrors.Register(ModuleName, 6, "insurance payout exceeds deposits")
	ErrInvalidTicker              = sdkerrors.Register(ModuleName, 7, "invalid ticker")
	ErrInvalidQuoteDenom          = sdkerrors.Register(ModuleName, 8, "invalid quote denom")
	ErrInvalidOracle              = sdkerrors.Register(ModuleName, 9, "invalid oracle")
	ErrInvalidExpirationTime      = sdkerrors.Register(ModuleName, 10, "invalid expiration time")
	ErrInvalidMarketID            = sdkerrors.Register(ModuleName, 11, "invalid marketID")
)

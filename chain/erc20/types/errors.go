package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrTokenPairExists        = errors.Register(ModuleName, 2, "attempting to create a token pair for bank denom that already has a pair associated")
	ErrUnauthorized           = errors.Register(ModuleName, 3, "unauthorized account")
	ErrInvalidGenesis         = errors.Register(ModuleName, 4, "invalid genesis")
	ErrInvalidTokenPair       = errors.Register(ModuleName, 5, "invalid token pair")
	ErrInvalidERC20Address    = errors.Register(ModuleName, 6, "invalid ERC20 contract address")
	ErrUnknownBankDenom       = errors.Register(ModuleName, 7, "unknown bank denom or zero supply")
	ErrUploadERC20Contract    = errors.Register(ModuleName, 8, "error uploading ERC20 contract")
	ErrInvalidTFDenom         = errors.Register(ModuleName, 9, "invalid token factory denom")
	ErrExistingEVMDenomSupply = errors.Register(ModuleName, 10, "respective evm/... denom has existing supply")
)

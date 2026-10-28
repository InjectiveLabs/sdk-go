package types

// DONTCOVER

import (
	"fmt"

	"cosmossdk.io/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrDenomExists              = errors.Register(ModuleName, 2, "attempting to create a denom that already exists (has bank metadata)")
	ErrUnauthorized             = errors.Register(ModuleName, 3, "unauthorized account")
	ErrInvalidDenom             = errors.Register(ModuleName, 4, "invalid denom")
	ErrInvalidCreator           = errors.Register(ModuleName, 5, "invalid creator")
	ErrInvalidAuthorityMetadata = errors.Register(ModuleName, 6, "invalid authority metadata")
	ErrInvalidGenesis           = errors.Register(ModuleName, 7, "invalid genesis")
	ErrSubdenomTooLong          = errors.Register(ModuleName, 8, fmt.Sprintf("subdenom too long, max length is %d bytes", MaxSubdenomLength))
	ErrSubdenomTooShort         = errors.Register(ModuleName, 9, fmt.Sprintf("subdenom too short, min length is %d bytes", MinSubdenomLength))
	ErrSubdenomNestedTooShort   = errors.Register(ModuleName, 10, fmt.Sprintf("nested subdenom too short, each one should have at least %d bytes", MinSubdenomLength))
	ErrCreatorTooLong           = errors.Register(ModuleName, 11, fmt.Sprintf("creator too long, max length is %d bytes", MaxCreatorLength))
	ErrDenomDoesNotExist        = errors.Register(ModuleName, 12, "denom does not exist")
	ErrAmountNotPositive        = errors.Register(ModuleName, 13, "amount has to be positive")
)

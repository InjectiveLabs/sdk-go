package types

// DONTCOVER

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/tokenfactory module sentinel errors
var (
	ErrDenomExists              = sdkerrors.Register(ModuleName, 2, "attempting to create a denom that already exists (has bank metadata)")
	ErrUnauthorized             = sdkerrors.Register(ModuleName, 3, "unauthorized account")
	ErrInvalidDenom             = sdkerrors.Register(ModuleName, 4, "invalid denom")
	ErrInvalidCreator           = sdkerrors.Register(ModuleName, 5, "invalid creator")
	ErrInvalidAuthorityMetadata = sdkerrors.Register(ModuleName, 6, "invalid authority metadata")
	ErrInvalidGenesis           = sdkerrors.Register(ModuleName, 7, "invalid genesis")
	ErrSubdenomTooLong          = sdkerrors.Register(ModuleName, 8, fmt.Sprintf("subdenom too long, max length is %d bytes", MaxSubdenomLength))
	ErrSubdenomTooShort         = sdkerrors.Register(ModuleName, 9, fmt.Sprintf("subdenom too short, min length is %d bytes", MinSubdenomLength))
	ErrSubdenomNestedTooShort   = sdkerrors.Register(ModuleName, 10, fmt.Sprintf("nested subdenom too short, each one should have at least %d bytes", MinSubdenomLength))
	ErrCreatorTooLong           = sdkerrors.Register(ModuleName, 11, fmt.Sprintf("creator too long, max length is %d bytes", MaxCreatorLength))
	ErrDenomDoesNotExist        = sdkerrors.Register(ModuleName, 12, "denom does not exist")
	ErrAmountNotPositive        = sdkerrors.Register(ModuleName, 13, "amount has to be positive")
)

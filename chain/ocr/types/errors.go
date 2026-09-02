package types

import (
	"cosmossdk.io/errors"
)

var (
	ErrStaleReport               = errors.Register(ModuleName, 1, "stale report")
	ErrIncompleteProposal        = errors.Register(ModuleName, 2, "incomplete proposal")
	ErrRepeatedAddress           = errors.Register(ModuleName, 3, "repeated oracle address")
	ErrTooManySigners            = errors.Register(ModuleName, 4, "too many signers")
	ErrIncorrectConfig           = errors.Register(ModuleName, 5, "incorrect config")
	ErrConfigDigestNotMatch      = errors.Register(ModuleName, 6, "config digest doesn't match")
	ErrWrongNumberOfSignatures   = errors.Register(ModuleName, 7, "wrong number of signatures")
	ErrIncorrectSignature        = errors.Register(ModuleName, 8, "incorrect signature")
	ErrNoTransmitter             = errors.Register(ModuleName, 9, "no transmitter specified")
	ErrIncorrectTransmissionData = errors.Register(ModuleName, 10, "incorrect transmission data")
	ErrNoTransmissionsFound      = errors.Register(ModuleName, 11, "no transmissions found")
	ErrMedianValueOutOfBounds    = errors.Register(ModuleName, 12, "median value is out of bounds")
	ErrIncorrectRewardPoolDenom  = errors.Register(ModuleName, 13, "LINK denom doesn't match")
	ErrNoRewardPool              = errors.Register(ModuleName, 14, "Reward Pool doesn't exist")
	ErrInvalidPayees             = errors.Register(ModuleName, 15, "wrong number of payees and transmitters")
	ErrModuleAdminRestricted     = errors.Register(ModuleName, 16, "action is restricted to the module admin")
	ErrFeedAlreadyExists         = errors.Register(ModuleName, 17, "feed already exists")
	ErrFeedDoesntExists          = errors.Register(ModuleName, 19, "feed doesnt exists")
	ErrAdminRestricted           = errors.Register(ModuleName, 20, "action is admin-restricted")
	ErrInsufficientRewardPool    = errors.Register(ModuleName, 21, "insufficient reward pool")
	ErrPayeeAlreadySet           = errors.Register(ModuleName, 22, "payee already set")
	ErrPayeeRestricted           = errors.Register(ModuleName, 23, "action is payee-restricted")
	ErrFeedConfigNotFound        = errors.Register(ModuleName, 24, "feed config not found")
)

package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptyRelayerAddr            = sdkerrors.Register(ModuleName, 1, "relayer address is empty")
	ErrBadRatesCount               = sdkerrors.Register(ModuleName, 2, "bad rates count")
	ErrBadResolveTimesCount        = sdkerrors.Register(ModuleName, 3, "bad resolve times")
	ErrBadRequestIDsCount          = sdkerrors.Register(ModuleName, 4, "bad request ID")
	ErrRelayerNotAuthorized        = sdkerrors.Register(ModuleName, 5, "relayer not authorized")
	ErrBadPriceFeedBaseCount       = sdkerrors.Register(ModuleName, 6, "bad price feed base count")
	ErrBadPriceFeedQuoteCount      = sdkerrors.Register(ModuleName, 7, "bad price feed quote count")
	ErrUnsupportedOracleType       = sdkerrors.Register(ModuleName, 8, "unsupported oracle type")
	ErrBadMessagesCount            = sdkerrors.Register(ModuleName, 9, "bad messages count")
	ErrBadCoinbaseMessage          = sdkerrors.Register(ModuleName, 10, "bad Coinbase message")
	ErrInvalidEthereumSignature    = sdkerrors.Register(ModuleName, 11, "bad Ethereum signature")
	ErrBadCoinbaseMessageTimestamp = sdkerrors.Register(ModuleName, 12, "bad Coinbase message timestamp")
	ErrCoinbasePriceNotFound       = sdkerrors.Register(ModuleName, 13, "Coinbase price not found")
)

package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// errors
var (
	ErrInvalidErc20Address   = sdkerrors.Register(ModuleName, 1, "invalid erc20 address")
	ErrUnmatchingCosmosDenom = sdkerrors.Register(ModuleName, 2, "unmatching cosmos denom")
	ErrNotAllowedBridge      = sdkerrors.Register(ModuleName, 3, "not allowed bridge")
	ErrInternalEthMinting    = sdkerrors.Register(ModuleName, 4, "internal ethereum minting error")
	ErrInitHubABI            = sdkerrors.Register(ModuleName, 5, "init hub abi error")
	ErrWritingEthTxPayload   = sdkerrors.Register(ModuleName, 6, "writing ethereum tx payload error")
)

package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	evmtypes "github.com/InjectiveLabs/sdk-go/chain/evm/types"
)

// BankKeeper defines the expected bank keeper methods
type BankKeeper interface {
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
}

// EVMKeeper defines the expected EVM keeper methods for Chainlink Data Streams verification
type EVMKeeper interface {
	EthCall(c context.Context, req *evmtypes.EthCallRequest) (*evmtypes.MsgEthereumTxResponse, error)
}

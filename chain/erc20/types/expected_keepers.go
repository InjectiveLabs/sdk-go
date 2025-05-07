package types

import (
	"context"

	"github.com/InjectiveLabs/sdk-go/chain/evm/statedb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	evmtypes "github.com/InjectiveLabs/sdk-go/chain/evm/types"
	tftypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type BankKeeper interface {
	HasSupply(ctx context.Context, denom string) bool
	GetDenomMetaData(ctx context.Context, denom string) (banktypes.Metadata, bool)
}

type AccountKeeper interface {
	GetSequence(ctx context.Context, addr sdk.AccAddress) (uint64, error)
}

type EVMKeeper interface {
	GetAccount(ctx sdk.Context, addr common.Address) *statedb.Account
	EthCall(c context.Context, req *evmtypes.EthCallRequest) (*evmtypes.MsgEthereumTxResponse, error)
	ApplyTransaction(ctx sdk.Context, msg *evmtypes.MsgEthereumTx) (*evmtypes.MsgEthereumTxResponse, error)
}

type TokenFactoryKeeper interface {
	GetAuthorityMetadata(ctx sdk.Context, denom string) (tftypes.DenomAuthorityMetadata, error)
}

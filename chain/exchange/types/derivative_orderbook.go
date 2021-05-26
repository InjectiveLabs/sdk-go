package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DerivativeOrderbook interface {
	GetNotional() sdk.Dec
	GetTotalQuantityFilled() sdk.Dec
	GetTransientOrderbookFills() *DerivativeOrderbookFills
	GetRestingOrderbookFills() *DerivativeOrderbookFills
	Peek(ctx sdk.Context) *PriceLevel
	Fill(fillQuantity sdk.Dec)
	Close()
}

type DerivativeOrderbookFills struct {
	Orders               []*DerivativeLimitOrder
	FillQuantities       []sdk.Dec
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type OrderbookStateChange struct {
	TransientBuyOrderbookFills  *OrderbookFills
	RestingBuyOrderbookFills    *OrderbookFills
	TransientSellOrderbookFills *OrderbookFills
	RestingSellOrderbookFills   *OrderbookFills
	ClearingPrice               sdk.Dec
}

type SpotOrderbook interface {
	GetNotional() sdk.Dec
	GetTotalQuantityFilled() sdk.Dec
	GetTransientOrderbookFills() *OrderbookFills
	GetRestingOrderbookFills() *OrderbookFills
	Peek() *PriceLevel
	Fill(sdk.Dec) error
	Close() error
}

type OrderbookFills struct {
	Orders         []*SpotLimitOrder
	FillQuantities []sdk.Dec
}

type PriceLevel struct {
	Price    sdk.Dec
	Quantity sdk.Dec
}

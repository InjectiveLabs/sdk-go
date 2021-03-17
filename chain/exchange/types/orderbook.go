package types

import (
	"github.com/cosmos/cosmos-sdk/types"
)

type OrderbookStateChange struct {
	NewBuyOrderbookFills      *OrderbookFills
	RestingBuyOrderbookFills  *OrderbookFills
	NewSellOrderbookFills     *OrderbookFills
	RestingSellOrderbookFills *OrderbookFills
	ClearingPrice             types.Dec
}

type Orderbook interface {
	GetNotional() types.Dec
	GetTotalQuantityFilled() types.Dec
	GetTransientOrderbookFills() *OrderbookFills
	GetRestingOrderbookFills() *OrderbookFills
	Done() bool
	Peek() *PriceLevel
	Fill(types.Dec) error
	Close() error
}

type OrderbookFills struct {
	Orders         []*SpotLimitOrder
	FillQuantities []types.Dec
}

type PriceLevel struct {
	Price    types.Dec
	Quantity types.Dec
}

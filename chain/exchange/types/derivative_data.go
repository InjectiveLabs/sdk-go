package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type DerivativeOrderbook interface {
	GetNotional() sdk.Dec
	GetTotalQuantityFilled() sdk.Dec
	GetTransientOrderbookFills() *DerivativeOrderbookFills
	GetRestingOrderbookFills() *DerivativeOrderbookFills
	Peek(ctx sdk.Context) *PriceLevel
	Fill(ctx sdk.Context, fillQuantity sdk.Dec)
	Close()
}

type DerivativeOrderbookFills struct {
	Orders         []*DerivativeLimitOrder
	FillQuantities []sdk.Dec
}

type DerivativeBatchExecutionData struct {
	Market *DerivativeMarket

	// update the deposits for margin deductions, payouts and refunds
	DepositDeltas        map[common.Hash]*DepositDelta
	DepositSubaccountIDs []common.Hash

	// updated positions
	Positions             []*Position
	PositionSubaccountIDs []common.Hash

	// resting limit order filled deltas to apply
	LimitOrderFilledDeltas []*DerivativeLimitOrderFilledDelta

	// events for batch market order and limit order execution
	MarketBuyOrderExecutionEvent          *EventBatchDerivativeExecution
	MarketSellOrderExecutionEvent         *EventBatchDerivativeExecution
	RestingLimitBuyOrderExecutionEvent    *EventBatchDerivativeExecution
	RestingLimitSellOrderExecutionEvent   *EventBatchDerivativeExecution
	TransientLimitBuyOrderExecutionEvent  *EventBatchDerivativeExecution
	TransientLimitSellOrderExecutionEvent *EventBatchDerivativeExecution
	PositionUpdateEvent                   *EventBatchDerivativePosition

	// event for new orders to add to the orderbook
	NewOrdersEvent          *EventNewDerivativeOrders
	CancelLimitOrderEvents  []*EventCancelDerivativeOrder
	CancelMarketOrderEvents []*EventCancelDerivativeOrder
}

type DerivativeLimitOrderFilledDelta struct {
	SubaccountIndexKey []byte
	FillableAmount     sdk.Dec
}

type PositionState struct {
	Position       *Position
	FundingPayment sdk.Dec
}

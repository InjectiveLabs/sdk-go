package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type DerivativeOrderStateExpansion struct {
	SubaccountID  common.Hash
	PositionDelta *PositionDelta
	Payout        sdk.Dec

	DepositChangeAmount sdk.Dec
	DepositRefundAmount sdk.Dec

	AuctionFeeReward   sdk.Dec
	FeeRecipientReward sdk.Dec
	FeeRecipient       common.Address
	Hash               common.Hash
	// For market orders, FillableAmount refers to the fillable quantity of the market order execution (if any)
	FillableAmount sdk.Dec
}

type DerivativeBatchExecutionData struct {
	Market *DerivativeMarket

	// update the deposits for margin deductions, payouts and refunds
	DepositMap           map[common.Hash]*DepositDelta
	DepositSubaccountIDs []common.Hash

	// resting limit order filled deltas to apply
	LimitOrderFilledDeltas []*DerivativeLimitOrderFilledDelta

	// events for batch market order and limit order execution
	MarketOrderExecutionEvent []*EventBatchDerivativeExecution
	LimitOrderExecutionEvent  []*EventBatchDerivativeExecution
	PositionUpdateEvent       *EventBatchDerivativePosition
	// event for new orders to add to the orderbook
	NewOrdersEvent *EventNewDerivativeOrders
}

type DerivativeLimitOrderFilledDelta struct {
	SubaccountIndexKey []byte
	FillableAmount     sdk.Dec
}

type PositionState struct {
	Position       *Position
	FundingPayment sdk.Dec
}

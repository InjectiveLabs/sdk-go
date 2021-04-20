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

func (o *DerivativeOrderbookFills) ProcessTransientDerivativeLimitOrderExpansions(
	isBuy bool,
	positionStates map[common.Hash]*PositionState,
	clearingPrice sdk.Dec,
	makerFeeRate, takerFeeRate, relayerFeeShareRate sdk.Dec,
) ([]*DerivativeOrderStateExpansion, []*DerivativeLimitOrder) {
	stateExpansions := make([]*DerivativeOrderStateExpansion, len(o.Orders))
	newRestingLimitOrders := make([]*DerivativeLimitOrder, 0, len(o.Orders))

	for idx := range o.Orders {
		order := o.Orders[idx]
		stateExpansions[idx] = GetDerivativeLimitOrderStateExpansion(
			isBuy,
			true,
			order,
			positionStates,
			o.FillQuantities[idx],
			clearingPrice,
			makerFeeRate,
			takerFeeRate,
			relayerFeeShareRate,
		)

		if !stateExpansions[idx].FillableAmount.IsZero() {
			newRestingLimitOrders = append(newRestingLimitOrders, order)
		}

	}
	return stateExpansions, newRestingLimitOrders
}

// ProcessRestingDerivativeLimitOrderExpansions processes the resting derivative limit order execution.
// NOTE: clearingPrice may be Nil
func (o *DerivativeOrderbookFills) ProcessRestingDerivativeLimitOrderExpansions(
	isBuy bool,
	positionStates map[common.Hash]*PositionState,
	clearingPrice sdk.Dec,
	tradeFeeRate, relayerFeeShareRate sdk.Dec,
) []*DerivativeOrderStateExpansion {
	stateExpansions := make([]*DerivativeOrderStateExpansion, len(o.Orders))
	// takerFeeRate is irrelevant for processing resting orders
	takerFeeRate := sdk.Dec{}

	for idx := range o.Orders {
		stateExpansions[idx] = GetDerivativeLimitOrderStateExpansion(
			isBuy,
			false,
			o.Orders[idx],
			positionStates,
			o.FillQuantities[idx],
			clearingPrice,
			tradeFeeRate,
			takerFeeRate,
			relayerFeeShareRate,
		)
	}
	return stateExpansions
}

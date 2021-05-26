package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type SpotOrderbookStateChange struct {
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

func NewSpotOrderbookStateChange(transientBuyOrders []*SpotLimitOrder, transientSellOrders []*SpotLimitOrder) *SpotOrderbookStateChange {
	orderbookStateChange := SpotOrderbookStateChange{
		TransientBuyOrderbookFills: &OrderbookFills{
			Orders: transientBuyOrders,
		},
		TransientSellOrderbookFills: &OrderbookFills{
			Orders: transientSellOrders,
		},
	}

	buyFillQuantities := make([]sdk.Dec, len(transientBuyOrders))
	for idx := range transientBuyOrders {
		buyFillQuantities[idx] = sdk.ZeroDec()
	}

	sellFillQuantities := make([]sdk.Dec, len(transientSellOrders))
	for idx := range transientSellOrders {
		sellFillQuantities[idx] = sdk.ZeroDec()
	}

	orderbookStateChange.TransientBuyOrderbookFills.FillQuantities = buyFillQuantities
	orderbookStateChange.TransientSellOrderbookFills.FillQuantities = sellFillQuantities
	return &orderbookStateChange
}

// ProcessBothRestingSpotLimitOrderExpansions processes both the orderbook state change to produce the spot execution batch events and filledDelta.
// Note: clearingPrice should be set to sdk.Dec{} for normal fills
func (o *SpotOrderbookStateChange) ProcessBothRestingSpotLimitOrderExpansions(
	marketID common.Hash,
	clearingPrice sdk.Dec,
	tradeFeeRate, relayerFeeShareRate sdk.Dec,
	baseDenomDepositDeltas DepositDeltas,
	quoteDenomDepositDeltas DepositDeltas,
) (limitBuyRestingOrderBatchEvent *EventBatchSpotExecution, limitSellRestingOrderBatchEvent *EventBatchSpotExecution, filledDeltas []*SpotLimitOrderDelta) {
	spotLimitBuyOrderStateExpansions := make([]*SpotOrderStateExpansion, 0)
	spotLimitSellOrderStateExpansions := make([]*SpotOrderStateExpansion, 0)
	filledDeltas = make([]*SpotLimitOrderDelta, 0)

	var currFilledDeltas []*SpotLimitOrderDelta

	if o.RestingBuyOrderbookFills != nil {
		spotLimitBuyOrderStateExpansions = o.ProcessRestingSpotLimitOrderExpansions(true, clearingPrice, tradeFeeRate, relayerFeeShareRate)
		// Process limit order events and filledDeltas
		limitBuyRestingOrderBatchEvent, currFilledDeltas = GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
			true,
			marketID,
			ExecutionType_LimitMatchRestingOrder,
			spotLimitBuyOrderStateExpansions,
			baseDenomDepositDeltas, quoteDenomDepositDeltas,
		)
		filledDeltas = append(filledDeltas, currFilledDeltas...)
	}

	if o.RestingSellOrderbookFills != nil {
		spotLimitSellOrderStateExpansions = o.ProcessRestingSpotLimitOrderExpansions(false, clearingPrice, tradeFeeRate, relayerFeeShareRate)
		// Process limit order events and filledDeltas
		limitSellRestingOrderBatchEvent, currFilledDeltas = GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
			false,
			marketID,
			ExecutionType_LimitMatchRestingOrder,
			spotLimitSellOrderStateExpansions,
			baseDenomDepositDeltas, quoteDenomDepositDeltas,
		)
		filledDeltas = append(filledDeltas, currFilledDeltas...)
	}
	return
}

// ProcessBothTransientSpotLimitOrderExpansions processes the transient spot limit orderbook state change.
// Note: clearingPrice should be set to sdk.Dec{} for normal fills
func (o *SpotOrderbookStateChange) ProcessBothTransientSpotLimitOrderExpansions(
	marketID common.Hash,
	clearingPrice sdk.Dec,
	makerFeeRate, takerFeeRate, relayerFeeShareRate sdk.Dec,
	baseDenomDepositDeltas DepositDeltas,
	quoteDenomDepositDeltas DepositDeltas,
) (
	limitBuyNewOrderBatchEvent *EventBatchSpotExecution,
	limitSellNewOrderBatchEvent *EventBatchSpotExecution,
	newRestingBuySpotLimitOrders []*SpotLimitOrder,
	newRestingSellSpotLimitOrders []*SpotLimitOrder,
) {
	var expansions []*SpotOrderStateExpansion
	if o.TransientBuyOrderbookFills != nil {
		expansions, newRestingBuySpotLimitOrders = o.processNewSpotLimitBuyExpansions(clearingPrice, makerFeeRate, takerFeeRate, relayerFeeShareRate)
		limitBuyNewOrderBatchEvent, _ = GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
			true,
			marketID,
			ExecutionType_LimitMatchNewOrder,
			expansions,
			baseDenomDepositDeltas, quoteDenomDepositDeltas,
		)
	}

	if o.TransientSellOrderbookFills != nil {
		expansions, newRestingSellSpotLimitOrders = o.processNewSpotLimitSellExpansions(clearingPrice, takerFeeRate, relayerFeeShareRate)
		limitSellNewOrderBatchEvent, _ = GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
			false,
			marketID,
			ExecutionType_LimitMatchNewOrder,
			expansions,
			baseDenomDepositDeltas, quoteDenomDepositDeltas,
		)
	}
	return
}

// ProcessRestingSpotLimitOrderExpansions processes the resting spot limit orderbook state change.
// Note: clearingPrice should be set to sdk.Dec{} for normal fills
func (o *SpotOrderbookStateChange) ProcessRestingSpotLimitOrderExpansions(
	isLimitBuy bool,
	clearingPrice sdk.Dec,
	makerFeeRate, relayerFeeShare sdk.Dec,
) []*SpotOrderStateExpansion {
	var fills *OrderbookFills
	if isLimitBuy {
		fills = o.RestingBuyOrderbookFills
	} else {
		fills = o.RestingSellOrderbookFills
	}
	return ProcessRestingSpotLimitOrderExpansions(fills, isLimitBuy, clearingPrice, makerFeeRate, relayerFeeShare)
}

// TODO: refactor to merge processNewSpotLimitBuyExpansions and processNewSpotLimitSellExpansions
func (o *SpotOrderbookStateChange) processNewSpotLimitBuyExpansions(
	clearingPrice sdk.Dec,
	makerFeeRate, takerFeeRate, relayerFeeShare sdk.Dec,
) ([]*SpotOrderStateExpansion, []*SpotLimitOrder) {
	orderbookFills := o.TransientBuyOrderbookFills
	stateExpansions := make([]*SpotOrderStateExpansion, len(orderbookFills.Orders))
	newRestingOrders := make([]*SpotLimitOrder, 0, len(orderbookFills.Orders))

	for idx, order := range orderbookFills.Orders {
		fillQuantity := sdk.ZeroDec()
		if orderbookFills.FillQuantities != nil {
			fillQuantity = orderbookFills.FillQuantities[idx]
		}
		stateExpansions[idx] = getNewSpotLimitBuyStateExpansion(
			order,
			common.BytesToHash(order.OrderHash),
			clearingPrice, fillQuantity,
			makerFeeRate, takerFeeRate, relayerFeeShare,
		)

		if fillQuantity.LT(order.OrderInfo.Quantity) {
			order.Fillable = order.Fillable.Sub(fillQuantity)
			newRestingOrders = append(newRestingOrders, order)
		}
	}
	return stateExpansions, newRestingOrders
}

// processNewSpotLimitSellExpansions processes.
// Note: clearingPrice should be set to sdk.Dec{} for normal fills
func (o *SpotOrderbookStateChange) processNewSpotLimitSellExpansions(
	clearingPrice sdk.Dec,
	takerFeeRate, relayerFeeShare sdk.Dec,
) ([]*SpotOrderStateExpansion, []*SpotLimitOrder) {
	orderbookFills := o.TransientSellOrderbookFills

	stateExpansions := make([]*SpotOrderStateExpansion, len(orderbookFills.Orders))
	newRestingOrders := make([]*SpotLimitOrder, 0, len(orderbookFills.Orders))

	for idx, order := range orderbookFills.Orders {
		fillQuantity, fillPrice := orderbookFills.FillQuantities[idx], order.OrderInfo.Price
		if !clearingPrice.IsNil() {
			fillPrice = clearingPrice
		}
		stateExpansions[idx] = getSpotLimitSellStateExpansion(
			order,
			fillQuantity,
			fillPrice,
			takerFeeRate,
			relayerFeeShare,
		)
		if order.Fillable.IsPositive() {
			newRestingOrders = append(newRestingOrders, order)
		}
	}
	return stateExpansions, newRestingOrders
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type DerivativeBatchExecutionData struct {
	Market *DerivativeMarket

	MarkPrice sdk.Dec
	Funding   *PerpetualMarketFunding

	// update the deposits for margin deductions, payouts and refunds
	DepositDeltas        DepositDeltas
	DepositSubaccountIDs []common.Hash

	// updated positions
	Positions             []*Position
	PositionSubaccountIDs []common.Hash

	// resting limit order filled deltas to apply
	LimitOrderFilledDeltas []*LimitOrderFilledDelta

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

	VwapData *VwapData
}

type DerivativeMatchingExpansionData struct {
	TransientLimitBuyExpansions  []*DerivativeOrderStateExpansion
	TransientLimitSellExpansions []*DerivativeOrderStateExpansion
	RestingLimitBuyExpansions    []*DerivativeOrderStateExpansion
	RestingLimitSellExpansions   []*DerivativeOrderStateExpansion
	RestingLimitBuyOrderCancels  []*DerivativeLimitOrder
	RestingLimitSellOrderCancels []*DerivativeLimitOrder
	ClearingPrice                sdk.Dec
	ClearingQuantity             sdk.Dec
	NewRestingLimitBuyOrders     []*DerivativeLimitOrder // transient buy orders that become new resting limit orders
	NewRestingLimitSellOrders    []*DerivativeLimitOrder // transient sell orders that become new resting limit orders
}

type DerivativeMarketOrderExpansionData struct {
	MarketBuyExpansions        []*DerivativeOrderStateExpansion
	MarketSellExpansions       []*DerivativeOrderStateExpansion
	LimitBuyExpansions         []*DerivativeOrderStateExpansion
	LimitSellExpansions        []*DerivativeOrderStateExpansion
	LimitBuyOrderCancels       []*DerivativeLimitOrder
	LimitSellOrderCancels      []*DerivativeLimitOrder
	MarketBuyOrderCancels      []*DerivativeMarketOrderCancel
	MarketSellOrderCancels     []*DerivativeMarketOrderCancel
	MarketBuyClearingPrice     sdk.Dec
	MarketSellClearingPrice    sdk.Dec
	MarketBuyClearingQuantity  sdk.Dec
	MarketSellClearingQuantity sdk.Dec
}

func (d *DerivativeMarketOrderExpansionData) SetExecutionData(
	isMarketBuy bool,
	marketOrderClearingPrice, marketOrderClearingQuantity sdk.Dec,
	limitOrderCancels []*DerivativeLimitOrder,
	marketOrderStateExpansions,
	restingLimitOrderStateExpansions []*DerivativeOrderStateExpansion,
	marketOrderCancels []*DerivativeMarketOrderCancel,
) {
	if isMarketBuy {
		d.MarketBuyClearingPrice = marketOrderClearingPrice
		d.MarketBuyClearingQuantity = marketOrderClearingQuantity
		d.LimitSellOrderCancels = limitOrderCancels
		d.MarketBuyExpansions = marketOrderStateExpansions
		d.LimitSellExpansions = restingLimitOrderStateExpansions
		d.MarketBuyOrderCancels = marketOrderCancels
	} else {
		d.MarketSellClearingPrice = marketOrderClearingPrice
		d.MarketSellClearingQuantity = marketOrderClearingQuantity
		d.LimitBuyOrderCancels = limitOrderCancels
		d.MarketSellExpansions = marketOrderStateExpansions
		d.LimitBuyExpansions = restingLimitOrderStateExpansions
		d.MarketSellOrderCancels = marketOrderCancels
	}
}

func (e *DerivativeMatchingExpansionData) GetDerivativeLimitOrderBatchExecution(
	market *DerivativeMarket,
	markPrice sdk.Dec,
	funding *PerpetualMarketFunding,
	positionStates PositionStates,
) *DerivativeBatchExecutionData {
	depositDeltas := NewDepositDeltas()

	// process undermargined resting limit order forced cancellations
	cancelLimitOrdersEvents := e.ApplyCancellationsAndGetDerivativeLimitCancelEvents(market.MarketID(), market.MakerFeeRate, depositDeltas, positionStates)

	positionUpdateEvent, positions, positionSubaccountIDs := positionStates.GetPositionUpdateEvent(market.MarketID(), funding)

	transientLimitBuyOrderBatchEvent, transientLimitBuyFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(true, ExecutionType_LimitMatchNewOrder, market, funding, e.TransientLimitBuyExpansions, depositDeltas)
	restingLimitBuyOrderBatchEvent, restingLimitBuyFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(true, ExecutionType_LimitMatchRestingOrder, market, funding, e.RestingLimitBuyExpansions, depositDeltas)
	transientLimitSellOrderBatchEvent, transientLimitSellFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(false, ExecutionType_LimitMatchNewOrder, market, funding, e.TransientLimitSellExpansions, depositDeltas)
	restingLimitSellOrderBatchEvent, restingLimitSellFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(false, ExecutionType_LimitMatchRestingOrder, market, funding, e.RestingLimitSellExpansions, depositDeltas)

	filledDeltas := mergeDerivativeLimitOrderFilledDeltas(transientLimitBuyFilledDeltas, restingLimitBuyFilledDeltas, transientLimitSellFilledDeltas, restingLimitSellFilledDeltas)

	// sort keys since map iteration is non-deterministic
	depositDeltaKeys := depositDeltas.GetSortedSubaccountKeys()

	var vwapData *VwapData
	if market.IsPerpetual {
		vwapData = vwapData.ApplyExecution(e.ClearingPrice, e.ClearingQuantity)
	}

	// Final Step: Store the DerivativeBatchExecutionData for future reduction/processing
	batch := &DerivativeBatchExecutionData{
		Market:                                market,
		MarkPrice:                             markPrice,
		Funding:                               funding,
		DepositDeltas:                         depositDeltas,
		DepositSubaccountIDs:                  depositDeltaKeys,
		Positions:                             positions,
		PositionSubaccountIDs:                 positionSubaccountIDs,
		LimitOrderFilledDeltas:                filledDeltas,
		MarketBuyOrderExecutionEvent:          nil,
		MarketSellOrderExecutionEvent:         nil,
		RestingLimitBuyOrderExecutionEvent:    restingLimitBuyOrderBatchEvent,
		RestingLimitSellOrderExecutionEvent:   restingLimitSellOrderBatchEvent,
		TransientLimitBuyOrderExecutionEvent:  transientLimitBuyOrderBatchEvent,
		TransientLimitSellOrderExecutionEvent: transientLimitSellOrderBatchEvent,
		PositionUpdateEvent:                   positionUpdateEvent,
		NewOrdersEvent:                        nil,
		CancelLimitOrderEvents:                cancelLimitOrdersEvents,
		CancelMarketOrderEvents:               nil,
		VwapData:                              vwapData,
	}

	if len(e.NewRestingLimitBuyOrders) > 0 || len(e.NewRestingLimitSellOrders) > 0 {
		batch.NewOrdersEvent = &EventNewDerivativeOrders{
			MarketId:   market.MarketId,
			BuyOrders:  e.NewRestingLimitBuyOrders,
			SellOrders: e.NewRestingLimitSellOrders,
		}
	}

	return batch
}

func (e *DerivativeMarketOrderExpansionData) GetDerivativeMarketCancelEvents(
	marketID common.Hash,
	depositDeltas DepositDeltas,
	positionStates PositionStates,
) []*EventCancelDerivativeOrder {

	marketIDHex := marketID.Hex()
	cancelOrdersEvent := make([]*EventCancelDerivativeOrder, 0, len(e.MarketBuyOrderCancels)+len(e.MarketSellOrderCancels))

	for idx := range e.MarketBuyOrderCancels {
		orderCancel := e.MarketBuyOrderCancels[idx]
		orderCancel.ApplyDerivativeMarketCancellation(depositDeltas, positionStates)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:          marketIDHex,
			IsLimitCancel:     false,
			MarketOrderCancel: orderCancel,
		})
	}

	for idx := range e.MarketSellOrderCancels {
		orderCancel := e.MarketSellOrderCancels[idx]
		orderCancel.ApplyDerivativeMarketCancellation(depositDeltas, positionStates)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:          marketIDHex,
			IsLimitCancel:     false,
			MarketOrderCancel: orderCancel,
		})
	}
	return cancelOrdersEvent
}

func applyDerivativeLimitCancellation(
	order *DerivativeLimitOrder,
	makerFeeRate sdk.Dec,
	depositDeltas DepositDeltas,
	positionStates PositionStates,
) {
	subaccountID := order.SubaccountID()
	// For vanilla orders, increment the available balance
	// For reduce-only orders, free the position hold quantity
	if order.IsVanilla() {
		depositDelta := order.GetCancelDepositDelta(makerFeeRate)
		depositDeltas.ApplyDepositDelta(subaccountID, depositDelta)
	} else if order.IsReduceOnly() {
		position := positionStates[subaccountID].Position
		position.HoldQuantity = position.HoldQuantity.Sub(order.Fillable)
	}
}

func (e *DerivativeMatchingExpansionData) ApplyCancellationsAndGetDerivativeLimitCancelEvents(
	marketID common.Hash,
	makerFeeRate sdk.Dec,
	depositDeltas DepositDeltas,
	positionStates PositionStates,
) []*EventCancelDerivativeOrder {
	marketIDHex := marketID.Hex()

	cancelOrdersEvent := make([]*EventCancelDerivativeOrder, 0, len(e.RestingLimitBuyOrderCancels)+len(e.RestingLimitSellOrderCancels))

	for idx := range e.RestingLimitBuyOrderCancels {
		order := e.RestingLimitBuyOrderCancels[idx]
		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas, positionStates)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
	}

	for idx := range e.RestingLimitSellOrderCancels {
		order := e.RestingLimitSellOrderCancels[idx]
		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas, positionStates)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
	}

	return cancelOrdersEvent
}

func (e *DerivativeMarketOrderExpansionData) ApplyCancellationsAndGetDerivativeLimitCancelEvents(
	marketID common.Hash,
	makerFeeRate sdk.Dec,
	depositDeltas DepositDeltas,
	positionStates PositionStates,
) []*EventCancelDerivativeOrder {
	marketIDHex := marketID.Hex()
	cancelOrdersEvent := make([]*EventCancelDerivativeOrder, 0, len(e.LimitBuyOrderCancels)+len(e.LimitSellOrderCancels))

	for idx := range e.LimitBuyOrderCancels {
		order := e.LimitBuyOrderCancels[idx]
		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas, positionStates)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
	}

	for idx := range e.LimitSellOrderCancels {
		order := e.LimitSellOrderCancels[idx]
		applyDerivativeLimitCancellation(order, makerFeeRate, depositDeltas, positionStates)
		cancelOrdersEvent = append(cancelOrdersEvent, &EventCancelDerivativeOrder{
			MarketId:      marketIDHex,
			IsLimitCancel: true,
			LimitOrder:    order,
		})
	}

	return cancelOrdersEvent
}

func (e *DerivativeMarketOrderExpansionData) GetDerivativeMarketOrderBatchExecution(
	market *DerivativeMarket,
	markPrice sdk.Dec,
	funding *PerpetualMarketFunding,
	positionStates PositionStates,
) *DerivativeBatchExecutionData {
	depositDeltas := NewDepositDeltas()

	// process undermargined limit order forced cancellations
	cancelLimitOrdersEvents := e.ApplyCancellationsAndGetDerivativeLimitCancelEvents(market.MarketID(), market.MakerFeeRate, depositDeltas, positionStates)

	// process unfilled market order cancellations
	cancelMarketOrdersEvents := e.GetDerivativeMarketCancelEvents(market.MarketID(), depositDeltas, positionStates)

	positionUpdateEvent, positions, positionSubaccountIDs := positionStates.GetPositionUpdateEvent(market.MarketID(), funding)

	buyMarketOrderBatchEvent, _ := ApplyDeltasAndGetDerivativeOrderBatchEvent(true, ExecutionType_Market, market, funding, e.MarketBuyExpansions, depositDeltas)
	sellMarketOrderBatchEvent, _ := ApplyDeltasAndGetDerivativeOrderBatchEvent(false, ExecutionType_Market, market, funding, e.MarketSellExpansions, depositDeltas)
	restingLimitBuyOrderBatchEvent, limitBuyFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(true, ExecutionType_LimitFill, market, funding, e.LimitBuyExpansions, depositDeltas)
	restingLimitSellOrderBatchEvent, limitSellFilledDeltas := ApplyDeltasAndGetDerivativeOrderBatchEvent(false, ExecutionType_LimitFill, market, funding, e.LimitSellExpansions, depositDeltas)

	filledDeltas := mergeDerivativeLimitOrderFilledDeltas(limitBuyFilledDeltas, limitSellFilledDeltas, nil, nil)

	// sort keys since map iteration is non-deterministic
	depositDeltaKeys := depositDeltas.GetSortedSubaccountKeys()

	var vwapData *VwapData
	if market.IsPerpetual {
		vwapData = vwapData.ApplyExecution(e.MarketBuyClearingPrice, e.MarketBuyClearingQuantity)
		vwapData = vwapData.ApplyExecution(e.MarketSellClearingPrice, e.MarketSellClearingQuantity)
	}

	// Final Step: Store the DerivativeBatchExecutionData for future reduction/processing
	batch := &DerivativeBatchExecutionData{
		Market:                                market,
		MarkPrice:                             markPrice,
		Funding:                               funding,
		DepositDeltas:                         depositDeltas,
		DepositSubaccountIDs:                  depositDeltaKeys,
		Positions:                             positions,
		PositionSubaccountIDs:                 positionSubaccountIDs,
		LimitOrderFilledDeltas:                filledDeltas,
		MarketBuyOrderExecutionEvent:          buyMarketOrderBatchEvent,
		MarketSellOrderExecutionEvent:         sellMarketOrderBatchEvent,
		RestingLimitBuyOrderExecutionEvent:    restingLimitBuyOrderBatchEvent,
		RestingLimitSellOrderExecutionEvent:   restingLimitSellOrderBatchEvent,
		TransientLimitBuyOrderExecutionEvent:  nil,
		TransientLimitSellOrderExecutionEvent: nil,
		PositionUpdateEvent:                   positionUpdateEvent,
		NewOrdersEvent:                        nil,
		CancelLimitOrderEvents:                cancelLimitOrdersEvents,
		CancelMarketOrderEvents:               cancelMarketOrdersEvents,
		VwapData:                              vwapData,
	}
	return batch
}

func mergeDerivativeLimitOrderFilledDeltas(d1, d2, d3, d4 []*LimitOrderFilledDelta) []*LimitOrderFilledDelta {
	filledDeltas := make([]*LimitOrderFilledDelta, 0, len(d1)+len(d2)+len(d3)+len(d4))
	if len(d1) > 0 {
		filledDeltas = append(filledDeltas, d1...)
	}
	if len(d2) > 0 {
		filledDeltas = append(filledDeltas, d2...)
	}
	if len(d3) > 0 {
		filledDeltas = append(filledDeltas, d3...)
	}
	if len(d4) > 0 {
		filledDeltas = append(filledDeltas, d4...)
	}
	return filledDeltas
}

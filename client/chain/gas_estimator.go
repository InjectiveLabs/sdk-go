package chain

import (
	"math"

	exchangetypes "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// refer to python sdk: https://github.com/InjectiveLabs/sdk-python/blob/master/pyinjective/core/gas_limit_estimator.py#L15
const (
	GENERAL_MESSAGE_GAS_LIMIT = 25_000
	BASIC_REFERENCE_GAS_LIMIT = 150_000

	SPOT_ORDER_CREATION_GAS_LIMIT          = 52_000
	DERIVATIVE_ORDER_CREATION_GAS_LIMIT    = 84_000
	SPOT_ORDER_CANCELATION_GAS_LIMIT       = 50_000
	DERIVATIVE_ORDER_CANCELATION_GAS_LIMIT = 68_000
	// POST ONLY orders take around 50% more gas to create than normal orders due to the required validations
	SPOT_POST_ONLY_ORDER_MULTIPLIER       = 0.62
	DERIVATIVE_POST_ONLY_ORDER_MULTIPLIER = 0.35

	// for batch update
	CANCEL_ALL_SPOT_MARKET_GAS_LIMIT       = 40_000
	CANCEL_ALL_DERIVATIVE_MARKET_GAS_LIMIT = 50_000
	MESSAGE_GAS_LIMIT                      = 30_000

	AVERAGE_CANCEL_ALL_AFFECTED_ORDERS = 20

	TRANSACTION_GAS_LIMIT = 60_000
)

type msgGasEstimator interface {
	appliesTo(msg sdk.Msg) bool
	estimateMsgGas(msgs sdk.Msg) uint64
}

type TXGasEstimator struct {
	msg_gas_estimators []msgGasEstimator
}

func newTXGasEstimator() *TXGasEstimator {
	return &TXGasEstimator{
		msg_gas_estimators: []msgGasEstimator{
			newMsgCreateSpotLimitOrderGasEstimator(),
			newMsgBatchCreateSpotLimitOrdersGasEstimator(),
			newMsgBatchCancelSpotOrdersGasEstimator(),
			newMsgCancelSpotOrderGasEstimator(),
		},
	}
}

func (g *TXGasEstimator) EstimateTXGas(msgs ...sdk.Msg) uint64 {

	// add up all types of messages' gas
	// messages_gas_limit = Decimal("0")
	// for message in messages:
	//    estimator = GasLimitEstimator.for_message(message=message)
	//    messages_gas_limit += estimator.gas_limit()
	// messages_gas_limit = math.ceil(messages_gas_limit)

	// add the transaction gas limit
	// gas_wanted = messages_gas_limit + self.TRANSACTION_GAS_LIMIT
	// transaction.with_gas(gas=gas_wanted)

	// msg types:
	// - spot order creation, with spot post only order multiplier
	// - derivative order creation, with derivative post only order multiplier
	// - spot order cancellation
	// - derivative order cancellation

	var totalMsgGas uint64
	for i := range msgs {
		msg := msgs[i]
		totalMsgGas += g.estimateMsgGas(msg)
	}
	txGas := totalMsgGas + TRANSACTION_GAS_LIMIT

	// for i := range msgs {
	// 	msg := msgs[i]
	// 	fmt.Printf("TypeURL: %s\n", proto.MessageName(msg))
	// }
	return txGas
}

func (g *TXGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// use chain responsibility pattern
	for i := range g.msg_gas_estimators {
		estimator := g.msg_gas_estimators[i]
		if estimator.appliesTo(msg) {
			return estimator.estimateMsgGas(msg)
		}
	}
	return GENERAL_MESSAGE_GAS_LIMIT
}

func selectPostOnlyOrders(orders []exchangetypes.SpotOrder) []exchangetypes.SpotOrder {
	var filtered []exchangetypes.SpotOrder
	for _, order := range orders {
		if order.OrderType == exchangetypes.OrderType_BUY_PO || order.OrderType == exchangetypes.OrderType_SELL_PO {
			filtered = append(filtered, order)
		}
	}
	return filtered
}

// MsgCreateSpotLimitOrder
type MsgCreateSpotLimitOrderGasEstimator struct {
}

func newMsgCreateSpotLimitOrderGasEstimator() *MsgCreateSpotLimitOrderGasEstimator {
	return &MsgCreateSpotLimitOrderGasEstimator{}
}

func (e *MsgCreateSpotLimitOrderGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgCreateSpotLimitOrder)
	return ok
}

func (e *MsgCreateSpotLimitOrderGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	postOnlyOrder := selectPostOnlyOrders([]exchangetypes.SpotOrder{msg.(*exchangetypes.MsgCreateSpotLimitOrder).Order})
	total := uint64(SPOT_ORDER_CREATION_GAS_LIMIT)
	if len(postOnlyOrder) > 0 {
		total += uint64(math.Ceil(SPOT_ORDER_CREATION_GAS_LIMIT * SPOT_POST_ONLY_ORDER_MULTIPLIER))
	}
	return total
}

// MsgBatchCreateSpotLimitOrders

type MsgBatchCreateSpotLimitOrdersGasEstimator struct {
}

func newMsgBatchCreateSpotLimitOrdersGasEstimator() *MsgBatchCreateSpotLimitOrdersGasEstimator {
	return &MsgBatchCreateSpotLimitOrdersGasEstimator{}
}

func (e *MsgBatchCreateSpotLimitOrdersGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgBatchCreateSpotLimitOrders)
	return ok
}

func (e *MsgBatchCreateSpotLimitOrdersGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// type assert exchangetypes.MsgBatchCreateSpotLimitOrders
	batchCreateMsg := msg.(*exchangetypes.MsgBatchCreateSpotLimitOrders)
	var total uint64
	postOnlyOrders := selectPostOnlyOrders(batchCreateMsg.Orders)
	total += GENERAL_MESSAGE_GAS_LIMIT
	total += uint64(len(batchCreateMsg.Orders)) * SPOT_ORDER_CREATION_GAS_LIMIT
	total += uint64(math.Ceil(float64(len(postOnlyOrders)) * float64(SPOT_ORDER_CREATION_GAS_LIMIT) * SPOT_POST_ONLY_ORDER_MULTIPLIER))
	return total
}

// MsgCancelSpotOrderGasEstimator
type MsgCancelSpotOrderGasEstimator struct {
}

func newMsgCancelSpotOrderGasEstimator() *MsgCancelSpotOrderGasEstimator {
	return &MsgCancelSpotOrderGasEstimator{}
}

func (e *MsgCancelSpotOrderGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgCancelSpotOrder)
	return ok
}

func (e *MsgCancelSpotOrderGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	return SPOT_ORDER_CANCELATION_GAS_LIMIT
}

// MsgBatchCancelSpotOrdersGasEstimator
type MsgBatchCancelSpotOrdersGasEstimator struct {
}

func newMsgBatchCancelSpotOrdersGasEstimator() *MsgBatchCancelSpotOrdersGasEstimator {
	return &MsgBatchCancelSpotOrdersGasEstimator{}
}

func (e *MsgBatchCancelSpotOrdersGasEstimator) appliesTo(msg sdk.Msg) bool {
	_, ok := msg.(*exchangetypes.MsgBatchCancelSpotOrders)
	return ok
}

func (e *MsgBatchCancelSpotOrdersGasEstimator) estimateMsgGas(msg sdk.Msg) uint64 {
	// type assert exchangetypes.MsgBatchCancelSpotOrders
	batchCancelMsg := msg.(*exchangetypes.MsgBatchCancelSpotOrders)
	var total uint64
	total += GENERAL_MESSAGE_GAS_LIMIT
	total += uint64(len(batchCancelMsg.Data)) * SPOT_ORDER_CANCELATION_GAS_LIMIT // Calculate gas based on number of orders to cancel
	return total
}

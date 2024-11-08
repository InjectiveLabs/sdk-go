package v2

import (
	"cosmossdk.io/math"
	"github.com/cosmos/gogoproto/proto"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	v1 "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

func (t OrderType) IsBuy() bool {
	switch t {
	case OrderType_BUY, OrderType_STOP_BUY, OrderType_TAKE_BUY, OrderType_BUY_PO, OrderType_BUY_ATOMIC:
		return true
	case OrderType_SELL, OrderType_STOP_SELL, OrderType_TAKE_SELL, OrderType_SELL_PO, OrderType_SELL_ATOMIC:
		return false
	}
	return false
}

func (t OrderType) IsPostOnly() bool {
	switch t {
	case OrderType_BUY_PO, OrderType_SELL_PO:
		return true
	default:
		return false
	}
}

func (t OrderType) IsConditional() bool {
	switch t {
	case OrderType_STOP_BUY,
		OrderType_STOP_SELL,
		OrderType_TAKE_BUY,
		OrderType_TAKE_SELL:
		return true
	}
	return false
}

func (t OrderType) IsAtomic() bool {
	switch t {
	case OrderType_BUY_ATOMIC,
		OrderType_SELL_ATOMIC:
		return true
	}
	return false
}

func NewV2OrderInfoFromV1(market MarketInterface, orderInfo v1.OrderInfo) *OrderInfo {
	humanPrice := market.PriceFromChainFormat(orderInfo.Price)
	humanQuantity := market.QuantityFromChainFormat(orderInfo.Quantity)

	return &OrderInfo{
		SubaccountId: orderInfo.SubaccountId,
		FeeRecipient: orderInfo.FeeRecipient,
		Price:        humanPrice,
		Quantity:     humanQuantity,
		Cid:          orderInfo.Cid,
	}
}

func (m *OrderInfo) GetNotional() math.LegacyDec {
	return m.Quantity.Mul(m.Price)
}

func (m *OrderInfo) GetFeeAmount(fee math.LegacyDec) math.LegacyDec {
	return m.GetNotional().Mul(fee)
}

func (m *OrderInfo) IsFromDefaultSubaccount() bool {
	return types.IsDefaultSubaccountID(common.HexToHash(m.SubaccountId))
}

func (m *OrderInfo) GetPrice() math.LegacyDec {
	return m.Price
}

func (m *OrderInfo) GetQuantity() math.LegacyDec {
	return m.Quantity
}

func (m *MarketStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MarketStatus_value, data, "MarketStatus")
	if err != nil {
		return err
	}
	*m = MarketStatus(value)
	return nil
}

func (m *ExecutionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExecutionType_value, data, "ExecutionType")
	if err != nil {
		return err
	}
	*m = ExecutionType(value)
	return nil
}

func (m *OrderType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OrderType_value, data, "OrderType")
	if err != nil {
		return err
	}
	*m = OrderType(value)
	return nil
}

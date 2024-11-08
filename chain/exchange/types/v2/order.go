package v2

import v1 "github.com/InjectiveLabs/sdk-go/chain/exchange/types"

func NewV1OrderInfoFromV2(market MarketInterface, orderInfo OrderInfo) v1.OrderInfo {
	return v1.OrderInfo{
		SubaccountId: orderInfo.SubaccountId,
		FeeRecipient: orderInfo.FeeRecipient,
		Price:        market.PriceToChainFormat(orderInfo.Price),
		Quantity:     market.QuantityToChainFormat(orderInfo.Quantity),
		Cid:          orderInfo.Cid,
	}
}

func NewV1SubaccountOrderDataFromV2(market MarketInterface, orderData *SubaccountOrderData) *v1.SubaccountOrderData {
	return &v1.SubaccountOrderData{
		Order: &v1.SubaccountOrder{
			Price:        market.PriceToChainFormat(orderData.Order.Price),
			Quantity:     market.QuantityToChainFormat(orderData.Order.Quantity),
			IsReduceOnly: orderData.Order.IsReduceOnly,
			Cid:          orderData.Order.Cid,
		},
		OrderHash: orderData.OrderHash,
	}
}

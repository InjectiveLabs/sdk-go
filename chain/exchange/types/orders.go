package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

var ZeroHash = common.Hash{}

func (o *SpotOrder) GetNewSpotLimitOrder(hash common.Hash) *SpotLimitOrder {
	return &SpotLimitOrder{
		OrderInfo:    o.OrderInfo,
		OrderType:    o.OrderType,
		Fillable:     o.OrderInfo.Quantity,
		TriggerPrice: o.TriggerPrice,
		Hash:         hash.Bytes(),
	}
}

type SpotOrderStateExpansion struct {
	BaseChangeAmount   sdk.Dec
	QuoteChangeAmount  sdk.Dec
	QuoteRefundAmount  sdk.Dec
	FeeRecipient       common.Address
	FeeRecipientReward sdk.Dec
	AuctionFeeReward   sdk.Dec
	Hash               common.Hash
	SubaccountID       common.Hash
	// for market orders, FillableAmount refers to the fillable quantity of the market order execution (if any)
	FillableAmount sdk.Dec
}

type SpotMarketBatchExecutionData struct {
	Market                         *SpotMarket
	BaseDenomDepositMap            map[common.Hash]*DepositDelta
	QuoteDenomDepositMap           map[common.Hash]*DepositDelta
	BaseDenomDepositSubaccountIDs  []common.Hash
	QuoteDenomDepositSubaccountIDs []common.Hash
	LimitOrderFilledDeltas         []*SpotLimitOrderFilledDelta
	MarketOrderExecutionEvent      *EventBatchSpotExecution
	LimitOrderExecutionEvent       []*EventBatchSpotExecution
	NewOrdersEvent                 *EventNewSpotOrders
}

type SpotLimitOrderFilledDelta struct {
	SubaccountIndexKey []byte
	FillableAmount     sdk.Dec
}

type DepositDelta struct {
	AvailableBalanceDelta sdk.Dec
	TotalBalanceDelta     sdk.Dec
}

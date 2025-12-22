// nolint // revive
package v2

import (
	"bytes"
	"sort"

	"cosmossdk.io/math"
	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/ethereum/go-ethereum/common"
)

type OrderFillType int

const (
	RestingLimitBuy    OrderFillType = 0
	RestingLimitSell   OrderFillType = 1
	TransientLimitBuy  OrderFillType = 2
	TransientLimitSell OrderFillType = 3
)

func (r *SpotOrderbookMatchingResults) GetOrderbookFills(fillType OrderFillType) *OrderbookFills {
	switch fillType {
	case RestingLimitBuy:
		return r.RestingBuyOrderbookFills
	case RestingLimitSell:
		return r.RestingSellOrderbookFills
	case TransientLimitBuy:
		return r.TransientBuyOrderbookFills
	case TransientLimitSell:
		return r.TransientSellOrderbookFills
	}

	return r.TransientSellOrderbookFills
}

type OrderbookFills struct {
	Orders         []*SpotLimitOrder
	FillQuantities []math.LegacyDec
}

type SpotOrderbookMatchingResults struct {
	TransientBuyOrderbookFills  *OrderbookFills
	RestingBuyOrderbookFills    *OrderbookFills
	TransientSellOrderbookFills *OrderbookFills
	RestingSellOrderbookFills   *OrderbookFills
	ClearingPrice               math.LegacyDec
	ClearingQuantity            math.LegacyDec
}

type SpotMarketFilter struct {
}

type FullSpotMarketFiller struct {
}

type SpotOrderStateExpansion struct {
	BaseChangeAmount        math.LegacyDec
	BaseRefundAmount        math.LegacyDec
	QuoteChangeAmount       math.LegacyDec
	QuoteRefundAmount       math.LegacyDec
	TradePrice              math.LegacyDec
	FeeRecipient            common.Address
	FeeRecipientReward      math.LegacyDec
	AuctionFeeReward        math.LegacyDec
	TraderFeeReward         math.LegacyDec
	TradingRewardPoints     math.LegacyDec
	LimitOrder              *SpotLimitOrder
	LimitOrderFillQuantity  math.LegacyDec
	MarketOrder             *SpotMarketOrder
	MarketOrderFillQuantity math.LegacyDec
	OrderHash               common.Hash
	OrderPrice              math.LegacyDec
	SubaccountID            common.Hash
	TraderAddress           string
	Cid                     string
}

type SpotBatchExecutionData struct {
	Market                         *SpotMarket
	BaseDenomDepositDeltas         types.DepositDeltas
	QuoteDenomDepositDeltas        types.DepositDeltas
	BaseDenomDepositSubaccountIDs  []common.Hash
	QuoteDenomDepositSubaccountIDs []common.Hash
	LimitOrderFilledDeltas         []*SpotLimitOrderDelta
	MarketOrderExecutionEvent      *EventBatchSpotExecution
	LimitOrderExecutionEvent       []*EventBatchSpotExecution
	NewOrdersEvent                 *EventNewSpotOrders
	TradingRewardPoints            types.TradingRewardPoints
	VwapData                       *SpotVwapData
}

func (e *SpotOrderStateExpansion) UpdateFromDepositDeltas(
	market *SpotMarket, baseDenomDepositDeltas, quoteDenomDepositDeltas types.DepositDeltas,
) {
	traderBaseDepositDelta := &types.DepositDelta{
		AvailableBalanceDelta: market.QuantityToChainFormat(e.BaseRefundAmount),
		TotalBalanceDelta:     market.QuantityToChainFormat(e.BaseChangeAmount),
	}

	traderQuoteDepositDelta := &types.DepositDelta{
		AvailableBalanceDelta: market.NotionalToChainFormat(e.QuoteRefundAmount),
		TotalBalanceDelta:     market.NotionalToChainFormat(e.QuoteChangeAmount),
	}

	if e.BaseChangeAmount.IsPositive() {
		traderBaseDepositDelta.AddAvailableBalance(market.QuantityToChainFormat(e.BaseChangeAmount))
	}

	if e.QuoteChangeAmount.IsPositive() {
		traderQuoteDepositDelta.AddAvailableBalance(market.NotionalToChainFormat(e.QuoteChangeAmount))
	}

	feeRecipientSubaccount := types.EthAddressToSubaccountID(e.FeeRecipient)
	if bytes.Equal(feeRecipientSubaccount.Bytes(), types.ZeroSubaccountID.Bytes()) {
		feeRecipientSubaccount = types.AuctionSubaccountID
	}

	baseDenomDepositDeltas.ApplyDepositDelta(e.SubaccountID, traderBaseDepositDelta)
	quoteDenomDepositDeltas.ApplyDepositDelta(e.SubaccountID, traderQuoteDepositDelta)

	quoteDenomDepositDeltas.ApplyUniformDelta(feeRecipientSubaccount, market.NotionalToChainFormat(e.FeeRecipientReward))
	quoteDenomDepositDeltas.ApplyUniformDelta(types.AuctionSubaccountID, market.NotionalToChainFormat(e.AuctionFeeReward))
}

func GetBatchExecutionEventsFromSpotLimitOrderStateExpansions(
	isBuy bool,
	market *SpotMarket,
	executionType ExecutionType,
	spotLimitOrderStateExpansions []*SpotOrderStateExpansion,
	baseDenomDepositDeltas,
	quoteDenomDepositDeltas types.DepositDeltas,
) (*EventBatchSpotExecution, []*SpotLimitOrderDelta, types.TradingRewardPoints) {
	limitOrderBatchEvent := &EventBatchSpotExecution{
		MarketId:      market.MarketID().Hex(),
		IsBuy:         isBuy,
		ExecutionType: executionType,
	}

	trades := make([]*TradeLog, 0, len(spotLimitOrderStateExpansions))

	// array of (SubaccountIndexKey, fillableAmount) to update/delete
	filledDeltas := make([]*SpotLimitOrderDelta, 0, len(spotLimitOrderStateExpansions))
	tradingRewardPoints := types.NewTradingRewardPoints()

	for idx := range spotLimitOrderStateExpansions {
		expansion := spotLimitOrderStateExpansions[idx]
		expansion.UpdateFromDepositDeltas(market, baseDenomDepositDeltas, quoteDenomDepositDeltas)

		// skip adding trade data if there was no trade (unfilled new order)
		fillQuantity := spotLimitOrderStateExpansions[idx].BaseChangeAmount
		if fillQuantity.IsZero() {
			continue
		}

		filledDeltas = append(filledDeltas, &SpotLimitOrderDelta{
			Order:        expansion.LimitOrder,
			FillQuantity: expansion.LimitOrderFillQuantity,
		})

		var realizedTradeFee math.LegacyDec

		isSelfRelayedTrade := expansion.FeeRecipient == types.SubaccountIDToEthAddress(expansion.SubaccountID)
		if isSelfRelayedTrade {
			// if negative fee, equals the full negative rebate
			// otherwise equals the fees going to auction
			realizedTradeFee = expansion.AuctionFeeReward
		} else {
			realizedTradeFee = expansion.FeeRecipientReward.Add(expansion.AuctionFeeReward)
		}

		tradingRewardPoints.AddPointsForAddress(expansion.TraderAddress, expansion.TradingRewardPoints)

		trades = append(trades, &TradeLog{
			Quantity:            expansion.BaseChangeAmount.Abs(),
			Price:               expansion.TradePrice,
			SubaccountId:        expansion.SubaccountID.Bytes(),
			Fee:                 realizedTradeFee,
			OrderHash:           expansion.OrderHash.Bytes(),
			FeeRecipientAddress: expansion.FeeRecipient.Bytes(),
			Cid:                 expansion.Cid,
		})
	}
	limitOrderBatchEvent.Trades = trades

	if len(trades) == 0 {
		limitOrderBatchEvent = nil
	}
	return limitOrderBatchEvent, filledDeltas, tradingRewardPoints
}

type SpotVwapData struct {
	Price    math.LegacyDec
	Quantity math.LegacyDec
}

func NewSpotVwapData() *SpotVwapData {
	return &SpotVwapData{
		Price:    math.LegacyZeroDec(),
		Quantity: math.LegacyZeroDec(),
	}
}

func (p *SpotVwapData) ApplyExecution(price, quantity math.LegacyDec) *SpotVwapData {
	if p == nil {
		p = NewSpotVwapData()
	}

	if price.IsNil() || quantity.IsNil() || quantity.IsZero() {
		return p
	}

	newQuantity := p.Quantity.Add(quantity)
	newPrice := p.Price.Mul(p.Quantity).Add(price.Mul(quantity)).Quo(newQuantity)

	return &SpotVwapData{
		Price:    newPrice,
		Quantity: newQuantity,
	}
}

type SpotVwapInfo map[common.Hash]*SpotVwapData

func NewSpotVwapInfo() SpotVwapInfo {
	return make(SpotVwapInfo)
}

func (p *SpotVwapInfo) ApplyVwap(marketID common.Hash, newVwapData *SpotVwapData) {
	var existingVwapData *SpotVwapData

	existingVwapData = (*p)[marketID]
	if existingVwapData == nil {
		existingVwapData = NewSpotVwapData()
		(*p)[marketID] = existingVwapData
	}

	if !newVwapData.Quantity.IsZero() {
		(*p)[marketID] = existingVwapData.ApplyExecution(newVwapData.Price, newVwapData.Quantity)
	}
}

func (p *SpotVwapInfo) GetSortedSpotMarketIDs() []common.Hash {
	spotMarketIDs := make([]common.Hash, 0)
	for k := range *p {
		spotMarketIDs = append(spotMarketIDs, k)
	}

	sort.SliceStable(spotMarketIDs, func(i, j int) bool {
		return bytes.Compare(spotMarketIDs[i].Bytes(), spotMarketIDs[j].Bytes()) < 0
	})
	return spotMarketIDs
}

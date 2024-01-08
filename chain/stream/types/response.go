package types

import (
	"fmt"
	"time"
)

type StreamResponseMap struct {
	tradeEventsCounter                   uint64
	BlockHeight                          uint64
	BlockTime                            time.Time
	BankBalancesByAccount                map[string][]*BankBalance
	SpotOrdersBySubaccount               map[string][]*SpotOrderUpdate
	SpotOrdersByMarketID                 map[string][]*SpotOrderUpdate
	DerivativeOrdersBySubaccount         map[string][]*DerivativeOrderUpdate
	DerivativeOrdersByMarketID           map[string][]*DerivativeOrderUpdate
	SpotOrderbookUpdatesByMarketID       map[string][]*OrderbookUpdate
	DerivativeOrderbookUpdatesByMarketID map[string][]*OrderbookUpdate
	SubaccountDepositsBySubaccountID     map[string][]*SubaccountDeposits
	SpotTradesBySubaccount               map[string][]*SpotTrade
	SpotTradesByMarketID                 map[string][]*SpotTrade
	DerivativeTradesBySubaccount         map[string][]*DerivativeTrade
	DerivativeTradesByMarketID           map[string][]*DerivativeTrade
	PositionsBySubaccount                map[string][]*Position
	PositionsByMarketID                  map[string][]*Position
	OraclePriceBySymbol                  map[string][]*OraclePrice
}

func NewStreamResponseMap() *StreamResponseMap {
	return &StreamResponseMap{
		BankBalancesByAccount:                map[string][]*BankBalance{},
		SpotOrdersBySubaccount:               map[string][]*SpotOrderUpdate{},
		SpotOrdersByMarketID:                 map[string][]*SpotOrderUpdate{},
		DerivativeOrdersBySubaccount:         map[string][]*DerivativeOrderUpdate{},
		DerivativeOrdersByMarketID:           map[string][]*DerivativeOrderUpdate{},
		SpotOrderbookUpdatesByMarketID:       map[string][]*OrderbookUpdate{},
		DerivativeOrderbookUpdatesByMarketID: map[string][]*OrderbookUpdate{},
		SubaccountDepositsBySubaccountID:     map[string][]*SubaccountDeposits{},
		SpotTradesBySubaccount:               map[string][]*SpotTrade{},
		SpotTradesByMarketID:                 map[string][]*SpotTrade{},
		DerivativeTradesBySubaccount:         map[string][]*DerivativeTrade{},
		DerivativeTradesByMarketID:           map[string][]*DerivativeTrade{},
		PositionsBySubaccount:                map[string][]*Position{},
		PositionsByMarketID:                  map[string][]*Position{},
		OraclePriceBySymbol:                  map[string][]*OraclePrice{},
	}
}

func NewChainStreamResponse() *StreamResponse {
	return &StreamResponse{
		BankBalances:               []*BankBalance{},
		SubaccountDeposits:         []*SubaccountDeposits{},
		SpotTrades:                 []*SpotTrade{},
		DerivativeTrades:           []*DerivativeTrade{},
		SpotOrders:                 []*SpotOrderUpdate{},
		DerivativeOrders:           []*DerivativeOrderUpdate{},
		SpotOrderbookUpdates:       []*OrderbookUpdate{},
		DerivativeOrderbookUpdates: []*OrderbookUpdate{},
		Positions:                  []*Position{},
		OraclePrices:               []*OraclePrice{},
	}
}

func (m *StreamRequest) Validate() error {
	if m.BankBalancesFilter == nil &&
		m.SubaccountDepositsFilter == nil &&
		m.SpotTradesFilter == nil &&
		m.DerivativeTradesFilter == nil &&
		m.SpotOrdersFilter == nil &&
		m.DerivativeOrdersFilter == nil &&
		m.SpotOrderbooksFilter == nil &&
		m.DerivativeOrderbooksFilter == nil &&
		m.PositionsFilter == nil &&
		m.OraclePriceFilter == nil {
		return fmt.Errorf("at least one filter must be set")
	}
	return nil
}

func (m *StreamResponseMap) NextTradeEventNumber() (tradeNumber uint64) {
	currentTradesNumber := m.tradeEventsCounter
	m.tradeEventsCounter++
	return currentTradesNumber
}

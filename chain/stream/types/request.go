package types

func NewFullStreamRequest() *StreamRequest {
	return &StreamRequest{
		BankBalancesFilter: &BankBalancesFilter{
			Accounts: []string{},
		},
		SpotOrdersFilter: &OrdersFilter{
			MarketIds:     []string{"*"},
			SubaccountIds: []string{"*"},
		},
		DerivativeOrdersFilter: &OrdersFilter{
			MarketIds:     []string{"*"},
			SubaccountIds: []string{"*"},
		},
		SpotTradesFilter: &TradesFilter{
			MarketIds:     []string{"*"},
			SubaccountIds: []string{"*"},
		},
		SubaccountDepositsFilter: &SubaccountDepositsFilter{
			SubaccountIds: []string{"*"},
		},
		DerivativeOrderbooksFilter: &OrderbookFilter{
			MarketIds: []string{"*"},
		},
		SpotOrderbooksFilter: &OrderbookFilter{
			MarketIds: []string{"*"},
		},
		PositionsFilter: &PositionsFilter{
			SubaccountIds: []string{"*"},
			MarketIds:     []string{"*"},
		},
		DerivativeTradesFilter: &TradesFilter{
			SubaccountIds: []string{"*"},
			MarketIds:     []string{"*"},
		},
		OraclePriceFilter: &OraclePriceFilter{
			Symbol: []string{"*"},
		},
	}
}

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

// Empty query matches any set of events.
type Empty struct {
}

// Matches always returns true.
func (Empty) Matches(tags map[string][]string) (bool, error) {
	return true, nil
}

func (Empty) String() string {
	return "empty"
}

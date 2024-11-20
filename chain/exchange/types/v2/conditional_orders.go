package v2

func (b *ConditionalDerivativeOrderBook) HasLimitBuyOrders() bool {
	return len(b.LimitBuyOrders) > 0
}

func (b *ConditionalDerivativeOrderBook) HasLimitSellOrders() bool {
	return len(b.LimitSellOrders) > 0
}

func (b *ConditionalDerivativeOrderBook) HasMarketBuyOrders() bool {
	return len(b.MarketBuyOrders) > 0
}

func (b *ConditionalDerivativeOrderBook) HasMarketSellOrders() bool {
	return len(b.MarketSellOrders) > 0
}

func (b *ConditionalDerivativeOrderBook) IsEmpty() bool {
	return !b.HasLimitBuyOrders() && !b.HasLimitSellOrders() && !b.HasMarketBuyOrders() && b.HasMarketSellOrders()
}

func (b *ConditionalDerivativeOrderBook) GetMarketOrders() []*DerivativeMarketOrder {
	return append(b.MarketBuyOrders, b.MarketSellOrders...)
}

func (b *ConditionalDerivativeOrderBook) GetLimitOrders() []*DerivativeLimitOrder {
	return append(b.LimitBuyOrders, b.LimitSellOrders...)
}

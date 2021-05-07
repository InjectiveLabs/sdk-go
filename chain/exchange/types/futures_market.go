package types

func (m *ExpiryFuturesMarketInfo) IsPremature(currBlockTime int64) bool {
	return currBlockTime < m.TwapStartTimestamp
}

func (m *ExpiryFuturesMarketInfo) IsStartingMaturation(currBlockTime int64) bool {
	return currBlockTime >= m.TwapStartTimestamp && m.ExpirationTwapStartPriceCumulative.IsNil()
}

func (m *ExpiryFuturesMarketInfo) IsMatured(currBlockTime int64) bool {
	return currBlockTime >= m.ExpirationTimestamp
}

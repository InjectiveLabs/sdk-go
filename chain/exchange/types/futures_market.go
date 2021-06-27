package types

func (m *ExpiryFuturesMarketInfo) IsPremature(currBlockTime int64) bool {
	return currBlockTime < m.TwapStartTimestamp
}

func (m *ExpiryFuturesMarketInfo) IsStartingMaturation(currBlockTime int64) bool {
	return currBlockTime >= m.TwapStartTimestamp && (m.ExpirationTwapStartPriceCumulative.IsNil() || m.ExpirationTwapStartPriceCumulative.IsZero())
}

func (m *ExpiryFuturesMarketInfo) IsMatured(currBlockTime int64) bool {
	return currBlockTime >= m.ExpirationTimestamp
}

package types

func (m *SpotMarket) IsActive() bool {
	return m.Status == MarketStatus_Active
}

func (m *SpotMarket) IsInactive() bool {
	return !m.IsActive()
}

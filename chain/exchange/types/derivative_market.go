package types

func (m *DerivativeMarket) IsTimeExpiry() bool {
	return !m.IsPerpetual
}

func (m *DerivativeMarket) IsActive() bool {
	return m.Status == MarketStatus_Active
}

func (m *DerivativeMarket) IsInactive() bool {
	return !m.IsActive()
}

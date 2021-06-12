package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type MatchedMarketDirection struct {
	MarketId    common.Hash
	BuysExists  bool
	SellsExists bool
}

func (m *DerivativeMarket) MarketID() common.Hash {
	return common.HexToHash(m.MarketId)
}

func (m *SpotMarket) MarketID() common.Hash {
	return common.HexToHash(m.MarketId)
}

func (e ExecutionType) IsMarket() bool {
	return e == ExecutionType_Market
}

func (m *SpotMarket) StatusSupportsOrderCancellations() bool {
	if m == nil {
		return false
	}
	return m.Status.supportsOrderCancellations()
}

func (m *DerivativeMarket) StatusSupportsOrderCancellations() bool {
	if m == nil {
		return false
	}
	return m.Status.supportsOrderCancellations()
}

func (s MarketStatus) supportsOrderCancellations() bool {
	switch s {
	case MarketStatus_Active, MarketStatus_Suspended, MarketStatus_Demolished, MarketStatus_Expired:
		return true
	case MarketStatus_Paused:
		return false
	default:
		return false
	}
}

package types

import (
	"bytes"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type MarketSummary struct {
	TotalUserQuantity     sdk.Dec
	TotalContractQuantity sdk.Dec
	TotalUserMargin       sdk.Dec
	TotalContractMargin   sdk.Dec
	netQuantity           sdk.Dec
}

func NewMarketSummary() *MarketSummary {
	return &MarketSummary{
		TotalUserQuantity:     sdk.ZeroDec(),
		TotalContractQuantity: sdk.ZeroDec(),
		TotalUserMargin:       sdk.ZeroDec(),
		TotalContractMargin:   sdk.ZeroDec(),
		netQuantity:           sdk.ZeroDec(),
	}
}

func NewSyntheticTradeActionSummary() *SyntheticTradeActionSummary {
	return &SyntheticTradeActionSummary{
		MarketSummary: make(map[common.Hash]*MarketSummary),
		MarketIDs:     make([]common.Hash, 0),
	}
}

type SyntheticTradeActionSummary struct {
	MarketSummary map[common.Hash]*MarketSummary
	MarketIDs     []common.Hash
}

func (s *SyntheticTradeActionSummary) GetMarketIDs() []common.Hash {
	var marketIDs []common.Hash
	for marketID := range s.MarketSummary {
		marketIDs = append(marketIDs, marketID)
	}

	sort.SliceStable(marketIDs, func(i, j int) bool {
		return bytes.Compare(marketIDs[i].Bytes(), marketIDs[j].Bytes()) < 0
	})
	s.MarketIDs = marketIDs
	return marketIDs
}

func (s SyntheticTradeActionSummary) Update(t *SyntheticTrade, isForUser bool) {
	if _, ok := s.MarketSummary[t.MarketID]; !ok {
		s.MarketSummary[t.MarketID] = NewMarketSummary()
	}
	summary := s.MarketSummary[t.MarketID]

	if t.IsBuy {
		summary.netQuantity = summary.netQuantity.Add(t.Quantity)
	} else {
		summary.netQuantity = summary.netQuantity.Sub(t.Quantity)
	}

	if isForUser {
		summary.TotalUserQuantity = summary.TotalUserQuantity.Add(t.Quantity)
		summary.TotalUserMargin = summary.TotalUserMargin.Add(t.Margin)
	} else {
		summary.TotalContractQuantity = summary.TotalContractQuantity.Add(t.Quantity)
		summary.TotalContractMargin = summary.TotalContractMargin.Add(t.Margin)
	}
}

// IsValid checks that all the net quantities are zero
func (s SyntheticTradeActionSummary) IsValid() bool {
	for _, v := range s.MarketSummary {
		if !v.netQuantity.IsZero() {
			return false
		}
	}
	return true
}

func (a *SyntheticTradeAction) Summarize() (*SyntheticTradeActionSummary, error) {
	summary := NewSyntheticTradeActionSummary()

	for _, t := range a.UserTrades {
		summary.Update(t, true)
	}

	for _, t := range a.ContractTrades {
		summary.Update(t, false)
	}

	// ensure that sum(buy quantity) == sum(sell quantity) for all markets
	if !summary.IsValid() {
		return nil, ErrInvalidQuantity
	}

	summary.GetMarketIDs()
	return summary, nil
}

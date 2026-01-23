package types

import (
	"bytes"
	"sort"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type SyntheticTradeActionSummary struct {
	Markets         map[common.Hash]bool
	ContractAddress sdk.Address
	UserAddress     sdk.Address
}

func (s *SyntheticTradeActionSummary) GetMarketIDs() []common.Hash {
	marketIDs := make([]common.Hash, 0, len(s.Markets))

	for marketID := range s.Markets {
		marketIDs = append(marketIDs, marketID)
	}

	sort.SliceStable(marketIDs, func(i, j int) bool {
		return bytes.Compare(marketIDs[i].Bytes(), marketIDs[j].Bytes()) < 0
	})

	return marketIDs
}

func (s *SyntheticTradeActionSummary) Update(t *SyntheticTrade, isForUser bool) error {
	if _, ok := s.Markets[t.MarketID]; !ok {
		s.Markets[t.MarketID] = true
	}

	address := SubaccountIDToSdkAddress(t.SubaccountID)

	if (isForUser && !address.Equals(s.UserAddress)) || (!isForUser && !address.Equals(s.ContractAddress)) {
		return ErrBadSubaccountID
	}

	return nil
}

func (a *SyntheticTradeAction) Summarize() (*SyntheticTradeActionSummary, error) {
	if err := a.validateTrades(); err != nil {
		return nil, err
	}

	summary := a.initSummary()

	if err := a.updateSummary(&summary); err != nil {
		return nil, err
	}

	return &summary, nil
}

func (a *SyntheticTradeAction) initSummary() SyntheticTradeActionSummary {
	return SyntheticTradeActionSummary{
		Markets:         make(map[common.Hash]bool),
		ContractAddress: SubaccountIDToSdkAddress(a.ContractTrades[0].SubaccountID),
		UserAddress:     SubaccountIDToSdkAddress(a.UserTrades[0].SubaccountID),
	}
}

func (a *SyntheticTradeAction) validateTrades() error {
	if len(a.UserTrades) == 0 || len(a.ContractTrades) == 0 {
		return errors.Wrapf(ErrInvalidTrade, "no trades in action")
	}

	if len(a.UserTrades) != len(a.ContractTrades) {
		return errors.Wrapf(
			ErrInvalidTrade,
			"mismatched user and contract trades: %d vs %d",
			len(a.UserTrades),
			len(a.ContractTrades),
		)
	}

	for i, userTrade := range a.UserTrades {
		contractTrade := a.ContractTrades[i]

		if userTrade.MarketID != contractTrade.MarketID {
			return errors.Wrapf(ErrInvalidTrade, "mismatched user and contract trade at index %d", i)
		}
		if userTrade.IsBuy == contractTrade.IsBuy {
			return errors.Wrapf(ErrInvalidTrade, "mismatched user and contract trade at index %d", i)
		}
		if !userTrade.Quantity.Equal(contractTrade.Quantity) {
			return errors.Wrapf(ErrInvalidTrade, "mismatched user and contract trade at index %d", i)
		}
		if !userTrade.Price.Equal(contractTrade.Price) {
			return errors.Wrapf(ErrInvalidTrade, "mismatched user and contract trade at index %d", i)
		}
	}

	return nil
}

func (a *SyntheticTradeAction) updateSummary(summary *SyntheticTradeActionSummary) error {
	for _, t := range a.UserTrades {
		if err := summary.Update(t, true); err != nil {
			return err
		}
	}

	for _, t := range a.ContractTrades {
		if err := summary.Update(t, false); err != nil {
			return err
		}
	}

	return nil
}

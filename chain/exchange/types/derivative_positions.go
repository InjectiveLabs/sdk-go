package types

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"sort"
)

type PositionStates map[common.Hash]*PositionState

func NewPositionStates() PositionStates {
	return make(PositionStates)
}

func (p *PositionStates) GetSortedSubaccountKeys() []common.Hash {
	subaccountKeys := make([]common.Hash, 0)
	for k := range *p {
		subaccountKeys = append(subaccountKeys, k)
	}
	sort.SliceStable(subaccountKeys, func(i, j int) bool {
		return bytes.Compare(subaccountKeys[i].Bytes(), subaccountKeys[j].Bytes()) < 0
	})
	return subaccountKeys
}

func (p *PositionStates) GetPositionUpdateEvent(
	marketID common.Hash,
	funding *PerpetualMarketFunding,
) (*EventBatchDerivativePosition, []*Position, []common.Hash) {
	positionSubaccountIDs := p.GetSortedSubaccountKeys()
	positions := make([]*Position, len(positionSubaccountIDs))
	positionLogs := make([]*DerivativePositionLog, len(positionSubaccountIDs))

	for idx := range positionSubaccountIDs {
		subaccountID := positionSubaccountIDs[idx]
		position := (*p)[subaccountID]
		positions[idx] = position.Position
		positionLogs[idx] = &DerivativePositionLog{
			SubaccountId:   subaccountID.Hex(),
			Position:       position.Position,
			FundingPayment: position.FundingPayment,
		}
	}

	event := &EventBatchDerivativePosition{
		MarketId:          marketID.Hex(),
		CumulativeFunding: &funding.CumulativeFunding,
		Positions:         positionLogs,
	}

	return event, positions, positionSubaccountIDs
}

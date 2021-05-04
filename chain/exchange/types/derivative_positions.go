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

func (p *PositionStates) GetPositionSliceData() ([]*Position, []common.Hash) {
	positionSubaccountIDs := p.GetSortedSubaccountKeys()
	positions := make([]*Position, len(positionSubaccountIDs))

	for idx := range positionSubaccountIDs {
		subaccountID := positionSubaccountIDs[idx]
		position := (*p)[subaccountID]
		positions[idx] = position.Position
	}

	return positions, positionSubaccountIDs
}

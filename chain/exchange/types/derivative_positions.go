package types

import (
	"bytes"
	"sort"

	"github.com/ethereum/go-ethereum/common"
)

func NewPositionStates() map[common.Hash]*PositionState {
	return make(map[common.Hash]*PositionState)
}

func GetSortedSubaccountKeys(p map[common.Hash]*PositionState) []common.Hash {
	subaccountKeys := make([]common.Hash, 0)
	for k := range p {
		subaccountKeys = append(subaccountKeys, k)
	}
	sort.SliceStable(subaccountKeys, func(i, j int) bool {
		return bytes.Compare(subaccountKeys[i].Bytes(), subaccountKeys[j].Bytes()) < 0
	})
	return subaccountKeys
}

func GetPositionSliceData(p map[common.Hash]*PositionState) ([]*Position, []common.Hash) {
	positionSubaccountIDs := GetSortedSubaccountKeys(p)
	positions := make([]*Position, 0, len(positionSubaccountIDs))

	nonNilPositionSubaccountIDs := make([]common.Hash, 0)
	for idx := range positionSubaccountIDs {
		subaccountID := positionSubaccountIDs[idx]
		position := p[subaccountID]
		if position.Position != nil {
			positions = append(positions, position.Position)
			nonNilPositionSubaccountIDs = append(nonNilPositionSubaccountIDs, subaccountID)
		} else {
			// fmt.Println("❌ position is nil for subaccount", subaccountID.Hex())
		}
	}

	return positions, nonNilPositionSubaccountIDs
}

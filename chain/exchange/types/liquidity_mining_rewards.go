package types

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LiquidityMiningRewards map[string]sdk.Dec

func NewLiquidityMiningRewards() LiquidityMiningRewards {
	liquidityMiningRewards := make(LiquidityMiningRewards)

	return liquidityMiningRewards
}

func (l LiquidityMiningRewards) GetSortedAccountKeys() []string {
	accountKeys := make([]string, 0, len(l))
	for k := range l {
		accountKeys = append(accountKeys, k)
	}
	sort.Strings(accountKeys)
	return accountKeys
}

func (l LiquidityMiningRewards) AddPointsForAddress(addr string, newPoints sdk.Dec) {
	v, found := l[addr]
	if !found {
		l[addr] = newPoints
	} else {
		l[addr] = v.Add(newPoints)
	}
}

func (l *LiquidityMiningRewards) GetAllAccountAddresses() []string {
	accountAddresses := make([]string, 0)

	for k := range *l {
		accountAddresses = append(accountAddresses, k)
	}

	return accountAddresses
}

func MergeLiquidityMiningRewards(m1, m2 LiquidityMiningRewards) LiquidityMiningRewards {
	if m1 == nil || len(m1) == 0 {
		return m2
	} else if m2 == nil || len(m2) == 0 {
		return m1
	}

	if len(m1) >= len(m2) {
		for k, v := range m2 {
			m1.AddPointsForAddress(k, v)
		}
		return m1
	} else {
		for k, v := range m1 {
			m2.AddPointsForAddress(k, v)
		}
		return m2
	}

}

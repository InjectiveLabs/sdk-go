package types

import (
	"sort"

	"cosmossdk.io/math"
)

type TradingRewardPoints map[string]math.LegacyDec

func NewTradingRewardPoints() TradingRewardPoints {
	tradingRewardPoints := make(TradingRewardPoints)

	return tradingRewardPoints
}

func (l TradingRewardPoints) GetSortedAccountKeys() []string {
	accountKeys := make([]string, 0, len(l))
	for k := range l {
		accountKeys = append(accountKeys, k)
	}
	sort.Strings(accountKeys)
	return accountKeys
}

func (l TradingRewardPoints) AddPointsForAddress(addr string, newPoints math.LegacyDec) {
	if !newPoints.IsPositive() {
		return
	}
	v, found := l[addr]
	if !found {
		l[addr] = newPoints
	} else {
		l[addr] = v.Add(newPoints)
	}
}

func (l *TradingRewardPoints) GetAllAccountAddresses() []string {
	accountAddresses := make([]string, 0)

	for k := range *l {
		accountAddresses = append(accountAddresses, k)
	}

	return accountAddresses
}

func MergeTradingRewardPoints(m1, m2 TradingRewardPoints) TradingRewardPoints {
	if len(m1) == 0 {
		return m2
	} else if len(m2) == 0 {
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

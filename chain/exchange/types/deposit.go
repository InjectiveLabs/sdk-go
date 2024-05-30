package types

import (
	"bytes"
	"fmt"
	"sort"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
)

func NewDeposit() *Deposit {
	return &Deposit{
		AvailableBalance: math.LegacyZeroDec(),
		TotalBalance:     math.LegacyZeroDec(),
	}
}

func (d *Deposit) IsEmpty() bool {
	return d.AvailableBalance.IsZero() && d.TotalBalance.IsZero()
}

func (d *Deposit) Display() string {
	return fmt.Sprintf("Deposit Available: %s, Total: %s", getReadableDec(d.AvailableBalance), getReadableDec(d.TotalBalance))
}

type DepositDelta struct {
	AvailableBalanceDelta math.LegacyDec
	TotalBalanceDelta     math.LegacyDec
}

func NewUniformDepositDelta(delta math.LegacyDec) *DepositDelta {
	return &DepositDelta{
		AvailableBalanceDelta: delta,
		TotalBalanceDelta:     delta,
	}
}

func NewDepositDelta() *DepositDelta {
	return NewUniformDepositDelta(math.LegacyZeroDec())
}

func (d *DepositDelta) AddAvailableBalance(amount math.LegacyDec) {
	d.AvailableBalanceDelta = d.AvailableBalanceDelta.Add(amount)
}

func (d *DepositDelta) IsEmpty() bool {
	if d == nil {
		return true
	}
	hasEmptyTotalBalanceDelta := d.TotalBalanceDelta.IsNil() || d.TotalBalanceDelta.IsZero()
	hasEmptyAvailableBalanceDelta := d.AvailableBalanceDelta.IsNil() || d.AvailableBalanceDelta.IsZero()
	return hasEmptyTotalBalanceDelta && hasEmptyAvailableBalanceDelta
}

type DepositDeltas map[common.Hash]*DepositDelta

func NewDepositDeltas() DepositDeltas {
	return make(DepositDeltas)
}

func (d *DepositDeltas) GetSortedSubaccountKeys() []common.Hash {
	subaccountKeys := make([]common.Hash, 0)
	for k := range *d {
		subaccountKeys = append(subaccountKeys, k)
	}
	sort.SliceStable(subaccountKeys, func(i, j int) bool {
		return bytes.Compare(subaccountKeys[i].Bytes(), subaccountKeys[j].Bytes()) < 0
	})
	return subaccountKeys
}

func (d *DepositDeltas) ApplyDepositDelta(subaccountID common.Hash, delta *DepositDelta) {
	d.ApplyDelta(subaccountID, delta.TotalBalanceDelta, delta.AvailableBalanceDelta)
}

func (d *DepositDeltas) ApplyUniformDelta(subaccountID common.Hash, delta math.LegacyDec) {
	d.ApplyDelta(subaccountID, delta, delta)
}

func (d *DepositDeltas) ApplyDelta(subaccountID common.Hash, totalBalanceDelta, availableBalanceDelta math.LegacyDec) {
	delta := (*d)[subaccountID]
	if delta == nil {
		delta = NewDepositDelta()
		(*d)[subaccountID] = delta
	}
	delta.AvailableBalanceDelta = delta.AvailableBalanceDelta.Add(availableBalanceDelta)
	delta.TotalBalanceDelta = delta.TotalBalanceDelta.Add(totalBalanceDelta)
}

func (d *Deposit) HasTransientOrRestingVanillaLimitOrders() bool {
	return d.AvailableBalance.LT(d.TotalBalance)
}

func GetSortedBalanceKeys(p map[string]*Deposit) []string {
	denoms := make([]string, 0)
	for k := range p {
		denoms = append(denoms, k)
	}
	sort.SliceStable(denoms, func(i, j int) bool {
		return denoms[i] < denoms[j]
	})
	return denoms
}

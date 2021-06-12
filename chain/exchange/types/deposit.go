package types

import (
	"bytes"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func NewDeposit() *Deposit {
	return &Deposit{
		AvailableBalance: sdk.ZeroDec(),
		TotalBalance:     sdk.ZeroDec(),
	}
}

func (d *Deposit) IsEmpty() bool {
	return d.AvailableBalance.IsZero() && d.TotalBalance.IsZero()
}

type DepositDelta struct {
	AvailableBalanceDelta sdk.Dec
	TotalBalanceDelta     sdk.Dec
}

func NewUniformDepositDelta(delta sdk.Dec) *DepositDelta {
	return &DepositDelta{
		AvailableBalanceDelta: delta,
		TotalBalanceDelta:     delta,
	}
}

func NewDepositDelta() *DepositDelta {
	return NewUniformDepositDelta(sdk.ZeroDec())
}

func (d *DepositDelta) AddAvailableBalance(amount sdk.Dec) {
	d.AvailableBalanceDelta = d.AvailableBalanceDelta.Add(amount)
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

func (d *DepositDeltas) ApplyUniformDelta(subaccountID common.Hash, delta sdk.Dec) {
	d.ApplyDelta(subaccountID, delta, delta)
}

func (d *DepositDeltas) ApplyDelta(subaccountID common.Hash, totalBalanceDelta, availableBalanceDelta sdk.Dec) {
	delta := (*d)[subaccountID]
	if delta == nil {
		delta = NewDepositDelta()
		(*d)[subaccountID] = delta
	}
	delta.AvailableBalanceDelta = delta.AvailableBalanceDelta.Add(availableBalanceDelta)
	delta.TotalBalanceDelta = delta.TotalBalanceDelta.Add(totalBalanceDelta)
}

func (d *Deposit) HasRestingVanillaLimitOrders() bool {
	return d.AvailableBalance.LT(d.TotalBalance)
}

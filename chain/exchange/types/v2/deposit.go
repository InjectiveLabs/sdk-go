package v2

import (
	"fmt"

	"cosmossdk.io/math"
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

func (d *Deposit) HasTransientOrRestingVanillaLimitOrders() bool {
	return d.AvailableBalance.LT(d.TotalBalance)
}

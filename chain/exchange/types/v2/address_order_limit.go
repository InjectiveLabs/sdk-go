package v2

import (
	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	"github.com/ethereum/go-ethereum/common"
)

// MaxOrdersPerAddress bounds retained order-book orders across all subaccounts.
// RFQ quotes and immediately executed atomic orders do not occupy the book.
const MaxOrdersPerAddress uint64 = 5_000

// ValidateAddressOrderLimit rejects genesis books that bypass runtime admission.
//
//nolint:revive // All four persistent order classes contribute to one owner count.
func (gs GenesisState) ValidateAddressOrderLimit() error {
	counts := make(map[common.Address]uint64)
	check := func(id string) error {
		subaccountID := common.HexToHash(id)
		address := common.BytesToAddress(subaccountID[:common.AddressLength])
		counts[address]++
		if counts[address] > MaxOrdersPerAddress {
			return types.ErrTooManyOrdersPerAddress.Wrapf("genesis address %s exceeds %d", address.Hex(), MaxOrdersPerAddress)
		}
		return nil
	}
	for _, book := range gs.SpotOrderbook {
		for _, order := range book.Orders {
			if order != nil {
				if err := check(order.OrderInfo.SubaccountId); err != nil {
					return err
				}
			}
		}
	}
	for _, book := range gs.DerivativeOrderbook {
		for _, order := range book.Orders {
			if order != nil {
				if err := check(order.OrderInfo.SubaccountId); err != nil {
					return err
				}
			}
		}
	}
	for _, book := range gs.ConditionalDerivativeOrderbooks {
		if book == nil {
			continue
		}
		for _, orders := range [][]*DerivativeLimitOrder{book.LimitBuyOrders, book.LimitSellOrders} {
			for _, order := range orders {
				if order != nil {
					if err := check(order.OrderInfo.SubaccountId); err != nil {
						return err
					}
				}
			}
		}
		for _, orders := range [][]*DerivativeMarketOrder{book.MarketBuyOrders, book.MarketSellOrders} {
			for _, order := range orders {
				if order != nil {
					if err := check(order.OrderInfo.SubaccountId); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

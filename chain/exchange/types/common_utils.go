package types

import (
	"bytes"
	"math/big"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/olekukonko/tablewriter"

	"github.com/ethereum/go-ethereum/common"
)

type PriceLevel struct {
	Price    sdk.Dec
	Quantity sdk.Dec
}

type SpotLimitOrderDelta struct {
	Order        *SpotLimitOrder
	FillQuantity sdk.Dec
}

type DerivativeLimitOrderDelta struct {
	Order          *DerivativeLimitOrder
	FillQuantity   sdk.Dec
	CancelQuantity sdk.Dec
}

type DerivativeMarketOrderDelta struct {
	Order        *DerivativeMarketOrder
	FillQuantity sdk.Dec
}

func (d *DerivativeMarketOrderDelta) UnfilledQuantity() sdk.Dec {
	return d.Order.OrderInfo.Quantity.Sub(d.FillQuantity)
}

func (d *DerivativeLimitOrderDelta) IsBuy() bool {
	return d.Order.IsBuy()
}

func (d *DerivativeLimitOrderDelta) SubaccountID() common.Hash {
	return d.Order.SubaccountID()
}

func (d *DerivativeLimitOrderDelta) Price() sdk.Dec {
	return d.Order.Price()
}

func (d *DerivativeLimitOrderDelta) FillableQuantity() sdk.Dec {
	return d.Order.Fillable.Sub(d.CancelQuantity)
}

func (d *DerivativeLimitOrderDelta) OrderHash() common.Hash {
	return d.Order.Hash()
}

var AuctionSubaccountID = common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
var ZeroSubaccountID = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")

func IsValidSubaccountID(subaccountID string) (*common.Address, bool) {
	if len(subaccountID) != 66 {
		return nil, false
	}
	subaccountIdBytes := common.FromHex(subaccountID)

	if len(subaccountIdBytes) != common.HashLength {
		return nil, false
	}
	addressBytes := subaccountIdBytes[:common.AddressLength]
	if !common.IsHexAddress(common.Bytes2Hex(addressBytes)) {
		return nil, false
	}
	address := common.BytesToAddress(addressBytes)
	return &address, true
}

func IsValidOrderHash(orderHash string) bool {
	if len(orderHash) != 66 {
		return false
	}
	orderHashBytes := common.FromHex(orderHash)

	if len(orderHashBytes) != common.HashLength {
		return false
	}
	return true
}

func BreachesMinimumTickSize(value sdk.Dec, minTickSize sdk.Dec) bool {
	// obviously breached if the value less than the minTickSize
	if value.LT(minTickSize) {
		return true
	}

	// prevent mod by 0
	if minTickSize.IsZero() {
		return true
	}

	residue := new(big.Int).Mod(value.BigInt(), minTickSize.BigInt())
	return !bytes.Equal(residue.Bytes(), big.NewInt(0).Bytes())
}

func (s *Subaccount) GetSubaccountID() (*common.Hash, error) {
	trader, err := sdk.AccAddressFromBech32(s.Trader)
	if err != nil {
		return nil, err
	}
	return SdkAddressWithNonceToSubaccountID(trader, s.SubaccountNonce)
}

func SdkAddressWithNonceToSubaccountID(addr sdk.Address, nonce uint32) (*common.Hash, error) {
	if len(addr.Bytes()) > common.AddressLength {
		return &AuctionSubaccountID, ErrBadSubaccountID
	}
	subaccountID := common.BytesToHash(append(addr.Bytes(), common.LeftPadBytes(big.NewInt(int64(nonce)).Bytes(), 12)...))
	return &subaccountID, nil
}

func SdkAddressToSubaccountID(addr sdk.Address) common.Hash {
	return common.BytesToHash(common.RightPadBytes(addr.Bytes(), 32))
}

func SdkAddressToEthAddress(addr sdk.Address) common.Address {
	return common.BytesToAddress(addr.Bytes())
}

func EthAddressToSubaccountID(addr common.Address) common.Hash {
	return common.BytesToHash(common.RightPadBytes(addr.Bytes(), 32))
}

// PrintSpotLimitOrderbookState is a helper debugger function to print a tabular view of the spot limit orderbook fill state
func PrintSpotLimitOrderbookState(buyOrderbookState *OrderbookFills, sellOrderbookState *OrderbookFills) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Buy Price", "Buy Quantity", "Buy Fill Quantity", "Sell Price", "Sell Quantity", "Sell Fill Quantity"})
	maxLength := 0
	if buyOrderbookState != nil {
		maxLength = len(buyOrderbookState.Orders)
	}
	if sellOrderbookState != nil {
		if len(sellOrderbookState.Orders) > maxLength {
			maxLength = len(sellOrderbookState.Orders)
		}
	}
	precision := 6

	for idx := 0; idx < maxLength; idx++ {
		row := make([]string, 0)
		if buyOrderbookState == nil || idx >= len(buyOrderbookState.Orders) {
			row = append(row, "-", "-", "-")
		} else {
			buyOrder := buyOrderbookState.Orders[idx]
			fillQuantity := buyOrderbookState.FillQuantities[idx]
			row = append(row, buyOrder.OrderInfo.Price.String()[:precision])
			row = append(row, buyOrder.Fillable.String()[:precision])
			row = append(row, fillQuantity.String()[:precision])
		}
		if sellOrderbookState == nil || idx >= len(sellOrderbookState.Orders) {
			row = append(row, "-", "-", "-")
		} else {
			sellOrder := sellOrderbookState.Orders[idx]
			fillQuantity := sellOrderbookState.FillQuantities[idx]
			row = append(row, sellOrder.OrderInfo.Price.String()[:precision])
			row = append(row, sellOrder.Fillable.String()[:precision])
			row = append(row, fillQuantity.String()[:precision])
		}

		table.Append(row)
	}
	table.Render()
}

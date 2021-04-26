package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/olekukonko/tablewriter"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type PriceLevel struct {
	Price    sdk.Dec
	Quantity sdk.Dec
}

type LimitOrderFilledDelta struct {
	SubaccountIndexKey []byte
	FillableAmount     sdk.Dec
}

var AuctionSubaccountID = common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")

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

func checkIfExceedDecimals(dec sdk.Dec, maxDecimals uint32) bool {
	powered := dec.Mul(sdk.NewDec(10).Power(uint64(maxDecimals)))
	return !powered.Equal(powered.Ceil())
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

package types

import (
	"os"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/olekukonko/tablewriter"
)

func NewOrderbook(marketID common.Hash) *Orderbook {
	return &Orderbook{
		MarketId:   marketID.Bytes(),
		BuyLevels:  make([]*Level, 0),
		SellLevels: make([]*Level, 0),
	}
}

func (o *Orderbook) AppendLevel(isBuy bool, level *Level) {
	if isBuy {
		o.BuyLevels = append(o.BuyLevels, level)
		return
	}
	o.SellLevels = append(o.SellLevels, level)
}

func (o *Orderbook) IsCrossed() bool {
	if len(o.BuyLevels) == 0 || len(o.SellLevels) == 0 {
		return false
	}

	highestBuyLevel := o.BuyLevels[len(o.BuyLevels)-1]

	if highestBuyLevel.Q.IsZero() {
		return false
	}

	lowestSellPrice := sdk.ZeroDec()

	isQuantityAllZero := true

	for _, level := range o.SellLevels {
		if level.Q.IsZero() {
			continue
		}
		lowestSellPrice = level.P
		isQuantityAllZero = false
		break
	}

	if isQuantityAllZero {
		return false
	}

	return highestBuyLevel.P.GTE(lowestSellPrice)
}

// PrintDisplay is a helper function to print the orderbook in a human readable format for debugging purposes
func (o *Orderbook) PrintDisplay() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Buy Price", "Buy Quantity", "Sell Price", "Sell Quantity"})

	maxLength := len(o.BuyLevels)
	if len(o.SellLevels) > maxLength {
		maxLength = len(o.SellLevels)
	}

	for idx := 0; idx < maxLength; idx++ {
		row := make([]string, 0)
		if idx < len(o.BuyLevels) {
			row = append(row, getReadableDec(o.BuyLevels[idx].P))
			row = append(row, getReadableDec(o.BuyLevels[idx].Q))
		} else {
			row = append(row, "-", "-")
		}
		if idx < len(o.SellLevels) {
			row = append(row, getReadableDec(o.SellLevels[idx].P))
			row = append(row, getReadableDec(o.SellLevels[idx].Q))
		} else {
			row = append(row, "-", "-")
		}
		table.Append(row)
	}
	table.Render()
}

// getReadableDec is a test utility function to return a readable representation of decimal strings
func getReadableDec(d sdk.Dec) string {
	if d.IsNil() {
		return d.String()
	}
	dec := strings.TrimRight(d.String(), "0")
	if len(dec) < 2 {
		return dec
	}

	if dec[len(dec)-1:] == "." {
		return dec + "0"
	}
	return dec
}

func NewLevel(price sdk.Dec, quantity sdk.Dec) *Level {
	return &Level{
		P: price,
		Q: quantity,
	}
}

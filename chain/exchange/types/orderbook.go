package types

import (
	"os"
	"sort"
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

func NewOrderbookWithLevels(marketID common.Hash, buyLevels, sellLevels []*Level) *Orderbook {
	return &Orderbook{
		MarketId:   marketID.Bytes(),
		BuyLevels:  buyLevels,
		SellLevels: sellLevels,
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
	PrintDisplayLevels(o.BuyLevels, o.SellLevels)
}

// PrintDisplayLevels is a helper function to print the orderbook in a human readable format for debugging purposes
func PrintDisplayLevels(buyLevels, sellLevels []*Level) {
	buyLevelsSorted := make([]*Level, len(buyLevels))
	copy(buyLevelsSorted, buyLevels)
	sellLevelsSorted := make([]*Level, len(sellLevels))
	copy(sellLevelsSorted, sellLevels)
	if len(buyLevels) > 1 {
		sort.SliceStable(buyLevelsSorted, func(i, j int) bool {
			return buyLevelsSorted[i].GetPrice().GT(buyLevelsSorted[j].GetPrice())
		})
	}
	if len(sellLevelsSorted) > 1 {
		sort.SliceStable(sellLevelsSorted, func(i, j int) bool {
			return sellLevelsSorted[i].GetPrice().LT(sellLevelsSorted[j].GetPrice())
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Buy Price", "Buy Quantity", "Sell Price", "Sell Quantity"})

	maxLength := len(buyLevelsSorted)
	if len(sellLevelsSorted) > maxLength {
		maxLength = len(sellLevelsSorted)
	}

	for idx := 0; idx < maxLength; idx++ {
		row := make([]string, 0)
		if idx < len(buyLevelsSorted) {
			row = append(row, getReadableDec(buyLevelsSorted[idx].P), getReadableDec(buyLevelsSorted[idx].Q))
		} else {
			row = append(row, "-", "-")
		}
		if idx < len(sellLevelsSorted) {
			row = append(row, getReadableDec(sellLevelsSorted[idx].P), getReadableDec(sellLevelsSorted[idx].Q))
		} else {
			row = append(row, "-", "-")
		}
		table.Append(row)
	}
	table.Render()
}

func (o *Orderbook) Equals(other *Orderbook) bool {
	if len(o.BuyLevels) != len(other.BuyLevels) || len(o.SellLevels) != len(other.SellLevels) {
		return false
	}
	for i, level := range o.BuyLevels {
		metaLevel := other.BuyLevels[i]
		if !level.Q.Equal(metaLevel.Q) || !level.P.Equal(metaLevel.P) {
			return false
		}
	}
	return true
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

func NewLevel(price, quantity sdk.Dec) *Level {
	return &Level{
		P: price,
		Q: quantity,
	}
}

func (l *Level) GetPrice() sdk.Dec {
	return l.P
}

func (l *Level) GetQuantity() sdk.Dec {
	return l.Q
}

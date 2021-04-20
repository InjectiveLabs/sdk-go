package types

import (
	"bytes"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type VwapData struct {
	Price    sdk.Dec
	Quantity sdk.Dec
}

func NewVwapData() *VwapData {
	return &VwapData{
		Price:    sdk.ZeroDec(),
		Quantity: sdk.ZeroDec(),
	}
}

func (p *VwapData) ApplyExecution(price, quantity sdk.Dec) *VwapData {
	if price.IsNil() || quantity.IsNil() || quantity.IsZero() {
		return NewVwapData()
	}

	if p == nil {
		return &VwapData{price, quantity}
	}

	newQuantity := p.Quantity.Add(quantity)
	newPrice := p.Price.Mul(p.Quantity).Add(price.Mul(quantity)).Quo(newQuantity)

	return &VwapData{
		Price:    newPrice,
		Quantity: newQuantity,
	}
}

type VwapInfo struct {
	MarkPrice sdk.Dec
	VwapData  *VwapData
}

func NewVwapInfo(markPrice sdk.Dec) *VwapInfo {
	return &VwapInfo{
		MarkPrice: markPrice,
		VwapData:  NewVwapData(),
	}
}

type PerpetualVwapInfo map[common.Hash]*VwapInfo

func NewPerpetualVwapInfo() PerpetualVwapInfo {
	return make(PerpetualVwapInfo)
}

func (p *PerpetualVwapInfo) ApplyVwap(marketID common.Hash, markPrice sdk.Dec, vwapData *VwapData) {
	vwapInfo := (*p)[marketID]
	if vwapInfo == nil {
		vwapInfo = NewVwapInfo(markPrice)
		(*p)[marketID] = vwapInfo
	}
	if !vwapData.Quantity.IsZero() {
		vwapInfo.VwapData = vwapInfo.VwapData.ApplyExecution(vwapData.Price, vwapData.Quantity)
	}
}

func (p *PerpetualVwapInfo) GetSortedMarketIDs() []common.Hash {
	marketIDs := make([]common.Hash, 0)
	for k := range *p {
		marketIDs = append(marketIDs, k)
	}
	sort.SliceStable(marketIDs, func(i, j int) bool {
		return bytes.Compare(marketIDs[i].Bytes(), marketIDs[j].Bytes()) < 0
	})
	return marketIDs
}

// ComputeSyntheticVwapUnitDelta returns (price - markPrice) / markPrice
func (p *PerpetualVwapInfo) ComputeSyntheticVwapUnitDelta(marketID common.Hash) sdk.Dec {
	vwapInfo := (*p)[marketID]
	return vwapInfo.VwapData.Price.Sub(vwapInfo.MarkPrice).Quo(vwapInfo.MarkPrice)
}

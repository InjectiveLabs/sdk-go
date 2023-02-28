package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func NewPythPriceState(
	priceID common.Hash,
	emaPrice, emaConf, conf sdk.Dec,
	publishTime int64,
	price sdk.Dec,
	blockTime int64,
) *PythPriceState {
	return &PythPriceState{
		PriceId:     priceID.Hex(),
		EmaPrice:    emaPrice,
		EmaConf:     emaConf,
		Conf:        conf,
		PublishTime: uint64(publishTime),
		PriceState:  *NewPriceState(price, blockTime),
	}
}

func (p *PythPriceState) Update(
	emaPrice, emaConf, conf sdk.Dec,
	publishTime uint64,
	price sdk.Dec,
	blockTime int64,
) {
	p.EmaPrice = emaPrice
	p.EmaConf = emaConf
	p.Conf = conf
	p.PublishTime = publishTime
	p.PriceState.UpdatePrice(price, blockTime)
}

func (p *PriceAttestation) GetPriceID() common.Hash {
	return common.BytesToHash(p.PriceId)
}

func (p *PriceAttestation) Validate() error {
	if len(p.PriceId) != 32 {
		return ErrInvalidPythPriceID
	}

	if !(p.Price > 0) {
		return ErrBadPrice
	}

	if p.Expo > MaxPythExponent || p.Expo < MinPythExponent {
		return ErrInvalidPythExponent
	}

	if !(p.PublishTime > 0) {
		return ErrInvalidPythPublishTime
	}

	// validating EMA price/conf is omitted since it's not mission critical and we don't want to block normal price updates
	return nil
}

func GetExponentiatedDec(value, expo int64) sdk.Dec {
	// price * 10^-expo
	if expo <= 0 {
		return sdk.NewDecWithPrec(value, -expo)
	}

	// price * 10^expo
	return sdk.NewDec(value).Power(uint64(expo))
}

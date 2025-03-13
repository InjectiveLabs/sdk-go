package v2

import (
	"cosmossdk.io/math"
	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type SpotLimitOrderDelta struct {
	Order        *SpotLimitOrder
	FillQuantity math.LegacyDec
}

type DerivativeLimitOrderDelta struct {
	Order          *DerivativeLimitOrder
	FillQuantity   math.LegacyDec
	CancelQuantity math.LegacyDec
}

type DerivativeMarketOrderDelta struct {
	Order        *DerivativeMarketOrder
	FillQuantity math.LegacyDec
}

func (d *DerivativeMarketOrderDelta) UnfilledQuantity() math.LegacyDec {
	return d.Order.OrderInfo.Quantity.Sub(d.FillQuantity)
}

func (d *DerivativeLimitOrderDelta) IsBuy() bool {
	return d.Order.IsBuy()
}

func (d *DerivativeLimitOrderDelta) SubaccountID() common.Hash {
	return d.Order.SubaccountID()
}

func (d *DerivativeLimitOrderDelta) Price() math.LegacyDec {
	return d.Order.Price()
}

func (d *DerivativeLimitOrderDelta) FillableQuantity() math.LegacyDec {
	return d.Order.Fillable.Sub(d.CancelQuantity)
}

func (d *DerivativeLimitOrderDelta) OrderHash() common.Hash {
	return d.Order.Hash()
}

func (d *DerivativeLimitOrderDelta) Cid() string {
	return d.Order.Cid()
}

func (s *Subaccount) GetSubaccountID() (*common.Hash, error) {
	trader, err := sdk.AccAddressFromBech32(s.Trader)
	if err != nil {
		return nil, err
	}
	return types.SdkAddressWithNonceToSubaccountID(trader, s.SubaccountNonce)
}

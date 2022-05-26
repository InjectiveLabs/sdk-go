package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

var (
	_ authz.Authorization = &CreateDerivativeLimitOrderAuthz{}
	_ authz.Authorization = &CreateDerivativeMarketOrderAuthz{}
	_ authz.Authorization = &BatchCreateDerivativeLimitOrdersAuthz{}
	_ authz.Authorization = &CancelDerivativeOrderAuthz{}
	_ authz.Authorization = &BatchCancelDerivativeOrdersAuthz{}
)

// CreateDerivativeLimitOrderAuthz impl
func (a CreateDerivativeLimitOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateDerivativeLimitOrder{})
}

func (a CreateDerivativeLimitOrderAuthz) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	derivativeOrder, ok := msg.(*MsgCreateDerivativeLimitOrder)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}
	// check authorized subaccount
	if derivativeOrder.Order.OrderInfo.SubaccountId != a.SubaccountId {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
	}
	// check authorized market
	if !find(a.MarketIds, derivativeOrder.Order.MarketId) {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested market is unauthorized")
	}
	return authz.AcceptResponse{Accept: true, Delete: false, Updated: nil}, nil
}

func (a CreateDerivativeLimitOrderAuthz) ValidateBasic() error {
	if !IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

// CreateDerivativeMarketOrderAuthz impl
func (a CreateDerivativeMarketOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateDerivativeMarketOrder{})
}

func (a CreateDerivativeMarketOrderAuthz) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	derivativeOrder, ok := msg.(*MsgCreateDerivativeMarketOrder)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}
	// check authorized subaccount
	if derivativeOrder.Order.OrderInfo.SubaccountId != a.SubaccountId {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
	}
	// check authorized market
	if !find(a.MarketIds, derivativeOrder.Order.MarketId) {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested market is unauthorized")
	}
	return authz.AcceptResponse{Accept: true, Delete: false, Updated: nil}, nil
}

func (a CreateDerivativeMarketOrderAuthz) ValidateBasic() error {
	if !IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

// BatchCreateDerivativeLimitOrdersAuthz impl
func (a BatchCreateDerivativeLimitOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchCreateDerivativeLimitOrders{})
}

func (a BatchCreateDerivativeLimitOrdersAuthz) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	derivativeOrders, ok := msg.(*MsgBatchCreateDerivativeLimitOrders)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}
	for _, o := range derivativeOrders.Orders {
		// check authorized subaccount
		if o.OrderInfo.SubaccountId != a.SubaccountId {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
		// check authorized markets
		if !find(a.MarketIds, o.MarketId) {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested market is unauthorized")
		}
	}
	return authz.AcceptResponse{Accept: true, Delete: false, Updated: nil}, nil
}

func (a BatchCreateDerivativeLimitOrdersAuthz) ValidateBasic() error {
	if !IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

// CancelDerivativeOrderAuthz impl
func (a CancelDerivativeOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCancelDerivativeOrder{})
}

func (a CancelDerivativeOrderAuthz) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	orderToCancel, ok := msg.(*MsgCancelDerivativeOrder)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}
	// check authorized subaccount
	if orderToCancel.SubaccountId != a.SubaccountId {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
	}
	// check authorized market
	if !find(a.MarketIds, orderToCancel.MarketId) {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested market is unauthorized")
	}
	return authz.AcceptResponse{Accept: true, Delete: false, Updated: nil}, nil
}

func (a CancelDerivativeOrderAuthz) ValidateBasic() error {
	if !IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

// BatchCancelDerivativeOrdersAuthz impl
func (a BatchCancelDerivativeOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchCancelDerivativeOrders{})
}

func (a BatchCancelDerivativeOrdersAuthz) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	ordersToCancel, ok := msg.(*MsgBatchCancelDerivativeOrders)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}
	for _, o := range ordersToCancel.Data {
		// check authorized subaccount
		if o.SubaccountId != a.SubaccountId {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
		// check authorized markets
		if !find(a.MarketIds, o.MarketId) {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested market is unauthorized")
		}
	}
	return authz.AcceptResponse{Accept: true, Delete: false, Updated: nil}, nil
}

func (a BatchCancelDerivativeOrdersAuthz) ValidateBasic() error {
	if !IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

var (
	_ authz.Authorization = &CreateSpotLimitOrderAuthz{}
	_ authz.Authorization = &CreateSpotMarketOrderAuthz{}
	_ authz.Authorization = &BatchCreateSpotLimitOrdersAuthz{}
	_ authz.Authorization = &CancelSpotOrderAuthz{}
	_ authz.Authorization = &BatchCancelSpotOrdersAuthz{}
)

// CreateSpotLimitOrderAuthz impl
func (a CreateSpotLimitOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateSpotLimitOrder{})
}

func (a CreateSpotLimitOrderAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	spotOrder, ok := msg.(*MsgCreateSpotLimitOrder)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}
	// check authorized subaccount
	if spotOrder.Order.OrderInfo.SubaccountId != a.SubaccountId {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
	}
	// check authorized market
	if !find(a.MarketIds, spotOrder.Order.MarketId) {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested market is unauthorized")
	}
	return authz.AcceptResponse{Accept: true, Delete: false, Updated: nil}, nil
}

func (a CreateSpotLimitOrderAuthz) ValidateBasic() error {
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

// CreateSpotMarketOrderAuthz impl
func (a CreateSpotMarketOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateSpotMarketOrder{})
}

func (a CreateSpotMarketOrderAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	spotOrder, ok := msg.(*MsgCreateSpotMarketOrder)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}
	// check authorized subaccount
	if spotOrder.Order.OrderInfo.SubaccountId != a.SubaccountId {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
	}
	// check authorized market
	if !find(a.MarketIds, spotOrder.Order.MarketId) {
		return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested market is unauthorized")
	}
	return authz.AcceptResponse{Accept: true, Delete: false, Updated: nil}, nil
}

func (a CreateSpotMarketOrderAuthz) ValidateBasic() error {
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

// BatchCreateSpotLimitOrdersAuthz impl
func (a BatchCreateSpotLimitOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchCreateSpotLimitOrders{})
}

func (a BatchCreateSpotLimitOrdersAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	spotOrders, ok := msg.(*MsgBatchCreateSpotLimitOrders)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}
	for _, o := range spotOrders.Orders {
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

func (a BatchCreateSpotLimitOrdersAuthz) ValidateBasic() error {
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

// CancelSpotOrderAuthz impl
func (a CancelSpotOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCancelSpotOrder{})
}

func (a CancelSpotOrderAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	orderToCancel, ok := msg.(*MsgCancelSpotOrder)
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

func (a CancelSpotOrderAuthz) ValidateBasic() error {
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

// BatchCancelSpotOrdersAuthz impl
func (a BatchCancelSpotOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchCancelSpotOrders{})
}

func (a BatchCancelSpotOrdersAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	ordersToCancel, ok := msg.(*MsgBatchCancelSpotOrders)
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

func (a BatchCancelSpotOrdersAuthz) ValidateBasic() error {
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

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func reduceToSet(slice []string) []string {
	set := map[string]bool{}
	for _, s := range slice {
		set[s] = true
	}
	output := make([]string, len(set))
	i := 0
	for k := range set {
		output[i] = k
		i += 1
	}
	return output
}

var (
	_                      authz.Authorization = &BatchUpdateOrdersAuthz{}
	AuthorizedMarketsLimit                     = 200
)

// BatchUpdateOrdersAuthz impl
func (a BatchUpdateOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchUpdateOrders{})
}

func (a BatchUpdateOrdersAuthz) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	ordersToUpdate, ok := msg.(*MsgBatchUpdateOrders)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	// check authorized spot markets
	for _, o := range ordersToUpdate.SpotOrdersToCreate {
		if !find(a.SpotMarkets, o.MarketId) {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested spot market to create orders is unauthorized")
		}
		if o.OrderInfo.SubaccountId != a.SubaccountId {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
	}
	for _, o := range ordersToUpdate.SpotOrdersToCancel {
		if !find(a.SpotMarkets, o.MarketId) {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested spot market to cancel orders is unauthorized")
		}
		if o.SubaccountId != a.SubaccountId {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
	}
	for _, id := range ordersToUpdate.SpotMarketIdsToCancelAll {
		if !find(a.SpotMarkets, id) {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested spot market to cancel all orders is unauthorized")
		}

		if ordersToUpdate.SubaccountId != a.SubaccountId {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
	}

	// check authorized derivative markets
	for _, o := range ordersToUpdate.DerivativeOrdersToCreate {
		if !find(a.DerivativeMarkets, o.MarketId) {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested derivative market to create orders is unauthorized")
		}
		if o.OrderInfo.SubaccountId != a.SubaccountId {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
	}
	for _, o := range ordersToUpdate.DerivativeOrdersToCancel {
		if !find(a.DerivativeMarkets, o.MarketId) {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested derivative market to cancel orders is unauthorized")
		}
		if o.SubaccountId != a.SubaccountId {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
	}
	for _, id := range ordersToUpdate.DerivativeMarketIdsToCancelAll {
		if !find(a.DerivativeMarkets, id) {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested derivative market to cancel all orders is unauthorized")
		}

		if ordersToUpdate.SubaccountId != a.SubaccountId {
			return authz.AcceptResponse{}, sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
	}

	// TODO add check for BO markets?

	return authz.AcceptResponse{Accept: true, Delete: false, Updated: nil}, nil
}

func (a BatchUpdateOrdersAuthz) ValidateBasic() error {
	if !IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.SpotMarkets) == 0 && len(a.DerivativeMarkets) == 0 {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	if len(a.SpotMarkets) > AuthorizedMarketsLimit || len(a.DerivativeMarkets) > AuthorizedMarketsLimit {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	spotMarketsSet := reduceToSet(a.SpotMarkets)
	derivativeMarketsSet := reduceToSet(a.DerivativeMarkets)
	if len(a.SpotMarkets) != len(spotMarketsSet) || len(a.DerivativeMarkets) != len(derivativeMarketsSet) {
		return sdkerrors.ErrLogic.Wrapf("cannot have duplicate markets")
	}
	for _, m := range a.SpotMarkets {
		if !IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid spot market id to authorize")
		}
	}
	for _, m := range a.DerivativeMarkets {
		if !IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid derivative market id to authorize")
		}
	}
	return nil
}

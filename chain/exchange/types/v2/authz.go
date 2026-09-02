package v2

import (
	"context"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

var (
	authorizedMarketsLimit    = 200
	authorizedMarketsLimitMux = new(sync.RWMutex)

	_ authz.Authorization = &CreateSpotLimitOrderAuthz{}
	_ authz.Authorization = &CreateSpotMarketOrderAuthz{}
	_ authz.Authorization = &BatchCreateSpotLimitOrdersAuthz{}
	_ authz.Authorization = &CancelSpotOrderAuthz{}
	_ authz.Authorization = &BatchCancelSpotOrdersAuthz{}

	_ authz.Authorization = &CreateDerivativeLimitOrderAuthz{}
	_ authz.Authorization = &CreateDerivativeMarketOrderAuthz{}
	_ authz.Authorization = &BatchCreateDerivativeLimitOrdersAuthz{}
	_ authz.Authorization = &CancelDerivativeOrderAuthz{}
	_ authz.Authorization = &BatchCancelDerivativeOrdersAuthz{}

	_ authz.Authorization = &BatchUpdateOrdersAuthz{}
)

// AuthorizedMarketsLimit returns the authorized markets limit.
func AuthorizedMarketsLimit() int {
	authorizedMarketsLimitMux.RLock()
	defer authorizedMarketsLimitMux.RUnlock()
	return authorizedMarketsLimit
}

// SetAuthorizedMarketsLimit sets the authorized markets limit.
func SetAuthorizedMarketsLimit(limit int) {
	authorizedMarketsLimitMux.Lock()
	authorizedMarketsLimit = limit
	authorizedMarketsLimitMux.Unlock()
}

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
		i++
	}
	return output
}

type batchUpdateOrdersAuthzCheck struct {
	field string
	check func(*BatchUpdateOrdersAuthz, *MsgBatchUpdateOrders) error
}

var batchUpdateOrdersAuthzChecks = []batchUpdateOrdersAuthzCheck{
	{field: "SpotOrdersToCreate", check: (*BatchUpdateOrdersAuthz).authorizeSpotOrdersToCreate},
	{field: "SpotOrdersToCancel", check: (*BatchUpdateOrdersAuthz).authorizeSpotOrdersToCancel},
	{field: "SpotMarketIdsToCancelAll", check: (*BatchUpdateOrdersAuthz).authorizeSpotMarketIDsToCancelAll},
	{field: "SpotMarketOrdersToCreate", check: (*BatchUpdateOrdersAuthz).authorizeSpotMarketOrdersToCreate},
	{field: "DerivativeOrdersToCreate", check: (*BatchUpdateOrdersAuthz).authorizeDerivativeOrdersToCreate},
	{field: "DerivativeOrdersToCancel", check: (*BatchUpdateOrdersAuthz).authorizeDerivativeOrdersToCancel},
	{field: "DerivativeMarketIdsToCancelAll", check: (*BatchUpdateOrdersAuthz).authorizeDerivativeMarketIDsToCancelAll},
	{field: "DerivativeMarketOrdersToCreate", check: (*BatchUpdateOrdersAuthz).authorizeDerivativeMarketOrdersToCreate},
	{field: "BinaryOptionsOrdersToCreate", check: (*BatchUpdateOrdersAuthz).authorizeBinaryOptionsOrdersToCreate},
	{field: "BinaryOptionsOrdersToCancel", check: (*BatchUpdateOrdersAuthz).authorizeBinaryOptionsOrdersToCancel},
	{field: "BinaryOptionsMarketIdsToCancelAll", check: (*BatchUpdateOrdersAuthz).authorizeBinaryOptionsMarketIDsToCancelAll},
	{field: "BinaryOptionsMarketOrdersToCreate", check: (*BatchUpdateOrdersAuthz).authorizeBinaryOptionsMarketOrdersToCreate},
}

func (a *BatchUpdateOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchUpdateOrders{})
}

func (a *BatchUpdateOrdersAuthz) Accept(_ context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	ordersToUpdate, ok := msg.(*MsgBatchUpdateOrders)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	for _, authzCheck := range batchUpdateOrdersAuthzChecks {
		if err := authzCheck.check(a, ordersToUpdate); err != nil {
			return authz.AcceptResponse{}, err
		}
	}

	return authz.AcceptResponse{Accept: true, Delete: false, Updated: nil}, nil
}

func (a *BatchUpdateOrdersAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.SpotMarkets) == 0 && len(a.DerivativeMarkets) == 0 {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	if err := types.ValidateAuthorizedMarkets(a.SpotMarkets, "invalid spot market id to authorize", AuthorizedMarketsLimit()); err != nil {
		return err
	}

	return types.ValidateAuthorizedMarkets(a.DerivativeMarkets, "invalid derivative market id to authorize", AuthorizedMarketsLimit())
}

func (a *BatchUpdateOrdersAuthz) authorizeSpotOrdersToCreate(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeOrders(
		a.SpotMarkets,
		a.SubaccountId,
		ordersToUpdate.SpotOrdersToCreate,
		func(order *SpotOrder) string { return order.MarketId },
		func(order *SpotOrder) string { return order.OrderInfo.SubaccountId },
		"requested spot market to create orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeSpotOrdersToCancel(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeOrders(
		a.SpotMarkets,
		a.SubaccountId,
		ordersToUpdate.SpotOrdersToCancel,
		func(order *OrderData) string { return order.MarketId },
		func(order *OrderData) string { return order.SubaccountId },
		"requested spot market to cancel orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeSpotMarketIDsToCancelAll(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeCancelAll(
		a.SpotMarkets,
		a.SubaccountId,
		ordersToUpdate.SubaccountId,
		ordersToUpdate.SpotMarketIdsToCancelAll,
		"requested spot market to cancel all orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeSpotMarketOrdersToCreate(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeOrders(
		a.SpotMarkets,
		a.SubaccountId,
		ordersToUpdate.SpotMarketOrdersToCreate,
		func(order *SpotOrder) string { return order.MarketId },
		func(order *SpotOrder) string { return order.OrderInfo.SubaccountId },
		"requested spot market to create orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeDerivativeOrdersToCreate(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeOrders(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.DerivativeOrdersToCreate,
		func(order *DerivativeOrder) string { return order.MarketId },
		func(order *DerivativeOrder) string { return order.OrderInfo.SubaccountId },
		"requested derivative market to create orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeDerivativeOrdersToCancel(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeOrders(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.DerivativeOrdersToCancel,
		func(order *OrderData) string { return order.MarketId },
		func(order *OrderData) string { return order.SubaccountId },
		"requested derivative market to cancel orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeDerivativeMarketIDsToCancelAll(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeCancelAll(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.SubaccountId,
		ordersToUpdate.DerivativeMarketIdsToCancelAll,
		"requested derivative market to cancel all orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeDerivativeMarketOrdersToCreate(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeOrders(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.DerivativeMarketOrdersToCreate,
		func(order *DerivativeOrder) string { return order.MarketId },
		func(order *DerivativeOrder) string { return order.OrderInfo.SubaccountId },
		"requested derivative market to create orders is unauthorized",
	)
}

// BatchUpdateOrdersAuthz only carries spot and derivative allowlists, so
// binary options share the derivative market authorization set.
func (a *BatchUpdateOrdersAuthz) authorizeBinaryOptionsOrdersToCreate(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeOrders(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.BinaryOptionsOrdersToCreate,
		func(order *DerivativeOrder) string { return order.MarketId },
		func(order *DerivativeOrder) string { return order.OrderInfo.SubaccountId },
		"requested binary options market to create orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeBinaryOptionsOrdersToCancel(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeOrders(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.BinaryOptionsOrdersToCancel,
		func(order *OrderData) string { return order.MarketId },
		func(order *OrderData) string { return order.SubaccountId },
		"requested binary options market to cancel orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeBinaryOptionsMarketIDsToCancelAll(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeCancelAll(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.SubaccountId,
		ordersToUpdate.BinaryOptionsMarketIdsToCancelAll,
		"requested binary options market to cancel all orders is unauthorized",
	)
}

func (a *BatchUpdateOrdersAuthz) authorizeBinaryOptionsMarketOrdersToCreate(ordersToUpdate *MsgBatchUpdateOrders) error {
	return types.AuthorizeOrders(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.BinaryOptionsMarketOrdersToCreate,
		func(order *DerivativeOrder) string { return order.MarketId },
		func(order *DerivativeOrder) string { return order.OrderInfo.SubaccountId },
		"requested binary options market to create orders is unauthorized",
	)
}

func (a *CreateDerivativeLimitOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateDerivativeLimitOrder{})
}

func (a *CreateDerivativeLimitOrderAuthz) Accept(_ context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *CreateDerivativeLimitOrderAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

func (a *CreateDerivativeMarketOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateDerivativeMarketOrder{})
}

func (a *CreateDerivativeMarketOrderAuthz) Accept(_ context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *CreateDerivativeMarketOrderAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

func (a *BatchCreateDerivativeLimitOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchCreateDerivativeLimitOrders{})
}

func (a *BatchCreateDerivativeLimitOrdersAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *BatchCreateDerivativeLimitOrdersAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

func (a *CancelDerivativeOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCancelDerivativeOrder{})
}

func (a *CancelDerivativeOrderAuthz) Accept(_ context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *CancelDerivativeOrderAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

func (a *BatchCancelDerivativeOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchCancelDerivativeOrders{})
}

func (a *BatchCancelDerivativeOrdersAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *BatchCancelDerivativeOrdersAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

func (a *CreateSpotLimitOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateSpotLimitOrder{})
}

func (a *CreateSpotLimitOrderAuthz) Accept(_ context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *CreateSpotLimitOrderAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

func (a *CreateSpotMarketOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateSpotMarketOrder{})
}

func (a *CreateSpotMarketOrderAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *CreateSpotMarketOrderAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

func (a *BatchCreateSpotLimitOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchCreateSpotLimitOrders{})
}

func (a *BatchCreateSpotLimitOrdersAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *BatchCreateSpotLimitOrdersAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

func (a *CancelSpotOrderAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCancelSpotOrder{})
}

func (a *CancelSpotOrderAuthz) Accept(_ context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *CancelSpotOrderAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

func (a *BatchCancelSpotOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchCancelSpotOrders{})
}

func (a *BatchCancelSpotOrdersAuthz) Accept(_ context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a *BatchCancelSpotOrdersAuthz) ValidateBasic() error {
	if !types.IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.MarketIds) == 0 || len(a.MarketIds) > AuthorizedMarketsLimit() {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	marketsSet := reduceToSet(a.MarketIds)
	if len(a.MarketIds) != len(marketsSet) {
		return sdkerrors.ErrLogic.Wrapf("Cannot have duplicate markets")
	}
	for _, m := range a.MarketIds {
		if !types.IsHexHash(m) {
			return sdkerrors.ErrLogic.Wrap("invalid market id to authorize")
		}
	}
	return nil
}

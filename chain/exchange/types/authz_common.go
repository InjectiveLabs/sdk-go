package types

import (
	"context"
	"sync"

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
		i++
	}
	return output
}

type batchUpdateOrdersAuthzCheck struct {
	field string
	check func(BatchUpdateOrdersAuthz, *MsgBatchUpdateOrders) error
}

var (
	_                            authz.Authorization = &BatchUpdateOrdersAuthz{}
	authorizedMarketsLimit                           = 200
	authorizedMarketsLimitMux                        = new(sync.RWMutex)
	batchUpdateOrdersAuthzChecks                     = []batchUpdateOrdersAuthzCheck{
		{field: "SpotOrdersToCreate", check: BatchUpdateOrdersAuthz.authorizeSpotOrdersToCreate},
		{field: "SpotOrdersToCancel", check: BatchUpdateOrdersAuthz.authorizeSpotOrdersToCancel},
		{field: "SpotMarketIdsToCancelAll", check: BatchUpdateOrdersAuthz.authorizeSpotMarketIDsToCancelAll},
		{field: "DerivativeOrdersToCreate", check: BatchUpdateOrdersAuthz.authorizeDerivativeOrdersToCreate},
		{field: "DerivativeOrdersToCancel", check: BatchUpdateOrdersAuthz.authorizeDerivativeOrdersToCancel},
		{field: "DerivativeMarketIdsToCancelAll", check: BatchUpdateOrdersAuthz.authorizeDerivativeMarketIDsToCancelAll},
		{field: "BinaryOptionsOrdersToCreate", check: BatchUpdateOrdersAuthz.authorizeBinaryOptionsOrdersToCreate},
		{field: "BinaryOptionsOrdersToCancel", check: BatchUpdateOrdersAuthz.authorizeBinaryOptionsOrdersToCancel},
		{field: "BinaryOptionsMarketIdsToCancelAll", check: BatchUpdateOrdersAuthz.authorizeBinaryOptionsMarketIDsToCancelAll},
	}
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

// BatchUpdateOrdersAuthz impl
func (a BatchUpdateOrdersAuthz) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgBatchUpdateOrders{})
}

func (a BatchUpdateOrdersAuthz) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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

func (a BatchUpdateOrdersAuthz) ValidateBasic() error {
	if !IsHexHash(a.SubaccountId) {
		return sdkerrors.ErrLogic.Wrap("invalid subaccount id to authorize")
	}
	if len(a.SpotMarkets) == 0 && len(a.DerivativeMarkets) == 0 {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	if err := ValidateAuthorizedMarkets(a.SpotMarkets, "invalid spot market id to authorize", AuthorizedMarketsLimit()); err != nil {
		return err
	}

	return ValidateAuthorizedMarkets(a.DerivativeMarkets, "invalid derivative market id to authorize", AuthorizedMarketsLimit())
}

// ValidateAuthorizedMarkets validates a batch-update authz market allowlist.
func ValidateAuthorizedMarkets(marketIDs []string, invalidMarketErr string, authorizedMarketsLimit int) error {
	if len(marketIDs) > authorizedMarketsLimit {
		return sdkerrors.ErrLogic.Wrapf("invalid markets array length")
	}
	if len(marketIDs) != len(reduceToSet(marketIDs)) {
		return sdkerrors.ErrLogic.Wrapf("cannot have duplicate markets")
	}
	for _, marketID := range marketIDs {
		if !IsHexHash(marketID) {
			return sdkerrors.ErrLogic.Wrap(invalidMarketErr)
		}
	}
	return nil
}

// AuthorizeOrders validates a batch of authz-protected orders against the allowed markets and subaccount.
func AuthorizeOrders[T any](
	authorizedMarkets []string,
	subaccountID string,
	orders []T,
	marketIDFn func(T) string,
	orderSubaccountIDFn func(T) string,
	unauthorizedMarketErr string,
) error {
	for _, order := range orders {
		if !find(authorizedMarkets, marketIDFn(order)) {
			return sdkerrors.ErrUnauthorized.Wrap(unauthorizedMarketErr)
		}
		if orderSubaccountIDFn(order) != subaccountID {
			return sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
	}

	return nil
}

// AuthorizeCancelAll validates cancel-all market requests against the allowed markets and subaccount.
func AuthorizeCancelAll(
	authorizedMarkets []string,
	authzSubaccountID, requestSubaccountID string,
	marketIDs []string,
	unauthorizedMarketErr string,
) error {
	for _, marketID := range marketIDs {
		if !find(authorizedMarkets, marketID) {
			return sdkerrors.ErrUnauthorized.Wrap(unauthorizedMarketErr)
		}
		if requestSubaccountID != authzSubaccountID {
			return sdkerrors.ErrUnauthorized.Wrapf("requested subaccount is unauthorized")
		}
	}

	return nil
}

func authorizeSpotOrders(authorizedMarkets []string, subaccountID string, orders []*SpotOrder, unauthorizedMarketErr string) error {
	return AuthorizeOrders(
		authorizedMarkets,
		subaccountID,
		orders,
		func(order *SpotOrder) string { return order.MarketId },
		func(order *SpotOrder) string { return order.OrderInfo.SubaccountId },
		unauthorizedMarketErr,
	)
}

func authorizeOrderCancellations(authorizedMarkets []string, subaccountID string, orders []*OrderData, unauthorizedMarketErr string) error {
	return AuthorizeOrders(
		authorizedMarkets,
		subaccountID,
		orders,
		func(order *OrderData) string { return order.MarketId },
		func(order *OrderData) string { return order.SubaccountId },
		unauthorizedMarketErr,
	)
}

func authorizeDerivativeOrders(authorizedMarkets []string, subaccountID string, orders []*DerivativeOrder, unauthorizedMarketErr string) error {
	return AuthorizeOrders(
		authorizedMarkets,
		subaccountID,
		orders,
		func(order *DerivativeOrder) string { return order.MarketId },
		func(order *DerivativeOrder) string { return order.OrderInfo.SubaccountId },
		unauthorizedMarketErr,
	)
}

func authorizeCancelAll(authorizedMarkets []string, authzSubaccountID, requestSubaccountID string, marketIDs []string, unauthorizedMarketErr string) error {
	return AuthorizeCancelAll(authorizedMarkets, authzSubaccountID, requestSubaccountID, marketIDs, unauthorizedMarketErr)
}

func (a BatchUpdateOrdersAuthz) authorizeSpotOrdersToCreate(ordersToUpdate *MsgBatchUpdateOrders) error {
	return authorizeSpotOrders(
		a.SpotMarkets,
		a.SubaccountId,
		ordersToUpdate.SpotOrdersToCreate,
		"requested spot market to create orders is unauthorized",
	)
}

func (a BatchUpdateOrdersAuthz) authorizeSpotOrdersToCancel(ordersToUpdate *MsgBatchUpdateOrders) error {
	return authorizeOrderCancellations(
		a.SpotMarkets,
		a.SubaccountId,
		ordersToUpdate.SpotOrdersToCancel,
		"requested spot market to cancel orders is unauthorized",
	)
}

func (a BatchUpdateOrdersAuthz) authorizeSpotMarketIDsToCancelAll(ordersToUpdate *MsgBatchUpdateOrders) error {
	return authorizeCancelAll(
		a.SpotMarkets,
		a.SubaccountId,
		ordersToUpdate.SubaccountId,
		ordersToUpdate.SpotMarketIdsToCancelAll,
		"requested spot market to cancel all orders is unauthorized",
	)
}

func (a BatchUpdateOrdersAuthz) authorizeDerivativeOrdersToCreate(ordersToUpdate *MsgBatchUpdateOrders) error {
	return authorizeDerivativeOrders(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.DerivativeOrdersToCreate,
		"requested derivative market to create orders is unauthorized",
	)
}

func (a BatchUpdateOrdersAuthz) authorizeDerivativeOrdersToCancel(ordersToUpdate *MsgBatchUpdateOrders) error {
	return authorizeOrderCancellations(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.DerivativeOrdersToCancel,
		"requested derivative market to cancel orders is unauthorized",
	)
}

func (a BatchUpdateOrdersAuthz) authorizeDerivativeMarketIDsToCancelAll(ordersToUpdate *MsgBatchUpdateOrders) error {
	return authorizeCancelAll(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.SubaccountId,
		ordersToUpdate.DerivativeMarketIdsToCancelAll,
		"requested derivative market to cancel all orders is unauthorized",
	)
}

// BatchUpdateOrdersAuthz only carries spot and derivative allowlists, so
// binary options share the derivative market authorization set.
func (a BatchUpdateOrdersAuthz) authorizeBinaryOptionsOrdersToCreate(ordersToUpdate *MsgBatchUpdateOrders) error {
	return authorizeDerivativeOrders(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.BinaryOptionsOrdersToCreate,
		"requested binary options market to create orders is unauthorized",
	)
}

func (a BatchUpdateOrdersAuthz) authorizeBinaryOptionsOrdersToCancel(ordersToUpdate *MsgBatchUpdateOrders) error {
	return authorizeOrderCancellations(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.BinaryOptionsOrdersToCancel,
		"requested binary options market to cancel orders is unauthorized",
	)
}

func (a BatchUpdateOrdersAuthz) authorizeBinaryOptionsMarketIDsToCancelAll(ordersToUpdate *MsgBatchUpdateOrders) error {
	return authorizeCancelAll(
		a.DerivativeMarkets,
		a.SubaccountId,
		ordersToUpdate.SubaccountId,
		ordersToUpdate.BinaryOptionsMarketIdsToCancelAll,
		"requested binary options market to cancel all orders is unauthorized",
	)
}

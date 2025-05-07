package v2

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

var _ authz.Authorization = &GenericExchangeAuthorization{}

const (
	ContextKeyHold = "hold"
)

// MsgType is the type of exchange message that is being approved.
type MsgType uint8

// This enum matches the one defined in Exchange.sol
const (
	MsgTypeDeposit                          MsgType = 1
	MsgTypeWithdraw                         MsgType = 2
	MsgTypeSubaccountTransfer               MsgType = 3
	MsgTypeExternalTransfer                 MsgType = 4
	MsgTypeIncreasePositionMargin           MsgType = 5
	MsgTypeDecreasePositionMargin           MsgType = 6
	MsgTypeBatchUpdateOrders                MsgType = 7
	MsgTypeCreateDerivativeLimitOrder       MsgType = 8
	MsgTypeBatchCreateDerivativeLimitOrders MsgType = 9
	MsgTypeCreateDerivativeMarketOrder      MsgType = 10
	MsgTypeCancelDerivativeOrder            MsgType = 11
	MsgTypeBatchCancelDerivativeOrders      MsgType = 12
	MsgTypeCreateSpotLimitOrder             MsgType = 13
	MsgTypeBatchCreateSpotLimitOrders       MsgType = 14
	MsgTypeCreateSpotMarketOrder            MsgType = 15
	MsgTypeCancelSpotOrder                  MsgType = 16
	MsgTypeBatchCancelSpotOrders            MsgType = 17
	MsgTypeUnknown                          MsgType = 18
)

var MsgTypeURLs = map[MsgType]string{
	MsgTypeDeposit:                          "/injective.exchange.v2.MsgDeposit",
	MsgTypeWithdraw:                         "/injective.exchange.v2.MsgWithdraw",
	MsgTypeSubaccountTransfer:               "/injective.exchange.v2.MsgSubaccountTransfer",
	MsgTypeExternalTransfer:                 "/injective.exchange.v2.MsgExternalTransfer",
	MsgTypeIncreasePositionMargin:           "/injective.exchange.v2.MsgIncreasePositionMargin",
	MsgTypeDecreasePositionMargin:           "/injective.exchange.v2.MsgDecreasePositionMargin",
	MsgTypeBatchUpdateOrders:                "/injective.exchange.v2.MsgBatchUpdateOrders",
	MsgTypeCreateDerivativeLimitOrder:       "/injective.exchange.v2.MsgCreateDerivativeLimitOrder",
	MsgTypeBatchCreateDerivativeLimitOrders: "/injective.exchange.v2.MsgBatchCreateDerivativeLimitOrders",
	MsgTypeCreateDerivativeMarketOrder:      "/injective.exchange.v2.MsgCreateDerivativeMarketOrder",
	MsgTypeCancelDerivativeOrder:            "/injective.exchange.v2.MsgCancelDerivativeOrder",
	MsgTypeBatchCancelDerivativeOrders:      "/injective.exchange.v2.MsgBatchCancelDerivativeOrders",
	MsgTypeCreateSpotLimitOrder:             "/injective.exchange.v2.MsgCreateSpotLimitOrder",
	MsgTypeBatchCreateSpotLimitOrders:       "/injective.exchange.v2.MsgBatchCreateSpotLimitOrders",
	MsgTypeCreateSpotMarketOrder:            "/injective.exchange.v2.MsgCreateSpotMarketOrder",
	MsgTypeCancelSpotOrder:                  "/injective.exchange.v2.MsgCancelSpotOrder",
	MsgTypeBatchCancelSpotOrders:            "/injective.exchange.v2.MsgBatchCancelSpotOrders",
}

// allowedMessages is a map of all the allowed messages that can be approved
var allowedMessages = map[string]bool{
	MsgTypeDeposit.URL():                          true,
	MsgTypeWithdraw.URL():                         true,
	MsgTypeSubaccountTransfer.URL():               true,
	MsgTypeExternalTransfer.URL():                 true,
	MsgTypeIncreasePositionMargin.URL():           true,
	MsgTypeDecreasePositionMargin.URL():           true,
	MsgTypeBatchUpdateOrders.URL():                true,
	MsgTypeCreateDerivativeLimitOrder.URL():       true,
	MsgTypeCreateDerivativeMarketOrder.URL():      true,
	MsgTypeCancelDerivativeOrder.URL():            true,
	MsgTypeBatchCancelDerivativeOrders.URL():      true,
	MsgTypeBatchCreateDerivativeLimitOrders.URL(): true,
	MsgTypeCreateSpotLimitOrder.URL():             true,
	MsgTypeBatchCreateSpotLimitOrders.URL():       true,
	MsgTypeCreateSpotMarketOrder.URL():            true,
	MsgTypeCancelSpotOrder.URL():                  true,
	MsgTypeBatchCancelSpotOrders.URL():            true,
}

// URL returns the URL of the message type.
func (m MsgType) URL() string {
	url, found := MsgTypeURLs[m]
	if found {
		return url
	}
	return "UNKNOWN"
}

// NewGenericExchangeAuthorization creates a new GenericExchangeAuthorization object.
func NewGenericExchangeAuthorization(msgTypeURL string, spendLimit sdk.Coins) *GenericExchangeAuthorization {
	return &GenericExchangeAuthorization{
		Msg:        msgTypeURL,
		SpendLimit: spendLimit,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a GenericExchangeAuthorization) MsgTypeURL() string {
	return a.Msg
}

// Accept implements Authorization.Accept.
// It is not necessary to check the message type here, as it is already checked
// in the ValidateBasic method. If we got here, it means a grant was found for
// the incoming message type, and it would not have been created if the
// message type was not allowed.
func (a GenericExchangeAuthorization) Accept(ctx context.Context, _ sdk.Msg) (authz.AcceptResponse, error) {
	newSpendLimit := a.SpendLimit

	// SpendLimit is optional, so we only check it if it is set
	if a.SpendLimit != nil {
		hold, ok := getHold(ctx)
		if !ok {
			return authz.AcceptResponse{Accept: false}, nil
		}
		for _, coin := range hold {
			allowed := a.SpendLimit.AmountOf(coin.Denom)
			if allowed.LT(coin.Amount) {
				return authz.AcceptResponse{Accept: false}, nil
			}
			newSpendLimit = a.SpendLimit.Sub(sdk.NewCoin(coin.Denom, coin.Amount))
		}
	}

	updatedAuthorization := NewGenericExchangeAuthorization(
		a.Msg,
		newSpendLimit,
	)

	return authz.AcceptResponse{Accept: true, Updated: updatedAuthorization}, nil
}

// getHold returns the hold from the context. Hold is the amount of coins that
// would be deducted from the account if the message is executed. We expect the
// caller to specify the hold in the context.
func getHold(ctx context.Context) (sdk.Coins, bool) {
	holdVal := ctx.Value(ContextKeyHold)
	if holdVal == nil {
		return nil, false
	}
	hold, ok := holdVal.(sdk.Coins)
	return hold, ok
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a GenericExchangeAuthorization) ValidateBasic() error {
	if a.Msg == "" {
		return errors.New("msg type cannot be empty")
	}

	if _, ok := allowedMessages[a.Msg]; !ok {
		return errors.New("msg type not allowed")
	}

	return nil
}

package types

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

// URL returns the URL of the message type.
func (m MsgType) URL() string {
	switch m {
	case MsgTypeDeposit:
		return "/injective.exchange.v1beta1.MsgDeposit"
	case MsgTypeWithdraw:
		return "/injective.exchange.v1beta1.MsgWithdraw"
	case MsgTypeSubaccountTransfer:
		return "/injective.exchange.v1beta1.MsgSubaccountTransfer"
	case MsgTypeExternalTransfer:
		return "/injective.exchange.v1beta1.MsgExternalTransfer"
	case MsgTypeIncreasePositionMargin:
		return "/injective.exchange.v1beta1.MsgIncreasePositionMargin"
	case MsgTypeDecreasePositionMargin:
		return "/injective.exchange.v1beta1.MsgDecreasePositionMargin"
	case MsgTypeBatchUpdateOrders:
		return "/injective.exchange.v1beta1.MsgBatchUpdateOrders"
	case MsgTypeCreateDerivativeLimitOrder:
		return "/injective.exchange.v1beta1.MsgCreateDerivativeLimitOrder"
	case MsgTypeBatchCreateDerivativeLimitOrders:
		return "/injective.exchange.v1beta1.MsgBatchCreateDerivativeLimitOrders"
	case MsgTypeCreateDerivativeMarketOrder:
		return "/injective.exchange.v1beta1.MsgCreateDerivativeMarketOrder"
	case MsgTypeCancelDerivativeOrder:
		return "/injective.exchange.v1beta1.MsgCancelDerivativeOrder"
	case MsgTypeBatchCancelDerivativeOrders:
		return "/injective.exchange.v1beta1.MsgBatchCancelDerivativeOrders"
	case MsgTypeCreateSpotLimitOrder:
		return "/injective.exchange.v1beta1.MsgCreateSpotLimitOrder"
	case MsgTypeBatchCreateSpotLimitOrders:
		return "/injective.exchange.v1beta1.MsgBatchCreateSpotLimitOrders"
	case MsgTypeCreateSpotMarketOrder:
		return "/injective.exchange.v1beta1.MsgCreateSpotMarketOrder"
	case MsgTypeCancelSpotOrder:
		return "/injective.exchange.v1beta1.MsgCancelSpotOrder"
	case MsgTypeBatchCancelSpotOrders:
		return "/injective.exchange.v1beta1.MsgBatchCancelSpotOrders"
	}
	return "UNKNOWN"
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
func (a GenericExchangeAuthorization) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
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
			a.SpendLimit = a.SpendLimit.Sub(sdk.NewCoin(coin.Denom, coin.Amount))
		}
	}
	return authz.AcceptResponse{Accept: true, Updated: &a}, nil
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

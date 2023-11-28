package types

import (
	"bytes"
	"encoding/json"

	"cosmossdk.io/errors"
	sdksecp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/injective-core/injective-chain/crypto/ethsecp256k1"
	oracletypes "github.com/InjectiveLabs/injective-core/injective-chain/modules/oracle/types"
	wasmxtypes "github.com/InjectiveLabs/injective-core/injective-chain/modules/wasmx/types"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgDeposit{}
	_ sdk.Msg = &MsgWithdraw{}
	_ sdk.Msg = &MsgCreateSpotLimitOrder{}
	_ sdk.Msg = &MsgBatchCreateSpotLimitOrders{}
	_ sdk.Msg = &MsgCreateSpotMarketOrder{}
	_ sdk.Msg = &MsgCancelSpotOrder{}
	_ sdk.Msg = &MsgBatchCancelSpotOrders{}
	_ sdk.Msg = &MsgCreateDerivativeLimitOrder{}
	_ sdk.Msg = &MsgBatchCreateDerivativeLimitOrders{}
	_ sdk.Msg = &MsgCreateDerivativeMarketOrder{}
	_ sdk.Msg = &MsgCancelDerivativeOrder{}
	_ sdk.Msg = &MsgBatchCancelDerivativeOrders{}
	_ sdk.Msg = &MsgSubaccountTransfer{}
	_ sdk.Msg = &MsgExternalTransfer{}
	_ sdk.Msg = &MsgIncreasePositionMargin{}
	_ sdk.Msg = &MsgLiquidatePosition{}
	_ sdk.Msg = &MsgInstantSpotMarketLaunch{}
	_ sdk.Msg = &MsgInstantPerpetualMarketLaunch{}
	_ sdk.Msg = &MsgInstantExpiryFuturesMarketLaunch{}
	_ sdk.Msg = &MsgBatchUpdateOrders{}
	_ sdk.Msg = &MsgPrivilegedExecuteContract{}
	_ sdk.Msg = &MsgRewardsOptOut{}
	_ sdk.Msg = &MsgInstantBinaryOptionsMarketLaunch{}
	_ sdk.Msg = &MsgCreateBinaryOptionsLimitOrder{}
	_ sdk.Msg = &MsgCreateBinaryOptionsMarketOrder{}
	_ sdk.Msg = &MsgCancelBinaryOptionsOrder{}
	_ sdk.Msg = &MsgAdminUpdateBinaryOptionsMarket{}
	_ sdk.Msg = &MsgBatchCancelBinaryOptionsOrders{}
	_ sdk.Msg = &MsgReclaimLockedFunds{}
	_ sdk.Msg = &MsgUpdateParams{}
)

// exchange message types
const (
	TypeMsgDeposit                          = "msgDeposit"
	TypeMsgWithdraw                         = "msgWithdraw"
	TypeMsgCreateSpotLimitOrder             = "createSpotLimitOrder"
	TypeMsgBatchCreateSpotLimitOrders       = "batchCreateSpotLimitOrders"
	TypeMsgCreateSpotMarketOrder            = "createSpotMarketOrder"
	TypeMsgCancelSpotOrder                  = "cancelSpotOrder"
	TypeMsgBatchCancelSpotOrders            = "batchCancelSpotOrders"
	TypeMsgCreateDerivativeLimitOrder       = "createDerivativeLimitOrder"
	TypeMsgBatchCreateDerivativeLimitOrders = "batchCreateDerivativeLimitOrder"
	TypeMsgCreateDerivativeMarketOrder      = "createDerivativeMarketOrder"
	TypeMsgCancelDerivativeOrder            = "cancelDerivativeOrder"
	TypeMsgBatchCancelDerivativeOrders      = "batchCancelDerivativeOrder"
	TypeMsgSubaccountTransfer               = "subaccountTransfer"
	TypeMsgExternalTransfer                 = "externalTransfer"
	TypeMsgIncreasePositionMargin           = "increasePositionMargin"
	TypeMsgLiquidatePosition                = "liquidatePosition"
	TypeMsgInstantSpotMarketLaunch          = "instantSpotMarketLaunch"
	TypeMsgInstantPerpetualMarketLaunch     = "instantPerpetualMarketLaunch"
	TypeMsgInstantExpiryFuturesMarketLaunch = "instantExpiryFuturesMarketLaunch"
	TypeMsgBatchUpdateOrders                = "batchUpdateOrders"
	TypeMsgPrivilegedExecuteContract        = "privilegedExecuteContract"
	TypeMsgRewardsOptOut                    = "rewardsOptOut"
	TypeMsgInstantBinaryOptionsMarketLaunch = "instantBinaryOptionsMarketLaunch"
	TypeMsgCreateBinaryOptionsLimitOrder    = "createBinaryOptionsLimitOrder"
	TypeMsgCreateBinaryOptionsMarketOrder   = "createBinaryOptionsMarketOrder"
	TypeMsgCancelBinaryOptionsOrder         = "cancelBinaryOptionsOrder"
	TypeMsgAdminUpdateBinaryOptionsMarket   = "adminUpdateBinaryOptionsMarket"
	TypeMsgBatchCancelBinaryOptionsOrders   = "batchCancelBinaryOptionsOrders"
	TypeMsgReclaimLockedFunds               = "reclaimLockedFunds"
	TypeMsgUpdateParams                     = "updateParams"
)

func (msg MsgUpdateParams) Route() string { return RouterKey }

func (msg MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

func (msg MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	if err := msg.Params.Validate(); err != nil {
		return err
	}

	return nil
}

func (msg *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(msg))
}

func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{addr}
}

func (o *SpotOrder) ValidateBasic(senderAddr sdk.AccAddress) error {
	if !IsHexHash(o.MarketId) {
		return errors.Wrap(ErrMarketInvalid, o.MarketId)
	}
	switch o.OrderType {
	case OrderType_BUY, OrderType_SELL, OrderType_BUY_PO, OrderType_SELL_PO, OrderType_BUY_ATOMIC, OrderType_SELL_ATOMIC:
		// do nothing
	default:
		return errors.Wrap(ErrUnrecognizedOrderType, string(o.OrderType))
	}

	// for legacy support purposes, allow non-conditional orders to send a 0 trigger price
	if o.TriggerPrice != nil && (o.TriggerPrice.IsNil() || o.TriggerPrice.IsNegative() || o.TriggerPrice.GT(MaxOrderPrice)) {
		return ErrInvalidTriggerPrice
	}

	if o.OrderInfo.FeeRecipient != "" {
		_, err := sdk.AccAddressFromBech32(o.OrderInfo.FeeRecipient)
		if err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, o.OrderInfo.FeeRecipient)
		}
	}
	return o.OrderInfo.ValidateBasic(senderAddr, false, false)
}

func (o *OrderInfo) ValidateBasic(senderAddr sdk.AccAddress, hasBinaryPriceBand, isDerivative bool) error {
	if err := CheckValidSubaccountIDOrNonce(senderAddr, o.SubaccountId); err != nil {
		return err
	}

	if o.Cid != "" && !IsValidCid(o.Cid) {
		return errors.Wrap(ErrInvalidCid, o.Cid)
	}

	if o.Quantity.IsNil() || o.Quantity.LTE(sdk.ZeroDec()) || o.Quantity.GT(MaxOrderQuantity) {
		return errors.Wrap(ErrInvalidQuantity, o.Quantity.String())
	}

	if hasBinaryPriceBand {
		// o.Price.GT(MaxOrderPrice) is correct (as opposed to o.Price.GT(sdk.OneDec())), because the price here is scaled
		// and we have no idea what the scale factor of the market is here when we execute ValidateBasic(), and thus we allow
		// very high ceiling price to cover all cases
		if o.Price.IsNil() || o.Price.LT(sdk.ZeroDec()) || o.Price.GT(MaxOrderPrice) {
			return errors.Wrap(ErrInvalidPrice, o.Price.String())
		}
	} else {
		if o.Price.IsNil() || o.Price.LTE(sdk.ZeroDec()) || o.Price.GT(MaxOrderPrice) {
			return errors.Wrap(ErrInvalidPrice, o.Price.String())
		}
	}

	if isDerivative && !hasBinaryPriceBand && o.Price.LT(MinDerivativeOrderPrice) {
		return errors.Wrap(ErrInvalidPrice, o.Price.String())
	}

	return nil
}

func (o *DerivativeOrder) ValidateBasic(senderAddr sdk.AccAddress, hasBinaryPriceBand bool) error {
	if !IsHexHash(o.MarketId) {
		return errors.Wrap(ErrMarketInvalid, o.MarketId)
	}

	switch o.OrderType {
	case OrderType_BUY, OrderType_SELL, OrderType_BUY_PO, OrderType_SELL_PO, OrderType_STOP_BUY, OrderType_STOP_SELL, OrderType_TAKE_BUY, OrderType_TAKE_SELL, OrderType_BUY_ATOMIC, OrderType_SELL_ATOMIC:
		// do nothing
	default:
		return errors.Wrap(ErrUnrecognizedOrderType, string(o.OrderType))
	}

	if o.Margin.IsNil() || o.Margin.LT(sdk.ZeroDec()) {
		return errors.Wrap(ErrInsufficientOrderMargin, o.Margin.String())
	}

	if o.Margin.GT(MaxOrderMargin) {
		return errors.Wrap(ErrTooMuchOrderMargin, o.Margin.String())
	}

	// for legacy support purposes, allow non-conditional orders to send a 0 trigger price
	if o.TriggerPrice != nil && (o.TriggerPrice.IsNil() || o.TriggerPrice.IsNegative() || o.TriggerPrice.GT(MaxOrderPrice)) {
		return ErrInvalidTriggerPrice
	}

	if o.IsConditional() && (o.TriggerPrice == nil || o.TriggerPrice.LT(MinDerivativeOrderPrice)) { /*||
		!o.IsConditional() && o.TriggerPrice != nil */ // commented out this check since FE is sending to us 0.0 trigger price for all orders
		return errors.Wrapf(ErrInvalidTriggerPrice, "Mismatch between triggerPrice: %v and orderType: %v, or triggerPrice is incorrect", o.TriggerPrice, o.OrderType)
	}

	if o.OrderInfo.FeeRecipient != "" {
		_, err := sdk.AccAddressFromBech32(o.OrderInfo.FeeRecipient)
		if err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, o.OrderInfo.FeeRecipient)
		}
	}
	return o.OrderInfo.ValidateBasic(senderAddr, hasBinaryPriceBand, !hasBinaryPriceBand)
}

func (o *OrderData) ValidateBasic(senderAddr sdk.AccAddress) error {
	if !IsHexHash(o.MarketId) {
		return errors.Wrap(ErrMarketInvalid, o.MarketId)
	}

	if err := CheckValidSubaccountIDOrNonce(senderAddr, o.SubaccountId); err != nil {
		return err
	}

	// order data must contain either an order hash or cid
	if o.Cid == "" && o.OrderHash == "" {
		return ErrOrderHashInvalid
	}

	if o.Cid != "" && !IsValidCid(o.Cid) {
		return errors.Wrap(ErrInvalidCid, o.Cid)
	}

	if o.OrderHash != "" && !IsValidOrderHash(o.OrderHash) {
		return errors.Wrap(ErrOrderHashInvalid, o.OrderHash)
	}
	return nil
}

func (o *OrderData) GetIdentifier() any {
	return GetOrderIdentifier(o.OrderHash, o.Cid)
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgDeposit) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgDeposit) Type() string { return TypeMsgDeposit }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgDeposit) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)

	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !msg.Amount.IsValid() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if IsNonceDerivedSubaccountID(msg.SubaccountId) {
		subaccountID, err := GetSubaccountIDOrDeriveFromNonce(senderAddr, msg.SubaccountId)
		if err != nil {
			return errors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
		}
		if IsDefaultSubaccountID(subaccountID) {
			return errors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
		}
	} else {
		// deposits to externally owned subaccounts are allowed but they MUST be explicitly specified
		_, ok := IsValidSubaccountID(msg.SubaccountId)
		if !ok {
			return errors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
		}
		if IsDefaultSubaccountID(common.HexToHash(msg.SubaccountId)) {
			return errors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
		}
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgDeposit) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgWithdraw) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgWithdraw) Type() string { return TypeMsgWithdraw }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgWithdraw) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !msg.Amount.IsValid() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if err := CheckValidSubaccountIDOrNonce(senderAddr, msg.SubaccountId); err != nil {
		return err
	}

	subaccountID, err := GetSubaccountIDOrDeriveFromNonce(senderAddr, msg.SubaccountId)
	if err != nil {
		return errors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}

	if IsDefaultSubaccountID(subaccountID) {
		return errors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgWithdraw) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgInstantSpotMarketLaunch) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgInstantSpotMarketLaunch) Type() string { return TypeMsgInstantSpotMarketLaunch }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgInstantSpotMarketLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" || len(msg.Ticker) > MaxTickerLength {
		return errors.Wrap(ErrInvalidTicker, "ticker should not be empty or exceed 30 characters")
	}
	if msg.BaseDenom == "" {
		return errors.Wrap(ErrInvalidBaseDenom, "base denom should not be empty")
	}
	if msg.QuoteDenom == "" {
		return errors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if msg.BaseDenom == msg.QuoteDenom {
		return ErrSameDenoms
	}

	if err := ValidateTickSize(msg.MinPriceTickSize); err != nil {
		return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(msg.MinQuantityTickSize); err != nil {
		return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgInstantSpotMarketLaunch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgInstantSpotMarketLaunch) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgInstantPerpetualMarketLaunch) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgInstantPerpetualMarketLaunch) Type() string { return TypeMsgInstantPerpetualMarketLaunch }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgInstantPerpetualMarketLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" || len(msg.Ticker) > MaxTickerLength {
		return errors.Wrap(ErrInvalidTicker, "ticker should not be empty or exceed 30 characters")
	}
	if msg.QuoteDenom == "" {
		return errors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	oracleParams := NewOracleParams(msg.OracleBase, msg.OracleQuote, msg.OracleScaleFactor, msg.OracleType)
	if err := oracleParams.ValidateBasic(); err != nil {
		return err
	}
	if err := ValidateMakerFee(msg.MakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(msg.TakerFeeRate); err != nil {
		return err
	}
	if err := ValidateMarginRatio(msg.InitialMarginRatio); err != nil {
		return err
	}
	if err := ValidateMarginRatio(msg.MaintenanceMarginRatio); err != nil {
		return err
	}
	if msg.MakerFeeRate.GT(msg.TakerFeeRate) {
		return ErrFeeRatesRelation
	}
	if msg.InitialMarginRatio.LT(msg.MaintenanceMarginRatio) {
		return ErrMarginsRelation
	}
	if err := ValidateTickSize(msg.MinPriceTickSize); err != nil {
		return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(msg.MinQuantityTickSize); err != nil {
		return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgInstantPerpetualMarketLaunch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgInstantPerpetualMarketLaunch) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgInstantBinaryOptionsMarketLaunch) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgInstantBinaryOptionsMarketLaunch) Type() string {
	return TypeMsgInstantBinaryOptionsMarketLaunch
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgInstantBinaryOptionsMarketLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" || len(msg.Ticker) > MaxTickerLength {
		return errors.Wrap(ErrInvalidTicker, "ticker should not be empty or exceed 30 characters")
	}
	if msg.OracleSymbol == "" {
		return errors.Wrap(ErrInvalidOracle, "oracle symbol should not be empty")
	}
	if msg.OracleProvider == "" {
		return errors.Wrap(ErrInvalidOracle, "oracle provider should not be empty")
	}
	if msg.OracleType != oracletypes.OracleType_Provider {
		return errors.Wrap(ErrInvalidOracleType, msg.OracleType.String())
	}
	if msg.OracleScaleFactor > MaxOracleScaleFactor {
		return ErrExceedsMaxOracleScaleFactor
	}
	if err := ValidateMakerFee(msg.MakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(msg.TakerFeeRate); err != nil {
		return err
	}
	if msg.MakerFeeRate.GT(msg.TakerFeeRate) {
		return ErrFeeRatesRelation
	}
	if msg.ExpirationTimestamp >= msg.SettlementTimestamp || msg.ExpirationTimestamp < 0 || msg.SettlementTimestamp < 0 {
		return ErrInvalidExpiry
	}
	if msg.Admin != "" {
		_, err := sdk.AccAddressFromBech32(msg.Admin)
		if err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Admin)
		}
	}
	if msg.QuoteDenom == "" {
		return errors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if err := ValidateTickSize(msg.MinPriceTickSize); err != nil {
		return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(msg.MinQuantityTickSize); err != nil {
		return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgInstantBinaryOptionsMarketLaunch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgInstantBinaryOptionsMarketLaunch) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgInstantExpiryFuturesMarketLaunch) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgInstantExpiryFuturesMarketLaunch) Type() string {
	return TypeMsgInstantExpiryFuturesMarketLaunch
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgInstantExpiryFuturesMarketLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" || len(msg.Ticker) > MaxTickerLength {
		return errors.Wrap(ErrInvalidTicker, "ticker should not be empty or exceed 30 characters")
	}
	if msg.QuoteDenom == "" {
		return errors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}

	oracleParams := NewOracleParams(msg.OracleBase, msg.OracleQuote, msg.OracleScaleFactor, msg.OracleType)
	if err := oracleParams.ValidateBasic(); err != nil {
		return err
	}
	if msg.Expiry <= 0 {
		return errors.Wrap(ErrInvalidExpiry, "expiry should not be empty")
	}
	if err := ValidateMakerFee(msg.MakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(msg.TakerFeeRate); err != nil {
		return err
	}
	if err := ValidateMarginRatio(msg.InitialMarginRatio); err != nil {
		return err
	}
	if err := ValidateMarginRatio(msg.MaintenanceMarginRatio); err != nil {
		return err
	}
	if msg.MakerFeeRate.GT(msg.TakerFeeRate) {
		return ErrFeeRatesRelation
	}
	if msg.InitialMarginRatio.LT(msg.MaintenanceMarginRatio) {
		return ErrMarginsRelation
	}
	if err := ValidateTickSize(msg.MinPriceTickSize); err != nil {
		return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(msg.MinQuantityTickSize); err != nil {
		return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgInstantExpiryFuturesMarketLaunch) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgInstantExpiryFuturesMarketLaunch) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgCreateSpotLimitOrder) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgCreateSpotLimitOrder) Type() string { return TypeMsgCreateSpotLimitOrder }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgCreateSpotLimitOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // We don't need to check if sender is empty.
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if err := msg.Order.ValidateBasic(senderAddr); err != nil {
		return err
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgCreateSpotLimitOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgCreateSpotLimitOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgBatchCreateSpotLimitOrders) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgBatchCreateSpotLimitOrders) Type() string { return TypeMsgBatchCreateSpotLimitOrders }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgBatchCreateSpotLimitOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // We don't need to check if sender is empty.
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Orders) == 0 {
		return errors.Wrap(ErrOrderDoesntExist, "must create at least 1 order")
	}

	for idx := range msg.Orders {
		order := msg.Orders[idx]
		if err := order.ValidateBasic(senderAddr); err != nil {
			return err
		}
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgBatchCreateSpotLimitOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgBatchCreateSpotLimitOrders) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgCreateSpotMarketOrder) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgCreateSpotMarketOrder) Type() string { return TypeMsgCreateSpotMarketOrder }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgCreateSpotMarketOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order.OrderType == OrderType_BUY_PO || msg.Order.OrderType == OrderType_SELL_PO {
		return errors.Wrap(ErrInvalidOrderTypeForMessage, "Spot market order can't be a post only order")
	}

	if err := msg.Order.ValidateBasic(senderAddr); err != nil {
		return err
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgCreateSpotMarketOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgCreateSpotMarketOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg *MsgCancelSpotOrder) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg *MsgCancelSpotOrder) Type() string { return TypeMsgCancelSpotOrder }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgCancelSpotOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	orderData := OrderData{
		MarketId:     msg.MarketId,
		SubaccountId: msg.SubaccountId,
		OrderHash:    msg.OrderHash,
		Cid:          msg.Cid,
	}
	return orderData.ValidateBasic(senderAddr)
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgCancelSpotOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg *MsgCancelSpotOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg *MsgBatchCancelSpotOrders) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg *MsgBatchCancelSpotOrders) Type() string { return TypeMsgBatchCancelSpotOrders }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgBatchCancelSpotOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Data) == 0 {
		return errors.Wrap(ErrOrderDoesntExist, "must cancel at least 1 order")
	}

	for idx := range msg.Data {
		if err := msg.Data[idx].ValidateBasic(senderAddr); err != nil {
			return err
		}
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgBatchCancelSpotOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg *MsgBatchCancelSpotOrders) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgCreateDerivativeLimitOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateDerivativeLimitOrder) Type() string { return TypeMsgCreateDerivativeLimitOrder }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateDerivativeLimitOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Order.OrderType == OrderType_BUY_ATOMIC || msg.Order.OrderType == OrderType_SELL_ATOMIC {
		return errors.Wrap(ErrInvalidOrderTypeForMessage, "Derivative limit orders can't be atomic orders")
	}
	if err := msg.Order.ValidateBasic(senderAddr, false); err != nil {
		return err
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateDerivativeLimitOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateDerivativeLimitOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func NewMsgCreateBinaryOptionsLimitOrder(
	sender sdk.AccAddress,
	market *BinaryOptionsMarket,
	subaccountID string,
	feeRecipient string,
	price, quantity sdk.Dec,
	orderType OrderType,
	isReduceOnly bool,
) *MsgCreateBinaryOptionsLimitOrder {
	margin := GetRequiredBinaryOptionsOrderMargin(price, quantity, market.OracleScaleFactor, orderType, isReduceOnly)

	return &MsgCreateBinaryOptionsLimitOrder{
		Sender: sender.String(),
		Order: DerivativeOrder{
			MarketId: market.MarketId,
			OrderInfo: OrderInfo{
				SubaccountId: subaccountID,
				FeeRecipient: feeRecipient,
				Price:        price,
				Quantity:     quantity,
			},
			OrderType:    orderType,
			Margin:       margin,
			TriggerPrice: nil,
		},
	}
}

// Route should return the name of the module
func (msg MsgCreateBinaryOptionsLimitOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateBinaryOptionsLimitOrder) Type() string {
	return TypeMsgCreateBinaryOptionsLimitOrder
}

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateBinaryOptionsLimitOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Order.OrderType.IsConditional() {
		return errors.Wrap(ErrUnrecognizedOrderType, string(msg.Order.OrderType))
	}
	if err := msg.Order.ValidateBasic(senderAddr, true); err != nil {
		return err
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateBinaryOptionsLimitOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateBinaryOptionsLimitOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgBatchCreateDerivativeLimitOrders) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBatchCreateDerivativeLimitOrders) Type() string {
	return TypeMsgBatchCreateDerivativeLimitOrders
}

// ValidateBasic runs stateless checks on the message
func (msg MsgBatchCreateDerivativeLimitOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Orders) == 0 {
		return errors.Wrap(ErrOrderDoesntExist, "must create at least 1 order")
	}

	for idx := range msg.Orders {
		order := msg.Orders[idx]
		if err := order.ValidateBasic(senderAddr, false); err != nil {
			return err
		}
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgBatchCreateDerivativeLimitOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBatchCreateDerivativeLimitOrders) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route should return the name of the module
func (msg MsgCreateDerivativeMarketOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateDerivativeMarketOrder) Type() string { return TypeMsgCreateDerivativeMarketOrder }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateDerivativeMarketOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order.OrderType == OrderType_BUY_PO || msg.Order.OrderType == OrderType_SELL_PO {
		return errors.Wrap(ErrInvalidOrderTypeForMessage, "Derivative market order can't be a post only order")
	}

	if err := msg.Order.ValidateBasic(senderAddr, false); err != nil {
		return err
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateDerivativeMarketOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateDerivativeMarketOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func NewMsgCreateBinaryOptionsMarketOrder(
	sender sdk.AccAddress,
	market *BinaryOptionsMarket,
	subaccountID string,
	feeRecipient string,
	price, quantity sdk.Dec,
	orderType OrderType,
	isReduceOnly bool,
) *MsgCreateBinaryOptionsMarketOrder {
	margin := GetRequiredBinaryOptionsOrderMargin(price, quantity, market.OracleScaleFactor, orderType, isReduceOnly)

	return &MsgCreateBinaryOptionsMarketOrder{
		Sender: sender.String(),
		Order: DerivativeOrder{
			MarketId: market.MarketId,
			OrderInfo: OrderInfo{
				SubaccountId: subaccountID,
				FeeRecipient: feeRecipient,
				Price:        price,
				Quantity:     quantity,
			},
			OrderType:    orderType,
			Margin:       margin,
			TriggerPrice: nil,
		},
	}
}

// Route should return the name of the module
func (msg MsgCreateBinaryOptionsMarketOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateBinaryOptionsMarketOrder) Type() string {
	return TypeMsgCreateBinaryOptionsMarketOrder
}

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateBinaryOptionsMarketOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order.OrderType == OrderType_BUY_PO || msg.Order.OrderType == OrderType_SELL_PO {
		return errors.Wrap(ErrInvalidOrderTypeForMessage, "market order can't be a post only order")
	}
	if msg.Order.OrderType.IsConditional() {
		return errors.Wrap(ErrUnrecognizedOrderType, string(msg.Order.OrderType))
	}

	if err := msg.Order.ValidateBasic(senderAddr, true); err != nil {
		return err
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateBinaryOptionsMarketOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateBinaryOptionsMarketOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg *MsgCancelDerivativeOrder) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface. It should return the action.
func (msg *MsgCancelDerivativeOrder) Type() string {
	return TypeMsgCancelDerivativeOrder
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgCancelDerivativeOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	orderData := OrderData{
		MarketId:     msg.MarketId,
		SubaccountId: msg.SubaccountId,
		OrderHash:    msg.OrderHash,
		Cid:          msg.Cid,
	}
	return orderData.ValidateBasic(senderAddr)
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgCancelDerivativeOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg *MsgCancelDerivativeOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg *MsgBatchCancelDerivativeOrders) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface. It should return the action.
func (msg *MsgBatchCancelDerivativeOrders) Type() string {
	return TypeMsgBatchCancelDerivativeOrders
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgBatchCancelDerivativeOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Data) == 0 {
		return errors.Wrap(ErrOrderDoesntExist, "must cancel at least 1 order")
	}

	for idx := range msg.Data {
		if err := msg.Data[idx].ValidateBasic(senderAddr); err != nil {
			return err
		}
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgBatchCancelDerivativeOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg *MsgBatchCancelDerivativeOrders) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg *MsgCancelBinaryOptionsOrder) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface. It should return the action.
func (msg *MsgCancelBinaryOptionsOrder) Type() string {
	return TypeMsgCancelBinaryOptionsOrder
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgCancelBinaryOptionsOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	orderData := OrderData{
		MarketId:     msg.MarketId,
		SubaccountId: msg.SubaccountId,
		OrderHash:    msg.OrderHash,
		Cid:          msg.Cid,
	}
	return orderData.ValidateBasic(senderAddr)
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgCancelBinaryOptionsOrder) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg *MsgCancelBinaryOptionsOrder) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgBatchCancelBinaryOptionsOrders) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface. It should return the action.
func (msg *MsgBatchCancelBinaryOptionsOrders) Type() string {
	return TypeMsgBatchCancelBinaryOptionsOrders
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgBatchCancelBinaryOptionsOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Data) == 0 {
		return errors.Wrap(ErrOrderDoesntExist, "must cancel at least 1 order")
	}

	for idx := range msg.Data {
		if err := msg.Data[idx].ValidateBasic(senderAddr); err != nil {
			return err
		}
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgBatchCancelBinaryOptionsOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg *MsgBatchCancelBinaryOptionsOrders) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSubaccountTransfer) Route() string {
	return RouterKey
}

func (msg *MsgSubaccountTransfer) Type() string {
	return TypeMsgSubaccountTransfer
}

func (msg *MsgSubaccountTransfer) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if !msg.Amount.IsValid() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if err := CheckValidSubaccountIDOrNonce(senderAddr, msg.SourceSubaccountId); err != nil {
		return err
	}

	if err := CheckValidSubaccountIDOrNonce(senderAddr, msg.DestinationSubaccountId); err != nil {
		return err
	}

	sourceSubaccount, err := GetSubaccountIDOrDeriveFromNonce(senderAddr, msg.SourceSubaccountId)
	if err != nil {
		return errors.Wrap(ErrBadSubaccountID, msg.SourceSubaccountId)
	}

	destinationSubaccount, err := GetSubaccountIDOrDeriveFromNonce(senderAddr, msg.DestinationSubaccountId)
	if err != nil {
		return errors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}

	if IsDefaultSubaccountID(sourceSubaccount) {
		return errors.Wrap(ErrBadSubaccountID, msg.SourceSubaccountId)
	}

	if IsDefaultSubaccountID(destinationSubaccount) {
		return errors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}

	if !bytes.Equal(SubaccountIDToSdkAddress(sourceSubaccount).Bytes(), SubaccountIDToSdkAddress(destinationSubaccount).Bytes()) {
		return errors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}

	return nil
}

func (msg *MsgSubaccountTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgSubaccountTransfer) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgExternalTransfer) Route() string {
	return RouterKey
}

func (msg *MsgExternalTransfer) Type() string {
	return TypeMsgExternalTransfer
}

func (msg *MsgExternalTransfer) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !msg.Amount.IsValid() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if err := CheckValidSubaccountIDOrNonce(senderAddr, msg.SourceSubaccountId); err != nil {
		return err
	}

	sourceSubaccountId, err := GetSubaccountIDOrDeriveFromNonce(senderAddr, msg.SourceSubaccountId)
	if err != nil {
		return errors.Wrap(ErrBadSubaccountID, msg.SourceSubaccountId)
	}

	if IsDefaultSubaccountID(common.HexToHash(msg.SourceSubaccountId)) {
		return errors.Wrap(ErrBadSubaccountID, msg.SourceSubaccountId)
	}

	_, ok := IsValidSubaccountID(msg.DestinationSubaccountId)
	if !ok || IsDefaultSubaccountID(common.HexToHash(msg.DestinationSubaccountId)) {
		return errors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}

	if !bytes.Equal(SubaccountIDToSdkAddress(sourceSubaccountId).Bytes(), senderAddr.Bytes()) {
		return errors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}
	return nil
}

func (msg *MsgExternalTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgExternalTransfer) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgIncreasePositionMargin) Route() string {
	return RouterKey
}

func (msg *MsgIncreasePositionMargin) Type() string {
	return TypeMsgIncreasePositionMargin
}

func (msg *MsgIncreasePositionMargin) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !IsHexHash(msg.MarketId) {
		return errors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	if !msg.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if msg.Amount.GT(MaxOrderMargin) {
		return errors.Wrap(ErrTooMuchOrderMargin, msg.Amount.String())
	}

	if err := CheckValidSubaccountIDOrNonce(senderAddr, msg.SourceSubaccountId); err != nil {
		return err
	}

	_, ok := IsValidSubaccountID(msg.DestinationSubaccountId)
	if !ok {
		return errors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}

	return nil
}

func (msg *MsgIncreasePositionMargin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgIncreasePositionMargin) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgPrivilegedExecuteContract) Route() string {
	return RouterKey
}

func (msg *MsgPrivilegedExecuteContract) Type() string {
	return TypeMsgPrivilegedExecuteContract
}

func (msg *MsgPrivilegedExecuteContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	// funds must either be "empty" or a valid funds coins string
	if !msg.HasEmptyFunds() {
		if coins, err := sdk.ParseDecCoins(msg.Funds); err != nil || !coins.IsAllPositive() {
			return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Funds)
		}
	}

	_, err = sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.ContractAddress)
	}

	var e wasmxtypes.ExecutionData
	if err := json.Unmarshal([]byte(msg.Data), &e); err != nil {
		return errors.Wrap(err, msg.Data)
	}

	if e.Name == "" {
		return errors.Wrap(ErrBadField, "name should not be empty")
	} else if e.Origin != "" && e.Origin != msg.Sender {
		return errors.Wrap(ErrBadField, "origin must match sender or be empty")
	}

	return nil
}

func (msg *MsgPrivilegedExecuteContract) HasEmptyFunds() bool {
	return msg.Funds == "" || msg.Funds == "0" || msg.Funds == "0inj"
}

func (msg *MsgPrivilegedExecuteContract) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgPrivilegedExecuteContract) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgRewardsOptOut) Route() string {
	return RouterKey
}

func (msg *MsgRewardsOptOut) Type() string {
	return TypeMsgRewardsOptOut
}

func (msg *MsgRewardsOptOut) ValidateBasic() error {

	return nil
}

func (msg *MsgRewardsOptOut) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRewardsOptOut) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgLiquidatePosition) Route() string {
	return RouterKey
}

func (msg *MsgLiquidatePosition) Type() string {
	return TypeMsgLiquidatePosition
}

func (msg *MsgLiquidatePosition) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)

	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !IsHexHash(msg.MarketId) {
		return errors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	_, ok := IsValidSubaccountID(msg.SubaccountId)
	if !ok {
		return errors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}

	if msg.Order != nil {
		liquidatorSubaccountID, err := GetSubaccountIDOrDeriveFromNonce(senderAddr, msg.Order.OrderInfo.SubaccountId)
		if err != nil {
			return err
		}

		// cannot liquidate own position with an order
		if liquidatorSubaccountID == common.HexToHash(msg.SubaccountId) {
			return ErrInvalidLiquidationOrder
		}

		if err := msg.Order.ValidateBasic(senderAddr, false); err != nil {
			return err
		}
	}

	return nil
}

func (msg *MsgLiquidatePosition) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgLiquidatePosition) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgBatchUpdateOrders) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgBatchUpdateOrders) Type() string { return TypeMsgBatchUpdateOrders }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgBatchUpdateOrders) ValidateBasic() error {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	hasCancelAllMarketId := len(msg.SpotMarketIdsToCancelAll) > 0 || len(msg.DerivativeMarketIdsToCancelAll) > 0 || len(msg.BinaryOptionsMarketIdsToCancelAll) > 0

	// for MsgBatchUpdateOrders, empty subaccountIDs do not count as the default subaccount
	hasSubaccountIdForCancelAll := msg.SubaccountId != ""

	if hasCancelAllMarketId && !hasSubaccountIdForCancelAll {
		return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains cancel all marketIDs but no subaccountID")
	}

	if hasSubaccountIdForCancelAll && !hasCancelAllMarketId {
		return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains subaccountID but no cancel all marketIDs")
	}

	if hasSubaccountIdForCancelAll {
		if err := CheckValidSubaccountIDOrNonce(sender, msg.SubaccountId); err != nil {
			return err
		}

		hasDuplicateSpotMarketIds := HasDuplicatesHexHash(msg.SpotMarketIdsToCancelAll)
		if hasDuplicateSpotMarketIds {
			return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate cancel all spot market ids")
		}

		hasDuplicateDerivativesMarketIds := HasDuplicatesHexHash(msg.DerivativeMarketIdsToCancelAll)
		if hasDuplicateDerivativesMarketIds {
			return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate cancel all derivative market ids")
		}
		hasDuplicateBinaryOptionsMarketIds := HasDuplicatesHexHash(msg.BinaryOptionsMarketIdsToCancelAll)
		if hasDuplicateBinaryOptionsMarketIds {
			return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate cancel all binary options market ids")
		}
	}

	if !hasSubaccountIdForCancelAll &&
		len(msg.DerivativeOrdersToCancel) == 0 &&
		len(msg.SpotOrdersToCancel) == 0 &&
		len(msg.DerivativeOrdersToCreate) == 0 &&
		len(msg.SpotOrdersToCreate) == 0 &&
		len(msg.BinaryOptionsOrdersToCreate) == 0 &&
		len(msg.BinaryOptionsOrdersToCancel) == 0 {
		return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg is empty")
	}

	hasDuplicateSpotOrderToCancel := HasDuplicatesOrder(msg.SpotOrdersToCancel)
	if hasDuplicateSpotOrderToCancel {
		return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate spot order to cancel")
	}

	hasDuplicateDerivativeOrderToCancel := HasDuplicatesOrder(msg.DerivativeOrdersToCancel)
	if hasDuplicateDerivativeOrderToCancel {
		return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate derivative order to cancel")
	}

	hasDuplicateBinaryOptionsOrderToCancel := HasDuplicatesOrder(msg.BinaryOptionsOrdersToCancel)
	if hasDuplicateBinaryOptionsOrderToCancel {
		return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate binary options order to cancel")
	}

	if len(msg.SpotMarketIdsToCancelAll) > 0 && len(msg.SpotOrdersToCancel) > 0 {
		seen := make(map[common.Hash]struct{})
		for _, marketID := range msg.SpotMarketIdsToCancelAll {
			if !IsHexHash(marketID) {
				return errors.Wrap(ErrMarketInvalid, marketID)
			}
			seen[common.HexToHash(marketID)] = struct{}{}
		}

		for idx := range msg.SpotOrdersToCancel {
			if _, ok := seen[common.HexToHash(msg.SpotOrdersToCancel[idx].MarketId)]; ok {
				return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains order to cancel in a spot market that is also in cancel all")
			}
		}
	}

	if len(msg.DerivativeMarketIdsToCancelAll) > 0 && len(msg.DerivativeOrdersToCancel) > 0 {
		seen := make(map[common.Hash]struct{})
		for _, marketID := range msg.DerivativeMarketIdsToCancelAll {
			if !IsHexHash(marketID) {
				return errors.Wrap(ErrMarketInvalid, marketID)
			}
			seen[common.HexToHash(marketID)] = struct{}{}
		}

		for idx := range msg.DerivativeOrdersToCancel {
			if _, ok := seen[common.HexToHash(msg.DerivativeOrdersToCancel[idx].MarketId)]; ok {
				return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains order to cancel in a derivative market that is also in cancel all")
			}
		}
	}

	if len(msg.BinaryOptionsMarketIdsToCancelAll) > 0 && len(msg.BinaryOptionsOrdersToCancel) > 0 {
		seen := make(map[common.Hash]struct{})
		for _, marketID := range msg.BinaryOptionsMarketIdsToCancelAll {
			if !IsHexHash(marketID) {
				return errors.Wrap(ErrMarketInvalid, marketID)
			}
			seen[common.HexToHash(marketID)] = struct{}{}
		}

		for idx := range msg.BinaryOptionsOrdersToCancel {
			if _, ok := seen[common.HexToHash(msg.BinaryOptionsOrdersToCancel[idx].MarketId)]; ok {
				return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains order to cancel in a binary options market that is also in cancel all")
			}
		}
	}

	for idx := range msg.SpotOrdersToCancel {
		if err := msg.SpotOrdersToCancel[idx].ValidateBasic(sender); err != nil {
			return err
		}
	}

	for idx := range msg.DerivativeOrdersToCancel {
		if err := msg.DerivativeOrdersToCancel[idx].ValidateBasic(sender); err != nil {
			return err
		}
	}
	for idx := range msg.BinaryOptionsOrdersToCancel {
		if err := msg.BinaryOptionsOrdersToCancel[idx].ValidateBasic(sender); err != nil {
			return err
		}
	}

	for idx := range msg.SpotOrdersToCreate {
		if err := msg.SpotOrdersToCreate[idx].ValidateBasic(sender); err != nil {
			return err
		}
		if msg.SpotOrdersToCreate[idx].OrderType.IsAtomic() { // must be checked separately as type is SpotOrder, so it won't check for atomic orders properly
			return errors.Wrap(ErrInvalidOrderTypeForMessage, "Spot limit orders can't be atomic orders")
		}
	}

	for idx := range msg.DerivativeOrdersToCreate {
		if err := msg.DerivativeOrdersToCreate[idx].ValidateBasic(sender, false); err != nil {
			return err
		}
		if msg.DerivativeOrdersToCreate[idx].OrderType.IsAtomic() {
			return errors.Wrap(ErrInvalidOrderTypeForMessage, "Derivative limit orders can't be atomic orders")
		}
	}

	for idx := range msg.BinaryOptionsOrdersToCreate {
		if err := msg.BinaryOptionsOrdersToCreate[idx].ValidateBasic(sender, true); err != nil {
			return err
		}
		if msg.BinaryOptionsOrdersToCreate[idx].OrderType.IsAtomic() {
			return errors.Wrap(ErrInvalidOrderTypeForMessage, "Binary limit orders can't be atomic orders")
		}
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgBatchUpdateOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgBatchUpdateOrders) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgAdminUpdateBinaryOptionsMarket) Route() string {
	return RouterKey
}

func (msg *MsgAdminUpdateBinaryOptionsMarket) Type() string {
	return TypeMsgAdminUpdateBinaryOptionsMarket
}

func (msg *MsgAdminUpdateBinaryOptionsMarket) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !IsHexHash(msg.MarketId) {
		return errors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	if (msg.SettlementTimestamp > 0 && msg.ExpirationTimestamp >= msg.SettlementTimestamp) ||
		msg.ExpirationTimestamp < 0 {
		return ErrInvalidExpiry
	}

	if msg.SettlementTimestamp < 0 {
		return ErrInvalidSettlement
	}

	// price is either nil (not set), -1 (demolish with refund) or [0..1] (demolish with settle)
	switch {
	case msg.SettlementPrice == nil,
		msg.SettlementPrice.IsNil():
		// ok
	case msg.SettlementPrice.Equal(BinaryOptionsMarketRefundFlagPrice),
		msg.SettlementPrice.GTE(sdk.ZeroDec()) && msg.SettlementPrice.LTE(MaxBinaryOptionsOrderPrice):
		if msg.Status != MarketStatus_Demolished {
			return errors.Wrapf(ErrInvalidMarketStatus, "status should be set to demolished when the settlement price is set, status: %s", msg.Status.String())
		}
		// ok
	default:
		return errors.Wrap(ErrInvalidPrice, msg.SettlementPrice.String())
	}
	// admin can only change status to demolished
	switch msg.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Demolished:
	default:
		return errors.Wrap(ErrInvalidMarketStatus, msg.Status.String())
	}

	return nil
}

func (msg *MsgAdminUpdateBinaryOptionsMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgAdminUpdateBinaryOptionsMarket) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgReclaimLockedFunds) Route() string {
	return RouterKey
}

func (msg *MsgReclaimLockedFunds) Type() string {
	return TypeMsgReclaimLockedFunds
}

func (msg *MsgReclaimLockedFunds) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	// TODO: restrict the msg.Sender to be a specific EOA?
	// Placeholder for now obviously, if we decide so, change this check to the actual address
	// if !senderAddr.Equals(senderAddr) {
	//	return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	// }

	lockedPubKey := sdksecp256k1.PubKey{
		Key: msg.LockedAccountPubKey,
	}
	correctPubKey := ethsecp256k1.PubKey{
		Key: msg.LockedAccountPubKey,
	}
	lockedAddress := sdk.AccAddress(lockedPubKey.Address())
	recipientAddress := sdk.AccAddress(correctPubKey.Address())

	data := ConstructFundsReclaimMessage(
		recipientAddress,
		lockedAddress,
	)

	msgSignData := MsgSignData{
		Signer: lockedAddress.Bytes(),
		Data:   data,
	}

	if err := msgSignData.ValidateBasic(); err != nil {
		return nil
	}

	tx := legacytx.NewStdTx(
		[]sdk.Msg{&MsgSignDoc{
			SignType: "sign/MsgSignData",
			Value:    msgSignData,
		}},
		legacytx.StdFee{
			Amount: sdk.Coins{},
			Gas:    0,
		},
		//nolint:staticcheck // we know it's deprecated and we think it's okay
		[]legacytx.StdSignature{
			{
				PubKey:    &lockedPubKey,
				Signature: msg.Signature,
			},
		},
		"",
	)

	if err := tx.ValidateBasic(); err != nil {
		return err
	}

	aminoJSONHandler := legacytx.NewStdTxSignModeHandler()

	signingData := signing.SignerData{
		ChainID:       "",
		AccountNumber: 0,
		Sequence:      0,
	}

	signBz, err := aminoJSONHandler.GetSignBytes(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, signingData, tx)
	if err != nil {
		return err
	}

	if !lockedPubKey.VerifySignature(signBz, tx.GetSignatures()[0]) {
		return errors.Wrapf(ErrBadField, "signature verification failed with signature %s on signBz %s, msg.Signature is %s", common.Bytes2Hex(tx.GetSignatures()[0]), string(signBz), common.Bytes2Hex(msg.Signature))
	}

	return nil
}

func (msg *MsgReclaimLockedFunds) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgReclaimLockedFunds) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// / Skeleton sdk.Msg interface implementation
var _ sdk.Msg = &MsgSignData{}
var _ legacytx.LegacyMsg = &MsgSignData{}

func (msg *MsgSignData) ValidateBasic() error {
	if msg.Signer.Empty() {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Signer.String())
	}

	return nil
}

func (msg *MsgSignData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

func (m *MsgSignData) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgSignData) Route() string {
	return RouterKey
}

func (m *MsgSignData) Type() string {
	return "signData"
}

// / Skeleton sdk.Msg interface implementation
var _ sdk.Msg = &MsgSignDoc{}
var _ legacytx.LegacyMsg = &MsgSignDoc{}

func (msg *MsgSignDoc) ValidateBasic() error {
	return nil
}

func (msg *MsgSignDoc) GetSigners() []sdk.AccAddress {
	return msg.Value.GetSigners()
}

func (m *MsgSignDoc) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgSignDoc) Route() string {
	return RouterKey
}

func (m *MsgSignDoc) Type() string {
	return "signDoc"
}

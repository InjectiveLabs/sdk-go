package types

import (
	"bytes"
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"

	wasmxtypes "github.com/InjectiveLabs/sdk-go/chain/wasmx/types"
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
	_ sdk.Msg = &MsgExec{}
	_ sdk.Msg = &MsgRegisterAsDMM{}
	_ sdk.Msg = &MsgInstantBinaryOptionsMarketLaunch{}
	_ sdk.Msg = &MsgCreateBinaryOptionsLimitOrder{}
	_ sdk.Msg = &MsgCreateBinaryOptionsMarketOrder{}
	_ sdk.Msg = &MsgCancelBinaryOptionsOrder{}
	_ sdk.Msg = &MsgAdminUpdateBinaryOptionsMarket{}
	_ sdk.Msg = &MsgBatchCancelBinaryOptionsOrders{}
)

func (o *SpotOrder) ValidateBasic(senderAddr sdk.AccAddress) error {
	if !IsHexHash(o.MarketId) {
		return sdkerrors.Wrap(ErrMarketInvalid, o.MarketId)
	}
	switch o.OrderType {
	case OrderType_BUY, OrderType_SELL, OrderType_BUY_PO, OrderType_SELL_PO:
		// do nothing
	default:
		return sdkerrors.Wrap(ErrUnrecognizedOrderType, string(o.OrderType))
	}
	if o.TriggerPrice != nil && (o.TriggerPrice.IsNil() || o.TriggerPrice.LT(sdk.ZeroDec()) || o.TriggerPrice.GT(MaxOrderPrice)) {
		return sdkerrors.Wrap(ErrInvalidTriggerPrice, o.TriggerPrice.String())
	}

	_, err := sdk.AccAddressFromBech32(o.OrderInfo.FeeRecipient)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, o.OrderInfo.FeeRecipient)
	}

	return o.OrderInfo.ValidateBasic(senderAddr, false)
}

func (o *OrderInfo) ValidateBasic(senderAddr sdk.AccAddress, hasBinaryPriceBand bool) error {
	subaccountAddress, ok := IsValidSubaccountID(o.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, o.SubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, senderAddr.String())
	}

	if o.Quantity.IsNil() || o.Quantity.LTE(sdk.ZeroDec()) || o.Quantity.GT(MaxOrderQuantity) {
		return sdkerrors.Wrap(ErrInvalidQuantity, o.Quantity.String())
	}

	if hasBinaryPriceBand {
		if o.Price.IsNil() || o.Price.LT(sdk.ZeroDec()) || o.Price.GT(MaxOrderPrice) {
			return sdkerrors.Wrap(ErrInvalidPrice, o.Price.String())
		}
	} else {
		if o.Price.IsNil() || o.Price.LTE(sdk.ZeroDec()) || o.Price.GT(MaxOrderPrice) {
			return sdkerrors.Wrap(ErrInvalidPrice, o.Price.String())
		}
	}

	return nil
}

func (o *DerivativeOrder) ValidateBasic(senderAddr sdk.AccAddress, hasBinaryPriceBand bool) error {
	if !IsHexHash(o.MarketId) {
		return sdkerrors.Wrap(ErrMarketInvalid, o.MarketId)
	}
	switch o.OrderType {
	case OrderType_BUY, OrderType_SELL, OrderType_BUY_PO, OrderType_SELL_PO:
		// do nothing
	default:
		return sdkerrors.Wrap(ErrUnrecognizedOrderType, string(o.OrderType))
	}

	if o.Margin.IsNil() || o.Margin.LT(sdk.ZeroDec()) || o.Margin.GT(MaxOrderPrice) {
		return sdkerrors.Wrap(ErrInsufficientOrderMargin, o.Margin.String())
	}
	if o.TriggerPrice != nil && (o.TriggerPrice.IsNil() || o.TriggerPrice.LT(sdk.ZeroDec()) || o.TriggerPrice.GT(MaxOrderPrice)) {
		return sdkerrors.Wrap(ErrInvalidTriggerPrice, o.TriggerPrice.String())
	}

	_, err := sdk.AccAddressFromBech32(o.OrderInfo.FeeRecipient)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, o.OrderInfo.FeeRecipient)
	}

	return o.OrderInfo.ValidateBasic(senderAddr, hasBinaryPriceBand)
}

func (o *OrderData) ValidateBasic(senderAddr sdk.AccAddress) error {
	if !IsHexHash(o.MarketId) {
		return sdkerrors.Wrap(ErrMarketInvalid, o.MarketId)
	}

	subaccountAddress, ok := IsValidSubaccountID(o.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, o.SubaccountId)
	}

	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, senderAddr.String())
	}

	if ok = IsValidOrderHash(o.OrderHash); !ok {
		return sdkerrors.Wrap(ErrOrderHashInvalid, o.OrderHash)
	}

	return nil
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgDeposit) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgDeposit) Type() string { return "msgDeposit" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)

	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if len(msg.SubaccountId) == 0 {
		return nil
	} else if _, ok := IsValidSubaccountID(msg.SubaccountId); !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
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
func (msg MsgWithdraw) Type() string { return "msgWithdraw" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgWithdraw) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	subaccountAddress, ok := IsValidSubaccountID(msg.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
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
func (msg MsgInstantSpotMarketLaunch) Type() string { return "instantSpotMarketLaunch" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgInstantSpotMarketLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" || len(msg.Ticker) > MaxTickerLength {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty or exceed 30 characters")
	}
	if msg.BaseDenom == "" {
		return sdkerrors.Wrap(ErrInvalidBaseDenom, "base denom should not be empty")
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if msg.BaseDenom == msg.QuoteDenom {
		return ErrSameDenoms
	}

	if err := ValidateTickSize(msg.MinPriceTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(msg.MinQuantityTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, err.Error())
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
func (msg MsgInstantPerpetualMarketLaunch) Type() string { return "instantPerpetualMarketLaunch" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgInstantPerpetualMarketLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" || len(msg.Ticker) > MaxTickerLength {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty or exceed 30 characters")
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
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
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(msg.MinQuantityTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, err.Error())
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
	return "instantBinaryOptionsMarketLaunch"
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgInstantBinaryOptionsMarketLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" || len(msg.Ticker) > MaxTickerLength {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty or exceed 30 characters")
	}
	if msg.OracleSymbol == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle symbol should not be empty")
	}
	if msg.OracleProvider == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle provider should not be empty")
	}
	if msg.OracleType != oracletypes.OracleType_Provider {
		return sdkerrors.Wrap(ErrInvalidOracleType, msg.OracleType.String())
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
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Admin)
		}
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if err := ValidateTickSize(msg.MinPriceTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(msg.MinQuantityTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, err.Error())
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
	return "instantExpiryFuturesMarketLaunch"
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgInstantExpiryFuturesMarketLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" || len(msg.Ticker) > MaxTickerLength {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty or exceed 30 characters")
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}

	oracleParams := NewOracleParams(msg.OracleBase, msg.OracleQuote, msg.OracleScaleFactor, msg.OracleType)
	if err := oracleParams.ValidateBasic(); err != nil {
		return err
	}
	if msg.Expiry <= 0 {
		return sdkerrors.Wrap(ErrInvalidExpiry, "expiry should not be empty")
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
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(msg.MinQuantityTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, err.Error())
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
func (msg MsgCreateSpotLimitOrder) Type() string { return "createSpotLimitOrder" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgCreateSpotLimitOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // We don't need to check if sender is empty.
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
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
func (msg MsgBatchCreateSpotLimitOrders) Type() string { return "batchCreateSpotLimitOrders" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgBatchCreateSpotLimitOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil { // We don't need to check if sender is empty.
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Orders) == 0 {
		return sdkerrors.Wrap(ErrOrderDoesntExist, "must create at least 1 order")
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
func (msg MsgCreateSpotMarketOrder) Type() string { return "createSpotMarketOrder" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgCreateSpotMarketOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order.OrderType == OrderType_BUY_PO || msg.Order.OrderType == OrderType_SELL_PO {
		return sdkerrors.Wrap(ErrInvalidOrderTypeForMessage, "Spot market order can't be a post only order")
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
func (msg *MsgCancelSpotOrder) Type() string { return "cancelSpotOrder" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgCancelSpotOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	orderData := OrderData{
		MarketId:     msg.MarketId,
		SubaccountId: msg.SubaccountId,
		OrderHash:    msg.OrderHash,
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
func (msg *MsgBatchCancelSpotOrders) Type() string { return "batchCancelSpotOrders" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgBatchCancelSpotOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Data) == 0 {
		return sdkerrors.Wrap(ErrOrderDoesntExist, "must cancel at least 1 order")
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
func (msg MsgCreateDerivativeLimitOrder) Type() string { return "createDerivativeLimitOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateDerivativeLimitOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
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
func (msg MsgCreateBinaryOptionsLimitOrder) Type() string { return "createBinaryOptionsLimitOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateBinaryOptionsLimitOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
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
	return "batchCreateDerivativeLimitOrder"
}

// ValidateBasic runs stateless checks on the message
func (msg MsgBatchCreateDerivativeLimitOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Orders) == 0 {
		return sdkerrors.Wrap(ErrOrderDoesntExist, "must create at least 1 order")
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
func (msg MsgCreateDerivativeMarketOrder) Type() string { return "createDerivativeMarketOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateDerivativeMarketOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order.OrderType == OrderType_BUY_PO || msg.Order.OrderType == OrderType_SELL_PO {
		return sdkerrors.Wrap(ErrInvalidOrderTypeForMessage, "Derivative market order can't be a post only order")
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
func (msg MsgCreateBinaryOptionsMarketOrder) Type() string { return "createBinaryOptionsMarketOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateBinaryOptionsMarketOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.Order.OrderType == OrderType_BUY_PO || msg.Order.OrderType == OrderType_SELL_PO {
		return sdkerrors.Wrap(ErrInvalidOrderTypeForMessage, "market order can't be a post only order")
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
	return "cancelDerivativeOrder"
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgCancelDerivativeOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	orderData := OrderData{
		MarketId:     msg.MarketId,
		SubaccountId: msg.SubaccountId,
		OrderHash:    msg.OrderHash,
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
	return "batchCancelDerivativeOrder"
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgBatchCancelDerivativeOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Data) == 0 {
		return sdkerrors.Wrap(ErrOrderDoesntExist, "must cancel at least 1 order")
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
	return "cancelBinaryOptionsOrder"
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgCancelBinaryOptionsOrder) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	orderData := OrderData{
		MarketId:     msg.MarketId,
		SubaccountId: msg.SubaccountId,
		OrderHash:    msg.OrderHash,
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
	return "batchCancelBinaryOptionsOrders"
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg *MsgBatchCancelBinaryOptionsOrders) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if len(msg.Data) == 0 {
		return sdkerrors.Wrap(ErrOrderDoesntExist, "must cancel at least 1 order")
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
	return "subaccountTransfer"
}

func (msg *MsgSubaccountTransfer) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	subaccountAddress, ok := IsValidSubaccountID(msg.SourceSubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SourceSubaccountId)
	}
	destSubaccountAddress, ok := IsValidSubaccountID(msg.DestinationSubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), destSubaccountAddress.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
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
	return "externalTransfer"
}

func (msg *MsgExternalTransfer) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	sourceSubaccountAddress, ok := IsValidSubaccountID(msg.SourceSubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SourceSubaccountId)
	}

	_, ok = IsValidSubaccountID(msg.DestinationSubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}
	if !bytes.Equal(sourceSubaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
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
	return "increasePositionMargin"
}

func (msg *MsgIncreasePositionMargin) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !IsHexHash(msg.MarketId) {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	sourceSubaccountAddress, ok := IsValidSubaccountID(msg.SourceSubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SourceSubaccountId)
	}
	if !bytes.Equal(sourceSubaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
	}

	_, ok = IsValidSubaccountID(msg.DestinationSubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
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

func (msg *MsgExec) Route() string {
	return RouterKey
}

func (msg *MsgExec) Type() string {
	return "exec"
}

func (msg *MsgExec) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !msg.BankFunds.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.BankFunds.String())
	}

	if msg.DepositsSubaccountId != "" {
		if addr, ok := IsValidSubaccountID(msg.DepositsSubaccountId); !ok {
			return sdkerrors.Wrap(ErrBadSubaccountID, msg.DepositsSubaccountId)
		} else if !bytes.Equal(addr.Bytes(), senderAddr.Bytes()) {
			return sdkerrors.Wrap(ErrBadSubaccountID, msg.DepositsSubaccountId)
		}
	}

	if !msg.DepositFunds.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.DepositFunds.String())
	}

	_, err = sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.ContractAddress)
	}

	var e wasmxtypes.ExecutionData
	if err := json.Unmarshal([]byte(msg.Data), &e); err != nil {
		return sdkerrors.Wrap(err, msg.Data)
	}

	if e.Name == "" {
		return sdkerrors.Wrap(ErrBadField, "name should not be empty")
	} else if e.Origin != "" && e.Origin != msg.Sender {
		return sdkerrors.Wrap(ErrBadField, "origin must match sender or be empty")
	}

	return nil
}

func (msg *MsgExec) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgExec) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgRegisterAsDMM) Route() string {
	return RouterKey
}

func (msg *MsgRegisterAsDMM) Type() string {
	return "registerAsDMM"
}

func (msg *MsgRegisterAsDMM) ValidateBasic() error {
	if msg.Sender != msg.DmmAccount {
		return sdkerrors.Wrap(ErrInvalidDMMSender, fmt.Sprintf("Sender: %s doesn't match dmm account: %s", msg.Sender, msg.DmmAccount))
	}

	return nil
}

func (msg *MsgRegisterAsDMM) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgRegisterAsDMM) GetSigners() []sdk.AccAddress {
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
	return "liquidatePosition"
}

func (msg *MsgLiquidatePosition) ValidateBasic() error {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)

	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !IsHexHash(msg.MarketId) {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	_, ok := IsValidSubaccountID(msg.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}

	if msg.Order != nil {
		// cannot liquidate own position with an order
		if msg.Order.OrderInfo.SubaccountId == msg.SubaccountId {
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
func (msg MsgBatchUpdateOrders) Type() string { return "batchUpdateOrders" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgBatchUpdateOrders) ValidateBasic() error {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	hasCancelAllMarketId := len(msg.SpotMarketIdsToCancelAll) > 0 || len(msg.DerivativeMarketIdsToCancelAll) > 0 || len(msg.BinaryOptionsMarketIdsToCancelAll) > 0
	hasSubaccountIdForCancelAll := msg.SubaccountId != ""

	if hasCancelAllMarketId && !hasSubaccountIdForCancelAll {
		return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains cancel all marketIDs but no subaccountID")
	}

	if hasSubaccountIdForCancelAll && !hasCancelAllMarketId {
		return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains subaccountID but no cancel all marketIDs")
	}

	if hasSubaccountIdForCancelAll {
		subaccountAddress, ok := IsValidSubaccountID(msg.SubaccountId)
		if !ok {
			return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
		}
		if !bytes.Equal(subaccountAddress.Bytes(), sender.Bytes()) {
			return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
		}

		hasDuplicateSpotMarketIds := HasDuplicatesHexHash(msg.SpotMarketIdsToCancelAll)
		if hasDuplicateSpotMarketIds {
			return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate cancel all spot market ids")
		}

		hasDuplicateDerivativesMarketIds := HasDuplicatesHexHash(msg.DerivativeMarketIdsToCancelAll)
		if hasDuplicateDerivativesMarketIds {
			return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate cancel all derivative market ids")
		}
		hasDuplicateBinaryOptionsMarketIds := HasDuplicatesHexHash(msg.BinaryOptionsMarketIdsToCancelAll)
		if hasDuplicateBinaryOptionsMarketIds {
			return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate cancel all binary options market ids")
		}
	}

	if !hasSubaccountIdForCancelAll &&
		len(msg.DerivativeOrdersToCancel) == 0 &&
		len(msg.SpotOrdersToCancel) == 0 &&
		len(msg.DerivativeOrdersToCreate) == 0 &&
		len(msg.SpotOrdersToCreate) == 0 &&
		len(msg.BinaryOptionsOrdersToCreate) == 0 &&
		len(msg.BinaryOptionsOrdersToCancel) == 0 {
		return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg is empty")
	}

	hasDuplicateSpotOrderToCancel := HasDuplicatesOrder(msg.SpotOrdersToCancel)
	if hasDuplicateSpotOrderToCancel {
		return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate spot order to cancel")
	}

	hasDuplicateDerivativeOrderToCancel := HasDuplicatesOrder(msg.DerivativeOrdersToCancel)
	if hasDuplicateDerivativeOrderToCancel {
		return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate derivative order to cancel")
	}

	hasDuplicateBinaryOptionsOrderToCancel := HasDuplicatesOrder(msg.BinaryOptionsOrdersToCancel)
	if hasDuplicateBinaryOptionsOrderToCancel {
		return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate binary options order to cancel")
	}

	if len(msg.SpotMarketIdsToCancelAll) > 0 && len(msg.SpotOrdersToCancel) > 0 {
		seen := make(map[common.Hash]struct{})
		for _, marketID := range msg.SpotMarketIdsToCancelAll {
			if !IsHexHash(marketID) {
				return sdkerrors.Wrap(ErrMarketInvalid, marketID)
			}
			seen[common.HexToHash(marketID)] = struct{}{}
		}

		for idx := range msg.SpotOrdersToCancel {
			if _, ok := seen[common.HexToHash(msg.SpotOrdersToCancel[idx].MarketId)]; ok {
				return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains order to cancel in a spot market that is also in cancel all")
			}
		}
	}

	if len(msg.DerivativeMarketIdsToCancelAll) > 0 && len(msg.DerivativeOrdersToCancel) > 0 {
		seen := make(map[common.Hash]struct{})
		for _, marketID := range msg.DerivativeMarketIdsToCancelAll {
			if !IsHexHash(marketID) {
				return sdkerrors.Wrap(ErrMarketInvalid, marketID)
			}
			seen[common.HexToHash(marketID)] = struct{}{}
		}

		for idx := range msg.DerivativeOrdersToCancel {
			if _, ok := seen[common.HexToHash(msg.DerivativeOrdersToCancel[idx].MarketId)]; ok {
				return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains order to cancel in a derivative market that is also in cancel all")
			}
		}
	}

	if len(msg.BinaryOptionsMarketIdsToCancelAll) > 0 && len(msg.BinaryOptionsOrdersToCancel) > 0 {
		seen := make(map[common.Hash]struct{})
		for _, marketID := range msg.BinaryOptionsMarketIdsToCancelAll {
			if !IsHexHash(marketID) {
				return sdkerrors.Wrap(ErrMarketInvalid, marketID)
			}
			seen[common.HexToHash(marketID)] = struct{}{}
		}

		for idx := range msg.BinaryOptionsOrdersToCancel {
			if _, ok := seen[common.HexToHash(msg.BinaryOptionsOrdersToCancel[idx].MarketId)]; ok {
				return sdkerrors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains order to cancel in a binary options market that is also in cancel all")
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
	}

	for idx := range msg.DerivativeOrdersToCreate {
		if err := msg.DerivativeOrdersToCreate[idx].ValidateBasic(sender, false); err != nil {
			return err
		}
	}

	for idx := range msg.BinaryOptionsOrdersToCreate {
		if err := msg.BinaryOptionsOrdersToCreate[idx].ValidateBasic(sender, true); err != nil {
			return err
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
	return "adminUpdateBinaryOptionsMarket"
}

func (msg *MsgAdminUpdateBinaryOptionsMarket) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !IsHexHash(msg.MarketId) {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
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
			return sdkerrors.Wrapf(ErrInvalidMarketStatus, "status should be set to demolished when the settlement price is set, status: %s", msg.Status.String())
		}
		// ok
	default:
		return sdkerrors.Wrap(ErrInvalidPrice, msg.SettlementPrice.String())
	}
	// admin can only change status to demolished
	switch msg.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Demolished:
	default:
		return sdkerrors.Wrap(ErrInvalidMarketStatus, msg.Status.String())
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

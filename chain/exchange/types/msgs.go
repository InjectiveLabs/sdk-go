package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgDeposit{}
	_ sdk.Msg = &MsgWithdraw{}
	_ sdk.Msg = &MsgCreateSpotLimitOrder{}
	_ sdk.Msg = &MsgCreateSpotMarketOrder{}
	_ sdk.Msg = &MsgCancelSpotOrder{}
	_ sdk.Msg = &MsgCreateDerivativeLimitOrder{}
	_ sdk.Msg = &MsgCreateDerivativeMarketOrder{}
	_ sdk.Msg = &MsgCancelDerivativeOrder{}
	_ sdk.Msg = &MsgSubaccountTransfer{}
	_ sdk.Msg = &MsgExternalTransfer{}
	_ sdk.Msg = &MsgIncreasePositionMargin{}
	_ sdk.Msg = &MsgLiquidatePosition{}
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgDeposit) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgDeposit) Type() string { return "msgDeposit" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgDeposit) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	if _, ok := IsValidSubaccountID(msg.SubaccountId); !ok {
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
	if msg.Sender == "" {
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

	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(err, "must provide a valid Bech32 address")
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
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty")
	}
	if msg.BaseDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "base denom should not be empty")
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidBaseDenom, "quote denom should not be empty")
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
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty")
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if msg.Oracle == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle should not be empty")
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
func (msg MsgInstantExpiryFuturesMarketLaunch) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgInstantExpiryFuturesMarketLaunch) Type() string {
	return "instantExpiryFuturesMarketLaunch"
}

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgInstantExpiryFuturesMarketLaunch) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty")
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if msg.Oracle == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle should not be empty")
	}
	if msg.Expiry == 0 {
		return sdkerrors.Wrap(ErrInvalidExpiry, "expirty should not be empty")
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
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Order.MarketId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.Order.MarketId)
	}
	if msg.Order.OrderType < 0 || msg.Order.OrderType > 5 {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, string(msg.Order.OrderType))
	}
	if msg.Order.TriggerPrice.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Order.TriggerPrice.String())
	}
	if msg.Order.OrderInfo.SubaccountId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.Order.OrderInfo.SubaccountId)
	}
	if msg.Order.OrderInfo.FeeRecipient == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Order.OrderInfo.FeeRecipient)
	}
	if msg.Order.OrderInfo.Price.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Order.OrderInfo.Price.String())
	}
	if msg.Order.OrderInfo.Quantity.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Order.OrderInfo.Quantity.String())
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
func (msg MsgCreateSpotMarketOrder) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgCreateSpotMarketOrder) Type() string { return "createSpotMarketOrder" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgCreateSpotMarketOrder) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Order.MarketId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.Order.MarketId)
	}
	if msg.Order.OrderType < 0 || msg.Order.OrderType > 5 {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, string(msg.Order.OrderType))
	}
	if msg.Order.TriggerPrice.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Order.TriggerPrice.String())
	}
	if msg.Order.OrderInfo.SubaccountId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, msg.Order.OrderInfo.SubaccountId)
	}
	if msg.Order.OrderInfo.FeeRecipient == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Order.OrderInfo.FeeRecipient)
	}
	if msg.Order.OrderInfo.Quantity.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Order.OrderInfo.Quantity.String())
	}
	if msg.Order.OrderInfo.Price.IsNil() {
		return sdkerrors.Wrap(ErrOrderInvalid, "order worst price cannot be nil")
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
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	// TODO: check if subaccountId and sender matches or not
	return nil
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

// Route should return the name of the module
func (msg MsgCreateDerivativeLimitOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateDerivativeLimitOrder) Type() string { return "createDerivativeLimitOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateDerivativeLimitOrder) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	//if msg.Order == nil {
	//	return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "no make order specified")
	//}
	//
	//order := msg.Order.ToSignedOrder()
	//quantity := order.TakerAssetAmount
	//price := order.MakerAssetAmount
	//orderHash, err := order.ComputeOrderHash()
	//makerAddress := common.HexToAddress(msg.Order.MakerAddress)
	//
	//if err != nil {
	//	return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("hash check failed: %v", err))
	//} else if quantity == nil || quantity.Cmp(big.NewInt(0)) <= 0 {
	//	return sdkerrors.Wrap(ErrInsufficientOrderQuantity, "insufficient quantity")
	//} else if price == nil || price.Cmp(big.NewInt(0)) <= 0 {
	//	return sdkerrors.Wrap(ErrInsufficientOrderQuantity, "insufficient price")
	//}

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

// Route should return the name of the module
func (msg MsgCreateDerivativeMarketOrder) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateDerivativeMarketOrder) Type() string { return "createDerivativeMarketOrder" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateDerivativeMarketOrder) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
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
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	// TODO: implement this
	return nil
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

func (msg *MsgSubaccountTransfer) Route() string {
	return RouterKey
}

func (msg *MsgSubaccountTransfer) Type() string {
	return "subaccountTransfer"
}

func (msg *MsgSubaccountTransfer) ValidateBasic() error {
	if msg.Sender == "" {
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

	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(err, "must provide a valid Bech32 address")
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
	if msg.Sender == "" {
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

	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(err, "must provide a valid Bech32 address")
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
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	subaccountAddress, ok := IsValidSubaccountID(msg.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}

	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(err, "must provide a valid Bech32 address")
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
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

func (msg *MsgLiquidatePosition) Route() string {
	return RouterKey
}

func (msg *MsgLiquidatePosition) Type() string {
	return "liquidatePosition"
}

func (msg *MsgLiquidatePosition) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	_, ok := IsValidSubaccountID(msg.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrap(err, "must provide a valid Bech32 address")
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

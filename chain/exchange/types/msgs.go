package types

import (
	"bytes"
	"errors"

	"github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	_ sdk.Msg = &MsgInstantSpotMarketLaunch{}
	_ sdk.Msg = &MsgInstantPerpetualMarketLaunch{}
	_ sdk.Msg = &MsgInstantExpiryFuturesMarketLaunch{}
)

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
	if msg.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty")
	}
	if msg.BaseDenom == "" {
		return sdkerrors.Wrap(ErrInvalidBaseDenom, "base denom should not be empty")
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}

	if msg.MinPriceTickSize.IsNil() || msg.MinPriceTickSize.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, msg.MinPriceTickSize.String())
	}
	if msg.MinQuantityTickSize.IsNil() || msg.MinQuantityTickSize.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, msg.MinQuantityTickSize.String())
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
	if msg.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty")
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if msg.OracleBase == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle base should not be empty")
	}
	if msg.OracleQuote == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle quote should not be empty")
	}
	switch msg.OracleType {
	case types.OracleType_Band, types.OracleType_PriceFeed, types.OracleType_Coinbase, types.OracleType_Chainlink, types.OracleType_Razor,
		types.OracleType_Dia, types.OracleType_API3, types.OracleType_Uma, types.OracleType_Pyth, types.OracleType_BandIBC:

	default:
		return sdkerrors.Wrap(ErrInvalidOracleType, msg.OracleType.String())
	}
	if err := ValidateFee(msg.MakerFeeRate); err != nil {
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
		return errors.New("MakerFeeRate cannot be greater than TakerFeeRate")
	}
	if msg.InitialMarginRatio.LT(msg.MaintenanceMarginRatio) {
		return errors.New("MaintenanceMarginRatio cannot be greater than InitialMarginRatio")
	}
	if msg.MinPriceTickSize.IsNil() || msg.MinPriceTickSize.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, msg.MinPriceTickSize.String())
	}
	if msg.MinQuantityTickSize.IsNil() || msg.MinQuantityTickSize.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, msg.MinQuantityTickSize.String())
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
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty")
	}
	if msg.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if msg.OracleBase == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle base should not be empty")
	}
	if msg.OracleQuote == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle quote should not be empty")
	}
	switch msg.OracleType {
	case types.OracleType_Band, types.OracleType_PriceFeed, types.OracleType_Coinbase, types.OracleType_Chainlink, types.OracleType_Razor,
		types.OracleType_Dia, types.OracleType_API3, types.OracleType_Uma, types.OracleType_Pyth, types.OracleType_BandIBC:

	default:
		return sdkerrors.Wrap(ErrInvalidOracleType, msg.OracleType.String())
	}
	if msg.Expiry <= 0 {
		return sdkerrors.Wrap(ErrInvalidExpiry, "expiry should not be empty")
	}
	if err := ValidateFee(msg.MakerFeeRate); err != nil {
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
		return errors.New("MakerFeeRate cannot be greater than TakerFeeRate")
	}
	if msg.InitialMarginRatio.LT(msg.MaintenanceMarginRatio) {
		return errors.New("MaintenanceMarginRatio cannot be greater than InitialMarginRatio")
	}
	if msg.MinPriceTickSize.IsNil() || msg.MinPriceTickSize.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, msg.MinPriceTickSize.String())
	}
	if msg.MinPriceTickSize.IsNil() || msg.MinQuantityTickSize.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, msg.MinQuantityTickSize.String())
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
	if err := msg.Order.ValidateBasic(); err != nil {
		return err
	}

	subaccountAddress, ok := IsValidSubaccountID(msg.Order.OrderInfo.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Order.OrderInfo.SubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
	}

	if msg.Order.OrderInfo.Price.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidPrice, msg.Order.OrderInfo.Price.String())
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
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if err := msg.Order.ValidateBasic(); err != nil {
		return err
	}

	subaccountAddress, ok := IsValidSubaccountID(msg.Order.OrderInfo.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Order.OrderInfo.SubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
	}
	if msg.Order.OrderInfo.Price.IsNil() {
		return sdkerrors.Wrap(ErrInvalidPrice, "order worst price cannot be nil")
	}

	return nil
}

func (msg SpotOrder) ValidateBasic() error {
	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}
	if msg.OrderType <= 0 || msg.OrderType > 2 {
		return sdkerrors.Wrap(ErrUnrecognizedOrderType, string(msg.OrderType))
	}
	if msg.TriggerPrice != nil && msg.TriggerPrice.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidTriggerPrice, msg.TriggerPrice.String())
	}

	_, err := sdk.AccAddressFromBech32(msg.OrderInfo.FeeRecipient)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.OrderInfo.FeeRecipient)
	}
	if msg.OrderInfo.Quantity.IsNil() || msg.OrderInfo.Quantity.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidQuantity, msg.OrderInfo.Quantity.String())
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
	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	subaccountAddress, ok := IsValidSubaccountID(msg.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
	}

	ok = IsValidOrderHash(msg.OrderHash)
	if !ok {
		return sdkerrors.Wrap(ErrOrderHashInvalid, msg.OrderHash)
	}

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
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if err := msg.Order.ValidateBasic(); err != nil {
		return err
	}

	subaccountAddress, ok := IsValidSubaccountID(msg.Order.OrderInfo.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Order.OrderInfo.SubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
	}
	if msg.Order.OrderInfo.Price.IsNil() || msg.Order.OrderInfo.Price.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidPrice, msg.Order.OrderInfo.Price.String())
	}
	if msg.Order.OrderInfo.Quantity.IsNil() || msg.Order.OrderInfo.Quantity.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidQuantity, msg.Order.OrderInfo.Quantity.String())
	}
	if msg.Order.Margin.IsNil() || msg.Order.Margin.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidMargin, msg.Order.Margin.String())
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
	if err := msg.Order.ValidateBasic(); err != nil {
		return err
	}

	subaccountAddress, ok := IsValidSubaccountID(msg.Order.OrderInfo.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Order.OrderInfo.SubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
	}
	if msg.Order.OrderInfo.Price.IsNil() {
		return sdkerrors.Wrap(ErrInvalidPrice, "order worst price cannot be nil")
	}

	return nil
}

func (msg DerivativeOrder) ValidateBasic() error {
	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}
	if msg.OrderType <= 0 || msg.OrderType > 2 {
		return sdkerrors.Wrap(ErrUnrecognizedOrderType, string(msg.OrderType))
	}
	if msg.Margin.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInsufficientOrderMargin, msg.Margin.String())
	}
	if msg.TriggerPrice != nil && msg.TriggerPrice.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidTriggerPrice, msg.TriggerPrice.String())
	}

	_, err := sdk.AccAddressFromBech32(msg.OrderInfo.FeeRecipient)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.OrderInfo.FeeRecipient)
	}
	if msg.OrderInfo.Quantity.IsNil() || msg.OrderInfo.Quantity.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalidQuantity, msg.OrderInfo.Quantity.String())
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
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	subaccountAddress, ok := IsValidSubaccountID(msg.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}
	if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
	}

	ok = IsValidOrderHash(msg.OrderHash)
	if !ok {
		return sdkerrors.Wrap(ErrOrderHashInvalid, msg.OrderHash)
	}

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
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}
	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	_, ok := IsValidSubaccountID(msg.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
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
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)

	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	_, ok := IsValidSubaccountID(msg.SubaccountId)
	if !ok {
		return sdkerrors.Wrap(ErrBadSubaccountID, msg.SubaccountId)
	}

	if msg.Order != nil {
		if err := msg.Order.ValidateBasic(); err != nil {
			return err
		}
		if msg.Order.MarketId != msg.MarketId {
			return sdkerrors.Wrap(ErrMarketInvalid, msg.Order.MarketId)
		}

		subaccountAddress, ok := IsValidSubaccountID(msg.Order.OrderInfo.SubaccountId)
		if !ok {
			return sdkerrors.Wrap(ErrBadSubaccountID, msg.Order.OrderInfo.SubaccountId)
		}
		if !bytes.Equal(subaccountAddress.Bytes(), senderAddr.Bytes()) {
			return sdkerrors.Wrap(ErrBadSubaccountID, msg.Sender)
		}

		if msg.Order.OrderInfo.Price.IsNil() || msg.Order.OrderInfo.Price.LTE(sdk.ZeroDec()) {
			return sdkerrors.Wrap(ErrInvalidPrice, msg.Order.OrderInfo.Price.String())
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

package types

import (
	"bytes"
	"encoding/json"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	wasmxtypes "github.com/InjectiveLabs/sdk-go/chain/wasmx/types"
	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"
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
	_ sdk.Msg = &MsgDecreasePositionMargin{}
	_ sdk.Msg = &MsgLiquidatePosition{}
	_ sdk.Msg = &MsgEmergencySettleMarket{}
	_ sdk.Msg = &MsgInstantSpotMarketLaunch{}
	_ sdk.Msg = &MsgBatchUpdateOrders{}
	_ sdk.Msg = &MsgPrivilegedExecuteContract{}
	_ sdk.Msg = &MsgRewardsOptOut{}
	_ sdk.Msg = &MsgInstantBinaryOptionsMarketLaunch{}
	_ sdk.Msg = &MsgCreateBinaryOptionsLimitOrder{}
	_ sdk.Msg = &MsgCreateBinaryOptionsMarketOrder{}
	_ sdk.Msg = &MsgCancelBinaryOptionsOrder{}
	_ sdk.Msg = &MsgAdminUpdateBinaryOptionsMarket{}
	_ sdk.Msg = &MsgBatchCancelBinaryOptionsOrders{}
	_ sdk.Msg = &MsgUpdateParams{}
	_ sdk.Msg = &MsgUpdateSpotMarket{}
	_ sdk.Msg = &MsgUpdateDerivativeMarket{}
	_ sdk.Msg = &MsgBatchExchangeModification{}
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
	TypeMsgDecreasePositionMargin           = "decreasePositionMargin"
	TypeMsgLiquidatePosition                = "liquidatePosition"
	TypeMsgEmergencySettleMarket            = "emergencySettleMarket"
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
	TypeMsgUpdateParams                     = "updateParams"
	TypeMsgUpdateSpotMarket                 = "updateSpotMarket"
	TypeMsgUpdateDerivativeMarket           = "updateDerivativeMarket"
	TypeMsgAuthorizeStakeGrants             = "authorizeStakeGrant"
	TypeMsgActivateStakeGrant               = "acceptStakeGrant"
	TypeMsgBatchExchangeModification        = "batchExchangeModification"
)

type UpdateSpotMarketMessage interface {
	GetAdmin() string
	GetMarketId() string
	GetNewTicker() string
	GetNewMinPriceTickSize() math.LegacyDec
	GetNewMinQuantityTickSize() math.LegacyDec
	GetNewMinNotional() math.LegacyDec
	HasTickerUpdate() bool
	HasMinPriceTickSizeUpdate() bool
	HasMinQuantityTickSizeUpdate() bool
	HasMinNotionalUpdate() bool
}

func (msg MsgUpdateParams) Route() string { return RouterKey }

func (msg MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

func (msg MsgUpdateParams) ValidateBasic() error {
	if err := ValidateAddress(msg.Authority); err != nil {
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

// Declaration to validate adherence to interface
var _ UpdateSpotMarketMessage = &MsgUpdateSpotMarket{}

func (msg MsgUpdateSpotMarket) GetNewMinPriceTickSize() math.LegacyDec {
	return msg.NewMinPriceTickSize
}

func (msg MsgUpdateSpotMarket) GetNewMinQuantityTickSize() math.LegacyDec {
	return msg.NewMinQuantityTickSize
}

func (msg MsgUpdateSpotMarket) GetNewMinNotional() math.LegacyDec {
	return msg.NewMinNotional
}

func (msg *MsgUpdateSpotMarket) ValidateBasic() error {
	return UpdateSpotMarketMessageValidateBasic(msg)
}

func (msg *MsgUpdateSpotMarket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Admin)}
}

func (msg *MsgUpdateSpotMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgUpdateSpotMarket) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSpotMarket) Type() string {
	return TypeMsgUpdateSpotMarket
}

func (msg *MsgUpdateSpotMarket) HasTickerUpdate() bool {
	return msg.NewTicker != ""
}

func (msg *MsgUpdateSpotMarket) HasMinPriceTickSizeUpdate() bool {
	return !msg.NewMinPriceTickSize.IsNil() && !msg.NewMinPriceTickSize.IsZero()
}

func (msg *MsgUpdateSpotMarket) HasMinQuantityTickSizeUpdate() bool {
	return !msg.NewMinQuantityTickSize.IsNil() && !msg.NewMinQuantityTickSize.IsZero()
}

func (msg *MsgUpdateSpotMarket) HasMinNotionalUpdate() bool {
	return !msg.NewMinNotional.IsNil() && !msg.NewMinNotional.IsZero()
}

func (msg *MsgUpdateDerivativeMarket) ValidateBasic() error {
	if err := ValidateAddress(msg.Admin); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Admin)
	}

	if !IsHexHash(msg.MarketId) {
		return errors.Wrap(ErrMarketInvalid, msg.MarketId)
	}

	hasNoUpdate := !msg.HasTickerUpdate() &&
		!msg.HasMinPriceTickSizeUpdate() &&
		!msg.HasMinNotionalUpdate() &&
		!msg.HasMinQuantityTickSizeUpdate() &&
		!msg.HasInitialMarginRatioUpdate() &&
		!msg.HasMaintenanceMarginRatioUpdate()

	if hasNoUpdate {
		return errors.Wrap(ErrBadField, "no update value present")
	}

	if len(msg.NewTicker) > MaxTickerLength {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not exceed %d characters", MaxTickerLength)
	}

	if msg.HasMinPriceTickSizeUpdate() {
		if err := ValidateTickSize(msg.NewMinPriceTickSize); err != nil {
			return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
		}
	}

	if msg.HasMinQuantityTickSizeUpdate() {
		if err := ValidateTickSize(msg.NewMinQuantityTickSize); err != nil {
			return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
		}
	}

	if msg.HasMinNotionalUpdate() {
		if err := ValidateMinNotional(msg.NewMinNotional); err != nil {
			return errors.Wrap(ErrInvalidNotional, err.Error())
		}
	}

	if msg.HasInitialMarginRatioUpdate() {
		if err := ValidateMarginRatio(msg.NewInitialMarginRatio); err != nil {
			return err
		}
	}

	if msg.HasMaintenanceMarginRatioUpdate() {
		if err := ValidateMarginRatio(msg.NewMaintenanceMarginRatio); err != nil {
			return err
		}
	}

	if msg.HasInitialMarginRatioUpdate() || msg.HasMaintenanceMarginRatioUpdate() {
		if msg.NewInitialMarginRatio.LTE(msg.NewMaintenanceMarginRatio) {
			return ErrMarginsRelation
		}
	}

	return nil
}

func (msg *MsgUpdateDerivativeMarket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Admin)}
}

func (msg *MsgUpdateDerivativeMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgUpdateDerivativeMarket) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDerivativeMarket) Type() string {
	return TypeMsgUpdateDerivativeMarket
}

func (msg *MsgUpdateDerivativeMarket) HasTickerUpdate() bool {
	return msg.NewTicker != ""
}

func (msg *MsgUpdateDerivativeMarket) HasMinPriceTickSizeUpdate() bool {
	return !msg.NewMinPriceTickSize.IsNil() && !msg.NewMinPriceTickSize.IsZero()
}

func (msg *MsgUpdateDerivativeMarket) HasMinQuantityTickSizeUpdate() bool {
	return !msg.NewMinQuantityTickSize.IsNil() && !msg.NewMinQuantityTickSize.IsZero()
}

func (msg *MsgUpdateDerivativeMarket) HasInitialMarginRatioUpdate() bool {
	return !msg.NewInitialMarginRatio.IsNil() && !msg.NewInitialMarginRatio.IsZero()
}

func (msg *MsgUpdateDerivativeMarket) HasMaintenanceMarginRatioUpdate() bool {
	return !msg.NewMaintenanceMarginRatio.IsNil() && !msg.NewMaintenanceMarginRatio.IsZero()
}

func (msg *MsgUpdateDerivativeMarket) HasMinNotionalUpdate() bool {
	return !msg.NewMinNotional.IsNil() && !msg.NewMinNotional.IsZero()
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
	if o.TriggerPrice != nil && (o.TriggerPrice.IsNil() || !o.TriggerPrice.IsZero()) {
		return ErrInvalidTriggerPrice
	}

	if o.OrderInfo.FeeRecipient != "" {
		if err := ValidateAddress(o.OrderInfo.FeeRecipient); err != nil {
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

	if o.Quantity.IsNil() || o.Quantity.LTE(math.LegacyZeroDec()) || o.Quantity.GT(MaxOrderQuantity) {
		return errors.Wrap(ErrInvalidQuantity, o.Quantity.String())
	}

	if hasBinaryPriceBand {
		// o.Price.GT(MaxOrderPrice) is correct (as opposed to o.Price.GT(math.LegacyOneDec())), because the price here is scaled
		// and we have no idea what the scale factor of the market is here when we execute ValidateBasic(), and thus we allow
		// very high ceiling price to cover all cases
		if o.Price.IsNil() || o.Price.LT(math.LegacyZeroDec()) || o.Price.GT(MaxOrderPrice) {
			return errors.Wrap(ErrInvalidPrice, o.Price.String())
		}
	} else {
		if o.Price.IsNil() || o.Price.LTE(math.LegacyZeroDec()) || o.Price.GT(MaxOrderPrice) {
			return errors.Wrap(ErrInvalidPrice, o.Price.String())
		}
	}

	if isDerivative && !hasBinaryPriceBand && o.Price.LT(MinDerivativeOrderPrice) {
		return errors.Wrap(ErrInvalidPrice, o.Price.String())
	}

	return nil
}

func (m *DerivativeOrder) ValidateBasic(senderAddr sdk.AccAddress, hasBinaryPriceBand bool) error {
	if !IsHexHash(m.MarketId) {
		return errors.Wrap(ErrMarketInvalid, m.MarketId)
	}

	switch m.OrderType {
	case OrderType_BUY,
		OrderType_SELL,
		OrderType_BUY_PO,
		OrderType_SELL_PO,
		OrderType_STOP_BUY,
		OrderType_STOP_SELL,
		OrderType_TAKE_BUY,
		OrderType_TAKE_SELL,
		OrderType_BUY_ATOMIC,
		OrderType_SELL_ATOMIC:
		// do nothing
	default:
		return errors.Wrap(ErrUnrecognizedOrderType, string(m.OrderType))
	}

	if m.Margin.IsNil() || m.Margin.LT(math.LegacyZeroDec()) {
		return errors.Wrap(ErrInsufficientMargin, m.Margin.String())
	}

	if m.Margin.GT(MaxOrderMargin) {
		return errors.Wrap(ErrTooMuchOrderMargin, m.Margin.String())
	}

	// for legacy support purposes, allow non-conditional orders to send a 0 trigger price
	if m.TriggerPrice != nil && (m.TriggerPrice.IsNil() || m.TriggerPrice.IsNegative() || m.TriggerPrice.GT(MaxOrderPrice)) {
		return ErrInvalidTriggerPrice
	}

	if m.OrderType.IsConditional() && (m.TriggerPrice == nil || m.TriggerPrice.LT(MinDerivativeOrderPrice)) { /*||
		!o.OrderType.IsConditional() && o.TriggerPrice != nil */ // commented out this check since FE is sending to us 0.0 trigger price for all orders
		return errors.Wrapf(
			ErrInvalidTriggerPrice,
			"Mismatch between triggerPrice: %v and orderType: %v, or triggerPrice is incorrect",
			m.TriggerPrice,
			m.OrderType,
		)
	}

	if m.OrderInfo.FeeRecipient != "" {
		_, err := sdk.AccAddressFromBech32(m.OrderInfo.FeeRecipient)
		if err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, m.OrderInfo.FeeRecipient)
		}
	}
	return m.OrderInfo.ValidateBasic(senderAddr, hasBinaryPriceBand, !hasBinaryPriceBand)
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
		return errors.Wrapf(ErrInvalidTicker, "ticker should not be empty or exceed %d characters", MaxTickerLength)
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
	if err := ValidateMinNotional(msg.MinNotional); err != nil {
		return errors.Wrap(ErrInvalidNotional, err.Error())
	}

	if msg.BaseDecimals > MaxDecimals {
		return errors.Wrap(ErrInvalidDenomDecimal, "base decimals is invalid")
	}
	if msg.QuoteDecimals > MaxDecimals {
		return errors.Wrap(ErrInvalidDenomDecimal, "quote decimals is invalid")
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
		return errors.Wrapf(ErrInvalidTicker, "ticker should not be empty or exceed %d characters", MaxTickerLength)
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
	if err := ValidateMinNotional(msg.MinNotional); err != nil {
		return errors.Wrap(ErrInvalidNotional, err.Error())
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
	price, quantity math.LegacyDec,
	orderType OrderType,
	isReduceOnly bool,
) *MsgCreateBinaryOptionsLimitOrder {
	margin := GetRequiredBinaryOptionsOrderMargin(price, quantity, market.OracleScaleFactor, orderType.IsBuy(), isReduceOnly)

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

func NewMsgCreateBinaryOptionsMarketOrderWithScaleFactor(
	sender sdk.AccAddress,
	market *BinaryOptionsMarket,
	subaccountID string,
	feeRecipient string,
	price, quantity math.LegacyDec,
	orderType OrderType,
	isReduceOnly bool,
	scaleFactor uint32,
) *MsgCreateBinaryOptionsMarketOrder {
	margin := GetRequiredBinaryOptionsOrderMargin(price, quantity, scaleFactor, orderType.IsBuy(), isReduceOnly)

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

func NewMsgCreateBinaryOptionsMarketOrder(
	sender sdk.AccAddress,
	market *BinaryOptionsMarket,
	subaccountID string,
	feeRecipient string,
	price, quantity math.LegacyDec,
	orderType OrderType,
	isReduceOnly bool,
) *MsgCreateBinaryOptionsMarketOrder {
	margin := GetRequiredBinaryOptionsOrderMargin(price, quantity, market.OracleScaleFactor, orderType.IsBuy(), isReduceOnly)

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

	sourceSubaccountID, err := GetSubaccountIDOrDeriveFromNonce(senderAddr, msg.SourceSubaccountId)
	if err != nil {
		return errors.Wrap(ErrBadSubaccountID, msg.SourceSubaccountId)
	}

	if IsDefaultSubaccountID(common.HexToHash(msg.SourceSubaccountId)) {
		return errors.Wrap(ErrBadSubaccountID, msg.SourceSubaccountId)
	}

	if _, ok := IsValidSubaccountID(msg.DestinationSubaccountId); !ok {
		return errors.Wrap(ErrBadSubaccountID, msg.DestinationSubaccountId)
	}

	if !bytes.Equal(SubaccountIDToSdkAddress(sourceSubaccountID).Bytes(), senderAddr.Bytes()) {
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

func (msg *MsgDecreasePositionMargin) Route() string {
	return RouterKey
}

func (msg *MsgDecreasePositionMargin) Type() string {
	return TypeMsgDecreasePositionMargin
}

func (msg *MsgDecreasePositionMargin) ValidateBasic() error {
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
	if err := CheckValidSubaccountIDOrNonce(senderAddr, msg.DestinationSubaccountId); err != nil {
		return err
	}

	return nil
}

func (msg *MsgDecreasePositionMargin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgDecreasePositionMargin) GetSigners() []sdk.AccAddress {
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

func (msg *MsgEmergencySettleMarket) Route() string {
	return RouterKey
}

func (msg *MsgEmergencySettleMarket) Type() string {
	return TypeMsgEmergencySettleMarket
}

func (msg *MsgEmergencySettleMarket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)

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

	return nil
}

func (msg *MsgEmergencySettleMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgEmergencySettleMarket) GetSigners() []sdk.AccAddress {
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

		hasDuplicateSpotMarketIDs := chaintypes.HasDuplicate(msg.SpotMarketIdsToCancelAll)
		if hasDuplicateSpotMarketIDs {
			return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate cancel all spot market ids")
		}

		hasDuplicateDerivativesMarketIDs := chaintypes.HasDuplicate(msg.DerivativeMarketIdsToCancelAll)
		if hasDuplicateDerivativesMarketIDs {
			return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate cancel all derivative market ids")
		}
		hasDuplicateBinaryOptionsMarketIDs := chaintypes.HasDuplicate(msg.BinaryOptionsMarketIdsToCancelAll)
		if hasDuplicateBinaryOptionsMarketIDs {
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

	hasDuplicateSpotOrderToCancel := hasDuplicatesOrder(msg.SpotOrdersToCancel)
	if hasDuplicateSpotOrderToCancel {
		return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate spot order to cancel")
	}

	hasDuplicateDerivativeOrderToCancel := hasDuplicatesOrder(msg.DerivativeOrdersToCancel)
	if hasDuplicateDerivativeOrderToCancel {
		return errors.Wrap(ErrInvalidBatchMsgUpdate, "msg contains duplicate derivative order to cancel")
	}

	hasDuplicateBinaryOptionsOrderToCancel := hasDuplicatesOrder(msg.BinaryOptionsOrdersToCancel)
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
		msg.SettlementPrice.GTE(math.LegacyZeroDec()) && msg.SettlementPrice.LTE(MaxBinaryOptionsOrderPrice):
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

func (msg *MsgAuthorizeStakeGrants) Route() string { return RouterKey }

func (msg *MsgAuthorizeStakeGrants) Type() string { return TypeMsgAuthorizeStakeGrants }

func (msg *MsgAuthorizeStakeGrants) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	for idx := range msg.Grants {
		grant := msg.Grants[idx]

		if _, err := sdk.AccAddressFromBech32(grant.Grantee); err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, grant.Grantee)
		}

		if grant.Amount.IsNegative() || grant.Amount.GT(MaxTokenInt) {
			return errors.Wrap(ErrInvalidStakeGrant, grant.Amount.String())

		}
	}
	return nil
}

func (msg *MsgAuthorizeStakeGrants) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgAuthorizeStakeGrants) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg *MsgActivateStakeGrant) Route() string { return RouterKey }

func (msg *MsgActivateStakeGrant) Type() string { return TypeMsgActivateStakeGrant }

func (msg *MsgActivateStakeGrant) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Granter); err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Granter)
	}
	return nil
}

func (msg *MsgActivateStakeGrant) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgActivateStakeGrant) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func ValidateAddress(address string) error {
	_, err := sdk.AccAddressFromBech32(address)

	return err
}

func hasDuplicatesOrder(slice []*OrderData) bool {
	seenHashes := make(map[string]struct{})
	seenCids := make(map[string]struct{})
	for _, item := range slice {
		var hash, cid string
		hash, cid = item.GetOrderHash(), item.GetCid()
		_, hashExists := seenHashes[hash]
		_, cidExists := seenCids[cid]

		if (hash != "" && hashExists) || (cid != "" && cidExists) {
			return true
		}
		seenHashes[hash] = struct{}{}
		seenCids[cid] = struct{}{}
	}
	return false
}

func (msg *MsgBatchExchangeModification) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgBatchExchangeModification) Route() string { return RouterKey }

func (msg *MsgBatchExchangeModification) Type() string { return TypeMsgBatchExchangeModification }

func (msg *MsgBatchExchangeModification) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return errors.Wrap(err, "invalid sender address")
	}

	if err := msg.Proposal.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

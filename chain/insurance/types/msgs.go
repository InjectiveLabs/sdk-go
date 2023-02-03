package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgCreateInsuranceFund{}
	_ sdk.Msg = &MsgUnderwrite{}
	_ sdk.Msg = &MsgRequestRedemption{}
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgCreateInsuranceFund) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgCreateInsuranceFund) Type() string { return "createInsuranceFund" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgCreateInsuranceFund) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty or exceed 30 characters")
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
	if msg.OracleType == oracletypes.OracleType_Unspecified {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle type should not be unspecified")
	}
	if msg.QuoteDenom != msg.InitialDeposit.Denom {
		return sdkerrors.Wrapf(ErrInvalidDepositDenom, "oracle quote denom %s does not match deposit denom %s", msg.QuoteDenom, msg.InitialDeposit.Denom)
	}

	if msg.OracleType == oracletypes.OracleType_Provider && msg.Expiry != BinaryOptionsExpiryFlag {
		return sdkerrors.Wrap(ErrInvalidExpirationTime, "oracle expiration time should be -2 for binary options markets")
	}

	if !msg.InitialDeposit.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.InitialDeposit.String())
	}
	if !msg.InitialDeposit.IsPositive() || msg.InitialDeposit.Amount.GT(MaxUnderwritingAmount) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.InitialDeposit.String())
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgCreateInsuranceFund) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgCreateInsuranceFund) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgUnderwrite) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgUnderwrite) Type() string { return "underwrite" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgUnderwrite) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrInvalidMarketID, msg.MarketId)
	}
	if !msg.Deposit.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Deposit.String())
	}
	if !msg.Deposit.IsPositive() || msg.Deposit.Amount.GT(MaxUnderwritingAmount) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Deposit.String())
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgUnderwrite) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgUnderwrite) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgRequestRedemption) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgRequestRedemption) Type() string { return "requestRedemption" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgRequestRedemption) ValidateBasic() error {
	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.MarketId == "" {
		return sdkerrors.Wrap(ErrInvalidMarketID, msg.MarketId)
	}

	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgRequestRedemption) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgRequestRedemption) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

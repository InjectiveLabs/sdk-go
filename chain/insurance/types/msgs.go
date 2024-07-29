package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
)

const RouterKey = ModuleName

var (
	_ sdk.Msg = &MsgCreateInsuranceFund{}
	_ sdk.Msg = &MsgUnderwrite{}
	_ sdk.Msg = &MsgRequestRedemption{}
	_ sdk.Msg = &MsgUpdateParams{}
)

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgUpdateParams) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgUpdateParams) Type() string { return "updateParams" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errors.Wrap(err, "invalid authority address")
	}

	if err := msg.Params.Validate(); err != nil {
		return err
	}

	return nil
}

// GetSignBytes implements the sdk.Msg interface. It encodes the message for signing
func (msg *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners implements the sdk.Msg interface. It defines whose signature is required
func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// Route implements the sdk.Msg interface. It should return the name of the module
func (msg MsgCreateInsuranceFund) Route() string { return RouterKey }

// Type implements the sdk.Msg interface. It should return the action.
func (msg MsgCreateInsuranceFund) Type() string { return "createInsuranceFund" }

// ValidateBasic implements the sdk.Msg interface. It runs stateless checks on the message
func (msg MsgCreateInsuranceFund) ValidateBasic() error {
	if msg.Sender == "" {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.Ticker == "" || len(msg.Ticker) > 40 {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not be empty or exceed 40 characters")
	}
	if msg.QuoteDenom == "" {
		return errors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if msg.OracleBase == "" {
		return errors.Wrap(ErrInvalidOracle, "oracle base should not be empty")
	}
	if msg.OracleQuote == "" {
		return errors.Wrap(ErrInvalidOracle, "oracle quote should not be empty")
	}
	if msg.OracleType == oracletypes.OracleType_Unspecified {
		return errors.Wrap(ErrInvalidOracle, "oracle type should not be unspecified")
	}
	if msg.QuoteDenom != msg.InitialDeposit.Denom {
		return errors.Wrapf(ErrInvalidDepositDenom, "oracle quote denom %s does not match deposit denom %s", msg.QuoteDenom, msg.InitialDeposit.Denom)
	}

	if msg.OracleType == oracletypes.OracleType_Provider && msg.Expiry != BinaryOptionsExpiryFlag && msg.Expiry != PerpetualExpiryFlag {
		return errors.Wrap(ErrInvalidExpirationTime, "oracle expiration time should be -2 or -1 for OracleType_Provider")
	}

	if !msg.InitialDeposit.IsValid() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.InitialDeposit.String())
	}
	if !msg.InitialDeposit.IsPositive() || msg.InitialDeposit.Amount.GT(MaxUnderwritingAmount) {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.InitialDeposit.String())
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
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.MarketId == "" {
		return errors.Wrap(ErrInvalidMarketID, msg.MarketId)
	}
	if !msg.Deposit.IsValid() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Deposit.String())
	}
	if !msg.Deposit.IsPositive() || msg.Deposit.Amount.GT(MaxUnderwritingAmount) {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Deposit.String())
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
		return errors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}
	if msg.MarketId == "" {
		return errors.Wrap(ErrInvalidMarketID, msg.MarketId)
	}

	if !msg.Amount.IsValid() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	if !msg.Amount.IsPositive() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
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

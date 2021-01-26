package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	_ sdk.Msg = &MsgERC20BridgeMint{}
)

// NewMsgERC20BridgeMint returns init hub msg instance
func NewMsgERC20BridgeMint(mapping string, amount string, address string, proposer sdk.AccAddress) *MsgERC20BridgeMint {
	return &MsgERC20BridgeMint{
		MappingId: mapping,
		Amount:    amount,
		Address:   address,
		Proposer:  proposer,
	}
}

// Route should return the name of the module
func (msg MsgERC20BridgeMint) Route() string { return RouterKey }

// Type should return the action
func (msg MsgERC20BridgeMint) Type() string { return TypeMsgMint }

// ValidateBasic runs stateless checks on the message
func (msg MsgERC20BridgeMint) ValidateBasic() error {
	coins, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return err
	}
	if len(coins) > 1 || coins.Empty() {
		return fmt.Errorf("invalid coins field: %s", coins.String())
	}
	if msg.Proposer.Empty() {
		return errors.New("empty proposer")
	}
	if msg.MappingId == "" {
		return errors.New("empty mapping id")
	}
	if !common.IsHexAddress(msg.Address) {
		return errors.New("invalid eth address")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgERC20BridgeMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgERC20BridgeMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Proposer}
}

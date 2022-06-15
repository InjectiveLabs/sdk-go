package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MintToUser struct {
	SubaccountIDSender string `json:"subaccount_id_sender"`
	Amount             string `json:"amount"`
}

type MintToUserMsg struct {
	MintToUser MintToUser `json:"mint_to_user"`
}

func NewMintToUserMsg(subaccountIDSender string, amount sdk.Int) MintToUserMsg {
	return MintToUserMsg{
		MintToUser: MintToUser{
			SubaccountIDSender: subaccountIDSender,
			Amount:             amount.String(),
		},
	}
}

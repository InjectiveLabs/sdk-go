package types

import (
	"cosmossdk.io/math"
)

type MintToUser struct {
	SubaccountIDSender string `json:"subaccount_id_sender"`
	Amount             string `json:"amount"`
}

type MintToUserMsg struct {
	MintToUser MintToUser `json:"mint_to_user"`
}

func NewMintToUserMsg(subaccountIDSender string, amount math.Int) MintToUserMsg {
	return MintToUserMsg{
		MintToUser: MintToUser{
			SubaccountIDSender: subaccountIDSender,
			Amount:             amount.String(),
		},
	}
}

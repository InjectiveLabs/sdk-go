package v2

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewActiveGrant(granter sdk.AccAddress, amount math.Int) *ActiveGrant {
	return &ActiveGrant{
		Granter: granter.String(),
		Amount:  amount,
	}
}

func NewEffectiveGrant(granter string, amount math.Int, isValid bool) *EffectiveGrant {
	return &EffectiveGrant{
		Granter:         granter,
		NetGrantedStake: amount,
		IsValid:         isValid,
	}
}

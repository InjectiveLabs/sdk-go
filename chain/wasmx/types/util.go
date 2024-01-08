package types

import (
	"github.com/CosmWasm/wasmd/x/wasm/types"
	types2 "github.com/cosmos/cosmos-sdk/types"
)

func IsAllowed(accessConfig types.AccessConfig, actor types2.AccAddress) bool {
	switch accessConfig.Permission {
	case types.AccessTypeAnyOfAddresses:
		for _, v := range accessConfig.Addresses {
			if v == actor.String() {
				return true
			}
		}
		return false
	}
	return false
}

package types

import (
	"github.com/CosmWasm/wasmd/x/wasm/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

func IsAllowed(accessConfig types.AccessConfig, actor sdktypes.AccAddress) bool {
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

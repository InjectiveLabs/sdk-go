package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	tftypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
)

const (
	EVERYONE = "EVERYONE"
	MaxPerm  = uint32(Action_MINT) | uint32(Action_RECEIVE) | uint32(Action_BURN)
)

type WasmHookMsg struct {
	From    sdk.AccAddress `json:"from_address"`
	To      sdk.AccAddress `json:"to_address"`
	Action  string         `json:"action"`
	Amounts sdk.Coins      `json:"amounts"`
}

func (n *Namespace) Validate() error {
	if _, _, err := tftypes.DeconstructDenom(n.Denom); err != nil {
		return errors.Wrap(err, "permissions namespace can only be applied to tokenfactory denoms")
	}

	// role_permissions
	hasEveryoneSet := false
	foundRoles := make(map[string]struct{}, len(n.RolePermissions))
	for _, rolePerm := range n.RolePermissions {
		if rolePerm.Role == EVERYONE {
			hasEveryoneSet = true
		}
		if _, ok := foundRoles[rolePerm.Role]; ok {
			return errors.Wrapf(ErrInvalidPermission, "permissions for the role %s set multiple times?", rolePerm.Role)
		}
		if rolePerm.Permissions > MaxPerm {
			return errors.Wrapf(ErrInvalidPermission, "permissions %d for the role %s is bigger than maximum expected %d", rolePerm.Permissions, rolePerm.Role, MaxPerm)
		}
		foundRoles[rolePerm.Role] = struct{}{}
	}

	if !hasEveryoneSet {
		return errors.Wrapf(ErrInvalidPermission, "permissions for role %s should be explicitly set", EVERYONE)
	}

	// address_roles
	foundAddresses := make(map[string]struct{}, len(n.AddressRoles))
	for _, addrRoles := range n.AddressRoles {
		if _, err := sdk.AccAddressFromBech32(addrRoles.Address); err != nil {
			return errors.Wrapf(err, "invalid address %s", addrRoles.Address)
		}
		if _, ok := foundAddresses[addrRoles.Address]; ok {
			return errors.Wrapf(ErrInvalidRole, "address %s is assigned new roles multiple times?", addrRoles.Address)
		}
		for _, role := range addrRoles.Roles {
			if _, ok := foundRoles[role]; !ok {
				return errors.Wrapf(ErrUnknownRole, "role %s has no defined permissions", role)
			}
			if role == EVERYONE {
				return errors.Wrapf(ErrInvalidRole, "role %s should not be explicitly attached to address, you need to remove address from the list completely instead", EVERYONE)
			}
		}
		foundAddresses[addrRoles.Address] = struct{}{}
	}

	// existing wasm hook contract
	if n.WasmHook != "" {
		if _, err := sdk.AccAddressFromBech32(n.WasmHook); err != nil {
			return errors.Wrap(err, "invalid WasmHook address")
		}
	}

	return nil
}

package types

import (
	"encoding/json"

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

func NewWasmHookMsg(fromAddr, toAddr sdk.AccAddress, action Action, amount sdk.Coin) WasmHookMsg {
	return WasmHookMsg{
		From:    fromAddr,
		To:      toAddr,
		Action:  action.String(),
		Amounts: sdk.NewCoins(amount),
	}
}

func GetWasmHookMsgBytes(fromAddr, toAddr sdk.AccAddress, action Action, amount sdk.Coin) ([]byte, error) {
	wasmHookMsg := struct {
		SendRestriction WasmHookMsg `json:"send_restriction"`
	}{NewWasmHookMsg(fromAddr, toAddr, action, amount)}

	bz, err := json.Marshal(wasmHookMsg)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func (n *Namespace) Validate() error {
	if err := n.ValidateBasicParams(); err != nil {
		return err
	}

	if err := n.ValidateEveryoneRole(); err != nil {
		return err
	}

	if err := n.ValidateRoles(false); err != nil {
		return err
	}

	return nil
}

func (n *Namespace) ValidateBasicParams() error {
	if _, _, err := tftypes.DeconstructDenom(n.Denom); err != nil {
		return errors.Wrap(err, "permissions namespace can only be applied to tokenfactory denoms")
	}

	// existing wasm hook contract
	if n.WasmHook != "" {
		if _, err := sdk.AccAddressFromBech32(n.WasmHook); err != nil {
			return ErrInvalidWasmHook
		}
	}

	return nil
}

func (n *Namespace) ValidateRoles(isForUpdate bool) error {
	// role_permissions
	foundRoles := make(map[string]struct{}, len(n.RolePermissions))
	for _, rolePerm := range n.RolePermissions {
		if _, ok := foundRoles[rolePerm.Role]; ok {
			return errors.Wrapf(ErrInvalidPermission, "permissions for the role %s set multiple times?", rolePerm.Role)
		}
		if rolePerm.Permissions > MaxPerm {
			return errors.Wrapf(ErrInvalidPermission, "permissions %d for the role %s is bigger than maximum expected %d", rolePerm.Permissions, rolePerm.Role, MaxPerm)
		}
		foundRoles[rolePerm.Role] = struct{}{}
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
			_, found := foundRoles[role]
			if !isForUpdate && !found {
				return errors.Wrapf(ErrUnknownRole, "role %s has no defined permissions", role)
			}
			if role == EVERYONE {
				return errors.Wrapf(ErrInvalidRole, "role %s should not be explicitly attached to address, you need to remove address from the list completely instead", EVERYONE)
			}
		}
		foundAddresses[addrRoles.Address] = struct{}{}
	}

	return nil
}

func (n *Namespace) ValidateEveryoneRole() error {
	// role_permissions
	for _, rolePerm := range n.RolePermissions {
		if rolePerm.Role == EVERYONE {
			return nil
		}
	}

	return errors.Wrapf(ErrInvalidPermission, "permissions for role %s should be explicitly set", EVERYONE)
}

func (n *Namespace) CheckActionValidity(action Action) error {
	// check that action is not paused
	switch action {
	case Action_MINT:
		if n.MintsPaused {
			return errors.Wrap(ErrRestrictedAction, "mints paused")
		}
	case Action_RECEIVE:
		if n.SendsPaused {
			return errors.Wrap(ErrRestrictedAction, "sends paused")
		}
	case Action_BURN:
		if n.BurnsPaused {
			return errors.Wrap(ErrRestrictedAction, "burns paused")
		}
	}
	return nil
}

func (a Action) DeriveActor(fromAddr, toAddr sdk.AccAddress) sdk.AccAddress {
	switch a {
	case Action_MINT, Action_RECEIVE:
		return toAddr
	case Action_BURN:
		return fromAddr
	}
	return fromAddr
}

func NewEmptyVoucher(denom string) sdk.Coin {
	return sdk.NewInt64Coin(denom, 0)
}

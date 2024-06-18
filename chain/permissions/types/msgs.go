package types

import (
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	tftypes "github.com/InjectiveLabs/sdk-go/chain/tokenfactory/types"
)

// constants
const (
	// RouterKey is the message route for slashing
	routerKey = ModuleName

	TypeMsgUpdateParams         = "update_params"
	TypeMsgCreateNamespace      = "create_namespace"
	TypeMsgDeleteNamespace      = "delete_namespace"
	TypeUpdateNamespace         = "update_namespace"
	TypeMsgUpdateNamespaceRoles = "update_namespace_roles"
	TypeMsgRevokeNamespaceRoles = "revoke_namespace_roles"
	TypeMsgClaimVoucher         = "claim_voucher"
)

var (
	_ sdk.Msg = &MsgUpdateParams{}
	_ sdk.Msg = &MsgCreateNamespace{}
	_ sdk.Msg = &MsgDeleteNamespace{}
	_ sdk.Msg = &MsgUpdateNamespace{}
	_ sdk.Msg = &MsgUpdateNamespaceRoles{}
	_ sdk.Msg = &MsgRevokeNamespaceRoles{}
	_ sdk.Msg = &MsgClaimVoucher{}
)

func (m MsgUpdateParams) Route() string { return routerKey }

func (m MsgUpdateParams) Type() string { return TypeMsgUpdateParams }

func (m MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return err
	}

	if err := m.Params.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

func (m MsgCreateNamespace) Route() string { return routerKey }

func (m MsgCreateNamespace) Type() string { return TypeMsgCreateNamespace }

func (msg MsgCreateNamespace) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return err
	}

	if err := msg.Namespace.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *MsgCreateNamespace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgCreateNamespace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m MsgDeleteNamespace) Route() string { return routerKey }

func (m MsgDeleteNamespace) Type() string { return TypeMsgDeleteNamespace }

func (m MsgDeleteNamespace) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return err
	}

	if _, _, err := tftypes.DeconstructDenom(m.NamespaceDenom); err != nil {
		return errors.Wrap(err, "permissions namespace can only be applied to tokenfactory denoms")
	}
	return nil
}

func (m *MsgDeleteNamespace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgDeleteNamespace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateNamespace) Route() string { return routerKey }

func (m MsgUpdateNamespace) Type() string { return TypeUpdateNamespace }

func (m MsgUpdateNamespace) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return err
	}

	if m.WasmHook != nil {
		if _, err := sdk.AccAddressFromBech32(m.WasmHook.NewValue); err != nil {
			return ErrInvalidWasmHook
		}
	}
	return nil
}

func (m *MsgUpdateNamespace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgUpdateNamespace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m MsgUpdateNamespaceRoles) Route() string { return routerKey }

func (m MsgUpdateNamespaceRoles) Type() string { return TypeMsgUpdateNamespaceRoles }

func (msg MsgUpdateNamespaceRoles) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return err
	}
	if _, _, err := tftypes.DeconstructDenom(msg.NamespaceDenom); err != nil {
		return errors.Wrap(err, "permissions namespace can only be applied to tokenfactory denoms")
	}

	for _, role := range msg.AddressRoles {
		if _, err := sdk.AccAddressFromBech32(role.Address); err != nil {
			return errors.Wrapf(err, "invalid address %s", role.Address)
		}
	}

	return nil
}

func (m *MsgUpdateNamespaceRoles) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgUpdateNamespaceRoles) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m MsgRevokeNamespaceRoles) Route() string { return routerKey }

func (m MsgRevokeNamespaceRoles) Type() string { return TypeMsgRevokeNamespaceRoles }

func (msg MsgRevokeNamespaceRoles) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return err
	}

	if _, _, err := tftypes.DeconstructDenom(msg.NamespaceDenom); err != nil {
		return errors.Wrap(err, "permissions namespace can only be applied to tokenfactory denoms")
	}

	// address_roles
	foundAddresses := make(map[string]struct{}, len(msg.AddressRolesToRevoke))
	for _, addrRoles := range msg.AddressRolesToRevoke {
		if _, err := sdk.AccAddressFromBech32(addrRoles.Address); err != nil {
			return errors.Wrapf(err, "invalid address %s", addrRoles.Address)
		}
		if _, ok := foundAddresses[addrRoles.Address]; ok {
			return errors.Wrapf(ErrInvalidRole, "address %s - revoking roles multiple times?", addrRoles.Address)
		}
		for _, role := range addrRoles.Roles {
			if role == EVERYONE {
				return errors.Wrapf(ErrInvalidRole, "role %s can not be set / revoked", EVERYONE)
			}
		}
		foundAddresses[addrRoles.Address] = struct{}{}
	}
	return nil
}

func (m *MsgRevokeNamespaceRoles) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgRevokeNamespaceRoles) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

func (m MsgClaimVoucher) Route() string { return routerKey }

func (m MsgClaimVoucher) Type() string { return TypeMsgClaimVoucher }

func (msg MsgClaimVoucher) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return err
	}

	if msg.Denom == "" {
		return fmt.Errorf("invalid denom")
	}
	return nil
}

func (m *MsgClaimVoucher) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgClaimVoucher) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

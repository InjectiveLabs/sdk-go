package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// constants
const (
	// RouterKey is the message route for slashing
	routerKey = ModuleName
)

var _ sdk.Msg = &MsgUpdateParams{}

func (m MsgUpdateParams) Route() string { return routerKey }

func (m MsgUpdateParams) Type() string { return "update_params" }

func (m MsgUpdateParams) ValidateBasic() error { return nil }

func (m *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgCreateNamespace{}

func (m MsgCreateNamespace) Route() string { return routerKey }

func (m MsgCreateNamespace) Type() string { return "create_namespace" }

func (m MsgCreateNamespace) ValidateBasic() error { return nil }

func (m *MsgCreateNamespace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgCreateNamespace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgDeleteNamespace{}

func (m MsgDeleteNamespace) Route() string { return routerKey }

func (m MsgDeleteNamespace) Type() string { return "delete_namespace" }

func (m MsgDeleteNamespace) ValidateBasic() error { return nil }

func (m *MsgDeleteNamespace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgDeleteNamespace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgUpdateNamespace{}

func (m MsgUpdateNamespace) Route() string { return routerKey }

func (m MsgUpdateNamespace) Type() string { return "update_namespace" }

func (m MsgUpdateNamespace) ValidateBasic() error { return nil }

func (m *MsgUpdateNamespace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgUpdateNamespace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgUpdateNamespaceRoles{}

func (m MsgUpdateNamespaceRoles) Route() string { return routerKey }

func (m MsgUpdateNamespaceRoles) Type() string { return "update_namespace_roles" }

func (m MsgUpdateNamespaceRoles) ValidateBasic() error { return nil }

func (m *MsgUpdateNamespaceRoles) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgUpdateNamespaceRoles) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgRevokeNamespaceRoles{}

func (m MsgRevokeNamespaceRoles) Route() string { return routerKey }

func (m MsgRevokeNamespaceRoles) Type() string { return "revoke_namespace_roles" }

func (m MsgRevokeNamespaceRoles) ValidateBasic() error { return nil }

func (m *MsgRevokeNamespaceRoles) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgRevokeNamespaceRoles) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

var _ sdk.Msg = &MsgClaimVoucher{}

func (m MsgClaimVoucher) Route() string { return routerKey }

func (m MsgClaimVoucher) Type() string { return "claim_voucher" }

func (m MsgClaimVoucher) ValidateBasic() error { return nil }

func (m *MsgClaimVoucher) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshal(m))
}

func (m MsgClaimVoucher) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{addr}
}

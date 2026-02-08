package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/ethereum/go-ethereum/common"
)

// constants
const (
	ProposalTypeGrantBandOraclePrivilege         string = "ProposalTypeGrantBandOraclePrivilege"
	ProposalTypeRevokeBandOraclePrivilege        string = "ProposalTypeRevokeBandOraclePrivilege"
	ProposalTypeGrantPriceFeederOraclePrivilege  string = "ProposalTypeGrantPriceFeederOraclePrivilege"
	ProposalTypeRevokePriceFeederOraclePrivilege string = "ProposalTypeRevokePriceFeederOraclePrivilege"
	ProposalAuthorizeBandOracleRequest           string = "ProposalTypeAuthorizeBandOracleRequest"
	ProposalUpdateBandOracleRequest              string = "ProposalUpdateBandOracleRequest"
	ProposalEnableBandIBC                        string = "ProposalTypeEnableBandIBC"
	ProposalTypeGrantProviderPrivilege           string = "ProposalTypeGrantProviderPrivilege"
	ProposalTypeRevokeProviderPrivilege          string = "ProposalTypeRevokeProviderPrivilege"
	ProposalTypeGrantStorkPublisherPrivilege     string = "ProposalTypeGrantStorkPublisherPrivilege"
	ProposalTypeRevokeStorkPublisherPrivilege    string = "ProposalTypeRevokeStorkPublisherPrivilege"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeGrantBandOraclePrivilege)
	govtypes.RegisterProposalType(ProposalTypeRevokeBandOraclePrivilege)
	govtypes.RegisterProposalType(ProposalTypeGrantPriceFeederOraclePrivilege)
	govtypes.RegisterProposalType(ProposalTypeRevokePriceFeederOraclePrivilege)
	govtypes.RegisterProposalType(ProposalAuthorizeBandOracleRequest)
	govtypes.RegisterProposalType(ProposalEnableBandIBC)
	govtypes.RegisterProposalType(ProposalUpdateBandOracleRequest)
	govtypes.RegisterProposalType(ProposalTypeGrantProviderPrivilege)
	govtypes.RegisterProposalType(ProposalTypeRevokeProviderPrivilege)
	govtypes.RegisterProposalType(ProposalTypeGrantStorkPublisherPrivilege)
	govtypes.RegisterProposalType(ProposalTypeRevokeStorkPublisherPrivilege)
}

// Implements Proposal Interface
var _ govtypes.Content = &GrantPriceFeederPrivilegeProposal{}
var _ govtypes.Content = &RevokePriceFeederPrivilegeProposal{}
var _ govtypes.Content = &GrantProviderPrivilegeProposal{}
var _ govtypes.Content = &RevokeProviderPrivilegeProposal{}
var _ govtypes.Content = &GrantStorkPublisherPrivilegeProposal{}
var _ govtypes.Content = &RevokeStorkPublisherPrivilegeProposal{}

// Deprecated: Band oracle proposal types kept for backward compatibility
var _ govtypes.Content = &GrantBandOraclePrivilegeProposal{}   //nolint:staticcheck // deprecated
var _ govtypes.Content = &RevokeBandOraclePrivilegeProposal{}  //nolint:staticcheck // deprecated
var _ govtypes.Content = &AuthorizeBandOracleRequestProposal{} //nolint:staticcheck // deprecated
var _ govtypes.Content = &UpdateBandOracleRequestProposal{}    //nolint:staticcheck // deprecated
var _ govtypes.Content = &EnableBandIBCProposal{}              //nolint:staticcheck // deprecated

// GetTitle returns the title of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) ProposalType() string {
	return ProposalTypeGrantPriceFeederOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) ValidateBasic() error {
	for _, relayer := range p.Relayers {
		r, err := sdk.AccAddressFromBech32(relayer)
		if err != nil {
			return err
		}
		if r.Empty() {
			return ErrEmptyRelayerAddr
		}
	}
	return govtypes.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) ProposalType() string {
	return ProposalTypeRevokePriceFeederOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) ValidateBasic() error {

	for _, relayer := range p.Relayers {
		r, err := sdk.AccAddressFromBech32(relayer)
		if err != nil {
			return err
		}
		if r.Empty() {
			return ErrEmptyRelayerAddr
		}
	}
	return govtypes.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *GrantProviderPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *GrantProviderPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *GrantProviderPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *GrantProviderPrivilegeProposal) ProposalType() string {
	return ProposalTypeGrantProviderPrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *GrantProviderPrivilegeProposal) ValidateBasic() error {

	if p.Provider == "" {
		return ErrEmptyProvider
	}

	for _, relayer := range p.Relayers {
		r, err := sdk.AccAddressFromBech32(relayer)
		if err != nil {
			return err
		}
		if r.Empty() {
			return ErrEmptyRelayerAddr
		}
	}
	return govtypes.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *RevokeProviderPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *RevokeProviderPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *RevokeProviderPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *RevokeProviderPrivilegeProposal) ProposalType() string {
	return ProposalTypeRevokeProviderPrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *RevokeProviderPrivilegeProposal) ValidateBasic() error {
	if p.Provider == "" {
		return ErrEmptyProvider
	}

	if strings.Contains(p.Provider, providerDelimiter) {
		return ErrInvalidProvider
	}

	if len(p.Relayers) == 0 {
		return ErrEmptyRelayerAddr
	}

	for _, relayer := range p.Relayers {
		r, err := sdk.AccAddressFromBech32(relayer)
		if err != nil {
			return err
		}
		if r.Empty() {
			return ErrEmptyRelayerAddr
		}
	}
	return govtypes.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *GrantStorkPublisherPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *GrantStorkPublisherPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *GrantStorkPublisherPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *GrantStorkPublisherPrivilegeProposal) ProposalType() string {
	return ProposalTypeGrantBandOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *GrantStorkPublisherPrivilegeProposal) ValidateBasic() error {
	for _, publisher := range p.StorkPublishers {
		if !common.IsHexAddress(publisher) {
			return fmt.Errorf("invalid publisher address: %s", publisher)
		}
	}

	return nil
}

// GetTitle returns the title of this proposal.
func (p *RevokeStorkPublisherPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *RevokeStorkPublisherPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *RevokeStorkPublisherPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *RevokeStorkPublisherPrivilegeProposal) ProposalType() string {
	return ProposalTypeGrantBandOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *RevokeStorkPublisherPrivilegeProposal) ValidateBasic() error {
	for _, publisher := range p.StorkPublishers {
		if !common.IsHexAddress(publisher) {
			return fmt.Errorf("invalid publisher address: %s", publisher)
		}
	}

	return nil
}

// Deprecated: Band oracle proposal types - kept for backward compatibility only

// GetTitle returns the title of this proposal.
func (p *GrantBandOraclePrivilegeProposal) GetTitle() string { return p.Title } //nolint:staticcheck // deprecated

// GetDescription returns the description of this proposal.
func (p *GrantBandOraclePrivilegeProposal) GetDescription() string { return p.Description } //nolint:staticcheck // deprecated

// ProposalRoute returns router key of this proposal.
func (*GrantBandOraclePrivilegeProposal) ProposalRoute() string { return RouterKey } //nolint:staticcheck // deprecated

// ProposalType returns proposal type of this proposal.
func (*GrantBandOraclePrivilegeProposal) ProposalType() string { //nolint:staticcheck // deprecated
	return ProposalTypeGrantBandOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (*GrantBandOraclePrivilegeProposal) ValidateBasic() error { return ErrBandOracleDeprecated } //nolint:staticcheck // deprecated

// GetTitle returns the title of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) GetTitle() string { return p.Title } //nolint:staticcheck // deprecated

// GetDescription returns the description of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) GetDescription() string { return p.Description } //nolint:staticcheck // deprecated

// ProposalRoute returns router key of this proposal.
func (*RevokeBandOraclePrivilegeProposal) ProposalRoute() string { return RouterKey } //nolint:staticcheck // deprecated

// ProposalType returns proposal type of this proposal.
func (*RevokeBandOraclePrivilegeProposal) ProposalType() string { //nolint:staticcheck // deprecated
	return ProposalTypeRevokeBandOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (*RevokeBandOraclePrivilegeProposal) ValidateBasic() error { return ErrBandOracleDeprecated } //nolint:staticcheck // deprecated

// GetTitle returns the title of this proposal.
func (p *AuthorizeBandOracleRequestProposal) GetTitle() string { return p.Title } //nolint:staticcheck // deprecated

// GetDescription returns the description of this proposal.
func (p *AuthorizeBandOracleRequestProposal) GetDescription() string { return p.Description } //nolint:staticcheck // deprecated

// ProposalRoute returns router key of this proposal.
func (*AuthorizeBandOracleRequestProposal) ProposalRoute() string { return RouterKey } //nolint:staticcheck // deprecated

// ProposalType returns proposal type of this proposal.
func (*AuthorizeBandOracleRequestProposal) ProposalType() string { //nolint:staticcheck // deprecated
	return ProposalAuthorizeBandOracleRequest
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (*AuthorizeBandOracleRequestProposal) ValidateBasic() error { return ErrBandOracleDeprecated } //nolint:staticcheck // deprecated

// GetTitle returns the title of this proposal.
func (p *UpdateBandOracleRequestProposal) GetTitle() string { return p.Title } //nolint:staticcheck // deprecated

// GetDescription returns the description of this proposal.
func (p *UpdateBandOracleRequestProposal) GetDescription() string { return p.Description } //nolint:staticcheck // deprecated

// ProposalRoute returns router key of this proposal.
func (*UpdateBandOracleRequestProposal) ProposalRoute() string { return RouterKey } //nolint:staticcheck // deprecated

// ProposalType returns proposal type of this proposal.
func (*UpdateBandOracleRequestProposal) ProposalType() string { //nolint:staticcheck // deprecated
	return ProposalUpdateBandOracleRequest
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (*UpdateBandOracleRequestProposal) ValidateBasic() error { return ErrBandOracleDeprecated } //nolint:staticcheck // deprecated

// GetTitle returns the title of this proposal.
func (p *EnableBandIBCProposal) GetTitle() string { return p.Title } //nolint:staticcheck // deprecated

// GetDescription returns the description of this proposal.
func (p *EnableBandIBCProposal) GetDescription() string { return p.Description } //nolint:staticcheck // deprecated

// ProposalRoute returns router key of this proposal.
func (*EnableBandIBCProposal) ProposalRoute() string { return RouterKey } //nolint:staticcheck // deprecated

// ProposalType returns proposal type of this proposal.
func (*EnableBandIBCProposal) ProposalType() string { //nolint:staticcheck // deprecated
	return ProposalEnableBandIBC
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (*EnableBandIBCProposal) ValidateBasic() error { return ErrBandOracleDeprecated } //nolint:staticcheck // deprecated

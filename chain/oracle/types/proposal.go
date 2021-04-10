package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// constants
const (
	ProposalTypeGrantBandOraclePrivilege         string = "ProposalTypeGrantBandOraclePrivilege"
	ProposalTypeRevokeBandOraclePrivilege        string = "ProposalTypeRevokeBandOraclePrivilege"
	ProposalTypeGrantPriceFeederOraclePrivilege  string = "ProposalTypeGrantPriceFeederOraclePrivilege"
	ProposalTypeRevokePriceFeederOraclePrivilege string = "ProposalTypeRevokePriceFeederOraclePrivilege"
)

func init() {
	gov.RegisterProposalType(ProposalTypeGrantBandOraclePrivilege)
	gov.RegisterProposalTypeCodec(&GrantBandOraclePrivilegeProposal{}, "injective/GrantBandOraclePrivilegeProposal")
	gov.RegisterProposalType(ProposalTypeRevokeBandOraclePrivilege)
	gov.RegisterProposalTypeCodec(&RevokeBandOraclePrivilegeProposal{}, "injective/RevokeBandOraclePrivilegeProposal")
	gov.RegisterProposalType(ProposalTypeGrantPriceFeederOraclePrivilege)
	gov.RegisterProposalTypeCodec(&GrantPriceFeederPrivilegeProposal{}, "injective/GrantPriceFeederPrivilegeProposal")
	gov.RegisterProposalType(ProposalTypeRevokePriceFeederOraclePrivilege)
	gov.RegisterProposalTypeCodec(&RevokePriceFeederPrivilegeProposal{}, "injective/RevokePriceFeederPrivilegeProposal")
}

// Implements Proposal Interface
var _ gov.Content = &GrantBandOraclePrivilegeProposal{}
var _ gov.Content = &RevokeBandOraclePrivilegeProposal{}
var _ gov.Content = &GrantPriceFeederPrivilegeProposal{}
var _ gov.Content = &RevokePriceFeederPrivilegeProposal{}

// GetTitle returns the title of this proposal.
func (p *GrantBandOraclePrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *GrantBandOraclePrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *GrantBandOraclePrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *GrantBandOraclePrivilegeProposal) ProposalType() string {
	return ProposalTypeGrantBandOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *GrantBandOraclePrivilegeProposal) ValidateBasic() error {
	for _, relayer := range p.Relayers {
		if _, err := sdk.AccAddressFromBech32(relayer); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) ProposalType() string {
	return ProposalTypeRevokeBandOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) ValidateBasic() error {
	for _, relayer := range p.Relayers {
		if _, err := sdk.AccAddressFromBech32(relayer); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)
}

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
		if _, err := sdk.AccAddressFromBech32(relayer); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)
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
		if _, err := sdk.AccAddressFromBech32(relayer); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)
}

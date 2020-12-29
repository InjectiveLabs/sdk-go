package types

import (
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// constants
const (
	ProposalTypeRegisterTokenMapping string = "RegisterTokenMapping"
)

// NewRegisterTokenMappingProposal returns new instance of TokenMappingProposal
func NewRegisterTokenMappingProposal(title, description string, mapping TokenMapping) gov.Content {
	return &RegisterTokenMappingProposal{title, description, mapping}
}

// Implements Proposal Interface
var _ gov.Content = &RegisterTokenMappingProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeRegisterTokenMapping)
	gov.RegisterProposalTypeCodec(&RegisterTokenMappingProposal{}, "cosmos-sdk/RegisterTokenMappingProposal")
}

// ProposalRoute returns router key for this proposal
func (sup *RegisterTokenMappingProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (sup *RegisterTokenMappingProposal) ProposalType() string {
	return ProposalTypeRegisterTokenMapping
}

// ValidateBasic returns ValidateBasic result for this proposal
func (sup *RegisterTokenMappingProposal) ValidateBasic() error {
	if err := sup.Mapping.ValidateBasic(); err != nil {
		return err
	}
	return gov.ValidateAbstract(sup)
}

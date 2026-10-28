package types

import (
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// constants
const (
	ProposalContractRegistrationRequest      string = "ProposalContractRegistrationRequest"
	ProposalBatchContractRegistrationRequest string = "ProposalBatchContractRegistrationRequest"
	ProposalBatchContractDeregistration      string = "ProposalBatchContractDeregistration"
	ProposalBatchStoreCode                   string = "ProposalBatchStoreCode"
)

func init() {
	govtypes.RegisterProposalType(ProposalContractRegistrationRequest)
	govtypes.RegisterProposalType(ProposalBatchContractRegistrationRequest)
	govtypes.RegisterProposalType(ProposalBatchContractDeregistration)
	govtypes.RegisterProposalType(ProposalBatchStoreCode)
}

// Implements Proposal Interface
var _ govtypes.Content = &ContractRegistrationRequestProposal{}
var _ govtypes.Content = &BatchContractRegistrationRequestProposal{}
var _ govtypes.Content = &BatchContractDeregistrationProposal{}
var _ govtypes.Content = &BatchStoreCodeProposal{}

// NewContractRegistrationRequestProposal returns new instance of ContractRegistrationRequestProposal
func NewContractRegistrationRequestProposal(title, description string, contractRegistrationRequest ContractRegistrationRequest) *ContractRegistrationRequestProposal {
	return &ContractRegistrationRequestProposal{
		Title:                       title,
		Description:                 description,
		ContractRegistrationRequest: contractRegistrationRequest,
	}
}

// GetTitle returns the title of this proposal.
func (p *ContractRegistrationRequestProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *ContractRegistrationRequestProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *ContractRegistrationRequestProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *ContractRegistrationRequestProposal) ProposalType() string {
	return ProposalContractRegistrationRequest
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *ContractRegistrationRequestProposal) ValidateBasic() error {
	return ErrModuleRemoved
}

// GetTitle returns the title of this proposal.
func (p *BatchContractRegistrationRequestProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *BatchContractRegistrationRequestProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *BatchContractRegistrationRequestProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *BatchContractRegistrationRequestProposal) ProposalType() string {
	return ProposalBatchContractRegistrationRequest
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *BatchContractRegistrationRequestProposal) ValidateBasic() error {
	return ErrModuleRemoved
}

// GetTitle returns the title of this proposal.
func (p *BatchContractDeregistrationProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *BatchContractDeregistrationProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *BatchContractDeregistrationProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *BatchContractDeregistrationProposal) ProposalType() string {
	return ProposalBatchContractDeregistration
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *BatchContractDeregistrationProposal) ValidateBasic() error {
	return ErrModuleRemoved
}

// GetTitle returns the title of this proposal.
func (p *BatchStoreCodeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *BatchStoreCodeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *BatchStoreCodeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *BatchStoreCodeProposal) ProposalType() string {
	return ProposalBatchStoreCode
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *BatchStoreCodeProposal) ValidateBasic() error {
	for idx := range p.Proposals {
		if err := p.Proposals[idx].ValidateBasic(); err != nil {
			return err
		}
	}
	return govtypes.ValidateAbstract(p)
}

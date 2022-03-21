package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// constants
const (
	ProposalContractRegistrationRequest      string = "ProposalContractRegistrationRequest"
	ProposalBatchContractRegistrationRequest string = "ProposalBatchContractRegistrationRequest"
)

func init() {
	gov.RegisterProposalType(ProposalContractRegistrationRequest)
	gov.RegisterProposalTypeCodec(&ContractRegistrationRequestProposal{}, "injective/ContractRegistrationRequestProposal")
	gov.RegisterProposalType(ProposalBatchContractRegistrationRequest)
	gov.RegisterProposalTypeCodec(&BatchContractRegistrationRequestProposal{}, "injective/BatchContractRegistrationRequestProposal")

}

// NewContractRegistrationRequestProposal returns new instance of ContractRegistrationRequestProposal
func NewContractRegistrationRequestProposal(title, description string, ContractRegistrationRequest ContractRegistrationRequest) *ContractRegistrationRequestProposal {
	return &ContractRegistrationRequestProposal{
		Title:                       title,
		Description:                 description,
		ContractRegistrationRequest: ContractRegistrationRequest,
	}
}

// Implements Proposal Interface
var _ gov.Content = &ContractRegistrationRequestProposal{}
var _ gov.Content = &BatchContractRegistrationRequestProposal{}

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
	// Check if contract address is valid
	if _, err := sdk.AccAddressFromBech32(p.ContractRegistrationRequest.ContractAddress); err != nil {
		return sdkerrors.Wrapf(ErrInvalidContractAddress, "ContractRegistrationRequestProposal: Error parsing registry contract address %s", err.Error())
	}

	return gov.ValidateAbstract(p)
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
	for _, req := range p.ContractRegistrationRequests {
		// Check if contract address is valid
		if _, err := sdk.AccAddressFromBech32(req.ContractAddress); err != nil {
			return sdkerrors.Wrapf(ErrInvalidContractAddress, "BatchContractRegistrationRequestProposal: Error parsing registry contract address %s", err.Error())
		}
	}

	return gov.ValidateAbstract(p)
}

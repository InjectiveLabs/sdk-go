package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// constants
const (
	ProposalContractRegistrationRequest string = "ProposalContractRegistrationRequest"
)

func init() {
	gov.RegisterProposalType(ProposalContractRegistrationRequest)
	gov.RegisterProposalTypeCodec(&ContractRegistrationRequestProposal{}, "injective/ContractRegistrationRequestProposal")

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
	// The gas price should be parsed to dec coin
	if _, err := sdk.ParseDecCoin(p.ContractRegistrationRequest.GasPrice); err != nil {
		return err
	}

	return gov.ValidateAbstract(p)
}

func SafeIsPositiveInt(v sdk.Int) bool {
	if v.IsNil() {
		return false
	}

	return v.IsPositive()
}

func SafeIsPositiveDec(v sdk.Dec) bool {
	if v.IsNil() {
		return false
	}

	return v.IsPositive()
}

func IsZeroOrNilInt(v sdk.Int) bool {
	return v.IsNil() || v.IsZero()
}

func IsZeroOrNilDec(v sdk.Dec) bool {
	return v.IsNil() || v.IsZero()
}

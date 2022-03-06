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
	// 	TODO
	// Add validation to check
	// - if code id exist
	// - if contract address is valid and exist
	// - if caller address is valid and exist
	// - if gasLimit is not zero
	// if p.ContractRegistrationRequest.GasLimit.IsZero() {
	// 	return fmt.Errorf("invalid parameter type: %T", i)
	// }

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

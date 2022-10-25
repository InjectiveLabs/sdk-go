package types

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// constants
const (
	ProposalContractRegistrationRequest      string = "ProposalContractRegistrationRequest"
	ProposalBatchContractRegistrationRequest string = "ProposalBatchContractRegistrationRequest"
	ProposalBatchContractDeregistration      string = "ProposalBatchContractDeregistration"
	ProposalBatchStoreCode                   string = "ProposalBatchStoreCode"
)

func init() {
	gov.RegisterProposalType(ProposalContractRegistrationRequest)
	gov.RegisterProposalTypeCodec(&ContractRegistrationRequestProposal{}, "injective/ContractRegistrationRequestProposal")
	gov.RegisterProposalType(ProposalBatchContractRegistrationRequest)
	gov.RegisterProposalTypeCodec(&BatchContractRegistrationRequestProposal{}, "injective/BatchContractRegistrationRequestProposal")
	gov.RegisterProposalType(ProposalBatchContractDeregistration)
	gov.RegisterProposalTypeCodec(&BatchContractDeregistrationProposal{}, "injective/BatchContractDeregistrationProposal")
	gov.RegisterProposalType(ProposalBatchStoreCode)
	gov.RegisterProposalTypeCodec(&BatchStoreCodeProposal{}, "injective/BatchStoreCodeProposal")

}

// Implements Proposal Interface
var _ gov.Content = &ContractRegistrationRequestProposal{}
var _ gov.Content = &BatchContractRegistrationRequestProposal{}
var _ gov.Content = &BatchContractDeregistrationProposal{}
var _ gov.Content = &BatchStoreCodeProposal{}

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

	if hasDuplicatesContractRegistrationRequest(p.ContractRegistrationRequests) {
		return sdkerrors.Wrapf(ErrDuplicateContract, "BatchContractRegistrationRequestProposal: Duplicate contract registration requests")
	}

	return gov.ValidateAbstract(p)
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
	found := make(map[string]struct{})

	if len(p.Contracts) == 0 {
		return sdkerrors.Wrapf(ErrNoContractAddresses, "BatchContractDeregistrationProposal: Contract list was empty")
	}

	for _, contract := range p.Contracts {
		// Check if contract address is valid
		addr, err := sdk.AccAddressFromBech32(contract)
		if err != nil {
			return sdkerrors.Wrapf(ErrInvalidContractAddress, "BatchContractDeregistrationProposal: Error parsing contract address %s", err.Error())
		}

		// Check that there are no duplicate contract addresses
		if _, ok := found[addr.String()]; ok {
			return sdkerrors.Wrapf(ErrDuplicateContract, "BatchContractDeregistrationProposal: Duplicate contract in contracts to deregister")
		} else {
			found[addr.String()] = struct{}{}
		}
	}

	return gov.ValidateAbstract(p)
}

// NewBatchStoreCodeProposal returns new instance of BatchStoreCodeProposal
func NewBatchStoreCodeProposal(title, description string, proposals []wasmtypes.StoreCodeProposal) *BatchStoreCodeProposal {
	return &BatchStoreCodeProposal{
		Title:       title,
		Description: description,
		Proposals:   proposals,
	}
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
	for _, proposal := range p.Proposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}

	return gov.ValidateAbstract(p)
}

func hasDuplicatesContractRegistrationRequest(slice []ContractRegistrationRequest) bool {
	seen := make(map[string]struct{})
	for _, item := range slice {
		addr := sdk.MustAccAddressFromBech32(item.ContractAddress).String()
		if _, ok := seen[addr]; ok {
			return true
		}
		seen[addr] = struct{}{}
	}
	return false
}

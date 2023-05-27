package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	if err := p.ContractRegistrationRequest.Validate(); err != nil {
		return err
	}

	return govtypes.ValidateAbstract(p)
}

func (req *ContractRegistrationRequest) Validate() error {
	// Check if contract address is valid
	if _, err := sdk.AccAddressFromBech32(req.ContractAddress); err != nil {
		return errors.Wrapf(ErrInvalidContractAddress, "ContractRegistrationRequestProposal: Error parsing registry contract address %s", err.Error())
	}

	if req.GranterAddress != "" {
		if _, err := sdk.AccAddressFromBech32(req.GranterAddress); err != nil {
			return errors.Wrapf(ErrInvalidContractAddress, "ContractRegistrationRequestProposal: Error parsing granter address %s", err.Error())
		}
	}

	if req.FundingMode == FundingMode_Unspecified {
		return errors.Wrapf(ErrInvalidFundingMode, "ContractRegistrationRequestProposal: FundingMode must be specified")
	}

	if (req.FundingMode == FundingMode_GrantOnly || req.FundingMode == FundingMode_Dual) && req.GranterAddress == "" {
		return errors.Wrapf(ErrInvalidFundingMode, "GranterAddress must be specified")
	}

	if req.FundingMode == FundingMode_SelfFunded && req.GranterAddress != "" {
		return errors.Wrapf(ErrInvalidFundingMode, "GranterAddress must be empty for self-funded contracts")
	}

	return nil
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
		if err := req.Validate(); err != nil {
			return err
		}
	}

	if hasDuplicatesContractRegistrationRequest(p.ContractRegistrationRequests) {
		return errors.Wrapf(ErrDuplicateContract, "BatchContractRegistrationRequestProposal: Duplicate contract registration requests")
	}

	return govtypes.ValidateAbstract(p)
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
	if len(p.Contracts) == 0 {
		return errors.Wrapf(ErrNoContractAddresses, "BatchContractDeregistrationProposal: Contract list was empty")
	}

	found := make(map[string]struct{})

	for _, contract := range p.Contracts {
		// Check if contract address is valid
		addr, err := sdk.AccAddressFromBech32(contract)
		if err != nil {
			return errors.Wrapf(ErrInvalidContractAddress, "BatchContractDeregistrationProposal: Error parsing contract address %s", err.Error())
		}

		// Check that there are no duplicate contract addresses
		if _, ok := found[addr.String()]; ok {
			return errors.Wrapf(ErrDuplicateContract, "BatchContractDeregistrationProposal: Duplicate contract in contracts to deregister")
		} else {
			found[addr.String()] = struct{}{}
		}
	}

	return govtypes.ValidateAbstract(p)
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

func HasDuplicates(slice []string) bool {
	seen := make(map[string]struct{})
	for _, item := range slice {
		if _, ok := seen[item]; ok {
			return true
		}
		seen[item] = struct{}{}
	}
	return false
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

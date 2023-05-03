package types

import (
	fmt "fmt"
	govcdc "github.com/cosmos/cosmos-sdk/x/gov/codec"

	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/ethereum/go-ethereum/common"
)

// constants
const (
	ProposalTypeRegisterTokenMapping string = "RegisterTokenMapping"
	ProposalTypeResetHub             string = "ResetHub"
)

func init() {
	gov.RegisterProposalType(ProposalTypeRegisterTokenMapping)
	gov.RegisterProposalType(ProposalTypeResetHub)
	govcdc.Amino.RegisterConcrete(&RegisterTokenMappingProposal{}, "cosmos-sdk/RegisterTokenMappingProposal", nil)
	govcdc.Amino.RegisterConcrete(&ResetHubProposal{}, "cosmos-sdk/ResetHubProposal", nil)
}

// NewTokenMapping returns an instance of TokenMapping
func NewTokenMapping(name string, erc20Address string, cosmosDenom string, enabled bool) TokenMapping {
	return TokenMapping{
		Name:         name,
		Erc20Address: erc20Address,
		CosmosDenom:  cosmosDenom,
		Enabled:      true,
	}
}

// NewRegisterTokenMappingProposal returns new instance of TokenMappingProposal
func NewRegisterTokenMappingProposal(title, description string, mapping TokenMapping) gov.Content {
	return &RegisterTokenMappingProposal{title, description, mapping}
}

// Implements Proposal Interface
var _ gov.Content = &RegisterTokenMappingProposal{}

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

// NewResetHubProposal returns new instance of ResetHubProposal
func NewResetHubProposal(title, description string, hub string) gov.Content {
	return &ResetHubProposal{title, description, hub}
}

// Implements Proposal Interface
var _ gov.Content = &ResetHubProposal{}

// ProposalRoute returns router key for this proposal
func (rh *ResetHubProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type for this proposal
func (rh *ResetHubProposal) ProposalType() string {
	return ProposalTypeResetHub
}

// ValidateBasic returns ValidateBasic result for this proposal
func (rh *ResetHubProposal) ValidateBasic() error {
	if !common.IsHexAddress(rh.HubAddress) {
		return fmt.Errorf("invalid hub address: %s", rh.HubAddress)
	}
	return gov.ValidateAbstract(rh)
}

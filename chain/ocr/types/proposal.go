package types

import (
	"cosmossdk.io/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// constants
const (
	ProposalTypeOcrSetConfig      string = "ProposalTypeOcrSetConfig"
	ProposalTypeOcrBatchSetConfig string = "ProposalTypeOcrBatchSetConfig"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeOcrSetConfig)
	govtypes.RegisterProposalType(ProposalTypeOcrBatchSetConfig)
}

// Implements Proposal Interface
var _ govtypes.Content = &SetConfigProposal{}
var _ govtypes.Content = &SetBatchConfigProposal{}

// GetTitle returns the title of this proposal.
func (p *SetConfigProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *SetConfigProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *SetConfigProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *SetConfigProposal) ProposalType() string {
	return ProposalTypeOcrSetConfig
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *SetConfigProposal) ValidateBasic() error {
	if p.Config == nil {
		return errors.Wrap(ErrIncompleteProposal, "proposal is missing config")
	}

	if err := p.Config.ValidateBasic(); err != nil {
		return err
	}

	return govtypes.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *SetBatchConfigProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *SetBatchConfigProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *SetBatchConfigProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *SetBatchConfigProposal) ProposalType() string {
	return ProposalTypeOcrBatchSetConfig
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *SetBatchConfigProposal) ValidateBasic() error {
	if len(p.FeedProperties) == 0 {
		return errors.Wrap(ErrIncompleteProposal, "proposal is missing feeds")
	}

	for _, feed := range p.FeedProperties {
		f := FeedConfig{
			Signers:               p.Signers,
			Transmitters:          p.Transmitters,
			F:                     feed.F,
			OnchainConfig:         feed.OnchainConfig,
			OffchainConfigVersion: feed.OffchainConfigVersion,
			OffchainConfig:        feed.OffchainConfig,
			ModuleParams: &ModuleParams{
				FeedId:              feed.FeedId,
				MinAnswer:           feed.MinAnswer,
				MaxAnswer:           feed.MaxAnswer,
				LinkPerObservation:  feed.LinkPerObservation,
				LinkPerTransmission: feed.LinkPerTransmission,
				LinkDenom:           p.LinkDenom,
				UniqueReports:       feed.UniqueReports,
				Description:         feed.Description,
			},
		}

		if err := f.ValidateBasic(); err != nil {
			return err
		}
	}
	return govtypes.ValidateAbstract(p)
}

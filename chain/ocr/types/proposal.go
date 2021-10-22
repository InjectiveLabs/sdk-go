package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// constants
const (
	ProposalTypeOcrSetConfig      string = "ProposalTypeOcrSetConfig"
	ProposalTypeOcrBatchSetConfig string = "ProposalTypeOcrBatchSetConfig"
)

func init() {
	gov.RegisterProposalType(ProposalTypeOcrSetConfig)
	gov.RegisterProposalTypeCodec(&SetConfigProposal{}, "injective/OcrSetConfigProposal")
	gov.RegisterProposalTypeCodec(&SetBatchConfigProposal{}, "injective/OcrSetBatchConfigProposal")
}

// Implements Proposal Interface
var _ gov.Content = &SetConfigProposal{}
var _ gov.Content = &SetBatchConfigProposal{}

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
		return sdkerrors.Wrap(ErrIncompleteProposal, "proposal is missing config")
	}

	if err := p.Config.ValidateBasic(); err != nil {
		return err
	}

	return gov.ValidateAbstract(p)
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
		return sdkerrors.Wrap(ErrIncompleteProposal, "proposal is missing feeds")
	}

	for _, feed := range p.FeedProperties {
		f := FeedConfig{
			Signers:      p.Signers,
			Transmitters: p.Transmitters,
			F:            feed.F,
			OnchainConfig: &OnchainConfig{
				FeedId:              feed.FeedId,
				MinAnswer:           feed.MinAnswer,
				MaxAnswer:           feed.MaxAnswer,
				LinkPerObservation:  feed.LinkPerObservation,
				LinkPerTransmission: feed.LinkPerTransmission,
				LinkDenom:           p.LinkDenom,
				UniqueReports:       feed.UniqueReports,
				Description:         feed.Description,
			},
			OffchainConfigVersion: feed.OffchainConfigVersion,
			OffchainConfig:        feed.OffchainConfig,
		}

		if err := f.ValidateBasic(); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)
}

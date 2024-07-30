package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeUpdateDenomsMetaData = "UpdateDenomsMetaData"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeUpdateDenomsMetaData)
}

var (
	_ govtypes.Content = &UpdateDenomsMetadataProposal{}
)

func NewUpdateDenomsMetadataProposal(title, description string, metadatas []banktypes.Metadata) govtypes.Content {
	return &UpdateDenomsMetadataProposal{
		Title:       title,
		Description: description,
		Metadatas:   metadatas,
	}
}

func (p *UpdateDenomsMetadataProposal) GetTitle() string { return p.Title }

func (p *UpdateDenomsMetadataProposal) GetDescription() string { return p.Description }

func (p *UpdateDenomsMetadataProposal) ProposalRoute() string { return RouterKey }

func (p *UpdateDenomsMetadataProposal) ProposalType() string {
	return ProposalTypeUpdateDenomsMetaData
}

func (p *UpdateDenomsMetadataProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(p)
	if err != nil {
		return err
	}

	for idx := range p.Metadatas {
		metadata := p.Metadatas[idx]
		if metadata.Base == "" {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denom (%s)", metadata.Base)
		}
	}

	return nil
}

package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"

	bandprice "github.com/InjectiveLabs/sdk-go/bandchain/hooks/price"
	bandoracle "github.com/InjectiveLabs/sdk-go/bandchain/oracle/types"
	bandobi "github.com/bandprotocol/bandchain-packet/obi"
)

// constants
const (
	ProposalTypeGrantBandOraclePrivilege         string = "ProposalTypeGrantBandOraclePrivilege"
	ProposalTypeRevokeBandOraclePrivilege        string = "ProposalTypeRevokeBandOraclePrivilege"
	ProposalTypeGrantPriceFeederOraclePrivilege  string = "ProposalTypeGrantPriceFeederOraclePrivilege"
	ProposalTypeRevokePriceFeederOraclePrivilege string = "ProposalTypeRevokePriceFeederOraclePrivilege"
	ProposalAuthorizeBandOracleRequest           string = "ProposalTypeAuthorizeBandOracleRequest"
	ProposalUpdateBandOracleRequest              string = "ProposalUpdateBandOracleRequest"
	ProposalEnableBandIBC                        string = "ProposalTypeEnableBandIBC"
	ProposalTypeGrantProviderPrivilege           string = "ProposalTypeGrantProviderPrivilege"
	ProposalTypeRevokeProviderPrivilege          string = "ProposalTypeRevokeProviderPrivilege"
)

func init() {
	gov.RegisterProposalType(ProposalTypeGrantBandOraclePrivilege)
	gov.RegisterProposalTypeCodec(&GrantBandOraclePrivilegeProposal{}, "injective/GrantBandOraclePrivilegeProposal")
	gov.RegisterProposalType(ProposalTypeRevokeBandOraclePrivilege)
	gov.RegisterProposalTypeCodec(&RevokeBandOraclePrivilegeProposal{}, "injective/RevokeBandOraclePrivilegeProposal")
	gov.RegisterProposalType(ProposalTypeGrantPriceFeederOraclePrivilege)
	gov.RegisterProposalTypeCodec(&GrantPriceFeederPrivilegeProposal{}, "injective/GrantPriceFeederPrivilegeProposal")
	gov.RegisterProposalType(ProposalTypeRevokePriceFeederOraclePrivilege)
	gov.RegisterProposalTypeCodec(&RevokePriceFeederPrivilegeProposal{}, "injective/RevokePriceFeederPrivilegeProposal")
	gov.RegisterProposalType(ProposalAuthorizeBandOracleRequest)
	gov.RegisterProposalTypeCodec(&AuthorizeBandOracleRequestProposal{}, "injective/AuthorizeBandOracleRequestProposal")
	gov.RegisterProposalType(ProposalEnableBandIBC)
	gov.RegisterProposalTypeCodec(&EnableBandIBCProposal{}, "injective/EnableBandIBCProposal")
	gov.RegisterProposalType(ProposalUpdateBandOracleRequest)
	gov.RegisterProposalTypeCodec(&UpdateBandOracleRequestProposal{}, "injective/UpdateBandOracleRequestProposal")
	gov.RegisterProposalType(ProposalTypeGrantProviderPrivilege)
	gov.RegisterProposalTypeCodec(&GrantProviderPrivilegeProposal{}, "injective/GrantProviderPrivilegeProposal")
	gov.RegisterProposalType(ProposalTypeRevokeProviderPrivilege)
	gov.RegisterProposalTypeCodec(&RevokeProviderPrivilegeProposal{}, "injective/RevokeProviderPrivilegeProposal")
}

// Implements Proposal Interface
var _ gov.Content = &GrantBandOraclePrivilegeProposal{}
var _ gov.Content = &RevokeBandOraclePrivilegeProposal{}
var _ gov.Content = &GrantPriceFeederPrivilegeProposal{}
var _ gov.Content = &RevokePriceFeederPrivilegeProposal{}
var _ gov.Content = &AuthorizeBandOracleRequestProposal{}
var _ gov.Content = &UpdateBandOracleRequestProposal{}
var _ gov.Content = &EnableBandIBCProposal{}
var _ gov.Content = &GrantProviderPrivilegeProposal{}
var _ gov.Content = &RevokeProviderPrivilegeProposal{}

// GetTitle returns the title of this proposal.
func (p *GrantBandOraclePrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *GrantBandOraclePrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *GrantBandOraclePrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *GrantBandOraclePrivilegeProposal) ProposalType() string {
	return ProposalTypeGrantBandOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *GrantBandOraclePrivilegeProposal) ValidateBasic() error {
	for _, relayer := range p.Relayers {
		if _, err := sdk.AccAddressFromBech32(relayer); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) ProposalType() string {
	return ProposalTypeRevokeBandOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *RevokeBandOraclePrivilegeProposal) ValidateBasic() error {
	for _, relayer := range p.Relayers {
		if _, err := sdk.AccAddressFromBech32(relayer); err != nil {
			return err
		}
	}
	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) ProposalType() string {
	return ProposalTypeGrantPriceFeederOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *GrantPriceFeederPrivilegeProposal) ValidateBasic() error {
	for _, relayer := range p.Relayers {
		r, err := sdk.AccAddressFromBech32(relayer)
		if err != nil {
			return err
		}
		if r.Empty() {
			return ErrEmptyRelayerAddr
		}
	}
	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) ProposalType() string {
	return ProposalTypeRevokePriceFeederOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *RevokePriceFeederPrivilegeProposal) ValidateBasic() error {

	for _, relayer := range p.Relayers {
		r, err := sdk.AccAddressFromBech32(relayer)
		if err != nil {
			return err
		}
		if r.Empty() {
			return ErrEmptyRelayerAddr
		}
	}
	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *GrantProviderPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *GrantProviderPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *GrantProviderPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *GrantProviderPrivilegeProposal) ProposalType() string {
	return ProposalTypeGrantProviderPrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *GrantProviderPrivilegeProposal) ValidateBasic() error {

	if p.Provider == "" {
		return ErrEmptyProvider
	}

	for _, relayer := range p.Relayers {
		r, err := sdk.AccAddressFromBech32(relayer)
		if err != nil {
			return err
		}
		if r.Empty() {
			return ErrEmptyRelayerAddr
		}
	}
	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *RevokeProviderPrivilegeProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *RevokeProviderPrivilegeProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *RevokeProviderPrivilegeProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *RevokeProviderPrivilegeProposal) ProposalType() string {
	return ProposalTypeRevokeProviderPrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *RevokeProviderPrivilegeProposal) ValidateBasic() error {
	if p.Provider == "" {
		return ErrEmptyProvider
	}

	if strings.Contains(p.Provider, providerDelimiter) {
		return ErrInvalidProvider
	}

	if len(p.Relayers) == 0 {
		return ErrEmptyRelayerAddr
	}

	for _, relayer := range p.Relayers {
		r, err := sdk.AccAddressFromBech32(relayer)
		if err != nil {
			return err
		}
		if r.Empty() {
			return ErrEmptyRelayerAddr
		}
	}
	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *AuthorizeBandOracleRequestProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *AuthorizeBandOracleRequestProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *AuthorizeBandOracleRequestProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *AuthorizeBandOracleRequestProposal) ProposalType() string {
	return ProposalTypeRevokePriceFeederOraclePrivilege
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *AuthorizeBandOracleRequestProposal) ValidateBasic() error {
	if p.Request.OracleScriptId <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBandIBCRequest, "AuthorizeBandOracleRequestProposal: Oracle script id (%d) must be positive.", p.Request.OracleScriptId)
	}

	if len(p.Request.Symbols) == 0 {
		return sdkerrors.Wrap(ErrBadSymbolsCount, "AuthorizeBandOracleRequestProposal")
	}

	requestCallData := bandprice.Input{
		Symbols:    p.Request.Symbols,
		Multiplier: BandPriceMultiplier,
	}
	callData := bandobi.MustEncode(requestCallData)

	if len(callData) > bandoracle.MaxDataSize {
		return bandoracle.WrapMaxError(bandoracle.ErrTooLargeCalldata, len(callData), bandoracle.MaxDataSize)
	}

	if p.Request.MinCount <= 0 {
		return sdkerrors.Wrapf(bandoracle.ErrInvalidMinCount, "AuthorizeBandOracleRequestProposal: Minimum validator count (%d) must be positive.", p.Request.MinCount)
	}

	if p.Request.AskCount <= 0 {
		return sdkerrors.Wrapf(bandoracle.ErrInvalidAskCount, "AuthorizeBandOracleRequestProposal: Request validator count (%d) must be positive.", p.Request.AskCount)
	}

	if p.Request.AskCount < p.Request.MinCount {
		return sdkerrors.Wrapf(bandoracle.ErrInvalidAskCount, "AuthorizeBandOracleRequestProposal: Request validator count (%d) must not be less than sufficient validator count (%d).", p.Request.AskCount, p.Request.MinCount)
	}

	if !p.Request.FeeLimit.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "AuthorizeBandOracleRequestProposal: Invalid Fee Limit (%s)", p.Request.GetFeeLimit().String())
	}

	if p.Request.PrepareGas <= 0 {
		return sdkerrors.Wrapf(bandoracle.ErrInvalidOwasmGas, "AuthorizeBandOracleRequestProposal: Invalid Prepare Gas (%d)", p.Request.GetPrepareGas())
	}

	if p.Request.ExecuteGas <= 0 {
		return sdkerrors.Wrapf(bandoracle.ErrInvalidOwasmGas, "AuthorizeBandOracleRequestProposal: Invalid Execute Gas (%d)", p.Request.ExecuteGas)
	}

	if p.Request.PrepareGas+p.Request.ExecuteGas > bandoracle.MaximumOwasmGas {
		return sdkerrors.Wrapf(bandoracle.ErrInvalidOwasmGas, "AuthorizeBandOracleRequestProposal: Sum of prepare gas and execute gas (%d) exceeds Maximum gas (%d)", (p.Request.PrepareGas + p.Request.ExecuteGas), bandoracle.MaximumOwasmGas)
	}

	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *EnableBandIBCProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *EnableBandIBCProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *EnableBandIBCProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *EnableBandIBCProposal) ProposalType() string {
	return ProposalEnableBandIBC
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *EnableBandIBCProposal) ValidateBasic() error {

	if p.BandIbcParams.IbcRequestInterval == 0 {
		return ErrBadRequestInterval
	}

	if p.BandIbcParams.IbcSourceChannel == "" {
		return sdkerrors.Wrap(ErrInvalidBandIBCRequest, "AuthorizeBandOracleRequestProposal: IBC Source Chanel must not be empty.")
	}
	if p.BandIbcParams.IbcVersion == "" {
		return sdkerrors.Wrap(bandoracle.ErrInvalidVersion, "AuthorizeBandOracleRequestProposal: IBC Version must not be empty.")
	}

	return gov.ValidateAbstract(p)
}

// GetTitle returns the title of this proposal.
func (p *UpdateBandOracleRequestProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *UpdateBandOracleRequestProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *UpdateBandOracleRequestProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *UpdateBandOracleRequestProposal) ProposalType() string {
	return ProposalUpdateBandOracleRequest
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *UpdateBandOracleRequestProposal) ValidateBasic() error {

	if p.DeleteRequestId == 0 && p.UpdateOracleRequest == nil {
		return ErrInvalidBandIBCUpdateRequest
	}

	if p.DeleteRequestId > 0 && p.UpdateOracleRequest != nil {
		return sdkerrors.Wrapf(ErrInvalidBandIBCUpdateRequest, "cannot update requestID %T and delete reqeustID %T at same time", p.UpdateOracleRequest.RequestId, p.DeleteRequestId)
	}

	if p.UpdateOracleRequest != nil && len(p.UpdateOracleRequest.Symbols) > 0 {
		requestCallData := bandprice.Input{
			Symbols:    p.UpdateOracleRequest.Symbols,
			Multiplier: BandPriceMultiplier,
		}
		callData := bandobi.MustEncode(requestCallData)

		if len(callData) > bandoracle.MaxDataSize {
			return bandoracle.WrapMaxError(bandoracle.ErrTooLargeCalldata, len(callData), bandoracle.MaxDataSize)
		}
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.AskCount > 0 && p.UpdateOracleRequest.MinCount > 0 && p.UpdateOracleRequest.AskCount < p.UpdateOracleRequest.MinCount {
		return sdkerrors.Wrapf(bandoracle.ErrInvalidAskCount, "UpdateBandOracleRequestProposal: Request validator count (%d) must not be less than sufficient validator count (%d).", p.UpdateOracleRequest.AskCount, p.UpdateOracleRequest.MinCount)
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.FeeLimit != nil && !p.UpdateOracleRequest.FeeLimit.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "UpdateBandOracleRequestProposal: Invalid Fee Limit (%s)", p.UpdateOracleRequest.GetFeeLimit().String())
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.PrepareGas > 0 && p.UpdateOracleRequest.ExecuteGas > 0 && p.UpdateOracleRequest.PrepareGas+p.UpdateOracleRequest.ExecuteGas > bandoracle.MaximumOwasmGas {
		return sdkerrors.Wrapf(bandoracle.ErrInvalidOwasmGas, "UpdateBandOracleRequestProposal: Sum of prepare gas and execute gas (%d) exceeds Maximum gas (%d)", (p.UpdateOracleRequest.PrepareGas + p.UpdateOracleRequest.ExecuteGas), bandoracle.MaximumOwasmGas)
	}

	return gov.ValidateAbstract(p)
}

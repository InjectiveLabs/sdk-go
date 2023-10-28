package types

import (
	"strings"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	bandobi "github.com/bandprotocol/bandchain-packet/obi"

	bandprice "github.com/InjectiveLabs/injective-core/injective-chain/modules/oracle/bandchain/hooks/price"
	bandoracle "github.com/InjectiveLabs/injective-core/injective-chain/modules/oracle/bandchain/oracle/types"
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
	govtypes.RegisterProposalType(ProposalTypeGrantBandOraclePrivilege)
	govtypes.RegisterProposalType(ProposalTypeRevokeBandOraclePrivilege)
	govtypes.RegisterProposalType(ProposalTypeGrantPriceFeederOraclePrivilege)
	govtypes.RegisterProposalType(ProposalTypeRevokePriceFeederOraclePrivilege)
	govtypes.RegisterProposalType(ProposalAuthorizeBandOracleRequest)
	govtypes.RegisterProposalType(ProposalEnableBandIBC)
	govtypes.RegisterProposalType(ProposalUpdateBandOracleRequest)
	govtypes.RegisterProposalType(ProposalTypeGrantProviderPrivilege)
	govtypes.RegisterProposalType(ProposalTypeRevokeProviderPrivilege)
}

// Implements Proposal Interface
var _ govtypes.Content = &GrantBandOraclePrivilegeProposal{}
var _ govtypes.Content = &RevokeBandOraclePrivilegeProposal{}
var _ govtypes.Content = &GrantPriceFeederPrivilegeProposal{}
var _ govtypes.Content = &RevokePriceFeederPrivilegeProposal{}
var _ govtypes.Content = &AuthorizeBandOracleRequestProposal{}
var _ govtypes.Content = &UpdateBandOracleRequestProposal{}
var _ govtypes.Content = &EnableBandIBCProposal{}
var _ govtypes.Content = &GrantProviderPrivilegeProposal{}
var _ govtypes.Content = &RevokeProviderPrivilegeProposal{}

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
	return govtypes.ValidateAbstract(p)
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
	return govtypes.ValidateAbstract(p)
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
	return govtypes.ValidateAbstract(p)
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
	return govtypes.ValidateAbstract(p)
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
	return govtypes.ValidateAbstract(p)
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
	return govtypes.ValidateAbstract(p)
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
		return errors.Wrapf(ErrInvalidBandIBCRequest, "AuthorizeBandOracleRequestProposal: Oracle script id (%d) must be positive.", p.Request.OracleScriptId)
	}

	if len(p.Request.Symbols) == 0 {
		return errors.Wrap(ErrBadSymbolsCount, "AuthorizeBandOracleRequestProposal")
	}

	callData, err := bandobi.Encode(bandprice.SymbolInput{
		Symbols:            p.Request.Symbols,
		MinimumSourceCount: uint8(p.Request.MinCount),
	})
	if err != nil {
		return err
	}

	if len(callData) > bandoracle.MaxDataSize {
		return bandoracle.WrapMaxError(bandoracle.ErrTooLargeCalldata, len(callData), bandoracle.MaxDataSize)
	}

	if p.Request.MinCount <= 0 {
		return errors.Wrapf(bandoracle.ErrInvalidMinCount, "AuthorizeBandOracleRequestProposal: Minimum validator count (%d) must be positive.", p.Request.MinCount)
	}

	if p.Request.AskCount <= 0 {
		return errors.Wrapf(bandoracle.ErrInvalidAskCount, "AuthorizeBandOracleRequestProposal: Request validator count (%d) must be positive.", p.Request.AskCount)
	}

	if p.Request.AskCount < p.Request.MinCount {
		return errors.Wrapf(bandoracle.ErrInvalidAskCount, "AuthorizeBandOracleRequestProposal: Request validator count (%d) must not be less than sufficient validator count (%d).", p.Request.AskCount, p.Request.MinCount)
	}

	if !p.Request.FeeLimit.IsValid() {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "AuthorizeBandOracleRequestProposal: Invalid Fee Limit (%s)", p.Request.GetFeeLimit().String())
	}

	if p.Request.PrepareGas <= 0 {
		return errors.Wrapf(bandoracle.ErrInvalidOwasmGas, "AuthorizeBandOracleRequestProposal: Invalid Prepare Gas (%d)", p.Request.GetPrepareGas())
	}

	if p.Request.ExecuteGas <= 0 {
		return errors.Wrapf(bandoracle.ErrInvalidOwasmGas, "AuthorizeBandOracleRequestProposal: Invalid Execute Gas (%d)", p.Request.ExecuteGas)
	}

	if p.Request.PrepareGas+p.Request.ExecuteGas > bandoracle.MaximumOwasmGas {
		return errors.Wrapf(bandoracle.ErrInvalidOwasmGas, "AuthorizeBandOracleRequestProposal: Sum of prepare gas and execute gas (%d) exceeds Maximum gas (%d)", (p.Request.PrepareGas + p.Request.ExecuteGas), bandoracle.MaximumOwasmGas)
	}

	return govtypes.ValidateAbstract(p)
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
		return errors.Wrap(ErrInvalidBandIBCRequest, "AuthorizeBandOracleRequestProposal: IBC Source Chanel must not be empty.")
	}
	if p.BandIbcParams.IbcVersion == "" {
		return errors.Wrap(bandoracle.ErrInvalidVersion, "AuthorizeBandOracleRequestProposal: IBC Version must not be empty.")
	}

	return govtypes.ValidateAbstract(p)
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
	if len(p.DeleteRequestIds) == 0 && p.UpdateOracleRequest == nil {
		return ErrInvalidBandIBCUpdateRequest
	}

	if len(p.DeleteRequestIds) > 0 && p.UpdateOracleRequest != nil {
		return errors.Wrapf(ErrInvalidBandIBCUpdateRequest, "cannot update requestID %T and delete reqeustID %T at same time", p.UpdateOracleRequest.RequestId, p.DeleteRequestIds)
	}

	if p.UpdateOracleRequest != nil && len(p.UpdateOracleRequest.Symbols) > 0 {
		callData, err := bandobi.Encode(bandprice.SymbolInput{
			Symbols:            p.UpdateOracleRequest.Symbols,
			MinimumSourceCount: uint8(p.UpdateOracleRequest.MinCount),
		})

		if err != nil {
			return err
		}

		if len(callData) > bandoracle.MaxDataSize {
			return bandoracle.WrapMaxError(bandoracle.ErrTooLargeCalldata, len(callData), bandoracle.MaxDataSize)
		}
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.AskCount > 0 && p.UpdateOracleRequest.MinCount > 0 && p.UpdateOracleRequest.AskCount < p.UpdateOracleRequest.MinCount {
		return errors.Wrapf(bandoracle.ErrInvalidAskCount, "UpdateBandOracleRequestProposal: Request validator count (%d) must not be less than sufficient validator count (%d).", p.UpdateOracleRequest.AskCount, p.UpdateOracleRequest.MinCount)
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.FeeLimit != nil && !p.UpdateOracleRequest.FeeLimit.IsValid() {
		return errors.Wrapf(sdkerrors.ErrInvalidCoins, "UpdateBandOracleRequestProposal: Invalid Fee Limit (%s)", p.UpdateOracleRequest.GetFeeLimit().String())
	}

	if p.UpdateOracleRequest != nil && p.UpdateOracleRequest.PrepareGas > 0 && p.UpdateOracleRequest.ExecuteGas > 0 && p.UpdateOracleRequest.PrepareGas+p.UpdateOracleRequest.ExecuteGas > bandoracle.MaximumOwasmGas {
		return errors.Wrapf(bandoracle.ErrInvalidOwasmGas, "UpdateBandOracleRequestProposal: Sum of prepare gas and execute gas (%d) exceeds Maximum gas (%d)", (p.UpdateOracleRequest.PrepareGas + p.UpdateOracleRequest.ExecuteGas), bandoracle.MaximumOwasmGas)
	}

	return govtypes.ValidateAbstract(p)
}

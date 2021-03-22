package types

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common"
)

// constants
const (
	ProposalTypeSpotMarketParamUpdate       string = "ProposalTypeSpotMarketParamUpdate"
	ProposalTypeSpotMarketLaunch            string = "ProposalTypeSpotMarketLaunch"
	ProposalTypeSetSpotMarketStatus         string = "ProposalTypeSetSpotMarketStatus"
	ProposalTypePerpetualMarketLaunch       string = "ProposalTypePerpetualMarketLaunch"
	ProposalTypeExpiryFuturesMarketLaunch   string = "ProposalTypeExpiryFuturesMarketLaunch"
	ProposalTypeDerivativeMarketParamUpdate string = "ProposalTypeDerivativeMarketParamUpdate"
	ProposalTypeSetDerivativeMarketStatus   string = "ProposalTypeSetDerivativeMarketStatus"
)

func init() {
	gov.RegisterProposalType(ProposalTypeSpotMarketParamUpdate)
	gov.RegisterProposalTypeCodec(&SpotMarketParamUpdateProposal{}, "injective/SpotMarketParamUpdateProposal")
	gov.RegisterProposalType(ProposalTypeSpotMarketLaunch)
	gov.RegisterProposalTypeCodec(&SpotMarketLaunchProposal{}, "injective/SpotMarketLaunchProposal")
	gov.RegisterProposalType(ProposalTypeSetSpotMarketStatus)
	gov.RegisterProposalTypeCodec(&SpotMarketStatusSetProposal{}, "injective/SpotMarketStatusSetProposal")
	gov.RegisterProposalType(ProposalTypePerpetualMarketLaunch)
	gov.RegisterProposalTypeCodec(&PerpetualMarketLaunchProposal{}, "injective/PerpetualMarketLaunchProposal")
	gov.RegisterProposalType(ProposalTypeExpiryFuturesMarketLaunch)
	gov.RegisterProposalTypeCodec(&ExpiryFuturesMarketLaunchProposal{}, "injective/ExpiryFuturesMarketLaunchProposal")
	gov.RegisterProposalType(ProposalTypeDerivativeMarketParamUpdate)
	gov.RegisterProposalTypeCodec(&DerivativeMarketParamUpdateProposal{}, "injective/DerivativeMarketParamUpdateProposal")
	gov.RegisterProposalType(ProposalTypeSetDerivativeMarketStatus)
	gov.RegisterProposalTypeCodec(&DerivativeMarketStatusSetProposal{}, "injective/DerivativeMarketStatusSetProposal")
}

// NewSpotMarketParamUpdateProposal returns new instance of SpotMarketParamUpdateProposal
func NewSpotMarketParamUpdateProposal(title, description string, marketID common.Hash, makerFeeRate, takerFeeRate, relayerFeeShareRate sdk.Dec, maxPriceScaleDecimals, maxQuantityScaleDecimals uint32) *SpotMarketParamUpdateProposal {
	return &SpotMarketParamUpdateProposal{
		title, description, marketID.Hex(), makerFeeRate, takerFeeRate, relayerFeeShareRate, maxPriceScaleDecimals, maxQuantityScaleDecimals,
	}
}

// Implements Proposal Interface
var _ gov.Content = &SpotMarketParamUpdateProposal{}

// GetTitle returns the title of this proposal.
func (p *SpotMarketParamUpdateProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *SpotMarketParamUpdateProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *SpotMarketParamUpdateProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *SpotMarketParamUpdateProposal) ProposalType() string {
	return ProposalTypeSpotMarketLaunch
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *SpotMarketParamUpdateProposal) ValidateBasic() error {

	if err := ValidateFee(p.MakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(p.TakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(p.RelayerFeeShareRate); err != nil {
		return err
	}
	if p.MakerFeeRate.GT(p.TakerFeeRate) {
		return errors.New("MakerFeeRate cannot be greater than TakerFeeRate")
	}
	return gov.ValidateAbstract(p)
}

// NewSpotMarketLaunchProposal returns new instance of SpotMarketLaunchProposal
func NewSpotMarketLaunchProposal(title, description, ticker, baseDenom, quoteDenom string, maxPriceScaleDecimals, maxQuantityScaleDecimals uint32) *SpotMarketLaunchProposal {
	return &SpotMarketLaunchProposal{
		title, description, ticker, baseDenom, quoteDenom, maxPriceScaleDecimals, maxQuantityScaleDecimals,
	}
}

// Implements Proposal Interface
var _ gov.Content = &SpotMarketLaunchProposal{}

// GetTitle returns the title of this proposal.
func (p *SpotMarketLaunchProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *SpotMarketLaunchProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *SpotMarketLaunchProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *SpotMarketLaunchProposal) ProposalType() string {
	return ProposalTypeSpotMarketLaunch
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *SpotMarketLaunchProposal) ValidateBasic() error {
	if err := types.ValidateDenom(p.BaseDenom); err != nil {
		return err
	}
	if err := types.ValidateDenom(p.QuoteDenom); err != nil {
		return err
	}
	return gov.ValidateAbstract(p)
}

// NewSpotMarketStatusSetProposal returns new instance of SpotMarketStatusSetProposal
func NewSpotMarketStatusSetProposal(title, description, baseDenom, quoteDenom string, status MarketStatus) *SpotMarketStatusSetProposal {
	return &SpotMarketStatusSetProposal{
		title, description, baseDenom, quoteDenom, status,
	}
}

// Implements Proposal Interface
var _ gov.Content = &SpotMarketStatusSetProposal{}

// GetTitle returns the title of this proposal.
func (p *SpotMarketStatusSetProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *SpotMarketStatusSetProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *SpotMarketStatusSetProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *SpotMarketStatusSetProposal) ProposalType() string {
	return ProposalTypeSetSpotMarketStatus
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *SpotMarketStatusSetProposal) ValidateBasic() error {
	if err := types.ValidateDenom(p.BaseDenom); err != nil {
		return err
	}
	if err := types.ValidateDenom(p.QuoteDenom); err != nil {
		return err
	}
	if p.Status.String() == "" {
		return errors.New("Invalid status")
	}
	return gov.ValidateAbstract(p)
}

// NewDerivativeMarketStatusSetProposal returns new instance of DerivativeMarketStatusSetProposal
func NewDerivativeMarketStatusSetProposal(title, description, marketID string, status MarketStatus) *DerivativeMarketStatusSetProposal {
	return &DerivativeMarketStatusSetProposal{
		title, description, marketID, status,
	}
}

// Implements Proposal Interface
var _ gov.Content = &DerivativeMarketStatusSetProposal{}

// GetTitle returns the title of this proposal.
func (p *DerivativeMarketStatusSetProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *DerivativeMarketStatusSetProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *DerivativeMarketStatusSetProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *DerivativeMarketStatusSetProposal) ProposalType() string {
	return ProposalTypeSetDerivativeMarketStatus
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *DerivativeMarketStatusSetProposal) ValidateBasic() error {
	return gov.ValidateAbstract(p)
}

// NewDerivativeMarketParamUpdateProposal returns new instance of DerivativeMarketParamUpdateProposal
func NewDerivativeMarketParamUpdateProposal(
	title, description string, marketID string,
	initialMarginRatio, maintenanceMarginRatio,
	makerFeeRate, takerFeeRate, relayerFeeShareRate *sdk.Dec,
	maxPriceScaleDecimals, maxQuantityScaleDecimals uint32,
) *DerivativeMarketParamUpdateProposal {
	return &DerivativeMarketParamUpdateProposal{
		title,
		description,
		marketID,
		initialMarginRatio,
		maintenanceMarginRatio,
		makerFeeRate,
		takerFeeRate,
		relayerFeeShareRate,
		maxPriceScaleDecimals,
		maxQuantityScaleDecimals,
	}
}

// Implements Proposal Interface
var _ gov.Content = &DerivativeMarketParamUpdateProposal{}

// GetTitle returns the title of this proposal
func (p *DerivativeMarketParamUpdateProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *DerivativeMarketParamUpdateProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *DerivativeMarketParamUpdateProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *DerivativeMarketParamUpdateProposal) ProposalType() string {
	return ProposalTypeDerivativeMarketParamUpdate
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *DerivativeMarketParamUpdateProposal) ValidateBasic() error {
	if p.MakerFeeRate != nil {
		if err := ValidateFee(*p.MakerFeeRate); err != nil {
			return err
		}
	}
	if p.TakerFeeRate != nil {
		if err := ValidateFee(*p.TakerFeeRate); err != nil {
			return err
		}
	}

	if p.RelayerFeeShareRate != nil {
		if err := ValidateFee(*p.RelayerFeeShareRate); err != nil {
			return err
		}
	}

	if p.InitialMarginRatio != nil {
		if err := ValidateMarginRatio(*p.InitialMarginRatio); err != nil {
			return err
		}
	}
	if p.MaintenanceMarginRatio != nil {
		if err := ValidateMarginRatio(*p.MaintenanceMarginRatio); err != nil {
			return err
		}
	}

	if p.MakerFeeRate != nil && p.TakerFeeRate != nil {
		if p.MakerFeeRate.GT(*p.TakerFeeRate) {
			return errors.New("MakerFeeRate cannot be greater than TakerFeeRate")
		}
	}

	if p.MaintenanceMarginRatio != nil && p.InitialMarginRatio != nil {
		if p.InitialMarginRatio.LT(*p.MaintenanceMarginRatio) {
			return errors.New("MaintenanceMarginRatio cannot be greater than InitialMarginRatio")
		}
	}

	return gov.ValidateAbstract(p)
}

// NewPerpetualMarketLaunchProposal returns new instance of PerpetualMarketLaunchProposal
func NewPerpetualMarketLaunchProposal(title, description, ticker, quoteDenom, oracle string, maxPriceScaleDecimals, maxQuantityScaleDecimals uint32) *PerpetualMarketLaunchProposal {
	return &PerpetualMarketLaunchProposal{
		title, description, ticker, quoteDenom, oracle, maxPriceScaleDecimals, maxQuantityScaleDecimals,
	}
}

// Implements Proposal Interface
var _ gov.Content = &PerpetualMarketLaunchProposal{}

// GetTitle returns the title of this proposal.
func (p *PerpetualMarketLaunchProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *PerpetualMarketLaunchProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *PerpetualMarketLaunchProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *PerpetualMarketLaunchProposal) ProposalType() string {
	return ProposalTypePerpetualMarketLaunch
}

// ValidateBasic returns ValidateBasic result of a perpetual market launch proposal.
func (p *PerpetualMarketLaunchProposal) ValidateBasic() error {
	if err := types.ValidateDenom(p.QuoteDenom); err != nil {
		return err
	}
	return gov.ValidateAbstract(p)
}

// NewExpiryFuturesMarketLaunchProposal returns new instance of ExpiryFuturesMarketLaunchProposal
func NewExpiryFuturesMarketLaunchProposal(title, description, ticker, quoteDenom, oracle string, expiry int64, maxPriceScaleDecimals, maxQuantityScaleDecimals uint32) *ExpiryFuturesMarketLaunchProposal {
	return &ExpiryFuturesMarketLaunchProposal{
		title, description, ticker, quoteDenom, oracle, expiry, maxPriceScaleDecimals, maxQuantityScaleDecimals,
	}
}

// Implements Proposal Interface
var _ gov.Content = &ExpiryFuturesMarketLaunchProposal{}

// GetTitle returns the title of this proposal.
func (p *ExpiryFuturesMarketLaunchProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *ExpiryFuturesMarketLaunchProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *ExpiryFuturesMarketLaunchProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *ExpiryFuturesMarketLaunchProposal) ProposalType() string {
	return ProposalTypeExpiryFuturesMarketLaunch
}

// ValidateBasic returns ValidateBasic result of a perpetual market launch proposal.
func (p *ExpiryFuturesMarketLaunchProposal) ValidateBasic() error {
	if err := types.ValidateDenom(p.QuoteDenom); err != nil {
		return err
	}
	return gov.ValidateAbstract(p)
}

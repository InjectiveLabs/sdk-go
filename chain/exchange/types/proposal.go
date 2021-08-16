package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common"
)

// constants
const (
	ProposalTypeExchangeEnable                      string = "ProposalTypeExchangeEnable"
	ProposalTypeSpotMarketParamUpdate               string = "ProposalTypeSpotMarketParamUpdate"
	ProposalTypeSpotMarketLaunch                    string = "ProposalTypeSpotMarketLaunch"
	ProposalTypePerpetualMarketLaunch               string = "ProposalTypePerpetualMarketLaunch"
	ProposalTypeExpiryFuturesMarketLaunch           string = "ProposalTypeExpiryFuturesMarketLaunch"
	ProposalTypeDerivativeMarketParamUpdate         string = "ProposalTypeDerivativeMarketParamUpdate"
	ProposalTypeDerivativeMarketBandOraclePromotion string = "ProposalTypeDerivativeMarketBandOraclePromotion"
)

func init() {
	gov.RegisterProposalType(ProposalTypeExchangeEnable)
	gov.RegisterProposalTypeCodec(&ExchangeEnableProposal{}, "injective/ExchangeEnableProposal")
	gov.RegisterProposalType(ProposalTypeSpotMarketParamUpdate)
	gov.RegisterProposalTypeCodec(&SpotMarketParamUpdateProposal{}, "injective/SpotMarketParamUpdateProposal")
	gov.RegisterProposalType(ProposalTypeSpotMarketLaunch)
	gov.RegisterProposalTypeCodec(&SpotMarketLaunchProposal{}, "injective/SpotMarketLaunchProposal")
	gov.RegisterProposalType(ProposalTypePerpetualMarketLaunch)
	gov.RegisterProposalTypeCodec(&PerpetualMarketLaunchProposal{}, "injective/PerpetualMarketLaunchProposal")
	gov.RegisterProposalType(ProposalTypeExpiryFuturesMarketLaunch)
	gov.RegisterProposalTypeCodec(&ExpiryFuturesMarketLaunchProposal{}, "injective/ExpiryFuturesMarketLaunchProposal")
	gov.RegisterProposalType(ProposalTypeDerivativeMarketParamUpdate)
	gov.RegisterProposalTypeCodec(&DerivativeMarketParamUpdateProposal{}, "injective/DerivativeMarketParamUpdateProposal")
	gov.RegisterProposalType(ProposalTypeDerivativeMarketBandOraclePromotion)
	gov.RegisterProposalTypeCodec(&DerivativeMarketBandOraclePromotionProposal{}, "injective/DerivativeMarketBandOraclePromotionProposal")
}

// Implements Proposal Interface
var _ gov.Content = &ExchangeEnableProposal{}

// GetTitle returns the title of this proposal.
func (p *ExchangeEnableProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *ExchangeEnableProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *ExchangeEnableProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *ExchangeEnableProposal) ProposalType() string {
	return ProposalTypeExchangeEnable
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *ExchangeEnableProposal) ValidateBasic() error {

	switch p.ExchangeType {
	case ExchangeType_SPOT, ExchangeType_DERIVATIVES:
	default:
		return ErrBadField
	}
	return gov.ValidateAbstract(p)
}

// NewSpotMarketParamUpdateProposal returns new instance of SpotMarketParamUpdateProposal
func NewSpotMarketParamUpdateProposal(title, description string, marketID common.Hash, makerFeeRate, takerFeeRate, relayerFeeShareRate, minPriceTickSize, minQuantityTickSize *sdk.Dec, status MarketStatus) *SpotMarketParamUpdateProposal {

	return &SpotMarketParamUpdateProposal{
		title,
		description,
		marketID.Hex(),
		makerFeeRate,
		takerFeeRate,
		relayerFeeShareRate,
		minPriceTickSize,
		minQuantityTickSize,
		status,
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
	return ProposalTypeSpotMarketParamUpdate
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *SpotMarketParamUpdateProposal) ValidateBasic() error {
	if p.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, p.MarketId)
	}
	if p.MakerFeeRate == nil && p.TakerFeeRate == nil && p.RelayerFeeShareRate == nil && p.MinPriceTickSize == nil && p.MinQuantityTickSize == nil && p.Status == MarketStatus_Unspecified {
		return sdkerrors.Wrap(gov.ErrInvalidProposalContent, "At least one field should not be nil")
	}

	if p.MakerFeeRate != nil {
		if err := ValidateMakerFee(*p.MakerFeeRate); err != nil {
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

	if p.MinPriceTickSize != nil {
		if err := ValidateTickSize(*p.MinPriceTickSize); err != nil {
			return sdkerrors.Wrap(ErrInvalidPriceTickSize, err.Error())
		}
	}
	if p.MinQuantityTickSize != nil {
		if err := ValidateTickSize(*p.MinQuantityTickSize); err != nil {
			return sdkerrors.Wrap(ErrInvalidQuantityTickSize, err.Error())
		}
	}

	switch p.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Active,
		MarketStatus_Paused,
		MarketStatus_Suspended,
		MarketStatus_Demolished,
		MarketStatus_Expired:

	default:
		return sdkerrors.Wrap(ErrInvalidMarketStatus, p.Status.String())
	}

	return gov.ValidateAbstract(p)
}

// NewSpotMarketLaunchProposal returns new instance of SpotMarketLaunchProposal
func NewSpotMarketLaunchProposal(title, description, ticker, baseDenom, quoteDenom string, minPriceTickSize, minQuantityTickSize sdk.Dec) *SpotMarketLaunchProposal {
	return &SpotMarketLaunchProposal{
		title, description, ticker, baseDenom, quoteDenom, minPriceTickSize, minQuantityTickSize,
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
	if p.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty")
	}
	if p.BaseDenom == "" {
		return sdkerrors.Wrap(ErrInvalidBaseDenom, "base denom should not be empty")
	}
	if p.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if p.BaseDenom == p.QuoteDenom {
		return ErrSameDenoms
	}

	if err := ValidateTickSize(p.MinPriceTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}
	return gov.ValidateAbstract(p)
}

// NewDerivativeMarketParamUpdateProposal returns new instance of DerivativeMarketParamUpdateProposal
func NewDerivativeMarketParamUpdateProposal(
	title, description string, marketID string,
	initialMarginRatio, maintenanceMarginRatio,
	makerFeeRate, takerFeeRate, relayerFeeShareRate, minPriceTickSize, minQuantityTickSize *sdk.Dec,
	status MarketStatus,
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
		minPriceTickSize,
		minQuantityTickSize,
		status,
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
	if p.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, p.MarketId)
	}
	if p.MakerFeeRate == nil &&
		p.TakerFeeRate == nil &&
		p.RelayerFeeShareRate == nil &&
		p.MinPriceTickSize == nil &&
		p.MinQuantityTickSize == nil &&
		p.InitialMarginRatio == nil &&
		p.MaintenanceMarginRatio == nil &&
		p.Status == MarketStatus_Unspecified {
		return sdkerrors.Wrap(gov.ErrInvalidProposalContent, "At least one field should not be nil")
	}

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

	if p.MinPriceTickSize != nil {
		if err := ValidateTickSize(*p.MinPriceTickSize); err != nil {
			return sdkerrors.Wrap(ErrInvalidPriceTickSize, err.Error())
		}
	}
	if p.MinQuantityTickSize != nil {
		if err := ValidateTickSize(*p.MinQuantityTickSize); err != nil {
			return sdkerrors.Wrap(ErrInvalidQuantityTickSize, err.Error())
		}
	}

	switch p.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Active,
		MarketStatus_Paused,
		MarketStatus_Suspended,
		MarketStatus_Demolished,
		MarketStatus_Expired:

	default:
		return sdkerrors.Wrap(ErrInvalidMarketStatus, p.Status.String())
	}

	return gov.ValidateAbstract(p)
}

// NewPerpetualMarketLaunchProposal returns new instance of PerpetualMarketLaunchProposal
func NewPerpetualMarketLaunchProposal(
	title, description, ticker, quoteDenom,
	oracleBase, oracleQuote string, oracleScaleFactor uint32, oracleType oracletypes.OracleType,
	initialMarginRatio, maintenanceMarginRatio, makerFeeRate, takerFeeRate, minPriceTickSize, minQuantityTickSize sdk.Dec,
) *PerpetualMarketLaunchProposal {
	return &PerpetualMarketLaunchProposal{
		Title:                  title,
		Description:            description,
		Ticker:                 ticker,
		QuoteDenom:             quoteDenom,
		OracleBase:             oracleBase,
		OracleQuote:            oracleQuote,
		OracleScaleFactor:      oracleScaleFactor,
		OracleType:             oracleType,
		InitialMarginRatio:     initialMarginRatio,
		MaintenanceMarginRatio: maintenanceMarginRatio,
		MakerFeeRate:           makerFeeRate,
		TakerFeeRate:           takerFeeRate,
		MinPriceTickSize:       minPriceTickSize,
		MinQuantityTickSize:    minQuantityTickSize,
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
	if p.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty")
	}
	if p.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if p.OracleBase == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle base should not be empty")
	}
	if p.OracleQuote == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle quote should not be empty")
	}
	if p.OracleBase == p.OracleQuote {
		return ErrSameOracles
	}
	switch p.OracleType {
	case oracletypes.OracleType_Band, oracletypes.OracleType_PriceFeed, oracletypes.OracleType_Coinbase, oracletypes.OracleType_Chainlink, oracletypes.OracleType_Razor,
		oracletypes.OracleType_Dia, oracletypes.OracleType_API3, oracletypes.OracleType_Uma, oracletypes.OracleType_Pyth, oracletypes.OracleType_BandIBC:

	default:
		return sdkerrors.Wrap(ErrInvalidOracleType, p.OracleType.String())
	}
	if p.OracleScaleFactor > MaxOracleScaleFactor {
		return ErrExceedsMaxOracleScaleFactor
	}
	if err := ValidateFee(p.MakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(p.TakerFeeRate); err != nil {
		return err
	}
	if err := ValidateMarginRatio(p.InitialMarginRatio); err != nil {
		return err
	}
	if err := ValidateMarginRatio(p.MaintenanceMarginRatio); err != nil {
		return err
	}
	if p.MakerFeeRate.GT(p.TakerFeeRate) {
		return ErrFeeRatesRelation
	}
	if p.InitialMarginRatio.LT(p.MaintenanceMarginRatio) {
		return ErrMarginsRelation
	}

	if err := ValidateTickSize(p.MinPriceTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}

	return gov.ValidateAbstract(p)
}

// NewExpiryFuturesMarketLaunchProposal returns new instance of ExpiryFuturesMarketLaunchProposal
func NewExpiryFuturesMarketLaunchProposal(
	title, description, ticker, quoteDenom,
	oracleBase, oracleQuote string, oracleScaleFactor uint32, oracleType oracletypes.OracleType, expiry int64,
	initialMarginRatio, maintenanceMarginRatio, makerFeeRate, takerFeeRate, minPriceTickSize, minQuantityTickSize sdk.Dec,
) *ExpiryFuturesMarketLaunchProposal {
	return &ExpiryFuturesMarketLaunchProposal{
		Title:                  title,
		Description:            description,
		Ticker:                 ticker,
		QuoteDenom:             quoteDenom,
		OracleBase:             oracleBase,
		OracleQuote:            oracleQuote,
		OracleScaleFactor:      oracleScaleFactor,
		OracleType:             oracleType,
		Expiry:                 expiry,
		InitialMarginRatio:     initialMarginRatio,
		MaintenanceMarginRatio: maintenanceMarginRatio,
		MakerFeeRate:           makerFeeRate,
		TakerFeeRate:           takerFeeRate,
		MinPriceTickSize:       minPriceTickSize,
		MinQuantityTickSize:    minQuantityTickSize,
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
	if p.Ticker == "" {
		return sdkerrors.Wrap(ErrInvalidTicker, "ticker should not be empty")
	}
	if p.QuoteDenom == "" {
		return sdkerrors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if p.OracleBase == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle base should not be empty")
	}
	if p.OracleQuote == "" {
		return sdkerrors.Wrap(ErrInvalidOracle, "oracle quote should not be empty")
	}
	if p.OracleBase == p.OracleQuote {
		return ErrSameOracles
	}
	switch p.OracleType {
	case oracletypes.OracleType_Band, oracletypes.OracleType_PriceFeed, oracletypes.OracleType_Coinbase, oracletypes.OracleType_Chainlink, oracletypes.OracleType_Razor,
		oracletypes.OracleType_Dia, oracletypes.OracleType_API3, oracletypes.OracleType_Uma, oracletypes.OracleType_Pyth, oracletypes.OracleType_BandIBC:

	default:
		return sdkerrors.Wrap(ErrInvalidOracleType, p.OracleType.String())
	}
	if p.OracleScaleFactor > MaxOracleScaleFactor {
		return ErrExceedsMaxOracleScaleFactor
	}
	if p.Expiry <= 0 {
		return sdkerrors.Wrap(ErrInvalidExpiry, "expiry should not be empty")
	}
	if err := ValidateFee(p.MakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(p.TakerFeeRate); err != nil {
		return err
	}
	if err := ValidateMarginRatio(p.InitialMarginRatio); err != nil {
		return err
	}
	if err := ValidateMarginRatio(p.MaintenanceMarginRatio); err != nil {
		return err
	}
	if p.MakerFeeRate.GT(p.TakerFeeRate) {
		return ErrFeeRatesRelation
	}
	if p.InitialMarginRatio.LT(p.MaintenanceMarginRatio) {
		return ErrMarginsRelation
	}

	if err := ValidateTickSize(p.MinPriceTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return sdkerrors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}

	return gov.ValidateAbstract(p)
}

// NewDerivativeMarketBandOraclePromotionProposal returns new instance of DerivativeMarketBandOraclePromotionProposal
func NewDerivativeMarketBandOraclePromotionProposal(title, description string, marketID string) *DerivativeMarketBandOraclePromotionProposal {
	return &DerivativeMarketBandOraclePromotionProposal{
		title,
		description,
		marketID,
	}
}

// Implements Proposal Interface
var _ gov.Content = &DerivativeMarketBandOraclePromotionProposal{}

// GetTitle returns the title of this proposal
func (p *DerivativeMarketBandOraclePromotionProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *DerivativeMarketBandOraclePromotionProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *DerivativeMarketBandOraclePromotionProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *DerivativeMarketBandOraclePromotionProposal) ProposalType() string {
	return ProposalTypeDerivativeMarketBandOraclePromotion
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *DerivativeMarketBandOraclePromotionProposal) ValidateBasic() error {
	if p.MarketId == "" {
		return sdkerrors.Wrap(ErrMarketInvalid, p.MarketId)
	}

	return gov.ValidateAbstract(p)
}

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
	ProposalTypeExchangeEnable              string = "ProposalTypeExchangeEnable"
	ProposalTypeBatchExchangeModification   string = "ProposalTypeBatchExchangeModification"
	ProposalTypeSpotMarketParamUpdate       string = "ProposalTypeSpotMarketParamUpdate"
	ProposalTypeSpotMarketLaunch            string = "ProposalTypeSpotMarketLaunch"
	ProposalTypePerpetualMarketLaunch       string = "ProposalTypePerpetualMarketLaunch"
	ProposalTypeExpiryFuturesMarketLaunch   string = "ProposalTypeExpiryFuturesMarketLaunch"
	ProposalTypeDerivativeMarketParamUpdate string = "ProposalTypeDerivativeMarketParamUpdate"
	ProposalTypeTradingRewardCampaign       string = "ProposalTypeTradingRewardCampaign"
	ProposalTypeTradingRewardCampaignUpdate string = "ProposalTypeTradingRewardCampaignUpdateProposal"
	ProposalTypeFeeDiscountProposal         string = "ProposalTypeFeeDiscountProposal"
)

func init() {
	gov.RegisterProposalType(ProposalTypeExchangeEnable)
	gov.RegisterProposalTypeCodec(&ExchangeEnableProposal{}, "injective/ExchangeEnableProposal")
	gov.RegisterProposalType(ProposalTypeBatchExchangeModification)
	gov.RegisterProposalTypeCodec(&BatchExchangeModificationProposal{}, "injective/BatchExchangeModificationProposal")
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
	gov.RegisterProposalType(ProposalTypeTradingRewardCampaign)
	gov.RegisterProposalTypeCodec(&TradingRewardCampaignLaunchProposal{}, "injective/TradingRewardCampaignLaunchProposal")
	gov.RegisterProposalType(ProposalTypeTradingRewardCampaignUpdate)
	gov.RegisterProposalTypeCodec(&TradingRewardCampaignUpdateProposal{}, "injective/TradingRewardCampaignUpdateProposal")
	gov.RegisterProposalType(ProposalTypeFeeDiscountProposal)
	gov.RegisterProposalTypeCodec(&FeeDiscountProposal{}, "injective/FeeDiscountProposal")
}

func SafeIsPositiveInt(v sdk.Int) bool {
	if v.IsNil() {
		return false
	}

	return v.IsPositive()
}

func SafeIsPositiveDec(v sdk.Dec) bool {
	if v.IsNil() {
		return false
	}

	return v.IsPositive()
}

func IsZeroOrNilInt(v sdk.Int) bool {
	return v.IsNil() || v.IsZero()
}

func IsZeroOrNilDec(v sdk.Dec) bool {
	return v.IsNil() || v.IsZero()
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

// Implements Proposal Interface
var _ gov.Content = &BatchExchangeModificationProposal{}

// GetTitle returns the title of this proposal.
func (p *BatchExchangeModificationProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *BatchExchangeModificationProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *BatchExchangeModificationProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *BatchExchangeModificationProposal) ProposalType() string {
	return ProposalTypeBatchExchangeModification
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *BatchExchangeModificationProposal) ValidateBasic() error {
	for _, proposal := range p.SpotMarketParamUpdateProposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}

	for _, proposal := range p.DerivativeMarketParamUpdateProposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}

	for _, proposal := range p.SpotMarketLaunchProposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}

	for _, proposal := range p.PerpetualMarketLaunchProposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}

	for _, proposal := range p.ExpiryFuturesMarketLaunchProposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}

	if p.TradingRewardCampaignUpdateProposal != nil {
		if err := p.TradingRewardCampaignUpdateProposal.ValidateBasic(); err != nil {
			return err
		}
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
func NewSpotMarketLaunchProposal(
	title string,
	description string,
	ticker string,
	baseDenom string,
	quoteDenom string,
	minPriceTickSize sdk.Dec,
	minQuantityTickSize sdk.Dec,
	makerFeeRate *sdk.Dec,
	takerFeeRate *sdk.Dec,
) *SpotMarketLaunchProposal {
	return &SpotMarketLaunchProposal{
		Title:               title,
		Description:         description,
		Ticker:              ticker,
		BaseDenom:           baseDenom,
		QuoteDenom:          quoteDenom,
		MinPriceTickSize:    minPriceTickSize,
		MinQuantityTickSize: minQuantityTickSize,
		MakerFeeRate:        makerFeeRate,
		TakerFeeRate:        takerFeeRate,
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

	if (p.MakerFeeRate == nil && p.TakerFeeRate != nil) || (p.MakerFeeRate != nil && p.TakerFeeRate == nil) {
		return sdkerrors.Wrap(ErrFeeRatesRelation, "maker and taker fee rate must either be both nil or both specified")
	}

	if p.MakerFeeRate != nil && p.TakerFeeRate != nil {
		if p.MakerFeeRate.GT(*p.TakerFeeRate) {
			return ErrFeeRatesRelation
		}
	}

	return gov.ValidateAbstract(p)
}

// NewDerivativeMarketParamUpdateProposal returns new instance of DerivativeMarketParamUpdateProposal
func NewDerivativeMarketParamUpdateProposal(
	title, description string, marketID string,
	initialMarginRatio, maintenanceMarginRatio,
	makerFeeRate, takerFeeRate, relayerFeeShareRate, minPriceTickSize, minQuantityTickSize *sdk.Dec,
	hourlyInterestRate, hourlyFundingRateCap *sdk.Dec,
	status MarketStatus, oracleParams *OracleParams,
) *DerivativeMarketParamUpdateProposal {
	return &DerivativeMarketParamUpdateProposal{
		Title:                  title,
		Description:            description,
		MarketId:               marketID,
		InitialMarginRatio:     initialMarginRatio,
		MaintenanceMarginRatio: maintenanceMarginRatio,
		MakerFeeRate:           makerFeeRate,
		TakerFeeRate:           takerFeeRate,
		RelayerFeeShareRate:    relayerFeeShareRate,
		MinPriceTickSize:       minPriceTickSize,
		MinQuantityTickSize:    minQuantityTickSize,
		HourlyInterestRate:     hourlyInterestRate,
		HourlyFundingRateCap:   hourlyFundingRateCap,
		Status:                 status,
		OracleParams:           oracleParams,
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
		p.HourlyInterestRate == nil &&
		p.HourlyFundingRateCap == nil &&
		p.Status == MarketStatus_Unspecified &&
		p.OracleParams == nil {
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

	if p.HourlyInterestRate != nil {
		if err := ValidateHourlyInterestRate(*p.HourlyInterestRate); err != nil {
			return sdkerrors.Wrap(ErrInvalidHourlyInterestRate, err.Error())
		}
	}

	if p.HourlyFundingRateCap != nil {
		if err := ValidateHourlyFundingRateCap(*p.HourlyFundingRateCap); err != nil {
			return sdkerrors.Wrap(ErrInvalidHourlyFundingRateCap, err.Error())
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

	if p.OracleParams != nil {
		if err := p.OracleParams.ValidateBasic(); err != nil {
			return err
		}
	}

	return gov.ValidateAbstract(p)
}

func NewOracleParams(
	oracleBase string,
	oracleQuote string,
	oracleScaleFactor uint32,
	oracleType oracletypes.OracleType,
) *OracleParams {
	return &OracleParams{
		OracleBase:        oracleBase,
		OracleQuote:       oracleQuote,
		OracleScaleFactor: oracleScaleFactor,
		OracleType:        oracleType,
	}
}

func (p *OracleParams) ValidateBasic() error {
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

	return nil
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

	oracleParams := NewOracleParams(p.OracleBase, p.OracleQuote, p.OracleScaleFactor, p.OracleType)
	if err := oracleParams.ValidateBasic(); err != nil {
		return err
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

	oracleParams := NewOracleParams(p.OracleBase, p.OracleQuote, p.OracleScaleFactor, p.OracleType)
	if err := oracleParams.ValidateBasic(); err != nil {
		return err
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

// NewTradingRewardCampaignUpdateProposal returns new instance of TradingRewardCampaignLaunchProposal
func NewTradingRewardCampaignUpdateProposal(
	title, description string,
	campaignInfo *TradingRewardCampaignInfo,
	rewardPoolsAdditions []*CampaignRewardPool,
	rewardPoolsUpdates []*CampaignRewardPool,
) *TradingRewardCampaignUpdateProposal {
	p := &TradingRewardCampaignUpdateProposal{
		Title:                        title,
		Description:                  description,
		CampaignInfo:                 campaignInfo,
		CampaignRewardPoolsAdditions: rewardPoolsAdditions,
		CampaignRewardPoolsUpdates:   rewardPoolsUpdates,
	}
	if err := p.ValidateBasic(); err != nil {
		panic(err)
	}
	return p
}

// Implements Proposal Interface
var _ gov.Content = &TradingRewardCampaignUpdateProposal{}

// GetTitle returns the title of this proposal
func (p *TradingRewardCampaignUpdateProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *TradingRewardCampaignUpdateProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *TradingRewardCampaignUpdateProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *TradingRewardCampaignUpdateProposal) ProposalType() string {
	return ProposalTypeTradingRewardCampaign
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *TradingRewardCampaignUpdateProposal) ValidateBasic() error {
	var err error

	if err := p.CampaignInfo.ValidateBasic(); err != nil {
		return err
	}

	prevStartTimestamp := int64(0)
	for _, pool := range p.CampaignRewardPoolsAdditions {
		if pool == nil {
			return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "campaign reward pool addition cannot be nil")
		}

		prevStartTimestamp, err = validateCampaignRewardPool(pool, p.CampaignInfo.CampaignDurationSeconds, prevStartTimestamp)
		if err != nil {
			return err
		}
	}

	prevStartTimestamp = int64(0)
	for _, pool := range p.CampaignRewardPoolsUpdates {
		prevStartTimestamp, err = validateCampaignRewardPool(pool, p.CampaignInfo.CampaignDurationSeconds, prevStartTimestamp)
		if err != nil {
			return err
		}
	}

	return gov.ValidateAbstract(p)
}

// NewTradingRewardCampaignLaunchProposal returns new instance of TradingRewardCampaignLaunchProposal
func NewTradingRewardCampaignLaunchProposal(
	title, description string,
	campaignInfo *TradingRewardCampaignInfo,
	campaignRewardPools []*CampaignRewardPool,
) *TradingRewardCampaignLaunchProposal {
	p := &TradingRewardCampaignLaunchProposal{
		Title:               title,
		Description:         description,
		CampaignInfo:        campaignInfo,
		CampaignRewardPools: campaignRewardPools,
	}
	if err := p.ValidateBasic(); err != nil {
		panic(err)
	}
	return p
}

// Implements Proposal Interface
var _ gov.Content = &TradingRewardCampaignLaunchProposal{}

// GetTitle returns the title of this proposal
func (p *TradingRewardCampaignLaunchProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *TradingRewardCampaignLaunchProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *TradingRewardCampaignLaunchProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *TradingRewardCampaignLaunchProposal) ProposalType() string {
	return ProposalTypeTradingRewardCampaign
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *TradingRewardCampaignLaunchProposal) ValidateBasic() error {
	var err error

	if p.CampaignInfo == nil {
		return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "new campaign info cannot be nil")
	}

	if p.CampaignInfo.CampaignDurationSeconds <= 0 {
		return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "new campaign duration cannot be zero")
	}

	if len(p.CampaignRewardPools) == 0 {
		return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "new campaign reward pools cannot be nil")
	}

	if err = p.CampaignInfo.ValidateBasic(); err != nil {
		return err
	}

	prevStartTimestamp := int64(0)
	for _, pool := range p.CampaignRewardPools {
		prevStartTimestamp, err = validateCampaignRewardPool(pool, p.CampaignInfo.CampaignDurationSeconds, prevStartTimestamp)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *TradingRewardCampaignBoostInfo) ValidateBasic() error {
	if len(t.BoostedSpotMarketIds) != len(t.SpotMarketMultipliers) {
		return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "boosted spot market ids is not matching spot market multipliers")
	}

	if len(t.BoostedDerivativeMarketIds) != len(t.DerivativeMarketMultipliers) {
		return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "boosted derivative market ids is not matching derivative market multipliers")
	}

	hasDuplicatesInMarkets := HasDuplicates(t.BoostedSpotMarketIds) || HasDuplicates(t.BoostedDerivativeMarketIds)
	if hasDuplicatesInMarkets {
		return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "campaign contains duplicate boosted market ids")
	}

	for _, multiplier := range t.SpotMarketMultipliers {
		if IsZeroOrNilDec(multiplier.MakerPointsMultiplier) {
			return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "spot market maker multiplier cannot be zero or nil")
		}

		if IsZeroOrNilDec(multiplier.TakerPointsMultiplier) {
			return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "spot market maker multiplier cannot be zero or nil")
		}

		if !SafeIsPositiveDec(multiplier.TakerPointsMultiplier) {
			return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "spot market maker multiplier cannot be negative")
		}
	}

	for _, multiplier := range t.DerivativeMarketMultipliers {
		if IsZeroOrNilDec(multiplier.MakerPointsMultiplier) {
			return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "derivative market maker multiplier cannot be zero or nil")
		}

		if IsZeroOrNilDec(multiplier.TakerPointsMultiplier) {
			return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "derivative market maker multiplier cannot be zero or nil")
		}

		if !SafeIsPositiveDec(multiplier.TakerPointsMultiplier) {
			return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "derivative market maker multiplier cannot be negative")
		}
	}
	return nil
}

func (c *TradingRewardCampaignInfo) ValidateBasic() error {
	if c == nil {
		return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "campaign info cannot be nil")
	}

	if len(c.QuoteDenoms) == 0 {
		return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "campaign quote denoms cannot be nil")
	}

	hasTradingRewardBoostInfoDefined := c != nil && c.TradingRewardBoostInfo != nil
	if hasTradingRewardBoostInfoDefined {
		if err := c.TradingRewardBoostInfo.ValidateBasic(); err != nil {
			return err
		}
	}

	hasDuplicatesInDisqualifiedMarkets := c != nil && HasDuplicates(c.DisqualifiedMarketIds)
	if hasDuplicatesInDisqualifiedMarkets {
		return sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "campaign contains duplicate disqualified market ids")
	}

	return nil
}

func validateCampaignRewardPool(pool *CampaignRewardPool, campaignDurationSeconds, prevStartTimestamp int64) (int64, error) {
	if pool == nil {
		return 0, sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "new campaign reward pool cannot be nil")
	}

	if pool.StartTimestamp <= prevStartTimestamp {
		return 0, sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "reward pool start timestamps must be in ascending order")
	}

	hasValidStartTimestamp := prevStartTimestamp == 0 || pool.StartTimestamp == (prevStartTimestamp+campaignDurationSeconds)
	if !hasValidStartTimestamp {
		return 0, sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "start timestamps not matching campaign duration")
	}

	prevStartTimestamp = pool.StartTimestamp

	hasDuplicatesInEpochRewards := HasDuplicatesCoin(pool.MaxCampaignRewards)
	if hasDuplicatesInEpochRewards {
		return 0, sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "reward pool campaign contains duplicate market coins")
	}

	for _, epochRewardDenom := range pool.MaxCampaignRewards {
		if !epochRewardDenom.IsValid() {
			return 0, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, epochRewardDenom.String())
		}

		if IsZeroOrNilInt(epochRewardDenom.Amount) {
			return 0, sdkerrors.Wrap(ErrInvalidTradingRewardCampaign, "reward pool contains zero or nil reward amount")
		}
	}

	return prevStartTimestamp, nil
}

// NewFeeDiscountProposal returns new instance of FeeDiscountProposal
func NewFeeDiscountProposal(title, description string, schedule *FeeDiscountSchedule) *FeeDiscountProposal {
	return &FeeDiscountProposal{
		Title:       title,
		Description: description,
		Schedule:    schedule,
	}
}

// Implements Proposal Interface
var _ gov.Content = &FeeDiscountProposal{}

// GetTitle returns the title of this proposal
func (p *FeeDiscountProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *FeeDiscountProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *FeeDiscountProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *FeeDiscountProposal) ProposalType() string {
	return ProposalTypeFeeDiscountProposal
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *FeeDiscountProposal) ValidateBasic() error {
	if p.Schedule == nil {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule cannot be nil")
	}

	if p.Schedule.BucketCount < 2 {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule must have at least 2 buckets")
	}

	if p.Schedule.BucketDuration < 10 {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule must have have bucket durations of at least 10 seconds")
	}

	if HasDuplicates(p.Schedule.QuoteDenoms) {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule cannot have duplicate quote denoms")
	}

	if HasDuplicates(p.Schedule.DisqualifiedMarketIds) {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule cannot have duplicate disqualified market ids")
	}

	if len(p.Schedule.TierInfos) < 1 {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule must have at least one discount tier")
	}

	for idx, tierInfo := range p.Schedule.TierInfos {
		if err := tierInfo.ValidateBasic(); err != nil {
			return err
		}

		if idx > 0 {
			prevTierInfo := p.Schedule.TierInfos[idx-1]

			if prevTierInfo.MakerDiscountRate.GT(tierInfo.MakerDiscountRate) {
				return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "successive MakerDiscountRates must be equal or larger than those of lower tiers")
			}

			if prevTierInfo.TakerDiscountRate.GT(tierInfo.TakerDiscountRate) {
				return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "successive TakerDiscountRates must be equal or larger than those of lower tiers")
			}

			if prevTierInfo.StakedAmount.GT(tierInfo.StakedAmount) {
				return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "successive StakedAmount must be equal or larger than those of lower tiers")
			}

			if prevTierInfo.FeePaidAmount.GT(tierInfo.FeePaidAmount) {
				return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "successive FeePaidAmount must be equal or larger than those of lower tiers")
			}
		}
	}

	return gov.ValidateAbstract(p)
}

func (t *FeeDiscountTierInfo) ValidateBasic() error {
	if !SafeIsPositiveDec(t.MakerDiscountRate) || t.MakerDiscountRate.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "MakerDiscountRate must be between 0 and 1")
	}

	if !SafeIsPositiveDec(t.TakerDiscountRate) || t.TakerDiscountRate.GT(sdk.OneDec()) {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "TakerDiscountRate must be between 0 and 1")
	}

	if !SafeIsPositiveInt(t.StakedAmount) {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "StakedAmount must be non-negative")
	}

	if !SafeIsPositiveDec(t.FeePaidAmount) {
		return sdkerrors.Wrap(ErrInvalidFeeDiscountSchedule, "FeePaidAmount must be non-negative")
	}
	return nil
}

package types

import (
	"fmt"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/common"
)

// constants
const (
	ProposalTypeExchangeEnable                     string = "ProposalTypeExchangeEnable"
	ProposalTypeBatchExchangeModification          string = "ProposalTypeBatchExchangeModification"
	ProposalTypeSpotMarketParamUpdate              string = "ProposalTypeSpotMarketParamUpdate"
	ProposalTypeSpotMarketLaunch                   string = "ProposalTypeSpotMarketLaunch"
	ProposalTypePerpetualMarketLaunch              string = "ProposalTypePerpetualMarketLaunch"
	ProposalTypeExpiryFuturesMarketLaunch          string = "ProposalTypeExpiryFuturesMarketLaunch"
	ProposalTypeDerivativeMarketParamUpdate        string = "ProposalTypeDerivativeMarketParamUpdate"
	ProposalTypeMarketForcedSettlement             string = "ProposalTypeMarketForcedSettlement"
	ProposalUpdateDenomDecimals                    string = "ProposalUpdateDenomDecimals"
	ProposalTypeTradingRewardCampaign              string = "ProposalTypeTradingRewardCampaign"
	ProposalTypeTradingRewardCampaignUpdate        string = "ProposalTypeTradingRewardCampaignUpdateProposal"
	ProposalTypeTradingRewardPointsUpdate          string = "ProposalTypeTradingRewardPointsUpdateProposal"
	ProposalTypeFeeDiscountProposal                string = "ProposalTypeFeeDiscountProposal"
	ProposalTypeBatchCommunityPoolSpendProposal    string = "ProposalTypeBatchCommunityPoolSpendProposal"
	ProposalTypeBinaryOptionsMarketLaunch          string = "ProposalTypeBinaryOptionsMarketLaunch"
	ProposalTypeBinaryOptionsMarketParamUpdate     string = "ProposalTypeBinaryOptionsMarketParamUpdate"
	ProposalAtomicMarketOrderFeeMultiplierSchedule string = "ProposalAtomicMarketOrderFeeMultiplierSchedule"
)

func init() {
	govtypes.RegisterProposalType(ProposalTypeExchangeEnable)
	govtypes.RegisterProposalType(ProposalTypeBatchExchangeModification)
	govtypes.RegisterProposalType(ProposalTypeSpotMarketParamUpdate)
	govtypes.RegisterProposalType(ProposalTypeSpotMarketLaunch)
	govtypes.RegisterProposalType(ProposalTypePerpetualMarketLaunch)
	govtypes.RegisterProposalType(ProposalTypeExpiryFuturesMarketLaunch)
	govtypes.RegisterProposalType(ProposalTypeDerivativeMarketParamUpdate)
	govtypes.RegisterProposalType(ProposalTypeMarketForcedSettlement)
	govtypes.RegisterProposalType(ProposalUpdateDenomDecimals)
	govtypes.RegisterProposalType(ProposalTypeTradingRewardCampaign)
	govtypes.RegisterProposalType(ProposalTypeTradingRewardCampaignUpdate)
	govtypes.RegisterProposalType(ProposalTypeTradingRewardPointsUpdate)
	govtypes.RegisterProposalType(ProposalTypeFeeDiscountProposal)
	govtypes.RegisterProposalType(ProposalTypeBatchCommunityPoolSpendProposal)
	govtypes.RegisterProposalType(ProposalTypeBinaryOptionsMarketLaunch)
	govtypes.RegisterProposalType(ProposalTypeBinaryOptionsMarketParamUpdate)
	govtypes.RegisterProposalType(ProposalAtomicMarketOrderFeeMultiplierSchedule)
}

func SafeIsPositiveInt(v math.Int) bool {
	if v.IsNil() {
		return false
	}

	return v.IsPositive()
}

func SafeIsPositiveDec(v math.LegacyDec) bool {
	if v.IsNil() {
		return false
	}

	return v.IsPositive()
}

func SafeIsNonNegativeDec(v math.LegacyDec) bool {
	if v.IsNil() {
		return false
	}

	return !v.IsNegative()
}

func IsZeroOrNilInt(v math.Int) bool {
	return v.IsNil() || v.IsZero()
}

func IsZeroOrNilDec(v math.LegacyDec) bool {
	return v.IsNil() || v.IsZero()
}

// Implements Proposal Interface
var _ govtypes.Content = &ExchangeEnableProposal{}

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
	return govtypes.ValidateAbstract(p)
}

// Implements Proposal Interface
var _ govtypes.Content = &BatchExchangeModificationProposal{}

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

	for _, proposal := range p.BinaryOptionsMarketLaunchProposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}

	for _, proposal := range p.BinaryOptionsParamUpdateProposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}

	if p.DenomDecimalsUpdateProposal != nil {
		if err := p.DenomDecimalsUpdateProposal.ValidateBasic(); err != nil {
			return err
		}
	}

	if p.FeeDiscountProposal != nil {
		if err := p.FeeDiscountProposal.ValidateBasic(); err != nil {
			return err
		}
	}

	for _, proposal := range p.MarketForcedSettlementProposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}

	return govtypes.ValidateAbstract(p)
}

// NewSpotMarketParamUpdateProposal returns new instance of SpotMarketParamUpdateProposal
func NewSpotMarketParamUpdateProposal(title, description string, marketID common.Hash, makerFeeRate, takerFeeRate, relayerFeeShareRate, minPriceTickSize, minQuantityTickSize, minNotional *math.LegacyDec, status MarketStatus, ticker string) *SpotMarketParamUpdateProposal {
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
		ticker,
		minNotional,
		nil,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &SpotMarketParamUpdateProposal{}

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
	if !IsHexHash(p.MarketId) {
		return errors.Wrap(ErrMarketInvalid, p.MarketId)
	}
	if p.MakerFeeRate == nil &&
		p.TakerFeeRate == nil &&
		p.RelayerFeeShareRate == nil &&
		p.MinPriceTickSize == nil &&
		p.MinQuantityTickSize == nil &&
		p.MinNotional == nil &&
		p.AdminInfo == nil &&
		p.Status == MarketStatus_Unspecified {
		return errors.Wrap(gov.ErrInvalidProposalContent, "At least one field should not be nil")
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
			return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
		}
	}
	if p.MinQuantityTickSize != nil {
		if err := ValidateTickSize(*p.MinQuantityTickSize); err != nil {
			return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
		}
	}
	if p.MinNotional != nil {
		if err := ValidateMinNotional(*p.MinNotional); err != nil {
			return errors.Wrap(ErrInvalidNotional, err.Error())
		}
	}

	if p.AdminInfo != nil {
		if p.AdminInfo.Admin != "" {
			if _, err := sdk.AccAddressFromBech32(p.AdminInfo.Admin); err != nil {
				return errors.Wrap(ErrInvalidAddress, err.Error())
			}
		}
		if p.AdminInfo.AdminPermissions > MaxPerm {
			return ErrInvalidPermissions
		}
	}

	if len(p.Ticker) > MaxTickerLength {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not exceed %d characters", MaxTickerLength)
	}

	switch p.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Active,
		MarketStatus_Paused,
		MarketStatus_Demolished,
		MarketStatus_Expired:

	default:
		return errors.Wrap(ErrInvalidMarketStatus, p.Status.String())
	}

	return govtypes.ValidateAbstract(p)
}

// NewSpotMarketLaunchProposal returns new instance of SpotMarketLaunchProposal
func NewSpotMarketLaunchProposal(
	title string,
	description string,
	ticker string,
	baseDenom string,
	quoteDenom string,
	minPriceTickSize math.LegacyDec,
	minQuantityTickSize math.LegacyDec,
	minNotional math.LegacyDec,
	makerFeeRate *math.LegacyDec,
	takerFeeRate *math.LegacyDec,
) *SpotMarketLaunchProposal {
	return &SpotMarketLaunchProposal{
		Title:               title,
		Description:         description,
		Ticker:              ticker,
		BaseDenom:           baseDenom,
		QuoteDenom:          quoteDenom,
		MinPriceTickSize:    minPriceTickSize,
		MinQuantityTickSize: minQuantityTickSize,
		MinNotional:         minNotional,
		MakerFeeRate:        makerFeeRate,
		TakerFeeRate:        takerFeeRate,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &SpotMarketLaunchProposal{}

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
	if p.Ticker == "" || len(p.Ticker) > MaxTickerLength {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not be empty or exceed %d characters", MaxTickerLength)
	}
	if p.BaseDenom == "" {
		return errors.Wrap(ErrInvalidBaseDenom, "base denom should not be empty")
	}
	if p.QuoteDenom == "" {
		return errors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if p.BaseDenom == p.QuoteDenom {
		return ErrSameDenoms
	}

	if err := ValidateTickSize(p.MinPriceTickSize); err != nil {
		return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}
	if err := ValidateMinNotional(p.MinNotional); err != nil {
		return errors.Wrap(ErrInvalidNotional, err.Error())
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
		return errors.Wrap(ErrFeeRatesRelation, "maker and taker fee rate must either be both nil or both specified")
	}

	if p.MakerFeeRate != nil && p.TakerFeeRate != nil {
		if p.MakerFeeRate.GT(*p.TakerFeeRate) {
			return ErrFeeRatesRelation
		}
	}

	return govtypes.ValidateAbstract(p)
}

// NewDerivativeMarketParamUpdateProposal returns new instance of DerivativeMarketParamUpdateProposal
func NewDerivativeMarketParamUpdateProposal(
	title string,
	description string,
	marketID string,
	initialMarginRatio *math.LegacyDec,
	maintenanceMarginRatio *math.LegacyDec,
	makerFeeRate *math.LegacyDec,
	takerFeeRate *math.LegacyDec,
	relayerFeeShareRate *math.LegacyDec,
	minPriceTickSize *math.LegacyDec,
	minQuantityTickSize *math.LegacyDec,
	minNotional *math.LegacyDec,
	hourlyInterestRate *math.LegacyDec,
	hourlyFundingRateCap *math.LegacyDec,
	status MarketStatus,
	oracleParams *OracleParams,
	ticker string,
	adminInfo *AdminInfo,
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
		Ticker:                 ticker,
		MinNotional:            minNotional,
		AdminInfo:              adminInfo,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &DerivativeMarketParamUpdateProposal{}

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
	if !IsHexHash(p.MarketId) {
		return errors.Wrap(ErrMarketInvalid, p.MarketId)
	}
	if p.MakerFeeRate == nil &&
		p.TakerFeeRate == nil &&
		p.RelayerFeeShareRate == nil &&
		p.MinPriceTickSize == nil &&
		p.MinQuantityTickSize == nil &&
		p.MinNotional == nil &&
		p.InitialMarginRatio == nil &&
		p.MaintenanceMarginRatio == nil &&
		p.HourlyInterestRate == nil &&
		p.HourlyFundingRateCap == nil &&
		p.Status == MarketStatus_Unspecified &&
		p.AdminInfo == nil &&
		p.OracleParams == nil {
		return errors.Wrap(gov.ErrInvalidProposalContent, "At least one field should not be nil")
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
			return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
		}
	}
	if p.MinQuantityTickSize != nil {
		if err := ValidateTickSize(*p.MinQuantityTickSize); err != nil {
			return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
		}
	}
	if p.MinNotional != nil {
		if err := ValidateMinNotional(*p.MinNotional); err != nil {
			return errors.Wrap(ErrInvalidNotional, err.Error())
		}
	}

	if p.HourlyInterestRate != nil {
		if err := ValidateHourlyInterestRate(*p.HourlyInterestRate); err != nil {
			return errors.Wrap(ErrInvalidHourlyInterestRate, err.Error())
		}
	}

	if p.HourlyFundingRateCap != nil {
		if err := ValidateHourlyFundingRateCap(*p.HourlyFundingRateCap); err != nil {
			return errors.Wrap(ErrInvalidHourlyFundingRateCap, err.Error())
		}
	}

	if p.AdminInfo != nil {
		if p.AdminInfo.Admin != "" {
			if _, err := sdk.AccAddressFromBech32(p.AdminInfo.Admin); err != nil {
				return errors.Wrap(ErrInvalidAddress, err.Error())
			}
		}
		if p.AdminInfo.AdminPermissions > MaxPerm {
			return ErrInvalidPermissions
		}
	}

	if len(p.Ticker) > MaxTickerLength {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not exceed %d characters", MaxTickerLength)
	}

	switch p.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Active,
		MarketStatus_Paused,
		MarketStatus_Demolished,
		MarketStatus_Expired:

	default:
		return errors.Wrap(ErrInvalidMarketStatus, p.Status.String())
	}

	if p.OracleParams != nil {
		if err := p.OracleParams.ValidateBasic(); err != nil {
			return err
		}
	}

	return govtypes.ValidateAbstract(p)
}

// NewMarketForcedSettlementProposal returns new instance of MarketForcedSettlementProposal
func NewMarketForcedSettlementProposal(
	title, description string, marketID string,
	settlementPrice *math.LegacyDec,
) *MarketForcedSettlementProposal {
	return &MarketForcedSettlementProposal{
		Title:           title,
		Description:     description,
		MarketId:        marketID,
		SettlementPrice: settlementPrice,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &MarketForcedSettlementProposal{}

// GetTitle returns the title of this proposal
func (p *MarketForcedSettlementProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *MarketForcedSettlementProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *MarketForcedSettlementProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *MarketForcedSettlementProposal) ProposalType() string {
	return ProposalTypeMarketForcedSettlement
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *MarketForcedSettlementProposal) ValidateBasic() error {
	if !IsHexHash(p.MarketId) {
		return errors.Wrap(ErrMarketInvalid, p.MarketId)
	}

	if p.SettlementPrice != nil && !SafeIsPositiveDec(*p.SettlementPrice) {
		return errors.Wrap(ErrInvalidSettlement, "settlement price must be positive for derivatives and nil for spot")
	}

	return govtypes.ValidateAbstract(p)
}

// NewUpdateDenomDecimalsProposal returns new instance of UpdateDenomDecimalsProposal
func NewUpdateDenomDecimalsProposal(
	title, description string,
	denomDecimals []*DenomDecimals,
) *UpdateDenomDecimalsProposal {
	return &UpdateDenomDecimalsProposal{
		Title:         title,
		Description:   description,
		DenomDecimals: denomDecimals,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &UpdateDenomDecimalsProposal{}

// GetTitle returns the title of this proposal
func (p *UpdateDenomDecimalsProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *UpdateDenomDecimalsProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *UpdateDenomDecimalsProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *UpdateDenomDecimalsProposal) ProposalType() string {
	return ProposalUpdateDenomDecimals
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *UpdateDenomDecimalsProposal) ValidateBasic() error {
	for _, d := range p.DenomDecimals {
		if err := d.Validate(); err != nil {
			return err
		}
	}
	return govtypes.ValidateAbstract(p)
}

func (d *DenomDecimals) Validate() error {
	if d.Denom == "" {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, d.Denom)
	}

	if d.Decimals <= 0 || d.Decimals > uint64(MaxOracleScaleFactor) {
		return errors.Wrapf(ErrInvalidDenomDecimal, "invalid decimals passed: %d", d.Decimals)
	}
	return nil
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
		return errors.Wrap(ErrInvalidOracle, "oracle base should not be empty")
	}
	if p.OracleQuote == "" {
		return errors.Wrap(ErrInvalidOracle, "oracle quote should not be empty")
	}
	if p.OracleBase == p.OracleQuote {
		return ErrSameOracles
	}
	switch p.OracleType {
	case oracletypes.OracleType_Band, oracletypes.OracleType_PriceFeed, oracletypes.OracleType_Coinbase, oracletypes.OracleType_Chainlink, oracletypes.OracleType_Razor,
		oracletypes.OracleType_Dia, oracletypes.OracleType_API3, oracletypes.OracleType_Uma, oracletypes.OracleType_Pyth, oracletypes.OracleType_BandIBC, oracletypes.OracleType_Provider:

	default:
		return errors.Wrap(ErrInvalidOracleType, p.OracleType.String())
	}

	if p.OracleScaleFactor > MaxOracleScaleFactor {
		return ErrExceedsMaxOracleScaleFactor
	}

	return nil
}

func NewProviderOracleParams(
	symbol string,
	oracleProvider string,
	oracleScaleFactor uint32,
	oracleType oracletypes.OracleType,
) *ProviderOracleParams {
	return &ProviderOracleParams{
		Symbol:            symbol,
		Provider:          oracleProvider,
		OracleScaleFactor: oracleScaleFactor,
		OracleType:        oracleType,
	}
}

func (p *ProviderOracleParams) ValidateBasic() error {
	if p.Symbol == "" {
		return errors.Wrap(ErrInvalidOracle, "oracle symbol should not be empty")
	}
	if p.Provider == "" {
		return errors.Wrap(ErrInvalidOracle, "oracle provider should not be empty")
	}

	if p.OracleType != oracletypes.OracleType_Provider {
		return errors.Wrap(ErrInvalidOracleType, p.OracleType.String())
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
	initialMarginRatio, maintenanceMarginRatio, makerFeeRate, takerFeeRate, minPriceTickSize, minQuantityTickSize, minNotional math.LegacyDec,
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
		MinNotional:            minNotional,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &PerpetualMarketLaunchProposal{}

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
	if p.Ticker == "" || len(p.Ticker) > MaxTickerLength {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not be empty or exceed %d characters", MaxTickerLength)
	}
	if p.QuoteDenom == "" {
		return errors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}

	oracleParams := NewOracleParams(p.OracleBase, p.OracleQuote, p.OracleScaleFactor, p.OracleType)
	if err := oracleParams.ValidateBasic(); err != nil {
		return err
	}
	if err := ValidateMakerFee(p.MakerFeeRate); err != nil {
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
		return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}
	if err := ValidateMinNotional(p.MinNotional); err != nil {
		return errors.Wrap(ErrInvalidNotional, err.Error())
	}

	return govtypes.ValidateAbstract(p)
}

// NewExpiryFuturesMarketLaunchProposal returns new instance of ExpiryFuturesMarketLaunchProposal
func NewExpiryFuturesMarketLaunchProposal(
	title, description, ticker, quoteDenom,
	oracleBase, oracleQuote string, oracleScaleFactor uint32, oracleType oracletypes.OracleType, expiry int64,
	initialMarginRatio, maintenanceMarginRatio, makerFeeRate, takerFeeRate, minPriceTickSize, minQuantityTickSize, minNotional math.LegacyDec,
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
		MinNotional:            minNotional,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &ExpiryFuturesMarketLaunchProposal{}

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
	if p.Ticker == "" || len(p.Ticker) > MaxTickerLength {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not be empty or exceed %d characters", MaxTickerLength)
	}
	if p.QuoteDenom == "" {
		return errors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}

	oracleParams := NewOracleParams(p.OracleBase, p.OracleQuote, p.OracleScaleFactor, p.OracleType)
	if err := oracleParams.ValidateBasic(); err != nil {
		return err
	}
	if p.Expiry <= 0 {
		return errors.Wrap(ErrInvalidExpiry, "expiry should not be empty")
	}
	if err := ValidateMakerFee(p.MakerFeeRate); err != nil {
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
		return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}
	if err := ValidateMinNotional(p.MinNotional); err != nil {
		return errors.Wrap(ErrInvalidNotional, err.Error())
	}

	return govtypes.ValidateAbstract(p)
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
var _ govtypes.Content = &TradingRewardCampaignUpdateProposal{}

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
			return errors.Wrap(ErrInvalidTradingRewardCampaign, "campaign reward pool addition cannot be nil")
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

	return govtypes.ValidateAbstract(p)
}

// Implements Proposal Interface
var _ govtypes.Content = &TradingRewardPendingPointsUpdateProposal{}

// GetTitle returns the title of this proposal
func (p *TradingRewardPendingPointsUpdateProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *TradingRewardPendingPointsUpdateProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *TradingRewardPendingPointsUpdateProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *TradingRewardPendingPointsUpdateProposal) ProposalType() string {
	return ProposalTypeTradingRewardPointsUpdate
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *TradingRewardPendingPointsUpdateProposal) ValidateBasic() error {
	if len(p.RewardPointUpdates) == 0 {
		return errors.Wrap(ErrInvalidTradingRewardsPendingPointsUpdate, "reward pending points update cannot be nil")
	}

	if p.PendingPoolTimestamp <= 0 {
		return errors.Wrap(ErrInvalidTradingRewardsPendingPointsUpdate, "pending pool timestamp cannot be zero")
	}

	accountAddresses := make([]string, 0)

	for _, rewardPointUpdate := range p.RewardPointUpdates {
		if rewardPointUpdate == nil {
			return errors.Wrap(ErrInvalidTradingRewardsPendingPointsUpdate, "reward pending point update cannot be nil")
		}

		_, err := sdk.AccAddressFromBech32(rewardPointUpdate.AccountAddress)

		if err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, rewardPointUpdate.AccountAddress)
		}

		accountAddresses = append(accountAddresses, rewardPointUpdate.AccountAddress)

		if rewardPointUpdate.NewPoints.IsNil() {
			return errors.Wrap(ErrInvalidTradingRewardsPendingPointsUpdate, "reward pending points cannot be nil")
		}

		if rewardPointUpdate.NewPoints.IsNegative() {
			return errors.Wrap(ErrInvalidTradingRewardsPendingPointsUpdate, "reward pending points cannot be negative")
		}
	}

	hasDuplicateAccountAddresses := HasDuplicates(accountAddresses)
	if hasDuplicateAccountAddresses {
		return errors.Wrap(ErrInvalidTradingRewardsPendingPointsUpdate, "account address cannot have duplicates")
	}

	return govtypes.ValidateAbstract(p)
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
var _ govtypes.Content = &TradingRewardCampaignLaunchProposal{}

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
		return errors.Wrap(ErrInvalidTradingRewardCampaign, "new campaign info cannot be nil")
	}

	if len(p.CampaignRewardPools) == 0 {
		return errors.Wrap(ErrInvalidTradingRewardCampaign, "new campaign reward pools cannot be nil")
	}

	err = p.CampaignInfo.ValidateBasic()
	if err != nil {
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
		return errors.Wrap(ErrInvalidTradingRewardCampaign, "boosted spot market ids is not matching spot market multipliers")
	}

	for _, marketID := range t.BoostedSpotMarketIds {
		if !IsHexHash(marketID) {
			return errors.Wrap(ErrMarketInvalid, marketID)
		}
	}

	for _, marketID := range t.BoostedDerivativeMarketIds {
		if !IsHexHash(marketID) {
			return errors.Wrap(ErrMarketInvalid, marketID)
		}
	}

	if len(t.BoostedDerivativeMarketIds) != len(t.DerivativeMarketMultipliers) {
		return errors.Wrap(ErrInvalidTradingRewardCampaign, "boosted derivative market ids is not matching derivative market multipliers")
	}

	hasDuplicatesInMarkets := HasDuplicates(t.BoostedSpotMarketIds) || HasDuplicates(t.BoostedDerivativeMarketIds)
	if hasDuplicatesInMarkets {
		return errors.Wrap(ErrInvalidTradingRewardCampaign, "campaign contains duplicate boosted market ids")
	}

	for _, multiplier := range t.SpotMarketMultipliers {
		if IsZeroOrNilDec(multiplier.MakerPointsMultiplier) {
			return errors.Wrap(ErrInvalidTradingRewardCampaign, "spot market maker multiplier cannot be zero or nil")
		}

		if IsZeroOrNilDec(multiplier.TakerPointsMultiplier) {
			return errors.Wrap(ErrInvalidTradingRewardCampaign, "spot market taker multiplier cannot be zero or nil")
		}

		if !SafeIsPositiveDec(multiplier.MakerPointsMultiplier) {
			return errors.Wrap(ErrInvalidTradingRewardCampaign, "spot market maker multiplier cannot be negative")
		}

		if !SafeIsPositiveDec(multiplier.TakerPointsMultiplier) {
			return errors.Wrap(ErrInvalidTradingRewardCampaign, "spot market taker multiplier cannot be negative")
		}
	}

	for _, multiplier := range t.DerivativeMarketMultipliers {
		if IsZeroOrNilDec(multiplier.MakerPointsMultiplier) {
			return errors.Wrap(ErrInvalidTradingRewardCampaign, "derivative market maker multiplier cannot be zero or nil")
		}

		if IsZeroOrNilDec(multiplier.TakerPointsMultiplier) {
			return errors.Wrap(ErrInvalidTradingRewardCampaign, "derivative market taker multiplier cannot be zero or nil")
		}

		if !SafeIsPositiveDec(multiplier.MakerPointsMultiplier) {
			return errors.Wrap(ErrInvalidTradingRewardCampaign, "derivative market maker multiplier cannot be negative")
		}

		if !SafeIsPositiveDec(multiplier.TakerPointsMultiplier) {
			return errors.Wrap(ErrInvalidTradingRewardCampaign, "derivative market taker multiplier cannot be negative")
		}
	}
	return nil
}

func (c *TradingRewardCampaignInfo) ValidateBasic() error {
	if c == nil {
		return errors.Wrap(ErrInvalidTradingRewardCampaign, "campaign info cannot be nil")
	}

	if c.CampaignDurationSeconds <= 0 {
		return errors.Wrap(ErrInvalidTradingRewardCampaign, "campaign duration cannot be zero")
	}

	if len(c.QuoteDenoms) == 0 {
		return errors.Wrap(ErrInvalidTradingRewardCampaign, "campaign quote denoms cannot be nil")
	}

	hasTradingRewardBoostInfoDefined := c != nil && c.TradingRewardBoostInfo != nil
	if hasTradingRewardBoostInfoDefined {
		if err := c.TradingRewardBoostInfo.ValidateBasic(); err != nil {
			return err
		}
	}

	for _, marketID := range c.DisqualifiedMarketIds {
		if !IsHexHash(marketID) {
			return errors.Wrap(ErrMarketInvalid, marketID)
		}
	}

	hasDuplicatesInDisqualifiedMarkets := c != nil && HasDuplicates(c.DisqualifiedMarketIds)
	if hasDuplicatesInDisqualifiedMarkets {
		return errors.Wrap(ErrInvalidTradingRewardCampaign, "campaign contains duplicate disqualified market ids")
	}

	return nil
}

func validateCampaignRewardPool(pool *CampaignRewardPool, campaignDurationSeconds, prevStartTimestamp int64) (int64, error) {
	if pool == nil {
		return 0, errors.Wrap(ErrInvalidTradingRewardCampaign, "new campaign reward pool cannot be nil")
	}

	if pool.StartTimestamp <= prevStartTimestamp {
		return 0, errors.Wrap(ErrInvalidTradingRewardCampaign, "reward pool start timestamps must be in ascending order")
	}

	hasValidStartTimestamp := prevStartTimestamp == 0 || pool.StartTimestamp == (prevStartTimestamp+campaignDurationSeconds)
	if !hasValidStartTimestamp {
		return 0, errors.Wrap(ErrInvalidTradingRewardCampaign, "start timestamps not matching campaign duration")
	}

	prevStartTimestamp = pool.StartTimestamp

	hasDuplicatesInEpochRewards := HasDuplicatesCoin(pool.MaxCampaignRewards)
	if hasDuplicatesInEpochRewards {
		return 0, errors.Wrap(ErrInvalidTradingRewardCampaign, "reward pool campaign contains duplicate market coins")
	}

	for _, epochRewardDenom := range pool.MaxCampaignRewards {
		if !epochRewardDenom.IsValid() {
			return 0, errors.Wrap(sdkerrors.ErrInvalidCoins, epochRewardDenom.String())
		}

		if IsZeroOrNilInt(epochRewardDenom.Amount) {
			return 0, errors.Wrap(ErrInvalidTradingRewardCampaign, "reward pool contains zero or nil reward amount")
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
var _ govtypes.Content = &FeeDiscountProposal{}

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
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule cannot be nil")
	}

	if p.Schedule.BucketCount < 2 {
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule must have at least 2 buckets")
	}

	if p.Schedule.BucketDuration < 10 {
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule must have have bucket durations of at least 10 seconds")
	}

	if HasDuplicates(p.Schedule.QuoteDenoms) {
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule cannot have duplicate quote denoms")
	}

	for _, marketID := range p.Schedule.DisqualifiedMarketIds {
		if !IsHexHash(marketID) {
			return errors.Wrap(ErrMarketInvalid, marketID)
		}
	}

	if HasDuplicates(p.Schedule.DisqualifiedMarketIds) {
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule cannot have duplicate disqualified market ids")
	}

	if len(p.Schedule.TierInfos) < 1 {
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "new fee discount schedule must have at least one discount tier")
	}

	for idx, tierInfo := range p.Schedule.TierInfos {
		if err := tierInfo.ValidateBasic(); err != nil {
			return err
		}

		if idx > 0 {
			prevTierInfo := p.Schedule.TierInfos[idx-1]

			if prevTierInfo.MakerDiscountRate.GT(tierInfo.MakerDiscountRate) {
				return errors.Wrap(ErrInvalidFeeDiscountSchedule, "successive MakerDiscountRates must be equal or larger than those of lower tiers")
			}

			if prevTierInfo.TakerDiscountRate.GT(tierInfo.TakerDiscountRate) {
				return errors.Wrap(ErrInvalidFeeDiscountSchedule, "successive TakerDiscountRates must be equal or larger than those of lower tiers")
			}

			if prevTierInfo.StakedAmount.GT(tierInfo.StakedAmount) {
				return errors.Wrap(ErrInvalidFeeDiscountSchedule, "successive StakedAmount must be equal or larger than those of lower tiers")
			}

			if prevTierInfo.Volume.GT(tierInfo.Volume) {
				return errors.Wrap(ErrInvalidFeeDiscountSchedule, "successive Volume must be equal or larger than those of lower tiers")
			}
		}
	}

	return govtypes.ValidateAbstract(p)
}

func (t *FeeDiscountTierInfo) ValidateBasic() error {
	if !SafeIsNonNegativeDec(t.MakerDiscountRate) || t.MakerDiscountRate.GT(math.LegacyOneDec()) {
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "MakerDiscountRate must be between 0 and 1")
	}

	if !SafeIsNonNegativeDec(t.TakerDiscountRate) || t.TakerDiscountRate.GT(math.LegacyOneDec()) {
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "TakerDiscountRate must be between 0 and 1")
	}

	if !SafeIsPositiveInt(t.StakedAmount) {
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "StakedAmount must be non-negative")
	}

	if !SafeIsPositiveDec(t.Volume) {
		return errors.Wrap(ErrInvalidFeeDiscountSchedule, "Volume must be non-negative")
	}
	return nil
}

// Implements Proposal Interface
var _ govtypes.Content = &BatchCommunityPoolSpendProposal{}

// GetTitle returns the title of this proposal.
func (p *BatchCommunityPoolSpendProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *BatchCommunityPoolSpendProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *BatchCommunityPoolSpendProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *BatchCommunityPoolSpendProposal) ProposalType() string {
	return ProposalTypeBatchCommunityPoolSpendProposal
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *BatchCommunityPoolSpendProposal) ValidateBasic() error {
	for _, proposal := range p.Proposals {
		if err := proposal.ValidateBasic(); err != nil {
			return err
		}
	}
	return govtypes.ValidateAbstract(p)
}

// NewBinaryOptionsMarketLaunchProposal returns new instance of BinaryOptionsMarketLaunchProposal
func NewBinaryOptionsMarketLaunchProposal(
	title, description, ticker, oracleSymbol, oracleProvider string,
	oracleType oracletypes.OracleType, oracleScaleFactor uint32,
	expirationTimestamp, settlementTimestamp int64,
	admin, quoteDenom string,
	makerFeeRate, takerFeeRate, minPriceTickSize, minQuantityTickSize, minNotional math.LegacyDec,

) *BinaryOptionsMarketLaunchProposal {
	return &BinaryOptionsMarketLaunchProposal{
		Title:               title,
		Description:         description,
		Ticker:              ticker,
		OracleSymbol:        oracleSymbol,
		OracleProvider:      oracleProvider,
		OracleType:          oracleType,
		OracleScaleFactor:   oracleScaleFactor,
		ExpirationTimestamp: expirationTimestamp,
		SettlementTimestamp: settlementTimestamp,
		Admin:               admin,
		QuoteDenom:          quoteDenom,
		MakerFeeRate:        makerFeeRate,
		TakerFeeRate:        takerFeeRate,
		MinPriceTickSize:    minPriceTickSize,
		MinQuantityTickSize: minQuantityTickSize,
		MinNotional:         minNotional,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &BinaryOptionsMarketLaunchProposal{}

// GetTitle returns the title of this proposal.
func (p *BinaryOptionsMarketLaunchProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal.
func (p *BinaryOptionsMarketLaunchProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *BinaryOptionsMarketLaunchProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *BinaryOptionsMarketLaunchProposal) ProposalType() string {
	return ProposalTypeBinaryOptionsMarketLaunch
}

// ValidateBasic returns ValidateBasic result of a perpetual market launch proposal.
func (p *BinaryOptionsMarketLaunchProposal) ValidateBasic() error {
	if p.Ticker == "" || len(p.Ticker) > MaxTickerLength {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not be empty or exceed %d characters", MaxTickerLength)
	}
	if p.OracleSymbol == "" {
		return errors.Wrap(ErrInvalidOracle, "oracle symbol should not be empty")
	}
	if p.OracleProvider == "" {
		return errors.Wrap(ErrInvalidOracle, "oracle provider should not be empty")
	}
	if p.OracleType != oracletypes.OracleType_Provider {
		return errors.Wrap(ErrInvalidOracleType, p.OracleType.String())
	}
	if p.OracleScaleFactor > MaxOracleScaleFactor {
		return ErrExceedsMaxOracleScaleFactor
	}

	if p.ExpirationTimestamp >= p.SettlementTimestamp || p.ExpirationTimestamp < 0 || p.SettlementTimestamp < 0 {
		return ErrInvalidExpiry
	}

	if p.Admin != "" {
		_, err := sdk.AccAddressFromBech32(p.Admin)
		if err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, p.Admin)
		}
	}
	if p.QuoteDenom == "" {
		return errors.Wrap(ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if err := ValidateMakerFee(p.MakerFeeRate); err != nil {
		return err
	}
	if err := ValidateFee(p.TakerFeeRate); err != nil {
		return err
	}

	if p.MakerFeeRate.GT(p.TakerFeeRate) {
		return ErrFeeRatesRelation
	}

	if err := ValidateTickSize(p.MinPriceTickSize); err != nil {
		return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
	}
	if err := ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
	}
	if err := ValidateMinNotional(p.MinNotional); err != nil {
		return errors.Wrap(ErrInvalidNotional, err.Error())
	}

	return govtypes.ValidateAbstract(p)
}

// NewBinaryOptionsMarketParamUpdateProposal returns new instance of BinaryOptionsMarketParamUpdateProposal
func NewBinaryOptionsMarketParamUpdateProposal(
	title string,
	description string,
	marketID string,
	makerFeeRate, takerFeeRate, relayerFeeShareRate, minPriceTickSize, minQuantityTickSize, minNotional *math.LegacyDec,
	expirationTimestamp, settlementTimestamp int64,
	admin string,
	status MarketStatus,
	oracleParams *ProviderOracleParams,
	ticker string,
) *BinaryOptionsMarketParamUpdateProposal {
	return &BinaryOptionsMarketParamUpdateProposal{
		Title:               title,
		Description:         description,
		MarketId:            marketID,
		MakerFeeRate:        makerFeeRate,
		TakerFeeRate:        takerFeeRate,
		RelayerFeeShareRate: relayerFeeShareRate,
		MinPriceTickSize:    minPriceTickSize,
		MinQuantityTickSize: minQuantityTickSize,
		MinNotional:         minNotional,
		ExpirationTimestamp: expirationTimestamp,
		SettlementTimestamp: settlementTimestamp,
		Admin:               admin,
		Status:              status,
		OracleParams:        oracleParams,
		Ticker:              ticker,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &BinaryOptionsMarketParamUpdateProposal{}

// GetTitle returns the title of this proposal
func (p *BinaryOptionsMarketParamUpdateProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *BinaryOptionsMarketParamUpdateProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *BinaryOptionsMarketParamUpdateProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *BinaryOptionsMarketParamUpdateProposal) ProposalType() string {
	return ProposalTypeBinaryOptionsMarketParamUpdate
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *BinaryOptionsMarketParamUpdateProposal) ValidateBasic() error {
	if !IsHexHash(p.MarketId) {
		return errors.Wrap(ErrMarketInvalid, p.MarketId)
	}
	if p.MakerFeeRate == nil &&
		p.TakerFeeRate == nil &&
		p.RelayerFeeShareRate == nil &&
		p.MinPriceTickSize == nil &&
		p.MinQuantityTickSize == nil &&
		p.MinNotional == nil &&
		p.Status == MarketStatus_Unspecified &&
		p.ExpirationTimestamp == 0 &&
		p.SettlementTimestamp == 0 &&
		p.SettlementPrice == nil &&
		p.Admin == "" &&
		p.OracleParams == nil {
		return errors.Wrap(gov.ErrInvalidProposalContent, "At least one field should not be nil")
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
			return errors.Wrap(ErrInvalidPriceTickSize, err.Error())
		}
	}

	if p.MinQuantityTickSize != nil {
		if err := ValidateTickSize(*p.MinQuantityTickSize); err != nil {
			return errors.Wrap(ErrInvalidQuantityTickSize, err.Error())
		}
	}

	if p.MinNotional != nil {
		if err := ValidateMinNotional(*p.MinNotional); err != nil {
			return errors.Wrap(ErrInvalidNotional, err.Error())
		}
	}

	if p.ExpirationTimestamp != 0 && p.SettlementTimestamp != 0 {
		if p.ExpirationTimestamp >= p.SettlementTimestamp || p.ExpirationTimestamp < 0 || p.SettlementTimestamp < 0 {
			return ErrInvalidExpiry
		}
	}

	if p.SettlementTimestamp < 0 {
		return ErrInvalidSettlement
	}

	if p.Admin != "" {
		if _, err := sdk.AccAddressFromBech32(p.Admin); err != nil {
			return err
		}
	}

	if len(p.Ticker) > MaxTickerLength {
		return errors.Wrapf(ErrInvalidTicker, "ticker should not exceed %d characters", MaxTickerLength)
	}

	// price is either nil (not set), -1 (demolish with refund) or [0..1] (demolish with settle)
	switch {
	case p.SettlementPrice == nil,
		p.SettlementPrice.IsNil():
		// ok
	case p.SettlementPrice.Equal(BinaryOptionsMarketRefundFlagPrice),
		p.SettlementPrice.GTE(math.LegacyZeroDec()) && p.SettlementPrice.LTE(MaxBinaryOptionsOrderPrice):
		if p.Status != MarketStatus_Demolished {
			return errors.Wrapf(ErrInvalidMarketStatus, "status should be set to demolished when the settlement price is set, status: %s", p.Status.String())
		}
		// ok
	default:
		return errors.Wrap(ErrInvalidPrice, p.SettlementPrice.String())
	}

	switch p.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Demolished:
	default:
		return errors.Wrap(ErrInvalidMarketStatus, p.Status.String())
	}

	if p.OracleParams != nil {
		if err := p.OracleParams.ValidateBasic(); err != nil {
			return err
		}
	}

	return govtypes.ValidateAbstract(p)
}

// Implements Proposal Interface
var _ govtypes.Content = &AtomicMarketOrderFeeMultiplierScheduleProposal{}

// GetTitle returns the title of this proposal
func (p *AtomicMarketOrderFeeMultiplierScheduleProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *AtomicMarketOrderFeeMultiplierScheduleProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (p *AtomicMarketOrderFeeMultiplierScheduleProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (p *AtomicMarketOrderFeeMultiplierScheduleProposal) ProposalType() string {
	return ProposalAtomicMarketOrderFeeMultiplierSchedule
}

func (p *AtomicMarketOrderFeeMultiplierScheduleProposal) ValidateBasic() error {
	if len(p.MarketFeeMultipliers) == 0 {
		return errors.Wrap(gov.ErrInvalidProposalContent, "At least one fee multiplier should be provided")
	}
	for _, m := range p.MarketFeeMultipliers {
		if !IsHexHash(m.MarketId) {
			return errors.Wrap(ErrMarketInvalid, m.MarketId)
		}
		multiplier := m.FeeMultiplier
		if multiplier.IsNil() {
			return fmt.Errorf("atomic taker fee multiplier cannot be nil: %v", multiplier)
		}

		if multiplier.LT(math.LegacyOneDec()) {
			return fmt.Errorf("atomic taker fee multiplier cannot be less than 1: %v", multiplier)
		}

		if multiplier.GT(MaxFeeMultiplier) {
			return fmt.Errorf("atomicMarketOrderFeeMultiplier cannot be bigger than %v: %v", multiplier, MaxFeeMultiplier)
		}
	}
	return govtypes.ValidateAbstract(p)
}

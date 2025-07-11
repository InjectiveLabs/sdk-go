package v2

import (
	"fmt"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
)

// constants
const (
	ProposalTypeExchangeEnable                         string = "ProposalTypeExchangeEnableV2"
	ProposalTypeBatchExchangeModification              string = "ProposalTypeBatchExchangeModificationV2"
	ProposalTypeSpotMarketParamUpdate                  string = "ProposalTypeSpotMarketParamUpdateV2"
	ProposalTypeSpotMarketLaunch                       string = "ProposalTypeSpotMarketLaunchV2"
	ProposalTypePerpetualMarketLaunch                  string = "ProposalTypePerpetualMarketLaunchV2"
	ProposalTypeExpiryFuturesMarketLaunch              string = "ProposalTypeExpiryFuturesMarketLaunchV2"
	ProposalTypeDerivativeMarketParamUpdate            string = "ProposalTypeDerivativeMarketParamUpdateV2"
	ProposalTypeMarketForcedSettlement                 string = "ProposalTypeMarketForcedSettlementV2"
	ProposalUpdateAuctionExchangeTransferDenomDecimals string = "ProposalUpdateAuctionExchangeTransferDenomDecimalsV2"
	ProposalTypeTradingRewardCampaign                  string = "ProposalTypeTradingRewardCampaignV2"
	ProposalTypeTradingRewardCampaignUpdate            string = "ProposalTypeTradingRewardCampaignUpdateProposalV2"
	ProposalTypeTradingRewardPointsUpdate              string = "ProposalTypeTradingRewardPointsUpdateProposalV2"
	ProposalTypeFeeDiscountProposal                    string = "ProposalTypeFeeDiscountProposalV2"
	ProposalTypeBatchCommunityPoolSpendProposal        string = "ProposalTypeBatchCommunityPoolSpendProposalV2"
	ProposalTypeBinaryOptionsMarketLaunch              string = "ProposalTypeBinaryOptionsMarketLaunchV2"
	ProposalTypeBinaryOptionsMarketParamUpdate         string = "ProposalTypeBinaryOptionsMarketParamUpdateV2"
	ProposalAtomicMarketOrderFeeMultiplierSchedule     string = "ProposalAtomicMarketOrderFeeMultiplierScheduleV2"
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
	govtypes.RegisterProposalType(ProposalUpdateAuctionExchangeTransferDenomDecimals)
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
		return types.ErrBadField
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

	if p.AuctionExchangeTransferDenomDecimalsUpdateProposal != nil {
		if err := p.AuctionExchangeTransferDenomDecimalsUpdateProposal.ValidateBasic(); err != nil {
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
func NewSpotMarketParamUpdateProposal(
	title, description string,
	marketID common.Hash,
	makerFeeRate, takerFeeRate, relayerFeeShareRate, minPriceTickSize, minQuantityTickSize, minNotional *math.LegacyDec,
	status MarketStatus,
	ticker string,
	baseDecimals, quoteDecimals uint32,
) *SpotMarketParamUpdateProposal {
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
		baseDecimals,
		quoteDecimals,
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
	if !types.IsHexHash(p.MarketId) {
		return errors.Wrap(types.ErrMarketInvalid, p.MarketId)
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
		if err := types.ValidateMakerFee(*p.MakerFeeRate); err != nil {
			return err
		}
	}
	if p.TakerFeeRate != nil {
		if err := types.ValidateFee(*p.TakerFeeRate); err != nil {
			return err
		}
	}
	if p.RelayerFeeShareRate != nil {
		if err := types.ValidateFee(*p.RelayerFeeShareRate); err != nil {
			return err
		}
	}

	if p.MinPriceTickSize != nil {
		if err := types.ValidateTickSize(*p.MinPriceTickSize); err != nil {
			return errors.Wrap(types.ErrInvalidPriceTickSize, err.Error())
		}
	}
	if p.MinQuantityTickSize != nil {
		if err := types.ValidateTickSize(*p.MinQuantityTickSize); err != nil {
			return errors.Wrap(types.ErrInvalidQuantityTickSize, err.Error())
		}
	}
	if p.MinNotional != nil {
		if err := types.ValidateMinNotional(*p.MinNotional); err != nil {
			return errors.Wrap(types.ErrInvalidNotional, err.Error())
		}
	}

	if p.AdminInfo != nil {
		if p.AdminInfo.Admin != "" {
			if _, err := sdk.AccAddressFromBech32(p.AdminInfo.Admin); err != nil {
				return errors.Wrap(types.ErrInvalidAddress, err.Error())
			}
		}
		if p.AdminInfo.AdminPermissions > types.MaxPerm {
			return types.ErrInvalidPermissions
		}
	}

	if len(p.Ticker) > types.MaxTickerLength {
		return errors.Wrapf(types.ErrInvalidTicker, "ticker should not exceed %d characters", types.MaxTickerLength)
	}

	switch p.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Active,
		MarketStatus_Paused,
		MarketStatus_Demolished,
		MarketStatus_Expired:

	default:
		return errors.Wrap(types.ErrInvalidMarketStatus, p.Status.String())
	}

	if p.BaseDecimals > types.MaxDecimals {
		return errors.Wrap(types.ErrInvalidDenomDecimal, "base decimals is invalid")
	}
	if p.QuoteDecimals > types.MaxDecimals {
		return errors.Wrap(types.ErrInvalidDenomDecimal, "quote decimals is invalid")
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
	baseDecimals uint32,
	quoteDecimals uint32,
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
		BaseDecimals:        baseDecimals,
		QuoteDecimals:       quoteDecimals,
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
	if p.Ticker == "" || len(p.Ticker) > types.MaxTickerLength {
		return errors.Wrapf(types.ErrInvalidTicker, "ticker should not be empty or exceed %d characters", types.MaxTickerLength)
	}
	if p.BaseDenom == "" {
		return errors.Wrap(types.ErrInvalidBaseDenom, "base denom should not be empty")
	}
	if p.QuoteDenom == "" {
		return errors.Wrap(types.ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if p.BaseDenom == p.QuoteDenom {
		return types.ErrSameDenoms
	}

	if err := types.ValidateTickSize(p.MinPriceTickSize); err != nil {
		return errors.Wrap(types.ErrInvalidPriceTickSize, err.Error())
	}
	if err := types.ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return errors.Wrap(types.ErrInvalidQuantityTickSize, err.Error())
	}
	if err := types.ValidateMinNotional(p.MinNotional); err != nil {
		return errors.Wrap(types.ErrInvalidNotional, err.Error())
	}

	if p.MakerFeeRate != nil {
		if err := types.ValidateMakerFee(*p.MakerFeeRate); err != nil {
			return err
		}
	}

	if p.TakerFeeRate != nil {
		if err := types.ValidateFee(*p.TakerFeeRate); err != nil {
			return err
		}
	}

	if (p.MakerFeeRate == nil && p.TakerFeeRate != nil) || (p.MakerFeeRate != nil && p.TakerFeeRate == nil) {
		return errors.Wrap(types.ErrFeeRatesRelation, "maker and taker fee rate must either be both nil or both specified")
	}

	if p.MakerFeeRate != nil && p.TakerFeeRate != nil {
		if p.MakerFeeRate.GT(*p.TakerFeeRate) {
			return types.ErrFeeRatesRelation
		}
	}

	if p.BaseDecimals > types.MaxDecimals {
		return errors.Wrap(types.ErrInvalidDenomDecimal, "base decimals is invalid")
	}
	if p.QuoteDecimals > types.MaxDecimals {
		return errors.Wrap(types.ErrInvalidDenomDecimal, "quote decimals is invalid")
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
	reduceMarginRatio *math.LegacyDec,
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
		ReduceMarginRatio:      reduceMarginRatio,
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
	if !types.IsHexHash(p.MarketId) {
		return errors.Wrap(types.ErrMarketInvalid, p.MarketId)
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
		if err := types.ValidateMakerFee(*p.MakerFeeRate); err != nil {
			return err
		}
	}
	if p.TakerFeeRate != nil {
		if err := types.ValidateFee(*p.TakerFeeRate); err != nil {
			return err
		}
	}

	if p.RelayerFeeShareRate != nil {
		if err := types.ValidateFee(*p.RelayerFeeShareRate); err != nil {
			return err
		}
	}

	if p.InitialMarginRatio != nil {
		if err := types.ValidateMarginRatio(*p.InitialMarginRatio); err != nil {
			return err
		}
	}
	if p.MaintenanceMarginRatio != nil {
		if err := types.ValidateMarginRatio(*p.MaintenanceMarginRatio); err != nil {
			return err
		}
	}

	if p.MinPriceTickSize != nil {
		if err := types.ValidateTickSize(*p.MinPriceTickSize); err != nil {
			return errors.Wrap(types.ErrInvalidPriceTickSize, err.Error())
		}
	}
	if p.MinQuantityTickSize != nil {
		if err := types.ValidateTickSize(*p.MinQuantityTickSize); err != nil {
			return errors.Wrap(types.ErrInvalidQuantityTickSize, err.Error())
		}
	}
	if p.MinNotional != nil {
		if err := types.ValidateMinNotional(*p.MinNotional); err != nil {
			return errors.Wrap(types.ErrInvalidNotional, err.Error())
		}
	}

	if p.HourlyInterestRate != nil {
		if err := types.ValidateHourlyInterestRate(*p.HourlyInterestRate); err != nil {
			return errors.Wrap(types.ErrInvalidHourlyInterestRate, err.Error())
		}
	}

	if p.HourlyFundingRateCap != nil {
		if err := types.ValidateHourlyFundingRateCap(*p.HourlyFundingRateCap); err != nil {
			return errors.Wrap(types.ErrInvalidHourlyFundingRateCap, err.Error())
		}
	}

	if p.AdminInfo != nil {
		if p.AdminInfo.Admin != "" {
			if _, err := sdk.AccAddressFromBech32(p.AdminInfo.Admin); err != nil {
				return errors.Wrap(types.ErrInvalidAddress, err.Error())
			}
		}
		if p.AdminInfo.AdminPermissions > types.MaxPerm {
			return types.ErrInvalidPermissions
		}
	}

	if len(p.Ticker) > types.MaxTickerLength {
		return errors.Wrapf(types.ErrInvalidTicker, "ticker should not exceed %d characters", types.MaxTickerLength)
	}

	switch p.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Active,
		MarketStatus_Paused,
		MarketStatus_Demolished,
		MarketStatus_Expired:

	default:
		return errors.Wrap(types.ErrInvalidMarketStatus, p.Status.String())
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
	if !types.IsHexHash(p.MarketId) {
		return errors.Wrap(types.ErrMarketInvalid, p.MarketId)
	}

	if p.SettlementPrice != nil && !SafeIsPositiveDec(*p.SettlementPrice) {
		return errors.Wrap(types.ErrInvalidSettlement, "settlement price must be positive for derivatives and nil for spot")
	}

	return govtypes.ValidateAbstract(p)
}

// NewUpdateDenomDecimalsProposal returns new instance of UpdateDenomDecimalsProposal
func NewUpdateDenomDecimalsProposal(
	title, description string,
	denomDecimals []*DenomDecimals,
) *UpdateAuctionExchangeTransferDenomDecimalsProposal {
	return &UpdateAuctionExchangeTransferDenomDecimalsProposal{
		Title:         title,
		Description:   description,
		DenomDecimals: denomDecimals,
	}
}

// Implements Proposal Interface
var _ govtypes.Content = &UpdateAuctionExchangeTransferDenomDecimalsProposal{}

// GetTitle returns the title of this proposal
func (p *UpdateAuctionExchangeTransferDenomDecimalsProposal) GetTitle() string {
	return p.Title
}

// GetDescription returns the description of this proposal
func (p *UpdateAuctionExchangeTransferDenomDecimalsProposal) GetDescription() string {
	return p.Description
}

// ProposalRoute returns router key of this proposal.
func (*UpdateAuctionExchangeTransferDenomDecimalsProposal) ProposalRoute() string { return RouterKey }

// ProposalType returns proposal type of this proposal.
func (*UpdateAuctionExchangeTransferDenomDecimalsProposal) ProposalType() string {
	return ProposalUpdateAuctionExchangeTransferDenomDecimals
}

// ValidateBasic returns ValidateBasic result of this proposal.
func (p *UpdateAuctionExchangeTransferDenomDecimalsProposal) ValidateBasic() error {
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

	if d.Decimals <= 0 || d.Decimals > uint64(types.MaxDecimals) {
		return errors.Wrapf(types.ErrInvalidDenomDecimal, "invalid decimals passed: %d", d.Decimals)
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
		return errors.Wrap(types.ErrInvalidOracle, "oracle base should not be empty")
	}
	if p.OracleQuote == "" {
		return errors.Wrap(types.ErrInvalidOracle, "oracle quote should not be empty")
	}
	if p.OracleBase == p.OracleQuote {
		return types.ErrSameOracles
	}
	switch p.OracleType {
	case oracletypes.OracleType_Band, oracletypes.OracleType_PriceFeed, oracletypes.OracleType_Coinbase,
		oracletypes.OracleType_Chainlink, oracletypes.OracleType_Razor, oracletypes.OracleType_Dia,
		oracletypes.OracleType_API3, oracletypes.OracleType_Uma, oracletypes.OracleType_Pyth,
		oracletypes.OracleType_BandIBC, oracletypes.OracleType_Provider, oracletypes.OracleType_Stork:

	default:
		return errors.Wrap(types.ErrInvalidOracleType, p.OracleType.String())
	}

	if p.OracleScaleFactor > types.MaxOracleScaleFactor {
		return types.ErrExceedsMaxOracleScaleFactor
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
		return errors.Wrap(types.ErrInvalidOracle, "oracle symbol should not be empty")
	}
	if p.Provider == "" {
		return errors.Wrap(types.ErrInvalidOracle, "oracle provider should not be empty")
	}

	if p.OracleType != oracletypes.OracleType_Provider {
		return errors.Wrap(types.ErrInvalidOracleType, p.OracleType.String())
	}

	if p.OracleScaleFactor > types.MaxOracleScaleFactor {
		return types.ErrExceedsMaxOracleScaleFactor
	}

	return nil
}

// NewPerpetualMarketLaunchProposal returns new instance of PerpetualMarketLaunchProposal
func NewPerpetualMarketLaunchProposal(
	title, description, ticker, quoteDenom,
	oracleBase, oracleQuote string, oracleScaleFactor uint32, oracleType oracletypes.OracleType,
	initialMarginRatio, maintenanceMarginRatio, reduceMarginRatio, makerFeeRate, takerFeeRate,
	minPriceTickSize, minQuantityTickSize, minNotional math.LegacyDec,
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
		ReduceMarginRatio:      reduceMarginRatio,
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
	if p.Ticker == "" || len(p.Ticker) > types.MaxTickerLength {
		return errors.Wrapf(types.ErrInvalidTicker, "ticker should not be empty or exceed %d characters", types.MaxTickerLength)
	}
	if p.QuoteDenom == "" {
		return errors.Wrap(types.ErrInvalidQuoteDenom, "quote denom should not be empty")
	}

	oracleParams := NewOracleParams(p.OracleBase, p.OracleQuote, p.OracleScaleFactor, p.OracleType)
	if err := oracleParams.ValidateBasic(); err != nil {
		return err
	}
	if err := types.ValidateMakerFee(p.MakerFeeRate); err != nil {
		return err
	}
	if err := types.ValidateFee(p.TakerFeeRate); err != nil {
		return err
	}
	if err := types.ValidateMarginRatio(p.InitialMarginRatio); err != nil {
		return err
	}
	if err := types.ValidateMarginRatio(p.MaintenanceMarginRatio); err != nil {
		return err
	}
	if err := types.ValidateMarginRatio(p.ReduceMarginRatio); err != nil {
		return err
	}
	if p.MakerFeeRate.GT(p.TakerFeeRate) {
		return types.ErrFeeRatesRelation
	}
	if p.InitialMarginRatio.LTE(p.MaintenanceMarginRatio) {
		return types.ErrMarginsRelation
	}
	if p.ReduceMarginRatio.LT(p.InitialMarginRatio) {
		return types.ErrMarginsRelation
	}

	if err := types.ValidateTickSize(p.MinPriceTickSize); err != nil {
		return errors.Wrap(types.ErrInvalidPriceTickSize, err.Error())
	}
	if err := types.ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return errors.Wrap(types.ErrInvalidQuantityTickSize, err.Error())
	}
	if err := types.ValidateMinNotional(p.MinNotional); err != nil {
		return errors.Wrap(types.ErrInvalidNotional, err.Error())
	}

	return govtypes.ValidateAbstract(p)
}

// NewExpiryFuturesMarketLaunchProposal returns new instance of ExpiryFuturesMarketLaunchProposal
func NewExpiryFuturesMarketLaunchProposal(
	title, description, ticker, quoteDenom,
	oracleBase, oracleQuote string, oracleScaleFactor uint32, oracleType oracletypes.OracleType, expiry int64,
	initialMarginRatio, maintenanceMarginRatio, reduceMarginRatio, makerFeeRate, takerFeeRate,
	minPriceTickSize, minQuantityTickSize, minNotional math.LegacyDec,
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
		ReduceMarginRatio:      reduceMarginRatio,
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
	if p.Ticker == "" || len(p.Ticker) > types.MaxTickerLength {
		return errors.Wrapf(types.ErrInvalidTicker, "ticker should not be empty or exceed %d characters", types.MaxTickerLength)
	}
	if p.QuoteDenom == "" {
		return errors.Wrap(types.ErrInvalidQuoteDenom, "quote denom should not be empty")
	}

	oracleParams := NewOracleParams(p.OracleBase, p.OracleQuote, p.OracleScaleFactor, p.OracleType)
	if err := oracleParams.ValidateBasic(); err != nil {
		return err
	}
	if p.Expiry <= 0 {
		return errors.Wrap(types.ErrInvalidExpiry, "expiry should not be empty")
	}
	if err := types.ValidateMakerFee(p.MakerFeeRate); err != nil {
		return err
	}
	if err := types.ValidateFee(p.TakerFeeRate); err != nil {
		return err
	}
	if err := types.ValidateMarginRatio(p.InitialMarginRatio); err != nil {
		return err
	}
	if err := types.ValidateMarginRatio(p.MaintenanceMarginRatio); err != nil {
		return err
	}
	if err := types.ValidateMarginRatio(p.ReduceMarginRatio); err != nil {
		return err
	}
	if p.MakerFeeRate.GT(p.TakerFeeRate) {
		return types.ErrFeeRatesRelation
	}
	if p.InitialMarginRatio.LTE(p.MaintenanceMarginRatio) {
		return types.ErrMarginsRelation
	}
	if p.ReduceMarginRatio.LT(p.InitialMarginRatio) {
		return types.ErrMarginsRelation
	}

	if err := types.ValidateTickSize(p.MinPriceTickSize); err != nil {
		return errors.Wrap(types.ErrInvalidPriceTickSize, err.Error())
	}
	if err := types.ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return errors.Wrap(types.ErrInvalidQuantityTickSize, err.Error())
	}
	if err := types.ValidateMinNotional(p.MinNotional); err != nil {
		return errors.Wrap(types.ErrInvalidNotional, err.Error())
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
			return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "campaign reward pool addition cannot be nil")
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
		return errors.Wrap(types.ErrInvalidTradingRewardsPendingPointsUpdate, "reward pending points update cannot be nil")
	}

	if p.PendingPoolTimestamp <= 0 {
		return errors.Wrap(types.ErrInvalidTradingRewardsPendingPointsUpdate, "pending pool timestamp cannot be zero")
	}

	accountAddresses := make([]string, 0)

	for _, rewardPointUpdate := range p.RewardPointUpdates {
		if rewardPointUpdate == nil {
			return errors.Wrap(types.ErrInvalidTradingRewardsPendingPointsUpdate, "reward pending point update cannot be nil")
		}

		_, err := sdk.AccAddressFromBech32(rewardPointUpdate.AccountAddress)

		if err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, rewardPointUpdate.AccountAddress)
		}

		accountAddresses = append(accountAddresses, rewardPointUpdate.AccountAddress)

		if rewardPointUpdate.NewPoints.IsNil() {
			return errors.Wrap(types.ErrInvalidTradingRewardsPendingPointsUpdate, "reward pending points cannot be nil")
		}

		if rewardPointUpdate.NewPoints.IsNegative() {
			return errors.Wrap(types.ErrInvalidTradingRewardsPendingPointsUpdate, "reward pending points cannot be negative")
		}
	}

	hasDuplicateAccountAddresses := types.HasDuplicates(accountAddresses)
	if hasDuplicateAccountAddresses {
		return errors.Wrap(types.ErrInvalidTradingRewardsPendingPointsUpdate, "account address cannot have duplicates")
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
		return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "new campaign info cannot be nil")
	}

	if len(p.CampaignRewardPools) == 0 {
		return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "new campaign reward pools cannot be nil")
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
	if err := t.validateMarketIdsAndMultipliers(); err != nil {
		return err
	}

	if err := t.validateMarketIds(); err != nil {
		return err
	}

	hasDuplicatesInMarkets := types.HasDuplicates(t.BoostedSpotMarketIds) || types.HasDuplicates(t.BoostedDerivativeMarketIds)
	if hasDuplicatesInMarkets {
		return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "campaign contains duplicate boosted market ids")
	}

	if err := t.validateSpotMultipliers(); err != nil {
		return err
	}

	if err := t.validateDerivativeMultipliers(); err != nil {
		return err
	}

	return nil
}

func (t *TradingRewardCampaignBoostInfo) validateMarketIdsAndMultipliers() error {
	if len(t.BoostedSpotMarketIds) != len(t.SpotMarketMultipliers) {
		return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "boosted spot market ids is not matching spot market multipliers")
	}

	if len(t.BoostedDerivativeMarketIds) != len(t.DerivativeMarketMultipliers) {
		return errors.Wrap(
			types.ErrInvalidTradingRewardCampaign,
			"boosted derivative market ids is not matching derivative market multipliers",
		)
	}

	return nil
}

func (t *TradingRewardCampaignBoostInfo) validateMarketIds() error {
	for _, marketID := range t.BoostedSpotMarketIds {
		if !types.IsHexHash(marketID) {
			return errors.Wrap(types.ErrMarketInvalid, marketID)
		}
	}

	for _, marketID := range t.BoostedDerivativeMarketIds {
		if !types.IsHexHash(marketID) {
			return errors.Wrap(types.ErrMarketInvalid, marketID)
		}
	}

	return nil
}

func (t *TradingRewardCampaignBoostInfo) validateSpotMultipliers() error {
	for _, multiplier := range t.SpotMarketMultipliers {
		if IsZeroOrNilDec(multiplier.MakerPointsMultiplier) {
			return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "spot market maker multiplier cannot be zero or nil")
		}

		if IsZeroOrNilDec(multiplier.TakerPointsMultiplier) {
			return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "spot market taker multiplier cannot be zero or nil")
		}

		if !SafeIsPositiveDec(multiplier.MakerPointsMultiplier) {
			return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "spot market maker multiplier cannot be negative")
		}

		if !SafeIsPositiveDec(multiplier.TakerPointsMultiplier) {
			return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "spot market taker multiplier cannot be negative")
		}
	}

	return nil
}

func (t *TradingRewardCampaignBoostInfo) validateDerivativeMultipliers() error {
	for _, multiplier := range t.DerivativeMarketMultipliers {
		if IsZeroOrNilDec(multiplier.MakerPointsMultiplier) {
			return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "derivative market maker multiplier cannot be zero or nil")
		}

		if IsZeroOrNilDec(multiplier.TakerPointsMultiplier) {
			return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "derivative market taker multiplier cannot be zero or nil")
		}

		if !SafeIsPositiveDec(multiplier.MakerPointsMultiplier) {
			return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "derivative market maker multiplier cannot be negative")
		}

		if !SafeIsPositiveDec(multiplier.TakerPointsMultiplier) {
			return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "derivative market taker multiplier cannot be negative")
		}
	}

	return nil
}

func (c *TradingRewardCampaignInfo) ValidateBasic() error {
	if c == nil {
		return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "campaign info cannot be nil")
	}

	if c.CampaignDurationSeconds <= 0 {
		return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "campaign duration cannot be zero")
	}

	if len(c.QuoteDenoms) == 0 {
		return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "campaign quote denoms cannot be nil")
	}

	hasTradingRewardBoostInfoDefined := c.TradingRewardBoostInfo != nil
	if hasTradingRewardBoostInfoDefined {
		if err := c.TradingRewardBoostInfo.ValidateBasic(); err != nil {
			return err
		}
	}

	for _, marketID := range c.DisqualifiedMarketIds {
		if !types.IsHexHash(marketID) {
			return errors.Wrap(types.ErrMarketInvalid, marketID)
		}
	}

	hasDuplicatesInDisqualifiedMarkets := types.HasDuplicates(c.DisqualifiedMarketIds)
	if hasDuplicatesInDisqualifiedMarkets {
		return errors.Wrap(types.ErrInvalidTradingRewardCampaign, "campaign contains duplicate disqualified market ids")
	}

	return nil
}

func validateCampaignRewardPool(pool *CampaignRewardPool, campaignDurationSeconds, prevStartTimestamp int64) (int64, error) {
	if pool == nil {
		return 0, errors.Wrap(types.ErrInvalidTradingRewardCampaign, "new campaign reward pool cannot be nil")
	}

	if pool.StartTimestamp <= prevStartTimestamp {
		return 0, errors.Wrap(types.ErrInvalidTradingRewardCampaign, "reward pool start timestamps must be in ascending order")
	}

	hasValidStartTimestamp := prevStartTimestamp == 0 || pool.StartTimestamp == (prevStartTimestamp+campaignDurationSeconds)
	if !hasValidStartTimestamp {
		return 0, errors.Wrap(types.ErrInvalidTradingRewardCampaign, "start timestamps not matching campaign duration")
	}

	prevStartTimestamp = pool.StartTimestamp

	hasDuplicatesInEpochRewards := types.HasDuplicatesCoin(pool.MaxCampaignRewards)
	if hasDuplicatesInEpochRewards {
		return 0, errors.Wrap(types.ErrInvalidTradingRewardCampaign, "reward pool campaign contains duplicate market coins")
	}

	for _, epochRewardDenom := range pool.MaxCampaignRewards {
		if !epochRewardDenom.IsValid() {
			return 0, errors.Wrap(sdkerrors.ErrInvalidCoins, epochRewardDenom.String())
		}

		if IsZeroOrNilInt(epochRewardDenom.Amount) {
			return 0, errors.Wrap(types.ErrInvalidTradingRewardCampaign, "reward pool contains zero or nil reward amount")
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
		return errors.Wrap(types.ErrInvalidFeeDiscountSchedule, "new fee discount schedule cannot be nil")
	}

	if p.Schedule.BucketCount < 2 {
		return errors.Wrap(types.ErrInvalidFeeDiscountSchedule, "new fee discount schedule must have at least 2 buckets")
	}

	if p.Schedule.BucketDuration < 10 {
		return errors.Wrap(
			types.ErrInvalidFeeDiscountSchedule,
			"new fee discount schedule must have have bucket durations of at least 10 seconds",
		)
	}

	if types.HasDuplicates(p.Schedule.QuoteDenoms) {
		return errors.Wrap(types.ErrInvalidFeeDiscountSchedule, "new fee discount schedule cannot have duplicate quote denoms")
	}

	for _, marketID := range p.Schedule.DisqualifiedMarketIds {
		if !types.IsHexHash(marketID) {
			return errors.Wrap(types.ErrMarketInvalid, marketID)
		}
	}

	if types.HasDuplicates(p.Schedule.DisqualifiedMarketIds) {
		return errors.Wrap(types.ErrInvalidFeeDiscountSchedule, "new fee discount schedule cannot have duplicate disqualified market ids")
	}

	if len(p.Schedule.TierInfos) < 1 {
		return errors.Wrap(types.ErrInvalidFeeDiscountSchedule, "new fee discount schedule must have at least one discount tier")
	}

	for idx, tierInfo := range p.Schedule.TierInfos {
		if err := tierInfo.ValidateBasic(); err != nil {
			return err
		}

		if idx > 0 {
			prevTierInfo := p.Schedule.TierInfos[idx-1]

			if prevTierInfo.MakerDiscountRate.GT(tierInfo.MakerDiscountRate) {
				return errors.Wrap(
					types.ErrInvalidFeeDiscountSchedule,
					"successive MakerDiscountRates must be equal or larger than those of lower tiers",
				)
			}

			if prevTierInfo.TakerDiscountRate.GT(tierInfo.TakerDiscountRate) {
				return errors.Wrap(
					types.ErrInvalidFeeDiscountSchedule,
					"successive TakerDiscountRates must be equal or larger than those of lower tiers",
				)
			}

			if prevTierInfo.StakedAmount.GT(tierInfo.StakedAmount) {
				return errors.Wrap(
					types.ErrInvalidFeeDiscountSchedule,
					"successive StakedAmount must be equal or larger than those of lower tiers",
				)
			}

			if prevTierInfo.Volume.GT(tierInfo.Volume) {
				return errors.Wrap(
					types.ErrInvalidFeeDiscountSchedule,
					"successive Volume must be equal or larger than those of lower tiers",
				)
			}
		}
	}

	return govtypes.ValidateAbstract(p)
}

func (t *FeeDiscountTierInfo) ValidateBasic() error {
	if !SafeIsNonNegativeDec(t.MakerDiscountRate) || t.MakerDiscountRate.GT(math.LegacyOneDec()) {
		return errors.Wrap(types.ErrInvalidFeeDiscountSchedule, "MakerDiscountRate must be between 0 and 1")
	}

	if !SafeIsNonNegativeDec(t.TakerDiscountRate) || t.TakerDiscountRate.GT(math.LegacyOneDec()) {
		return errors.Wrap(types.ErrInvalidFeeDiscountSchedule, "TakerDiscountRate must be between 0 and 1")
	}

	if !SafeIsPositiveInt(t.StakedAmount) {
		return errors.Wrap(types.ErrInvalidFeeDiscountSchedule, "StakedAmount must be non-negative")
	}

	if !SafeIsPositiveDec(t.Volume) {
		return errors.Wrap(types.ErrInvalidFeeDiscountSchedule, "Volume must be non-negative")
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
	if p.Ticker == "" || len(p.Ticker) > types.MaxTickerLength {
		return errors.Wrapf(types.ErrInvalidTicker, "ticker should not be empty or exceed %d characters", types.MaxTickerLength)
	}
	if p.OracleSymbol == "" {
		return errors.Wrap(types.ErrInvalidOracle, "oracle symbol should not be empty")
	}
	if p.OracleProvider == "" {
		return errors.Wrap(types.ErrInvalidOracle, "oracle provider should not be empty")
	}
	if p.OracleType != oracletypes.OracleType_Provider {
		return errors.Wrap(types.ErrInvalidOracleType, p.OracleType.String())
	}
	if p.OracleScaleFactor > types.MaxOracleScaleFactor {
		return types.ErrExceedsMaxOracleScaleFactor
	}

	if p.ExpirationTimestamp >= p.SettlementTimestamp || p.ExpirationTimestamp < 0 || p.SettlementTimestamp < 0 {
		return types.ErrInvalidExpiry
	}

	if p.Admin != "" {
		_, err := sdk.AccAddressFromBech32(p.Admin)
		if err != nil {
			return errors.Wrap(sdkerrors.ErrInvalidAddress, p.Admin)
		}
	}
	if p.QuoteDenom == "" {
		return errors.Wrap(types.ErrInvalidQuoteDenom, "quote denom should not be empty")
	}
	if err := types.ValidateMakerFee(p.MakerFeeRate); err != nil {
		return err
	}
	if err := types.ValidateFee(p.TakerFeeRate); err != nil {
		return err
	}

	if p.MakerFeeRate.GT(p.TakerFeeRate) {
		return types.ErrFeeRatesRelation
	}

	if err := types.ValidateTickSize(p.MinPriceTickSize); err != nil {
		return errors.Wrap(types.ErrInvalidPriceTickSize, err.Error())
	}
	if err := types.ValidateTickSize(p.MinQuantityTickSize); err != nil {
		return errors.Wrap(types.ErrInvalidQuantityTickSize, err.Error())
	}
	if err := types.ValidateMinNotional(p.MinNotional); err != nil {
		return errors.Wrap(types.ErrInvalidNotional, err.Error())
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
	if !types.IsHexHash(p.MarketId) {
		return errors.Wrap(types.ErrMarketInvalid, p.MarketId)
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
		if err := types.ValidateMakerFee(*p.MakerFeeRate); err != nil {
			return err
		}
	}
	if p.TakerFeeRate != nil {
		if err := types.ValidateFee(*p.TakerFeeRate); err != nil {
			return err
		}
	}

	if p.RelayerFeeShareRate != nil {
		if err := types.ValidateFee(*p.RelayerFeeShareRate); err != nil {
			return err
		}
	}

	if p.MinPriceTickSize != nil {
		if err := types.ValidateTickSize(*p.MinPriceTickSize); err != nil {
			return errors.Wrap(types.ErrInvalidPriceTickSize, err.Error())
		}
	}

	if p.MinQuantityTickSize != nil {
		if err := types.ValidateTickSize(*p.MinQuantityTickSize); err != nil {
			return errors.Wrap(types.ErrInvalidQuantityTickSize, err.Error())
		}
	}

	if p.MinNotional != nil {
		if err := types.ValidateMinNotional(*p.MinNotional); err != nil {
			return errors.Wrap(types.ErrInvalidNotional, err.Error())
		}
	}

	if p.ExpirationTimestamp != 0 && p.SettlementTimestamp != 0 {
		if p.ExpirationTimestamp >= p.SettlementTimestamp || p.ExpirationTimestamp < 0 || p.SettlementTimestamp < 0 {
			return types.ErrInvalidExpiry
		}
	}

	if p.SettlementTimestamp < 0 {
		return types.ErrInvalidSettlement
	}

	if p.Admin != "" {
		if _, err := sdk.AccAddressFromBech32(p.Admin); err != nil {
			return err
		}
	}

	if len(p.Ticker) > types.MaxTickerLength {
		return errors.Wrapf(types.ErrInvalidTicker, "ticker should not exceed %d characters", types.MaxTickerLength)
	}

	// price is either nil (not set), -1 (demolish with refund) or [0..1] (demolish with settle)
	switch {
	case p.SettlementPrice == nil,
		p.SettlementPrice.IsNil():
		// ok
	case p.SettlementPrice.Equal(BinaryOptionsMarketRefundFlagPrice),
		p.SettlementPrice.GTE(math.LegacyZeroDec()) && p.SettlementPrice.LTE(types.MaxBinaryOptionsOrderPrice):
		if p.Status != MarketStatus_Demolished {
			return errors.Wrapf(
				types.ErrInvalidMarketStatus,
				"status should be set to demolished when the settlement price is set, status: %s",
				p.Status.String(),
			)
		}
		// ok
	default:
		return errors.Wrap(types.ErrInvalidPrice, p.SettlementPrice.String())
	}

	switch p.Status {
	case
		MarketStatus_Unspecified,
		MarketStatus_Demolished:
	default:
		return errors.Wrap(types.ErrInvalidMarketStatus, p.Status.String())
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
		if !types.IsHexHash(m.MarketId) {
			return errors.Wrap(types.ErrMarketInvalid, m.MarketId)
		}
		multiplier := m.FeeMultiplier
		if multiplier.IsNil() {
			return fmt.Errorf("atomic taker fee multiplier cannot be nil: %v", multiplier)
		}

		if multiplier.LT(math.LegacyOneDec()) {
			return fmt.Errorf("atomic taker fee multiplier cannot be less than 1: %v", multiplier)
		}

		if multiplier.GT(types.MaxFeeMultiplier) {
			return fmt.Errorf("atomicMarketOrderFeeMultiplier cannot be bigger than %v: %v", multiplier, types.MaxFeeMultiplier)
		}
	}
	return govtypes.ValidateAbstract(p)
}

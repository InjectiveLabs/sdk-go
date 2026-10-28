package types

import "errors"

func (gs GenesisState) Validate() error {
	if len(gs.BinaryOptionsMarkets) != 0 || len(gs.BinaryOptionsMarketIdsScheduledForSettlement) != 0 {
		return errors.New("binary-options genesis state is no longer supported")
	}
	if len(gs.HistoricalTradeRecords) != 0 {
		return errors.New("historical trades cannot be imported after removal")
	}
	if gs.TradingRewardCampaignInfo != nil ||
		len(gs.TradingRewardPoolCampaignSchedule) != 0 ||
		len(gs.TradingRewardCampaignAccountPoints) != 0 ||
		len(gs.PendingTradingRewardPoolCampaignSchedule) != 0 ||
		len(gs.PendingTradingRewardCampaignAccountPoints) != 0 ||
		len(gs.RewardsOptOutAddresses) != 0 {
		return errors.New("retired Trade&Earn state cannot be imported after removal")
	}

	if err := gs.ValidateRetiredOracleGenesis(); err != nil {
		return err
	}
	// TODO: validate stuff in genesis
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}

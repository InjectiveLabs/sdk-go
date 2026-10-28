package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	oracletypes "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
)

// ValidateRetiredOracleGenesis retains terminal metadata while rejecting any
// imported obligations that would require an oracle removed in v1.21.
//
//nolint:revive // Keep each historical genesis collection prerequisite together for auditability.
func (gs GenesisState) ValidateRetiredOracleGenesis() error {
	// Genesis import keys these records by decoded hash, including references
	// with different hex casing or without the optional 0x prefix.
	retired := make(map[common.Hash]bool)
	for _, market := range gs.DerivativeMarkets {
		if market == nil || !oracletypes.IsRetiredOracle(market.OracleType) {
			continue
		}
		if market.Status != MarketStatus_Demolished {
			return fmt.Errorf("genesis market %s uses retired %s oracle and must be demolished", market.MarketId, market.OracleType)
		}
		if !IsHexHash(market.MarketId) {
			return fmt.Errorf("genesis market with retired oracle has invalid market_id %q", market.MarketId)
		}
		retired[common.HexToHash(market.MarketId)] = true
	}
	for _, position := range gs.Positions {
		if retired[common.HexToHash(position.MarketId)] {
			return fmt.Errorf("genesis market %s with retired oracle has positions", position.MarketId)
		}
	}
	for _, book := range gs.DerivativeOrderbook {
		if retired[common.HexToHash(book.MarketId)] && len(book.Orders) > 0 {
			return fmt.Errorf("genesis market %s with retired oracle has orders", book.MarketId)
		}
	}
	for _, book := range gs.ConditionalDerivativeOrderbooks {
		if book != nil && retired[common.HexToHash(book.MarketId)] &&
			(len(book.LimitBuyOrders) > 0 || len(book.LimitSellOrders) > 0 || len(book.MarketBuyOrders) > 0 || len(book.MarketSellOrders) > 0) {
			return fmt.Errorf("genesis market %s with retired oracle has conditional orders", book.MarketId)
		}
	}
	for _, settlement := range gs.DerivativeMarketSettlementScheduled {
		if retired[common.HexToHash(settlement.MarketId)] {
			return fmt.Errorf("genesis market %s with retired oracle has scheduled settlement", settlement.MarketId)
		}
	}
	// Import rebuilds the natural-expiry timestamp queue even for demolished
	// markets, so these records must not resume against a retired oracle.
	for _, expiry := range gs.ExpiryFuturesMarketInfoState {
		if retired[common.HexToHash(expiry.MarketId)] {
			return fmt.Errorf("genesis market %s with retired oracle has expiry settlement metadata", expiry.MarketId)
		}
	}
	return nil
}

package v2

import (
	"fmt"
	"math/big"

	"cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
)

// CrossMarginGenesisSpotHoldBudgets returns the final total balance minus
// effective-ISOLATED derivative and binary-options order reservations for each
// cross-exposed owner and denom. A missing balance is zero; duplicate balance
// records retain InitGenesis's last-write-wins semantics. Underfunded isolated
// orders cannot be repaired by pruning spot orders and are rejected here.
//
//nolint:revive // Keep the persistent order classes and deterministic validation order together.
func (gs GenesisState) CrossMarginGenesisSpotHoldBudgets(
	membership GenesisCrossMarginMembership,
) (map[common.Hash]map[string]*big.Int, error) {
	budgets := make(map[common.Hash]map[string]*big.Int)
	for i, balance := range gs.Balances {
		subaccountID := common.HexToHash(balance.SubaccountId)
		if !membership.HasCrossMarginExposure(subaccountID) {
			continue
		}
		if !isWellFormedGenesisBalance(balance) || balance.Deposits.TotalBalance.IsNegative() {
			return nil, fmt.Errorf("balances[%d]: invalid cross-margin total balance", i)
		}
		if budgets[subaccountID] == nil {
			budgets[subaccountID] = make(map[string]*big.Int)
		}
		budgets[subaccountID][balance.Denom] = new(big.Int).Set(balance.Deposits.TotalBalance.BigInt())
	}

	markets := make(map[common.Hash]DerivativeMarketI, len(gs.DerivativeMarkets)+len(gs.BinaryOptionsMarkets))
	for _, market := range gs.DerivativeMarkets {
		if market != nil {
			markets[market.MarketID()] = market
		}
	}
	for _, market := range gs.BinaryOptionsMarkets {
		if market != nil {
			markets[market.MarketID()] = market
		}
	}
	isolatedMarket := func(subaccountID, marketID common.Hash, field string) (DerivativeMarketI, error) {
		if !membership.HasCrossMarginExposure(subaccountID) || membership.IsCrossMarginMarket(subaccountID, marketID) {
			return nil, nil
		}
		market := markets[marketID]
		if market == nil || market.GetQuoteDenom() == "" {
			return nil, fmt.Errorf("%s: unknown isolated derivative market %s", field, marketID.Hex())
		}
		return market, nil
	}
	reserve := func(subaccountID common.Hash, denom, field string, hold math.LegacyDec) error {
		if budgets[subaccountID] == nil {
			budgets[subaccountID] = make(map[string]*big.Int)
		}
		if budgets[subaccountID][denom] == nil {
			budgets[subaccountID][denom] = new(big.Int)
		}
		remaining := budgets[subaccountID][denom]
		remaining.Sub(remaining, hold.BigInt())
		if remaining.Sign() < 0 {
			return fmt.Errorf("%s: isolated derivative holds exceed total balance for subaccount %s denom %s", field, subaccountID.Hex(), denom)
		}
		return nil
	}
	reserveLimit := func(marketID common.Hash, order *DerivativeLimitOrder, conditional bool, field string) error {
		if order == nil {
			return fmt.Errorf("%s: missing order", field)
		}
		market, err := isolatedMarket(order.SubaccountID(), marketID, field)
		if err != nil || market == nil {
			return err
		}
		if order.Margin.IsNil() || order.Margin.IsNegative() || !order.Margin.IsInValidRange() {
			return fmt.Errorf("%s: invalid isolated order margin", field)
		}
		hold := math.LegacyZeroDec()
		if order.IsVanilla() {
			feeRate := market.GetMakerFeeRate()
			if conditional && !order.OrderType.IsPostOnly() {
				feeRate = market.GetTakerFeeRate()
			}
			hold, err = CanonicalIsolatedLimitHold(order, market, feeRate)
			if err != nil {
				return fmt.Errorf("%s: %w", field, err)
			}
		}
		return reserve(order.SubaccountID(), market.GetQuoteDenom(), field, hold)
	}
	for i, orderbook := range gs.DerivativeOrderbook {
		for j, order := range orderbook.Orders {
			field := fmt.Sprintf("derivative_orderbook[%d].orders[%d]", i, j)
			if err := reserveLimit(common.HexToHash(orderbook.MarketId), order, false, field); err != nil {
				return nil, err
			}
		}
	}
	for i, orderbook := range gs.ConditionalDerivativeOrderbooks {
		if orderbook == nil {
			return nil, fmt.Errorf("conditional_derivative_orderbooks[%d]: missing orderbook", i)
		}
		marketID := common.HexToHash(orderbook.MarketId)
		for side, orders := range [][]*DerivativeLimitOrder{orderbook.LimitBuyOrders, orderbook.LimitSellOrders} {
			for j, order := range orders {
				field := fmt.Sprintf("conditional_derivative_orderbooks[%d].limit_orders[%d][%d]", i, side, j)
				if err := reserveLimit(marketID, order, true, field); err != nil {
					return nil, err
				}
			}
		}
		for side, orders := range [][]*DerivativeMarketOrder{orderbook.MarketBuyOrders, orderbook.MarketSellOrders} {
			for j, order := range orders {
				field := fmt.Sprintf("conditional_derivative_orderbooks[%d].market_orders[%d][%d]", i, side, j)
				if order == nil {
					return nil, fmt.Errorf("%s: missing order", field)
				}
				market, err := isolatedMarket(order.SubaccountID(), marketID, field)
				if err != nil {
					return nil, err
				}
				if market == nil {
					continue
				}
				if order.Margin.IsNil() || order.Margin.IsNegative() || !order.Margin.IsInValidRange() {
					return nil, fmt.Errorf("%s: invalid isolated order margin", field)
				}
				hold, err := CanonicalIsolatedMarketHold(order.GetCancelRefundAmount(), market)
				if err != nil {
					return nil, fmt.Errorf("%s: %w", field, err)
				}
				if err := reserve(order.SubaccountID(), market.GetQuoteDenom(), field, hold); err != nil {
					return nil, err
				}
			}
		}
	}
	return budgets, nil
}

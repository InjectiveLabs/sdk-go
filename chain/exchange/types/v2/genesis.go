package v2

import (
	"bytes"
	"errors"
	"fmt"
	stdmath "math"
	"math/big"
	"slices"

	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

func NewGenesisState() GenesisState {
	return GenesisState{}
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:                       DefaultParams(),
		IsSpotExchangeEnabled:        true,
		IsDerivativesExchangeEnabled: true,
	}
}

func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	routers, err := ParseCrossMarginRFQRouterSet(gs.Params.CrossMarginParams.LiquidationRfqContractAddress)
	if err != nil {
		return err
	}
	if err := gs.validateReservedRFQInventories(routers); err != nil {
		return err
	}

	if err := gs.validateSpotOrderbookMarkets(); err != nil {
		return err
	}

	if err := gs.validateDerivativeMarketSettlementScheduled(); err != nil {
		return err
	}

	membership, err := BuildGenesisCrossMarginMembership(gs)
	if err != nil {
		return err
	}
	if err := gs.validateCrossMarginGenesisState(membership); err != nil {
		return err
	}

	return nil
}

type genesisMarketRiskModeKey struct {
	subaccountID common.Hash
	marketID     common.Hash
}

// GenesisCrossMarginMembership resolves the two related cross-margin concepts
// needed during genesis validation and import. HasCrossMarginExposure is the
// subaccount-wide predicate used by spot and balance accounting; IsCrossMarginMarket
// resolves derivative pool membership using binary-options pinning, then an
// explicit per-market override, then the subaccount profile.
type GenesisCrossMarginMembership struct {
	crossExposedSubaccounts map[common.Hash]struct{}
	profileModes            map[common.Hash]RiskMode
	marketModes             map[genesisMarketRiskModeKey]RiskMode
	binaryMarketIDs         map[common.Hash]struct{}
}

// BuildGenesisCrossMarginMembership validates the profile and per-market mode
// records and constructs the effective membership used by every genesis path.
func BuildGenesisCrossMarginMembership(gs GenesisState) (GenesisCrossMarginMembership, error) {
	membership := GenesisCrossMarginMembership{
		crossExposedSubaccounts: make(map[common.Hash]struct{}, len(gs.SubaccountRiskProfiles)+len(gs.SubaccountMarketRiskModes)),
		profileModes:            make(map[common.Hash]RiskMode, len(gs.SubaccountRiskProfiles)),
		marketModes:             make(map[genesisMarketRiskModeKey]RiskMode, len(gs.SubaccountMarketRiskModes)),
		binaryMarketIDs:         make(map[common.Hash]struct{}, len(gs.BinaryOptionsMarkets)),
	}
	routers, err := ParseCrossMarginRFQRouterSet(gs.Params.CrossMarginParams.LiquidationRfqContractAddress)
	if err != nil {
		return GenesisCrossMarginMembership{}, err
	}

	seenProfiles := make(map[common.Hash]struct{}, len(gs.SubaccountRiskProfiles))
	for i, record := range gs.SubaccountRiskProfiles {
		if err := validateRiskProfileRecord(record); err != nil {
			return GenesisCrossMarginMembership{}, fmt.Errorf("subaccount_risk_profiles[%d]: %w", i, err)
		}
		mode, _ := validateRiskMode(&record.RiskProfile)
		subaccountID := common.HexToHash(record.SubaccountId)
		if _, duplicate := seenProfiles[subaccountID]; duplicate {
			return GenesisCrossMarginMembership{}, fmt.Errorf(
				"subaccount_risk_profiles[%d]: duplicate subaccount_id %s", i, subaccountID.Hex(),
			)
		}
		seenProfiles[subaccountID] = struct{}{}
		membership.profileModes[subaccountID] = mode
		if mode != RiskMode_RISK_MODE_CROSS {
			continue
		}
		if routers.Len() == 0 {
			return GenesisCrossMarginMembership{}, fmt.Errorf(
				"subaccount_risk_profiles[%d]: cross-margin profile requires at least one liquidation RFQ router", i,
			)
		}
		if routers.Len() == 1 && routers.Contains(types.SubaccountIDToSdkAddress(subaccountID)) {
			return GenesisCrossMarginMembership{}, fmt.Errorf(
				"subaccount_risk_profiles[%d]: singleton RFQ router cannot own a cross-margin subaccount", i,
			)
		}
		membership.crossExposedSubaccounts[subaccountID] = struct{}{}
	}

	if err := gs.validateSubaccountMarketRiskModeRecords(); err != nil {
		return GenesisCrossMarginMembership{}, err
	}
	for i, record := range gs.SubaccountMarketRiskModes {
		subaccountID := common.HexToHash(record.SubaccountId)
		membership.marketModes[genesisMarketRiskModeKey{
			subaccountID: subaccountID,
			marketID:     common.HexToHash(record.MarketId),
		}] = record.Mode
		if record.Mode != RiskMode_RISK_MODE_CROSS {
			continue
		}
		if routers.Len() == 0 {
			return GenesisCrossMarginMembership{}, fmt.Errorf(
				"subaccount_market_risk_modes[%d]: cross-margin override requires at least one liquidation RFQ router", i,
			)
		}
		if routers.Len() == 1 && routers.Contains(types.SubaccountIDToSdkAddress(subaccountID)) {
			return GenesisCrossMarginMembership{}, fmt.Errorf(
				"subaccount_market_risk_modes[%d]: singleton RFQ router cannot own a cross-margin subaccount", i,
			)
		}
		membership.crossExposedSubaccounts[subaccountID] = struct{}{}
	}

	for _, market := range gs.BinaryOptionsMarkets {
		if market != nil {
			membership.binaryMarketIDs[common.HexToHash(market.MarketId)] = struct{}{}
		}
	}

	return membership, nil
}

// HasCrossMarginExposure reports whether genesis state gives the subaccount any
// cross-margin pool semantics through its profile or a CROSS override.
func (m GenesisCrossMarginMembership) HasCrossMarginExposure(subaccountID common.Hash) bool {
	_, ok := m.crossExposedSubaccounts[subaccountID]
	return ok
}

// IsCrossMarginMarket reports whether a derivative market is an effective
// cross-margin member for the subaccount in genesis state.
func (m GenesisCrossMarginMembership) IsCrossMarginMarket(subaccountID, marketID common.Hash) bool {
	if _, isBinary := m.binaryMarketIDs[marketID]; isBinary {
		return false
	}
	if mode, hasOverride := m.marketModes[genesisMarketRiskModeKey{
		subaccountID: subaccountID,
		marketID:     marketID,
	}]; hasOverride {
		return mode == RiskMode_RISK_MODE_CROSS
	}
	return m.profileModes[subaccountID] == RiskMode_RISK_MODE_CROSS
}

// CrossMarginSubaccountIDs returns all cross-exposed subaccounts in byte-lexicographic order.
func (m GenesisCrossMarginMembership) CrossMarginSubaccountIDs() []common.Hash {
	subaccountIDs := make([]common.Hash, 0, len(m.crossExposedSubaccounts))
	for subaccountID := range m.crossExposedSubaccounts {
		subaccountIDs = append(subaccountIDs, subaccountID)
	}
	slices.SortFunc(subaccountIDs, func(a, b common.Hash) int {
		return bytes.Compare(a.Bytes(), b.Bytes())
	})
	return subaccountIDs
}

// ValidateCrossMarginGenesisState validates cross-margin balances, persistent
// order identity, spot holds, the all-carrier terminal G/U/C/B bounds, and
// pool-specific position state using the supplied effective genesis membership.
// InitGenesis uses this before persisting state so direct imports receive the
// same cross-margin validation as GenesisState.Validate.
func ValidateCrossMarginGenesisState(gs GenesisState, membership GenesisCrossMarginMembership) error {
	return gs.validateCrossMarginGenesisState(membership)
}

func (gs GenesisState) validateSubaccountMarketRiskModes() error {
	return gs.validateSubaccountMarketRiskModeRecords()
}

func (gs GenesisState) validateSubaccountMarketRiskModeRecords() error {
	derivativeMarkets := gs.derivativeMarketIDs()
	seen := make(map[string]struct{}, len(gs.SubaccountMarketRiskModes))

	for i, record := range gs.SubaccountMarketRiskModes {
		if err := validateMarketRiskModeRecord(record, derivativeMarkets); err != nil {
			return fmt.Errorf("subaccount_market_risk_modes[%d]: %w", i, err)
		}

		// Normalize before keying: IsHexHash accepts mixed-case hex, and the
		// store keys on the decoded bytes, so mixed-case aliases of one
		// (subaccount, market) pair must collide here rather than slip through.
		pairKey := common.HexToHash(record.SubaccountId).Hex() + common.HexToHash(record.MarketId).Hex()
		if _, ok := seen[pairKey]; ok {
			return fmt.Errorf(
				"subaccount_market_risk_modes[%d]: duplicate record for subaccount %s market %s",
				i, record.SubaccountId, record.MarketId,
			)
		}
		seen[pairKey] = struct{}{}
	}

	return nil
}

func validateMarketRiskModeRecord(record *SubaccountMarketRiskModeRecord, derivativeMarkets map[string]struct{}) error {
	if err := ValidateSubaccountMarketRiskModeRecordShape(record); err != nil {
		return err
	}

	if _, ok := derivativeMarkets[common.HexToHash(record.MarketId).Hex()]; !ok {
		return fmt.Errorf("unknown derivative market_id %s", record.MarketId)
	}

	return nil
}

// ValidateSubaccountMarketRiskModeRecordShape runs every stateless per-record
// check except market existence, which callers verify against their own
// market source (the genesis record set or the live store). InitGenesis
// re-runs these checks so a direct import — which bypasses stateless genesis
// validation — cannot silently accept malformed IDs or an UNSPECIFIED mode
// (which the setter would treat as a delete).
func ValidateSubaccountMarketRiskModeRecordShape(record *SubaccountMarketRiskModeRecord) error {
	if record == nil {
		return errors.New("nil record")
	}

	if _, ok := types.IsValidSubaccountID(record.SubaccountId); !ok {
		return fmt.Errorf("invalid subaccount_id %q: must be a 32-byte hex hash", record.SubaccountId)
	}

	if !types.IsHexHash(record.MarketId) {
		return fmt.Errorf("invalid market_id %q: must be a 32-byte hex hash", record.MarketId)
	}

	if record.Mode != RiskMode_RISK_MODE_ISOLATED && record.Mode != RiskMode_RISK_MODE_CROSS {
		return fmt.Errorf("unsupported risk mode %v", record.Mode)
	}

	if record.Mode == RiskMode_RISK_MODE_CROSS && types.IsDefaultSubaccountID(common.HexToHash(record.SubaccountId)) {
		return fmt.Errorf("default subaccount %s cannot use cross-margin mode", record.SubaccountId)
	}

	return nil
}

func (gs GenesisState) derivativeMarketIDs() map[string]struct{} {
	markets := make(map[string]struct{}, len(gs.DerivativeMarkets))
	for _, market := range gs.DerivativeMarkets {
		if market == nil {
			continue
		}
		markets[common.HexToHash(market.MarketId).Hex()] = struct{}{}
	}

	return markets
}

type genesisCrossPoolKey struct {
	subaccountID common.Hash
	quoteDenom   string
}

type genesisCrossMarketState struct {
	positionRaw *big.Int
	queuedRaw   *big.Int
}

type genesisSpotHold struct {
	marketID     common.Hash
	isBuy        bool
	orderHash    common.Hash
	lockingDenom string
	holdRaw      *big.Int
}

// validateCrossMarginGenesisState validates persistent order identity and the
// post-repair B/C/U/G terminal-cleanup state. Those bounds intentionally count
// isolated and binary carriers too; effective membership only selects positions
// that require cross-pool RFQ representability. Underfunded CM spot orders are
// simulated using the same canonical hold helper and deterministic identity
// ordering used by InitGenesis; custom historical state may therefore be
// repaired instead of being rejected merely because its exported Available
// value is stale.
func (gs GenesisState) validateCrossMarginGenesisState(membership GenesisCrossMarginMembership) error { //nolint:revive // genesis validation is intentionally whole-state
	if err := gs.validateGenesisBalances(membership); err != nil {
		return err
	}

	derivativeMarkets := make(map[common.Hash]DerivativeMarketI, len(gs.DerivativeMarkets)+len(gs.BinaryOptionsMarkets))
	for _, market := range gs.DerivativeMarkets {
		if market != nil {
			derivativeMarkets[common.HexToHash(market.MarketId)] = market
		}
	}
	for _, market := range gs.BinaryOptionsMarkets {
		if market != nil {
			derivativeMarkets[common.HexToHash(market.MarketId)] = market
		}
	}
	spotMarkets := make(map[common.Hash]*SpotMarket, len(gs.SpotMarkets))
	for _, market := range gs.SpotMarkets {
		if market != nil {
			spotMarkets[common.HexToHash(market.MarketId)] = market
		}
	}

	seenHashes := make(map[common.Hash]string)
	seenCIDs := make(map[common.Hash]map[string]string)
	registerIdentity := func(field string, subaccountID, orderHash common.Hash, hashBytes []byte, cid string) error {
		if len(hashBytes) != common.HashLength || orderHash == (common.Hash{}) {
			return fmt.Errorf("%s: malformed order hash", field)
		}
		if previous, duplicate := seenHashes[orderHash]; duplicate {
			return fmt.Errorf("%s: duplicate order hash %s previously used by %s", field, orderHash.Hex(), previous)
		}
		seenHashes[orderHash] = field
		if cid == "" {
			return nil
		}
		if seenCIDs[subaccountID] == nil {
			seenCIDs[subaccountID] = make(map[string]string)
		}
		if previous, duplicate := seenCIDs[subaccountID][cid]; duplicate {
			return fmt.Errorf("%s: duplicate cid %q previously used by %s", field, cid, previous)
		}
		seenCIDs[subaccountID][cid] = field
		return nil
	}

	carriers := make(map[common.Hash]map[common.Hash]struct{})
	nonBinaryByPool := make(map[genesisCrossPoolKey]map[common.Hash]struct{})
	derivativeOrderCount := make(map[genesisCrossPoolKey]uint64)
	marketStates := make(map[common.Hash]map[common.Hash]*genesisCrossMarketState)
	quoteDecimalsByPool := make(map[genesisCrossPoolKey]uint32)
	ensureCarrier := func(subaccountID, marketID common.Hash, field string) (DerivativeMarketI, error) {
		market := derivativeMarkets[marketID]
		if market == nil || market.GetQuoteDenom() == "" {
			return nil, fmt.Errorf("%s: unknown derivative market %s", field, marketID.Hex())
		}
		if carriers[subaccountID] == nil {
			carriers[subaccountID] = make(map[common.Hash]struct{})
		}
		carriers[subaccountID][marketID] = struct{}{}
		if uint32(len(carriers[subaccountID])) > MaxCrossMarginDerivativeCarriers {
			return nil, fmt.Errorf("%s: cross-margin G bound exceeded", field)
		}
		if !market.GetMarketType().IsBinaryOptions() {
			poolKey := genesisCrossPoolKey{subaccountID: subaccountID, quoteDenom: market.GetQuoteDenom()}
			if nonBinaryByPool[poolKey] == nil {
				nonBinaryByPool[poolKey] = make(map[common.Hash]struct{})
			}
			nonBinaryByPool[poolKey][marketID] = struct{}{}
			if uint32(len(nonBinaryByPool[poolKey])) > MaxCrossMarginDerivativeCarriersPerQuote {
				return nil, fmt.Errorf("%s: cross-margin U bound exceeded for %s", field, market.GetQuoteDenom())
			}
		}
		return market, nil
	}
	stateFor := func(subaccountID, marketID common.Hash) *genesisCrossMarketState {
		if marketStates[subaccountID] == nil {
			marketStates[subaccountID] = make(map[common.Hash]*genesisCrossMarketState)
		}
		state := marketStates[subaccountID][marketID]
		if state == nil {
			state = &genesisCrossMarketState{positionRaw: new(big.Int), queuedRaw: new(big.Int)}
			marketStates[subaccountID][marketID] = state
		}
		return state
	}

	lastPosition := make(map[[2]common.Hash]int, len(gs.Positions))
	for i, record := range gs.Positions {
		if record.Position == nil || record.Position.Quantity.IsNil() {
			return fmt.Errorf("positions[%d]: position is incomplete", i)
		}
		lastPosition[[2]common.Hash{common.HexToHash(record.SubaccountId), common.HexToHash(record.MarketId)}] = i
	}
	for i, record := range gs.Positions {
		subaccountID := common.HexToHash(record.SubaccountId)
		marketID := common.HexToHash(record.MarketId)
		if !membership.HasCrossMarginExposure(subaccountID) ||
			lastPosition[[2]common.Hash{subaccountID, marketID}] != i ||
			record.Position.Quantity.IsZero() {
			continue
		}
		if record.Position.Quantity.IsNegative() {
			return fmt.Errorf("positions[%d]: negative position quantity", i)
		}
		marketI, err := ensureCarrier(subaccountID, marketID, fmt.Sprintf("positions[%d]", i))
		if err != nil {
			return err
		}
		state := stateFor(subaccountID, marketID)
		state.positionRaw.Set(record.Position.Quantity.BigInt())
		if !record.Position.IsLong {
			state.positionRaw.Neg(state.positionRaw)
		}
		if !membership.IsCrossMarginMarket(subaccountID, marketID) {
			continue
		}
		market, ok := marketI.(*DerivativeMarket)
		if !ok || record.Position.EntryPrice.IsNil() || record.Position.Margin.IsNil() ||
			(market.IsPerpetual && record.Position.CumulativeFundingEntry.IsNil()) {
			return fmt.Errorf("positions[%d]: live cross-margin position is incomplete", i)
		}
		if market.QuoteDecimals == 0 || market.QuoteDecimals > types.MaxDecimals {
			return fmt.Errorf("positions[%d]: quote decimals must be between 1 and %d", i, types.MaxDecimals)
		}
		if !market.CanRepresentNotionalInChainFormat(record.Position.Margin) ||
			!record.Position.EntryPrice.IsPositive() || !record.Position.EntryPrice.IsInValidRange() ||
			!canonicalRFQGenesisChunkEntryNotionalFits(record.Position, market) {
			return fmt.Errorf("positions[%d]: position is not representable by canonical RFQ chunks", i)
		}
		poolKey := genesisCrossPoolKey{subaccountID: subaccountID, quoteDenom: market.QuoteDenom}
		if decimals, found := quoteDecimalsByPool[poolKey]; found && decimals != market.QuoteDecimals {
			return fmt.Errorf("positions[%d]: inconsistent quote decimals for %s", i, market.QuoteDenom)
		}
		quoteDecimalsByPool[poolKey] = market.QuoteDecimals
	}

	addDerivativeOrder := func(field string, marketID common.Hash, orderHash []byte, orderSubaccountID common.Hash, cid string, isBuy, outerIsBuy, isVanilla bool, quantity sdkmath.LegacyDec) error {
		if isBuy != outerIsBuy {
			return fmt.Errorf("%s: outer side does not match order body", field)
		}
		if err := registerIdentity(field, orderSubaccountID, common.BytesToHash(orderHash), orderHash, cid); err != nil {
			return err
		}
		if !membership.HasCrossMarginExposure(orderSubaccountID) {
			return nil
		}
		if quantity.IsNil() || quantity.IsNegative() || !quantity.IsInValidRange() {
			return fmt.Errorf("%s: invalid derivative order quantity", field)
		}
		market, err := ensureCarrier(orderSubaccountID, marketID, field)
		if err != nil {
			return err
		}
		poolKey := genesisCrossPoolKey{subaccountID: orderSubaccountID, quoteDenom: market.GetQuoteDenom()}
		derivativeOrderCount[poolKey]++
		if derivativeOrderCount[poolKey] > uint64(MaxCrossMarginCancellableOrdersPerQuote) {
			return fmt.Errorf("%s: cross-margin C bound exceeded for %s", field, market.GetQuoteDenom())
		}
		if isVanilla {
			stateFor(orderSubaccountID, marketID).queuedRaw.Add(
				stateFor(orderSubaccountID, marketID).queuedRaw, quantity.BigInt(),
			)
		}
		return nil
	}

	for i, orderbook := range gs.DerivativeOrderbook {
		if !types.IsHexHash(orderbook.MarketId) {
			return fmt.Errorf("derivative_orderbook[%d]: invalid market_id %q", i, orderbook.MarketId)
		}
		marketID := common.HexToHash(orderbook.MarketId)
		for j, order := range orderbook.Orders {
			if order == nil {
				return fmt.Errorf("derivative_orderbook[%d].orders[%d]: missing order", i, j)
			}
			field := fmt.Sprintf("derivative_orderbook[%d].orders[%d]", i, j)
			if order.OrderInfo.Price.IsNil() || !order.OrderInfo.Price.IsInValidRange() ||
				order.OrderInfo.Quantity.IsNil() || !order.OrderInfo.Quantity.IsInValidRange() ||
				order.Fillable.IsNil() || !order.Fillable.IsInValidRange() ||
				order.Margin.IsNil() || !order.Margin.IsInValidRange() {
				return fmt.Errorf("%s: incomplete order", field)
			}
			if err := addDerivativeOrder(field, marketID, order.OrderHash, order.SubaccountID(), order.Cid(), order.IsBuy(), orderbook.IsBuySide, order.IsVanilla(), order.Fillable); err != nil {
				return err
			}
		}
	}

	for i, orderbook := range gs.ConditionalDerivativeOrderbooks {
		if orderbook == nil || !types.IsHexHash(orderbook.MarketId) {
			return fmt.Errorf("conditional_derivative_orderbooks[%d]: invalid orderbook or market_id", i)
		}
		marketID := common.HexToHash(orderbook.MarketId)
		if _, ok := derivativeMarkets[marketID].(*DerivativeMarket); !ok {
			return fmt.Errorf("conditional_derivative_orderbooks[%d]: market must be a non-binary derivative market", i)
		}
		for j, order := range orderbook.LimitBuyOrders {
			if order == nil || order.TriggerPrice == nil || order.TriggerPrice.IsNil() || !order.TriggerPrice.IsInValidRange() ||
				order.OrderInfo.Price.IsNil() || !order.OrderInfo.Price.IsInValidRange() ||
				order.OrderInfo.Quantity.IsNil() || !order.OrderInfo.Quantity.IsInValidRange() ||
				order.Fillable.IsNil() || !order.Fillable.IsInValidRange() ||
				order.Margin.IsNil() || !order.Margin.IsInValidRange() {
				return fmt.Errorf("conditional_derivative_orderbooks[%d].limit_buy_orders[%d]: incomplete order", i, j)
			}
			field := fmt.Sprintf("conditional_derivative_orderbooks[%d].limit_buy_orders[%d]", i, j)
			if err := addDerivativeOrder(field, marketID, order.OrderHash, order.SubaccountID(), order.Cid(), order.IsBuy(), true, order.IsVanilla(), order.Fillable); err != nil {
				return err
			}
		}
		for j, order := range orderbook.LimitSellOrders {
			if order == nil || order.TriggerPrice == nil || order.TriggerPrice.IsNil() || !order.TriggerPrice.IsInValidRange() ||
				order.OrderInfo.Price.IsNil() || !order.OrderInfo.Price.IsInValidRange() ||
				order.OrderInfo.Quantity.IsNil() || !order.OrderInfo.Quantity.IsInValidRange() ||
				order.Fillable.IsNil() || !order.Fillable.IsInValidRange() ||
				order.Margin.IsNil() || !order.Margin.IsInValidRange() {
				return fmt.Errorf("conditional_derivative_orderbooks[%d].limit_sell_orders[%d]: incomplete order", i, j)
			}
			field := fmt.Sprintf("conditional_derivative_orderbooks[%d].limit_sell_orders[%d]", i, j)
			if err := addDerivativeOrder(field, marketID, order.OrderHash, order.SubaccountID(), order.Cid(), order.IsBuy(), false, order.IsVanilla(), order.Fillable); err != nil {
				return err
			}
		}
		for j, order := range orderbook.MarketBuyOrders {
			if order == nil || order.TriggerPrice == nil || order.TriggerPrice.IsNil() || !order.TriggerPrice.IsInValidRange() ||
				order.OrderInfo.Price.IsNil() || !order.OrderInfo.Price.IsInValidRange() ||
				order.OrderInfo.Quantity.IsNil() || !order.OrderInfo.Quantity.IsInValidRange() ||
				order.Margin.IsNil() || !order.Margin.IsInValidRange() {
				return fmt.Errorf("conditional_derivative_orderbooks[%d].market_buy_orders[%d]: incomplete order", i, j)
			}
			field := fmt.Sprintf("conditional_derivative_orderbooks[%d].market_buy_orders[%d]", i, j)
			if err := addDerivativeOrder(field, marketID, order.OrderHash, order.SubaccountID(), order.Cid(), order.IsBuy(), true, order.IsVanilla(), order.OrderInfo.Quantity); err != nil {
				return err
			}
		}
		for j, order := range orderbook.MarketSellOrders {
			if order == nil || order.TriggerPrice == nil || order.TriggerPrice.IsNil() || !order.TriggerPrice.IsInValidRange() ||
				order.OrderInfo.Price.IsNil() || !order.OrderInfo.Price.IsInValidRange() ||
				order.OrderInfo.Quantity.IsNil() || !order.OrderInfo.Quantity.IsInValidRange() ||
				order.Margin.IsNil() || !order.Margin.IsInValidRange() {
				return fmt.Errorf("conditional_derivative_orderbooks[%d].market_sell_orders[%d]: incomplete order", i, j)
			}
			field := fmt.Sprintf("conditional_derivative_orderbooks[%d].market_sell_orders[%d]", i, j)
			if err := addDerivativeOrder(field, marketID, order.OrderHash, order.SubaccountID(), order.Cid(), order.IsBuy(), false, order.IsVanilla(), order.OrderInfo.Quantity); err != nil {
				return err
			}
		}
	}

	spotHolds := make(map[common.Hash][]genesisSpotHold)
	for i, orderbook := range gs.SpotOrderbook {
		marketID := common.HexToHash(orderbook.MarketId)
		market := spotMarkets[marketID]
		for j, order := range orderbook.Orders {
			if order == nil {
				return fmt.Errorf("spot_orderbook[%d].orders[%d]: missing order", i, j)
			}
			field := fmt.Sprintf("spot_orderbook[%d].orders[%d]", i, j)
			if order.IsBuy() != orderbook.IsBuySide {
				return fmt.Errorf("%s: outer side does not match order body", field)
			}
			subaccountID := order.SubaccountID()
			if err := registerIdentity(field, subaccountID, order.Hash(), order.OrderHash, order.Cid()); err != nil {
				return err
			}
			if !membership.HasCrossMarginExposure(subaccountID) {
				continue
			}
			hold, lockingDenom, err := CanonicalCrossMarginSpotLimitHold(order, market, false)
			if err != nil {
				return fmt.Errorf("%s: %w", field, err)
			}
			spotHolds[subaccountID] = append(spotHolds[subaccountID], genesisSpotHold{
				marketID: marketID, isBuy: order.IsBuy(), orderHash: order.Hash(),
				lockingDenom: lockingDenom, holdRaw: new(big.Int).Set(hold.BigInt()),
			})
		}
	}

	totals := gs.finalGenesisBalanceTotals()
	spotSurvivors := make(map[genesisCrossPoolKey]uint64)
	for subaccountID, holds := range spotHolds {
		slices.SortFunc(holds, func(a, b genesisSpotHold) int {
			if cmp := bytes.Compare(a.marketID.Bytes(), b.marketID.Bytes()); cmp != 0 {
				return cmp
			}
			if a.isBuy != b.isBuy {
				if !a.isBuy {
					return -1
				}
				return 1
			}
			return bytes.Compare(a.orderHash.Bytes(), b.orderHash.Bytes())
		})
		byDenom := make(map[string][]genesisSpotHold)
		for _, hold := range holds {
			byDenom[hold.lockingDenom] = append(byDenom[hold.lockingDenom], hold)
		}
		for denom, denomHolds := range byDenom {
			total := totals[genesisCrossPoolKey{subaccountID: subaccountID, quoteDenom: denom}]
			totalRaw := new(big.Int)
			if !total.IsNil() {
				totalRaw.Set(total.BigInt())
			}
			holdRaw := new(big.Int)
			for _, hold := range denomHolds {
				holdRaw.Add(holdRaw, hold.holdRaw)
			}
			firstSurvivor := 0
			for holdRaw.Cmp(totalRaw) > 0 && firstSurvivor < len(denomHolds) {
				holdRaw.Sub(holdRaw, denomHolds[firstSurvivor].holdRaw)
				firstSurvivor++
			}
			spotSurvivors[genesisCrossPoolKey{subaccountID: subaccountID, quoteDenom: denom}] = uint64(len(denomHolds) - firstSurvivor)
		}
	}

	for poolKey, derivativeCount := range derivativeOrderCount {
		if derivativeCount+spotSurvivors[poolKey] > uint64(MaxCrossMarginCancellableOrdersPerQuote) {
			return fmt.Errorf("cross-margin C bound exceeded for %s", poolKey.quoteDenom)
		}
	}
	for poolKey, spotCount := range spotSurvivors {
		if derivativeOrderCount[poolKey]+spotCount > uint64(MaxCrossMarginCancellableOrdersPerQuote) {
			return fmt.Errorf("cross-margin C bound exceeded for %s", poolKey.quoteDenom)
		}
	}
	for subaccountID, states := range marketStates {
		chunksByDenom := make(map[string]*big.Int)
		for marketID, state := range states {
			market := derivativeMarkets[marketID]
			chunks, err := CrossMarginCanonicalChunkCountRaw(state.positionRaw, state.queuedRaw)
			if err != nil {
				return err
			}
			if chunksByDenom[market.GetQuoteDenom()] == nil {
				chunksByDenom[market.GetQuoteDenom()] = new(big.Int)
			}
			chunksByDenom[market.GetQuoteDenom()].Add(chunksByDenom[market.GetQuoteDenom()], chunks)
			if chunksByDenom[market.GetQuoteDenom()].Cmp(big.NewInt(int64(MaxCrossMarginCanonicalChunksPerQuote))) > 0 {
				return fmt.Errorf("cross-margin B bound exceeded for subaccount %s quote %s", subaccountID.Hex(), market.GetQuoteDenom())
			}
		}
	}
	return nil
}

func (gs GenesisState) validateGenesisBalances(membership GenesisCrossMarginMembership) error {
	for i, balance := range gs.Balances {
		if _, ok := types.IsValidSubaccountID(balance.SubaccountId); !ok || balance.Denom == "" ||
			balance.Deposits == nil || balance.Deposits.TotalBalance.IsNil() || balance.Deposits.AvailableBalance.IsNil() ||
			!balance.Deposits.TotalBalance.IsInValidRange() || !balance.Deposits.AvailableBalance.IsInValidRange() {
			return fmt.Errorf("balances[%d]: invalid balance", i)
		}
		if membership.HasCrossMarginExposure(common.HexToHash(balance.SubaccountId)) &&
			balance.Deposits.TotalBalance.IsNegative() {
			return fmt.Errorf("balances[%d]: cross-margin total balance cannot be negative", i)
		}
	}
	return nil
}

func (gs GenesisState) finalGenesisBalanceTotals() map[genesisCrossPoolKey]sdkmath.LegacyDec {
	totals := make(map[genesisCrossPoolKey]sdkmath.LegacyDec, len(gs.Balances))
	for _, balance := range gs.Balances {
		totals[genesisCrossPoolKey{
			subaccountID: common.HexToHash(balance.SubaccountId),
			quoteDenom:   balance.Denom,
		}] = balance.Deposits.TotalBalance
	}
	return totals
}

func canonicalRFQGenesisChunkEntryNotionalFits(position *Position, market *DerivativeMarket) bool {
	if position == nil || market == nil || !position.Quantity.IsPositive() || !position.EntryPrice.IsPositive() {
		return false
	}
	chunkQuantity := sdkmath.LegacyMinDec(position.Quantity, types.MaxOrderQuantity)
	precision := sdkmath.LegacyOneDec().BigInt()
	raw, remainder := new(big.Int), new(big.Int)
	raw.QuoRem(new(big.Int).Mul(chunkQuantity.BigInt(), position.EntryPrice.BigInt()), precision, remainder)
	halfPrecision := new(big.Int).Quo(precision, big.NewInt(2))
	if cmp := remainder.Cmp(halfPrecision); cmp > 0 || cmp == 0 && raw.Bit(0) == 1 {
		raw.Add(raw, big.NewInt(1))
	}
	entryNotional := sdkmath.LegacyNewDecFromBigIntWithPrec(raw, sdkmath.LegacyPrecision)
	return entryNotional.IsInValidRange() && market.CanRepresentNotionalInChainFormat(entryNotional)
}

func (gs GenesisState) validateSpotOrderbookMarkets() error {
	spotMarkets, err := gs.validatedSpotMarkets()
	if err != nil {
		return err
	}
	for i, orderbook := range gs.SpotOrderbook {
		marketID := orderbook.MarketId
		if !types.IsHexHash(marketID) {
			return fmt.Errorf("spot_orderbook[%d]: invalid market_id %q", i, marketID)
		}
		market, ok := spotMarkets[marketID]
		if !ok {
			return fmt.Errorf("spot_orderbook[%d]: unknown market_id %s", i, marketID)
		}
		if market.IsActive() {
			if err := validateSpotOrderbookTicks(i, orderbook, market); err != nil {
				return err
			}
		}
	}

	return nil
}

func (gs GenesisState) validatedSpotMarkets() (map[string]*SpotMarket, error) {
	markets := make(map[string]*SpotMarket, len(gs.SpotMarkets))
	for i, market := range gs.SpotMarkets {
		if market == nil {
			continue
		}
		if market.IsActive() {
			if err := ValidateSpotMarketTickSizes(market.MinPriceTickSize, market.MinQuantityTickSize); err != nil {
				return nil, fmt.Errorf("spot_markets[%d]: %w", i, err)
			}
		}
		markets[market.MarketId] = market
	}

	return markets, nil
}

func validateSpotOrderbookTicks(i int, orderbook SpotOrderBook, market *SpotMarket) error {
	for j, order := range orderbook.Orders {
		if order == nil {
			return fmt.Errorf("spot_orderbook[%d].orders[%d]: missing order", i, j)
		}
		if order.OrderInfo.Price.IsNil() || types.BreachesMinimumTickSize(order.OrderInfo.Price, market.MinPriceTickSize) {
			return fmt.Errorf("spot_orderbook[%d].orders[%d]: price does not match market tick size", i, j)
		}
		if order.Fillable.IsNil() || types.BreachesMinimumTickSize(order.Fillable, market.MinQuantityTickSize) {
			return fmt.Errorf("spot_orderbook[%d].orders[%d]: fillable does not match market tick size", i, j)
		}
	}

	return nil
}

func (gs GenesisState) validateDerivativeMarketSettlementScheduled() error {
	expiryInfos, err := gs.validatedExpiryInfoMarketIDs()
	if err != nil {
		return err
	}

	derivativeMarkets := gs.derivativeMarketByID()
	binaryMarkets := gs.binaryMarketIDs()
	seen := make(map[string]struct{}, len(gs.DerivativeMarketSettlementScheduled))
	for i, marker := range gs.DerivativeMarketSettlementScheduled {
		if err := validateScheduledSettlementMarkerID(i, marker, seen); err != nil {
			return err
		}
		if err := validateScheduledSettlementMarkerMarket(i, marker, derivativeMarkets, binaryMarkets, expiryInfos); err != nil {
			return err
		}
	}

	return nil
}

func (gs GenesisState) derivativeMarketByID() map[string]*DerivativeMarket {
	markets := make(map[string]*DerivativeMarket, len(gs.DerivativeMarkets))
	for _, market := range gs.DerivativeMarkets {
		if market == nil {
			continue
		}
		markets[market.MarketId] = market
	}

	return markets
}

func (gs GenesisState) binaryMarketIDs() map[string]struct{} {
	markets := make(map[string]struct{}, len(gs.BinaryOptionsMarkets))
	for _, market := range gs.BinaryOptionsMarkets {
		if market == nil {
			continue
		}
		markets[market.MarketId] = struct{}{}
	}

	return markets
}

func (gs GenesisState) validatedExpiryInfoMarketIDs() (map[string]struct{}, error) {
	expiryInfos := make(map[string]struct{}, len(gs.ExpiryFuturesMarketInfoState))
	for i, info := range gs.ExpiryFuturesMarketInfoState {
		if err := validateExpiryInfoState(i, info); err != nil {
			return nil, err
		}
		expiryInfos[info.MarketId] = struct{}{}
	}

	return expiryInfos, nil
}

func validateExpiryInfoState(i int, info ExpiryFuturesMarketInfoState) error {
	if !types.IsHexHash(info.MarketId) {
		return fmt.Errorf("expiry_futures_market_info_state[%d]: invalid market_id %q", i, info.MarketId)
	}
	if info.MarketInfo == nil {
		return fmt.Errorf("expiry_futures_market_info_state[%d]: missing market_info", i)
	}
	if info.MarketInfo.MarketId != "" && info.MarketInfo.MarketId != info.MarketId {
		return fmt.Errorf("expiry_futures_market_info_state[%d]: market_id mismatch %q != %q",
			i, info.MarketInfo.MarketId, info.MarketId)
	}

	return nil
}

func validateScheduledSettlementMarkerID(
	i int,
	marker DerivativeMarketSettlementInfo,
	seen map[string]struct{},
) error {
	marketID := marker.MarketId
	if !types.IsHexHash(marketID) {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: invalid market_id %q", i, marketID)
	}
	if _, ok := seen[marketID]; ok {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: duplicate market_id %s", i, marketID)
	}
	seen[marketID] = struct{}{}

	return nil
}

func validateScheduledSettlementMarkerMarket(
	i int,
	marker DerivativeMarketSettlementInfo,
	derivativeMarkets map[string]*DerivativeMarket,
	binaryMarkets map[string]struct{},
	expiryInfos map[string]struct{},
) error {
	marketID := marker.MarketId
	if derivativeMarket, ok := derivativeMarkets[marketID]; ok {
		return validateDerivativeSettlementMarker(i, marker, derivativeMarket, expiryInfos)
	}
	if _, ok := binaryMarkets[marketID]; ok {
		return validateBinarySettlementMarker(i, marker)
	}

	return fmt.Errorf("derivative_market_settlement_scheduled[%d]: unknown market_id %s", i, marketID)
}

func validateDerivativeSettlementMarker(
	i int,
	marker DerivativeMarketSettlementInfo,
	market *DerivativeMarket,
	expiryInfos map[string]struct{},
) error {
	marketID := marker.MarketId
	if market.IsTimeExpiry() {
		if _, hasInfo := expiryInfos[marketID]; !hasInfo {
			return fmt.Errorf("derivative_market_settlement_scheduled[%d]: expiry market %s missing expiry info",
				i, marketID)
		}
	}
	if marker.SettlementPrice.IsNil() {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: derivative market %s missing settlement price",
			i, marketID)
	}
	if marker.IsForcedSettlement && marker.SettlementPrice.IsNegative() {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: forced derivative market %s has negative settlement price %s",
			i, marketID, marker.SettlementPrice.String())
	}
	if !marker.IsForcedSettlement && marker.SettlementPrice.IsNegative() {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: derivative market %s has negative settlement price %s",
			i, marketID, marker.SettlementPrice.String())
	}

	return nil
}

func validateBinarySettlementMarker(i int, marker DerivativeMarketSettlementInfo) error {
	marketID := marker.MarketId
	if marker.SettlementPrice.IsNil() {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: binary options market %s missing settlement price",
			i, marketID)
	}
	if marker.IsForcedSettlement && !isValidForcedBinarySettlementMarkerPrice(marker) {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: forced binary options market %s has invalid settlement price %s",
			i, marketID, marker.SettlementPrice.String())
	}
	if !marker.IsForcedSettlement && !isValidNonForcedBinarySettlementMarkerPrice(marker) {
		return fmt.Errorf("derivative_market_settlement_scheduled[%d]: binary options market %s has invalid settlement price %s",
			i, marketID, marker.SettlementPrice.String())
	}

	return nil
}

func isValidForcedBinarySettlementMarkerPrice(marker DerivativeMarketSettlementInfo) bool {
	return marker.SettlementPrice.IsPositive() &&
		marker.SettlementPrice.LTE(types.MaxBinaryOptionsOrderPrice)
}

func isValidNonForcedBinarySettlementMarkerPrice(marker DerivativeMarketSettlementInfo) bool {
	return marker.SettlementPrice.Equal(BinaryOptionsMarketRefundFlagPrice) ||
		(!marker.SettlementPrice.IsNegative() &&
			marker.SettlementPrice.LTE(types.MaxBinaryOptionsOrderPrice))
}

func validateRiskProfileRecord(record *SubaccountRiskProfileRecord) error {
	if record == nil {
		return errors.New("nil record")
	}

	if _, ok := types.IsValidSubaccountID(record.SubaccountId); !ok {
		return fmt.Errorf("invalid subaccount_id %q: must be a 32-byte hex hash", record.SubaccountId)
	}

	mode, err := validateRiskMode(&record.RiskProfile)
	if err != nil {
		return err
	}

	if mode == RiskMode_RISK_MODE_CROSS && types.IsDefaultSubaccountID(common.HexToHash(record.SubaccountId)) {
		return fmt.Errorf("default subaccount %s cannot use cross-margin mode", record.SubaccountId)
	}

	return nil
}

func validateRiskMode(p *SubaccountRiskProfile) (RiskMode, error) {
	mode := p.Mode
	if mode == RiskMode_RISK_MODE_UNSPECIFIED {
		mode = RiskMode_RISK_MODE_ISOLATED
	}
	if mode != RiskMode_RISK_MODE_ISOLATED && mode != RiskMode_RISK_MODE_CROSS {
		return 0, fmt.Errorf("unsupported risk mode %v", p.Mode)
	}

	policy := p.ReservationPolicy
	if policy == ReservationPolicy_RESERVATION_POLICY_UNSPECIFIED {
		policy = ReservationPolicy_RESERVATION_POLICY_FULL_HOLD
	}
	if policy != ReservationPolicy_RESERVATION_POLICY_FULL_HOLD {
		return 0, fmt.Errorf("unsupported reservation policy %v", p.ReservationPolicy)
	}

	if p.CreditLineId != "" {
		return 0, errors.New("credit lines are not supported")
	}

	return mode, nil
}

// validateReservedRFQInventories mirrors, on the genesis payload, the shape
// InitGenesis enforces on every configured router's two reserved provider
// inventories (nonces 2^32-2 long and 2^32-1 short): no CROSS profile or
// per-market CROSS override, no derivative orders, at most
// MaxCrossMarginRFQReservedInventoryMarkets position markets, and every
// position in the inventory's assigned direction. A payload that fails here
// would otherwise pass preflight and panic InitChain.
func (gs GenesisState) validateReservedRFQInventories(routers CrossMarginRFQRouterSet) error {
	reserved := make(map[common.Hash]bool) // subaccount -> assigned direction is long
	for _, router := range routers.Addresses() {
		for _, nonce := range []uint32{stdmath.MaxUint32 - 1, stdmath.MaxUint32} {
			subaccountID, err := types.SdkAddressWithNonceToSubaccountID(router, nonce)
			if err != nil || subaccountID == nil {
				return fmt.Errorf("router %s has no reserved provider inventory for nonce %d", router, nonce)
			}
			reserved[*subaccountID] = nonce == stdmath.MaxUint32-1
		}
	}
	if len(reserved) == 0 {
		return nil
	}
	for i, record := range gs.SubaccountRiskProfiles {
		if record == nil {
			continue
		}
		if _, isReserved := reserved[common.HexToHash(record.SubaccountId)]; isReserved && record.RiskProfile.Mode == RiskMode_RISK_MODE_CROSS {
			return fmt.Errorf("subaccount_risk_profiles[%d]: reserved RFQ provider inventory %s must not be cross-margin", i, record.SubaccountId)
		}
	}
	for i, record := range gs.SubaccountMarketRiskModes {
		if record == nil {
			continue
		}
		if _, isReserved := reserved[common.HexToHash(record.SubaccountId)]; isReserved && record.Mode == RiskMode_RISK_MODE_CROSS {
			return fmt.Errorf("subaccount_market_risk_modes[%d]: reserved RFQ provider inventory %s must not carry a cross-margin override", i, record.SubaccountId)
		}
	}
	marketRecords := make(map[common.Hash]int, len(gs.DerivativeMarkets)+len(gs.BinaryOptionsMarkets))
	for _, market := range gs.DerivativeMarkets {
		if market != nil {
			marketRecords[common.HexToHash(market.MarketId)]++
		}
	}
	for _, market := range gs.BinaryOptionsMarkets {
		if market != nil {
			marketRecords[common.HexToHash(market.MarketId)]++
		}
	}
	markets := make(map[common.Hash]map[common.Hash]struct{})
	for i, record := range gs.Positions {
		subaccountID := common.HexToHash(record.SubaccountId)
		isLong, isReserved := reserved[subaccountID]
		if !isReserved {
			continue
		}
		if record.Position == nil || record.Position.Quantity.IsNil() || !record.Position.Quantity.IsPositive() {
			return fmt.Errorf("positions[%d]: reserved RFQ provider inventory %s holds a non-positive position", i, record.SubaccountId)
		}
		if record.Position.EntryPrice.IsNil() || record.Position.Margin.IsNil() || record.Position.CumulativeFundingEntry.IsNil() {
			return fmt.Errorf("positions[%d]: reserved RFQ provider inventory %s holds an incomplete position body", i, record.SubaccountId)
		}
		if marketRecords[common.HexToHash(record.MarketId)] != 1 {
			return fmt.Errorf(
				"positions[%d]: reserved RFQ provider inventory %s position market %s must resolve to exactly one market record",
				i, record.SubaccountId, record.MarketId,
			)
		}
		if record.Position.IsLong != isLong {
			return fmt.Errorf("positions[%d]: reserved RFQ provider inventory %s holds a position against its assigned direction in market %s", i, record.SubaccountId, record.MarketId)
		}
		if markets[subaccountID] == nil {
			markets[subaccountID] = make(map[common.Hash]struct{})
		}
		markets[subaccountID][common.HexToHash(record.MarketId)] = struct{}{}
		if len(markets[subaccountID]) > MaxCrossMarginRFQReservedInventoryMarkets {
			return fmt.Errorf("positions[%d]: reserved RFQ provider inventory %s exceeds %d position markets", i, record.SubaccountId, MaxCrossMarginRFQReservedInventoryMarkets)
		}
	}
	for i, orderbook := range gs.DerivativeOrderbook {
		for j, order := range orderbook.Orders {
			if order == nil {
				continue
			}
			if _, isReserved := reserved[common.HexToHash(order.OrderInfo.SubaccountId)]; isReserved {
				return fmt.Errorf("derivative_orderbook[%d].orders[%d]: reserved RFQ provider inventory %s must not hold orders", i, j, order.OrderInfo.SubaccountId)
			}
		}
	}
	for i, orderbook := range gs.SpotOrderbook {
		for j, order := range orderbook.Orders {
			if order == nil {
				continue
			}
			if _, isReserved := reserved[common.HexToHash(order.OrderInfo.SubaccountId)]; isReserved {
				return fmt.Errorf("spot_orderbook[%d].orders[%d]: reserved RFQ provider inventory %s must not hold orders", i, j, order.OrderInfo.SubaccountId)
			}
		}
	}
	enabledQuoteDenoms := make(map[string]struct{}, len(gs.Params.CrossMarginParams.EnabledQuoteDenoms))
	for _, denom := range gs.Params.CrossMarginParams.EnabledQuoteDenoms {
		enabledQuoteDenoms[denom] = struct{}{}
	}
	for i, balance := range gs.Balances {
		if balance.Deposits == nil {
			continue
		}
		if _, isReserved := reserved[common.HexToHash(balance.SubaccountId)]; !isReserved {
			continue
		}
		if _, enabled := enabledQuoteDenoms[balance.Denom]; !enabled {
			continue
		}
		if balance.Deposits.AvailableBalance.IsNil() || balance.Deposits.TotalBalance.IsNil() ||
			balance.Deposits.AvailableBalance.LT(balance.Deposits.TotalBalance) {
			return fmt.Errorf("balances[%d]: reserved RFQ provider inventory %s holds locked %s funds", i, balance.SubaccountId, balance.Denom)
		}
	}
	for i, orderbook := range gs.ConditionalDerivativeOrderbooks {
		if orderbook == nil {
			continue
		}
		check := func(kind string, j int, subaccountID string) error {
			if _, isReserved := reserved[common.HexToHash(subaccountID)]; isReserved {
				return fmt.Errorf("conditional_derivative_orderbooks[%d].%s[%d]: reserved RFQ provider inventory %s must not hold orders", i, kind, j, subaccountID)
			}
			return nil
		}
		for j, order := range orderbook.LimitBuyOrders {
			if order != nil {
				if err := check("limit_buy_orders", j, order.OrderInfo.SubaccountId); err != nil {
					return err
				}
			}
		}
		for j, order := range orderbook.LimitSellOrders {
			if order != nil {
				if err := check("limit_sell_orders", j, order.OrderInfo.SubaccountId); err != nil {
					return err
				}
			}
		}
		for j, order := range orderbook.MarketBuyOrders {
			if order != nil {
				if err := check("market_buy_orders", j, order.OrderInfo.SubaccountId); err != nil {
					return err
				}
			}
		}
		for j, order := range orderbook.MarketSellOrders {
			if order != nil {
				if err := check("market_sell_orders", j, order.OrderInfo.SubaccountId); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

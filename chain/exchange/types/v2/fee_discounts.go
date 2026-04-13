package v2

import (
	"bytes"
	"sort"
	"sync"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/chain/exchange/types"
)

func (s *FeeDiscountSchedule) CalculateFeeDiscountTier(
	stakedAmount math.Int,
	tradingVolume math.LegacyDec,
) (
	feeDiscountRates *types.FeeDiscountRates,
	tierLevel uint64,
) {
	highestTierLevel := 0
	// O(N) but probably the most efficient way nonetheless since we just have 10 tiers and most will be on low tiers
	for idx, tier := range s.TierInfos {
		// both tier conditions must be satisfied to reach a tier level
		if stakedAmount.LT(tier.StakedAmount) || tradingVolume.LT(tier.Volume) {
			break
		}
		highestTierLevel = idx + 1
	}

	tierLevel = uint64(highestTierLevel)

	if tierLevel == 0 {
		feeDiscountRates = &types.FeeDiscountRates{
			MakerDiscountRate: math.LegacyZeroDec(),
			TakerDiscountRate: math.LegacyZeroDec(),
		}
	} else {
		feeDiscountRates = &types.FeeDiscountRates{
			MakerDiscountRate: s.TierInfos[highestTierLevel-1].MakerDiscountRate,
			TakerDiscountRate: s.TierInfos[highestTierLevel-1].TakerDiscountRate,
		}
	}

	return feeDiscountRates, tierLevel
}

func (s *FeeDiscountSchedule) TierOneRequirements() (
	minStakedAmount math.Int,
	minTradingFeePaid math.LegacyDec,
) {
	// s.TierInfos[0] is tier one, since tier 0 is implicit
	return s.TierInfos[0].StakedAmount, s.TierInfos[0].Volume
}

func (s *FeeDiscountSchedule) GetFeeDiscountRatesMap() types.FeeDiscountRatesMap {
	if s == nil {
		return make(types.FeeDiscountRatesMap)
	}

	feeDiscountRatesMap := make(types.FeeDiscountRatesMap, len(s.TierInfos))
	feeDiscountRatesMap[0] = &types.FeeDiscountRates{
		MakerDiscountRate: math.LegacyZeroDec(),
		TakerDiscountRate: math.LegacyZeroDec(),
	}

	for idx, tierInfo := range s.TierInfos {
		feeDiscountRatesMap[uint64(idx+1)] = &types.FeeDiscountRates{
			MakerDiscountRate: tierInfo.MakerDiscountRate,
			TakerDiscountRate: tierInfo.TakerDiscountRate,
		}
	}
	return feeDiscountRatesMap
}

func NewFeeDiscountTierTTL(tier uint64, ttlTimestamp int64) *FeeDiscountTierTTL {
	return &FeeDiscountTierTTL{
		Tier:         tier,
		TtlTimestamp: ttlTimestamp,
	}
}

type ValidatorCache map[string]stakingtypes.ValidatorI

func NewFeeDiscountConfig(isQualified bool, stakingInfo *FeeDiscountStakingInfo) *FeeDiscountConfig {
	if stakingInfo == nil {
		isQualified = false
	}
	return &FeeDiscountConfig{
		IsMarketQualified:      isQualified,
		FeeDiscountStakingInfo: stakingInfo,
	}
}

type FeeDiscountConfig struct {
	IsMarketQualified bool
	*FeeDiscountStakingInfo
}

func (c *FeeDiscountConfig) GetFeeDiscountRate(account sdk.AccAddress, isMaker bool) *math.LegacyDec { // nolint:revive // ok
	if c == nil {
		return nil
	}

	c.AccountFeeTiersMux.RLock()
	defer c.AccountFeeTiersMux.RUnlock()

	tier, ok := c.AccountFeeTiers[types.SdkAccAddressToAccount(account)]
	if !ok || tier == nil {
		// should never happen but just in case
		return nil
	}

	if isMaker {
		return &tier.MakerDiscountRate
	}

	return &tier.TakerDiscountRate
}

func (c *FeeDiscountConfig) IncrementAccountVolumeContribution(
	subaccountID common.Hash,
	marketID common.Hash,
	amount math.LegacyDec,
	isMaker bool,
) {
	// defensive programming: should never happen
	if amount.IsNegative() {
		return
	}

	shouldIncrementAccountVolumeContributions := c.IsMarketQualified
	account := types.SubaccountIDToAccount(subaccountID)

	c.AccountVolumesMux.Lock()
	defer c.AccountVolumesMux.Unlock()

	// skip account volume contributions if the market isn't qualified for fee discounts
	if shouldIncrementAccountVolumeContributions {
		if v, ok := c.AccountVolumeContributions[account]; !ok {
			c.AccountVolumeContributions[account] = amount
		} else {
			c.AccountVolumeContributions[account] = v.Add(amount)
		}
	}

	newVolume := NewVolumeWithSingleType(amount, isMaker)
	// the SubaccountMarketVolumeContributions is still fine to update though since volumes are recorded on a per-market level
	if innerMap, ok := c.SubaccountMarketVolumeContributions[subaccountID]; !ok {
		c.SubaccountMarketVolumeContributions[subaccountID] = map[common.Hash]VolumeRecord{
			marketID: newVolume,
		}
	} else {
		if v, ok := innerMap[marketID]; !ok {
			c.SubaccountMarketVolumeContributions[subaccountID][marketID] = newVolume
		} else {
			c.SubaccountMarketVolumeContributions[subaccountID][marketID] = v.Add(newVolume)
		}
	}
}

func NewFeeDiscountStakingInfo(
	schedule *FeeDiscountSchedule,
	currBucketStartTimestamp, oldestBucketStartTimestamp int64,
	maxTTLTimestamp int64,
	nextTTLTimestamp int64,
	isFirstFeeCycleFinished bool,
) *FeeDiscountStakingInfo {
	return &FeeDiscountStakingInfo{
		SubaccountMarketVolumeContributions: make(map[common.Hash]map[common.Hash]VolumeRecord),
		AccountVolumeContributions:          make(map[types.Account]math.LegacyDec),
		AccountFeeTiers:                     make(map[types.Account]*types.FeeDiscountRates),
		Validators:                          make(ValidatorCache),
		NewAccounts:                         make(map[types.Account]*FeeDiscountTierTTL),
		GrantCheckpoints:                    make(map[string]struct{}),
		InvalidGrants:                       make(map[string]string),
		AccountFeeTiersMux:                  new(sync.RWMutex),
		AccountVolumesMux:                   new(sync.RWMutex),
		ValidatorsMux:                       new(sync.RWMutex),
		NewAccountsMux:                      new(sync.RWMutex),
		GrantsMux:                           new(sync.RWMutex),
		Schedule:                            schedule,
		CurrBucketStartTimestamp:            currBucketStartTimestamp,
		OldestBucketStartTimestamp:          oldestBucketStartTimestamp,
		MaxTTLTimestamp:                     maxTTLTimestamp,
		NextTTLTimestamp:                    nextTTLTimestamp,
		FeeDiscountRatesCache:               schedule.GetFeeDiscountRatesMap(),
		IsFirstFeeCycleFinished:             isFirstFeeCycleFinished,
	}
}

type FeeDiscountStakingInfo struct {
	// subaccountID => marketID => volume
	SubaccountMarketVolumeContributions map[common.Hash]map[common.Hash]VolumeRecord
	AccountVolumeContributions          map[types.Account]math.LegacyDec
	AccountFeeTiers                     map[types.Account]*types.FeeDiscountRates
	Validators                          ValidatorCache
	NewAccounts                         map[types.Account]*FeeDiscountTierTTL
	GrantCheckpoints                    map[string]struct{}
	InvalidGrants                       map[string]string // grantee => granter

	AccountFeeTiersMux *sync.RWMutex
	AccountVolumesMux  *sync.RWMutex
	ValidatorsMux      *sync.RWMutex
	NewAccountsMux     *sync.RWMutex
	GrantsMux          *sync.RWMutex

	Schedule                   *FeeDiscountSchedule
	CurrBucketStartTimestamp   int64
	OldestBucketStartTimestamp int64
	MaxTTLTimestamp            int64
	NextTTLTimestamp           int64
	FeeDiscountRatesCache      types.FeeDiscountRatesMap
	IsFirstFeeCycleFinished    bool
}

type AccountTierTTL struct {
	Account sdk.AccAddress
	TierTTL *FeeDiscountTierTTL
}

type AccountContribution struct {
	Account sdk.AccAddress
	Amount  math.LegacyDec
}

type SubaccountVolumeContribution struct {
	SubaccountID common.Hash
	MarketID     common.Hash
	Volume       VolumeRecord
}

type MarketVolumeContribution struct {
	MarketID common.Hash
	Volume   VolumeRecord
}

func (info *FeeDiscountStakingInfo) GetSortedNewFeeDiscountAccountTiers() []*AccountTierTTL {
	accountTiers := make([]*AccountTierTTL, 0, len(info.AccountFeeTiers))
	info.NewAccountsMux.RLock()
	for k, v := range info.NewAccounts {
		accountTiers = append(accountTiers, &AccountTierTTL{
			Account: sdk.AccAddress([]byte(string(k[:]))),
			TierTTL: v,
		})
	}
	info.NewAccountsMux.RUnlock()
	sort.SliceStable(accountTiers, func(i, j int) bool {
		return bytes.Compare(accountTiers[i].Account.Bytes(), accountTiers[j].Account.Bytes()) < 0
	})
	return accountTiers
}

func (info *FeeDiscountStakingInfo) GetSortedAccountVolumeContributions() []*AccountContribution {
	accountContributions := make([]*AccountContribution, 0, len(info.AccountFeeTiers))
	info.AccountVolumesMux.RLock()
	for k, v := range info.AccountVolumeContributions {
		accountContributions = append(accountContributions, &AccountContribution{
			// use copy of value in closure, since the memory is not copied, it's reused.
			// So if your closure captures it, instead of copying via call args, you'll get same index in all goroutines
			Account: sdk.AccAddress([]byte(string(k[:]))),
			Amount:  v,
		})
	}
	info.AccountVolumesMux.RUnlock()
	sort.SliceStable(accountContributions, func(i, j int) bool {
		return bytes.Compare(accountContributions[i].Account.Bytes(), accountContributions[j].Account.Bytes()) < 0
	})
	return accountContributions
}

func (info *FeeDiscountStakingInfo) GetSortedSubaccountAndMarketVolumes() (
	[]*SubaccountVolumeContribution,
	[]*MarketVolumeContribution,
) {
	subaccountVolumes := make([]*SubaccountVolumeContribution, 0, len(info.AccountFeeTiers))
	marketVolumeTracker := make(map[common.Hash]VolumeRecord)

	info.AccountVolumesMux.RLock()
	for subaccountID, innerMap := range info.SubaccountMarketVolumeContributions {
		for marketID, volume := range innerMap {
			subaccountVolumes = append(subaccountVolumes, &SubaccountVolumeContribution{
				// use copy of value in closure, since the memory is not copied, it's reused.
				// So if your closure captures it, instead of copying via call args, you'll get same index in all goroutines
				SubaccountID: common.BytesToHash(subaccountID.Bytes()),
				MarketID:     common.BytesToHash(marketID.Bytes()),
				Volume:       volume,
			})

			if prevVolume, ok := marketVolumeTracker[marketID]; ok {
				marketVolumeTracker[marketID] = prevVolume.Add(volume)
			} else {
				marketVolumeTracker[marketID] = volume
			}
		}
	}
	info.AccountVolumesMux.RUnlock()

	sort.SliceStable(subaccountVolumes, func(i, j int) bool {
		return bytes.Compare(
			append(subaccountVolumes[i].SubaccountID.Bytes(), subaccountVolumes[i].MarketID.Bytes()...),
			append(subaccountVolumes[j].SubaccountID.Bytes(), subaccountVolumes[j].MarketID.Bytes()...),
		) < 0
	})

	marketVolumes := make([]*MarketVolumeContribution, 0, len(marketVolumeTracker))
	for k, v := range marketVolumeTracker {
		marketVolumes = append(marketVolumes, &MarketVolumeContribution{
			MarketID: common.BytesToHash(k.Bytes()),
			Volume:   v,
		})
	}

	sort.SliceStable(marketVolumes, func(i, j int) bool {
		return bytes.Compare(marketVolumes[i].MarketID.Bytes(), marketVolumes[j].MarketID.Bytes()) < 0
	})

	return subaccountVolumes, marketVolumes
}

func (info *FeeDiscountStakingInfo) GetSortedGrantCheckpointGrantersAndInvalidGrants() (
	granters []string,
	invalidGrants []*EventInvalidGrant,
) {
	info.GrantsMux.RLock()
	defer info.GrantsMux.RUnlock()

	granters = make([]string, 0, len(info.GrantCheckpoints))
	invalidGrants = make([]*EventInvalidGrant, 0, len(info.InvalidGrants))

	for k := range info.GrantCheckpoints {
		granters = append(granters, k)
	}

	sort.SliceStable(granters, func(i, j int) bool {
		return granters[i] < granters[j]
	})

	for k, v := range info.InvalidGrants {
		invalidGrants = append(invalidGrants, &EventInvalidGrant{
			Grantee: k,
			Granter: v,
		})
	}

	sort.SliceStable(invalidGrants, func(i, j int) bool {
		return invalidGrants[i].Grantee < invalidGrants[j].Grantee
	})
	return granters, invalidGrants
}

func (info *FeeDiscountStakingInfo) SetAccountTierInfo(accAddress sdk.AccAddress, discountRates *types.FeeDiscountRates) {
	info.AccountFeeTiersMux.Lock()
	info.AccountFeeTiers[types.SdkAccAddressToAccount(accAddress)] = discountRates
	info.AccountFeeTiersMux.Unlock()
}

func (info *FeeDiscountStakingInfo) SetNewAccountTierTTL(accAddress sdk.AccAddress, tier uint64) {
	info.NewAccountsMux.Lock()
	info.NewAccounts[types.SdkAccAddressToAccount(accAddress)] = &FeeDiscountTierTTL{
		Tier:         tier,
		TtlTimestamp: info.NextTTLTimestamp,
	}
	info.NewAccountsMux.Unlock()
}

func (info *FeeDiscountStakingInfo) GetIsPastTradingFeesCheckRequired() bool {
	return info.IsFirstFeeCycleFinished
}

func (info *FeeDiscountStakingInfo) AddCheckpoint(granter string) {
	info.GrantsMux.Lock()
	info.GrantCheckpoints[granter] = struct{}{}
	info.GrantsMux.Unlock()
}

func (info *FeeDiscountStakingInfo) AddInvalidGrant(grantee, granter string) {
	info.GrantsMux.Lock()
	info.InvalidGrants[grantee] = granter
	info.GrantsMux.Unlock()
}

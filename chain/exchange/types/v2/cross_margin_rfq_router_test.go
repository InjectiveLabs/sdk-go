package v2_test

import (
	"bytes"
	"math/rand"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	exchangev2 "github.com/InjectiveLabs/sdk-go/chain/exchange/types/v2"
)

type legacyScalarCrossMarginRFQRouter struct {
	Address string `protobuf:"bytes,15,opt,name=liquidation_rfq_contract_address,json=liquidationRfqContractAddress,proto3"`
}

func (m *legacyScalarCrossMarginRFQRouter) Reset()         { *m = legacyScalarCrossMarginRFQRouter{} }
func (m *legacyScalarCrossMarginRFQRouter) String() string { return proto.CompactTextString(m) }
func (*legacyScalarCrossMarginRFQRouter) ProtoMessage()    {}

func rfqRouterAddress(fill byte) sdk.AccAddress {
	return sdk.AccAddress(bytes.Repeat([]byte{fill}, 20))
}

func TestParseCrossMarginRFQRouterSetCanonicalizesByAddressBytes(t *testing.T) {
	t.Parallel()

	a := rfqRouterAddress(1)
	b := rfqRouterAddress(2)
	c := rfqRouterAddress(3)

	routers, err := exchangev2.ParseCrossMarginRFQRouterSet([]string{c.String(), a.String(), b.String()})
	require.NoError(t, err)
	require.Equal(t, []string{a.String(), b.String(), c.String()}, routers.CanonicalStrings())
	require.True(t, routers.Contains(b))
	require.False(t, routers.Contains(rfqRouterAddress(4)))

	addresses := routers.Addresses()
	addresses[0][0] = 9
	require.Equal(t, a.String(), routers.CanonicalStrings()[0], "returned addresses must not alias the parsed set")
}

func TestParseCrossMarginRFQRouterSetRejectsByteAliasesAndBounds(t *testing.T) {
	t.Parallel()

	a := rfqRouterAddress(1).String()
	_, err := exchangev2.ParseCrossMarginRFQRouterSet([]string{a, strings.ToUpper(a)})
	require.ErrorContains(t, err, "duplicate address")

	_, err = exchangev2.ParseCrossMarginRFQRouterSet([]string{"not-an-address"})
	require.ErrorContains(t, err, "is invalid")

	atLimit := make([]string, exchangev2.MaxCrossMarginRFQContractAddresses)
	for i := range atLimit {
		atLimit[i] = rfqRouterAddress(byte(i + 1)).String()
	}
	_, err = exchangev2.ParseCrossMarginRFQRouterSet(atLimit)
	require.NoError(t, err)

	overLimit := append(append([]string(nil), atLimit...), rfqRouterAddress(9).String())
	_, err = exchangev2.ParseCrossMarginRFQRouterSet(overLimit)
	require.ErrorContains(t, err, "at most 8")
}

func TestCrossMarginRFQRouterSetMaintainsAlternativeRouterReachability(t *testing.T) {
	t.Parallel()

	owner := rfqRouterAddress(1)
	empty, err := exchangev2.ParseCrossMarginRFQRouterSet(nil)
	require.NoError(t, err)
	require.Zero(t, countAlternativeRFQRouters(empty, owner))

	singleton, err := exchangev2.ParseCrossMarginRFQRouterSet([]string{owner.String()})
	require.NoError(t, err)
	require.Zero(t, countAlternativeRFQRouters(singleton, owner))
	require.Equal(t, 1, countAlternativeRFQRouters(singleton, rfqRouterAddress(2)))

	rng := rand.New(rand.NewSource(0x524651)) //nolint:gosec // deterministic property-test input
	for sample := 0; sample < 64; sample++ {
		routerCount := 2 + sample%(exchangev2.MaxCrossMarginRFQContractAddresses-1)
		values := make([]string, routerCount)
		for i := range values {
			raw := make([]byte, 20)
			_, err := rng.Read(raw)
			require.NoError(t, err)
			raw[0] = byte(i + 1) // guarantee uniqueness within this sample
			values[i] = sdk.AccAddress(raw).String()
		}

		routers, err := exchangev2.ParseCrossMarginRFQRouterSet(values)
		require.NoError(t, err)
		require.Equal(t, routerCount, routers.Len())
		for _, member := range routers.Addresses() {
			require.Equal(t, routerCount-1, countAlternativeRFQRouters(routers, member))
			require.Positive(t, countAlternativeRFQRouters(routers, member))
		}
		require.Equal(t, routerCount, countAlternativeRFQRouters(routers, rfqRouterAddress(0xff)))
	}
}

func countAlternativeRFQRouters(routers exchangev2.CrossMarginRFQRouterSet, owner sdk.AccAddress) int {
	count := 0
	for _, router := range routers.Addresses() {
		if !router.Equals(owner) {
			count++
		}
	}
	return count
}

func TestCrossMarginRFQRouterTag15WireCompatibility(t *testing.T) {
	t.Parallel()

	a := rfqRouterAddress(1).String()
	b := rfqRouterAddress(2).String()

	legacyBytes, err := proto.Marshal(&legacyScalarCrossMarginRFQRouter{Address: a})
	require.NoError(t, err)
	current := exchangev2.DefaultCrossMarginParams()
	require.NoError(t, proto.Unmarshal(legacyBytes, &current))
	require.Equal(t, []string{a}, current.LiquidationRfqContractAddress)

	current = exchangev2.DefaultCrossMarginParams()
	current.LiquidationRfqContractAddress = []string{a, b}
	currentBytes, err := proto.Marshal(&current)
	require.NoError(t, err)
	legacy := &legacyScalarCrossMarginRFQRouter{}
	require.NoError(t, proto.Unmarshal(currentBytes, legacy))
	require.Equal(t, b, legacy.Address, "a scalar decoder must retain the final repeated occurrence")
}

package types

import (
	"bytes"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

func IsEqualDenoms(denoms1, denoms2 []string) bool {
	denom1Map := make(map[string]struct{})
	denom2Map := make(map[string]struct{})

	for _, denom := range denoms1 {
		denom1Map[denom] = struct{}{}
	}
	for _, denom := range denoms2 {
		denom2Map[denom] = struct{}{}
	}

	return reflect.DeepEqual(denom1Map, denom2Map)
}

func IsPeggyToken(denom string) bool {
	return strings.HasPrefix(denom, "peggy0x")
}

func IsIBCDenom(denom string) bool {
	return strings.HasPrefix(denom, "ibc/")
}

type SpotLimitOrderDelta struct {
	Order        *SpotLimitOrder
	FillQuantity math.LegacyDec
}

type DerivativeLimitOrderDelta struct {
	Order          *DerivativeLimitOrder
	FillQuantity   math.LegacyDec
	CancelQuantity math.LegacyDec
}

type DerivativeMarketOrderDelta struct {
	Order        *DerivativeMarketOrder
	FillQuantity math.LegacyDec
}

func (d *DerivativeMarketOrderDelta) UnfilledQuantity() math.LegacyDec {
	return d.Order.OrderInfo.Quantity.Sub(d.FillQuantity)
}

func (d *DerivativeLimitOrderDelta) IsBuy() bool {
	return d.Order.IsBuy()
}

func (d *DerivativeLimitOrderDelta) SubaccountID() common.Hash {
	return d.Order.SubaccountID()
}

func (d *DerivativeLimitOrderDelta) Price() math.LegacyDec {
	return d.Order.Price()
}

func (d *DerivativeLimitOrderDelta) FillableQuantity() math.LegacyDec {
	return d.Order.Fillable.Sub(d.CancelQuantity)
}

func (d *DerivativeLimitOrderDelta) OrderHash() common.Hash {
	return d.Order.Hash()
}

func (d *DerivativeLimitOrderDelta) Cid() string {
	return d.Order.Cid()
}

var AuctionSubaccountID = common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
var ZeroSubaccountID = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")

// inj1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqe2hm49
var TempRewardsSenderAddress = sdk.AccAddress(common.HexToAddress(ZeroSubaccountID.Hex()).Bytes())

// inj1qqq3zyg3zyg3zyg3zyg3zyg3zyg3zyg3c9gg96
var AuctionFeesAddress = sdk.AccAddress(common.HexToAddress(AuctionSubaccountID.Hex()).Bytes())

var hexRegex = regexp.MustCompile("^(0x)?[0-9a-fA-F]+$")

func StringInSlice(a string, list *[]string) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}

func IsDefaultSubaccountID(subaccountID common.Hash) bool {
	// empty 12 bytes
	emptyBytes := make([]byte, common.HashLength-common.AddressLength)
	return bytes.Equal(subaccountID[common.AddressLength:], emptyBytes)
}

func IsNonceDerivedSubaccountID(subaccountID string) bool {
	return len(subaccountID) <= MaxSubaccountNonceLength
}

// CheckValidSubaccountIDOrNonce checks that either 1) the subaccountId is well-formatted subaccount nonce or
// 2) the normal subaccountId is a valid subaccount ID and the sender is the owner of the subaccount
func CheckValidSubaccountIDOrNonce(sender sdk.AccAddress, subaccountId string) error {
	if IsNonceDerivedSubaccountID(subaccountId) {
		if _, err := GetNonceDerivedSubaccountID(sender, subaccountId); err != nil {
			return errors.Wrap(ErrBadSubaccountID, subaccountId)
		}
		return nil
	}

	subaccountAddress, ok := IsValidSubaccountID(subaccountId)
	if !ok {
		return errors.Wrap(ErrBadSubaccountID, subaccountId)
	}

	if !bytes.Equal(subaccountAddress.Bytes(), sender.Bytes()) {
		return errors.Wrap(ErrBadSubaccountID, subaccountId)
	}
	return nil
}

func IsValidSubaccountID(subaccountID string) (*common.Address, bool) {
	if !IsHexHash(subaccountID) {
		return nil, false
	}
	subaccountIdBytes := common.FromHex(subaccountID)
	addressBytes := subaccountIdBytes[:common.AddressLength]
	address := common.BytesToAddress(addressBytes)
	return &address, true
}

func IsValidOrderHash(orderHash string) bool {
	return IsHexHash(orderHash)
}

func IsValidCid(cid string) bool {
	// Arbitrarily setting max length of cid to uuid length
	return len(cid) <= 36
}

// IsHexHash verifies whether a string can represent a valid hex-encoded hash or not.
func IsHexHash(s string) bool {
	if !isHexString(s) {
		return false
	}

	if strings.HasPrefix(s, "0x") {
		return len(s) == 2*common.HashLength+2
	}

	return len(s) == 2*common.HashLength
}

func isHexString(str string) bool {
	return hexRegex.MatchString(str)
}

func BreachesMinimumTickSize(value, minTickSize math.LegacyDec) bool {
	// obviously breached if the value less than the minTickSize
	if value.LT(minTickSize) {
		return true
	}

	// prevent mod by 0
	if minTickSize.IsZero() {
		return true
	}

	// is breaching when value % minTickSize != 0
	residue := new(big.Int).Mod(value.BigInt(), minTickSize.BigInt())
	return !bytes.Equal(residue.Bytes(), big.NewInt(0).Bytes())
}

func (s *Subaccount) GetSubaccountID() (*common.Hash, error) {
	trader, err := sdk.AccAddressFromBech32(s.Trader)
	if err != nil {
		return nil, err
	}
	return SdkAddressWithNonceToSubaccountID(trader, s.SubaccountNonce)
}

type Account [20]byte

func SdkAccAddressToAccount(account sdk.AccAddress) Account {
	var accAddr Account
	copy(accAddr[:], account.Bytes())
	return accAddr
}

func SubaccountIDToAccount(subaccountID common.Hash) Account {
	var accAddr Account
	copy(accAddr[:], subaccountID.Bytes()[:20])
	return accAddr
}

func GetSubaccountNonce(subaccountIdStr string) (uint32, error) {
	// maximum of 10^3 = 1000 numeric subaccounts allowed
	isValidLength := len(subaccountIdStr) <= MaxSubaccountNonceLength

	if !isValidLength {
		return 0, ErrBadSubaccountNonce
	}

	if subaccountIdStr == "" || subaccountIdStr == "0" {
		return 0, nil
	}

	num, err := strconv.Atoi(subaccountIdStr)
	if err != nil {
		return 0, err
	}

	if num < 0 || num > 999 {
		return 0, ErrBadSubaccountNonce
	}

	return uint32(num), nil
}

func MustGetSubaccountIDOrDeriveFromNonce(sender sdk.AccAddress, subaccountId string) common.Hash {
	subaccountID, err := GetSubaccountIDOrDeriveFromNonce(sender, subaccountId)
	if err != nil {
		panic(err)
	}

	return subaccountID
}

func GetNonceDerivedSubaccountID(sender sdk.AccAddress, subaccountId string) (common.Hash, error) {
	nonce, err := GetSubaccountNonce(subaccountId)
	if err != nil {
		return common.Hash{}, err
	}

	subaccountID, err := SdkAddressWithNonceToSubaccountID(sender, nonce)
	if err != nil {
		return common.Hash{}, err
	}

	return *subaccountID, nil
}

func GetSubaccountIDOrDeriveFromNonce(sender sdk.AccAddress, subaccountId string) (common.Hash, error) {
	if IsNonceDerivedSubaccountID(subaccountId) {
		return GetNonceDerivedSubaccountID(sender, subaccountId)
	}

	if !IsHexHash(subaccountId) {
		return common.Hash{}, errors.Wrap(ErrBadSubaccountID, subaccountId)
	}

	return common.HexToHash(subaccountId), nil
}

func SdkAddressWithNonceToSubaccountID(addr sdk.AccAddress, nonce uint32) (*common.Hash, error) {
	if len(addr.Bytes()) > common.AddressLength {
		return &AuctionSubaccountID, ErrBadSubaccountID
	}
	subaccountID := common.BytesToHash(append(addr.Bytes(), common.LeftPadBytes(big.NewInt(int64(nonce)).Bytes(), 12)...))

	return &subaccountID, nil
}

func MustSdkAddressWithNonceToSubaccountID(addr sdk.AccAddress, nonce uint32) common.Hash {
	if len(addr.Bytes()) > common.AddressLength {
		panic(ErrBadSubaccountID)
	}
	subaccountID := common.BytesToHash(append(addr.Bytes(), common.LeftPadBytes(big.NewInt(int64(nonce)).Bytes(), 12)...))

	return subaccountID
}

func SdkAddressToSubaccountID(addr sdk.AccAddress) common.Hash {
	return common.BytesToHash(common.RightPadBytes(addr.Bytes(), 32))
}

func SdkAddressToEthAddress(addr sdk.AccAddress) common.Address {
	return common.BytesToAddress(addr.Bytes())
}

func SubaccountIDToSdkAddress(subaccountID common.Hash) sdk.AccAddress {
	return sdk.AccAddress(subaccountID[:common.AddressLength])
}

func SubaccountIDToEthAddress(subaccountID common.Hash) common.Address {
	return common.BytesToAddress(subaccountID[:common.AddressLength])
}

func EthAddressToSubaccountID(addr common.Address) common.Hash {
	return common.BytesToHash(common.RightPadBytes(addr.Bytes(), 32))
}

func DecToDecBytes(dec math.LegacyDec) []byte {
	return dec.BigInt().Bytes()
}

func DecBytesToDec(bz []byte) math.LegacyDec {
	dec := math.LegacyNewDecFromBigIntWithPrec(new(big.Int).SetBytes(bz), math.LegacyPrecision)
	if dec.IsNil() {
		return math.LegacyZeroDec()
	}
	return dec
}

func IntToIntBytes(i math.Int) []byte {
	return i.BigInt().Bytes()
}

func IntBytesToInt(bz []byte) math.Int {
	return math.NewIntFromBigInt(new(big.Int).SetBytes(bz))
}

func HasDuplicatesOrder(slice []*OrderData) bool {
	seenHashes := make(map[string]struct{})
	seenCids := make(map[string]struct{})
	for _, item := range slice {
		hash, cid := item.GetOrderHash(), item.GetCid()
		_, hashExists := seenHashes[hash]
		_, cidExists := seenCids[cid]

		if (hash != "" && hashExists) || (cid != "" && cidExists) {
			return true
		}
		seenHashes[hash] = struct{}{}
		seenCids[cid] = struct{}{}
	}
	return false
}

package types

import (
	"bytes"
	"math/big"
	"reflect"

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

type Account [20]byte

func SdkAccAddressToAccount(account sdk.AccAddress) Account {
	var accAddr Account
	copy(accAddr[:], account.Bytes())
	return accAddr
}

type SpotLimitOrderDelta struct {
	Order        *SpotLimitOrder
	FillQuantity sdk.Dec
}

type DerivativeLimitOrderDelta struct {
	Order          *DerivativeLimitOrder
	FillQuantity   sdk.Dec
	CancelQuantity sdk.Dec
}

type DerivativeMarketOrderDelta struct {
	Order        *DerivativeMarketOrder
	FillQuantity sdk.Dec
}

func (d *DerivativeMarketOrderDelta) UnfilledQuantity() sdk.Dec {
	return d.Order.OrderInfo.Quantity.Sub(d.FillQuantity)
}

func (d *DerivativeLimitOrderDelta) IsBuy() bool {
	return d.Order.IsBuy()
}

func (d *DerivativeLimitOrderDelta) SubaccountID() common.Hash {
	return d.Order.SubaccountID()
}

func (d *DerivativeLimitOrderDelta) Price() sdk.Dec {
	return d.Order.Price()
}

func (d *DerivativeLimitOrderDelta) FillableQuantity() sdk.Dec {
	return d.Order.Fillable.Sub(d.CancelQuantity)
}

func (d *DerivativeLimitOrderDelta) OrderHash() common.Hash {
	return d.Order.Hash()
}

var AuctionSubaccountID = common.HexToHash("0x1111111111111111111111111111111111111111111111111111111111111111")
var ZeroSubaccountID = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
var TempRewardsSenderAddress = sdk.AccAddress(common.HexToAddress(AuctionSubaccountID.Hex()).Bytes())

func StringInSlice(a string, list *[]string) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}
	return false
}

func IsValidSubaccountID(subaccountID string) (*common.Address, bool) {
	if len(subaccountID) != 66 {
		return nil, false
	}
	subaccountIdBytes := common.FromHex(subaccountID)

	if len(subaccountIdBytes) != common.HashLength {
		return nil, false
	}
	addressBytes := subaccountIdBytes[:common.AddressLength]
	if !common.IsHexAddress(common.Bytes2Hex(addressBytes)) {
		return nil, false
	}
	address := common.BytesToAddress(addressBytes)
	return &address, true
}

func IsValidOrderHash(orderHash string) bool {
	if len(orderHash) != 66 {
		return false
	}

	orderHashBytes := common.FromHex(orderHash)
	return len(orderHashBytes) == common.HashLength
}

func BreachesMinimumTickSize(value sdk.Dec, minTickSize sdk.Dec) bool {
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

func SdkAddressWithNonceToSubaccountID(addr sdk.AccAddress, nonce uint32) (*common.Hash, error) {
	if len(addr.Bytes()) > common.AddressLength {
		return &AuctionSubaccountID, ErrBadSubaccountID
	}
	subaccountID := common.BytesToHash(append(addr.Bytes(), common.LeftPadBytes(big.NewInt(int64(nonce)).Bytes(), 12)...))

	return &subaccountID, nil
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

func EthAddressToSubaccountID(addr common.Address) common.Hash {
	return common.BytesToHash(common.RightPadBytes(addr.Bytes(), 32))
}

func DecToDecBytes(dec sdk.Dec) []byte {
	return dec.BigInt().Bytes()
}

func DecBytesToDec(bz []byte) sdk.Dec {
	dec := sdk.NewDecFromBigIntWithPrec(new(big.Int).SetBytes(bz), sdk.Precision)
	if dec.IsNil() {
		return sdk.ZeroDec()
	}
	return dec
}

func HasDuplicates(slice []string) bool {
	seen := make(map[string]struct{})
	for _, item := range slice {
		if _, ok := seen[item]; ok {
			return true
		}
		seen[item] = struct{}{}
	}
	return false
}

func HasDuplicatesHexHash(slice []string) bool {
	seen := make(map[common.Hash]struct{})
	for _, item := range slice {
		if _, ok := seen[common.HexToHash(item)]; ok {
			return true
		}
		seen[common.HexToHash(item)] = struct{}{}
	}
	return false
}

func HasDuplicatesCoin(slice []sdk.Coin) bool {
	seen := make(map[string]struct{})
	for _, item := range slice {
		if _, ok := seen[item.Denom]; ok {
			return true
		}
		seen[item.Denom] = struct{}{}
	}
	return false
}

func HasDuplicatesOrder(slice []*OrderData) bool {
	seen := make(map[string]struct{})
	for _, item := range slice {
		if _, ok := seen[item.OrderHash]; ok {
			return true
		}
		seen[item.OrderHash] = struct{}{}
	}
	return false
}

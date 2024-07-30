package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"

	peggytypes "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
)

func getMessageHash(publisher common.Address, assetPairID, timeStamp, price string) (hash common.Hash) {
	enc := encodePacked(publisher.Bytes(), []byte(assetPairID), encodeUint256(timeStamp), encodeUint256(price))

	return crypto.Keccak256Hash(enc)
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func encodeUint256(v string) []byte {
	bn := new(big.Int)
	bn.SetString(v, 10)
	return math.U256Bytes(bn)
}

func VerifyStorkMsgSignature(oraclePubKey common.Address, assetPairID, timeStamp, price string, signature []byte) bool {
	hash := getMessageHash(oraclePubKey, assetPairID, timeStamp, price)

	recoveredPubKey, err := peggytypes.EthAddressFromSignature(hash, signature)
	if err != nil {
		return false
	}

	return recoveredPubKey == oraclePubKey
}

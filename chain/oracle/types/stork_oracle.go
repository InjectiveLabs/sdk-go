package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

func getETHSignedMessageHash(message []byte) common.Hash {
	preamblePrefix := []byte(preamblePrefix)
	preambleMessage := append(preamblePrefix, message...) // nolint:gocritic
	preambleHash := crypto.Keccak256Hash(preambleMessage)

	return preambleHash
}

func getMessageHash(oracleName common.Address, assetPairID, timeStamp, price string) (hash common.Hash) {
	return crypto.Keccak256Hash(encodePacked(oracleName.Bytes(), []byte(assetPairID), encodeUint256(timeStamp), encodeUint256(price)))
}

func encodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func encodeUint256(v string) []byte {
	bn := new(big.Int)
	bn.SetString(v, 10)
	return math.U256Bytes(bn)
}

func getSigner(msgHash, signature []byte) common.Address {
	signedMsgHash := getETHSignedMessageHash(msgHash)
	pubKey, err := crypto.SigToPub(signedMsgHash.Bytes(), signature)
	if err != nil {
		return common.Address{}
	}
	return crypto.PubkeyToAddress(*pubKey)
}

func VerifyStorkMsgSignature(oraclePubKey common.Address, assetPairID, timeStamp, price string, signature []byte) bool {
	hash := getMessageHash(oraclePubKey, assetPairID, timeStamp, price)
	bz := hash.Bytes()
	address := getSigner(bz, signature)
	println("check address: ", address.String())
	return address == oraclePubKey
}

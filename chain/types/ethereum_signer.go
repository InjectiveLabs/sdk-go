package types

import (
	"math/big"

	"cosmossdk.io/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const signaturePrefix = "\x19Ethereum Signed Message:\n32"

// normalizeEthereumSignature returns a canonical 65-byte signature. Coinbase may relay 96-byte
// blobs where the recovery id is at index 95 instead of 64.
func normalizeEthereumSignature(signature []byte) ([]byte, error) {
	switch len(signature) {
	case 65:
		trimmedSig := make([]byte, 65)
		copy(trimmedSig, signature)
		return trimmedSig, nil
	case 96:
		trimmedSig := make([]byte, 65)
		copy(trimmedSig, signature[:65])
		trimmedSig[64] = signature[95]
		return trimmedSig, nil
	default:
		if len(signature) < 65 {
			return nil, errors.Wrap(ErrInvalidEthereumSignature, "signature too short")
		}
		return nil, errors.Wrapf(ErrInvalidEthereumSignature, "unexpected signature length: %d", len(signature))
	}
}

// recoverEthAddress performs the core EIP-191 address recovery from a normalized 65-byte sig.
func recoverEthAddress(hash common.Hash, sig []byte) (common.Address, error) {
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}
	if sig[64] != 0 && sig[64] != 1 {
		return common.Address{}, errors.Wrapf(ErrInvalidEthereumSignature, "invalid recovery id: %d", sig[64])
	}

	protectedHash := crypto.Keccak256Hash(append([]byte(signaturePrefix), hash[:]...))

	pubkey, err := crypto.SigToPub(protectedHash.Bytes(), sig)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "signature to public key")
	}

	return crypto.PubkeyToAddress(*pubkey), nil
}

// EthAddressFromSignature recovers the Ethereum address that signed hash using an EIP-191
// "\x19Ethereum Signed Message:\n32" prefixed message.
//
// This function accepts both low-S and high-S signatures. Coinbase oracle relayers produce
// high-S signatures, so callers that need malleability protection must use
// StrictEthAddressFromSignature instead.
func EthAddressFromSignature(hash common.Hash, signature []byte) (common.Address, error) {
	sig, err := normalizeEthereumSignature(signature)
	if err != nil {
		return common.Address{}, err
	}
	return recoverEthAddress(hash, sig)
}

// StrictEthAddressFromSignature is like EthAddressFromSignature but additionally enforces
// low-S (homestead) signature values, rejecting malleable high-S variants.
// Use this for any context where signature uniqueness must be guaranteed (e.g. peggy, stork).
func StrictEthAddressFromSignature(hash common.Hash, signature []byte) (common.Address, error) {
	sig, err := normalizeEthereumSignature(signature)
	if err != nil {
		return common.Address{}, err
	}

	v := sig[64]
	if v == 27 || v == 28 {
		v -= 27
	}
	if v != 0 && v != 1 {
		return common.Address{}, errors.Wrapf(ErrInvalidEthereumSignature, "invalid recovery id: %d", v)
	}

	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])
	if !crypto.ValidateSignatureValues(v, r, s, true) {
		return common.Address{}, errors.Wrap(ErrInvalidEthereumSignature, "signature values failed validation")
	}

	sig[64] = v
	return recoverEthAddress(hash, sig)
}

// ValidateEthereumSignature returns an error if signature does not recover to ethAddress for hash.
func ValidateEthereumSignature(hash common.Hash, signature []byte, ethAddress common.Address) error {
	addr, err := EthAddressFromSignature(hash, signature)
	if err != nil {
		return errors.Wrap(err, "")
	}

	if addr != ethAddress {
		return errors.Wrap(ErrInvalidEthereumSignature, "signature not matching")
	}

	return nil
}

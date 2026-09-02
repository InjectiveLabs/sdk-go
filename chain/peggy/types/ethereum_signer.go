package types

import (
	"crypto/ecdsa"

	"cosmossdk.io/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"
)

const (
	signaturePrefix = "\x19Ethereum Signed Message:\n32"
)

// NewEthereumSignature creates a new signuature over a given byte array
func NewEthereumSignature(hash common.Hash, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.Wrap(ErrEmpty, "private key")
	}
	protectedHash := crypto.Keccak256Hash(append([]uint8(signaturePrefix), hash[:]...))
	return crypto.Sign(protectedHash.Bytes(), privateKey)
}

// EthAddressFromSignature recovers the Ethereum address that signed hash using an EIP-191
// "\x19Ethereum Signed Message:\n32" prefixed message. Enforces low-S (homestead) validation
// to reject malleable high-S signature variants.
func EthAddressFromSignature(hash common.Hash, signature []byte) (common.Address, error) {
	return chaintypes.StrictEthAddressFromSignature(hash, signature)
}

// ValidateEthereumSignature takes a message, an associated signature and public key and
// returns an error if the signature isn't valid. Enforces low-S (homestead) validation.
func ValidateEthereumSignature(hash common.Hash, signature []byte, ethAddress common.Address) error {
	addr, err := chaintypes.StrictEthAddressFromSignature(hash, signature)
	if err != nil {
		return err
	}
	if addr != ethAddress {
		return errors.Wrap(ErrInvalid, "signature not matching")
	}
	return nil
}

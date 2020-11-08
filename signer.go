package sdk

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

// Signer defines the methods needed to act as a elliptic curve signer
type Signer interface {
	EthSign(message []byte, signerAddress common.Address) (*ECSignature, error)
	EcRecover(message []byte, sig []byte) (common.Address, error)
}

// SignatureType represents the type of 0x signature encountered
type SignatureType uint8

// SignatureType values
const (
	IllegalSignature SignatureType = iota
	InvalidSignature
	EIP712Signature
	EthSignSignature
	WalletSignature
	ValidatorSignature
	PreSignedSignature
	EIP1271WalletSignature
	NSignatureTypesSignature
)

// ECSignature contains the parameters of an elliptic curve signature
type ECSignature struct {
	V byte
	R common.Hash
	S common.Hash
}

// Bytes generates 0x EthSign Signature (append the signature type byte) data
func (ecSignature ECSignature) Bytes() []byte {
	signature := make([]byte, 66)

	signature[0] = ecSignature.V
	copy(signature[1:33], ecSignature.R[:])
	copy(signature[33:65], ecSignature.S[:])
	signature[65] = byte(EthSignSignature)

	return signature
}

// EthRPCSigner is a signer that uses a call to Ethereum JSON-RPC method `eth_call`
// to produce a signature
type EthRPCSigner struct {
	rpcClient *rpc.Client
}

// NewEthRPCSigner instantiates a new EthRPCSigner
func NewEthRPCSigner(rpcClient *rpc.Client) Signer {
	return &EthRPCSigner{
		rpcClient: rpcClient,
	}
}

// EthSign signs a message via the `eth_sign` Ethereum JSON-RPC call
func (e *EthRPCSigner) EthSign(message []byte, signerAddress common.Address) (*ECSignature, error) {
	var signatureHex string
	if err := e.rpcClient.Call(&signatureHex, "eth_sign", signerAddress.Hex(), common.Bytes2Hex(message)); err != nil {
		return nil, err
	}
	// `eth_sign` returns the signature in the [R || S || V] format where V is 0 or 1.
	signatureBytes := common.Hex2Bytes(signatureHex[2:])
	vParam := signatureBytes[64]
	if vParam == byte(0) {
		vParam = byte(27)
	} else if vParam == byte(1) {
		vParam = byte(28)
	}

	ecSignature := &ECSignature{
		V: vParam,
		R: common.BytesToHash(signatureBytes[0:32]),
		S: common.BytesToHash(signatureBytes[32:64]),
	}
	return ecSignature, nil
}

func (e *EthRPCSigner) EcRecover(message []byte, sig []byte) (common.Address, error) {
	return ecRecover(message, sig)
}

// LocalSigner is a signer that produces an `eth_sign`-compatible signature locally using
// a private key
type LocalSigner struct {
	privateKey *ecdsa.PrivateKey
}

// NewLocalSigner instantiates a new LocalSigner
func NewLocalSigner(privateKey *ecdsa.PrivateKey) Signer {
	return &LocalSigner{
		privateKey: privateKey,
	}
}

// GetSignerAddress returns the signerAddress corresponding to LocalSigner's private key
func (l *LocalSigner) GetSignerAddress() common.Address {
	return crypto.PubkeyToAddress(l.privateKey.PublicKey)
}

// EthSign mimicks the signing of `eth_sign` locally its supplied private key
func (l *LocalSigner) EthSign(message []byte, signerAddress common.Address) (*ECSignature, error) {
	// Add message prefix: "\x19Ethereum Signed Message:\n"${message length}
	messageWithPrefix, _ := textAndHash(message)

	ecSignature, err := l.sign(messageWithPrefix, signerAddress)
	if err != nil {
		err = errors.Wrap(err, "failed to sign message with prefix")
		return nil, err
	}

	return ecSignature, nil
}

func (l *LocalSigner) EcRecover(message []byte, sig []byte) (common.Address, error) {
	return ecRecover(message, sig)
}

func ecRecover(message []byte, sig []byte) (address common.Address, err error) {
	if len(sig) < 65 {
		err = errors.New("signature is too short")
		return
	}

	digestHash, _ := textAndHash(message)

	ecSignature := make([]byte, 65)
	copy(ecSignature[:32], sig[1:33])    // R
	copy(ecSignature[32:64], sig[33:65]) // S
	ecSignature[64] = sig[0] - 27        // V (0 or 1)

	var pubKey *ecdsa.PublicKey

	if pubKey, err = crypto.SigToPub(digestHash, ecSignature); err != nil {
		err = errors.Wrap(err, "failed get pub key from ecSignature")
		return
	}

	address = crypto.PubkeyToAddress(*pubKey)
	return address, nil
}

// Sign signs the message with the corresponding private key to the supplied signerAddress and returns
// the raw signature byte array
func (l *LocalSigner) simpleSign(message []byte, signerAddress common.Address) ([]byte, error) {
	expectedSignerAddress := l.GetSignerAddress()
	if signerAddress != expectedSignerAddress {
		err := errors.Errorf("Cannot sign with signerAddress %s since LocalSigner contains private key for %s", signerAddress, expectedSignerAddress)
		return nil, err
	}

	// The produced signature is in the [R || S || V] format where V is 0 or 1.
	signatureBytes, err := crypto.Sign(message, l.privateKey)
	if err != nil {
		err = errors.Wrap(err, "crypto sign error")
		return nil, err
	}

	return signatureBytes, nil
}

// Sign signs the message with the corresponding private key to the supplied signerAddress and returns
// the parsed V, R, S values
func (l *LocalSigner) sign(message []byte, signerAddress common.Address) (*ECSignature, error) {
	signatureBytes, err := l.simpleSign(message, signerAddress)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	vParam := signatureBytes[64]
	if vParam == byte(0) {
		vParam = byte(27)
	} else if vParam == byte(1) {
		vParam = byte(28)
	}

	ecSignature := &ECSignature{
		V: vParam,
		R: common.BytesToHash(signatureBytes[0:32]),
		S: common.BytesToHash(signatureBytes[32:64]),
	}

	return ecSignature, nil
}

// ================
// ================

// LocalKeystoreSigner is a signer that produces an `eth_sign`-compatible signature locally using
// a private key
type LocalKeystoreSigner struct {
	keystore *keystore.KeyStore
}

// NewLocalKeystoreSigner instantiates a new LocalKeystoreSigner
func NewLocalKeystoreSigner(keystore *keystore.KeyStore) Signer {
	return &LocalKeystoreSigner{
		keystore: keystore,
	}
}

func (l *LocalKeystoreSigner) EthSign(message []byte, signerAddress common.Address) (*ECSignature, error) {
	// Add message prefix: "\x19Ethereum Signed Message:\n"${message length}
	messageWithPrefix, _ := textAndHash(message)

	ecSignature, err := l.sign(messageWithPrefix, signerAddress)
	if err != nil {
		err = errors.Wrap(err, "failed to sign message with prefix")
		return nil, err
	}

	return ecSignature, nil
}

func (l *LocalKeystoreSigner) EcRecover(message []byte, sig []byte) (common.Address, error) {
	return ecRecover(message, sig)
}

// Sign signs the message with the corresponding private key to the supplied signerAddress and returns
// the raw signature byte array
func (l *LocalKeystoreSigner) simpleSign(message []byte, signerAddress common.Address) ([]byte, error) {
	var signerAccount accounts.Account
	signerAccount.Address = signerAddress

	// The produced signature is in the [R || S || V] format where V is 0 or 1.
	signatureBytes, err := l.keystore.SignHash(signerAccount, message)
	if err != nil {
		err = errors.Wrap(err, "keystore sign hash error")
		return nil, err
	}

	return signatureBytes, nil
}

// Sign signs the message with the corresponding private key to the supplied signerAddress and returns
// the parsed V, R, S values
func (l *LocalKeystoreSigner) sign(message []byte, signerAddress common.Address) (*ECSignature, error) {
	signatureBytes, err := l.simpleSign(message, signerAddress)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	vParam := signatureBytes[64]
	if vParam == byte(0) {
		vParam = byte(27)
	} else if vParam == byte(1) {
		vParam = byte(28)
	}

	ecSignature := &ECSignature{
		V: vParam,
		R: common.BytesToHash(signatureBytes[0:32]),
		S: common.BytesToHash(signatureBytes[32:64]),
	}

	return ecSignature, nil
}

// textAndHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calulcated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func textAndHash(data []byte) ([]byte, string) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
	hasher := sha3.NewLegacyKeccak256()
	// Note: Write will never return an error here. We added placeholders in order
	// to satisfy the linter.
	_, _ = hasher.Write([]byte(msg))
	return hasher.Sum(nil), msg
}

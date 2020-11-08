package sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	gethsigner "github.com/ethereum/go-ethereum/signer/core"
	"github.com/pkg/errors"
	"golang.org/x/crypto/sha3"
)

type CoordinatorApproval struct {
	TxOrigin             common.Address `json:"txOrigin"`
	TransactionHash      common.Hash    `json:"txHash"`
	TransactionSignature []byte         `json:"txSignature"`
	Domain               EIP712Domain   `json:"domain"`

	hash *common.Hash `json:"-"`
}

type SignedCoordinatorApproval struct {
	CoordinatorApproval
	Signature []byte `json:"signature"`
}

// ComputeApprovalHash computes a hash of 0x coordinator approval
func (a *CoordinatorApproval) ComputeApprovalHash() (common.Hash, error) {
	if a.hash != nil {
		return *a.hash, nil
	}

	var message = map[string]interface{}{
		"txOrigin":             a.TxOrigin.Hex(),
		"transactionHash":      a.TransactionHash.Bytes(),
		"transactionSignature": a.TransactionSignature,
	}

	var typedData = gethsigner.TypedData{
		Types:       eip712CoordinatorApprovalTypes,
		PrimaryType: "CoordinatorApproval",
		Domain:      a.CoordinatorDomain(),
		Message:     message,
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, err
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return common.Hash{}, err
	}

	w := sha3.NewLegacyKeccak256()
	w.Write([]byte("\x19\x01"))
	w.Write([]byte(domainSeparator))
	w.Write([]byte(typedDataHash))

	hash := common.BytesToHash(w.Sum(nil))
	a.hash = &hash

	return hash, nil
}

// SignCoordinatorApproval signs the 0x coordinator approval with the supplied Signer
func SignCoordinatorApproval(
	coordinatorFrom common.Address,
	ethSigner Signer,
	approval *CoordinatorApproval,
) (*SignedCoordinatorApproval, error) {
	if approval == nil {
		return nil, errors.New("cannot sign nil coordinator approval")
	}
	approvalHash, err := approval.ComputeApprovalHash()
	if err != nil {
		return nil, err
	}

	ecSignature, err := ethSigner.EthSign(approvalHash.Bytes(), coordinatorFrom)
	if err != nil {
		return nil, err
	}

	// Generate 0x Ethereum Signature (append the signature type byte)
	signature := make([]byte, 66)
	signature[0] = ecSignature.V
	copy(signature[1:33], ecSignature.R[:])
	copy(signature[33:65], ecSignature.S[:])
	signature[65] = byte(EthSignSignature)
	signedApproval := &SignedCoordinatorApproval{
		CoordinatorApproval: *approval,
		Signature:           signature,
	}

	return signedApproval, nil
}

func (a *CoordinatorApproval) CoordinatorDomain() gethsigner.TypedDataDomain {
	return makeCoordinatorDomain(a.Domain.ChainID, a.Domain.VerifyingContract)
}

func makeCoordinatorDomain(
	chainID *big.Int,
	verifyingContract common.Address,
) gethsigner.TypedDataDomain {
	return gethsigner.TypedDataDomain{
		Name:              "0x Protocol Coordinator",
		Version:           "3.0.0",
		ChainId:           math.NewHexOrDecimal256(chainID.Int64()),
		VerifyingContract: verifyingContract.Hex(),
	}
}

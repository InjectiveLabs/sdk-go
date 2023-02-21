package types

import (
	"fmt"

	sdksecp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	"github.com/InjectiveLabs/injective-core/injective-chain/crypto/ethsecp256k1"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ValidateFundsReclaimSignature(
	lockedAccountPubKey, signature []byte,
) error {
	// validate the signature of the locked account
	lockedPubKey := sdksecp256k1.PubKey{
		Key: lockedAccountPubKey,
	}

	correctPubKey := ethsecp256k1.PubKey{
		Key: lockedAccountPubKey,
	}

	lockedAddress := sdk.AccAddress(lockedPubKey.Address())
	recipient := sdk.AccAddress(correctPubKey.Address())

	signMessage := ConstructFundsReclaimMessage(recipient, lockedAddress)

	if !lockedPubKey.VerifySignature(signMessage, signature) {
		return ErrInvalidSignature
	}
	return nil
}

func ConstructFundsReclaimMessage(
	recipient, signer sdk.AccAddress,
) []byte {
	message := fmt.Sprintf("I authorize %s to reclaim my funds from locked account %s on Injective", recipient.String(), signer.String())
	return []byte(message)
}

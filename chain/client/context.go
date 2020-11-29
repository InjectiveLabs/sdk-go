package client

import (
	"os"

	"github.com/pkg/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"

	keyscodec "github.com/InjectiveLabs/sdk-go/chain/crypto/codec"
	"github.com/InjectiveLabs/sdk-go/chain/crypto/hd"
	orders "github.com/InjectiveLabs/sdk-go/chain/orders/types"
	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"
)

func NewClientContext(
	chainId string,
	privKey tmcrypto.PrivKey,
) (client.Context, error) {
	clientCtx := client.Context{}

	interfaceRegistry := types.NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	orders.RegisterInterfaces(interfaceRegistry)
	chaintypes.RegisterInterfaces(interfaceRegistry)

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	encodingConfig := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig: tx.NewTxConfig(marshaler, []signingtypes.SignMode{
			signingtypes.SignMode_SIGN_MODE_DIRECT,
		}),
	}

	var kb keyring.Keyring
	var info keyring.Info

	if privKey != nil {
		kb = keyring.NewInMemory(hd.EthSecp256k1Option())
		tmpPhrase := randPhrase(64)
		armor := crypto.EncryptArmorPrivKey(privKey, tmpPhrase, "eth_secp256k1")
		err := kb.ImportPrivKey(clientKeyName, armor, tmpPhrase)
		if err != nil {
			err = errors.Wrap(err, "failed to import privkey")
			return clientCtx, err
		}

		info, err = kb.Key(clientKeyName)
		if err != nil {
			err = errors.Wrap(err, "failed to get info about imported privkey")
			return clientCtx, err
		}
	}

	clientCtx = newContext(
		chainId,
		encodingConfig,
		kb,
		info,
	)

	return clientCtx, nil
}

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Marshaler
	TxConfig          client.TxConfig
}

var clientKeyName = "client"

func newContext(
	chainId string,
	encodingConfig EncodingConfig,
	kb keyring.Keyring,
	account keyring.Info,
) client.Context {
	clientCtx := client.Context{
		ChainID:           chainId,
		JSONMarshaler:     encodingConfig.Marshaler,
		InterfaceRegistry: encodingConfig.InterfaceRegistry,
		Keyring:           kb,
		Output:            os.Stderr,
		OutputFormat:      "json",
		From:              clientKeyName,
		BroadcastMode:     "block",
		FromName:          clientKeyName,
		UseLedger:         false,
		Simulate:          false,
		GenerateOnly:      false,
		Offline:           false,
		SkipConfirm:       true,
		TxConfig:          encodingConfig.TxConfig,
		AccountRetriever:  authtypes.AccountRetriever{},
	}

	if account != nil {
		clientCtx = clientCtx.WithFromAddress(account.GetAddress())
	}

	return clientCtx
}

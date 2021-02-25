package client

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/pkg/errors"

	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	keyscodec "github.com/InjectiveLabs/sdk-go/chain/crypto/codec"

	evm "github.com/InjectiveLabs/sdk-go/chain/evm/types"
	orders "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	peggy "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
	ctypes "github.com/InjectiveLabs/sdk-go/chain/types"
)

// NewTxConfig initializes new Cosmos TxConfig with certain signModes enabled.
func NewTxConfig(signModes []signingtypes.SignMode) client.TxConfig {
	interfaceRegistry := types.NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	orders.RegisterInterfaces(interfaceRegistry)
	evm.RegisterInterfaces(interfaceRegistry)
	peggy.RegisterInterfaces(interfaceRegistry)
	ctypes.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	bank.RegisterInterfaces(interfaceRegistry)
	staking.RegisterInterfaces(interfaceRegistry)
	gov.RegisterInterfaces(interfaceRegistry)
	evidence.RegisterInterfaces(interfaceRegistry)

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	return tx.NewTxConfig(marshaler, signModes)
}

// NewClientContext creates a new Cosmos Client context, where chainID
// corresponds to Cosmos chain ID, fromSpec is either name of the key, or bech32-address
// of the Cosmos account. Keyring is required to contain the specified key.
func NewClientContext(
	chainId, fromSpec string, kb keyring.Keyring,
) (client.Context, error) {
	clientCtx := client.Context{}

	interfaceRegistry := types.NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	orders.RegisterInterfaces(interfaceRegistry)
	evm.RegisterInterfaces(interfaceRegistry)
	peggy.RegisterInterfaces(interfaceRegistry)
	ctypes.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	bank.RegisterInterfaces(interfaceRegistry)
	staking.RegisterInterfaces(interfaceRegistry)
	gov.RegisterInterfaces(interfaceRegistry)
	evidence.RegisterInterfaces(interfaceRegistry)

	marshaler := codec.NewProtoCodec(interfaceRegistry)
	encodingConfig := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig: NewTxConfig([]signingtypes.SignMode{
			signingtypes.SignMode_SIGN_MODE_DIRECT,
		}),
	}

	var keyInfo keyring.Info
	var err error

	addr, err := cosmtypes.AccAddressFromBech32(fromSpec)
	if err == nil {
		keyInfo, err = kb.KeyByAddress(addr)
		if err != nil {
			err = errors.Wrapf(err, "failed to load key info by address %s", addr.String())
			return clientCtx, err
		}
	} else {
		// failed to parse Bech32, is it a name?
		keyInfo, err = kb.Key(fromSpec)
		if err != nil {
			err = errors.Wrapf(err, "no key in keyring for name: %s", fromSpec)
			return clientCtx, err
		}
	}

	clientCtx = newContext(
		chainId,
		encodingConfig,
		kb,
		keyInfo,
	)

	return clientCtx, nil
}

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Marshaler         codec.Marshaler
	TxConfig          client.TxConfig
}

func newContext(
	chainId string,
	encodingConfig EncodingConfig,
	kb keyring.Keyring,
	keyInfo keyring.Info,
) client.Context {
	clientCtx := client.Context{
		ChainID:           chainId,
		JSONMarshaler:     encodingConfig.Marshaler,
		InterfaceRegistry: encodingConfig.InterfaceRegistry,
		Keyring:           kb,
		Output:            os.Stderr,
		OutputFormat:      "json",
		From:              keyInfo.GetName(),
		BroadcastMode:     "block",
		FromName:          keyInfo.GetName(),
		UseLedger:         false,
		Simulate:          false,
		GenerateOnly:      false,
		Offline:           false,
		SkipConfirm:       true,
		TxConfig:          encodingConfig.TxConfig,
		AccountRetriever:  authtypes.AccountRetriever{},
	}

	if keyInfo != nil {
		clientCtx = clientCtx.WithFromAddress(keyInfo.GetAddress())
	}

	return clientCtx
}

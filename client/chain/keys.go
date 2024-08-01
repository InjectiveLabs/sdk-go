package chain

import (
	"bytes"
	"crypto/rand"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/codec"
	cosmcrypto "github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	crypto_cdc "github.com/InjectiveLabs/sdk-go/chain/crypto/codec"
	"github.com/InjectiveLabs/sdk-go/chain/crypto/ethsecp256k1"
	"github.com/InjectiveLabs/sdk-go/chain/crypto/hd"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

const defaultKeyringKeyName = "validator"

var emptyCosmosAddress = cosmtypes.AccAddress{}

func InitCosmosKeyring(
	cosmosKeyringDir string,
	cosmosKeyringAppName string,
	cosmosKeyringBackend string,
	cosmosKeyFrom string,
	cosmosKeyPassphrase string,
	cosmosPrivKey string,
	cosmosUseLedger bool,
) (cosmtypes.AccAddress, keyring.Keyring, error) {
	switch {
	case cosmosPrivKey != "":
		if cosmosUseLedger {
			err := errors.New("cannot combine ledger and privkey options")
			return emptyCosmosAddress, nil, err
		}

		pkBytes, err := common.HexToBytes(cosmosPrivKey)
		if err != nil {
			err = errors.Wrap(err, "failed to hex-decode cosmos account privkey")
			return emptyCosmosAddress, nil, err
		}

		// Specific to Injective chain with Ethermint keys
		// Should be secp256k1.PrivKey for generic Cosmos chain
		cosmosAccPk := &ethsecp256k1.PrivKey{
			Key: pkBytes,
		}

		addressFromPk := cosmtypes.AccAddress(cosmosAccPk.PubKey().Address().Bytes())

		var keyName string

		// check that if cosmos 'From' specified separately, it must match the provided privkey,
		if cosmosKeyFrom != "" {
			addressFrom, err := cosmtypes.AccAddressFromBech32(cosmosKeyFrom)
			if err == nil {
				if !bytes.Equal(addressFrom.Bytes(), addressFromPk.Bytes()) {
					err = errors.Errorf("expected account address %s but got %s from the private key", addressFrom.String(), addressFromPk.String())
					return emptyCosmosAddress, nil, err
				}
			} else {
				// use it as a name then
				keyName = cosmosKeyFrom
			}
		}

		if keyName == "" {
			keyName = defaultKeyringKeyName
		}

		// wrap a PK into a Keyring
		kb, err := KeyringForPrivKey(keyName, cosmosAccPk)
		return addressFromPk, kb, err

	case cosmosKeyFrom != "":
		var fromIsAddress bool
		addressFrom, err := cosmtypes.AccAddressFromBech32(cosmosKeyFrom)
		if err == nil {
			fromIsAddress = true
		}

		var passReader io.Reader = os.Stdin
		if cosmosKeyPassphrase != "" {
			passReader = newPassReader(cosmosKeyPassphrase)
		}

		var absoluteKeyringDir string
		if filepath.IsAbs(cosmosKeyringDir) {
			absoluteKeyringDir = cosmosKeyringDir
		} else {
			absoluteKeyringDir, _ = filepath.Abs(cosmosKeyringDir)
		}

		kb, err := keyring.New(
			cosmosKeyringAppName,
			cosmosKeyringBackend,
			absoluteKeyringDir,
			passReader,
			getCryptoCodec(),
			hd.EthSecp256k1Option(),
		)
		if err != nil {
			err = errors.Wrap(err, "failed to init keyring")
			return emptyCosmosAddress, nil, err
		}

		var keyInfo *keyring.Record
		if fromIsAddress {
			if keyInfo, err = kb.KeyByAddress(addressFrom); err != nil {
				err = errors.Wrapf(err, "couldn't find an entry for the key %s in keybase", addressFrom.String())
				return emptyCosmosAddress, nil, err
			}
		} else {
			if keyInfo, err = kb.Key(cosmosKeyFrom); err != nil {
				err = errors.Wrapf(err, "could not find an entry for the key '%s' in keybase", cosmosKeyFrom)
				return emptyCosmosAddress, nil, err
			}
		}

		switch keyType := keyInfo.GetType(); keyType {
		case keyring.TypeLocal:
			// kb has a key and it's totally usable
			addr, err := keyInfo.GetAddress()
			if err != nil {
				err = errors.Wrapf(err, "failed to get address for key '%s'", keyInfo.Name)
				return emptyCosmosAddress, nil, err
			}
			return addr, kb, nil
		case keyring.TypeLedger:
			// the kb stores references to ledger keys, so we must explicitly
			// check that. kb doesn't know how to scan HD keys - they must be added manually before
			if cosmosUseLedger {
				addr, err := keyInfo.GetAddress()
				if err != nil {
					err = errors.Wrapf(err, "failed to get address for key '%s'", keyInfo.Name)
					return emptyCosmosAddress, nil, err
				}
				return addr, kb, nil
			}
			err := errors.Errorf("'%s' key is a ledger reference, enable ledger option", keyInfo.Name)
			return emptyCosmosAddress, nil, err
		case keyring.TypeOffline:
			err := errors.Errorf("'%s' key is an offline key, not supported yet", keyInfo.Name)
			return emptyCosmosAddress, nil, err
		case keyring.TypeMulti:
			err := errors.Errorf("'%s' key is an multisig key, not supported yet", keyInfo.Name)
			return emptyCosmosAddress, nil, err
		default:
			err := errors.Errorf("'%s' key  has unsupported type: %s", keyInfo.Name, keyType)
			return emptyCosmosAddress, nil, err
		}

	default:
		err := errors.New("insufficient cosmos key details provided")
		return emptyCosmosAddress, nil, err
	}
}

func newPassReader(pass string) io.Reader {
	return &passReader{
		pass: pass,
		buf:  new(bytes.Buffer),
	}
}

type passReader struct {
	pass string
	buf  *bytes.Buffer
}

var _ io.Reader = &passReader{}

func (r *passReader) Read(p []byte) (n int, err error) {
	n, err = r.buf.Read(p)
	if err == io.EOF || n == 0 {
		r.buf.WriteString(r.pass + "\n")

		n, err = r.buf.Read(p)
	}

	return
}

func getCryptoCodec() *codec.ProtoCodec {
	registry := NewInterfaceRegistry()
	crypto_cdc.RegisterInterfaces(registry)
	return codec.NewProtoCodec(registry)
}

// KeyringForPrivKey creates a temporary in-mem keyring for a PrivKey.
// Allows to init Context when the key has been provided in plaintext and parsed.
func KeyringForPrivKey(name string, privKey cryptotypes.PrivKey) (keyring.Keyring, error) {
	kb := keyring.NewInMemory(getCryptoCodec(), hd.EthSecp256k1Option())
	tmpPhrase := randPhrase(64)
	armored := cosmcrypto.EncryptArmorPrivKey(privKey, tmpPhrase, privKey.Type())
	err := kb.ImportPrivKey(name, armored, tmpPhrase)
	if err != nil {
		err = errors.Wrap(err, "failed to import privkey")
		return nil, err
	}

	return kb, nil
}

func randPhrase(size int) string {
	buf := make([]byte, size)
	_, err := rand.Read(buf)
	orPanic(err)

	return string(buf)
}

func orPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

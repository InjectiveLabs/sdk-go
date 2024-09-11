package keyring

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/InjectiveLabs/sdk-go/chain/crypto/hd"
	"github.com/cosmos/cosmos-sdk/codec"
	cosmcrypto "github.com/cosmos/cosmos-sdk/crypto"
	cosmkeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

var (
	defaultKeyringKeyName = "default"
	emptyCosmosAddress    = sdk.AccAddress{}
)

// NewCosmosKeyring creates a new keyring from a variety of options. See ConfigOpt and related options.
func NewCosmosKeyring(cdc codec.Codec, opts ...ConfigOpt) (sdk.AccAddress, cosmkeyring.Keyring, error) {
	config := &cosmosKeyringConfig{}
	for optIdx, optFn := range opts {
		if err := optFn(config); err != nil {
			err = errors.Wrapf(ErrFailedToApplyConfigOption, "option #%d: %s", optIdx+1, err.Error())
			return emptyCosmosAddress, nil, err
		}
	}

	if len(config.Keys) == 0 {
		return emptyCosmosAddress, nil, ErrInsufficientKeyDetails
	}

	var kb cosmkeyring.Keyring
	var realKB cosmkeyring.Keyring
	var usingRealKeyring bool
	var firstKey *sdk.AccAddress

	for keyIdx, keyConfig := range config.Keys {
		switch {
		case keyConfig.Mnemonic != "":
			if usingRealKeyring {
				return emptyCosmosAddress, nil, ErrMultipleKeysWithDifferentSecurity
			} else if kb == nil {
				kb = cosmkeyring.NewInMemory(cdc, hd.EthSecp256k1Option())
			}

			if config.UseLedger {
				err := errors.Wrap(ErrIncompatibleOptionsProvided, "cannot combine ledger and mnemonic options")
				return emptyCosmosAddress, nil, err
			}

			addr, err := fromMnemonic(kb, keyConfig)
			if err != nil {
				return addr, kb, err
			}

			if keyIdx == 0 {
				firstKey = &addr
			}

		case keyConfig.PrivKeyHex != "":
			if usingRealKeyring {
				return emptyCosmosAddress, nil, ErrMultipleKeysWithDifferentSecurity
			} else if kb == nil {
				kb = cosmkeyring.NewInMemory(cdc, hd.EthSecp256k1Option())
			}

			if config.UseLedger {
				err := errors.Wrap(ErrIncompatibleOptionsProvided, "cannot combine ledger and privkey options")
				return emptyCosmosAddress, nil, err
			}

			addr, err := fromPrivkeyHex(kb, keyConfig)
			if err != nil {
				return addr, kb, err
			}

			if keyIdx == 0 {
				firstKey = &addr
			}

		case keyConfig.KeyFrom != "":
			if kb != nil {
				return emptyCosmosAddress, nil, ErrMultipleKeysWithDifferentSecurity
			} else {
				usingRealKeyring = true
			}

			var fromIsAddress bool

			addressFrom, err := sdk.AccAddressFromBech32(keyConfig.KeyFrom)
			if err == nil {
				fromIsAddress = true
			}

			addr, kb, err := fromCosmosKeyring(cdc, config, keyConfig, addressFrom, fromIsAddress)
			if err != nil {
				return addr, kb, err
			}

			realKB = kb
			if keyIdx == 0 {
				firstKey = &addr
			}

		default:
			err := errors.Wrapf(ErrInsufficientKeyDetails, "key %d details", keyIdx+1)
			return emptyCosmosAddress, nil, err
		}
	}

	if realKB != nil {
		if config.DefaultKey != "" {
			defaultKeyAddr, err := findKeyInKeyring(realKB, config, config.DefaultKey)
			if err != nil {
				return emptyCosmosAddress, nil, err
			}

			return defaultKeyAddr, realKB, nil
		}

		return *firstKey, realKB, nil
	}

	if config.DefaultKey != "" {
		defaultKeyAddr, err := findKeyInKeyring(kb, config, config.DefaultKey)
		if err != nil {
			return emptyCosmosAddress, nil, err
		}

		return defaultKeyAddr, kb, nil
	}

	return *firstKey, kb, nil
}

func fromPrivkeyHex(
	kb cosmkeyring.Keyring,
	keyConfig *cosmosKeyConfig,
) (sdk.AccAddress, error) {
	pkBytes, err := hexToBytes(keyConfig.PrivKeyHex)
	if err != nil {
		err = errors.Wrapf(ErrHexFormatError, "failed to decode cosmos account privkey: %s", err.Error())
		return emptyCosmosAddress, err
	}

	cosmosAccPk := hd.EthSecp256k1.Generate()(pkBytes)
	addressFromPk := sdk.AccAddress(cosmosAccPk.PubKey().Address().Bytes())

	keyName := keyConfig.Name

	// check that if cosmos 'From' specified separately, it must match the provided privkey
	if keyConfig.KeyFrom != "" {
		addressFrom, err := sdk.AccAddressFromBech32(keyConfig.KeyFrom)

		switch {
		case err == nil:
			if !bytes.Equal(addressFrom.Bytes(), addressFromPk.Bytes()) {
				err = errors.Wrapf(
					ErrUnexpectedAddress,
					"expected account address %s but got %s from the private key",
					addressFrom.String(), addressFromPk.String(),
				)

				return emptyCosmosAddress, err
			}

		case keyName == "":
			// use it as a name then
			keyName = keyConfig.KeyFrom

		case keyName != keyConfig.KeyFrom:
			err := errors.Errorf(
				"key 'from' opt is a name, but doesn't match given key name: %s != %s",
				keyConfig.KeyFrom, keyName,
			)

			return emptyCosmosAddress, err
		}
	}

	if keyName == "" {
		keyName = defaultKeyringKeyName
	}

	// add a PK into a Keyring
	err = addFromPrivKey(kb, keyName, cosmosAccPk)
	if err != nil {
		err = errors.WithStack(err)
	}

	return addressFromPk, err
}

func fromMnemonic(
	kb cosmkeyring.Keyring,
	keyConfig *cosmosKeyConfig,
) (sdk.AccAddress, error) {
	cfg := sdk.GetConfig()

	pkBytes, err := hd.EthSecp256k1.Derive()(
		keyConfig.Mnemonic,
		cosmkeyring.DefaultBIP39Passphrase,
		cfg.GetFullBIP44Path(),
	)
	if err != nil {
		err = errors.Wrapf(ErrDeriveFailed, "failed to derive secp256k1 private key: %s", err.Error())
		return emptyCosmosAddress, err
	}

	cosmosAccPk := hd.EthSecp256k1.Generate()(pkBytes)
	addressFromPk := sdk.AccAddress(cosmosAccPk.PubKey().Address().Bytes())

	keyName := keyConfig.Name

	// check that if cosmos 'From' specified separately, it must match the derived privkey
	if keyConfig.KeyFrom != "" {
		addressFrom, err := sdk.AccAddressFromBech32(keyConfig.KeyFrom)
		switch {
		case err == nil:
			if !bytes.Equal(addressFrom.Bytes(), addressFromPk.Bytes()) {
				err = errors.Wrapf(
					ErrUnexpectedAddress,
					"expected account address %s but got %s from the mnemonic at /0",
					addressFrom.String(), addressFromPk.String(),
				)

				return emptyCosmosAddress, err
			}
		case keyName == "":
			// use it as a name then
			keyName = keyConfig.KeyFrom
		case keyName != keyConfig.KeyFrom:
			err := errors.Errorf(
				"key 'from' opt is a name, but doesn't match given key name: %s != %s",
				keyConfig.KeyFrom, keyName,
			)
			return emptyCosmosAddress, err
		}
	}

	// check that if 'PrivKeyHex' specified separately, it must match the derived privkey too
	if keyConfig.PrivKeyHex != "" {
		if err := checkPrivkeyHexMatchesMnemonic(keyConfig.PrivKeyHex, pkBytes); err != nil {
			return emptyCosmosAddress, err
		}
	}

	if keyName == "" {
		keyName = defaultKeyringKeyName
	}

	// add a PK into a Keyring
	err = addFromPrivKey(kb, keyName, cosmosAccPk)
	if err != nil {
		err = errors.WithStack(err)
	}

	return addressFromPk, err
}

func checkPrivkeyHexMatchesMnemonic(pkHex string, mnemonicDerivedPkBytes []byte) error {
	pkBytesFromHex, err := hexToBytes(pkHex)
	if err != nil {
		err = errors.Wrapf(ErrHexFormatError, "failed to decode cosmos account privkey: %s", err.Error())
		return err
	}

	if !bytes.Equal(mnemonicDerivedPkBytes, pkBytesFromHex) {
		err := errors.Wrap(
			ErrPrivkeyConflict,
			"both mnemonic and privkey hex options provided, but privkey doesn't match mnemonic",
		)
		return err
	}

	return nil
}

func fromCosmosKeyring(
	cdc codec.Codec,
	config *cosmosKeyringConfig,
	keyConfig *cosmosKeyConfig,
	fromAddress sdk.AccAddress,
	fromIsAddress bool,
) (sdk.AccAddress, cosmkeyring.Keyring, error) {
	var passReader io.Reader = os.Stdin
	if keyConfig.KeyPassphrase != "" {
		passReader = newPassReader(keyConfig.KeyPassphrase)
	}

	var err error
	absoluteKeyringDir := config.KeyringDir
	if !filepath.IsAbs(config.KeyringDir) {
		absoluteKeyringDir, err = filepath.Abs(config.KeyringDir)
		if err != nil {
			err = errors.Wrapf(ErrFilepathIncorrect, "failed to get abs path for keyring dir: %s", err.Error())
			return emptyCosmosAddress, nil, err
		}
	}

	kb, err := cosmkeyring.New(
		config.KeyringAppName,
		string(config.KeyringBackend),
		absoluteKeyringDir,
		passReader,
		cdc,
		hd.EthSecp256k1Option(),
	)
	if err != nil {
		err = errors.Wrapf(ErrCosmosKeyringCreationFailed, "failed to init cosmos keyring: %s", err.Error())
		return emptyCosmosAddress, nil, err
	}

	var keyRecord *cosmkeyring.Record
	if fromIsAddress {
		keyRecord, err = kb.KeyByAddress(fromAddress)
	} else {
		keyName := keyConfig.Name
		if keyName != "" && keyConfig.KeyFrom != keyName {
			err := errors.Errorf(
				"key 'from' opt is a name, but doesn't match given key name: %s != %s",
				keyConfig.KeyFrom, keyName,
			)

			return emptyCosmosAddress, nil, err
		}

		keyRecord, err = kb.Key(keyConfig.KeyFrom)
	}

	if err != nil {
		err = errors.Wrapf(
			ErrKeyRecordNotFound, "couldn't find an entry for the key '%s' in keybase: %s",
			keyConfig.KeyFrom, err.Error())

		return emptyCosmosAddress, nil, err
	}

	if err := checkKeyRecord(config, keyRecord); err != nil {
		return emptyCosmosAddress, nil, err
	}

	addr, err := keyRecord.GetAddress()
	if err != nil {
		return emptyCosmosAddress, nil, err
	}

	return addr, kb, nil
}

func findKeyInKeyring(kb cosmkeyring.Keyring, config *cosmosKeyringConfig, fromSpec string) (sdk.AccAddress, error) {
	var fromIsAddress bool

	addressFrom, err := sdk.AccAddressFromBech32(fromSpec)
	if err == nil {
		fromIsAddress = true
	}

	var keyRecord *cosmkeyring.Record
	if fromIsAddress {
		keyRecord, err = kb.KeyByAddress(addressFrom)
	} else {
		keyRecord, err = kb.Key(fromSpec)
	}

	if err != nil {
		err = errors.Wrapf(
			ErrKeyRecordNotFound, "couldn't find an entry for the key '%s' in keybase: %s",
			fromSpec, err.Error())

		return emptyCosmosAddress, err
	}

	if err := checkKeyRecord(config, keyRecord); err != nil {
		return emptyCosmosAddress, err
	}

	addr, err := keyRecord.GetAddress()
	if err != nil {
		return emptyCosmosAddress, err
	}

	return addr, nil
}

func checkKeyRecord(
	config *cosmosKeyringConfig,
	keyRecord *cosmkeyring.Record,
) error {
	switch keyType := keyRecord.GetType(); keyType {
	case cosmkeyring.TypeLocal:
		// kb has a key and it's totally usable
		return nil

	case cosmkeyring.TypeLedger:
		// the kb stores references to ledger keys, so we must explicitly
		// check that. kb doesn't know how to scan HD keys - they must be added manually before
		if config.UseLedger {
			return nil
		}
		err := errors.Wrapf(
			ErrKeyIncompatible,
			"'%s' key is a ledger reference, enable ledger option",
			keyRecord.Name,
		)
		return err

	case cosmkeyring.TypeOffline:
		err := errors.Wrapf(
			ErrKeyIncompatible,
			"'%s' key is an offline key, not supported yet",
			keyRecord.Name,
		)
		return err

	case cosmkeyring.TypeMulti:
		err := errors.Wrapf(
			ErrKeyIncompatible,
			"'%s' key is an multisig key, not supported yet",
			keyRecord.Name,
		)
		return err

	default:
		err := errors.Wrapf(
			ErrKeyIncompatible,
			"'%s' key  has unsupported type: %s",
			keyRecord.Name, keyType,
		)
		return err
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

	return n, err
}

// addFromPrivKey adds a PrivKey into temporary in-mem keyring.
// Allows to init Context when the key has been provided in plaintext and parsed.
func addFromPrivKey(kb cosmkeyring.Keyring, name string, privKey cryptotypes.PrivKey) error {
	tmpPhrase := randPhrase(64)
	armored := cosmcrypto.EncryptArmorPrivKey(privKey, tmpPhrase, privKey.Type())
	err := kb.ImportPrivKey(name, armored, tmpPhrase)
	if err != nil {
		err = errors.Wrapf(ErrCosmosKeyringImportFailed, "failed to import privkey: %s", err.Error())
		return err
	}

	return nil
}

func hexToBytes(str string) ([]byte, error) {
	data, err := hex.DecodeString(strings.TrimPrefix(str, "0x"))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func randPhrase(size int) string {
	buf := make([]byte, size)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}

	return string(buf)
}

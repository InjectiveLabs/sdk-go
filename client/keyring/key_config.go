package keyring

import (
	bip39 "github.com/cosmos/go-bip39"
	"github.com/pkg/errors"
)

type cosmosKeyConfig struct {
	Name          string
	KeyFrom       string
	KeyPassphrase string
	PrivKeyHex    string
	Mnemonic      string
}

// KeyConfigOpt defines a known cosmos keyring key option.
type KeyConfigOpt func(c *cosmosKeyConfig) error

// WithKeyFrom sets the key name to use for signing. Must exist in the provided keyring.
func WithKeyFrom(v string) KeyConfigOpt {
	return func(c *cosmosKeyConfig) error {
		if v != "" {
			c.KeyFrom = v
		}

		return nil
	}
}

// WithKeyPassphrase sets the passphrase for keyring files. Insecure option, use for testing only.
// The package will fallback to os.Stdin if this option was not provided, but pass is required.
func WithKeyPassphrase(v string) KeyConfigOpt {
	return func(c *cosmosKeyConfig) error {
		if v != "" {
			c.KeyPassphrase = v
		}

		return nil
	}
}

// WithPrivKeyHex allows to specify a private key as plaintext hex. Insecure option, use for testing only.
// The package will create a virtual keyring holding that key, to meet all the interfaces.
func WithPrivKeyHex(v string) KeyConfigOpt {
	return func(c *cosmosKeyConfig) error {
		if v != "" {
			c.PrivKeyHex = v
		}

		return nil
	}
}

// WithMnemonic allows to specify a mnemonic pharse as plaintext. Insecure option, use for testing only.
// The package will create a virtual keyring to derive the keys and meet all the interfaces.
func WithMnemonic(v string) KeyConfigOpt {
	return func(c *cosmosKeyConfig) error {
		if v != "" {
			if !bip39.IsMnemonicValid(v) {
				err := errors.New("provided mnemonic is not a valid BIP39 mnemonic")
				return err
			}

			c.Mnemonic = v
		}

		return nil
	}
}

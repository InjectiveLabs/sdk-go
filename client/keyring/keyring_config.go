package keyring

import (
	"github.com/pkg/errors"
)

// ConfigOpt defines a known cosmos keyring option.
type ConfigOpt func(c *cosmosKeyringConfig) error

type cosmosKeyringConfig struct {
	KeyringDir     string
	KeyringAppName string
	KeyringBackend Backend
	UseLedger      bool

	Keys       []*cosmosKeyConfig
	DefaultKey string
}

// Backend defines a known keyring backend name.
type Backend string

const (
	// BackendTest is a testing backend, no passphrases required.
	BackendTest Backend = "test"
	// BackendFile is a backend where keys are stored as encrypted files.
	BackendFile Backend = "file"
	// BackendOS is a backend where keys are stored in the OS key chain. Platform specific.
	BackendOS Backend = "os"
)

// WithKeyringDir option sets keyring path in the filesystem, useful when keyring backend is `file`.
func WithKeyringDir(v string) ConfigOpt {
	return func(c *cosmosKeyringConfig) error {
		if len(v) > 0 {
			c.KeyringDir = v
		}

		return nil
	}
}

// WithKeyringAppName option sets keyring application name (used by Cosmos to separate keyrings).
func WithKeyringAppName(v string) ConfigOpt {
	return func(c *cosmosKeyringConfig) error {
		if len(v) > 0 {
			c.KeyringAppName = v
		}

		return nil
	}
}

// WithKeyringBackend sets the keyring backend. Expected values: test, file, os.
func WithKeyringBackend(v Backend) ConfigOpt {
	return func(c *cosmosKeyringConfig) error {
		if len(v) > 0 {
			c.KeyringBackend = v
		}

		return nil
	}
}

// WithUseLedger sets the option to use hardware wallet, if available on the system.
func WithUseLedger(b bool) ConfigOpt {
	return func(c *cosmosKeyringConfig) error {
		c.UseLedger = b

		return nil
	}
}

// WithKey adds an unnamed key into the keyring, based on its individual options.
func WithKey(opts ...KeyConfigOpt) ConfigOpt {
	return func(c *cosmosKeyringConfig) error {
		config := &cosmosKeyConfig{}

		for optIdx, optFn := range opts {
			if err := optFn(config); err != nil {
				err = errors.Wrapf(ErrFailedToApplyKeyConfigOption, "key option #%d: %s", optIdx+1, err.Error())
				return err
			}
		}

		c.Keys = append(c.Keys, config)
		return nil
	}
}

// WithNamedKey adds a key into the keyring, based on its individual options, with a given name (alias).
func WithNamedKey(name string, opts ...KeyConfigOpt) ConfigOpt {
	return func(c *cosmosKeyringConfig) error {
		config := &cosmosKeyConfig{
			Name: name,
		}

		for optIdx, optFn := range opts {
			if err := optFn(config); err != nil {
				err = errors.Wrapf(ErrFailedToApplyKeyConfigOption, "key option #%d: %s", optIdx+1, err.Error())
				return err
			}
		}

		c.Keys = append(c.Keys, config)
		return nil
	}
}

// WithDefaultKey specifies the default key (name or address) to be fetched during keyring init.
// This key must exist in specified keys.
func WithDefaultKey(v string) ConfigOpt {
	return func(c *cosmosKeyringConfig) error {
		if len(v) > 0 {
			c.DefaultKey = v
		}

		return nil
	}
}

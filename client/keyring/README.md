# Injective Chain Keyring Helper

Creates a new keyring from a variety of options. See `ConfigOpt` and related options. This keyring helper allows to initialize Cosmos SDK keyring used for signing transactions.

It allows flexibly define a static configuration of keys, supports multiple pre-defined keys in the same keyring and allows to load keys from a file, derive from mnemonic or read plain private key bytes from a HEX string. Extremely useful for testing and local development, but also robust for production use cases.

## Usage

```go
NewCosmosKeyring(cdc codec.Codec, opts ...ConfigOpt) (sdk.AccAddress, cosmkeyring.Keyring, error)
```

**ConfigOpts:**

These options are global on the keyring level.

* `WithKeyringDir` option sets keyring path in the filesystem, useful when keyring backend is `file`.
* `WithKeyringAppName` option sets keyring application name (defaults to `injectived`)
* `WithKeyringBackend` sets the keyring backend. Expected values: `test`, `file`, `os`.
* `WithUseLedger` sets the option to use hardware wallet, if available on the system.

These options allow to add keys to the keyring during initialization.

* `WithKey` adds a single key to the keyring, without having alias name.
* `WithNamedKey` addes a single key to the keyring, with a name.
* `WithDefaultKey` sets a default key reference to use for signing (by name).

**KeyConfigOpts:**

These options are set per key.

* `WithKeyFrom` sets the key name to use for signing. Must exist in the provided keyring.
* `WithKeyPassphrase` sets the passphrase for keyring files. The package will fallback to `os.Stdin` if this option was not provided, but passphrase is required.
* `WithPrivKeyHex` allows to specify a private key as plain-text hex. Insecure option, use for testing only. The package will create a virtual keyring holding that key, to meet all the interfaces.
* `WithMnemonic` allows to specify a mnemonic pharse as plain-text hex. Insecure option, use for testing only. The package will create a virtual keyring to derive the keys and meet all the interfaces.

## Examples

Initialize an in-memory keyring with a private key hex:

```go
NewCosmosKeyring(
    cdc,
    WithKey(
        WithPrivKeyHex("e6888cb164d52e4880e08a8a5dbe69cd62f67fde3d5906f2c5c951be553b2267"),
        WithKeyFrom("sender"),
    ),
)
```

Initialize an in-memory keyring with a mnemonic phrase:

```go
NewCosmosKeyring(
    s.cdc,
    WithKey(
        WithMnemonic("real simple naive ....... love"),
        WithKeyFrom("sender"),
    ),
)
```

Real world use case of keyring initialization from CLI flags, with a single named key set as default:

```go
NewCosmosKeyring(
    cdc,
    WithKeyringDir(*keyringDir),
    WithKeyringAppName(*keyringAppName),
    WithKeyringBackend(Backend(*keyringBackend)),
    
    WithNamedKey(
        "dispatcher",
        WithKeyFrom(*dispatcherKeyFrom),
        WithKeyPassphrase(*dispatcherKeyPassphrase),
        WithPrivKeyHex(*dispatcherKeyPrivateHex),
        WithMnemonic(*dispatcherKeyMnemonic),
    ),

    WithDefaultKey(
        "dispatcher",
    ),
)
```

## Testing

```bash
go test -v -cover

PASS
coverage: 83.1% of statements
```

## Generating a Test Fixture

```bash
> cd testdata

> injectived keys --keyring-dir `pwd` --keyring-backend file add test
```

Passphrase should be `test12345678` for this fixture to work.

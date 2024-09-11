package keyring

import (
	"os"

	"github.com/InjectiveLabs/sdk-go/chain/crypto/hd"
	cosmkeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
)

func (s *KeyringTestSuite) TestErrCosmosKeyringCreationFailed() {
	requireT := s.Require()

	_, _, err := NewCosmosKeyring(
		s.cdc,
		WithKeyringBackend("kowabunga"),
		WithKey(
			WithKeyFrom(testAccAddressBech),
		),
	)

	requireT.ErrorIs(err, ErrCosmosKeyringCreationFailed)
}

func (s *KeyringTestSuite) TestErrFailedToApplyConfigOption() {
	requireT := s.Require()

	_, _, err := NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithMnemonic(`???`),
		),
	)

	requireT.ErrorIs(err, ErrFailedToApplyConfigOption)
}

func (s *KeyringTestSuite) TestErrHexFormatError() {
	requireT := s.Require()

	_, _, err := NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithPrivKeyHex("nothex"),
		),
	)

	requireT.ErrorIs(err, ErrHexFormatError)

	_, _, err = NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithMnemonic(testMnemonic),
			WithPrivKeyHex("nothex"),
		),
	)

	requireT.ErrorIs(err, ErrHexFormatError)
}

func (s *KeyringTestSuite) TestErrIncompatibleOptionsProvided() {
	requireT := s.Require()

	_, _, err := NewCosmosKeyring(
		s.cdc,
		WithUseLedger(true),
		WithKey(
			WithMnemonic(testMnemonic),
		),
	)

	requireT.ErrorIs(err, ErrIncompatibleOptionsProvided)

	_, _, err = NewCosmosKeyring(
		s.cdc,
		WithUseLedger(true),
		WithKey(
			WithPrivKeyHex(testPrivKeyHex),
		),
	)

	requireT.ErrorIs(err, ErrIncompatibleOptionsProvided)
}

func (s *KeyringTestSuite) TestErrInsufficientKeyDetails() {
	requireT := s.Require()

	_, _, err := NewCosmosKeyring(s.cdc)

	requireT.ErrorIs(err, ErrInsufficientKeyDetails)
}

func (s *KeyringTestSuite) TestErrKeyIncompatible() {
	requireT := s.Require()

	addr, kb, err := NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithPrivKeyHex(testPrivKeyHex),
		),
	)
	requireT.NoError(err)

	testRecord, err := kb.KeyByAddress(addr)
	requireT.NoError(err)
	testRecordPubKey, err := testRecord.GetPubKey()
	requireT.NoError(err)

	kbDir, err := os.MkdirTemp(os.TempDir(), "keyring-test-kbroot-*")
	requireT.NoError(err)
	s.T().Cleanup(func() {
		_ = os.RemoveAll(kbDir)
	})

	testKeyring, err := cosmkeyring.New(
		KeyringAppName,
		cosmkeyring.BackendTest,
		kbDir,
		nil,
		s.cdc,
		hd.EthSecp256k1Option(),
	)
	requireT.NoError(err)

	_, err = testKeyring.SaveOfflineKey("test_pubkey", testRecordPubKey)
	requireT.NoError(err)

	_, _, err = NewCosmosKeyring(
		s.cdc,
		WithKeyringBackend(BackendTest),
		WithKeyringDir(kbDir),
		WithKeyringAppName(KeyringAppName),
		WithKey(
			WithKeyFrom("test_pubkey"),
		),
	)
	requireT.ErrorIs(err, ErrKeyIncompatible)

	// TODO: add test for unsupported multisig keys
}

func (s *KeyringTestSuite) TestErrKeyRecordNotFound() {
	requireT := s.Require()

	_, _, err := NewCosmosKeyring(
		s.cdc,
		WithKeyringBackend(BackendFile),
		WithKeyringDir("./testdata"),
		WithKey(
			WithKeyFrom("kowabunga"),
			WithKeyPassphrase("test12345678"),
		),
	)

	requireT.ErrorIs(err, ErrKeyRecordNotFound)
}

func (s *KeyringTestSuite) TestErrPrivkeyConflict() {
	requireT := s.Require()

	_, _, err := NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithPrivKeyHex(testOtherPrivKeyHex),
			WithMnemonic(testMnemonic), // different mnemonic
		),
	)

	requireT.ErrorIs(err, ErrPrivkeyConflict)
}

func (s *KeyringTestSuite) TestErrUnexpectedAddress() {
	requireT := s.Require()

	_, _, err := NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithPrivKeyHex(testOtherPrivKeyHex),
			WithKeyFrom(testAccAddressBech), // will not match privkey above
		),
	)

	requireT.ErrorIs(err, ErrUnexpectedAddress)

	_, _, err = NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithMnemonic(testMnemonic),
			WithKeyFrom("inj1xypj9l9sjdaduaafhgx39ru70utnzfuklcpxz9"), // will not match mnemonic above
		),
	)

	requireT.ErrorIs(err, ErrUnexpectedAddress)
}

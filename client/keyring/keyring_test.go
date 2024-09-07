package keyring

import (
	"encoding/hex"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	cosmcrypto "github.com/cosmos/cosmos-sdk/crypto"
	cosmkeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	crypto_cdc "github.com/InjectiveLabs/sdk-go/chain/crypto/codec"
	"github.com/InjectiveLabs/sdk-go/chain/crypto/hd"
	ctypes "github.com/InjectiveLabs/sdk-go/chain/types"
	"github.com/InjectiveLabs/sdk-go/client/chain"
)

type KeyringTestSuite struct {
	suite.Suite

	cdc codec.Codec
}

func TestKeyringTestSuite(t *testing.T) {
	suite.Run(t, new(KeyringTestSuite))
}

func getCryptoCodec() *codec.ProtoCodec {
	registry := chain.NewInterfaceRegistry()
	crypto_cdc.RegisterInterfaces(registry)
	return codec.NewProtoCodec(registry)
}

func (s *KeyringTestSuite) SetupTest() {
	config := sdk.GetConfig()
	ctypes.SetBech32Prefixes(config)
	ctypes.SetBip44CoinType(config)

	s.cdc = getCryptoCodec()
}

func (s *KeyringTestSuite) TestKeyFromPrivkey() {
	requireT := s.Require()

	accAddr, kb, err := NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithPrivKeyHex(testPrivKeyHex),
			WithKeyFrom(testAccAddressBech), // must match the privkey above
		),
	)
	requireT.NoError(err)
	requireT.Equal(testAccAddressBech, accAddr.String())

	record, err := kb.KeyByAddress(accAddr)
	requireT.NoError(err)
	requireT.Equal(cosmkeyring.TypeLocal, record.GetType())
	requireT.Equal(expectedPubKeyType, record.PubKey.TypeUrl)
	recordPubKey, err := record.GetPubKey()
	requireT.NoError(err)

	logPrivKey(s.T(), kb, accAddr)

	res, pubkey, err := kb.SignByAddress(accAddr, []byte("test"), signing.SignMode_SIGN_MODE_DIRECT)
	requireT.NoError(err)
	requireT.EqualValues(recordPubKey, pubkey)
	requireT.Equal(testSig, res)
}

func (s *KeyringTestSuite) TestKeyFromMnemonic() {
	requireT := s.Require()

	accAddr, kb, err := NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithMnemonic(testMnemonic),
			WithPrivKeyHex(testPrivKeyHex),  // must match mnemonic above
			WithKeyFrom(testAccAddressBech), // must match mnemonic above
		),
	)
	requireT.NoError(err)
	requireT.Equal(testAccAddressBech, accAddr.String())

	record, err := kb.KeyByAddress(accAddr)
	requireT.NoError(err)
	requireT.Equal(cosmkeyring.TypeLocal, record.GetType())
	requireT.Equal(expectedPubKeyType, record.PubKey.TypeUrl)
	recordPubKey, err := record.GetPubKey()
	requireT.NoError(err)

	logPrivKey(s.T(), kb, accAddr)

	res, pubkey, err := kb.SignByAddress(accAddr, []byte("test"), signing.SignMode_SIGN_MODE_DIRECT)
	requireT.NoError(err)
	requireT.Equal(recordPubKey, pubkey)
	requireT.Equal(testSig, res)
}

func (s *KeyringTestSuite) TestKeyringFile() {
	requireT := s.Require()

	accAddr, _, err := NewCosmosKeyring(
		s.cdc,
		WithKeyringBackend(BackendFile),
		WithKeyringDir("./testdata"),
		WithKey(
			WithKeyFrom("test"),
			WithKeyPassphrase("test12345678"),
		),
	)
	requireT.NoError(err)
	requireT.Equal(testAccAddressBech, accAddr.String())

	accAddr, kb, err := NewCosmosKeyring(
		s.cdc,
		WithKeyringBackend(BackendFile),
		WithKeyringDir("./testdata"),
		WithKey(
			WithKeyFrom(testAccAddressBech),
			WithKeyPassphrase("test12345678"),
		),
	)
	requireT.NoError(err)
	requireT.Equal(testAccAddressBech, accAddr.String())

	record, err := kb.KeyByAddress(accAddr)
	requireT.NoError(err)
	requireT.Equal(cosmkeyring.TypeLocal, record.GetType())
	requireT.Equal(expectedPubKeyType, record.PubKey.TypeUrl)
	requireT.Equal("test", record.Name)
	recordPubKey, err := record.GetPubKey()
	requireT.NoError(err)

	logPrivKey(s.T(), kb, accAddr)

	res, pubkey, err := kb.SignByAddress(accAddr, []byte("test"), signing.SignMode_SIGN_MODE_DIRECT)
	requireT.NoError(err)
	requireT.Equal(recordPubKey, pubkey)
	requireT.Equal(testSig, res)
}

func (s *KeyringTestSuite) TestKeyringOsWithAppName() {
	if testing.Short() {
		s.T().Skip("skipping testing in short mode")
		return
	}

	requireT := require.New(s.T())

	osKeyring, err := cosmkeyring.New(
		"keyring_test",
		cosmkeyring.BackendOS,
		"",
		nil,
		s.cdc,
		hd.EthSecp256k1Option(),
	)
	requireT.NoError(err)

	var accRecord *cosmkeyring.Record
	if accRecord, err = osKeyring.Key("test"); err != nil {
		accRecord, err = osKeyring.NewAccount(
			"test",
			testMnemonic,
			cosmkeyring.DefaultBIP39Passphrase,
			sdk.GetConfig().GetFullBIP44Path(),
			hd.EthSecp256k1,
		)

		requireT.NoError(err)

		accAddr, err := accRecord.GetAddress()
		requireT.NoError(err)
		requireT.Equal(testAccAddressBech, accAddr.String())
	}

	s.T().Cleanup(func() {
		// cleanup
		addr, err := accRecord.GetAddress()
		if err == nil {
			_ = osKeyring.DeleteByAddress(addr)
		}
	})

	accAddr, kb, err := NewCosmosKeyring(
		s.cdc,
		WithKeyringBackend(BackendOS),
		WithKeyringAppName("keyring_test"),
		WithKey(
			WithKeyFrom("test"),
		),
	)
	requireT.NoError(err)
	requireT.Equal(testAccAddressBech, accAddr.String())

	record, err := kb.KeyByAddress(accAddr)
	requireT.NoError(err)
	requireT.Equal(cosmkeyring.TypeLocal, record.GetType())
	requireT.Equal(expectedPubKeyType, record.PubKey.TypeUrl)
	recordPubKey, err := record.GetPubKey()
	requireT.NoError(err)

	requireT.Equal("test", record.Name)

	res, pubkey, err := kb.SignByAddress(accAddr, []byte("test"), signing.SignMode_SIGN_MODE_DIRECT)
	requireT.NoError(err)
	requireT.Equal(recordPubKey, pubkey)
	requireT.Equal(testSig, res)
}

func (s *KeyringTestSuite) TestUseFromAsName() {
	requireT := s.Require()

	accAddr, _, err := NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithPrivKeyHex(testPrivKeyHex),
			WithKeyFrom("kowabunga"),
		),
		WithDefaultKey("kowabunga"),
	)
	requireT.NoError(err)
	requireT.Equal(testAccAddressBech, accAddr.String())

	accAddr, _, err = NewCosmosKeyring(
		s.cdc,
		WithKey(
			WithMnemonic(testMnemonic),
			WithKeyFrom("kowabunga"),
		),
		WithDefaultKey("kowabunga"),
	)
	requireT.NoError(err)
	requireT.Equal(testAccAddressBech, accAddr.String())
}

func (s *KeyringTestSuite) TestNamedKeys() {
	requireT := s.Require()

	accAddr, kb, err := NewCosmosKeyring(
		s.cdc,
		WithNamedKey(
			"bad",
			WithPrivKeyHex(testOtherPrivKeyHex),
		),

		WithNamedKey(
			"good",
			WithPrivKeyHex(testPrivKeyHex),
		),

		WithDefaultKey("good"),
	)

	requireT.NoError(err)
	requireT.Equal(testAccAddressBech, accAddr.String())

	record, err := kb.KeyByAddress(accAddr)
	requireT.NoError(err)
	requireT.Equal(cosmkeyring.TypeLocal, record.GetType())
	requireT.Equal(expectedPubKeyType, record.PubKey.TypeUrl)
	recordPubKey, err := record.GetPubKey()
	requireT.NoError(err)

	logPrivKey(s.T(), kb, accAddr)

	res, pubkey, err := kb.SignByAddress(accAddr, []byte("test"), signing.SignMode_SIGN_MODE_DIRECT)
	requireT.NoError(err)
	requireT.EqualValues(recordPubKey, pubkey)
	requireT.Equal(testSig, res)
}

const expectedPubKeyType = "/injective.crypto.v1beta1.ethsecp256k1.PubKey"

const testAccAddressBech = "inj1ycc302kea06htx5zw2kj4eyk3hgj63sz206fq0"

//nolint:lll // mnemonic fixture
const testMnemonic = `real simple naive tissue alcohol bar short joy maze shoe reason item tray attitude panda century pulse skirt original autumn sea shop exhaust love`

var testPrivKeyHex = "e6888cb164d52e4880e08a8a5dbe69cd62f67fde3d5906f2c5c951be553b2267"
var testOtherPrivKeyHex = "ef3bc8bc1e1bae12268e0192787673a4137af840bfcbd1aa4c535bbd95fe6837"

var testSig = []byte{
	0xf9, 0x04, 0x3e, 0x81, 0x83, 0xb2, 0x73, 0xf6,
	0xdd, 0xf7, 0xd6, 0x91, 0x6f, 0xb5, 0x63, 0xf4,
	0x8a, 0xa2, 0x4a, 0x51, 0x63, 0xe1, 0x04, 0x18,
	0xd2, 0xe6, 0xed, 0x9e, 0xda, 0x52, 0x2f, 0x0a,
	0x69, 0x74, 0x04, 0x73, 0x7b, 0x9a, 0xf1, 0xc8,
	0xdf, 0xe7, 0xf3, 0x4a, 0x48, 0xe6, 0x5f, 0xc0,
	0x69, 0x5e, 0x6e, 0x03, 0x9e, 0x6e, 0x5f, 0x31,
	0xa6, 0x40, 0x19, 0x1b, 0x76, 0x07, 0xd9, 0x65,
	0x00,
}

func logPrivKey(t *testing.T, kb cosmkeyring.Keyring, accAddr sdk.AccAddress) {
	armor, _ := kb.ExportPrivKeyArmorByAddress(accAddr, "")
	privKey, _, _ := cosmcrypto.UnarmorDecryptPrivKey(armor, "")
	t.Log("[PRIV]", hex.EncodeToString(privKey.Bytes()))
}

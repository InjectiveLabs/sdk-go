package chain

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	eth "github.com/ethereum/go-ethereum/common"
)

func accountForTests() (cosmtypes.AccAddress, keyring.Keyring, error) {
	senderAddress, cosmosKeyring, err := InitCosmosKeyring(
		os.Getenv("HOME")+"/.injectived",
		"injectived",
		"file",
		"inj-user",
		"12345678",
		"5d386fbdbf11f1141010f81a46b40f94887367562bd33b452bbaa6ce1cd1381e", // keyring will be used if pk not provided
		false,
	)

	return senderAddress, cosmosKeyring, err
}

func createClient(senderAddress cosmtypes.AccAddress, cosmosKeyring keyring.Keyring, network common.Network) (ChainClient, error) {
	tmClient, _ := rpchttp.New(network.TmEndpoint, "/websocket")
	clientCtx, err := NewClientContext(
		network.ChainId,
		senderAddress.String(),
		cosmosKeyring,
	)

	if err != nil {
		return nil, err
	}

	clientCtx = clientCtx.WithNodeURI(network.TmEndpoint).WithClient(tmClient)
	// configure Keyring as nil to avoid the account initialization request when running unit tests
	clientCtx.Keyring = nil

	chainClient, err := NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	return chainClient, err
}

type OfacTestSuite struct {
	suite.Suite
	network       common.Network
	tmClient      *rpchttp.HTTP
	senderAddress cosmtypes.AccAddress
	cosmosKeyring keyring.Keyring
}

func (suite *OfacTestSuite) SetupTest() {
	var err error
	suite.network = common.LoadNetwork("testnet", "lb")
	suite.tmClient, err = rpchttp.New(suite.network.TmEndpoint, "/websocket")
	suite.NoError(err)

	suite.senderAddress, suite.cosmosKeyring, err = accountForTests()
	suite.NoError(err)

	// Prepare OFAC list file
	testList := []string{
		suite.senderAddress.String(),
	}
	jsonData, err := json.Marshal(testList)
	suite.NoError(err)

	ofacListFilename = "ofac_test.json"
	file, err := os.Create(getOfacListPath())
	suite.NoError(err)

	_, err = io.WriteString(file, string(jsonData))
	suite.NoError(err)

	err = file.Close()
	suite.NoError(err)
}

func (suite *OfacTestSuite) TearDownTest() {
	err := os.Remove(getOfacListPath())
	suite.NoError(err)
	ofacListFilename = defaultofacListFilename
}

func (suite *OfacTestSuite) TestOfacList() {
	clientCtx, err := NewClientContext(
		suite.network.ChainId,
		suite.senderAddress.String(),
		suite.cosmosKeyring,
	)
	suite.NoError(err)

	clientCtx = clientCtx.WithNodeURI(suite.network.TmEndpoint).WithClient(suite.tmClient)
	testChecker, err := NewOfacChecker()
	suite.NoError(err)
	suite.Equal(1, len(testChecker.ofacList))

	_, err = NewChainClient(
		clientCtx,
		suite.network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)
	suite.Error(err)
}

func TestOfacTestSuite(t *testing.T) {
	suite.Run(t, new(OfacTestSuite))
}

func TestDefaultSubaccount(t *testing.T) {
	network := common.LoadNetwork("devnet", "lb")
	senderAddress, cosmosKeyring, err := accountForTests()

	if err != nil {
		t.Errorf("Error creating the address %v", err)
	}

	chainClient, err := createClient(senderAddress, cosmosKeyring, network)

	if err != nil {
		t.Errorf("Error creating the client %v", err)
	}

	defaultSubaccountID := chainClient.DefaultSubaccount(senderAddress)

	expectedSubaccountId := "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000000"
	expectedSubaccountIdHash := eth.HexToHash(expectedSubaccountId)
	if defaultSubaccountID != expectedSubaccountIdHash {
		t.Error("The default subaccount is calculated incorrectly")
	}
}

func TestGetSubaccountWithIndex(t *testing.T) {
	network := common.LoadNetwork("devnet", "lb")
	senderAddress, cosmosKeyring, err := accountForTests()

	if err != nil {
		t.Errorf("Error creating the address %v", err)
	}

	chainClient, err := createClient(senderAddress, cosmosKeyring, network)

	if err != nil {
		t.Errorf("Error creating the client %v", err)
	}

	subaccountOne := chainClient.Subaccount(senderAddress, 1)
	subaccountThirty := chainClient.Subaccount(senderAddress, 30)

	expectedSubaccounOnetId := "0xaf79152ac5df276d9a8e1e2e22822f9713474902000000000000000000000001"
	expectedSubaccountOneIdHash := eth.HexToHash(expectedSubaccounOnetId)

	expectedSubaccounThirtytId := "0xaf79152ac5df276d9a8e1e2e22822f971347490200000000000000000000001e"
	expectedSubaccountThirtyIdHash := eth.HexToHash(expectedSubaccounThirtytId)

	if subaccountOne != expectedSubaccountOneIdHash {
		t.Error("The subaccount with index 1 was calculated incorrectly")
	}
	if subaccountThirty != expectedSubaccountThirtyIdHash {
		t.Error("The subaccount with index 30 was calculated incorrectly")
	}
}

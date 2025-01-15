package chain_test

import (
	"encoding/json"
	"os"
	"testing"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	eth "github.com/ethereum/go-ethereum/common"

	"github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

var (
	originalOfacListPath string
)

func setUpChainTest(t *testing.T) {
	// Store the original OfacListPath
	originalOfacListPath = chain.OfacListPath

	// Update OfacListPath to point to the temporary directory
	chain.OfacListPath = t.TempDir()

	// Get the sender address from accountForTests
	senderAddress, _, err := accountForTests()
	if err != nil {
		t.Fatalf("Failed to get sender address: %v", err)
	}

	// Create the OFAC list JSON file with the sender address
	ofacListPath := chain.GetOfacListPath()
	ofacList := []string{senderAddress.String()}

	// Ensure the directory exists
	if err := os.MkdirAll(chain.OfacListPath, 0755); err != nil {
		t.Fatalf("Failed to create OFAC list directory: %v", err)
	}

	// Create the OFAC list file
	file, err := os.Create(ofacListPath)
	if err != nil {
		t.Fatalf("Failed to create OFAC list file: %v", err)
	}
	defer file.Close()

	// Encode the list as JSON
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(ofacList); err != nil {
		t.Fatalf("Failed to write OFAC list: %v", err)
	}
}

func tearDownChainTest() {
	// Restore the original OfacListPath
	chain.OfacListPath = originalOfacListPath
}

func TestDefaultSubaccount(t *testing.T) {
	setUpChainTest(t)
	defer tearDownChainTest()

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
	setUpChainTest(t)
	defer tearDownChainTest()

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
func accountForTests() (cosmtypes.AccAddress, keyring.Keyring, error) {
	senderAddress, cosmosKeyring, err := chain.InitCosmosKeyring(
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

func createClient(senderAddress cosmtypes.AccAddress, cosmosKeyring keyring.Keyring, network common.Network) (chain.ChainClient, error) {
	tmClient, _ := rpchttp.New(network.TmEndpoint, "/websocket")
	clientCtx, err := chain.NewClientContext(
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

	chainClient, err := chain.NewChainClient(
		clientCtx,
		network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	return chainClient, err
}

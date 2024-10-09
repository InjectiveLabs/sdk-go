package chain_test

import (
	"encoding/json"
	"github.com/InjectiveLabs/sdk-go/client"
	"github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmtypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"io"
	"os"
	"testing"
)

type OfacTestSuite struct {
	suite.Suite
	network          common.Network
	tmClient         *rpchttp.HTTP
	senderAddress    cosmtypes.AccAddress
	cosmosKeyring    keyring.Keyring
	originalOfacPath string
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

	suite.originalOfacPath = chain.OfacListPath
	chain.OfacListPath = suite.T().TempDir()
	suite.NoError(err)
	file, err := os.Create(chain.GetOfacListPath())
	suite.NoError(err)

	_, err = io.WriteString(file, string(jsonData))
	suite.NoError(err)

	err = file.Close()
	suite.NoError(err)
}

func (suite *OfacTestSuite) TearDownTest() {
	chain.OfacListPath = suite.originalOfacPath
}

func (suite *OfacTestSuite) TestOfacList() {
	clientCtx, err := chain.NewClientContext(
		suite.network.ChainId,
		suite.senderAddress.String(),
		suite.cosmosKeyring,
	)
	suite.NoError(err)

	clientCtx = clientCtx.WithNodeURI(suite.network.TmEndpoint).WithClient(suite.tmClient)
	testChecker, err := chain.NewOfacChecker()
	suite.NoError(err)
	suite.True(testChecker.IsBlacklisted(suite.senderAddress.String()))
	suite.False(testChecker.IsBlacklisted("inj1"))

	_, err = chain.NewChainClient(
		clientCtx,
		suite.network,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)
	suite.Error(err)
}

func TestOfacTestSuite(t *testing.T) {
	suite.Run(t, new(OfacTestSuite))
}

package tm

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
)

type TendermintClient interface {
	GetBlock(ctx context.Context, height int64) (*tmctypes.ResultBlock, error)
	GetLatestBlockHeight(ctx context.Context) (int64, error)
	GetTxs(ctx context.Context, block *tmctypes.ResultBlock) ([]*ctypes.ResultTx, error)
	GetBlockResults(ctx context.Context, height int64) (*ctypes.ResultBlockResults, error)
	GetValidatorSet(ctx context.Context, height int64) (*tmctypes.ResultValidators, error)
	GetABCIInfo(ctx context.Context) (*ctypes.ResultABCIInfo, error)
}

type tmClient struct {
	rpcClient rpcclient.Client
	logger    *logrus.Logger
}

func NewRPCClient(rpcNodeAddr string, logger *logrus.Logger) TendermintClient {
	rpcClient, err := rpchttp.NewWithTimeout(rpcNodeAddr, "/websocket", 10)
	if err != nil {
		logger.Errorln("[INJ-GO-SDK] Failed to init rpcClient: ", err)
	}

	return &tmClient{
		rpcClient: rpcClient,
		logger:    logger,
	}
}

// GetBlock queries for a block by height. An error is returned if the query fails.
func (c *tmClient) GetBlock(ctx context.Context, height int64) (*tmctypes.ResultBlock, error) {
	return c.rpcClient.Block(ctx, &height)
}

// GetBlock queries for a block by height. An error is returned if the query fails.
func (c *tmClient) GetBlockResults(ctx context.Context, height int64) (*ctypes.ResultBlockResults, error) {
	return c.rpcClient.BlockResults(ctx, &height)
}

// GetLatestBlockHeight returns the latest block height on the active chain.
func (c *tmClient) GetLatestBlockHeight(ctx context.Context) (int64, error) {
	status, err := c.rpcClient.Status(ctx)
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight

	return height, nil
}

// GetTxs queries for all the transactions in a block height.
// It uses `Tx` RPC method to query for the transaction.
func (c *tmClient) GetTxs(ctx context.Context, block *tmctypes.ResultBlock) ([]*ctypes.ResultTx, error) {
	txs := make([]*ctypes.ResultTx, 0, len(block.Block.Txs))

	for _, tmTx := range block.Block.Txs {
		tx, err := c.rpcClient.Tx(ctx, tmTx.Hash(), true)
		if err != nil {
			if strings.HasSuffix(err.Error(), "not found") {
				c.logger.Errorln("[INJ-GO-SDK] Failed to get Tx by hash: ", err)
				continue
			}

			return nil, err
		}

		txs = append(txs, tx)
	}

	return txs, nil
}

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c *tmClient) GetValidatorSet(ctx context.Context, height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(ctx, &height, nil, nil)
}

// GetABCIInfo returns the node abci version
func (c *tmClient) GetABCIInfo(ctx context.Context) (*tmctypes.ResultABCIInfo, error) {
	return c.rpcClient.ABCIInfo(ctx)
}

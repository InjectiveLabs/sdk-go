package tm

import (
	"context"
	"strings"

	log "github.com/InjectiveLabs/suplog"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
)

type TendermintClient interface {
	GetBlock(ctx context.Context, height int64) (*ctypes.ResultBlock, error)
	GetLatestBlockHeight(ctx context.Context) (int64, error)
	GetTxs(ctx context.Context, block *ctypes.ResultBlock) ([]*ctypes.ResultTx, error)
	GetBlockResults(ctx context.Context, height int64) (*ctypes.ResultBlockResults, error)
	GetValidatorSet(ctx context.Context, height int64) (*ctypes.ResultValidators, error)
	GetABCIInfo(ctx context.Context) (*ctypes.ResultABCIInfo, error)
	GetStatus(ctx context.Context) (*ctypes.ResultStatus, error)
}

type tmClient struct {
	rpcClient rpcclient.Client
}

func NewRPCClient(rpcNodeAddr string) TendermintClient {
	rpcClient, err := rpchttp.NewWithTimeout(rpcNodeAddr, "/websocket", 10)
	if err != nil {
		log.WithError(err).Fatalln("failed to init rpcClient")
	}

	return &tmClient{
		rpcClient: rpcClient,
	}
}

// GetBlock queries for a block by height. An error is returned if the query fails.
func (c *tmClient) GetBlock(ctx context.Context, height int64) (*ctypes.ResultBlock, error) {
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
func (c *tmClient) GetTxs(ctx context.Context, block *ctypes.ResultBlock) ([]*ctypes.ResultTx, error) {
	txs := make([]*ctypes.ResultTx, 0, len(block.Block.Txs))

	for _, tmTx := range block.Block.Txs {
		tx, err := c.rpcClient.Tx(ctx, tmTx.Hash(), true)
		if err != nil {
			if strings.HasSuffix(err.Error(), "not found") {
				log.WithError(err).Errorln("failed to get Tx by hash")
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
func (c *tmClient) GetValidatorSet(ctx context.Context, height int64) (*ctypes.ResultValidators, error) {
	return c.rpcClient.Validators(ctx, &height, nil, nil)
}

// GetABCIInfo returns the node abci version
func (c *tmClient) GetABCIInfo(ctx context.Context) (*ctypes.ResultABCIInfo, error) {
	return c.rpcClient.ABCIInfo(ctx)
}

// GetStatus returns the node status.
func (c *tmClient) GetStatus(ctx context.Context) (*ctypes.ResultStatus, error) {
	return c.rpcClient.Status(ctx)
}

package explorer

import (
	"context"

	log "github.com/InjectiveLabs/suplog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/InjectiveLabs/sdk-go/client/common"
	explorerPB "github.com/InjectiveLabs/sdk-go/exchange/explorer_rpc/pb"
)

type ExplorerClient interface {
	QueryClient() *grpc.ClientConn
	GetAccountTxs(ctx context.Context, req *explorerPB.GetAccountTxsRequest) (*explorerPB.GetAccountTxsResponse, error)
	FetchContractTxs(ctx context.Context, req *explorerPB.GetContractTxsRequest) (*explorerPB.GetContractTxsResponse, error)
	FetchContractTxsV2(ctx context.Context, req *explorerPB.GetContractTxsV2Request) (*explorerPB.GetContractTxsV2Response, error)
	GetBlocks(ctx context.Context) (*explorerPB.GetBlocksResponse, error)
	GetBlock(ctx context.Context, blockHeight string) (*explorerPB.GetBlockResponse, error)
	FetchValidators(ctx context.Context) (*explorerPB.GetValidatorsResponse, error)
	FetchValidator(ctx context.Context, address string) (*explorerPB.GetValidatorResponse, error)
	FetchValidatorUptime(ctx context.Context, address string) (*explorerPB.GetValidatorUptimeResponse, error)
	GetTxs(ctx context.Context, req *explorerPB.GetTxsRequest) (*explorerPB.GetTxsResponse, error)
	GetTxByTxHash(ctx context.Context, hash string) (*explorerPB.GetTxByTxHashResponse, error)
	GetPeggyDeposits(ctx context.Context, req *explorerPB.GetPeggyDepositTxsRequest) (*explorerPB.GetPeggyDepositTxsResponse, error)
	GetPeggyWithdrawals(ctx context.Context, req *explorerPB.GetPeggyWithdrawalTxsRequest) (*explorerPB.GetPeggyWithdrawalTxsResponse, error)
	GetIBCTransfers(ctx context.Context, req *explorerPB.GetIBCTransferTxsRequest) (*explorerPB.GetIBCTransferTxsResponse, error)
	GetWasmCodes(ctx context.Context, req *explorerPB.GetWasmCodesRequest) (*explorerPB.GetWasmCodesResponse, error)
	GetWasmCodeByID(ctx context.Context, req *explorerPB.GetWasmCodeByIDRequest) (*explorerPB.GetWasmCodeByIDResponse, error)
	GetWasmContracts(ctx context.Context, req *explorerPB.GetWasmContractsRequest) (*explorerPB.GetWasmContractsResponse, error)
	GetWasmContractByAddress(ctx context.Context, req *explorerPB.GetWasmContractByAddressRequest) (*explorerPB.GetWasmContractByAddressResponse, error)
	GetCW20Balance(ctx context.Context, req *explorerPB.GetCw20BalanceRequest) (*explorerPB.GetCw20BalanceResponse, error)
	FetchRelayers(ctx context.Context, marketIDs []string) (*explorerPB.RelayersResponse, error)
	FetchBankTransfers(ctx context.Context, req *explorerPB.GetBankTransfersRequest) (*explorerPB.GetBankTransfersResponse, error)
	StreamTxs(ctx context.Context) (explorerPB.InjectiveExplorerRPC_StreamTxsClient, error)
	StreamBlocks(ctx context.Context) (explorerPB.InjectiveExplorerRPC_StreamBlocksClient, error)
	Close()
}

func NewExplorerClient(network common.Network, options ...common.ClientOption) (ExplorerClient, error) {
	// process options
	opts := common.DefaultClientOptions()
	if network.ChainTLSCert != nil {
		options = append(options, common.OptionTLSCert(network.ExchangeTLSCert))
	}
	for _, opt := range options {
		if err := opt(opts); err != nil {
			err = errors.Wrap(err, "error in client option")
			return nil, err
		}
	}

	// create grpc client
	var conn *grpc.ClientConn
	var err error
	if opts.TLSCert != nil {
		conn, err = grpc.NewClient(network.ExplorerGrpcEndpoint, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		conn, err = grpc.NewClient(network.ExplorerGrpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(common.DialerFunc))
	}
	if err != nil {
		err := errors.Wrapf(err, "failed to connect to the gRPC: %s", network.ExplorerGrpcEndpoint)
		return nil, err
	}

	// build client
	cc := &explorerClient{
		opts:    opts,
		network: network,
		conn:    conn,

		explorerClient: explorerPB.NewInjectiveExplorerRPCClient(conn),
		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "exchangeClient",
		}),
	}

	return cc, nil
}

type explorerClient struct {
	opts    *common.ClientOptions
	network common.Network
	conn    *grpc.ClientConn
	logger  log.Logger

	explorerClient explorerPB.InjectiveExplorerRPCClient
}

func (c *explorerClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *explorerClient) GetTxByTxHash(ctx context.Context, hash string) (*explorerPB.GetTxByTxHashResponse, error) {
	req := explorerPB.GetTxByTxHashRequest{
		Hash: hash,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetTxByTxHash, &req)

	if err != nil {
		return &explorerPB.GetTxByTxHashResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetAccountTxs(ctx context.Context, req *explorerPB.GetAccountTxsRequest) (*explorerPB.GetAccountTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetAccountTxs, req)

	if err != nil {
		return &explorerPB.GetAccountTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetBlocks(ctx context.Context) (*explorerPB.GetBlocksResponse, error) {
	req := explorerPB.GetBlocksRequest{}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetBlocks, &req)

	if err != nil {
		return &explorerPB.GetBlocksResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetBlock(ctx context.Context, blockHeight string) (*explorerPB.GetBlockResponse, error) {
	req := explorerPB.GetBlockRequest{
		Id: blockHeight,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetBlock, &req)

	if err != nil {
		return &explorerPB.GetBlockResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchValidators(ctx context.Context) (*explorerPB.GetValidatorsResponse, error) {
	req := &explorerPB.GetValidatorsRequest{}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetValidators, req)

	if err != nil {
		return &explorerPB.GetValidatorsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchValidator(
	ctx context.Context,
	address string,
) (*explorerPB.GetValidatorResponse, error) {
	req := &explorerPB.GetValidatorRequest{
		Address: address,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetValidator, req)

	if err != nil {
		return &explorerPB.GetValidatorResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchValidatorUptime(
	ctx context.Context,
	address string,
) (*explorerPB.GetValidatorUptimeResponse, error) {
	req := &explorerPB.GetValidatorUptimeRequest{
		Address: address,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetValidatorUptime, req)

	if err != nil {
		return &explorerPB.GetValidatorUptimeResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetTxs(ctx context.Context, req *explorerPB.GetTxsRequest) (*explorerPB.GetTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetTxs, req)

	if err != nil {
		return &explorerPB.GetTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetPeggyDeposits(ctx context.Context, req *explorerPB.GetPeggyDepositTxsRequest) (*explorerPB.GetPeggyDepositTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetPeggyDepositTxs, req)

	if err != nil {
		return &explorerPB.GetPeggyDepositTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetPeggyWithdrawals(ctx context.Context, req *explorerPB.GetPeggyWithdrawalTxsRequest) (*explorerPB.GetPeggyWithdrawalTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetPeggyWithdrawalTxs, req)

	if err != nil {
		return &explorerPB.GetPeggyWithdrawalTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetIBCTransfers(ctx context.Context, req *explorerPB.GetIBCTransferTxsRequest) (*explorerPB.GetIBCTransferTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetIBCTransferTxs, req)

	if err != nil {
		return &explorerPB.GetIBCTransferTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) StreamTxs(ctx context.Context) (explorerPB.InjectiveExplorerRPC_StreamTxsClient, error) {
	req := explorerPB.StreamTxsRequest{}

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.StreamTxs, &req)

	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *explorerClient) StreamBlocks(ctx context.Context) (explorerPB.InjectiveExplorerRPC_StreamBlocksClient, error) {
	req := explorerPB.StreamBlocksRequest{}

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.StreamBlocks, &req)

	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *explorerClient) GetWasmCodes(ctx context.Context, req *explorerPB.GetWasmCodesRequest) (*explorerPB.GetWasmCodesResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetWasmCodes, req)

	if err != nil {
		return &explorerPB.GetWasmCodesResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetWasmCodeByID(ctx context.Context, req *explorerPB.GetWasmCodeByIDRequest) (*explorerPB.GetWasmCodeByIDResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetWasmCodeByID, req)

	if err != nil {
		return &explorerPB.GetWasmCodeByIDResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetWasmContracts(ctx context.Context, req *explorerPB.GetWasmContractsRequest) (*explorerPB.GetWasmContractsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetWasmContracts, req)

	if err != nil {
		return &explorerPB.GetWasmContractsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetWasmContractByAddress(ctx context.Context, req *explorerPB.GetWasmContractByAddressRequest) (*explorerPB.GetWasmContractByAddressResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetWasmContractByAddress, req)

	if err != nil {
		return &explorerPB.GetWasmContractByAddressResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetCW20Balance(ctx context.Context, req *explorerPB.GetCw20BalanceRequest) (*explorerPB.GetCw20BalanceResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetCw20Balance, req)

	if err != nil {
		return &explorerPB.GetCw20BalanceResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchContractTxs(
	ctx context.Context,
	req *explorerPB.GetContractTxsRequest,
) (*explorerPB.GetContractTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetContractTxs, req)

	if err != nil {
		return &explorerPB.GetContractTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchContractTxsV2(
	ctx context.Context,
	req *explorerPB.GetContractTxsV2Request,
) (*explorerPB.GetContractTxsV2Response, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetContractTxsV2, req)

	if err != nil {
		return &explorerPB.GetContractTxsV2Response{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchRelayers(
	ctx context.Context,
	marketIDs []string,
) (*explorerPB.RelayersResponse, error) {
	req := &explorerPB.RelayersRequest{
		MarketIDs: marketIDs,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.Relayers, req)

	if err != nil {
		return &explorerPB.RelayersResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchBankTransfers(
	ctx context.Context,
	req *explorerPB.GetBankTransfersRequest,
) (*explorerPB.GetBankTransfersResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetBankTransfers, req)

	if err != nil {
		return &explorerPB.GetBankTransfersResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) Close() {
	c.conn.Close()
}

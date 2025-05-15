package explorer

import (
	"context"

	log "github.com/InjectiveLabs/suplog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/InjectiveLabs/sdk-go/client/common"
	explorerpb "github.com/InjectiveLabs/sdk-go/exchange/explorer_rpc/pb"
)

type ExplorerClient interface {
	QueryClient() *grpc.ClientConn
	GetAccountTxs(ctx context.Context, req *explorerpb.GetAccountTxsRequest) (*explorerpb.GetAccountTxsResponse, error)
	FetchContractTxs(ctx context.Context, req *explorerpb.GetContractTxsRequest) (*explorerpb.GetContractTxsResponse, error)
	FetchContractTxsV2(ctx context.Context, req *explorerpb.GetContractTxsV2Request) (*explorerpb.GetContractTxsV2Response, error)
	GetBlocks(ctx context.Context) (*explorerpb.GetBlocksResponse, error)
	GetBlock(ctx context.Context, blockHeight string) (*explorerpb.GetBlockResponse, error)
	FetchValidators(ctx context.Context) (*explorerpb.GetValidatorsResponse, error)
	FetchValidator(ctx context.Context, address string) (*explorerpb.GetValidatorResponse, error)
	FetchValidatorUptime(ctx context.Context, address string) (*explorerpb.GetValidatorUptimeResponse, error)
	GetTxs(ctx context.Context, req *explorerpb.GetTxsRequest) (*explorerpb.GetTxsResponse, error)
	GetTxByTxHash(ctx context.Context, hash string) (*explorerpb.GetTxByTxHashResponse, error)
	GetPeggyDeposits(ctx context.Context, req *explorerpb.GetPeggyDepositTxsRequest) (*explorerpb.GetPeggyDepositTxsResponse, error)
	GetPeggyWithdrawals(
		ctx context.Context,
		req *explorerpb.GetPeggyWithdrawalTxsRequest,
	) (*explorerpb.GetPeggyWithdrawalTxsResponse, error)
	GetIBCTransfers(ctx context.Context, req *explorerpb.GetIBCTransferTxsRequest) (*explorerpb.GetIBCTransferTxsResponse, error)
	GetWasmCodes(ctx context.Context, req *explorerpb.GetWasmCodesRequest) (*explorerpb.GetWasmCodesResponse, error)
	GetWasmCodeByID(ctx context.Context, req *explorerpb.GetWasmCodeByIDRequest) (*explorerpb.GetWasmCodeByIDResponse, error)
	GetWasmContracts(ctx context.Context, req *explorerpb.GetWasmContractsRequest) (*explorerpb.GetWasmContractsResponse, error)
	GetWasmContractByAddress(
		ctx context.Context,
		req *explorerpb.GetWasmContractByAddressRequest,
	) (*explorerpb.GetWasmContractByAddressResponse, error)
	GetCW20Balance(ctx context.Context, req *explorerpb.GetCw20BalanceRequest) (*explorerpb.GetCw20BalanceResponse, error)
	FetchRelayers(ctx context.Context, marketIDs []string) (*explorerpb.RelayersResponse, error)
	FetchBankTransfers(ctx context.Context, req *explorerpb.GetBankTransfersRequest) (*explorerpb.GetBankTransfersResponse, error)
	StreamTxs(ctx context.Context) (explorerpb.InjectiveExplorerRPC_StreamTxsClient, error)
	StreamBlocks(ctx context.Context) (explorerpb.InjectiveExplorerRPC_StreamBlocksClient, error)
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
		conn, err = grpc.NewClient(
			network.ExplorerGrpcEndpoint,
			grpc.WithTransportCredentials(opts.TLSCert),
			grpc.WithContextDialer(common.DialerFunc),
		)
	} else {
		conn, err = grpc.NewClient(
			network.ExplorerGrpcEndpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(common.DialerFunc),
		)
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

		explorerClient: explorerpb.NewInjectiveExplorerRPCClient(conn),
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

	explorerClient explorerpb.InjectiveExplorerRPCClient
}

func (c *explorerClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *explorerClient) GetTxByTxHash(ctx context.Context, hash string) (*explorerpb.GetTxByTxHashResponse, error) {
	req := explorerpb.GetTxByTxHashRequest{
		Hash: hash,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetTxByTxHash, &req)

	if err != nil {
		return &explorerpb.GetTxByTxHashResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetAccountTxs(
	ctx context.Context,
	req *explorerpb.GetAccountTxsRequest,
) (*explorerpb.GetAccountTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetAccountTxs, req)

	if err != nil {
		return &explorerpb.GetAccountTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetBlocks(ctx context.Context) (*explorerpb.GetBlocksResponse, error) {
	req := explorerpb.GetBlocksRequest{}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetBlocks, &req)

	if err != nil {
		return &explorerpb.GetBlocksResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetBlock(ctx context.Context, blockHeight string) (*explorerpb.GetBlockResponse, error) {
	req := explorerpb.GetBlockRequest{
		Id: blockHeight,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetBlock, &req)

	if err != nil {
		return &explorerpb.GetBlockResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchValidators(ctx context.Context) (*explorerpb.GetValidatorsResponse, error) {
	req := &explorerpb.GetValidatorsRequest{}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetValidators, req)

	if err != nil {
		return &explorerpb.GetValidatorsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchValidator(
	ctx context.Context,
	address string,
) (*explorerpb.GetValidatorResponse, error) {
	req := &explorerpb.GetValidatorRequest{
		Address: address,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetValidator, req)

	if err != nil {
		return &explorerpb.GetValidatorResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchValidatorUptime(
	ctx context.Context,
	address string,
) (*explorerpb.GetValidatorUptimeResponse, error) {
	req := &explorerpb.GetValidatorUptimeRequest{
		Address: address,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetValidatorUptime, req)

	if err != nil {
		return &explorerpb.GetValidatorUptimeResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetTxs(ctx context.Context, req *explorerpb.GetTxsRequest) (*explorerpb.GetTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetTxs, req)

	if err != nil {
		return &explorerpb.GetTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetPeggyDeposits(
	ctx context.Context,
	req *explorerpb.GetPeggyDepositTxsRequest,
) (*explorerpb.GetPeggyDepositTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetPeggyDepositTxs, req)

	if err != nil {
		return &explorerpb.GetPeggyDepositTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetPeggyWithdrawals(
	ctx context.Context,
	req *explorerpb.GetPeggyWithdrawalTxsRequest,
) (*explorerpb.GetPeggyWithdrawalTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetPeggyWithdrawalTxs, req)

	if err != nil {
		return &explorerpb.GetPeggyWithdrawalTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetIBCTransfers(
	ctx context.Context,
	req *explorerpb.GetIBCTransferTxsRequest,
) (*explorerpb.GetIBCTransferTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetIBCTransferTxs, req)

	if err != nil {
		return &explorerpb.GetIBCTransferTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) StreamTxs(ctx context.Context) (explorerpb.InjectiveExplorerRPC_StreamTxsClient, error) {
	req := explorerpb.StreamTxsRequest{}

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.StreamTxs, &req)

	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *explorerClient) StreamBlocks(ctx context.Context) (explorerpb.InjectiveExplorerRPC_StreamBlocksClient, error) {
	req := explorerpb.StreamBlocksRequest{}

	stream, err := common.ExecuteStreamCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.StreamBlocks, &req)

	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *explorerClient) GetWasmCodes(
	ctx context.Context,
	req *explorerpb.GetWasmCodesRequest,
) (*explorerpb.GetWasmCodesResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetWasmCodes, req)

	if err != nil {
		return &explorerpb.GetWasmCodesResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetWasmCodeByID(
	ctx context.Context,
	req *explorerpb.GetWasmCodeByIDRequest,
) (*explorerpb.GetWasmCodeByIDResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetWasmCodeByID, req)

	if err != nil {
		return &explorerpb.GetWasmCodeByIDResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetWasmContracts(
	ctx context.Context,
	req *explorerpb.GetWasmContractsRequest,
) (*explorerpb.GetWasmContractsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetWasmContracts, req)

	if err != nil {
		return &explorerpb.GetWasmContractsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetWasmContractByAddress(
	ctx context.Context,
	req *explorerpb.GetWasmContractByAddressRequest,
) (*explorerpb.GetWasmContractByAddressResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetWasmContractByAddress, req)

	if err != nil {
		return &explorerpb.GetWasmContractByAddressResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) GetCW20Balance(
	ctx context.Context,
	req *explorerpb.GetCw20BalanceRequest,
) (*explorerpb.GetCw20BalanceResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetCw20Balance, req)

	if err != nil {
		return &explorerpb.GetCw20BalanceResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchContractTxs(
	ctx context.Context,
	req *explorerpb.GetContractTxsRequest,
) (*explorerpb.GetContractTxsResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetContractTxs, req)

	if err != nil {
		return &explorerpb.GetContractTxsResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchContractTxsV2(
	ctx context.Context,
	req *explorerpb.GetContractTxsV2Request,
) (*explorerpb.GetContractTxsV2Response, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetContractTxsV2, req)

	if err != nil {
		return &explorerpb.GetContractTxsV2Response{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchRelayers(
	ctx context.Context,
	marketIDs []string,
) (*explorerpb.RelayersResponse, error) {
	req := &explorerpb.RelayersRequest{
		MarketIDs: marketIDs,
	}

	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.Relayers, req)

	if err != nil {
		return &explorerpb.RelayersResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) FetchBankTransfers(
	ctx context.Context,
	req *explorerpb.GetBankTransfersRequest,
) (*explorerpb.GetBankTransfersResponse, error) {
	res, err := common.ExecuteCall(ctx, c.network.ExplorerCookieAssistant, c.explorerClient.GetBankTransfers, req)

	if err != nil {
		return &explorerpb.GetBankTransfersResponse{}, err
	}

	return res, nil
}

func (c *explorerClient) Close() {
	c.conn.Close()
}

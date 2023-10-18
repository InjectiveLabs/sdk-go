package exchange

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/common"
	explorerPB "github.com/InjectiveLabs/sdk-go/exchange/explorer_rpc/pb"
	"google.golang.org/grpc/metadata"

	log "github.com/InjectiveLabs/suplog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type ExplorerClient interface {
	QueryClient() *grpc.ClientConn
	GetTxByTxHash(ctx context.Context, hash string) (explorerPB.GetTxByTxHashResponse, error)
	GetTxs(ctx context.Context, req explorerPB.GetTxsRequest) (explorerPB.GetTxsResponse, error)
	GetBlock(ctx context.Context, blockHeight string) (explorerPB.GetBlockResponse, error)
	GetBlocks(ctx context.Context) (explorerPB.GetBlocksResponse, error)
	GetAccountTxs(ctx context.Context, req explorerPB.GetAccountTxsRequest) (explorerPB.GetAccountTxsResponse, error)
	GetPeggyDeposits(ctx context.Context, req explorerPB.GetPeggyDepositTxsRequest) (explorerPB.GetPeggyDepositTxsResponse, error)
	GetPeggyWithdrawals(ctx context.Context, req explorerPB.GetPeggyWithdrawalTxsRequest) (explorerPB.GetPeggyWithdrawalTxsResponse, error)
	GetIBCTransfers(ctx context.Context, req explorerPB.GetIBCTransferTxsRequest) (explorerPB.GetIBCTransferTxsResponse, error)
	StreamTxs(ctx context.Context) (explorerPB.InjectiveExplorerRPC_StreamTxsClient, error)
	StreamBlocks(ctx context.Context) (explorerPB.InjectiveExplorerRPC_StreamBlocksClient, error)
	GetWasmCodes(ctx context.Context, req explorerPB.GetWasmCodesRequest) (explorerPB.GetWasmCodesResponse, error)
	GetWasmCodeByID(ctx context.Context, req explorerPB.GetWasmCodeByIDRequest) (explorerPB.GetWasmCodeByIDResponse, error)
	GetWasmContracts(ctx context.Context, req explorerPB.GetWasmContractsRequest) (explorerPB.GetWasmContractsResponse, error)
	GetWasmContractByAddress(ctx context.Context, req explorerPB.GetWasmContractByAddressRequest) (explorerPB.GetWasmContractByAddressResponse, error)
	GetCW20Balance(ctx context.Context, req explorerPB.GetCw20BalanceRequest) (explorerPB.GetCw20BalanceResponse, error)
	Close()
}

func NewExplorerClient(network common.Network, options ...common.ClientOption) (ExplorerClient, error) {
	// process options
	opts := common.DefaultClientOptions()
	if network.ChainTlsCert != nil {
		options = append(options, common.OptionTLSCert(network.ExchangeTlsCert))
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
		conn, err = grpc.Dial(network.ExplorerGrpcEndpoint, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {
		conn, err = grpc.Dial(network.ExplorerGrpcEndpoint, grpc.WithInsecure(), grpc.WithContextDialer(common.DialerFunc))
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

	sessionCookie  string
	explorerClient explorerPB.InjectiveExplorerRPCClient
}

func (c *explorerClient) requestCookie() metadata.MD {
	var header metadata.MD
	req := explorerPB.GetTxsRequest{}
	c.explorerClient.GetTxs(context.Background(), &req, grpc.Header(&header))
	return header
}

func (c *explorerClient) getCookie(ctx context.Context) context.Context {
	provider := common.NewMetadataProvider(c.requestCookie)
	cookie, _ := c.network.ExplorerMetadata(provider)
	return metadata.AppendToOutgoingContext(ctx, "cookie", cookie)
}

func (c *explorerClient) setCookie(metadata metadata.MD) {
	md := metadata.Get("set-cookie")
	if len(md) > 0 {
		c.sessionCookie = md[0]
	}
}

func (c *explorerClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *explorerClient) GetTxByTxHash(ctx context.Context, hash string) (explorerPB.GetTxByTxHashResponse, error) {
	req := explorerPB.GetTxByTxHashRequest{
		Hash: hash,
	}

	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetTxByTxHash(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetTxByTxHashResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetAccountTxs(ctx context.Context, req explorerPB.GetAccountTxsRequest) (explorerPB.GetAccountTxsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetAccountTxs(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetAccountTxsResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetBlocks(ctx context.Context) (explorerPB.GetBlocksResponse, error) {
	req := explorerPB.GetBlocksRequest{}

	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetBlocks(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetBlocksResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetBlock(ctx context.Context, blockHeight string) (explorerPB.GetBlockResponse, error) {
	req := explorerPB.GetBlockRequest{
		Id: blockHeight,
	}

	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetBlock(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetBlockResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetTxs(ctx context.Context, req explorerPB.GetTxsRequest) (explorerPB.GetTxsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetTxs(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetTxsResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetPeggyDeposits(ctx context.Context, req explorerPB.GetPeggyDepositTxsRequest) (explorerPB.GetPeggyDepositTxsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetPeggyDepositTxs(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetPeggyDepositTxsResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetPeggyWithdrawals(ctx context.Context, req explorerPB.GetPeggyWithdrawalTxsRequest) (explorerPB.GetPeggyWithdrawalTxsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetPeggyWithdrawalTxs(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetPeggyWithdrawalTxsResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetIBCTransfers(ctx context.Context, req explorerPB.GetIBCTransferTxsRequest) (explorerPB.GetIBCTransferTxsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetIBCTransferTxs(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetIBCTransferTxsResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) StreamTxs(ctx context.Context) (explorerPB.InjectiveExplorerRPC_StreamTxsClient, error) {
	req := explorerPB.StreamTxsRequest{}

	ctx = c.getCookie(ctx)
	stream, err := c.explorerClient.StreamTxs(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *explorerClient) StreamBlocks(ctx context.Context) (explorerPB.InjectiveExplorerRPC_StreamBlocksClient, error) {
	req := explorerPB.StreamBlocksRequest{}

	ctx = c.getCookie(ctx)
	stream, err := c.explorerClient.StreamBlocks(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return stream, nil
}

func (c *explorerClient) GetWasmCodes(ctx context.Context, req explorerPB.GetWasmCodesRequest) (explorerPB.GetWasmCodesResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetWasmCodes(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetWasmCodesResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetWasmCodeByID(ctx context.Context, req explorerPB.GetWasmCodeByIDRequest) (explorerPB.GetWasmCodeByIDResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetWasmCodeByID(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetWasmCodeByIDResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetWasmContracts(ctx context.Context, req explorerPB.GetWasmContractsRequest) (explorerPB.GetWasmContractsResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetWasmContracts(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetWasmContractsResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetWasmContractByAddress(ctx context.Context, req explorerPB.GetWasmContractByAddressRequest) (explorerPB.GetWasmContractByAddressResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetWasmContractByAddress(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetWasmContractByAddressResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) GetCW20Balance(ctx context.Context, req explorerPB.GetCw20BalanceRequest) (explorerPB.GetCw20BalanceResponse, error) {
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetCw20Balance(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetCw20BalanceResponse{}, err
	}

	return *res, nil
}

func (c *explorerClient) Close() {
	c.conn.Close()
}

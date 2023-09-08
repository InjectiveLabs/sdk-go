package exchange

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"time"

	"github.com/InjectiveLabs/sdk-go/client/common"
	explorerPB "github.com/InjectiveLabs/sdk-go/exchange/explorer_rpc/pb"
	"google.golang.org/grpc/metadata"

	log "github.com/InjectiveLabs/suplog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var keepaliveParameters = keepalive.ClientParameters{
	Time:                45 * time.Second, // send pings every 45 seconds if there is no activity
	Timeout:             5 * time.Second,  // wait 5 seconds for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

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

func NewExplorerClient(protoAddr string, options ...common.ClientOption) (ExplorerClient, error) {
	// process options
	opts := common.DefaultClientOptions()
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
		conn, err = grpc.Dial(
			protoAddr,
			grpc.WithTransportCredentials(opts.TLSCert),
			grpc.WithContextDialer(common.DialerFunc),
			grpc.WithKeepaliveParams(keepaliveParameters),
		)
	} else {
		conn, err = grpc.Dial(
			protoAddr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(common.DialerFunc),
			grpc.WithKeepaliveParams(keepaliveParameters),
		)
	}
	if err != nil {
		err := errors.Wrapf(err, "failed to connect to the gRPC: %s", protoAddr)
		return nil, err
	}

	// build client
	cc := &explorerClient{
		opts: opts,
		conn: conn,

		explorerClient: explorerPB.NewInjectiveExplorerRPCClient(conn),
		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "exchangeClient",
		}),
	}

	return cc, nil
}

type explorerClient struct {
	opts   *common.ClientOptions
	conn   *grpc.ClientConn
	logger log.Logger

	sessionCookie  string
	explorerClient explorerPB.InjectiveExplorerRPCClient
}

func (c *explorerClient) setCookie(metadata metadata.MD) {
	md := metadata.Get("set-cookie")
	if len(md) > 0 {
		c.sessionCookie = md[0]
	}
}

func (c *explorerClient) getCookie(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "cookie", c.sessionCookie)
}

func (c *explorerClient) QueryClient() *grpc.ClientConn {
	return c.conn
}

func (c *explorerClient) GetTxByTxHash(ctx context.Context, hash string) (explorerPB.GetTxByTxHashResponse, error) {
	req := explorerPB.GetTxByTxHashRequest{
		Hash: hash,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetTxByTxHash(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetTxByTxHashResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetAccountTxs(ctx context.Context, req explorerPB.GetAccountTxsRequest) (explorerPB.GetAccountTxsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetAccountTxs(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetAccountTxsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetBlocks(ctx context.Context) (explorerPB.GetBlocksResponse, error) {
	req := explorerPB.GetBlocksRequest{}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetBlocks(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetBlocksResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetBlock(ctx context.Context, blockHeight string) (explorerPB.GetBlockResponse, error) {
	req := explorerPB.GetBlockRequest{
		Id: blockHeight,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetBlock(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetBlockResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetTxs(ctx context.Context, req explorerPB.GetTxsRequest) (explorerPB.GetTxsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetTxs(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetTxsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetPeggyDeposits(ctx context.Context, req explorerPB.GetPeggyDepositTxsRequest) (explorerPB.GetPeggyDepositTxsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetPeggyDepositTxs(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetPeggyDepositTxsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetPeggyWithdrawals(ctx context.Context, req explorerPB.GetPeggyWithdrawalTxsRequest) (explorerPB.GetPeggyWithdrawalTxsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetPeggyWithdrawalTxs(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetPeggyWithdrawalTxsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetIBCTransfers(ctx context.Context, req explorerPB.GetIBCTransferTxsRequest) (explorerPB.GetIBCTransferTxsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetIBCTransferTxs(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetIBCTransferTxsResponse{}, err
	}
	c.setCookie(header)

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
	header, err := stream.Header()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

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
	header, err := stream.Header()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	c.setCookie(header)

	return stream, nil
}

func (c *explorerClient) GetWasmCodes(ctx context.Context, req explorerPB.GetWasmCodesRequest) (explorerPB.GetWasmCodesResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetWasmCodes(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetWasmCodesResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetWasmCodeByID(ctx context.Context, req explorerPB.GetWasmCodeByIDRequest) (explorerPB.GetWasmCodeByIDResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetWasmCodeByID(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetWasmCodeByIDResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetWasmContracts(ctx context.Context, req explorerPB.GetWasmContractsRequest) (explorerPB.GetWasmContractsResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetWasmContracts(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetWasmContractsResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetWasmContractByAddress(ctx context.Context, req explorerPB.GetWasmContractByAddressRequest) (explorerPB.GetWasmContractByAddressResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetWasmContractByAddress(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetWasmContractByAddressResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) GetCW20Balance(ctx context.Context, req explorerPB.GetCw20BalanceRequest) (explorerPB.GetCw20BalanceResponse, error) {
	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.explorerClient.GetCw20Balance(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return explorerPB.GetCw20BalanceResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *explorerClient) Close() {
	c.conn.Close()
}

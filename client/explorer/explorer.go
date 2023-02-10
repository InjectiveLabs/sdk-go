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
		conn, err = grpc.Dial(protoAddr, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(common.DialerFunc))
	} else {

		conn, err = grpc.Dial(protoAddr, grpc.WithInsecure(), grpc.WithContextDialer(common.DialerFunc))
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

func (c *explorerClient) Close() {
	c.conn.Close()
}

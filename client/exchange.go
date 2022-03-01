package client

import (
	"context"
	"fmt"
	accountPB "github.com/InjectiveLabs/sdk-go/exchange/accounts_rpc/pb"
	auctionPB "github.com/InjectiveLabs/sdk-go/exchange/auction_rpc/pb"
	derivativeExchangePB "github.com/InjectiveLabs/sdk-go/exchange/derivative_exchange_rpc/pb"
	insurancePB "github.com/InjectiveLabs/sdk-go/exchange/insurance_rpc/pb"
	metaPB "github.com/InjectiveLabs/sdk-go/exchange/meta_rpc/pb"
	oraclePB "github.com/InjectiveLabs/sdk-go/exchange/oracle_rpc/pb"
	spotExchangePB "github.com/InjectiveLabs/sdk-go/exchange/spot_exchange_rpc/pb"
	"google.golang.org/grpc/metadata"

	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
	"google.golang.org/grpc"
)

type ExchangeClient interface {
	GetOrderbook(ctx context.Context, marketId string) (derivativeExchangePB.OrderbookResponse, error)
	StreamOrderbook(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookClient, error)
	Close()
}

// NewExchangeClient creates a new gRPC client that communicates with gRPC server at protoAddr.
// protoAddr must be in form "tcp://127.0.0.1:8080" or "unix:///tmp/test.sock", protocol is required.
func NewExchangeClient(
	protoAddr string,
	options ...clientOption,
) (ExchangeClient, error) {

	// process options
	opts := defaultClientOptions()
	for _, opt := range options {
		if err := opt(opts); err != nil {
			err = errors.Wrap(err, "error in a exchange client option")
			return nil, err
		}
	}

	// create grpc client
	var conn *grpc.ClientConn
	var err error
	if opts.TLSCert != nil {
		conn, err = grpc.Dial(protoAddr, grpc.WithTransportCredentials(opts.TLSCert), grpc.WithContextDialer(dialerFunc))
	} else {
		conn, err = grpc.Dial(protoAddr, grpc.WithInsecure(), grpc.WithContextDialer(dialerFunc))
	}
	if err != nil {
		err := errors.Wrapf(err, "failed to connect to the gRPC: %s", protoAddr)
		return nil, err
	}

	// build client
	cc := &exchangeClient{
		opts: opts,
		conn: conn,

		metaClient:               metaPB.NewInjectiveMetaRPCClient(conn),
		accountClient:            accountPB.NewInjectiveAccountsRPCClient(conn),
		auctionClient:            auctionPB.NewInjectiveAuctionRPCClient(conn),
		oracleClient:             oraclePB.NewInjectiveOracleRPCClient(conn),
		insuranceClient:          insurancePB.NewInjectiveInsuranceRPCClient(conn),
		spotExchangeClient:       spotExchangePB.NewInjectiveSpotExchangeRPCClient(conn),
		derivativeExchangeClient: derivativeExchangePB.NewInjectiveDerivativeExchangeRPCClient(conn),

		logger: log.WithFields(log.Fields{
			"module": "sdk-go",
			"svc":    "cosmosClient",
		}),
	}

	return cc, nil
}

type exchangeClient struct {
	opts   *clientOptions
	conn   *grpc.ClientConn
	logger log.Logger
	client *grpc.ClientConn

	sessionCookie			string

	metaClient               metaPB.InjectiveMetaRPCClient
	accountClient            accountPB.InjectiveAccountsRPCClient
	auctionClient            auctionPB.InjectiveAuctionRPCClient
	oracleClient             oraclePB.InjectiveOracleRPCClient
	insuranceClient          insurancePB.InjectiveInsuranceRPCClient
	spotExchangeClient       spotExchangePB.InjectiveSpotExchangeRPCClient
	derivativeExchangeClient derivativeExchangePB.InjectiveDerivativeExchangeRPCClient

	closed int64
}

func (c *exchangeClient) setCookie(metadata metadata.MD) {
	md := metadata.Get("set-cookie")
	if len(md) > 0 {
		c.sessionCookie = md[0]
	}
}

func (c *exchangeClient) getCookie(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx,"cookie", c.sessionCookie)
}

func (c *exchangeClient) GetOrderbook(ctx context.Context, marketId string) (derivativeExchangePB.OrderbookResponse, error) {
	req := derivativeExchangePB.OrderbookRequest{
		MarketId: marketId,
	}

	var header metadata.MD
	ctx = c.getCookie(ctx)
	res, err := c.derivativeExchangeClient.Orderbook(ctx, &req, grpc.Header(&header))
	if err != nil {
		fmt.Println(err)
		return derivativeExchangePB.OrderbookResponse{}, err
	}
	c.setCookie(header)

	return *res, nil
}

func (c *exchangeClient) StreamOrderbook(ctx context.Context, marketIds []string) (derivativeExchangePB.InjectiveDerivativeExchangeRPC_StreamOrderbookClient, error) {
	req := derivativeExchangePB.StreamOrderbookRequest{
		MarketIds: marketIds,
	}

	ctx = c.getCookie(ctx)
	stream, err := c.derivativeExchangeClient.StreamOrderbook(ctx, &req)
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

func (c *exchangeClient) Close() {
	c.Close()
}

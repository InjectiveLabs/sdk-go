package rpc

import (
	"context"
	"fmt"

	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"google.golang.org/grpc/metadata"
)

// ContextWithHeight wraps a context with the a gRPC block height header. If the provided height is
// 0, it will return an empty context and the gRPC query will use the latest block height for querying.
func ContextWithHeight(height int64) context.Context {
	if height == 0 {
		return context.Background()
	}

	return metadata.AppendToOutgoingContext(context.Background(), grpctypes.GRPCBlockHeightHeader, fmt.Sprintf("%d", height))
}

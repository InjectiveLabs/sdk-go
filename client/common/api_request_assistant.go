package common

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ApiCall[Q any, R any] func(ctx context.Context, in *Q, opts ...grpc.CallOption) (*R, error)
type ApiStreamCall[Q any, S grpc.ClientStream] func(ctx context.Context, in *Q, opts ...grpc.CallOption) (S, error)

func ExecuteCall[Q any, R any](ctx context.Context, cookieAssistant CookieAssistant, call ApiCall[Q, R], in *Q) (*R, error) {
	var header metadata.MD
	localCtx := metadata.NewOutgoingContext(ctx, cookieAssistant.RealMetadata())

	response, err := call(localCtx, in, grpc.Header(&header))

	cookieAssistant.ProcessResponseMetadata(header)

	return response, err
}

func ExecuteStreamCall[Q any, S grpc.ClientStream](ctx context.Context, cookieAssistant CookieAssistant, call ApiStreamCall[Q, S], in *Q) (S, error) {
	localCtx := metadata.NewOutgoingContext(ctx, cookieAssistant.RealMetadata())

	stream, callError := call(localCtx, in)

	if callError == nil {
		header, err := stream.Header()
		if err == nil {
			cookieAssistant.ProcessResponseMetadata(header)
		}
	}

	return stream, callError
}

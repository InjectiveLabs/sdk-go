package common

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type APICall[Q any, R any] func(ctx context.Context, in *Q, opts ...grpc.CallOption) (*R, error)
type APIStreamCall[Q any, S grpc.ClientStream] func(ctx context.Context, in *Q, opts ...grpc.CallOption) (S, error)

func ExecuteCall[Q any, R any](
	ctx context.Context,
	cookieAssistant CookieAssistant,
	call APICall[Q, R],
	in *Q,
	opts ...grpc.CallOption,
) (*R, error) {
	md := cookieAssistant.RealMetadata()

	if upstreamMetadata, ok := metadata.FromOutgoingContext(ctx); ok {
		// If metadata already exists in the context, merge it with the cookie metadata
		md = metadata.Join(md, upstreamMetadata)
	}

	localCtx := metadata.NewOutgoingContext(ctx, md)

	allOpts := make([]grpc.CallOption, 0, len(opts)+1)
	allOpts = append(allOpts, opts...)
	allOpts = append(allOpts, grpc.Header(&md))
	response, err := call(localCtx, in, allOpts...)

	cookieAssistant.ProcessResponseMetadata(md)

	return response, err
}

func ExecuteStreamCall[Q any, S grpc.ClientStream](
	ctx context.Context,
	cookieAssistant CookieAssistant,
	call APIStreamCall[Q, S],
	in *Q,
	opts ...grpc.CallOption,
) (S, error) {
	md := cookieAssistant.RealMetadata()

	if upstreamMetadata, ok := metadata.FromOutgoingContext(ctx); ok {
		// If metadata already exists in the context, merge it with the cookie metadata
		md = metadata.Join(md, upstreamMetadata)
	}

	localCtx := metadata.NewOutgoingContext(ctx, md)

	allOpts := make([]grpc.CallOption, 0, len(opts)+1)
	allOpts = append(allOpts, opts...)
	allOpts = append(allOpts, grpc.Header(&md))
	stream, callError := call(localCtx, in, allOpts...)

	if callError == nil {
		header, err := stream.Header()
		if err == nil {
			cookieAssistant.ProcessResponseMetadata(header)
		}
	}

	return stream, callError
}

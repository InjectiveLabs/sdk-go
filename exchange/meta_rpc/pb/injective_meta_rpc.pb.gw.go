// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: injective_meta_rpc.proto

/*
Package injective_meta_rpcpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package injective_meta_rpcpb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_InjectiveMetaRPC_Ping_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveMetaRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PingRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Ping(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InjectiveMetaRPC_Ping_0(ctx context.Context, marshaler runtime.Marshaler, server InjectiveMetaRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PingRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Ping(ctx, &protoReq)
	return msg, metadata, err

}

func request_InjectiveMetaRPC_Version_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveMetaRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VersionRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Version(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InjectiveMetaRPC_Version_0(ctx context.Context, marshaler runtime.Marshaler, server InjectiveMetaRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VersionRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Version(ctx, &protoReq)
	return msg, metadata, err

}

func request_InjectiveMetaRPC_Info_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveMetaRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq InfoRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Info(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InjectiveMetaRPC_Info_0(ctx context.Context, marshaler runtime.Marshaler, server InjectiveMetaRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq InfoRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Info(ctx, &protoReq)
	return msg, metadata, err

}

func request_InjectiveMetaRPC_StreamKeepalive_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveMetaRPCClient, req *http.Request, pathParams map[string]string) (InjectiveMetaRPC_StreamKeepaliveClient, runtime.ServerMetadata, error) {
	var protoReq StreamKeepaliveRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.StreamKeepalive(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

func request_InjectiveMetaRPC_TokenMetadata_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveMetaRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TokenMetadataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.TokenMetadata(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InjectiveMetaRPC_TokenMetadata_0(ctx context.Context, marshaler runtime.Marshaler, server InjectiveMetaRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TokenMetadataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.TokenMetadata(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterInjectiveMetaRPCHandlerServer registers the http handlers for service InjectiveMetaRPC to "mux".
// UnaryRPC     :call InjectiveMetaRPCServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterInjectiveMetaRPCHandlerFromEndpoint instead.
func RegisterInjectiveMetaRPCHandlerServer(ctx context.Context, mux *runtime.ServeMux, server InjectiveMetaRPCServer) error {

	mux.Handle("POST", pattern_InjectiveMetaRPC_Ping_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/injective_meta_rpc.InjectiveMetaRPC/Ping", runtime.WithHTTPPathPattern("/injective_meta_rpc.InjectiveMetaRPC/Ping"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InjectiveMetaRPC_Ping_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveMetaRPC_Ping_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveMetaRPC_Version_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/injective_meta_rpc.InjectiveMetaRPC/Version", runtime.WithHTTPPathPattern("/injective_meta_rpc.InjectiveMetaRPC/Version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InjectiveMetaRPC_Version_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveMetaRPC_Version_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveMetaRPC_Info_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/injective_meta_rpc.InjectiveMetaRPC/Info", runtime.WithHTTPPathPattern("/injective_meta_rpc.InjectiveMetaRPC/Info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InjectiveMetaRPC_Info_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveMetaRPC_Info_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveMetaRPC_StreamKeepalive_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("POST", pattern_InjectiveMetaRPC_TokenMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/injective_meta_rpc.InjectiveMetaRPC/TokenMetadata", runtime.WithHTTPPathPattern("/injective_meta_rpc.InjectiveMetaRPC/TokenMetadata"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InjectiveMetaRPC_TokenMetadata_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveMetaRPC_TokenMetadata_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterInjectiveMetaRPCHandlerFromEndpoint is same as RegisterInjectiveMetaRPCHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterInjectiveMetaRPCHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterInjectiveMetaRPCHandler(ctx, mux, conn)
}

// RegisterInjectiveMetaRPCHandler registers the http handlers for service InjectiveMetaRPC to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterInjectiveMetaRPCHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterInjectiveMetaRPCHandlerClient(ctx, mux, NewInjectiveMetaRPCClient(conn))
}

// RegisterInjectiveMetaRPCHandlerClient registers the http handlers for service InjectiveMetaRPC
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "InjectiveMetaRPCClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "InjectiveMetaRPCClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "InjectiveMetaRPCClient" to call the correct interceptors.
func RegisterInjectiveMetaRPCHandlerClient(ctx context.Context, mux *runtime.ServeMux, client InjectiveMetaRPCClient) error {

	mux.Handle("POST", pattern_InjectiveMetaRPC_Ping_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_meta_rpc.InjectiveMetaRPC/Ping", runtime.WithHTTPPathPattern("/injective_meta_rpc.InjectiveMetaRPC/Ping"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveMetaRPC_Ping_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveMetaRPC_Ping_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveMetaRPC_Version_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_meta_rpc.InjectiveMetaRPC/Version", runtime.WithHTTPPathPattern("/injective_meta_rpc.InjectiveMetaRPC/Version"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveMetaRPC_Version_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveMetaRPC_Version_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveMetaRPC_Info_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_meta_rpc.InjectiveMetaRPC/Info", runtime.WithHTTPPathPattern("/injective_meta_rpc.InjectiveMetaRPC/Info"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveMetaRPC_Info_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveMetaRPC_Info_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveMetaRPC_StreamKeepalive_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_meta_rpc.InjectiveMetaRPC/StreamKeepalive", runtime.WithHTTPPathPattern("/injective_meta_rpc.InjectiveMetaRPC/StreamKeepalive"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveMetaRPC_StreamKeepalive_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveMetaRPC_StreamKeepalive_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveMetaRPC_TokenMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_meta_rpc.InjectiveMetaRPC/TokenMetadata", runtime.WithHTTPPathPattern("/injective_meta_rpc.InjectiveMetaRPC/TokenMetadata"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveMetaRPC_TokenMetadata_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveMetaRPC_TokenMetadata_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_InjectiveMetaRPC_Ping_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_meta_rpc.InjectiveMetaRPC", "Ping"}, ""))

	pattern_InjectiveMetaRPC_Version_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_meta_rpc.InjectiveMetaRPC", "Version"}, ""))

	pattern_InjectiveMetaRPC_Info_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_meta_rpc.InjectiveMetaRPC", "Info"}, ""))

	pattern_InjectiveMetaRPC_StreamKeepalive_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_meta_rpc.InjectiveMetaRPC", "StreamKeepalive"}, ""))

	pattern_InjectiveMetaRPC_TokenMetadata_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_meta_rpc.InjectiveMetaRPC", "TokenMetadata"}, ""))
)

var (
	forward_InjectiveMetaRPC_Ping_0 = runtime.ForwardResponseMessage

	forward_InjectiveMetaRPC_Version_0 = runtime.ForwardResponseMessage

	forward_InjectiveMetaRPC_Info_0 = runtime.ForwardResponseMessage

	forward_InjectiveMetaRPC_StreamKeepalive_0 = runtime.ForwardResponseStream

	forward_InjectiveMetaRPC_TokenMetadata_0 = runtime.ForwardResponseMessage
)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: injective_meta_rpc.proto

package injective_meta_rpcpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InjectiveMetaRPCClient is the client API for InjectiveMetaRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InjectiveMetaRPCClient interface {
	// Endpoint for checking server health.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Returns injective-exchange version.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Gets connection info
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	// Stream keepalive, if server exits, a shutdown event will be sent over this
	// channel.
	StreamKeepalive(ctx context.Context, in *StreamKeepaliveRequest, opts ...grpc.CallOption) (InjectiveMetaRPC_StreamKeepaliveClient, error)
	// Get tokens metadata. Can be filtered by denom
	TokenMetadata(ctx context.Context, in *TokenMetadataRequest, opts ...grpc.CallOption) (*TokenMetadataResponse, error)
}

type injectiveMetaRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewInjectiveMetaRPCClient(cc grpc.ClientConnInterface) InjectiveMetaRPCClient {
	return &injectiveMetaRPCClient{cc}
}

func (c *injectiveMetaRPCClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/injective_meta_rpc.InjectiveMetaRPC/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveMetaRPCClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/injective_meta_rpc.InjectiveMetaRPC/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveMetaRPCClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/injective_meta_rpc.InjectiveMetaRPC/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveMetaRPCClient) StreamKeepalive(ctx context.Context, in *StreamKeepaliveRequest, opts ...grpc.CallOption) (InjectiveMetaRPC_StreamKeepaliveClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveMetaRPC_ServiceDesc.Streams[0], "/injective_meta_rpc.InjectiveMetaRPC/StreamKeepalive", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectiveMetaRPCStreamKeepaliveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveMetaRPC_StreamKeepaliveClient interface {
	Recv() (*StreamKeepaliveResponse, error)
	grpc.ClientStream
}

type injectiveMetaRPCStreamKeepaliveClient struct {
	grpc.ClientStream
}

func (x *injectiveMetaRPCStreamKeepaliveClient) Recv() (*StreamKeepaliveResponse, error) {
	m := new(StreamKeepaliveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectiveMetaRPCClient) TokenMetadata(ctx context.Context, in *TokenMetadataRequest, opts ...grpc.CallOption) (*TokenMetadataResponse, error) {
	out := new(TokenMetadataResponse)
	err := c.cc.Invoke(ctx, "/injective_meta_rpc.InjectiveMetaRPC/TokenMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InjectiveMetaRPCServer is the server API for InjectiveMetaRPC service.
// All implementations must embed UnimplementedInjectiveMetaRPCServer
// for forward compatibility
type InjectiveMetaRPCServer interface {
	// Endpoint for checking server health.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Returns injective-exchange version.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Gets connection info
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	// Stream keepalive, if server exits, a shutdown event will be sent over this
	// channel.
	StreamKeepalive(*StreamKeepaliveRequest, InjectiveMetaRPC_StreamKeepaliveServer) error
	// Get tokens metadata. Can be filtered by denom
	TokenMetadata(context.Context, *TokenMetadataRequest) (*TokenMetadataResponse, error)
	mustEmbedUnimplementedInjectiveMetaRPCServer()
}

// UnimplementedInjectiveMetaRPCServer must be embedded to have forward compatible implementations.
type UnimplementedInjectiveMetaRPCServer struct {
}

func (UnimplementedInjectiveMetaRPCServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedInjectiveMetaRPCServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedInjectiveMetaRPCServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedInjectiveMetaRPCServer) StreamKeepalive(*StreamKeepaliveRequest, InjectiveMetaRPC_StreamKeepaliveServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamKeepalive not implemented")
}
func (UnimplementedInjectiveMetaRPCServer) TokenMetadata(context.Context, *TokenMetadataRequest) (*TokenMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenMetadata not implemented")
}
func (UnimplementedInjectiveMetaRPCServer) mustEmbedUnimplementedInjectiveMetaRPCServer() {}

// UnsafeInjectiveMetaRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InjectiveMetaRPCServer will
// result in compilation errors.
type UnsafeInjectiveMetaRPCServer interface {
	mustEmbedUnimplementedInjectiveMetaRPCServer()
}

func RegisterInjectiveMetaRPCServer(s grpc.ServiceRegistrar, srv InjectiveMetaRPCServer) {
	s.RegisterService(&InjectiveMetaRPC_ServiceDesc, srv)
}

func _InjectiveMetaRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveMetaRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_meta_rpc.InjectiveMetaRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveMetaRPCServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveMetaRPC_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveMetaRPCServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_meta_rpc.InjectiveMetaRPC/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveMetaRPCServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveMetaRPC_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveMetaRPCServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_meta_rpc.InjectiveMetaRPC/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveMetaRPCServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveMetaRPC_StreamKeepalive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamKeepaliveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveMetaRPCServer).StreamKeepalive(m, &injectiveMetaRPCStreamKeepaliveServer{stream})
}

type InjectiveMetaRPC_StreamKeepaliveServer interface {
	Send(*StreamKeepaliveResponse) error
	grpc.ServerStream
}

type injectiveMetaRPCStreamKeepaliveServer struct {
	grpc.ServerStream
}

func (x *injectiveMetaRPCStreamKeepaliveServer) Send(m *StreamKeepaliveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveMetaRPC_TokenMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveMetaRPCServer).TokenMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_meta_rpc.InjectiveMetaRPC/TokenMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveMetaRPCServer).TokenMetadata(ctx, req.(*TokenMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InjectiveMetaRPC_ServiceDesc is the grpc.ServiceDesc for InjectiveMetaRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InjectiveMetaRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "injective_meta_rpc.InjectiveMetaRPC",
	HandlerType: (*InjectiveMetaRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _InjectiveMetaRPC_Ping_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _InjectiveMetaRPC_Version_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _InjectiveMetaRPC_Info_Handler,
		},
		{
			MethodName: "TokenMetadata",
			Handler:    _InjectiveMetaRPC_TokenMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamKeepalive",
			Handler:       _InjectiveMetaRPC_StreamKeepalive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "injective_meta_rpc.proto",
}

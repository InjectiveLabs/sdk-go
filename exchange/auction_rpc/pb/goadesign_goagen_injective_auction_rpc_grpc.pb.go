// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: goadesign_goagen_injective_auction_rpc.proto

package injective_auction_rpcpb

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

// InjectiveAuctionRPCClient is the client API for InjectiveAuctionRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InjectiveAuctionRPCClient interface {
	// Provide historical auction info for a given auction
	AuctionEndpoint(ctx context.Context, in *AuctionEndpointRequest, opts ...grpc.CallOption) (*AuctionEndpointResponse, error)
	// Provide the historical auctions info
	Auctions(ctx context.Context, in *AuctionsRequest, opts ...grpc.CallOption) (*AuctionsResponse, error)
	// StreamBids streams new bids of an auction.
	StreamBids(ctx context.Context, in *StreamBidsRequest, opts ...grpc.CallOption) (InjectiveAuctionRPC_StreamBidsClient, error)
	// InjBurntEndpoint returns the total amount of INJ burnt in auctions.
	InjBurntEndpoint(ctx context.Context, in *InjBurntEndpointRequest, opts ...grpc.CallOption) (*InjBurntEndpointResponse, error)
}

type injectiveAuctionRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewInjectiveAuctionRPCClient(cc grpc.ClientConnInterface) InjectiveAuctionRPCClient {
	return &injectiveAuctionRPCClient{cc}
}

func (c *injectiveAuctionRPCClient) AuctionEndpoint(ctx context.Context, in *AuctionEndpointRequest, opts ...grpc.CallOption) (*AuctionEndpointResponse, error) {
	out := new(AuctionEndpointResponse)
	err := c.cc.Invoke(ctx, "/injective_auction_rpc.InjectiveAuctionRPC/AuctionEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveAuctionRPCClient) Auctions(ctx context.Context, in *AuctionsRequest, opts ...grpc.CallOption) (*AuctionsResponse, error) {
	out := new(AuctionsResponse)
	err := c.cc.Invoke(ctx, "/injective_auction_rpc.InjectiveAuctionRPC/Auctions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *injectiveAuctionRPCClient) StreamBids(ctx context.Context, in *StreamBidsRequest, opts ...grpc.CallOption) (InjectiveAuctionRPC_StreamBidsClient, error) {
	stream, err := c.cc.NewStream(ctx, &InjectiveAuctionRPC_ServiceDesc.Streams[0], "/injective_auction_rpc.InjectiveAuctionRPC/StreamBids", opts...)
	if err != nil {
		return nil, err
	}
	x := &injectiveAuctionRPCStreamBidsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InjectiveAuctionRPC_StreamBidsClient interface {
	Recv() (*StreamBidsResponse, error)
	grpc.ClientStream
}

type injectiveAuctionRPCStreamBidsClient struct {
	grpc.ClientStream
}

func (x *injectiveAuctionRPCStreamBidsClient) Recv() (*StreamBidsResponse, error) {
	m := new(StreamBidsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *injectiveAuctionRPCClient) InjBurntEndpoint(ctx context.Context, in *InjBurntEndpointRequest, opts ...grpc.CallOption) (*InjBurntEndpointResponse, error) {
	out := new(InjBurntEndpointResponse)
	err := c.cc.Invoke(ctx, "/injective_auction_rpc.InjectiveAuctionRPC/InjBurntEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InjectiveAuctionRPCServer is the server API for InjectiveAuctionRPC service.
// All implementations must embed UnimplementedInjectiveAuctionRPCServer
// for forward compatibility
type InjectiveAuctionRPCServer interface {
	// Provide historical auction info for a given auction
	AuctionEndpoint(context.Context, *AuctionEndpointRequest) (*AuctionEndpointResponse, error)
	// Provide the historical auctions info
	Auctions(context.Context, *AuctionsRequest) (*AuctionsResponse, error)
	// StreamBids streams new bids of an auction.
	StreamBids(*StreamBidsRequest, InjectiveAuctionRPC_StreamBidsServer) error
	// InjBurntEndpoint returns the total amount of INJ burnt in auctions.
	InjBurntEndpoint(context.Context, *InjBurntEndpointRequest) (*InjBurntEndpointResponse, error)
	mustEmbedUnimplementedInjectiveAuctionRPCServer()
}

// UnimplementedInjectiveAuctionRPCServer must be embedded to have forward compatible implementations.
type UnimplementedInjectiveAuctionRPCServer struct {
}

func (UnimplementedInjectiveAuctionRPCServer) AuctionEndpoint(context.Context, *AuctionEndpointRequest) (*AuctionEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuctionEndpoint not implemented")
}
func (UnimplementedInjectiveAuctionRPCServer) Auctions(context.Context, *AuctionsRequest) (*AuctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auctions not implemented")
}
func (UnimplementedInjectiveAuctionRPCServer) StreamBids(*StreamBidsRequest, InjectiveAuctionRPC_StreamBidsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamBids not implemented")
}
func (UnimplementedInjectiveAuctionRPCServer) InjBurntEndpoint(context.Context, *InjBurntEndpointRequest) (*InjBurntEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InjBurntEndpoint not implemented")
}
func (UnimplementedInjectiveAuctionRPCServer) mustEmbedUnimplementedInjectiveAuctionRPCServer() {}

// UnsafeInjectiveAuctionRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InjectiveAuctionRPCServer will
// result in compilation errors.
type UnsafeInjectiveAuctionRPCServer interface {
	mustEmbedUnimplementedInjectiveAuctionRPCServer()
}

func RegisterInjectiveAuctionRPCServer(s grpc.ServiceRegistrar, srv InjectiveAuctionRPCServer) {
	s.RegisterService(&InjectiveAuctionRPC_ServiceDesc, srv)
}

func _InjectiveAuctionRPC_AuctionEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuctionEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAuctionRPCServer).AuctionEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_auction_rpc.InjectiveAuctionRPC/AuctionEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAuctionRPCServer).AuctionEndpoint(ctx, req.(*AuctionEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveAuctionRPC_Auctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAuctionRPCServer).Auctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_auction_rpc.InjectiveAuctionRPC/Auctions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAuctionRPCServer).Auctions(ctx, req.(*AuctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InjectiveAuctionRPC_StreamBids_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamBidsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InjectiveAuctionRPCServer).StreamBids(m, &injectiveAuctionRPCStreamBidsServer{stream})
}

type InjectiveAuctionRPC_StreamBidsServer interface {
	Send(*StreamBidsResponse) error
	grpc.ServerStream
}

type injectiveAuctionRPCStreamBidsServer struct {
	grpc.ServerStream
}

func (x *injectiveAuctionRPCStreamBidsServer) Send(m *StreamBidsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InjectiveAuctionRPC_InjBurntEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InjBurntEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InjectiveAuctionRPCServer).InjBurntEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/injective_auction_rpc.InjectiveAuctionRPC/InjBurntEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InjectiveAuctionRPCServer).InjBurntEndpoint(ctx, req.(*InjBurntEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InjectiveAuctionRPC_ServiceDesc is the grpc.ServiceDesc for InjectiveAuctionRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InjectiveAuctionRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "injective_auction_rpc.InjectiveAuctionRPC",
	HandlerType: (*InjectiveAuctionRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuctionEndpoint",
			Handler:    _InjectiveAuctionRPC_AuctionEndpoint_Handler,
		},
		{
			MethodName: "Auctions",
			Handler:    _InjectiveAuctionRPC_Auctions_Handler,
		},
		{
			MethodName: "InjBurntEndpoint",
			Handler:    _InjectiveAuctionRPC_InjBurntEndpoint_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamBids",
			Handler:       _InjectiveAuctionRPC_StreamBids_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "goadesign_goagen_injective_auction_rpc.proto",
}

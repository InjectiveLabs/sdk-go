// Code generated by MockGen. DO NOT EDIT.
// Source: api/gen/grpc/injective_trading_rpc/pb/injective_trading_rpc_grpc.pb.go

// Package injective_trading_rpcpb is a generated GoMock package.
package injective_trading_rpcpb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockInjectiveTradingRPCClient is a mock of InjectiveTradingRPCClient interface.
type MockInjectiveTradingRPCClient struct {
	ctrl     *gomock.Controller
	recorder *MockInjectiveTradingRPCClientMockRecorder
}

// MockInjectiveTradingRPCClientMockRecorder is the mock recorder for MockInjectiveTradingRPCClient.
type MockInjectiveTradingRPCClientMockRecorder struct {
	mock *MockInjectiveTradingRPCClient
}

// NewMockInjectiveTradingRPCClient creates a new mock instance.
func NewMockInjectiveTradingRPCClient(ctrl *gomock.Controller) *MockInjectiveTradingRPCClient {
	mock := &MockInjectiveTradingRPCClient{ctrl: ctrl}
	mock.recorder = &MockInjectiveTradingRPCClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInjectiveTradingRPCClient) EXPECT() *MockInjectiveTradingRPCClientMockRecorder {
	return m.recorder
}

// ListTradingStrategies mocks base method.
func (m *MockInjectiveTradingRPCClient) ListTradingStrategies(ctx context.Context, in *ListTradingStrategiesRequest, opts ...grpc.CallOption) (*ListTradingStrategiesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTradingStrategies", varargs...)
	ret0, _ := ret[0].(*ListTradingStrategiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTradingStrategies indicates an expected call of ListTradingStrategies.
func (mr *MockInjectiveTradingRPCClientMockRecorder) ListTradingStrategies(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTradingStrategies", reflect.TypeOf((*MockInjectiveTradingRPCClient)(nil).ListTradingStrategies), varargs...)
}

// MockInjectiveTradingRPCServer is a mock of InjectiveTradingRPCServer interface.
type MockInjectiveTradingRPCServer struct {
	ctrl     *gomock.Controller
	recorder *MockInjectiveTradingRPCServerMockRecorder
}

// MockInjectiveTradingRPCServerMockRecorder is the mock recorder for MockInjectiveTradingRPCServer.
type MockInjectiveTradingRPCServerMockRecorder struct {
	mock *MockInjectiveTradingRPCServer
}

// NewMockInjectiveTradingRPCServer creates a new mock instance.
func NewMockInjectiveTradingRPCServer(ctrl *gomock.Controller) *MockInjectiveTradingRPCServer {
	mock := &MockInjectiveTradingRPCServer{ctrl: ctrl}
	mock.recorder = &MockInjectiveTradingRPCServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInjectiveTradingRPCServer) EXPECT() *MockInjectiveTradingRPCServerMockRecorder {
	return m.recorder
}

// ListTradingStrategies mocks base method.
func (m *MockInjectiveTradingRPCServer) ListTradingStrategies(arg0 context.Context, arg1 *ListTradingStrategiesRequest) (*ListTradingStrategiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTradingStrategies", arg0, arg1)
	ret0, _ := ret[0].(*ListTradingStrategiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTradingStrategies indicates an expected call of ListTradingStrategies.
func (mr *MockInjectiveTradingRPCServerMockRecorder) ListTradingStrategies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTradingStrategies", reflect.TypeOf((*MockInjectiveTradingRPCServer)(nil).ListTradingStrategies), arg0, arg1)
}

// mustEmbedUnimplementedInjectiveTradingRPCServer mocks base method.
func (m *MockInjectiveTradingRPCServer) mustEmbedUnimplementedInjectiveTradingRPCServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInjectiveTradingRPCServer")
}

// mustEmbedUnimplementedInjectiveTradingRPCServer indicates an expected call of mustEmbedUnimplementedInjectiveTradingRPCServer.
func (mr *MockInjectiveTradingRPCServerMockRecorder) mustEmbedUnimplementedInjectiveTradingRPCServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInjectiveTradingRPCServer", reflect.TypeOf((*MockInjectiveTradingRPCServer)(nil).mustEmbedUnimplementedInjectiveTradingRPCServer))
}

// MockUnsafeInjectiveTradingRPCServer is a mock of UnsafeInjectiveTradingRPCServer interface.
type MockUnsafeInjectiveTradingRPCServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeInjectiveTradingRPCServerMockRecorder
}

// MockUnsafeInjectiveTradingRPCServerMockRecorder is the mock recorder for MockUnsafeInjectiveTradingRPCServer.
type MockUnsafeInjectiveTradingRPCServerMockRecorder struct {
	mock *MockUnsafeInjectiveTradingRPCServer
}

// NewMockUnsafeInjectiveTradingRPCServer creates a new mock instance.
func NewMockUnsafeInjectiveTradingRPCServer(ctrl *gomock.Controller) *MockUnsafeInjectiveTradingRPCServer {
	mock := &MockUnsafeInjectiveTradingRPCServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeInjectiveTradingRPCServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeInjectiveTradingRPCServer) EXPECT() *MockUnsafeInjectiveTradingRPCServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedInjectiveTradingRPCServer mocks base method.
func (m *MockUnsafeInjectiveTradingRPCServer) mustEmbedUnimplementedInjectiveTradingRPCServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInjectiveTradingRPCServer")
}

// mustEmbedUnimplementedInjectiveTradingRPCServer indicates an expected call of mustEmbedUnimplementedInjectiveTradingRPCServer.
func (mr *MockUnsafeInjectiveTradingRPCServerMockRecorder) mustEmbedUnimplementedInjectiveTradingRPCServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInjectiveTradingRPCServer", reflect.TypeOf((*MockUnsafeInjectiveTradingRPCServer)(nil).mustEmbedUnimplementedInjectiveTradingRPCServer))
}

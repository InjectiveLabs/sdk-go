// Code generated by MockGen. DO NOT EDIT.
// Source: api/gen/grpc/event_provider_api/pb/goadesign_goagen_event_provider_api_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=api/gen/grpc/event_provider_api/pb/goadesign_goagen_event_provider_api_grpc.pb.go -destination=api/gen/grpc/event_provider_api/pb/event_provider_api_grpc.pb_mock.go -package=event_provider_apipb
//

// Package event_provider_apipb is a generated GoMock package.
package event_provider_apipb

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockEventProviderAPIClient is a mock of EventProviderAPIClient interface.
type MockEventProviderAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockEventProviderAPIClientMockRecorder
}

// MockEventProviderAPIClientMockRecorder is the mock recorder for MockEventProviderAPIClient.
type MockEventProviderAPIClientMockRecorder struct {
	mock *MockEventProviderAPIClient
}

// NewMockEventProviderAPIClient creates a new mock instance.
func NewMockEventProviderAPIClient(ctrl *gomock.Controller) *MockEventProviderAPIClient {
	mock := &MockEventProviderAPIClient{ctrl: ctrl}
	mock.recorder = &MockEventProviderAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventProviderAPIClient) EXPECT() *MockEventProviderAPIClientMockRecorder {
	return m.recorder
}

// GetABCIBlockEvents mocks base method.
func (m *MockEventProviderAPIClient) GetABCIBlockEvents(ctx context.Context, in *GetABCIBlockEventsRequest, opts ...grpc.CallOption) (*GetABCIBlockEventsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetABCIBlockEvents", varargs...)
	ret0, _ := ret[0].(*GetABCIBlockEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetABCIBlockEvents indicates an expected call of GetABCIBlockEvents.
func (mr *MockEventProviderAPIClientMockRecorder) GetABCIBlockEvents(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetABCIBlockEvents", reflect.TypeOf((*MockEventProviderAPIClient)(nil).GetABCIBlockEvents), varargs...)
}

// GetABCIBlockEventsAtHeight mocks base method.
func (m *MockEventProviderAPIClient) GetABCIBlockEventsAtHeight(ctx context.Context, in *GetABCIBlockEventsAtHeightRequest, opts ...grpc.CallOption) (*GetABCIBlockEventsAtHeightResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetABCIBlockEventsAtHeight", varargs...)
	ret0, _ := ret[0].(*GetABCIBlockEventsAtHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetABCIBlockEventsAtHeight indicates an expected call of GetABCIBlockEventsAtHeight.
func (mr *MockEventProviderAPIClientMockRecorder) GetABCIBlockEventsAtHeight(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetABCIBlockEventsAtHeight", reflect.TypeOf((*MockEventProviderAPIClient)(nil).GetABCIBlockEventsAtHeight), varargs...)
}

// GetBlockEventsRPC mocks base method.
func (m *MockEventProviderAPIClient) GetBlockEventsRPC(ctx context.Context, in *GetBlockEventsRPCRequest, opts ...grpc.CallOption) (*GetBlockEventsRPCResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlockEventsRPC", varargs...)
	ret0, _ := ret[0].(*GetBlockEventsRPCResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockEventsRPC indicates an expected call of GetBlockEventsRPC.
func (mr *MockEventProviderAPIClientMockRecorder) GetBlockEventsRPC(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockEventsRPC", reflect.TypeOf((*MockEventProviderAPIClient)(nil).GetBlockEventsRPC), varargs...)
}

// GetCustomEventsRPC mocks base method.
func (m *MockEventProviderAPIClient) GetCustomEventsRPC(ctx context.Context, in *GetCustomEventsRPCRequest, opts ...grpc.CallOption) (*GetCustomEventsRPCResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCustomEventsRPC", varargs...)
	ret0, _ := ret[0].(*GetCustomEventsRPCResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomEventsRPC indicates an expected call of GetCustomEventsRPC.
func (mr *MockEventProviderAPIClientMockRecorder) GetCustomEventsRPC(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomEventsRPC", reflect.TypeOf((*MockEventProviderAPIClient)(nil).GetCustomEventsRPC), varargs...)
}

// GetLatestHeight mocks base method.
func (m *MockEventProviderAPIClient) GetLatestHeight(ctx context.Context, in *GetLatestHeightRequest, opts ...grpc.CallOption) (*GetLatestHeightResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLatestHeight", varargs...)
	ret0, _ := ret[0].(*GetLatestHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestHeight indicates an expected call of GetLatestHeight.
func (mr *MockEventProviderAPIClientMockRecorder) GetLatestHeight(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestHeight", reflect.TypeOf((*MockEventProviderAPIClient)(nil).GetLatestHeight), varargs...)
}

// StreamBlockEvents mocks base method.
func (m *MockEventProviderAPIClient) StreamBlockEvents(ctx context.Context, in *StreamBlockEventsRequest, opts ...grpc.CallOption) (EventProviderAPI_StreamBlockEventsClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StreamBlockEvents", varargs...)
	ret0, _ := ret[0].(EventProviderAPI_StreamBlockEventsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StreamBlockEvents indicates an expected call of StreamBlockEvents.
func (mr *MockEventProviderAPIClientMockRecorder) StreamBlockEvents(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamBlockEvents", reflect.TypeOf((*MockEventProviderAPIClient)(nil).StreamBlockEvents), varargs...)
}

// MockEventProviderAPI_StreamBlockEventsClient is a mock of EventProviderAPI_StreamBlockEventsClient interface.
type MockEventProviderAPI_StreamBlockEventsClient struct {
	ctrl     *gomock.Controller
	recorder *MockEventProviderAPI_StreamBlockEventsClientMockRecorder
}

// MockEventProviderAPI_StreamBlockEventsClientMockRecorder is the mock recorder for MockEventProviderAPI_StreamBlockEventsClient.
type MockEventProviderAPI_StreamBlockEventsClientMockRecorder struct {
	mock *MockEventProviderAPI_StreamBlockEventsClient
}

// NewMockEventProviderAPI_StreamBlockEventsClient creates a new mock instance.
func NewMockEventProviderAPI_StreamBlockEventsClient(ctrl *gomock.Controller) *MockEventProviderAPI_StreamBlockEventsClient {
	mock := &MockEventProviderAPI_StreamBlockEventsClient{ctrl: ctrl}
	mock.recorder = &MockEventProviderAPI_StreamBlockEventsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventProviderAPI_StreamBlockEventsClient) EXPECT() *MockEventProviderAPI_StreamBlockEventsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockEventProviderAPI_StreamBlockEventsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockEventProviderAPI_StreamBlockEventsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockEventProviderAPI_StreamBlockEventsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsClient) Recv() (*StreamBlockEventsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*StreamBlockEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockEventProviderAPI_StreamBlockEventsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockEventProviderAPI_StreamBlockEventsClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockEventProviderAPI_StreamBlockEventsClientMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockEventProviderAPI_StreamBlockEventsClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockEventProviderAPI_StreamBlockEventsClientMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockEventProviderAPI_StreamBlockEventsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsClient)(nil).Trailer))
}

// MockEventProviderAPIServer is a mock of EventProviderAPIServer interface.
type MockEventProviderAPIServer struct {
	ctrl     *gomock.Controller
	recorder *MockEventProviderAPIServerMockRecorder
}

// MockEventProviderAPIServerMockRecorder is the mock recorder for MockEventProviderAPIServer.
type MockEventProviderAPIServerMockRecorder struct {
	mock *MockEventProviderAPIServer
}

// NewMockEventProviderAPIServer creates a new mock instance.
func NewMockEventProviderAPIServer(ctrl *gomock.Controller) *MockEventProviderAPIServer {
	mock := &MockEventProviderAPIServer{ctrl: ctrl}
	mock.recorder = &MockEventProviderAPIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventProviderAPIServer) EXPECT() *MockEventProviderAPIServerMockRecorder {
	return m.recorder
}

// GetABCIBlockEvents mocks base method.
func (m *MockEventProviderAPIServer) GetABCIBlockEvents(arg0 context.Context, arg1 *GetABCIBlockEventsRequest) (*GetABCIBlockEventsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetABCIBlockEvents", arg0, arg1)
	ret0, _ := ret[0].(*GetABCIBlockEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetABCIBlockEvents indicates an expected call of GetABCIBlockEvents.
func (mr *MockEventProviderAPIServerMockRecorder) GetABCIBlockEvents(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetABCIBlockEvents", reflect.TypeOf((*MockEventProviderAPIServer)(nil).GetABCIBlockEvents), arg0, arg1)
}

// GetABCIBlockEventsAtHeight mocks base method.
func (m *MockEventProviderAPIServer) GetABCIBlockEventsAtHeight(arg0 context.Context, arg1 *GetABCIBlockEventsAtHeightRequest) (*GetABCIBlockEventsAtHeightResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetABCIBlockEventsAtHeight", arg0, arg1)
	ret0, _ := ret[0].(*GetABCIBlockEventsAtHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetABCIBlockEventsAtHeight indicates an expected call of GetABCIBlockEventsAtHeight.
func (mr *MockEventProviderAPIServerMockRecorder) GetABCIBlockEventsAtHeight(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetABCIBlockEventsAtHeight", reflect.TypeOf((*MockEventProviderAPIServer)(nil).GetABCIBlockEventsAtHeight), arg0, arg1)
}

// GetBlockEventsRPC mocks base method.
func (m *MockEventProviderAPIServer) GetBlockEventsRPC(arg0 context.Context, arg1 *GetBlockEventsRPCRequest) (*GetBlockEventsRPCResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlockEventsRPC", arg0, arg1)
	ret0, _ := ret[0].(*GetBlockEventsRPCResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockEventsRPC indicates an expected call of GetBlockEventsRPC.
func (mr *MockEventProviderAPIServerMockRecorder) GetBlockEventsRPC(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockEventsRPC", reflect.TypeOf((*MockEventProviderAPIServer)(nil).GetBlockEventsRPC), arg0, arg1)
}

// GetCustomEventsRPC mocks base method.
func (m *MockEventProviderAPIServer) GetCustomEventsRPC(arg0 context.Context, arg1 *GetCustomEventsRPCRequest) (*GetCustomEventsRPCResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCustomEventsRPC", arg0, arg1)
	ret0, _ := ret[0].(*GetCustomEventsRPCResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCustomEventsRPC indicates an expected call of GetCustomEventsRPC.
func (mr *MockEventProviderAPIServerMockRecorder) GetCustomEventsRPC(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCustomEventsRPC", reflect.TypeOf((*MockEventProviderAPIServer)(nil).GetCustomEventsRPC), arg0, arg1)
}

// GetLatestHeight mocks base method.
func (m *MockEventProviderAPIServer) GetLatestHeight(arg0 context.Context, arg1 *GetLatestHeightRequest) (*GetLatestHeightResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestHeight", arg0, arg1)
	ret0, _ := ret[0].(*GetLatestHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestHeight indicates an expected call of GetLatestHeight.
func (mr *MockEventProviderAPIServerMockRecorder) GetLatestHeight(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestHeight", reflect.TypeOf((*MockEventProviderAPIServer)(nil).GetLatestHeight), arg0, arg1)
}

// StreamBlockEvents mocks base method.
func (m *MockEventProviderAPIServer) StreamBlockEvents(arg0 *StreamBlockEventsRequest, arg1 EventProviderAPI_StreamBlockEventsServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StreamBlockEvents", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StreamBlockEvents indicates an expected call of StreamBlockEvents.
func (mr *MockEventProviderAPIServerMockRecorder) StreamBlockEvents(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StreamBlockEvents", reflect.TypeOf((*MockEventProviderAPIServer)(nil).StreamBlockEvents), arg0, arg1)
}

// mustEmbedUnimplementedEventProviderAPIServer mocks base method.
func (m *MockEventProviderAPIServer) mustEmbedUnimplementedEventProviderAPIServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedEventProviderAPIServer")
}

// mustEmbedUnimplementedEventProviderAPIServer indicates an expected call of mustEmbedUnimplementedEventProviderAPIServer.
func (mr *MockEventProviderAPIServerMockRecorder) mustEmbedUnimplementedEventProviderAPIServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedEventProviderAPIServer", reflect.TypeOf((*MockEventProviderAPIServer)(nil).mustEmbedUnimplementedEventProviderAPIServer))
}

// MockUnsafeEventProviderAPIServer is a mock of UnsafeEventProviderAPIServer interface.
type MockUnsafeEventProviderAPIServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeEventProviderAPIServerMockRecorder
}

// MockUnsafeEventProviderAPIServerMockRecorder is the mock recorder for MockUnsafeEventProviderAPIServer.
type MockUnsafeEventProviderAPIServerMockRecorder struct {
	mock *MockUnsafeEventProviderAPIServer
}

// NewMockUnsafeEventProviderAPIServer creates a new mock instance.
func NewMockUnsafeEventProviderAPIServer(ctrl *gomock.Controller) *MockUnsafeEventProviderAPIServer {
	mock := &MockUnsafeEventProviderAPIServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeEventProviderAPIServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeEventProviderAPIServer) EXPECT() *MockUnsafeEventProviderAPIServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedEventProviderAPIServer mocks base method.
func (m *MockUnsafeEventProviderAPIServer) mustEmbedUnimplementedEventProviderAPIServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedEventProviderAPIServer")
}

// mustEmbedUnimplementedEventProviderAPIServer indicates an expected call of mustEmbedUnimplementedEventProviderAPIServer.
func (mr *MockUnsafeEventProviderAPIServerMockRecorder) mustEmbedUnimplementedEventProviderAPIServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedEventProviderAPIServer", reflect.TypeOf((*MockUnsafeEventProviderAPIServer)(nil).mustEmbedUnimplementedEventProviderAPIServer))
}

// MockEventProviderAPI_StreamBlockEventsServer is a mock of EventProviderAPI_StreamBlockEventsServer interface.
type MockEventProviderAPI_StreamBlockEventsServer struct {
	ctrl     *gomock.Controller
	recorder *MockEventProviderAPI_StreamBlockEventsServerMockRecorder
}

// MockEventProviderAPI_StreamBlockEventsServerMockRecorder is the mock recorder for MockEventProviderAPI_StreamBlockEventsServer.
type MockEventProviderAPI_StreamBlockEventsServerMockRecorder struct {
	mock *MockEventProviderAPI_StreamBlockEventsServer
}

// NewMockEventProviderAPI_StreamBlockEventsServer creates a new mock instance.
func NewMockEventProviderAPI_StreamBlockEventsServer(ctrl *gomock.Controller) *MockEventProviderAPI_StreamBlockEventsServer {
	mock := &MockEventProviderAPI_StreamBlockEventsServer{ctrl: ctrl}
	mock.recorder = &MockEventProviderAPI_StreamBlockEventsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventProviderAPI_StreamBlockEventsServer) EXPECT() *MockEventProviderAPI_StreamBlockEventsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockEventProviderAPI_StreamBlockEventsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockEventProviderAPI_StreamBlockEventsServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockEventProviderAPI_StreamBlockEventsServerMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsServer) Send(arg0 *StreamBlockEventsResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockEventProviderAPI_StreamBlockEventsServerMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockEventProviderAPI_StreamBlockEventsServerMockRecorder) SendHeader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockEventProviderAPI_StreamBlockEventsServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockEventProviderAPI_StreamBlockEventsServerMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockEventProviderAPI_StreamBlockEventsServerMockRecorder) SetHeader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockEventProviderAPI_StreamBlockEventsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockEventProviderAPI_StreamBlockEventsServerMockRecorder) SetTrailer(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockEventProviderAPI_StreamBlockEventsServer)(nil).SetTrailer), arg0)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: cdnsystem/cdnsystem.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	base "d7y.io/dragonfly/v2/pkg/rpc/base"
	cdnsystem "d7y.io/dragonfly/v2/pkg/rpc/cdnsystem"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockSeederClient is a mock of SeederClient interface.
type MockSeederClient struct {
	ctrl     *gomock.Controller
	recorder *MockSeederClientMockRecorder
}

// MockSeederClientMockRecorder is the mock recorder for MockSeederClient.
type MockSeederClientMockRecorder struct {
	mock *MockSeederClient
}

// NewMockSeederClient creates a new mock instance.
func NewMockSeederClient(ctrl *gomock.Controller) *MockSeederClient {
	mock := &MockSeederClient{ctrl: ctrl}
	mock.recorder = &MockSeederClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeederClient) EXPECT() *MockSeederClientMockRecorder {
	return m.recorder
}

// GetPieceTasks mocks base method.
func (m *MockSeederClient) GetPieceTasks(ctx context.Context, in *base.PieceTaskRequest, opts ...grpc.CallOption) (*base.PiecePacket, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPieceTasks", varargs...)
	ret0, _ := ret[0].(*base.PiecePacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieceTasks indicates an expected call of GetPieceTasks.
func (mr *MockSeederClientMockRecorder) GetPieceTasks(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieceTasks", reflect.TypeOf((*MockSeederClient)(nil).GetPieceTasks), varargs...)
}

// ObtainSeeds mocks base method.
func (m *MockSeederClient) ObtainSeeds(ctx context.Context, in *cdnsystem.SeedRequest, opts ...grpc.CallOption) (cdnsystem.Seeder_ObtainSeedsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObtainSeeds", varargs...)
	ret0, _ := ret[0].(cdnsystem.Seeder_ObtainSeedsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObtainSeeds indicates an expected call of ObtainSeeds.
func (mr *MockSeederClientMockRecorder) ObtainSeeds(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObtainSeeds", reflect.TypeOf((*MockSeederClient)(nil).ObtainSeeds), varargs...)
}

// SyncPieceTasks mocks base method.
func (m *MockSeederClient) SyncPieceTasks(ctx context.Context, opts ...grpc.CallOption) (cdnsystem.Seeder_SyncPieceTasksClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncPieceTasks", varargs...)
	ret0, _ := ret[0].(cdnsystem.Seeder_SyncPieceTasksClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncPieceTasks indicates an expected call of SyncPieceTasks.
func (mr *MockSeederClientMockRecorder) SyncPieceTasks(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPieceTasks", reflect.TypeOf((*MockSeederClient)(nil).SyncPieceTasks), varargs...)
}

// MockSeeder_ObtainSeedsClient is a mock of Seeder_ObtainSeedsClient interface.
type MockSeeder_ObtainSeedsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSeeder_ObtainSeedsClientMockRecorder
}

// MockSeeder_ObtainSeedsClientMockRecorder is the mock recorder for MockSeeder_ObtainSeedsClient.
type MockSeeder_ObtainSeedsClientMockRecorder struct {
	mock *MockSeeder_ObtainSeedsClient
}

// NewMockSeeder_ObtainSeedsClient creates a new mock instance.
func NewMockSeeder_ObtainSeedsClient(ctrl *gomock.Controller) *MockSeeder_ObtainSeedsClient {
	mock := &MockSeeder_ObtainSeedsClient{ctrl: ctrl}
	mock.recorder = &MockSeeder_ObtainSeedsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeeder_ObtainSeedsClient) EXPECT() *MockSeeder_ObtainSeedsClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockSeeder_ObtainSeedsClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockSeeder_ObtainSeedsClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockSeeder_ObtainSeedsClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockSeeder_ObtainSeedsClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSeeder_ObtainSeedsClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSeeder_ObtainSeedsClient)(nil).Context))
}

// Header mocks base method.
func (m *MockSeeder_ObtainSeedsClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockSeeder_ObtainSeedsClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockSeeder_ObtainSeedsClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockSeeder_ObtainSeedsClient) Recv() (*cdnsystem.PieceSeed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*cdnsystem.PieceSeed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockSeeder_ObtainSeedsClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockSeeder_ObtainSeedsClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockSeeder_ObtainSeedsClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockSeeder_ObtainSeedsClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSeeder_ObtainSeedsClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockSeeder_ObtainSeedsClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockSeeder_ObtainSeedsClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSeeder_ObtainSeedsClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockSeeder_ObtainSeedsClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockSeeder_ObtainSeedsClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockSeeder_ObtainSeedsClient)(nil).Trailer))
}

// MockSeeder_SyncPieceTasksClient is a mock of Seeder_SyncPieceTasksClient interface.
type MockSeeder_SyncPieceTasksClient struct {
	ctrl     *gomock.Controller
	recorder *MockSeeder_SyncPieceTasksClientMockRecorder
}

// MockSeeder_SyncPieceTasksClientMockRecorder is the mock recorder for MockSeeder_SyncPieceTasksClient.
type MockSeeder_SyncPieceTasksClientMockRecorder struct {
	mock *MockSeeder_SyncPieceTasksClient
}

// NewMockSeeder_SyncPieceTasksClient creates a new mock instance.
func NewMockSeeder_SyncPieceTasksClient(ctrl *gomock.Controller) *MockSeeder_SyncPieceTasksClient {
	mock := &MockSeeder_SyncPieceTasksClient{ctrl: ctrl}
	mock.recorder = &MockSeeder_SyncPieceTasksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeeder_SyncPieceTasksClient) EXPECT() *MockSeeder_SyncPieceTasksClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockSeeder_SyncPieceTasksClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockSeeder_SyncPieceTasksClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockSeeder_SyncPieceTasksClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockSeeder_SyncPieceTasksClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSeeder_SyncPieceTasksClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSeeder_SyncPieceTasksClient)(nil).Context))
}

// Header mocks base method.
func (m *MockSeeder_SyncPieceTasksClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockSeeder_SyncPieceTasksClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockSeeder_SyncPieceTasksClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockSeeder_SyncPieceTasksClient) Recv() (*base.PiecePacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*base.PiecePacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockSeeder_SyncPieceTasksClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockSeeder_SyncPieceTasksClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockSeeder_SyncPieceTasksClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockSeeder_SyncPieceTasksClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSeeder_SyncPieceTasksClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockSeeder_SyncPieceTasksClient) Send(arg0 *base.PieceTaskRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockSeeder_SyncPieceTasksClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSeeder_SyncPieceTasksClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockSeeder_SyncPieceTasksClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockSeeder_SyncPieceTasksClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSeeder_SyncPieceTasksClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockSeeder_SyncPieceTasksClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockSeeder_SyncPieceTasksClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockSeeder_SyncPieceTasksClient)(nil).Trailer))
}

// MockSeederServer is a mock of SeederServer interface.
type MockSeederServer struct {
	ctrl     *gomock.Controller
	recorder *MockSeederServerMockRecorder
}

// MockSeederServerMockRecorder is the mock recorder for MockSeederServer.
type MockSeederServerMockRecorder struct {
	mock *MockSeederServer
}

// NewMockSeederServer creates a new mock instance.
func NewMockSeederServer(ctrl *gomock.Controller) *MockSeederServer {
	mock := &MockSeederServer{ctrl: ctrl}
	mock.recorder = &MockSeederServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeederServer) EXPECT() *MockSeederServerMockRecorder {
	return m.recorder
}

// GetPieceTasks mocks base method.
func (m *MockSeederServer) GetPieceTasks(arg0 context.Context, arg1 *base.PieceTaskRequest) (*base.PiecePacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPieceTasks", arg0, arg1)
	ret0, _ := ret[0].(*base.PiecePacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPieceTasks indicates an expected call of GetPieceTasks.
func (mr *MockSeederServerMockRecorder) GetPieceTasks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPieceTasks", reflect.TypeOf((*MockSeederServer)(nil).GetPieceTasks), arg0, arg1)
}

// ObtainSeeds mocks base method.
func (m *MockSeederServer) ObtainSeeds(arg0 *cdnsystem.SeedRequest, arg1 cdnsystem.Seeder_ObtainSeedsServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObtainSeeds", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ObtainSeeds indicates an expected call of ObtainSeeds.
func (mr *MockSeederServerMockRecorder) ObtainSeeds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObtainSeeds", reflect.TypeOf((*MockSeederServer)(nil).ObtainSeeds), arg0, arg1)
}

// SyncPieceTasks mocks base method.
func (m *MockSeederServer) SyncPieceTasks(arg0 cdnsystem.Seeder_SyncPieceTasksServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncPieceTasks", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncPieceTasks indicates an expected call of SyncPieceTasks.
func (mr *MockSeederServerMockRecorder) SyncPieceTasks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPieceTasks", reflect.TypeOf((*MockSeederServer)(nil).SyncPieceTasks), arg0)
}

// MockSeeder_ObtainSeedsServer is a mock of Seeder_ObtainSeedsServer interface.
type MockSeeder_ObtainSeedsServer struct {
	ctrl     *gomock.Controller
	recorder *MockSeeder_ObtainSeedsServerMockRecorder
}

// MockSeeder_ObtainSeedsServerMockRecorder is the mock recorder for MockSeeder_ObtainSeedsServer.
type MockSeeder_ObtainSeedsServerMockRecorder struct {
	mock *MockSeeder_ObtainSeedsServer
}

// NewMockSeeder_ObtainSeedsServer creates a new mock instance.
func NewMockSeeder_ObtainSeedsServer(ctrl *gomock.Controller) *MockSeeder_ObtainSeedsServer {
	mock := &MockSeeder_ObtainSeedsServer{ctrl: ctrl}
	mock.recorder = &MockSeeder_ObtainSeedsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeeder_ObtainSeedsServer) EXPECT() *MockSeeder_ObtainSeedsServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockSeeder_ObtainSeedsServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSeeder_ObtainSeedsServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSeeder_ObtainSeedsServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockSeeder_ObtainSeedsServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockSeeder_ObtainSeedsServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSeeder_ObtainSeedsServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockSeeder_ObtainSeedsServer) Send(arg0 *cdnsystem.PieceSeed) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockSeeder_ObtainSeedsServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSeeder_ObtainSeedsServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockSeeder_ObtainSeedsServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockSeeder_ObtainSeedsServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockSeeder_ObtainSeedsServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockSeeder_ObtainSeedsServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockSeeder_ObtainSeedsServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSeeder_ObtainSeedsServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockSeeder_ObtainSeedsServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockSeeder_ObtainSeedsServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockSeeder_ObtainSeedsServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockSeeder_ObtainSeedsServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockSeeder_ObtainSeedsServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockSeeder_ObtainSeedsServer)(nil).SetTrailer), arg0)
}

// MockSeeder_SyncPieceTasksServer is a mock of Seeder_SyncPieceTasksServer interface.
type MockSeeder_SyncPieceTasksServer struct {
	ctrl     *gomock.Controller
	recorder *MockSeeder_SyncPieceTasksServerMockRecorder
}

// MockSeeder_SyncPieceTasksServerMockRecorder is the mock recorder for MockSeeder_SyncPieceTasksServer.
type MockSeeder_SyncPieceTasksServerMockRecorder struct {
	mock *MockSeeder_SyncPieceTasksServer
}

// NewMockSeeder_SyncPieceTasksServer creates a new mock instance.
func NewMockSeeder_SyncPieceTasksServer(ctrl *gomock.Controller) *MockSeeder_SyncPieceTasksServer {
	mock := &MockSeeder_SyncPieceTasksServer{ctrl: ctrl}
	mock.recorder = &MockSeeder_SyncPieceTasksServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeeder_SyncPieceTasksServer) EXPECT() *MockSeeder_SyncPieceTasksServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockSeeder_SyncPieceTasksServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockSeeder_SyncPieceTasksServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockSeeder_SyncPieceTasksServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockSeeder_SyncPieceTasksServer) Recv() (*base.PieceTaskRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*base.PieceTaskRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockSeeder_SyncPieceTasksServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockSeeder_SyncPieceTasksServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockSeeder_SyncPieceTasksServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockSeeder_SyncPieceTasksServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockSeeder_SyncPieceTasksServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockSeeder_SyncPieceTasksServer) Send(arg0 *base.PiecePacket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockSeeder_SyncPieceTasksServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSeeder_SyncPieceTasksServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockSeeder_SyncPieceTasksServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockSeeder_SyncPieceTasksServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockSeeder_SyncPieceTasksServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockSeeder_SyncPieceTasksServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockSeeder_SyncPieceTasksServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockSeeder_SyncPieceTasksServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockSeeder_SyncPieceTasksServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockSeeder_SyncPieceTasksServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockSeeder_SyncPieceTasksServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockSeeder_SyncPieceTasksServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockSeeder_SyncPieceTasksServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockSeeder_SyncPieceTasksServer)(nil).SetTrailer), arg0)
}
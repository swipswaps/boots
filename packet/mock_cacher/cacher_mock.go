// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/packethost/cacher/protos/cacher (interfaces: CacherClient)

// Package mock_cacher is a generated GoMock package.
package mock_cacher

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cacher "github.com/packethost/cacher/protos/cacher"
	grpc "google.golang.org/grpc"
)

// MockCacherClient is a mock of CacherClient interface
type MockCacherClient struct {
	ctrl     *gomock.Controller
	recorder *MockCacherClientMockRecorder
}

// MockCacherClientMockRecorder is the mock recorder for MockCacherClient
type MockCacherClientMockRecorder struct {
	mock *MockCacherClient
}

// NewMockCacherClient creates a new mock instance
func NewMockCacherClient(ctrl *gomock.Controller) *MockCacherClient {
	mock := &MockCacherClient{ctrl: ctrl}
	mock.recorder = &MockCacherClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCacherClient) EXPECT() *MockCacherClientMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockCacherClient) All(arg0 context.Context, arg1 *cacher.Empty, arg2 ...grpc.CallOption) (cacher.Cacher_AllClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "All", varargs...)
	ret0, _ := ret[0].(cacher.Cacher_AllClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockCacherClientMockRecorder) All(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockCacherClient)(nil).All), varargs...)
}

// ByID mocks base method
func (m *MockCacherClient) ByID(arg0 context.Context, arg1 *cacher.GetRequest, arg2 ...grpc.CallOption) (*cacher.Hardware, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ByID", varargs...)
	ret0, _ := ret[0].(*cacher.Hardware)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByID indicates an expected call of ByID
func (mr *MockCacherClientMockRecorder) ByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByID", reflect.TypeOf((*MockCacherClient)(nil).ByID), varargs...)
}

// ByIP mocks base method
func (m *MockCacherClient) ByIP(arg0 context.Context, arg1 *cacher.GetRequest, arg2 ...grpc.CallOption) (*cacher.Hardware, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ByIP", varargs...)
	ret0, _ := ret[0].(*cacher.Hardware)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByIP indicates an expected call of ByIP
func (mr *MockCacherClientMockRecorder) ByIP(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByIP", reflect.TypeOf((*MockCacherClient)(nil).ByIP), varargs...)
}

// ByMAC mocks base method
func (m *MockCacherClient) ByMAC(arg0 context.Context, arg1 *cacher.GetRequest, arg2 ...grpc.CallOption) (*cacher.Hardware, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ByMAC", varargs...)
	ret0, _ := ret[0].(*cacher.Hardware)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByMAC indicates an expected call of ByMAC
func (mr *MockCacherClientMockRecorder) ByMAC(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByMAC", reflect.TypeOf((*MockCacherClient)(nil).ByMAC), varargs...)
}

// Ingest mocks base method
func (m *MockCacherClient) Ingest(arg0 context.Context, arg1 *cacher.Empty, arg2 ...grpc.CallOption) (*cacher.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Ingest", varargs...)
	ret0, _ := ret[0].(*cacher.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ingest indicates an expected call of Ingest
func (mr *MockCacherClientMockRecorder) Ingest(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ingest", reflect.TypeOf((*MockCacherClient)(nil).Ingest), varargs...)
}

// Push mocks base method
func (m *MockCacherClient) Push(arg0 context.Context, arg1 *cacher.PushRequest, arg2 ...grpc.CallOption) (*cacher.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Push", varargs...)
	ret0, _ := ret[0].(*cacher.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Push indicates an expected call of Push
func (mr *MockCacherClientMockRecorder) Push(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockCacherClient)(nil).Push), varargs...)
}

// Watch mocks base method
func (m *MockCacherClient) Watch(arg0 context.Context, arg1 *cacher.GetRequest, arg2 ...grpc.CallOption) (cacher.Cacher_WatchClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(cacher.Cacher_WatchClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockCacherClientMockRecorder) Watch(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockCacherClient)(nil).Watch), varargs...)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/southbound/defs.go

// Package southbound is a generated GoMock package.
package southbound

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/onosproject/onos-config/pkg/southbound"
	topodevice "github.com/onosproject/onos-topo/api/device"
	"github.com/openconfig/gnmi/client"
	"github.com/openconfig/gnmi/proto/gnmi"
	"reflect"
)

// MockTargetIf is a mock of TargetIf interface
type MockTargetIf struct {
	ctrl     *gomock.Controller
	recorder *MockTargetIfMockRecorder
}

// MockTargetIfMockRecorder is the mock recorder for MockTargetIf
type MockTargetIfMockRecorder struct {
	mock *MockTargetIf
}

// NewMockTargetIf creates a new mock instance
func NewMockTargetIf(ctrl *gomock.Controller) *MockTargetIf {
	mock := &MockTargetIf{ctrl: ctrl}
	mock.recorder = &MockTargetIfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTargetIf) EXPECT() *MockTargetIfMockRecorder {
	return m.recorder
}

// ConnectTarget mocks base method
func (m *MockTargetIf) ConnectTarget(ctx context.Context, device topodevice.Device) (topodevice.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectTarget", ctx, device)
	ret0, _ := ret[0].(topodevice.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectTarget indicates an expected call of ConnectTarget
func (mr *MockTargetIfMockRecorder) ConnectTarget(ctx, device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectTarget", reflect.TypeOf((*MockTargetIf)(nil).ConnectTarget), ctx, device)
}

// CapabilitiesWithString mocks base method
func (m *MockTargetIf) CapabilitiesWithString(ctx context.Context, request string) (*gnmi.CapabilityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CapabilitiesWithString", ctx, request)
	ret0, _ := ret[0].(*gnmi.CapabilityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CapabilitiesWithString indicates an expected call of CapabilitiesWithString
func (mr *MockTargetIfMockRecorder) CapabilitiesWithString(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CapabilitiesWithString", reflect.TypeOf((*MockTargetIf)(nil).CapabilitiesWithString), ctx, request)
}

// Get mocks base method
func (m *MockTargetIf) Get(ctx context.Context, request *gnmi.GetRequest) (*gnmi.GetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, request)
	ret0, _ := ret[0].(*gnmi.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTargetIfMockRecorder) Get(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTargetIf)(nil).Get), ctx, request)
}

// GetWithString mocks base method
func (m *MockTargetIf) GetWithString(ctx context.Context, request string) (*gnmi.GetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWithString", ctx, request)
	ret0, _ := ret[0].(*gnmi.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithString indicates an expected call of GetWithString
func (mr *MockTargetIfMockRecorder) GetWithString(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithString", reflect.TypeOf((*MockTargetIf)(nil).GetWithString), ctx, request)
}

// Set mocks base method
func (m *MockTargetIf) Set(ctx context.Context, request *gnmi.SetRequest) (*gnmi.SetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, request)
	ret0, _ := ret[0].(*gnmi.SetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockTargetIfMockRecorder) Set(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockTargetIf)(nil).Set), ctx, request)
}

// SetWithString mocks base method
func (m *MockTargetIf) SetWithString(ctx context.Context, request string) (*gnmi.SetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWithString", ctx, request)
	ret0, _ := ret[0].(*gnmi.SetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetWithString indicates an expected call of SetWithString
func (mr *MockTargetIfMockRecorder) SetWithString(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWithString", reflect.TypeOf((*MockTargetIf)(nil).SetWithString), ctx, request)
}

// Subscribe mocks base method
func (m *MockTargetIf) Subscribe(ctx context.Context, request *gnmi.SubscribeRequest, handler client.ProtoHandler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, request, handler)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockTargetIfMockRecorder) Subscribe(ctx, request, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockTargetIf)(nil).Subscribe), ctx, request, handler)
}

// Context mocks base method
func (m *MockTargetIf) Context() *context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(*context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockTargetIfMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTargetIf)(nil).Context))
}

// Destination mocks base method
func (m *MockTargetIf) Destination() *client.Destination {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destination")
	ret0, _ := ret[0].(*client.Destination)
	return ret0
}

// Destination indicates an expected call of Destination
func (mr *MockTargetIfMockRecorder) Dest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destination", reflect.TypeOf((*MockTargetIf)(nil).Destination))
}

// Client mocks base method
func (m *MockTargetIf) Client() *southbound.GnmiClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*southbound.GnmiClient)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockTargetIfMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockTargetIf)(nil).Client))
}

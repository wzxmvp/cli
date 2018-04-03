// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/cloudfoundry/bosh-agent/agentclient (interfaces: AgentClient)

package mocks

import (
	agentclient "github.com/cloudfoundry/bosh-agent/agentclient"
	applyspec "github.com/cloudfoundry/bosh-agent/agentclient/applyspec"
	gomock "github.com/golang/mock/gomock"
)

// Mock of AgentClient interface
type MockAgentClient struct {
	ctrl     *gomock.Controller
	recorder *_MockAgentClientRecorder
}

// Recorder for MockAgentClient (not exported)
type _MockAgentClientRecorder struct {
	mock *MockAgentClient
}

func NewMockAgentClient(ctrl *gomock.Controller) *MockAgentClient {
	mock := &MockAgentClient{ctrl: ctrl}
	mock.recorder = &_MockAgentClientRecorder{mock}
	return mock
}

func (_m *MockAgentClient) EXPECT() *_MockAgentClientRecorder {
	return _m.recorder
}

func (_m *MockAgentClient) Apply(_param0 applyspec.ApplySpec) error {
	ret := _m.ctrl.Call(_m, "Apply", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAgentClientRecorder) Apply(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Apply", arg0)
}

func (_m *MockAgentClient) CompilePackage(_param0 agentclient.BlobRef, _param1 []agentclient.BlobRef) (agentclient.BlobRef, error) {
	ret := _m.ctrl.Call(_m, "CompilePackage", _param0, _param1)
	ret0, _ := ret[0].(agentclient.BlobRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAgentClientRecorder) CompilePackage(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CompilePackage", arg0, arg1)
}

func (_m *MockAgentClient) DeleteARPEntries(_param0 []string) error {
	ret := _m.ctrl.Call(_m, "DeleteARPEntries", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAgentClientRecorder) DeleteARPEntries(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteARPEntries", arg0)
}

func (_m *MockAgentClient) GetState() (agentclient.AgentState, error) {
	ret := _m.ctrl.Call(_m, "GetState")
	ret0, _ := ret[0].(agentclient.AgentState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAgentClientRecorder) GetState() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetState")
}

func (_m *MockAgentClient) ListDisk() ([]string, error) {
	ret := _m.ctrl.Call(_m, "ListDisk")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAgentClientRecorder) ListDisk() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListDisk")
}

func (_m *MockAgentClient) MigrateDisk() error {
	ret := _m.ctrl.Call(_m, "MigrateDisk")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAgentClientRecorder) MigrateDisk() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MigrateDisk")
}

func (_m *MockAgentClient) MountDisk(_param0 string) error {
	ret := _m.ctrl.Call(_m, "MountDisk", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAgentClientRecorder) MountDisk(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MountDisk", arg0)
}

func (_m *MockAgentClient) Ping() (string, error) {
	ret := _m.ctrl.Call(_m, "Ping")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAgentClientRecorder) Ping() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ping")
}

func (_m *MockAgentClient) RunScript(_param0 string, _param1 map[string]interface{}) error {
	ret := _m.ctrl.Call(_m, "RunScript", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAgentClientRecorder) RunScript(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RunScript", arg0, arg1)
}

func (_m *MockAgentClient) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAgentClientRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockAgentClient) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAgentClientRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

func (_m *MockAgentClient) SyncDNS(_param0 string, _param1 string, _param2 uint64) (string, error) {
	ret := _m.ctrl.Call(_m, "SyncDNS", _param0, _param1, _param2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAgentClientRecorder) SyncDNS(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SyncDNS", arg0, arg1, arg2)
}

func (_m *MockAgentClient) UnmountDisk(_param0 string) error {
	ret := _m.ctrl.Call(_m, "UnmountDisk", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAgentClientRecorder) UnmountDisk(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnmountDisk", arg0)
}

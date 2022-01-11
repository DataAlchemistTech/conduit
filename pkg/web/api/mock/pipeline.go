// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit/pkg/web/api (interfaces: PipelineOrchestrator)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	pipeline "github.com/conduitio/conduit/pkg/pipeline"
	gomock "github.com/golang/mock/gomock"
)

// PipelineOrchestrator is a mock of PipelineOrchestrator interface.
type PipelineOrchestrator struct {
	ctrl     *gomock.Controller
	recorder *PipelineOrchestratorMockRecorder
}

// PipelineOrchestratorMockRecorder is the mock recorder for PipelineOrchestrator.
type PipelineOrchestratorMockRecorder struct {
	mock *PipelineOrchestrator
}

// NewPipelineOrchestrator creates a new mock instance.
func NewPipelineOrchestrator(ctrl *gomock.Controller) *PipelineOrchestrator {
	mock := &PipelineOrchestrator{ctrl: ctrl}
	mock.recorder = &PipelineOrchestratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *PipelineOrchestrator) EXPECT() *PipelineOrchestratorMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *PipelineOrchestrator) Create(arg0 context.Context, arg1 pipeline.Config) (*pipeline.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*pipeline.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *PipelineOrchestratorMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*PipelineOrchestrator)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *PipelineOrchestrator) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *PipelineOrchestratorMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*PipelineOrchestrator)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *PipelineOrchestrator) Get(arg0 context.Context, arg1 string) (*pipeline.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*pipeline.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *PipelineOrchestratorMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*PipelineOrchestrator)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *PipelineOrchestrator) List(arg0 context.Context) map[string]*pipeline.Instance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(map[string]*pipeline.Instance)
	return ret0
}

// List indicates an expected call of List.
func (mr *PipelineOrchestratorMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*PipelineOrchestrator)(nil).List), arg0)
}

// Start mocks base method.
func (m *PipelineOrchestrator) Start(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *PipelineOrchestratorMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*PipelineOrchestrator)(nil).Start), arg0, arg1)
}

// Stop mocks base method.
func (m *PipelineOrchestrator) Stop(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *PipelineOrchestratorMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*PipelineOrchestrator)(nil).Stop), arg0, arg1)
}

// Update mocks base method.
func (m *PipelineOrchestrator) Update(arg0 context.Context, arg1 string, arg2 pipeline.Config) (*pipeline.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*pipeline.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *PipelineOrchestratorMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*PipelineOrchestrator)(nil).Update), arg0, arg1, arg2)
}

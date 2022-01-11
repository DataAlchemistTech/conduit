// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit/pkg/connector (interfaces: Source,Destination)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	connector "github.com/conduitio/conduit/pkg/connector"
	record "github.com/conduitio/conduit/pkg/record"
	gomock "github.com/golang/mock/gomock"
)

// Source is a mock of Source interface.
type Source struct {
	ctrl     *gomock.Controller
	recorder *SourceMockRecorder
}

// SourceMockRecorder is the mock recorder for Source.
type SourceMockRecorder struct {
	mock *Source
}

// NewSource creates a new mock instance.
func NewSource(ctrl *gomock.Controller) *Source {
	mock := &Source{ctrl: ctrl}
	mock.recorder = &SourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Source) EXPECT() *SourceMockRecorder {
	return m.recorder
}

// Ack mocks base method.
func (m *Source) Ack(arg0 context.Context, arg1 record.Position) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack.
func (mr *SourceMockRecorder) Ack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*Source)(nil).Ack), arg0, arg1)
}

// Config mocks base method.
func (m *Source) Config() connector.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(connector.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *SourceMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*Source)(nil).Config))
}

// Errors mocks base method.
func (m *Source) Errors() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Errors")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// Errors indicates an expected call of Errors.
func (mr *SourceMockRecorder) Errors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errors", reflect.TypeOf((*Source)(nil).Errors))
}

// ID mocks base method.
func (m *Source) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *SourceMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*Source)(nil).ID))
}

// IsRunning mocks base method.
func (m *Source) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning.
func (mr *SourceMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*Source)(nil).IsRunning))
}

// Open mocks base method.
func (m *Source) Open(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *SourceMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*Source)(nil).Open), arg0)
}

// Read mocks base method.
func (m *Source) Read(arg0 context.Context) (record.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(record.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *SourceMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*Source)(nil).Read), arg0)
}

// SetConfig mocks base method.
func (m *Source) SetConfig(arg0 connector.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConfig", arg0)
}

// SetConfig indicates an expected call of SetConfig.
func (mr *SourceMockRecorder) SetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*Source)(nil).SetConfig), arg0)
}

// SetState mocks base method.
func (m *Source) SetState(arg0 connector.SourceState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetState", arg0)
}

// SetState indicates an expected call of SetState.
func (mr *SourceMockRecorder) SetState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*Source)(nil).SetState), arg0)
}

// State mocks base method.
func (m *Source) State() connector.SourceState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(connector.SourceState)
	return ret0
}

// State indicates an expected call of State.
func (mr *SourceMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*Source)(nil).State))
}

// Teardown mocks base method.
func (m *Source) Teardown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teardown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Teardown indicates an expected call of Teardown.
func (mr *SourceMockRecorder) Teardown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teardown", reflect.TypeOf((*Source)(nil).Teardown), arg0)
}

// Type mocks base method.
func (m *Source) Type() connector.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(connector.Type)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *SourceMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*Source)(nil).Type))
}

// Validate mocks base method.
func (m *Source) Validate(arg0 context.Context, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *SourceMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*Source)(nil).Validate), arg0, arg1)
}

// Destination is a mock of Destination interface.
type Destination struct {
	ctrl     *gomock.Controller
	recorder *DestinationMockRecorder
}

// DestinationMockRecorder is the mock recorder for Destination.
type DestinationMockRecorder struct {
	mock *Destination
}

// NewDestination creates a new mock instance.
func NewDestination(ctrl *gomock.Controller) *Destination {
	mock := &Destination{ctrl: ctrl}
	mock.recorder = &DestinationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Destination) EXPECT() *DestinationMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *Destination) Config() connector.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(connector.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *DestinationMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*Destination)(nil).Config))
}

// Errors mocks base method.
func (m *Destination) Errors() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Errors")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// Errors indicates an expected call of Errors.
func (mr *DestinationMockRecorder) Errors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errors", reflect.TypeOf((*Destination)(nil).Errors))
}

// ID mocks base method.
func (m *Destination) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *DestinationMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*Destination)(nil).ID))
}

// IsRunning mocks base method.
func (m *Destination) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning.
func (mr *DestinationMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*Destination)(nil).IsRunning))
}

// Open mocks base method.
func (m *Destination) Open(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *DestinationMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*Destination)(nil).Open), arg0)
}

// SetConfig mocks base method.
func (m *Destination) SetConfig(arg0 connector.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConfig", arg0)
}

// SetConfig indicates an expected call of SetConfig.
func (mr *DestinationMockRecorder) SetConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*Destination)(nil).SetConfig), arg0)
}

// SetState mocks base method.
func (m *Destination) SetState(arg0 connector.DestinationState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetState", arg0)
}

// SetState indicates an expected call of SetState.
func (mr *DestinationMockRecorder) SetState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*Destination)(nil).SetState), arg0)
}

// State mocks base method.
func (m *Destination) State() connector.DestinationState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(connector.DestinationState)
	return ret0
}

// State indicates an expected call of State.
func (mr *DestinationMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*Destination)(nil).State))
}

// Teardown mocks base method.
func (m *Destination) Teardown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Teardown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Teardown indicates an expected call of Teardown.
func (mr *DestinationMockRecorder) Teardown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Teardown", reflect.TypeOf((*Destination)(nil).Teardown), arg0)
}

// Type mocks base method.
func (m *Destination) Type() connector.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(connector.Type)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *DestinationMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*Destination)(nil).Type))
}

// Validate mocks base method.
func (m *Destination) Validate(arg0 context.Context, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *DestinationMockRecorder) Validate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*Destination)(nil).Validate), arg0, arg1)
}

// Write mocks base method.
func (m *Destination) Write(arg0 context.Context, arg1 record.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *DestinationMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*Destination)(nil).Write), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit/pkg/record/schema (interfaces: MutableSchema,MutableStructDescriptor,MutableField,MutableMapDescriptor,MutableArrayDescriptor,MutablePrimitiveDescriptor,MutableEnumDescriptor,MutableEnumValueDescriptor)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	schema "github.com/conduitio/conduit/pkg/record/schema"
	gomock "github.com/golang/mock/gomock"
)

// MutableSchema is a mock of MutableSchema interface.
type MutableSchema struct {
	ctrl     *gomock.Controller
	recorder *MutableSchemaMockRecorder
}

// MutableSchemaMockRecorder is the mock recorder for MutableSchema.
type MutableSchemaMockRecorder struct {
	mock *MutableSchema
}

// NewMutableSchema creates a new mock instance.
func NewMutableSchema(ctrl *gomock.Controller) *MutableSchema {
	mock := &MutableSchema{ctrl: ctrl}
	mock.recorder = &MutableSchemaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MutableSchema) EXPECT() *MutableSchemaMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MutableSchema) Build() (schema.Schema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build")
	ret0, _ := ret[0].(schema.Schema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MutableSchemaMockRecorder) Build() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MutableSchema)(nil).Build))
}

// Descriptors mocks base method.
func (m *MutableSchema) Descriptors() []schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptors")
	ret0, _ := ret[0].([]schema.Descriptor)
	return ret0
}

// Descriptors indicates an expected call of Descriptors.
func (mr *MutableSchemaMockRecorder) Descriptors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptors", reflect.TypeOf((*MutableSchema)(nil).Descriptors))
}

// SetDescriptors mocks base method.
func (m *MutableSchema) SetDescriptors(arg0 []schema.MutableDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDescriptors", arg0)
}

// SetDescriptors indicates an expected call of SetDescriptors.
func (mr *MutableSchemaMockRecorder) SetDescriptors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDescriptors", reflect.TypeOf((*MutableSchema)(nil).SetDescriptors), arg0)
}

// SetVersion mocks base method.
func (m *MutableSchema) SetVersion(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVersion", arg0)
}

// SetVersion indicates an expected call of SetVersion.
func (mr *MutableSchemaMockRecorder) SetVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVersion", reflect.TypeOf((*MutableSchema)(nil).SetVersion), arg0)
}

// Type mocks base method.
func (m *MutableSchema) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MutableSchemaMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MutableSchema)(nil).Type))
}

// Version mocks base method.
func (m *MutableSchema) Version() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(int)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MutableSchemaMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MutableSchema)(nil).Version))
}

// MutableStructDescriptor is a mock of MutableStructDescriptor interface.
type MutableStructDescriptor struct {
	ctrl     *gomock.Controller
	recorder *MutableStructDescriptorMockRecorder
}

// MutableStructDescriptorMockRecorder is the mock recorder for MutableStructDescriptor.
type MutableStructDescriptorMockRecorder struct {
	mock *MutableStructDescriptor
}

// NewMutableStructDescriptor creates a new mock instance.
func NewMutableStructDescriptor(ctrl *gomock.Controller) *MutableStructDescriptor {
	mock := &MutableStructDescriptor{ctrl: ctrl}
	mock.recorder = &MutableStructDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MutableStructDescriptor) EXPECT() *MutableStructDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *MutableStructDescriptor) DescriptorType(arg0 schema.StructDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *MutableStructDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*MutableStructDescriptor)(nil).DescriptorType), arg0)
}

// Fields mocks base method.
func (m *MutableStructDescriptor) Fields() []schema.Field {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].([]schema.Field)
	return ret0
}

// Fields indicates an expected call of Fields.
func (mr *MutableStructDescriptorMockRecorder) Fields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MutableStructDescriptor)(nil).Fields))
}

// Name mocks base method.
func (m *MutableStructDescriptor) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MutableStructDescriptorMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MutableStructDescriptor)(nil).Name))
}

// Parameters mocks base method.
func (m *MutableStructDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *MutableStructDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MutableStructDescriptor)(nil).Parameters))
}

// SetFields mocks base method.
func (m *MutableStructDescriptor) SetFields(arg0 []schema.MutableField) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFields", arg0)
}

// SetFields indicates an expected call of SetFields.
func (mr *MutableStructDescriptorMockRecorder) SetFields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFields", reflect.TypeOf((*MutableStructDescriptor)(nil).SetFields), arg0)
}

// SetName mocks base method.
func (m *MutableStructDescriptor) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MutableStructDescriptorMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MutableStructDescriptor)(nil).SetName), arg0)
}

// MutableField is a mock of MutableField interface.
type MutableField struct {
	ctrl     *gomock.Controller
	recorder *MutableFieldMockRecorder
}

// MutableFieldMockRecorder is the mock recorder for MutableField.
type MutableFieldMockRecorder struct {
	mock *MutableField
}

// NewMutableField creates a new mock instance.
func NewMutableField(ctrl *gomock.Controller) *MutableField {
	mock := &MutableField{ctrl: ctrl}
	mock.recorder = &MutableFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MutableField) EXPECT() *MutableFieldMockRecorder {
	return m.recorder
}

// Descriptor mocks base method.
func (m *MutableField) Descriptor() schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptor")
	ret0, _ := ret[0].(schema.Descriptor)
	return ret0
}

// Descriptor indicates an expected call of Descriptor.
func (mr *MutableFieldMockRecorder) Descriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptor", reflect.TypeOf((*MutableField)(nil).Descriptor))
}

// DescriptorType mocks base method.
func (m *MutableField) DescriptorType(arg0 schema.Field) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *MutableFieldMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*MutableField)(nil).DescriptorType), arg0)
}

// Index mocks base method.
func (m *MutableField) Index() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index")
	ret0, _ := ret[0].(int)
	return ret0
}

// Index indicates an expected call of Index.
func (mr *MutableFieldMockRecorder) Index() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MutableField)(nil).Index))
}

// Name mocks base method.
func (m *MutableField) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MutableFieldMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MutableField)(nil).Name))
}

// SetDescriptor mocks base method.
func (m *MutableField) SetDescriptor(arg0 schema.MutableDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDescriptor", arg0)
}

// SetDescriptor indicates an expected call of SetDescriptor.
func (mr *MutableFieldMockRecorder) SetDescriptor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDescriptor", reflect.TypeOf((*MutableField)(nil).SetDescriptor), arg0)
}

// SetIndex mocks base method.
func (m *MutableField) SetIndex(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIndex", arg0)
}

// SetIndex indicates an expected call of SetIndex.
func (mr *MutableFieldMockRecorder) SetIndex(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIndex", reflect.TypeOf((*MutableField)(nil).SetIndex), arg0)
}

// SetName mocks base method.
func (m *MutableField) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MutableFieldMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MutableField)(nil).SetName), arg0)
}

// MutableMapDescriptor is a mock of MutableMapDescriptor interface.
type MutableMapDescriptor struct {
	ctrl     *gomock.Controller
	recorder *MutableMapDescriptorMockRecorder
}

// MutableMapDescriptorMockRecorder is the mock recorder for MutableMapDescriptor.
type MutableMapDescriptorMockRecorder struct {
	mock *MutableMapDescriptor
}

// NewMutableMapDescriptor creates a new mock instance.
func NewMutableMapDescriptor(ctrl *gomock.Controller) *MutableMapDescriptor {
	mock := &MutableMapDescriptor{ctrl: ctrl}
	mock.recorder = &MutableMapDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MutableMapDescriptor) EXPECT() *MutableMapDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *MutableMapDescriptor) DescriptorType(arg0 schema.MapDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *MutableMapDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*MutableMapDescriptor)(nil).DescriptorType), arg0)
}

// KeyDescriptor mocks base method.
func (m *MutableMapDescriptor) KeyDescriptor() schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyDescriptor")
	ret0, _ := ret[0].(schema.Descriptor)
	return ret0
}

// KeyDescriptor indicates an expected call of KeyDescriptor.
func (mr *MutableMapDescriptorMockRecorder) KeyDescriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyDescriptor", reflect.TypeOf((*MutableMapDescriptor)(nil).KeyDescriptor))
}

// Parameters mocks base method.
func (m *MutableMapDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *MutableMapDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MutableMapDescriptor)(nil).Parameters))
}

// SetKeyDescriptor mocks base method.
func (m *MutableMapDescriptor) SetKeyDescriptor(arg0 schema.MutableDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKeyDescriptor", arg0)
}

// SetKeyDescriptor indicates an expected call of SetKeyDescriptor.
func (mr *MutableMapDescriptorMockRecorder) SetKeyDescriptor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKeyDescriptor", reflect.TypeOf((*MutableMapDescriptor)(nil).SetKeyDescriptor), arg0)
}

// SetValueDescriptor mocks base method.
func (m *MutableMapDescriptor) SetValueDescriptor(arg0 schema.MutableDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetValueDescriptor", arg0)
}

// SetValueDescriptor indicates an expected call of SetValueDescriptor.
func (mr *MutableMapDescriptorMockRecorder) SetValueDescriptor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValueDescriptor", reflect.TypeOf((*MutableMapDescriptor)(nil).SetValueDescriptor), arg0)
}

// ValueDescriptor mocks base method.
func (m *MutableMapDescriptor) ValueDescriptor() schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueDescriptor")
	ret0, _ := ret[0].(schema.Descriptor)
	return ret0
}

// ValueDescriptor indicates an expected call of ValueDescriptor.
func (mr *MutableMapDescriptorMockRecorder) ValueDescriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueDescriptor", reflect.TypeOf((*MutableMapDescriptor)(nil).ValueDescriptor))
}

// MutableArrayDescriptor is a mock of MutableArrayDescriptor interface.
type MutableArrayDescriptor struct {
	ctrl     *gomock.Controller
	recorder *MutableArrayDescriptorMockRecorder
}

// MutableArrayDescriptorMockRecorder is the mock recorder for MutableArrayDescriptor.
type MutableArrayDescriptorMockRecorder struct {
	mock *MutableArrayDescriptor
}

// NewMutableArrayDescriptor creates a new mock instance.
func NewMutableArrayDescriptor(ctrl *gomock.Controller) *MutableArrayDescriptor {
	mock := &MutableArrayDescriptor{ctrl: ctrl}
	mock.recorder = &MutableArrayDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MutableArrayDescriptor) EXPECT() *MutableArrayDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *MutableArrayDescriptor) DescriptorType(arg0 schema.ArrayDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *MutableArrayDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*MutableArrayDescriptor)(nil).DescriptorType), arg0)
}

// Parameters mocks base method.
func (m *MutableArrayDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *MutableArrayDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MutableArrayDescriptor)(nil).Parameters))
}

// SetValueDescriptor mocks base method.
func (m *MutableArrayDescriptor) SetValueDescriptor(arg0 schema.MutableDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetValueDescriptor", arg0)
}

// SetValueDescriptor indicates an expected call of SetValueDescriptor.
func (mr *MutableArrayDescriptorMockRecorder) SetValueDescriptor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValueDescriptor", reflect.TypeOf((*MutableArrayDescriptor)(nil).SetValueDescriptor), arg0)
}

// ValueDescriptor mocks base method.
func (m *MutableArrayDescriptor) ValueDescriptor() schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueDescriptor")
	ret0, _ := ret[0].(schema.Descriptor)
	return ret0
}

// ValueDescriptor indicates an expected call of ValueDescriptor.
func (mr *MutableArrayDescriptorMockRecorder) ValueDescriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueDescriptor", reflect.TypeOf((*MutableArrayDescriptor)(nil).ValueDescriptor))
}

// MutablePrimitiveDescriptor is a mock of MutablePrimitiveDescriptor interface.
type MutablePrimitiveDescriptor struct {
	ctrl     *gomock.Controller
	recorder *MutablePrimitiveDescriptorMockRecorder
}

// MutablePrimitiveDescriptorMockRecorder is the mock recorder for MutablePrimitiveDescriptor.
type MutablePrimitiveDescriptorMockRecorder struct {
	mock *MutablePrimitiveDescriptor
}

// NewMutablePrimitiveDescriptor creates a new mock instance.
func NewMutablePrimitiveDescriptor(ctrl *gomock.Controller) *MutablePrimitiveDescriptor {
	mock := &MutablePrimitiveDescriptor{ctrl: ctrl}
	mock.recorder = &MutablePrimitiveDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MutablePrimitiveDescriptor) EXPECT() *MutablePrimitiveDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *MutablePrimitiveDescriptor) DescriptorType(arg0 schema.PrimitiveDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *MutablePrimitiveDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*MutablePrimitiveDescriptor)(nil).DescriptorType), arg0)
}

// Parameters mocks base method.
func (m *MutablePrimitiveDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *MutablePrimitiveDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MutablePrimitiveDescriptor)(nil).Parameters))
}

// SetType mocks base method.
func (m *MutablePrimitiveDescriptor) SetType(arg0 schema.PrimitiveDescriptorType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetType", arg0)
}

// SetType indicates an expected call of SetType.
func (mr *MutablePrimitiveDescriptorMockRecorder) SetType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetType", reflect.TypeOf((*MutablePrimitiveDescriptor)(nil).SetType), arg0)
}

// Type mocks base method.
func (m *MutablePrimitiveDescriptor) Type() schema.PrimitiveDescriptorType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(schema.PrimitiveDescriptorType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MutablePrimitiveDescriptorMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MutablePrimitiveDescriptor)(nil).Type))
}

// MutableEnumDescriptor is a mock of MutableEnumDescriptor interface.
type MutableEnumDescriptor struct {
	ctrl     *gomock.Controller
	recorder *MutableEnumDescriptorMockRecorder
}

// MutableEnumDescriptorMockRecorder is the mock recorder for MutableEnumDescriptor.
type MutableEnumDescriptorMockRecorder struct {
	mock *MutableEnumDescriptor
}

// NewMutableEnumDescriptor creates a new mock instance.
func NewMutableEnumDescriptor(ctrl *gomock.Controller) *MutableEnumDescriptor {
	mock := &MutableEnumDescriptor{ctrl: ctrl}
	mock.recorder = &MutableEnumDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MutableEnumDescriptor) EXPECT() *MutableEnumDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *MutableEnumDescriptor) DescriptorType(arg0 schema.EnumDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *MutableEnumDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*MutableEnumDescriptor)(nil).DescriptorType), arg0)
}

// Name mocks base method.
func (m *MutableEnumDescriptor) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MutableEnumDescriptorMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MutableEnumDescriptor)(nil).Name))
}

// Parameters mocks base method.
func (m *MutableEnumDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *MutableEnumDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MutableEnumDescriptor)(nil).Parameters))
}

// SetName mocks base method.
func (m *MutableEnumDescriptor) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MutableEnumDescriptorMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MutableEnumDescriptor)(nil).SetName), arg0)
}

// SetValueDescriptors mocks base method.
func (m *MutableEnumDescriptor) SetValueDescriptors(arg0 []schema.MutableEnumValueDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetValueDescriptors", arg0)
}

// SetValueDescriptors indicates an expected call of SetValueDescriptors.
func (mr *MutableEnumDescriptorMockRecorder) SetValueDescriptors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValueDescriptors", reflect.TypeOf((*MutableEnumDescriptor)(nil).SetValueDescriptors), arg0)
}

// ValueDescriptors mocks base method.
func (m *MutableEnumDescriptor) ValueDescriptors() []schema.EnumValueDescriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueDescriptors")
	ret0, _ := ret[0].([]schema.EnumValueDescriptor)
	return ret0
}

// ValueDescriptors indicates an expected call of ValueDescriptors.
func (mr *MutableEnumDescriptorMockRecorder) ValueDescriptors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueDescriptors", reflect.TypeOf((*MutableEnumDescriptor)(nil).ValueDescriptors))
}

// MutableEnumValueDescriptor is a mock of MutableEnumValueDescriptor interface.
type MutableEnumValueDescriptor struct {
	ctrl     *gomock.Controller
	recorder *MutableEnumValueDescriptorMockRecorder
}

// MutableEnumValueDescriptorMockRecorder is the mock recorder for MutableEnumValueDescriptor.
type MutableEnumValueDescriptorMockRecorder struct {
	mock *MutableEnumValueDescriptor
}

// NewMutableEnumValueDescriptor creates a new mock instance.
func NewMutableEnumValueDescriptor(ctrl *gomock.Controller) *MutableEnumValueDescriptor {
	mock := &MutableEnumValueDescriptor{ctrl: ctrl}
	mock.recorder = &MutableEnumValueDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MutableEnumValueDescriptor) EXPECT() *MutableEnumValueDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *MutableEnumValueDescriptor) DescriptorType(arg0 schema.EnumValueDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *MutableEnumValueDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*MutableEnumValueDescriptor)(nil).DescriptorType), arg0)
}

// Name mocks base method.
func (m *MutableEnumValueDescriptor) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MutableEnumValueDescriptorMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MutableEnumValueDescriptor)(nil).Name))
}

// Parameters mocks base method.
func (m *MutableEnumValueDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *MutableEnumValueDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MutableEnumValueDescriptor)(nil).Parameters))
}

// SetName mocks base method.
func (m *MutableEnumValueDescriptor) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MutableEnumValueDescriptorMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MutableEnumValueDescriptor)(nil).SetName), arg0)
}

// SetValue mocks base method.
func (m *MutableEnumValueDescriptor) SetValue(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetValue", arg0)
}

// SetValue indicates an expected call of SetValue.
func (mr *MutableEnumValueDescriptorMockRecorder) SetValue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValue", reflect.TypeOf((*MutableEnumValueDescriptor)(nil).SetValue), arg0)
}

// Value mocks base method.
func (m *MutableEnumValueDescriptor) Value() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(string)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MutableEnumValueDescriptorMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MutableEnumValueDescriptor)(nil).Value))
}

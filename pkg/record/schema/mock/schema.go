// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit/pkg/record/schema (interfaces: Schema,StructDescriptor,Field,MapDescriptor,ArrayDescriptor,PrimitiveDescriptor,EnumDescriptor,EnumValueDescriptor)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	schema "github.com/conduitio/conduit/pkg/record/schema"
	gomock "github.com/golang/mock/gomock"
)

// Schema is a mock of Schema interface.
type Schema struct {
	ctrl     *gomock.Controller
	recorder *SchemaMockRecorder
}

// SchemaMockRecorder is the mock recorder for Schema.
type SchemaMockRecorder struct {
	mock *Schema
}

// NewSchema creates a new mock instance.
func NewSchema(ctrl *gomock.Controller) *Schema {
	mock := &Schema{ctrl: ctrl}
	mock.recorder = &SchemaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Schema) EXPECT() *SchemaMockRecorder {
	return m.recorder
}

// Descriptors mocks base method.
func (m *Schema) Descriptors() []schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptors")
	ret0, _ := ret[0].([]schema.Descriptor)
	return ret0
}

// Descriptors indicates an expected call of Descriptors.
func (mr *SchemaMockRecorder) Descriptors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptors", reflect.TypeOf((*Schema)(nil).Descriptors))
}

// ToMutable mocks base method.
func (m *Schema) ToMutable() schema.MutableSchema {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToMutable")
	ret0, _ := ret[0].(schema.MutableSchema)
	return ret0
}

// ToMutable indicates an expected call of ToMutable.
func (mr *SchemaMockRecorder) ToMutable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToMutable", reflect.TypeOf((*Schema)(nil).ToMutable))
}

// Type mocks base method.
func (m *Schema) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *SchemaMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*Schema)(nil).Type))
}

// Version mocks base method.
func (m *Schema) Version() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(int)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *SchemaMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*Schema)(nil).Version))
}

// StructDescriptor is a mock of StructDescriptor interface.
type StructDescriptor struct {
	ctrl     *gomock.Controller
	recorder *StructDescriptorMockRecorder
}

// StructDescriptorMockRecorder is the mock recorder for StructDescriptor.
type StructDescriptorMockRecorder struct {
	mock *StructDescriptor
}

// NewStructDescriptor creates a new mock instance.
func NewStructDescriptor(ctrl *gomock.Controller) *StructDescriptor {
	mock := &StructDescriptor{ctrl: ctrl}
	mock.recorder = &StructDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *StructDescriptor) EXPECT() *StructDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *StructDescriptor) DescriptorType(arg0 schema.StructDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *StructDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*StructDescriptor)(nil).DescriptorType), arg0)
}

// Fields mocks base method.
func (m *StructDescriptor) Fields() []schema.Field {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].([]schema.Field)
	return ret0
}

// Fields indicates an expected call of Fields.
func (mr *StructDescriptorMockRecorder) Fields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*StructDescriptor)(nil).Fields))
}

// Name mocks base method.
func (m *StructDescriptor) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *StructDescriptorMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*StructDescriptor)(nil).Name))
}

// Parameters mocks base method.
func (m *StructDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *StructDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*StructDescriptor)(nil).Parameters))
}

// Field is a mock of Field interface.
type Field struct {
	ctrl     *gomock.Controller
	recorder *FieldMockRecorder
}

// FieldMockRecorder is the mock recorder for Field.
type FieldMockRecorder struct {
	mock *Field
}

// NewField creates a new mock instance.
func NewField(ctrl *gomock.Controller) *Field {
	mock := &Field{ctrl: ctrl}
	mock.recorder = &FieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Field) EXPECT() *FieldMockRecorder {
	return m.recorder
}

// Descriptor mocks base method.
func (m *Field) Descriptor() schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Descriptor")
	ret0, _ := ret[0].(schema.Descriptor)
	return ret0
}

// Descriptor indicates an expected call of Descriptor.
func (mr *FieldMockRecorder) Descriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptor", reflect.TypeOf((*Field)(nil).Descriptor))
}

// DescriptorType mocks base method.
func (m *Field) DescriptorType(arg0 schema.Field) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *FieldMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*Field)(nil).DescriptorType), arg0)
}

// Index mocks base method.
func (m *Field) Index() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index")
	ret0, _ := ret[0].(int)
	return ret0
}

// Index indicates an expected call of Index.
func (mr *FieldMockRecorder) Index() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*Field)(nil).Index))
}

// Name mocks base method.
func (m *Field) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *FieldMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*Field)(nil).Name))
}

// MapDescriptor is a mock of MapDescriptor interface.
type MapDescriptor struct {
	ctrl     *gomock.Controller
	recorder *MapDescriptorMockRecorder
}

// MapDescriptorMockRecorder is the mock recorder for MapDescriptor.
type MapDescriptorMockRecorder struct {
	mock *MapDescriptor
}

// NewMapDescriptor creates a new mock instance.
func NewMapDescriptor(ctrl *gomock.Controller) *MapDescriptor {
	mock := &MapDescriptor{ctrl: ctrl}
	mock.recorder = &MapDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MapDescriptor) EXPECT() *MapDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *MapDescriptor) DescriptorType(arg0 schema.MapDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *MapDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*MapDescriptor)(nil).DescriptorType), arg0)
}

// KeyDescriptor mocks base method.
func (m *MapDescriptor) KeyDescriptor() schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyDescriptor")
	ret0, _ := ret[0].(schema.Descriptor)
	return ret0
}

// KeyDescriptor indicates an expected call of KeyDescriptor.
func (mr *MapDescriptorMockRecorder) KeyDescriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyDescriptor", reflect.TypeOf((*MapDescriptor)(nil).KeyDescriptor))
}

// Parameters mocks base method.
func (m *MapDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *MapDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MapDescriptor)(nil).Parameters))
}

// ValueDescriptor mocks base method.
func (m *MapDescriptor) ValueDescriptor() schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueDescriptor")
	ret0, _ := ret[0].(schema.Descriptor)
	return ret0
}

// ValueDescriptor indicates an expected call of ValueDescriptor.
func (mr *MapDescriptorMockRecorder) ValueDescriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueDescriptor", reflect.TypeOf((*MapDescriptor)(nil).ValueDescriptor))
}

// ArrayDescriptor is a mock of ArrayDescriptor interface.
type ArrayDescriptor struct {
	ctrl     *gomock.Controller
	recorder *ArrayDescriptorMockRecorder
}

// ArrayDescriptorMockRecorder is the mock recorder for ArrayDescriptor.
type ArrayDescriptorMockRecorder struct {
	mock *ArrayDescriptor
}

// NewArrayDescriptor creates a new mock instance.
func NewArrayDescriptor(ctrl *gomock.Controller) *ArrayDescriptor {
	mock := &ArrayDescriptor{ctrl: ctrl}
	mock.recorder = &ArrayDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *ArrayDescriptor) EXPECT() *ArrayDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *ArrayDescriptor) DescriptorType(arg0 schema.ArrayDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *ArrayDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*ArrayDescriptor)(nil).DescriptorType), arg0)
}

// Parameters mocks base method.
func (m *ArrayDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *ArrayDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*ArrayDescriptor)(nil).Parameters))
}

// ValueDescriptor mocks base method.
func (m *ArrayDescriptor) ValueDescriptor() schema.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueDescriptor")
	ret0, _ := ret[0].(schema.Descriptor)
	return ret0
}

// ValueDescriptor indicates an expected call of ValueDescriptor.
func (mr *ArrayDescriptorMockRecorder) ValueDescriptor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueDescriptor", reflect.TypeOf((*ArrayDescriptor)(nil).ValueDescriptor))
}

// PrimitiveDescriptor is a mock of PrimitiveDescriptor interface.
type PrimitiveDescriptor struct {
	ctrl     *gomock.Controller
	recorder *PrimitiveDescriptorMockRecorder
}

// PrimitiveDescriptorMockRecorder is the mock recorder for PrimitiveDescriptor.
type PrimitiveDescriptorMockRecorder struct {
	mock *PrimitiveDescriptor
}

// NewPrimitiveDescriptor creates a new mock instance.
func NewPrimitiveDescriptor(ctrl *gomock.Controller) *PrimitiveDescriptor {
	mock := &PrimitiveDescriptor{ctrl: ctrl}
	mock.recorder = &PrimitiveDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *PrimitiveDescriptor) EXPECT() *PrimitiveDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *PrimitiveDescriptor) DescriptorType(arg0 schema.PrimitiveDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *PrimitiveDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*PrimitiveDescriptor)(nil).DescriptorType), arg0)
}

// Parameters mocks base method.
func (m *PrimitiveDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *PrimitiveDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*PrimitiveDescriptor)(nil).Parameters))
}

// Type mocks base method.
func (m *PrimitiveDescriptor) Type() schema.PrimitiveDescriptorType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(schema.PrimitiveDescriptorType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *PrimitiveDescriptorMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*PrimitiveDescriptor)(nil).Type))
}

// EnumDescriptor is a mock of EnumDescriptor interface.
type EnumDescriptor struct {
	ctrl     *gomock.Controller
	recorder *EnumDescriptorMockRecorder
}

// EnumDescriptorMockRecorder is the mock recorder for EnumDescriptor.
type EnumDescriptorMockRecorder struct {
	mock *EnumDescriptor
}

// NewEnumDescriptor creates a new mock instance.
func NewEnumDescriptor(ctrl *gomock.Controller) *EnumDescriptor {
	mock := &EnumDescriptor{ctrl: ctrl}
	mock.recorder = &EnumDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *EnumDescriptor) EXPECT() *EnumDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *EnumDescriptor) DescriptorType(arg0 schema.EnumDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *EnumDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*EnumDescriptor)(nil).DescriptorType), arg0)
}

// Name mocks base method.
func (m *EnumDescriptor) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *EnumDescriptorMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*EnumDescriptor)(nil).Name))
}

// Parameters mocks base method.
func (m *EnumDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *EnumDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*EnumDescriptor)(nil).Parameters))
}

// ValueDescriptors mocks base method.
func (m *EnumDescriptor) ValueDescriptors() []schema.EnumValueDescriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValueDescriptors")
	ret0, _ := ret[0].([]schema.EnumValueDescriptor)
	return ret0
}

// ValueDescriptors indicates an expected call of ValueDescriptors.
func (mr *EnumDescriptorMockRecorder) ValueDescriptors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValueDescriptors", reflect.TypeOf((*EnumDescriptor)(nil).ValueDescriptors))
}

// EnumValueDescriptor is a mock of EnumValueDescriptor interface.
type EnumValueDescriptor struct {
	ctrl     *gomock.Controller
	recorder *EnumValueDescriptorMockRecorder
}

// EnumValueDescriptorMockRecorder is the mock recorder for EnumValueDescriptor.
type EnumValueDescriptorMockRecorder struct {
	mock *EnumValueDescriptor
}

// NewEnumValueDescriptor creates a new mock instance.
func NewEnumValueDescriptor(ctrl *gomock.Controller) *EnumValueDescriptor {
	mock := &EnumValueDescriptor{ctrl: ctrl}
	mock.recorder = &EnumValueDescriptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *EnumValueDescriptor) EXPECT() *EnumValueDescriptorMockRecorder {
	return m.recorder
}

// DescriptorType mocks base method.
func (m *EnumValueDescriptor) DescriptorType(arg0 schema.EnumValueDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescriptorType", arg0)
}

// DescriptorType indicates an expected call of DescriptorType.
func (mr *EnumValueDescriptorMockRecorder) DescriptorType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescriptorType", reflect.TypeOf((*EnumValueDescriptor)(nil).DescriptorType), arg0)
}

// Name mocks base method.
func (m *EnumValueDescriptor) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *EnumValueDescriptorMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*EnumValueDescriptor)(nil).Name))
}

// Parameters mocks base method.
func (m *EnumValueDescriptor) Parameters() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Parameters indicates an expected call of Parameters.
func (mr *EnumValueDescriptorMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*EnumValueDescriptor)(nil).Parameters))
}

// Value mocks base method.
func (m *EnumValueDescriptor) Value() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(string)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *EnumValueDescriptorMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*EnumValueDescriptor)(nil).Value))
}

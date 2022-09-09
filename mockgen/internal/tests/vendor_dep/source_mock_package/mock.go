// Code generated by MockGen. DO NOT EDIT.
// Source: vendor_dep.go

// Package mock_vendor_dep is a generated GoMock package.
package mock_vendor_dep

import (
	reflect "reflect"

	gomock "github.com/uberbrodt/mock/gomock"
	present "golang.org/x/tools/present"
)

// MockVendorsDep is a mock of VendorsDep interface.
type MockVendorsDep struct {
	ctrl     *gomock.Controller
	recorder *MockVendorsDepMockRecorder
}

// MockVendorsDepMockRecorder is the mock recorder for MockVendorsDep.
type MockVendorsDepMockRecorder struct {
	mock *MockVendorsDep
}

// NewMockVendorsDep creates a new mock instance.
func NewMockVendorsDep(ctrl *gomock.Controller) *MockVendorsDep {
	mock := &MockVendorsDep{ctrl: ctrl}
	mock.recorder = &MockVendorsDepMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVendorsDep) EXPECT() *MockVendorsDepMockRecorder {
	return m.recorder
}

// Foo mocks base method.
func (m *MockVendorsDep) Foo() present.Elem {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Foo")
	ret0, _ := ret[0].(present.Elem)
	return ret0
}

// Foo indicates an expected call of Foo.
func (mr *MockVendorsDepMockRecorder) Foo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockVendorsDep)(nil).Foo))
}

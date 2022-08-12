// Code generated by MockGen. DO NOT EDIT.
// Source: external.go

// Package source is a generated GoMock package.
package source

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	generics "github.com/golang/mock/mockgen/internal/tests/generics"
	other "github.com/golang/mock/mockgen/internal/tests/generics/other"
	constraints "golang.org/x/exp/constraints"
)

// MockExternalConstraint is a mock of ExternalConstraint interface.
type MockExternalConstraint[I constraints.Integer, F constraints.Float] struct {
	ctrl     *gomock.Controller
	recorder *MockExternalConstraintMockRecorder[I, F]
}

// MockExternalConstraintMockRecorder is the mock recorder for MockExternalConstraint.
type MockExternalConstraintMockRecorder[I constraints.Integer, F constraints.Float] struct {
	mock *MockExternalConstraint[I, F]
}

// NewMockExternalConstraint creates a new mock instance.
func NewMockExternalConstraint[I constraints.Integer, F constraints.Float](ctrl *gomock.Controller) *MockExternalConstraint[I, F] {
	mock := &MockExternalConstraint[I, F]{ctrl: ctrl}
	mock.recorder = &MockExternalConstraintMockRecorder[I, F]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExternalConstraint[I, F]) EXPECT() *MockExternalConstraintMockRecorder[I, F] {
	return m.recorder
}

// Eight mocks base method.
func (m *MockExternalConstraint[I, F]) Eight(arg0 F) other.Two[I, F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eight", arg0)
	ret0, _ := ret[0].(other.Two[I, F])
	return ret0
}

// Eight indicates an expected call of Eight.
func (mr *MockExternalConstraintMockRecorder[I, F]) Eight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eight", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Eight), arg0)
}

// Five mocks base method.
func (m *MockExternalConstraint[I, F]) Five(arg0 I) generics.Baz[F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Five", arg0)
	ret0, _ := ret[0].(generics.Baz[F])
	return ret0
}

// Five indicates an expected call of Five.
func (mr *MockExternalConstraintMockRecorder[I, F]) Five(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Five", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Five), arg0)
}

// Four mocks base method.
func (m *MockExternalConstraint[I, F]) Four(arg0 I) generics.Foo[I, F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Four", arg0)
	ret0, _ := ret[0].(generics.Foo[I, F])
	return ret0
}

// Four indicates an expected call of Four.
func (mr *MockExternalConstraintMockRecorder[I, F]) Four(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Four", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Four), arg0)
}

// Nine mocks base method.
func (m *MockExternalConstraint[I, F]) Nine(arg0 generics.Iface[I]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Nine", arg0)
}

// Nine indicates an expected call of Nine.
func (mr *MockExternalConstraintMockRecorder[I, F]) Nine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nine", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Nine), arg0)
}

// One mocks base method.
func (m *MockExternalConstraint[I, F]) One(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// One indicates an expected call of One.
func (mr *MockExternalConstraintMockRecorder[I, F]) One(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).One), arg0)
}

// Seven mocks base method.
func (m *MockExternalConstraint[I, F]) Seven(arg0 I) other.One[I] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seven", arg0)
	ret0, _ := ret[0].(other.One[I])
	return ret0
}

// Seven indicates an expected call of Seven.
func (mr *MockExternalConstraintMockRecorder[I, F]) Seven(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seven", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Seven), arg0)
}

// Six mocks base method.
func (m *MockExternalConstraint[I, F]) Six(arg0 I) *generics.Baz[F] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Six", arg0)
	ret0, _ := ret[0].(*generics.Baz[F])
	return ret0
}

// Six indicates an expected call of Six.
func (mr *MockExternalConstraintMockRecorder[I, F]) Six(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Six", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Six), arg0)
}

// Ten mocks base method.
func (m *MockExternalConstraint[I, F]) Ten(arg0 *I) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ten", arg0)
}

// Ten indicates an expected call of Ten.
func (mr *MockExternalConstraintMockRecorder[I, F]) Ten(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ten", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Ten), arg0)
}

// Three mocks base method.
func (m *MockExternalConstraint[I, F]) Three(arg0 I) F {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Three", arg0)
	ret0, _ := ret[0].(F)
	return ret0
}

// Three indicates an expected call of Three.
func (mr *MockExternalConstraintMockRecorder[I, F]) Three(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Three", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Three), arg0)
}

// Two mocks base method.
func (m *MockExternalConstraint[I, F]) Two(arg0 I) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Two", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Two indicates an expected call of Two.
func (mr *MockExternalConstraintMockRecorder[I, F]) Two(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Two", reflect.TypeOf((*MockExternalConstraint[I, F])(nil).Two), arg0)
}

// MockEmbeddingIface is a mock of EmbeddingIface interface.
type MockEmbeddingIface[T constraints.Integer, R constraints.Float] struct {
	ctrl     *gomock.Controller
	recorder *MockEmbeddingIfaceMockRecorder[T, R]
}

// MockEmbeddingIfaceMockRecorder is the mock recorder for MockEmbeddingIface.
type MockEmbeddingIfaceMockRecorder[T constraints.Integer, R constraints.Float] struct {
	mock *MockEmbeddingIface[T, R]
}

// NewMockEmbeddingIface creates a new mock instance.
func NewMockEmbeddingIface[T constraints.Integer, R constraints.Float](ctrl *gomock.Controller) *MockEmbeddingIface[T, R] {
	mock := &MockEmbeddingIface[T, R]{ctrl: ctrl}
	mock.recorder = &MockEmbeddingIfaceMockRecorder[T, R]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbeddingIface[T, R]) EXPECT() *MockEmbeddingIfaceMockRecorder[T, R] {
	return m.recorder
}

// Eight mocks base method.
func (m *MockEmbeddingIface[T, R]) Eight(arg0 R) other.Two[T, R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eight", arg0)
	ret0, _ := ret[0].(other.Two[T, R])
	return ret0
}

// Eight indicates an expected call of Eight.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Eight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eight", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Eight), arg0)
}

// Five mocks base method.
func (m *MockEmbeddingIface[T, R]) Five(arg0 T) generics.Baz[R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Five", arg0)
	ret0, _ := ret[0].(generics.Baz[R])
	return ret0
}

// Five indicates an expected call of Five.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Five(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Five", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Five), arg0)
}

// Foo mocks base method.
func (m *MockEmbeddingIface[T, R]) Foo() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Foo")
	ret0, _ := ret[0].(error)
	return ret0
}

// Foo indicates an expected call of Foo.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Foo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Foo", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Foo))
}

// Four mocks base method.
func (m *MockEmbeddingIface[T, R]) Four(arg0 T) generics.Foo[T, R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Four", arg0)
	ret0, _ := ret[0].(generics.Foo[T, R])
	return ret0
}

// Four indicates an expected call of Four.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Four(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Four", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Four), arg0)
}

// Nine mocks base method.
func (m *MockEmbeddingIface[T, R]) Nine(arg0 generics.Iface[T]) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Nine", arg0)
}

// Nine indicates an expected call of Nine.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Nine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nine", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Nine), arg0)
}

// One mocks base method.
func (m *MockEmbeddingIface[T, R]) One(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// One indicates an expected call of One.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) One(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).One), arg0)
}

// Seven mocks base method.
func (m *MockEmbeddingIface[T, R]) Seven(arg0 T) other.One[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seven", arg0)
	ret0, _ := ret[0].(other.One[T])
	return ret0
}

// Seven indicates an expected call of Seven.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Seven(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seven", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Seven), arg0)
}

// Six mocks base method.
func (m *MockEmbeddingIface[T, R]) Six(arg0 T) *generics.Baz[R] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Six", arg0)
	ret0, _ := ret[0].(*generics.Baz[R])
	return ret0
}

// Six indicates an expected call of Six.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Six(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Six", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Six), arg0)
}

// Ten mocks base method.
func (m *MockEmbeddingIface[T, R]) Ten(arg0 *T) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Ten", arg0)
}

// Ten indicates an expected call of Ten.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Ten(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ten", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Ten), arg0)
}

// Three mocks base method.
func (m *MockEmbeddingIface[T, R]) Three(arg0 T) R {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Three", arg0)
	ret0, _ := ret[0].(R)
	return ret0
}

// Three indicates an expected call of Three.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Three(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Three", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Three), arg0)
}

// Twenty mocks base method.
func (m *MockEmbeddingIface[T, R]) Twenty(arg0 generics.StructType, arg1 T) (R, other.Five) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Twenty", arg0, arg1)
	ret0, _ := ret[0].(R)
	ret1, _ := ret[1].(other.Five)
	return ret0, ret1
}

// Twenty indicates an expected call of Twenty.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Twenty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Twenty", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Twenty), arg0, arg1)
}

// TwentyFour mocks base method.
func (m *MockEmbeddingIface[T, R]) TwentyFour() other.StructType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwentyFour")
	ret0, _ := ret[0].(other.StructType)
	return ret0
}

// TwentyFour indicates an expected call of TwentyFour.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) TwentyFour() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwentyFour", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).TwentyFour))
}

// TwentyThree mocks base method.
func (m *MockEmbeddingIface[T, R]) TwentyThree(arg0 generics.TwentyTwo[R], arg1 generics.TwentyTwo[T]) other.StructType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwentyThree", arg0, arg1)
	ret0, _ := ret[0].(other.StructType)
	return ret0
}

// TwentyThree indicates an expected call of TwentyThree.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) TwentyThree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwentyThree", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).TwentyThree), arg0, arg1)
}

// TwentyTwo mocks base method.
func (m *MockEmbeddingIface[T, R]) TwentyTwo() generics.StructType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwentyTwo")
	ret0, _ := ret[0].(generics.StructType)
	return ret0
}

// TwentyTwo indicates an expected call of TwentyTwo.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) TwentyTwo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwentyTwo", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).TwentyTwo))
}

// Two mocks base method.
func (m *MockEmbeddingIface[T, R]) Two(arg0 T) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Two", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Two indicates an expected call of Two.
func (mr *MockEmbeddingIfaceMockRecorder[T, R]) Two(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Two", reflect.TypeOf((*MockEmbeddingIface[T, R])(nil).Two), arg0)
}

// MockTwentyOne is a mock of TwentyOne interface.
type MockTwentyOne[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockTwentyOneMockRecorder[T]
}

// MockTwentyOneMockRecorder is the mock recorder for MockTwentyOne.
type MockTwentyOneMockRecorder[T any] struct {
	mock *MockTwentyOne[T]
}

// NewMockTwentyOne creates a new mock instance.
func NewMockTwentyOne[T any](ctrl *gomock.Controller) *MockTwentyOne[T] {
	mock := &MockTwentyOne[T]{ctrl: ctrl}
	mock.recorder = &MockTwentyOneMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTwentyOne[T]) EXPECT() *MockTwentyOneMockRecorder[T] {
	return m.recorder
}

// TwentyOne mocks base method.
func (m *MockTwentyOne[T]) TwentyOne() T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwentyOne")
	ret0, _ := ret[0].(T)
	return ret0
}

// TwentyOne indicates an expected call of TwentyOne.
func (mr *MockTwentyOneMockRecorder[T]) TwentyOne() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwentyOne", reflect.TypeOf((*MockTwentyOne[T])(nil).TwentyOne))
}

// MockTwentyFour is a mock of TwentyFour interface.
type MockTwentyFour[T other.StructType] struct {
	ctrl     *gomock.Controller
	recorder *MockTwentyFourMockRecorder[T]
}

// MockTwentyFourMockRecorder is the mock recorder for MockTwentyFour.
type MockTwentyFourMockRecorder[T other.StructType] struct {
	mock *MockTwentyFour[T]
}

// NewMockTwentyFour creates a new mock instance.
func NewMockTwentyFour[T other.StructType](ctrl *gomock.Controller) *MockTwentyFour[T] {
	mock := &MockTwentyFour[T]{ctrl: ctrl}
	mock.recorder = &MockTwentyFourMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTwentyFour[T]) EXPECT() *MockTwentyFourMockRecorder[T] {
	return m.recorder
}

// TwentyFour mocks base method.
func (m *MockTwentyFour[T]) TwentyFour() T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TwentyFour")
	ret0, _ := ret[0].(T)
	return ret0
}

// TwentyFour indicates an expected call of TwentyFour.
func (mr *MockTwentyFourMockRecorder[T]) TwentyFour() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TwentyFour", reflect.TypeOf((*MockTwentyFour[T])(nil).TwentyFour))
}

// MockClonable is a mock of Clonable interface.
type MockClonable[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockClonableMockRecorder[T]
}

// MockClonableMockRecorder is the mock recorder for MockClonable.
type MockClonableMockRecorder[T any] struct {
	mock *MockClonable[T]
}

// NewMockClonable creates a new mock instance.
func NewMockClonable[T any](ctrl *gomock.Controller) *MockClonable[T] {
	mock := &MockClonable[T]{ctrl: ctrl}
	mock.recorder = &MockClonableMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClonable[T]) EXPECT() *MockClonableMockRecorder[T] {
	return m.recorder
}

// Clone mocks base method.
func (m *MockClonable[T]) Clone() T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(T)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockClonableMockRecorder[T]) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockClonable[T])(nil).Clone))
}

// MockFinder is a mock of Finder interface.
type MockFinder[T generics.Clonable[T]] struct {
	ctrl     *gomock.Controller
	recorder *MockFinderMockRecorder[T]
}

// MockFinderMockRecorder is the mock recorder for MockFinder.
type MockFinderMockRecorder[T generics.Clonable[T]] struct {
	mock *MockFinder[T]
}

// NewMockFinder creates a new mock instance.
func NewMockFinder[T generics.Clonable[T]](ctrl *gomock.Controller) *MockFinder[T] {
	mock := &MockFinder[T]{ctrl: ctrl}
	mock.recorder = &MockFinderMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFinder[T]) EXPECT() *MockFinderMockRecorder[T] {
	return m.recorder
}

// Find mocks base method.
func (m *MockFinder[T]) Find(ctx context.Context) ([]T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx)
	ret0, _ := ret[0].([]T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockFinderMockRecorder[T]) Find(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFinder[T])(nil).Find), ctx)
}

// MockUpdateNotifier is a mock of UpdateNotifier interface.
type MockUpdateNotifier[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateNotifierMockRecorder[T]
}

// MockUpdateNotifierMockRecorder is the mock recorder for MockUpdateNotifier.
type MockUpdateNotifierMockRecorder[T any] struct {
	mock *MockUpdateNotifier[T]
}

// NewMockUpdateNotifier creates a new mock instance.
func NewMockUpdateNotifier[T any](ctrl *gomock.Controller) *MockUpdateNotifier[T] {
	mock := &MockUpdateNotifier[T]{ctrl: ctrl}
	mock.recorder = &MockUpdateNotifierMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateNotifier[T]) EXPECT() *MockUpdateNotifierMockRecorder[T] {
	return m.recorder
}

// NotifyC mocks base method.
func (m *MockUpdateNotifier[T]) NotifyC(ctx context.Context) <-chan []T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyC", ctx)
	ret0, _ := ret[0].(<-chan []T)
	return ret0
}

// NotifyC indicates an expected call of NotifyC.
func (mr *MockUpdateNotifierMockRecorder[T]) NotifyC(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyC", reflect.TypeOf((*MockUpdateNotifier[T])(nil).NotifyC), ctx)
}

// Refresh mocks base method.
func (m *MockUpdateNotifier[T]) Refresh(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Refresh", ctx)
}

// Refresh indicates an expected call of Refresh.
func (mr *MockUpdateNotifierMockRecorder[T]) Refresh(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockUpdateNotifier[T])(nil).Refresh), ctx)
}

// MockEmbeddedW is a mock of EmbeddedW interface.
type MockEmbeddedW[W generics.StructType] struct {
	ctrl     *gomock.Controller
	recorder *MockEmbeddedWMockRecorder[W]
}

// MockEmbeddedWMockRecorder is the mock recorder for MockEmbeddedW.
type MockEmbeddedWMockRecorder[W generics.StructType] struct {
	mock *MockEmbeddedW[W]
}

// NewMockEmbeddedW creates a new mock instance.
func NewMockEmbeddedW[W generics.StructType](ctrl *gomock.Controller) *MockEmbeddedW[W] {
	mock := &MockEmbeddedW[W]{ctrl: ctrl}
	mock.recorder = &MockEmbeddedWMockRecorder[W]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbeddedW[W]) EXPECT() *MockEmbeddedWMockRecorder[W] {
	return m.recorder
}

// EmbeddedZ mocks base method.
func (m *MockEmbeddedW[W]) EmbeddedZ(arg0 W) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EmbeddedZ", arg0)
}

// EmbeddedZ indicates an expected call of EmbeddedZ.
func (mr *MockEmbeddedWMockRecorder[W]) EmbeddedZ(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmbeddedZ", reflect.TypeOf((*MockEmbeddedW[W])(nil).EmbeddedZ), arg0)
}

// MockEmbeddedX is a mock of EmbeddedX interface.
type MockEmbeddedX[X generics.StructType] struct {
	ctrl     *gomock.Controller
	recorder *MockEmbeddedXMockRecorder[X]
}

// MockEmbeddedXMockRecorder is the mock recorder for MockEmbeddedX.
type MockEmbeddedXMockRecorder[X generics.StructType] struct {
	mock *MockEmbeddedX[X]
}

// NewMockEmbeddedX creates a new mock instance.
func NewMockEmbeddedX[X generics.StructType](ctrl *gomock.Controller) *MockEmbeddedX[X] {
	mock := &MockEmbeddedX[X]{ctrl: ctrl}
	mock.recorder = &MockEmbeddedXMockRecorder[X]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbeddedX[X]) EXPECT() *MockEmbeddedXMockRecorder[X] {
	return m.recorder
}

// EmbeddedZ mocks base method.
func (m *MockEmbeddedX[X]) EmbeddedZ(arg0 X) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EmbeddedZ", arg0)
}

// EmbeddedZ indicates an expected call of EmbeddedZ.
func (mr *MockEmbeddedXMockRecorder[X]) EmbeddedZ(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmbeddedZ", reflect.TypeOf((*MockEmbeddedX[X])(nil).EmbeddedZ), arg0)
}

// MockEmbeddedY is a mock of EmbeddedY interface.
type MockEmbeddedY[Y generics.StructType] struct {
	ctrl     *gomock.Controller
	recorder *MockEmbeddedYMockRecorder[Y]
}

// MockEmbeddedYMockRecorder is the mock recorder for MockEmbeddedY.
type MockEmbeddedYMockRecorder[Y generics.StructType] struct {
	mock *MockEmbeddedY[Y]
}

// NewMockEmbeddedY creates a new mock instance.
func NewMockEmbeddedY[Y generics.StructType](ctrl *gomock.Controller) *MockEmbeddedY[Y] {
	mock := &MockEmbeddedY[Y]{ctrl: ctrl}
	mock.recorder = &MockEmbeddedYMockRecorder[Y]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbeddedY[Y]) EXPECT() *MockEmbeddedYMockRecorder[Y] {
	return m.recorder
}

// EmbeddedZ mocks base method.
func (m *MockEmbeddedY[Y]) EmbeddedZ(arg0 Y) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EmbeddedZ", arg0)
}

// EmbeddedZ indicates an expected call of EmbeddedZ.
func (mr *MockEmbeddedYMockRecorder[Y]) EmbeddedZ(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmbeddedZ", reflect.TypeOf((*MockEmbeddedY[Y])(nil).EmbeddedZ), arg0)
}

// MockEmbeddedZ is a mock of EmbeddedZ interface.
type MockEmbeddedZ[Z any] struct {
	ctrl     *gomock.Controller
	recorder *MockEmbeddedZMockRecorder[Z]
}

// MockEmbeddedZMockRecorder is the mock recorder for MockEmbeddedZ.
type MockEmbeddedZMockRecorder[Z any] struct {
	mock *MockEmbeddedZ[Z]
}

// NewMockEmbeddedZ creates a new mock instance.
func NewMockEmbeddedZ[Z any](ctrl *gomock.Controller) *MockEmbeddedZ[Z] {
	mock := &MockEmbeddedZ[Z]{ctrl: ctrl}
	mock.recorder = &MockEmbeddedZMockRecorder[Z]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbeddedZ[Z]) EXPECT() *MockEmbeddedZMockRecorder[Z] {
	return m.recorder
}

// EmbeddedZ mocks base method.
func (m *MockEmbeddedZ[Z]) EmbeddedZ(arg0 Z) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EmbeddedZ", arg0)
}

// EmbeddedZ indicates an expected call of EmbeddedZ.
func (mr *MockEmbeddedZMockRecorder[Z]) EmbeddedZ(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmbeddedZ", reflect.TypeOf((*MockEmbeddedZ[Z])(nil).EmbeddedZ), arg0)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: foo.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFoo is a mock of Foo interface
type MockFoo struct {
	ctrl     *gomock.Controller
	recorder *MockFooMockRecorder
}

// MockFooMockRecorder is the mock recorder for MockFoo
type MockFooMockRecorder struct {
	mock *MockFoo
}

// NewMockFoo creates a new mock instance
func NewMockFoo(ctrl *gomock.Controller) *MockFoo {
	mock := &MockFoo{ctrl: ctrl}
	mock.recorder = &MockFooMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFoo) EXPECT() *MockFooMockRecorder {
	return m.recorder
}

// Max mocks base method
func (m *MockFoo) Max(a, b int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Max", a, b)
	ret0, _ := ret[0].(int)
	return ret0
}

// Max indicates an expected call of Max
func (mr *MockFooMockRecorder) Max(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Max", reflect.TypeOf((*MockFoo)(nil).Max), a, b)
}

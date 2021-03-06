// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/morvanzhou/unittest-demo/5mock (interfaces: YourInterface)

// Package _mock is a generated GoMock package.
package _mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockYourInterface is a mock of YourInterface interface
type MockYourInterface struct {
	ctrl     *gomock.Controller
	recorder *MockYourInterfaceMockRecorder
}

// MockYourInterfaceMockRecorder is the mock recorder for MockYourInterface
type MockYourInterfaceMockRecorder struct {
	mock *MockYourInterface
}

// NewMockYourInterface creates a new mock instance
func NewMockYourInterface(ctrl *gomock.Controller) *MockYourInterface {
	mock := &MockYourInterface{ctrl: ctrl}
	mock.recorder = &MockYourInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockYourInterface) EXPECT() *MockYourInterfaceMockRecorder {
	return m.recorder
}

// YourInterFunc mocks base method
func (m *MockYourInterface) YourInterFunc(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "YourInterFunc", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// YourInterFunc indicates an expected call of YourInterFunc
func (mr *MockYourInterfaceMockRecorder) YourInterFunc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "YourInterFunc", reflect.TypeOf((*MockYourInterface)(nil).YourInterFunc), arg0)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/base-org/pessimism/internal/client/alert_client (interfaces: AlertClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	alert_client "github.com/base-org/pessimism/internal/client"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAlertClient is a mock of AlertClient interface.
type MockAlertClient struct {
	ctrl     *gomock.Controller
	recorder *MockAlertClientMockRecorder
}

// MockAlertClientMockRecorder is the mock recorder for MockAlertClient.
type MockAlertClientMockRecorder struct {
	mock *MockAlertClient
}

// NewMockAlertClient creates a new mock instance.
func NewMockAlertClient(ctrl *gomock.Controller) *MockAlertClient {
	mock := &MockAlertClient{ctrl: ctrl}
	mock.recorder = &MockAlertClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertClient) EXPECT() *MockAlertClientMockRecorder {
	return m.recorder
}

// PostEvent mocks base method.
func (m *MockAlertClient) PostEvent(arg0 context.Context, arg1 *alert_client.AlertEventTrigger) (*alert_client.AlertAPIResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostEvent", arg0, arg1)
	ret0, _ := ret[0].(*alert_client.AlertAPIResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostEvent indicates an expected call of PostEvent.
func (mr *MockAlertClientMockRecorder) PostEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostEvent", reflect.TypeOf((*MockAlertClient)(nil).PostEvent), arg0, arg1)
}

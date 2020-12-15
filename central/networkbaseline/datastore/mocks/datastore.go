// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// GetNetworkBaseline mocks base method
func (m *MockDataStore) GetNetworkBaseline(ctx context.Context, deploymentID string) (*storage.NetworkBaseline, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkBaseline", ctx, deploymentID)
	ret0, _ := ret[0].(*storage.NetworkBaseline)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNetworkBaseline indicates an expected call of GetNetworkBaseline
func (mr *MockDataStoreMockRecorder) GetNetworkBaseline(ctx, deploymentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkBaseline", reflect.TypeOf((*MockDataStore)(nil).GetNetworkBaseline), ctx, deploymentID)
}

// Walk mocks base method
func (m *MockDataStore) Walk(ctx context.Context, f func(*storage.NetworkBaseline) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk
func (mr *MockDataStoreMockRecorder) Walk(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockDataStore)(nil).Walk), ctx, f)
}

// UpsertNetworkBaselines mocks base method
func (m *MockDataStore) UpsertNetworkBaselines(ctx context.Context, baselines []*storage.NetworkBaseline) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertNetworkBaselines", ctx, baselines)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertNetworkBaselines indicates an expected call of UpsertNetworkBaselines
func (mr *MockDataStoreMockRecorder) UpsertNetworkBaselines(ctx, baselines interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertNetworkBaselines", reflect.TypeOf((*MockDataStore)(nil).UpsertNetworkBaselines), ctx, baselines)
}

// DeleteNetworkBaseline mocks base method
func (m *MockDataStore) DeleteNetworkBaseline(ctx context.Context, deploymentID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetworkBaseline", ctx, deploymentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetworkBaseline indicates an expected call of DeleteNetworkBaseline
func (mr *MockDataStoreMockRecorder) DeleteNetworkBaseline(ctx, deploymentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetworkBaseline", reflect.TypeOf((*MockDataStore)(nil).DeleteNetworkBaseline), ctx, deploymentID)
}

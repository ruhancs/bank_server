// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/account/domain/gateway/account_repository.go

// Package mock_gateway_account is a generated GoMock package.
package mock_gateway_account

import (
	account_entity "bank_server/internal/account/domain/entity"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountRepositoryInterface is a mock of AccountRepositoryInterface interface.
type MockAccountRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryInterfaceMockRecorder
}

// MockAccountRepositoryInterfaceMockRecorder is the mock recorder for MockAccountRepositoryInterface.
type MockAccountRepositoryInterfaceMockRecorder struct {
	mock *MockAccountRepositoryInterface
}

// NewMockAccountRepositoryInterface creates a new mock instance.
func NewMockAccountRepositoryInterface(ctrl *gomock.Controller) *MockAccountRepositoryInterface {
	mock := &MockAccountRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepositoryInterface) EXPECT() *MockAccountRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountRepositoryInterface) Create(ctx context.Context, account *account_entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccountRepositoryInterfaceMockRecorder) Create(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).Create), ctx, account)
}

// Delete mocks base method.
func (m *MockAccountRepositoryInterface) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAccountRepositoryInterfaceMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockAccountRepositoryInterface) Get(ctx context.Context, id string) (*account_entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*account_entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAccountRepositoryInterfaceMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).Get), ctx, id)
}

// GetToUpdate mocks base method.
func (m *MockAccountRepositoryInterface) GetToUpdate(ctx context.Context, id string) (*account_entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToUpdate", ctx, id)
	ret0, _ := ret[0].(*account_entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToUpdate indicates an expected call of GetToUpdate.
func (mr *MockAccountRepositoryInterfaceMockRecorder) GetToUpdate(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToUpdate", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).GetToUpdate), ctx, id)
}

// UpdateBalance mocks base method.
func (m *MockAccountRepositoryInterface) UpdateBalance(ctx context.Context, id string, balance int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", ctx, id, balance)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockAccountRepositoryInterfaceMockRecorder) UpdateBalance(ctx, id, balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).UpdateBalance), ctx, id, balance)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\augusto.lourencatto\go\src\github.com\augusto\full_cycle_esquenta_go\entity\repository.go

// Package mock_entity is a generated GoMock package.
package mock_entity

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockTransactionRepository) Insert(id, accountId string, amount float64, status, errorMessage string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", id, accountId, amount, status, errorMessage)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockTransactionRepositoryMockRecorder) Insert(id, accountId, amount, status, errorMessage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTransactionRepository)(nil).Insert), id, accountId, amount, status, errorMessage)
}

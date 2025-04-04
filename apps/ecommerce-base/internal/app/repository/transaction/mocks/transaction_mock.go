// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/app/repository/transaction/transaction.go

// Package mock_transaction is a generated GoMock package.
package mock_transaction

import (
	context "context"
	sql "database/sql"
	entity "ecommerce/internal/entity"
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

// BeginTx mocks base method.
func (m *MockTransactionRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockTransactionRepositoryMockRecorder) BeginTx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockTransactionRepository)(nil).BeginTx), ctx)
}

// CommitTx mocks base method.
func (m *MockTransactionRepository) CommitTx(ctx context.Context, tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTx", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitTx indicates an expected call of CommitTx.
func (mr *MockTransactionRepositoryMockRecorder) CommitTx(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTx", reflect.TypeOf((*MockTransactionRepository)(nil).CommitTx), ctx, tx)
}

// CreateOrder mocks base method.
func (m *MockTransactionRepository) CreateOrder(ctx context.Context, tx *sql.Tx, request entity.OrderRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, tx, request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockTransactionRepositoryMockRecorder) CreateOrder(ctx, tx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockTransactionRepository)(nil).CreateOrder), ctx, tx, request)
}

// CreateOrderItem mocks base method.
func (m *MockTransactionRepository) CreateOrderItem(ctx context.Context, tx *sql.Tx, request entity.OrderRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderItem", ctx, tx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrderItem indicates an expected call of CreateOrderItem.
func (mr *MockTransactionRepositoryMockRecorder) CreateOrderItem(ctx, tx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderItem", reflect.TypeOf((*MockTransactionRepository)(nil).CreateOrderItem), ctx, tx, request)
}

// GetMultipleSku mocks base method.
func (m *MockTransactionRepository) GetMultipleSku(ctx context.Context, skuIds []entity.SkuRequest) ([]entity.WarehouseStock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultipleSku", ctx, skuIds)
	ret0, _ := ret[0].([]entity.WarehouseStock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultipleSku indicates an expected call of GetMultipleSku.
func (mr *MockTransactionRepositoryMockRecorder) GetMultipleSku(ctx, skuIds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultipleSku", reflect.TypeOf((*MockTransactionRepository)(nil).GetMultipleSku), ctx, skuIds)
}

// GetOrderById mocks base method.
func (m *MockTransactionRepository) GetOrderById(ctx context.Context, id string) (entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", ctx, id)
	ret0, _ := ret[0].(entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockTransactionRepositoryMockRecorder) GetOrderById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockTransactionRepository)(nil).GetOrderById), ctx, id)
}

// ReverseStock mocks base method.
func (m *MockTransactionRepository) ReverseStock(ctx context.Context, tx *sql.Tx, warehouseId string, sku entity.SkuRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReverseStock", ctx, tx, warehouseId, sku)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReverseStock indicates an expected call of ReverseStock.
func (mr *MockTransactionRepositoryMockRecorder) ReverseStock(ctx, tx, warehouseId, sku interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseStock", reflect.TypeOf((*MockTransactionRepository)(nil).ReverseStock), ctx, tx, warehouseId, sku)
}

// RollbackTx mocks base method.
func (m *MockTransactionRepository) RollbackTx(ctx context.Context, tx *sql.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTx", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackTx indicates an expected call of RollbackTx.
func (mr *MockTransactionRepositoryMockRecorder) RollbackTx(ctx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTx", reflect.TypeOf((*MockTransactionRepository)(nil).RollbackTx), ctx, tx)
}

// UpdateOrderStatus mocks base method.
func (m *MockTransactionRepository) UpdateOrderStatus(ctx context.Context, tx *sql.Tx, id, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderStatus", ctx, tx, id, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderStatus indicates an expected call of UpdateOrderStatus.
func (mr *MockTransactionRepositoryMockRecorder) UpdateOrderStatus(ctx, tx, id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderStatus", reflect.TypeOf((*MockTransactionRepository)(nil).UpdateOrderStatus), ctx, tx, id, status)
}

// UpdateStock mocks base method.
func (m *MockTransactionRepository) UpdateStock(ctx context.Context, tx *sql.Tx, request entity.SkuRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStock", ctx, tx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStock indicates an expected call of UpdateStock.
func (mr *MockTransactionRepositoryMockRecorder) UpdateStock(ctx, tx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStock", reflect.TypeOf((*MockTransactionRepository)(nil).UpdateStock), ctx, tx, request)
}

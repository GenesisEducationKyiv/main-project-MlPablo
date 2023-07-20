// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_currency is a generated GoMock package.
package mock_currency

import (
	context "context"
	rate "currency/internal/domain/rate"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICurrencyService is a mock of ICurrencyService interface.
type MockICurrencyService struct {
	ctrl     *gomock.Controller
	recorder *MockICurrencyServiceMockRecorder
}

// MockICurrencyServiceMockRecorder is the mock recorder for MockICurrencyService.
type MockICurrencyServiceMockRecorder struct {
	mock *MockICurrencyService
}

// NewMockICurrencyService creates a new mock instance.
func NewMockICurrencyService(ctrl *gomock.Controller) *MockICurrencyService {
	mock := &MockICurrencyService{ctrl: ctrl}
	mock.recorder = &MockICurrencyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICurrencyService) EXPECT() *MockICurrencyServiceMockRecorder {
	return m.recorder
}

// GetCurrency mocks base method.
func (m *MockICurrencyService) GetCurrency(ctx context.Context, c *rate.Rate) (*rate.Currency, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrency", ctx, c)
	ret0, _ := ret[0].(*rate.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrency indicates an expected call of GetCurrency.
func (mr *MockICurrencyServiceMockRecorder) GetCurrency(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrency", reflect.TypeOf((*MockICurrencyService)(nil).GetCurrency), ctx, c)
}
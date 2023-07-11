// Code generated by MockGen. DO NOT EDIT.
// Source: chain.go

// Package mock_currency is a generated GoMock package.
package mock_currency

import (
	context "context"
	rate "currency/internal/domain/rate"
	currency "currency/internal/infrastructure/currency"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICryptoProvider is a mock of ICryptoProvider interface.
type MockICryptoProvider struct {
	ctrl     *gomock.Controller
	recorder *MockICryptoProviderMockRecorder
}

// MockICryptoProviderMockRecorder is the mock recorder for MockICryptoProvider.
type MockICryptoProviderMockRecorder struct {
	mock *MockICryptoProvider
}

// NewMockICryptoProvider creates a new mock instance.
func NewMockICryptoProvider(ctrl *gomock.Controller) *MockICryptoProvider {
	mock := &MockICryptoProvider{ctrl: ctrl}
	mock.recorder = &MockICryptoProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICryptoProvider) EXPECT() *MockICryptoProviderMockRecorder {
	return m.recorder
}

// GetCurrency mocks base method.
func (m *MockICryptoProvider) GetCurrency(ctx context.Context, data *rate.Rate) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrency", ctx, data)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrency indicates an expected call of GetCurrency.
func (mr *MockICryptoProviderMockRecorder) GetCurrency(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrency", reflect.TypeOf((*MockICryptoProvider)(nil).GetCurrency), ctx, data)
}

// MockIChain is a mock of IChain interface.
type MockIChain struct {
	ctrl     *gomock.Controller
	recorder *MockIChainMockRecorder
}

// MockIChainMockRecorder is the mock recorder for MockIChain.
type MockIChainMockRecorder struct {
	mock *MockIChain
}

// NewMockIChain creates a new mock instance.
func NewMockIChain(ctrl *gomock.Controller) *MockIChain {
	mock := &MockIChain{ctrl: ctrl}
	mock.recorder = &MockIChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIChain) EXPECT() *MockIChainMockRecorder {
	return m.recorder
}

// GetCurrency mocks base method.
func (m *MockIChain) GetCurrency(ctx context.Context, data *rate.Rate) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrency", ctx, data)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrency indicates an expected call of GetCurrency.
func (mr *MockIChainMockRecorder) GetCurrency(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrency", reflect.TypeOf((*MockIChain)(nil).GetCurrency), ctx, data)
}

// SetNext mocks base method.
func (m *MockIChain) SetNext(arg0 currency.IChain) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNext", arg0)
}

// SetNext indicates an expected call of SetNext.
func (mr *MockIChainMockRecorder) SetNext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNext", reflect.TypeOf((*MockIChain)(nil).SetNext), arg0)
}

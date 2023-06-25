// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go

// Package mock_http is a generated GoMock package.
package mock_http

import (
	context "context"
	event_domain "exchange/internal/domain/event_domain"
	rate_domain "exchange/internal/domain/rate_domain"
	user_domain "exchange/internal/domain/user_domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockINotificationService is a mock of INotificationService interface.
type MockINotificationService struct {
	ctrl     *gomock.Controller
	recorder *MockINotificationServiceMockRecorder
}

// MockINotificationServiceMockRecorder is the mock recorder for MockINotificationService.
type MockINotificationServiceMockRecorder struct {
	mock *MockINotificationService
}

// NewMockINotificationService creates a new mock instance.
func NewMockINotificationService(ctrl *gomock.Controller) *MockINotificationService {
	mock := &MockINotificationService{ctrl: ctrl}
	mock.recorder = &MockINotificationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINotificationService) EXPECT() *MockINotificationServiceMockRecorder {
	return m.recorder
}

// Notify mocks base method.
func (m *MockINotificationService) Notify(ctx context.Context, n *event_domain.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notify", ctx, n)
	ret0, _ := ret[0].(error)
	return ret0
}

// Notify indicates an expected call of Notify.
func (mr *MockINotificationServiceMockRecorder) Notify(ctx, n interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockINotificationService)(nil).Notify), ctx, n)
}

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
func (m *MockICurrencyService) GetCurrency(ctx context.Context, c *rate_domain.Rate) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrency", ctx, c)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrency indicates an expected call of GetCurrency.
func (mr *MockICurrencyServiceMockRecorder) GetCurrency(ctx, c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrency", reflect.TypeOf((*MockICurrencyService)(nil).GetCurrency), ctx, c)
}

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// NewUser mocks base method.
func (m *MockIUserService) NewUser(ctx context.Context, eu *user_domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUser", ctx, eu)
	ret0, _ := ret[0].(error)
	return ret0
}

// NewUser indicates an expected call of NewUser.
func (mr *MockIUserServiceMockRecorder) NewUser(ctx, eu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUser", reflect.TypeOf((*MockIUserService)(nil).NewUser), ctx, eu)
}
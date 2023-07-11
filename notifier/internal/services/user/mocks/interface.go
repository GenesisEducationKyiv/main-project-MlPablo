// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	user "os/user"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

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
func (m *MockIUserService) NewUser(ctx context.Context, eu *user.User) error {
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

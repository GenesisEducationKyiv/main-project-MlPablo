package services_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain/user"
	mock_user "exchange/internal/domain/user/mocks"
	"exchange/internal/services"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_user.NewMockUserRepository(ctrl)

	const email = "test@email.com"

	gomock.InOrder(
		userRepoMock.EXPECT().EmailExist(context.Background(), email).Return(false, nil),
		userRepoMock.EXPECT().SaveUser(context.Background(), user.NewUser(email)).Return(nil),
	)

	userService := services.NewUserService(userRepoMock)

	err := userService.NewUser(context.Background(), user.NewUser(email))
	require.NoError(t, err)
}

func TestCreateUserAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_user.NewMockUserRepository(ctrl)

	const email = "test@email.com"

	userRepoMock.EXPECT().EmailExist(context.Background(), email).Return(true, nil)

	userService := services.NewUserService(userRepoMock)

	err := userService.NewUser(context.Background(), user.NewUser(email))
	require.ErrorIs(t, err, user.ErrAlreadyExist)
}

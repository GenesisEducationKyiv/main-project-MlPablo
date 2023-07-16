package user_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	user_domain "notifier/internal/domain/user"
	"notifier/internal/services/user"
	mock_user "notifier/internal/services/user/mocks"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_user.NewMockUserRepository(ctrl)

	const email = "test@email.com"

	gomock.InOrder(
		userRepoMock.EXPECT().EmailExist(context.Background(), email).Return(false, nil),
		userRepoMock.EXPECT().
			SaveUser(context.Background(), user_domain.NewUser(email)).
			Return(nil),
	)

	userService := user.NewUserService(userRepoMock)

	err := userService.NewUser(context.Background(), user_domain.NewUser(email))
	require.NoError(t, err)
}

func TestCreateUserAlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_user.NewMockUserRepository(ctrl)

	const email = "test@email.com"

	userRepoMock.EXPECT().EmailExist(context.Background(), email).Return(true, nil)

	userService := user.NewUserService(userRepoMock)

	err := userService.NewUser(context.Background(), user_domain.NewUser(email))
	require.ErrorIs(t, err, user_domain.ErrAlreadyExist)
}

package functional

import (
	"context"

	"github.com/go-faker/faker/v4"

	"notifier/internal/domain/notification"
	"notifier/internal/domain/user"
)

func (suite *Suite) TestUserExist() {
	u := user.NewUser(faker.Email())

	err := suite.srv.userService.NewUser(context.Background(), u)
	suite.Require().NoError(err)

	err = suite.srv.userService.NewUser(context.Background(), u)
	suite.Require().ErrorIs(err, user.ErrAlreadyExist)
}

func (suite *Suite) TestSendEmails() {
	err := suite.srv.notifyService.Notify(
		context.Background(),
		notification.DefaultNotification(),
	)
	suite.Require().NoError(err)
}

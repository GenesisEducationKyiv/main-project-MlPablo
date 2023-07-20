package functional

import (
	"context"

	"github.com/go-faker/faker/v4"

	"exchange/internal/domain/notification"
	"exchange/internal/domain/rate"
	"exchange/internal/domain/user"
)

func (suite *Suite) TestUserExist() {
	u := user.NewUser(faker.Email())

	err := suite.srv.userService.NewUser(context.Background(), u)
	suite.Require().NoError(err)

	err = suite.srv.userService.NewUser(context.Background(), u)
	suite.Require().ErrorIs(err, user.ErrAlreadyExist)
}

func (suite *Suite) TestGetCurrency() {
	res, err := suite.srv.currecnyService.GetCurrency(
		context.Background(),
		rate.GetBitcoinToUAH(),
	)
	suite.Require().NoError(err)
	suite.NotZero(res)
}

func (suite *Suite) TestSendEmails() {
	err := suite.srv.notifyService.Notify(
		context.Background(),
		notification.DefaultNotification(),
	)
	suite.Require().NoError(err)
}

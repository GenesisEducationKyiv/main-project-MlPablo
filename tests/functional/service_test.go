package functional

import (
	"context"

	"github.com/go-faker/faker/v4"

	"exchange/internal/domain/notification"
	rate_domain "exchange/internal/domain/rate"
	user_domain "exchange/internal/domain/user"
)

func (suite *Suite) TestUserExist() {
	u := user_domain.NewUser(faker.Email())

	err := suite.srv.userService.NewUser(context.Background(), u)
	suite.Require().NoError(err)

	err = suite.srv.userService.NewUser(context.Background(), u)
	suite.Require().ErrorIs(err, user_domain.ErrAlreadyExist)
}

func (suite *Suite) TestGetCurrency() {
	res, err := suite.srv.currecnyService.GetCurrency(
		context.Background(),
		rate_domain.GetBitcoinToUAH(),
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

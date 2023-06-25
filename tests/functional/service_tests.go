package functional

import (
	"context"

	"github.com/bxcodec/faker/v3"

	"exchange/internal/domain/event_domain"
	"exchange/internal/domain/rate_domain"
	"exchange/internal/domain/user_domain"
)

func (suite *Suite) TestValidCreateUser() {
	user := user_domain.NewUser(faker.Email())

	err := suite.srv.userService.NewUser(context.Background(), user)
	suite.Require().NoError(err)

	ok, err := suite.checkRowExistInFileDB(user.Email)
	suite.Require().NoError(err)
	suite.True(ok)
}

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
		event_domain.DefaultNotification(),
	)
	suite.Require().NoError(err)
}

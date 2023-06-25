package functional

import (
	"context"

	"github.com/bxcodec/faker/v3"

	"exchange/internal/domain/event"
	"exchange/internal/domain/rate"
	"exchange/internal/domain/user"
)

func (suite *Suite) TestValidCreateUser() {
	user := user.NewUser(faker.Email())

	err := suite.srv.userService.NewUser(context.Background(), user)
	suite.Require().NoError(err)

	ok, err := suite.checkRowExistInFileDB(user.Email)
	suite.Require().NoError(err)
	suite.True(ok)
}

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
		event.DefaultNotification(),
	)
	suite.Require().NoError(err)
}

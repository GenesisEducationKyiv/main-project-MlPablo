package functional

import (
	"context"

	"github.com/bxcodec/faker/v3"

	"exchange/internal/domain"
)

func (suite *Suite) TestValidCreateUser() {
	user := domain.NewUser(faker.Email())

	err := suite.srv.userService.NewUser(context.Background(), user)
	suite.Require().NoError(err)

	ok, err := suite.checkRowExistInFileDB(user.Email)
	suite.Require().NoError(err)
	suite.True(ok)
}

func (suite *Suite) TestUserExist() {
	user := domain.NewUser(faker.Email())

	err := suite.srv.userService.NewUser(context.Background(), user)
	suite.Require().NoError(err)

	err = suite.srv.userService.NewUser(context.Background(), user)
	suite.Require().ErrorIs(err, domain.ErrAlreadyExist)
}

func (suite *Suite) TestGetCurrency() {
	res, err := suite.srv.currecnyService.GetCurrency(
		context.Background(),
		domain.GetBitcoinToUAH(),
	)
	suite.Require().NoError(err)
	suite.NotZero(res)
}

func (suite *Suite) TestSendEmails() {
	err := suite.srv.notifyService.Notify(
		context.Background(),
		domain.DefaultNotification(),
	)
	suite.Require().NoError(err)
}

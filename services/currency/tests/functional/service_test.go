package functional

import (
	"context"

	"currency/internal/domain/rate"
)

func (suite *Suite) TestGetCurrency() {
	res, err := suite.srv.currecnyService.GetCurrency(
		context.Background(),
		rate.GetBitcoinToUAH(),
	)
	suite.Require().NoError(err)
	suite.NotZero(res)
}

package functional

import (
	"os"

	"github.com/stretchr/testify/suite"

	"currency/internal/services/currency"
)

const testFilePath = "test_path.txt"

type Services struct {
	currecnyService *currency.Service
}

type Suite struct {
	suite.Suite
	srv *Services
}

func (suite *Suite) SetupSuite() {
	stubs := new(thirdParyStubs)

	currecnyService := currency.NewCurrencyService(stubs)

	srvs := &Services{
		currecnyService: currecnyService,
	}

	suite.srv = srvs
}

func (suite *Suite) AfterTest(_, _ string) {
	os.Remove(testFilePath)
}

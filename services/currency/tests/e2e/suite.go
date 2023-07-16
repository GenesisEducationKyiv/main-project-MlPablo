package e2e

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"currency/internal/controller/http"
	"currency/internal/infrastructure/currency/currencyapi"
	"currency/internal/services/currency"
	"currency/utils"
)

type Services struct {
	currecnyService *currency.Service
}

type Suite struct {
	suite.Suite
	srv *Services
	e   *echo.Echo
}

func (suite *Suite) SetupSuite() {
	if err := godotenv.Load("../../.env"); err != nil {
		logrus.Fatal(err)
	}

	envGet := utils.TryGetEnv[string]

	currencyAPI := currencyapi.NewCurrencyAPI(
		currencyapi.NewConfig(
			envGet("CURR_API_KEY"),
			envGet("CURR_URL"),
		),
	)

	currecnyService := currency.NewCurrencyService(currencyAPI)

	srvs := &Services{
		currecnyService: currecnyService,
	}

	e := echo.New()

	http.RegisterHandlers(e, &http.Services{
		CurrencyService: currecnyService,
	})

	suite.srv = srvs
	suite.e = e
}

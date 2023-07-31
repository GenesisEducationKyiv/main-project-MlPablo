package e2e

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"currency/internal/controller/http"
	"currency/internal/infrastructure/currency/binance"
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

	binanceAPI := binance.NewBinanceApi(binance.NewConfig(envGet("BINANCE_URL")))

	currecnyService := currency.NewCurrencyService(binanceAPI)

	srvs := &Services{
		currecnyService: currecnyService,
	}

	e := echo.New()

	http.RegisterHandlers(e, &http.Services{
		CurrencyService: currecnyService,
	}, new(eventer))

	suite.srv = srvs
	suite.e = e
}

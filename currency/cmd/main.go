package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"currency/internal/controller/http"
	"currency/internal/infrastructure/currency/binance"
	"currency/internal/infrastructure/currency/coingecko"
	"currency/internal/infrastructure/currency/currencyapi"
	"currency/internal/services/currency"
	echoserver "currency/pkg/echo"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("starting application...")

	fx.New(CreateApp()).Run()

	logrus.Info("application stopped.")
}

func CreateApp() fx.Option { //nolint: ireturn // ok
	return fx.Options(
		fx.Provide(
			NewServices,
			NewBinanceConfig,
			NewEchoConfig,
			NewCurrencyapiConfig,
			NewCoingeckoConfig,
			createChan,
			func(cfg *currencyapi.Config) *currencyapi.CurrencyAPI {
				return currencyapi.NewCurrencyAPI(
					cfg,
					currencyapi.WithLogger(logrus.StandardLogger()),
				)
			},
			func(cfg *binance.Config) *binance.BinanceAPI {
				return binance.NewBinanceApi(
					cfg,
					binance.WithLogger(logrus.StandardLogger()),
				)
			},
			func(cfg *coingecko.Config) *coingecko.CoingeckoAPI {
				return coingecko.NewCoingeckoApi(
					cfg,
					coingecko.WithLogger(logrus.StandardLogger()),
				)
			},
			fx.Annotate(
				func(c *currencyapi.CurrencyAPI) *currencyapi.CurrencyAPI {
					return c
				},
				fx.As(new(currency.ICurrencyAPI)),
			),
			fx.Annotate(
				currency.NewCurrencyService,
				fx.As(new(http.ICurrencyService)),
			),
			echoserver.New,
		),
		fx.Invoke(
			startErrorHandling,
			registerCryptoChain,
			registerHttpHandlers,
			startServer,
		),
	)
}

func createChan() chan error {
	return make(chan error)
}

func startServer(srv *echoserver.Server, ls fx.Lifecycle, errChan chan error) {
	ls.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				go srv.Start(errChan)
				return nil
			},
			OnStop: func(_ context.Context) error {
				return srv.Stop()
			},
		},
	)
}

func startErrorHandling(shutdowner fx.Shutdowner, errChan chan error) {
	go func(ch chan error) {
		err := <-ch
		logrus.Error(err)
		shutdowner.Shutdown() //nolint: errcheck// no need here to check
	}(errChan)
}

func registerCryptoChain(
	cu *currencyapi.CurrencyAPI,
	b *binance.BinanceAPI,
	co *coingecko.CoingeckoAPI,
) error {
	if err := cu.SetNext(b); err != nil {
		return err
	}

	if err := b.SetNext(co); err != nil {
		return err
	}

	return nil
}

func registerHttpHandlers(srv *http.Services, e *echoserver.Server) {
	http.RegisterHandlers(e.GetEchoServer(), srv)
}

func NewServices(
	c http.ICurrencyService,
) *http.Services {
	return &http.Services{
		CurrencyService: c,
	}
}
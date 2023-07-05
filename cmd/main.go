package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"exchange/internal/controller/http"
	"exchange/internal/infrastructure/currency/binance"
	"exchange/internal/infrastructure/currency/coingecko"
	"exchange/internal/infrastructure/currency/currencyapi"
	"exchange/internal/infrastructure/mail"
	"exchange/internal/infrastructure/repository/filesystem"
	"exchange/internal/services/currency"
	"exchange/internal/services/event"
	"exchange/internal/services/user"
	echoserver "exchange/pkg/echo"
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
			NewMailConfig,
			NewCurrencyapiConfig,
			NewFileSystemConfig,
			NewCoingechoConfig,
			createChan,
			fx.Annotate(
				filesystem.NewFileSystemRepository,
				fx.As(new(user.UserRepository)),
				fx.As(new(event.UserRepository)),
			),
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
			fx.Annotate(mail.NewMailService, fx.As(new(event.IMailService))),
			fx.Annotate(
				currency.NewCurrencyService,
				fx.As(new(event.ICurrencyService)),
				fx.As(new(http.ICurrencyService)),
			),
			fx.Annotate(user.NewUserService, fx.As(new(http.IUserService))),
			fx.Annotate(event.NewNotificationService, fx.As(new(http.INotificationService))),
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
	ls.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go srv.Start(errChan)
			return nil
		},
		OnStop: func(_ context.Context) error {
			return srv.Stop()
		},
	})
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

func registerHttpHandlers(srv *http.Services, srva *echoserver.Server) {
	http.RegisterHandlers(srva.GetEchoServer(), srv)
}

func NewServices(
	c http.ICurrencyService,
	u http.IUserService,
	e http.INotificationService,
) *http.Services {
	return &http.Services{
		CurrencyService:     c,
		UserService:         u,
		NotificationService: e,
	}
}

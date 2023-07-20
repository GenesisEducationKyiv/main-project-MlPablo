package main

import (
	"context"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	_http "exchange/internal/controller/http"
	"exchange/internal/infrastructure/currency/binance"
	"exchange/internal/infrastructure/currency/coingecko"
	"exchange/internal/infrastructure/currency/currencyapi"
	"exchange/internal/infrastructure/mail"
	"exchange/internal/infrastructure/repository/filesystem"
	"exchange/internal/services/currency"
	"exchange/internal/services/event"
	"exchange/internal/services/user"
	echoserver "exchange/pkg/echo"
	"exchange/pkg/http/client"
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
			NewCoingeckoConfig,
			createChan,
			fx.Annotate(
				filesystem.NewFileSystemRepository,
				fx.As(new(user.UserRepository)),
				fx.As(new(event.UserRepository)),
			),
			func() *http.Client {
				return client.New(client.WithLogger(logrus.StandardLogger()))
			},
			func(cfg *currencyapi.Config, cli *http.Client) *currencyapi.CurrencyAPI {
				return currencyapi.NewCurrencyAPI(
					cfg,
					currencyapi.WithClient(cli),
				)
			},
			func(cfg *binance.Config, cli *http.Client) *binance.BinanceAPI {
				return binance.NewBinanceApi(
					cfg,
					binance.WithClient(cli),
				)
			},
			func(cfg *coingecko.Config, cli *http.Client) *coingecko.CoingeckoAPI {
				return coingecko.NewCoingeckoApi(
					cfg,
					coingecko.WithClient(cli),
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
				fx.As(new(currency.ICurrencyService)),
				fx.As(new(event.ICurrencyService)),
			),
			fx.Annotate(user.NewUserService, fx.As(new(user.IUserService))),
			fx.Annotate(event.NewNotificationService, fx.As(new(event.INotificationService))),
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
) {
	cu.SetNext(b)
	b.SetNext(co)
}

func registerHttpHandlers(srv *_http.Services, e *echoserver.Server) {
	_http.RegisterHandlers(e.GetEchoServer(), srv)
}

func NewServices(
	c currency.ICurrencyService,
	u user.IUserService,
	e event.INotificationService,
) *_http.Services {
	return &_http.Services{
		CurrencyService:     c,
		UserService:         u,
		NotificationService: e,
	}
}

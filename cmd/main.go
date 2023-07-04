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
	"exchange/utils"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal(err)
	}

	// ctx, cancel := signal.NotifyContext(
	// 	context.Background(),
	// 	syscall.SIGINT,
	// 	syscall.SIGTERM,
	// 	syscall.SIGHUP,
	// )
	// defer cancel()

	logrus.Info("starting application...")

	// app, err := app.New()
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	//
	// if err = app.Run(ctx, cancel); err != nil {
	// 	logrus.Fatal(err)
	// }
	//
	// logrus.Info("application started =)")

	fx.New(CreateApp(make(chan error))).Run()
	// <-ctx.Done()
	logrus.Info("application stopped.")
}

func CreateApp(errChan chan error) fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				filesystem.NewFileSystemRepository,
				fx.As(new(user.UserRepository)),
				fx.As(new(event.UserRepository)),
			),
			fx.Annotate(
				currencyapi.NewCurrencyAPI,
				fx.As(new(currency.ICurrencyAPI)),
			),
			binance.NewBinanceApi,
			coingecko.NewCoingeckoApi,
			fx.Annotate(mail.NewMailService, fx.As(new(event.IMailService))),
			fx.Annotate(
				currency.NewCurrencyService,
				fx.As(new(event.ICurrencyService)),
				fx.As(new(http.ICurrencyService)),
				// fx.As(new(http.ICurrencyService)),
			),
			fx.Annotate(user.NewUserService, fx.As(new(http.IUserService))),
			fx.Annotate(event.NewNotificationService, fx.As(new(http.INotificationService))),
			echoserver.New,
			// http.RegisterHandlers,
			NewServices,
			NewConfig1,
			NewConfig2,
			NewConfig3,
			Config57,
		),
		fx.Invoke(
			func(srv *http.Services, srva *echoserver.Server) {
				http.RegisterHandlers(srva.GetEchoServer(), srv)
			},
			func(srv *echoserver.Server, ls fx.Lifecycle) {
				ls.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						go srv.Start(errChan)
						return nil
					},
					OnStop: func(_ context.Context) error {
						return srv.Stop()
					},
				})
			}),
		// ...fx.Invoke(
	)
}

func NewConfig1() *echoserver.Config {
	return &echoserver.Config{
		Address: utils.TryGetEnvDefault("SERVER_ADDR", "8080"),
	}
}

func NewConfig2() *mail.Config {
	envGet := utils.TryGetEnv[string]
	return mail.NewConfig(
		envGet("EMAIL_LOGIN"),
		envGet("EMAIL_APP_PASSWORD"),
		envGet("SMTP_HOST"),
		envGet("SMTP_PORT"),
	)
}

func NewConfig3() *currencyapi.Config {
	envGet := utils.TryGetEnv[string]
	return currencyapi.NewConfig(
		envGet("CURR_API_KEY"),
		envGet("CURR_URL"),
	)
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

func Config57() *filesystem.Config {
	return &filesystem.Config{Path: utils.TryGetEnvDefault("FILE_STORE_PATH", "./file_storage.txt")}
}

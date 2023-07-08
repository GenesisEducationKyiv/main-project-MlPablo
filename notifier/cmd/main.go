package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"notifier/internal/controller/http"
	"notifier/internal/infrastructure/mail"
	"notifier/internal/infrastructure/repository/filesystem"
	"notifier/internal/services/currency"
	"notifier/internal/services/event"
	"notifier/internal/services/user"
	echoserver "notifier/pkg/echo"
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
			NewEchoConfig,
			NewMailConfig,
			NewFileSystemConfig,
			createChan,
			fx.Annotate(
				filesystem.NewFileSystemRepository,
				fx.As(new(user.UserRepository)),
				fx.As(new(event.UserRepository)),
			),
			fx.Annotate(mail.NewMailService, fx.As(new(event.IMailService))),
			fx.Annotate(
				currency.New,
				fx.As(new(event.ICurrencyService)),
			),
			fx.Annotate(user.NewUserService, fx.As(new(http.IUserService))),
			fx.Annotate(event.NewNotificationService, fx.As(new(http.INotificationService))),
			echoserver.New,
		),
		fx.Invoke(
			startErrorHandling,
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

func registerHttpHandlers(srv *http.Services, e *echoserver.Server) {
	http.RegisterHandlers(e.GetEchoServer(), srv)
}

func NewServices(
	u http.IUserService,
	e http.INotificationService,
) *http.Services {
	return &http.Services{
		UserService:         u,
		NotificationService: e,
	}
}
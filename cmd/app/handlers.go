package app

import (
	"os"

	_http "exchange/internal/controller/http"
	"exchange/internal/domain"
	"exchange/internal/infrastructure/currency/currencyapi"
	"exchange/internal/infrastructure/mail"
	"exchange/internal/repository/filesysytem"
	"exchange/internal/services"
)

type Services struct {
	CurrencyService     domain.ICurrencyService
	UserService         domain.IUserService
	NotificationService domain.INotificationService
}

func creaeteServicesAndHandlers(server *Servers) error {
	mailRepo, err := filesysytem.NewFileSystemRepository(os.Getenv("FILE_STORE_PATH"))
	if err != nil {
		return err
	}

	mCfg, err := mail.NewConfig(
		os.Getenv("EMAIL_LOGIN"),
		os.Getenv("EMAIL_APP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
	)
	if err != nil {
		return err
	}

	mailPusher := mail.NewMailService(mCfg)

	currencyGetter := currencyapi.NewCurrencyAPI(
		currencyapi.NewConfig(os.Getenv("CURR_API_KEY")),
		os.Getenv("CURR_URL"),
	)

	userMailService := services.NewUserService(mailRepo)
	notifierService := services.NewNotificationService(mailRepo, currencyGetter, mailPusher)

	_http.RegisterHandlers(server.HTTPServer.GetEchoServer(), &_http.Services{
		CurrencyService:     currencyGetter,
		UserService:         userMailService,
		NotificationService: notifierService,
	})

	return nil
}

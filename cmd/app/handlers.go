package app

import (
	_http "exchange/internal/controller/http"
	"exchange/internal/domain"
	"exchange/internal/infrastructure/currency/currencyapi"
	"exchange/internal/infrastructure/mail"
	"exchange/internal/repository/filesysytem"
	"exchange/internal/services"
	"exchange/utils"
)

type Services struct {
	CurrencyService     domain.ICurrencyService
	UserService         domain.IUserService
	NotificationService domain.INotificationService
}

func creaeteServicesAndHandlers(server *Servers) error {
	envGet := utils.TryGetEnv[string]

	mailRepo, err := filesysytem.NewFileSystemRepository(
		utils.TryGetEnvDefault[string]("FILE_STORE_PATH", "./file steerage.txt"),
	)
	if err != nil {
		return err
	}

	mCfg, err := mail.NewConfig(
		envGet("EMAIL_LOGIN"),
		envGet("EMAIL_APP_PASSWORD"),
		envGet("SMTP_HOST"),
		envGet("SMTP_PORT"),
	)
	if err != nil {
		return err
	}

	mailPusher := mail.NewMailService(mCfg)

	currencyGetter := currencyapi.NewCurrencyAPI(
		currencyapi.NewConfig(
			envGet("CURR_API_KEY"),
			envGet("CURR_URL"),
		),
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

package app

import (
	_http "exchange/internal/controller/http"
	"exchange/internal/domain"
	"exchange/internal/infrastructure/currency/currencyapi"
	"exchange/internal/infrastructure/mail"
	"exchange/internal/repository/filesystem"
	"exchange/internal/services"
	"exchange/utils"
)

type Services struct {
	CurrencyService     domain.ICurrencyService
	UserService         domain.IUserService
	NotificationService domain.INotificationService
}

func createServicesAndHandlers(server *Servers) error {
	envGet := utils.TryGetEnv[string]

	mailRepo, err := filesystem.NewFileSystemRepository(
		utils.TryGetEnvDefault("FILE_STORE_PATH", "./file_storage.txt"),
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

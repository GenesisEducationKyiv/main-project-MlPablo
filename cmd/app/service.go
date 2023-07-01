package app

import (
	"exchange/internal/infrastructure/currency/currencyapi"
	"exchange/internal/infrastructure/mail"
	"exchange/internal/repository/filesystem"
	"exchange/internal/services/currency"
	"exchange/internal/services/event"
	"exchange/internal/services/user"
	"exchange/utils"
)

type Services struct {
	CurrencyService     *currency.Service
	UserService         *user.Service
	NotificationService *event.Service
}

type ThirdPartyServices struct {
	MailSender  *mail.EmailSender
	CurrencyAPI *currencyapi.CurrencyAPI
}

type Repositories struct {
	fileRepo *filesystem.Repository
}

func createServices() (*Services, error) {
	repo, err := createRepositories()
	if err != nil {
		return nil, err
	}

	tds := createThirdPartyServices()

	return &Services{
		UserService: user.NewUserService(repo.fileRepo),
		NotificationService: event.NewNotificationService(
			repo.fileRepo,
			tds.CurrencyAPI,
			tds.MailSender,
		),
		CurrencyService: currency.NewCurrencyService(tds.CurrencyAPI),
	}, nil
}

func createThirdPartyServices() *ThirdPartyServices {
	envGet := utils.TryGetEnv[string]

	mailSender := mail.NewMailService(mail.NewConfig(
		envGet("EMAIL_LOGIN"),
		envGet("EMAIL_APP_PASSWORD"),
		envGet("SMTP_HOST"),
		envGet("SMTP_PORT"),
	))

	currencyAPI := currencyapi.NewCurrencyAPI(
		currencyapi.NewConfig(
			envGet("CURR_API_KEY"),
			envGet("CURR_URL"),
		),
	)

	return &ThirdPartyServices{
		MailSender:  mailSender,
		CurrencyAPI: currencyAPI,
	}
}

func createRepositories() (*Repositories, error) {
	fileRepo, err := filesystem.NewFileSystemRepository(
		utils.TryGetEnvDefault("FILE_STORE_PATH", "./file_storage.txt"),
	)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		fileRepo: fileRepo,
	}, nil
}

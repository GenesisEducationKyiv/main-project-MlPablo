package app

import (
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

type ThirdPartyServices struct {
	MailSender  *mail.EmailSender
	CurrencyAPI *currencyapi.CurrencyAPI
}

type Repositories struct {
	UserRepo domain.UserRepository
}

func createServices() (*Services, error) {
	repo, err := createRepositories()
	if err != nil {
		return nil, err
	}

	tds := createThirdPartyServices()

	return &Services{
		UserService: services.NewUserService(repo.UserRepo),
		NotificationService: services.NewNotificationService(
			repo.UserRepo,
			tds.CurrencyAPI,
			tds.MailSender,
		),
		CurrencyService: services.NewCurrencyService(tds.CurrencyAPI),
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
	userRepo, err := filesystem.NewFileSystemRepository(
		utils.TryGetEnvDefault("FILE_STORE_PATH", "./file_storage.txt"),
	)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		UserRepo: userRepo,
	}, nil
}

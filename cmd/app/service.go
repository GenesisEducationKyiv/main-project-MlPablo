package app

import (
	"github.com/sirupsen/logrus"

	"exchange/internal/infrastructure/currency/binance"
	"exchange/internal/infrastructure/currency/coingecko"
	"exchange/internal/infrastructure/currency/currencyapi"
	"exchange/internal/infrastructure/mail"
	"exchange/internal/infrastructure/repository/filesystem"
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
	MailSender   *mail.EmailSender
	CurrencyAPI  *currencyapi.CurrencyAPI
	BinanceAPI   *binance.BinanceAPI
	CoingeckoAPI *coingecko.CoingeckoAPI
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

	if err = tds.BinanceAPI.SetNext(tds.CurrencyAPI); err != nil {
		return nil, err
	}

	if err = tds.CurrencyAPI.SetNext(tds.CoingeckoAPI); err != nil {
		return nil, err
	}

	currencyService := currency.NewCurrencyService(tds.BinanceAPI)

	return &Services{
		UserService: user.NewUserService(repo.fileRepo),
		NotificationService: event.NewNotificationService(
			repo.fileRepo,
			currencyService,
			tds.MailSender,
		),
		CurrencyService: currencyService,
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
		), currencyapi.WithLogger(logrus.StandardLogger()),
	)

	binanceAPI := binance.NewBinanceApi(
		binance.NewConfig(
			envGet("BINANCE_URL"),
		), binance.WithLogger(logrus.StandardLogger()),
	)

	coingeckoAPI := coingecko.NewCoingeckoApi(
		coingecko.NewConfig(
			envGet("COINGECKO_URL"),
		), coingecko.WithLogger(logrus.StandardLogger()),
	)

	return &ThirdPartyServices{
		MailSender:   mailSender,
		CurrencyAPI:  currencyAPI,
		BinanceAPI:   binanceAPI,
		CoingeckoAPI: coingeckoAPI,
	}
}

func createRepositories() (*Repositories, error) {
	fileRepo, err := filesystem.NewFileSystemRepository(
		&filesystem.Config{utils.TryGetEnvDefault("FILE_STORE_PATH", "./file_storage.txt")},
	)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		fileRepo: fileRepo,
	}, nil
}

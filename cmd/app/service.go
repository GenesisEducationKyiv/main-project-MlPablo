package app

import (
	"exchange/internal/infrastructure/currency/binance"
	"exchange/internal/infrastructure/currency/coingecko"
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
		),
	)

	binanceAPI := binance.NewBinanceApi(
		binance.NewConfig(
			envGet("BINANCE_URL"),
		),
	)

	coingeckoAPI := coingecko.NewCoingeckoApi(
		coingecko.NewConfig(
			envGet("COINGECKO_URL"),
		),
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
		utils.TryGetEnvDefault("FILE_STORE_PATH", "./file_storage.txt"),
	)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		fileRepo: fileRepo,
	}, nil
}

// type Chain interface {
// 	GetCurrency(ctx context.Context, r *rate.Rate) (float64, error)
// 	SetNext(next Chain)
// }
//
// type provider struct {
// 	next Chain
// }
//
// func NewProvider() *provider {
// 	return new(provider)
// }
//
// func (p *provider) SetNext(next Chain) {
// 	p.next = next
// }
//
// func (p *provider) GetCurrency(ctx context.Context, r *rate.Rate) (float64, error) {
// 	rate, err := p.next.GetCurrency(ctx, r)
// 	if err != nil {
// 		if p.next == nil {
// 			return 0, err
// 		}
// 		return p.next.GetCurrency(ctx, r)
// 	}
//
// 	return rate, nil
// }

package e2e

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"exchange/internal/controller/http"
	"exchange/internal/infrastructure/currency/currencyapi"
	"exchange/internal/infrastructure/mail"
	"exchange/internal/repository/filesystem"
	"exchange/internal/services/currency"
	"exchange/internal/services/event"
	"exchange/internal/services/user"
	"exchange/utils"
)

const testFilePath = "test_path.txt"

type Services struct {
	currecnyService *currency.Service
	notifyService   *event.Service
	userService     *user.Service
}

type Suite struct {
	suite.Suite
	srv  *Services
	repo *filesystem.Repository
	e    *echo.Echo
}

func (suite *Suite) SetupSuite() {
	if err := godotenv.Load("../../.env"); err != nil {
		logrus.Fatal(err)
	}

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

	fileRepo, err := filesystem.NewFileSystemRepository(testFilePath)
	if err != nil {
		logrus.Fatal(err)
	}

	currecnyService := currency.NewCurrencyService(currencyAPI)
	userService := user.NewUserService(fileRepo)
	notificationService := event.NewNotificationService(
		fileRepo,
		currecnyService,
		mailSender,
	)

	srvs := &Services{
		currecnyService: currecnyService,
		userService:     userService,
		notifyService:   notificationService,
	}

	e := echo.New()

	http.RegisterHandlers(e, &http.Services{
		UserService:         userService,
		CurrencyService:     currecnyService,
		NotificationService: notificationService,
	})

	suite.srv = srvs
	suite.repo = fileRepo
	suite.e = e
}

func (suite *Suite) AfterTest(_, _ string) {
	os.Remove(testFilePath)
}

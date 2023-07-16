package e2e

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"notifier/internal/controller/http"
	"notifier/internal/infrastructure/mail"
	"notifier/internal/infrastructure/repository/filesystem"
	"notifier/internal/services/event"
	"notifier/internal/services/user"
	"notifier/utils"
)

const testFilePath = "test_path.txt"

type Services struct {
	notifyService *event.Service
	userService   *user.Service
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

	fileRepo, err := filesystem.NewFileSystemRepository(&filesystem.Config{
		Path: testFilePath,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	stubs := new(thirdParyStubs)

	userService := user.NewUserService(fileRepo)
	notificationService := event.NewNotificationService(
		fileRepo,
		stubs,
		mailSender,
	)

	srvs := &Services{
		userService:   userService,
		notifyService: notificationService,
	}

	e := echo.New()

	http.RegisterHandlers(e, &http.Services{
		UserService:         userService,
		NotificationService: notificationService,
	})

	suite.srv = srvs
	suite.repo = fileRepo
	suite.e = e
}

func (suite *Suite) AfterTest(_, _ string) {
	os.Remove(testFilePath)
}

package functional

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"exchange/internal/repository/filesystem"
	"exchange/internal/services/currency_service"
	"exchange/internal/services/event_service"
	"exchange/internal/services/user_service"
)

const testFilePath = "test_path.txt"

type Services struct {
	currecnyService *currency_service.Service
	notifyService   *event_service.Service
	userService     *user_service.Service
}

type Suite struct {
	suite.Suite
	srv  *Services
	repo *filesystem.Repository
}

func (suite *Suite) SetupSuite() {
	fileRepo, err := filesystem.NewFileSystemRepository(testFilePath)
	if err != nil {
		logrus.Fatal(err)
	}

	stubs := new(thirdParyStubs)

	currecnyService := currency_service.NewCurrencyService(stubs)
	userService := user_service.NewUserService(fileRepo)
	notificationService := event_service.NewNotificationService(fileRepo, currecnyService, stubs)

	srvs := &Services{
		currecnyService: currecnyService,
		userService:     userService,
		notifyService:   notificationService,
	}

	suite.srv = srvs
	suite.repo = fileRepo
}

func (suite *Suite) AfterTest(_, _ string) {
	suite.NoError(suite.repo.DeleteFile())
}

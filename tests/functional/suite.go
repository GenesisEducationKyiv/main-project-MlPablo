package functional

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"exchange/internal/infrastructure/repository/filesystem"
	"exchange/internal/services/currency"
	"exchange/internal/services/event"
	"exchange/internal/services/user"
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
}

func (suite *Suite) SetupSuite() {
	fileRepo, err := filesystem.NewFileSystemRepository(testFilePath)
	if err != nil {
		logrus.Fatal(err)
	}

	stubs := new(thirdParyStubs)

	currecnyService := currency.NewCurrencyService(stubs)
	userService := user.NewUserService(fileRepo)
	notificationService := event.NewNotificationService(fileRepo, currecnyService, stubs)

	srvs := &Services{
		currecnyService: currecnyService,
		userService:     userService,
		notifyService:   notificationService,
	}

	suite.srv = srvs
	suite.repo = fileRepo
}

func (suite *Suite) AfterTest(_, _ string) {
	os.Remove(testFilePath)
}

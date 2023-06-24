package functional

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"exchange/internal/domain"
	"exchange/internal/repository/filesystem"
	"exchange/internal/services"
)

const testFilePath = "test_path.txt"

type Services struct {
	currecnyService domain.ICurrencyService
	notifyService   domain.INotificationService
	userService     domain.IUserService
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

	currecnyService := services.NewCurrencyService(stubs)
	userService := services.NewUserService(fileRepo)
	notificationService := services.NewNotificationService(fileRepo, currecnyService, stubs)

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

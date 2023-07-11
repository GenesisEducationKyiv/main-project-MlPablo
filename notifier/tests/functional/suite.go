package functional

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"

	"notifier/internal/infrastructure/repository/filesystem"
	"notifier/internal/services/event"
	"notifier/internal/services/user"
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
}

func (suite *Suite) SetupSuite() {
	fileRepo, err := filesystem.NewFileSystemRepository(&filesystem.Config{
		Path: testFilePath,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	stubs := new(thirdParyStubs)

	userService := user.NewUserService(fileRepo)
	notificationService := event.NewNotificationService(fileRepo, stubs, stubs)

	srvs := &Services{
		userService:   userService,
		notifyService: notificationService,
	}

	suite.srv = srvs
	suite.repo = fileRepo
}

func (suite *Suite) AfterTest(_, _ string) {
	os.Remove(testFilePath)
}

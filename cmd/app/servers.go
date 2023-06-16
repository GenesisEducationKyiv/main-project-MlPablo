package app

import (
	"os"

	"github.com/sirupsen/logrus"

	"exchange/pkg/echo"
)

type Servers struct {
	HTTPServer *echo.Server
}

func (servers *Servers) Stop() {
	if err := servers.HTTPServer.Stop(); err != nil {
		logrus.Error(err)
	}
}

func createServers() (*Servers, error) {
	httpServer, err := echo.New(&echo.Config{
		Address: os.Getenv("SERVER_ADDR"),
	})
	if err != nil {
		return nil, err
	}

	return &Servers{
		HTTPServer: httpServer,
	}, nil
}

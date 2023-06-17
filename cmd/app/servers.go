package app

import (
	"github.com/sirupsen/logrus"

	"exchange/pkg/echo"
	"exchange/utils"
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
		Address: utils.TryGetEnvDefault[string]("SERVER_ADDR", "8080"),
	})
	if err != nil {
		return nil, err
	}

	return &Servers{
		HTTPServer: httpServer,
	}, nil
}

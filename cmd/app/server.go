package app

import (
	"github.com/sirupsen/logrus"

	echo_server "exchange/pkg/echo"
	"exchange/utils"
)

type Servers struct {
	HTTPServer *echo_server.Server
}

func (servers *Servers) Stop() {
	if err := servers.HTTPServer.Stop(); err != nil {
		logrus.Error(err)
	}
}

func createServers() (*Servers, error) {
	httpServer, err := echo_server.New(&echo_server.Config{
		Address: utils.TryGetEnvDefault("SERVER_ADDR", "8080"),
	})
	if err != nil {
		return nil, err
	}

	return &Servers{
		HTTPServer: httpServer,
	}, nil
}

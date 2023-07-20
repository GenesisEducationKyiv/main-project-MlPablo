package echoserver

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *Server) Start(errChan chan error) {
	logrus.Info("Starting server...")

	if err := s.e.Start(":" + s.c.Address); !errors.Is(err, http.ErrServerClosed) {
		errChan <- err
	}
}

func (s *Server) Stop() error {
	return s.e.Close()
}

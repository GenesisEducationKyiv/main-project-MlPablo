package server

import (
	"github.com/sirupsen/logrus"
)

func (s *Server) Run(errChan chan error) {
	logrus.Info("Starting gRPC server...")

	go func() {
		if err := s.Server.Serve(s.netListener); err != nil {
			errChan <- err
		}
	}()

	logrus.Info("gRPC server started")
}

func (s *Server) Stop() {
	logrus.Info("Stopping gRPC server...")
	s.Server.Stop()
	logrus.Info("gRPC server stopped")
}

package echoserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Server struct {
	e *echo.Echo
	c *Config
}

func New(cfg *Config) (*Server, error) {
	e := echo.New()
	e.Use(getServerLogger())

	s := &Server{e: e, c: cfg}

	return s, nil
}

func getServerLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(_ echo.Context, values middleware.RequestLoggerValues) error {
			logrus.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	})
}

func (s *Server) GetEchoServer() *echo.Echo {
	return s.e
}

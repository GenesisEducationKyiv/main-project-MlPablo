package echoserver

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	})
}

func (s *Server) GetEchoServer() *echo.Echo {
	return s.e
}

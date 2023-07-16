package http

import (
	"github.com/labstack/echo/v4"

	"notifier/internal/services/event"
	"notifier/internal/services/user"
)

type Services struct {
	UserService         user.IUserService
	NotificationService event.INotificationService
}

type exchangeHandler struct {
	services *Services
	c        *Config
}

func RegisterHandlers(e *echo.Echo, services *Services, c *Config) {
	handler := &exchangeHandler{
		services: services,
		c:        c,
	}

	group := e.Group("/api")

	group.POST("/subscribe", handler.CreateMailSubscriberSaga)
	group.POST("/add_email", handler.CreateMailSubscriber)
	group.POST("/add_email_compensate", handler.CreateMailCompensate)
	// dtmutil.DefaultHTTPServer
	group.POST("/sendEmails", handler.SendEmails)
}

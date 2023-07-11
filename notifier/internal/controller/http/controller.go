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
}

func RegisterHandlers(e *echo.Echo, services *Services) {
	handler := &exchangeHandler{
		services: services,
	}

	group := e.Group("/api")

	group.POST("/subscribe", handler.CreateMailSubscriber)
	group.POST("/sendEmails", handler.SendEmails)
}

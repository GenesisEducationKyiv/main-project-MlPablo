package http

import (
	"context"

	"github.com/labstack/echo/v4"

	"notifier/internal/domain/notification"
	"notifier/internal/domain/user"
)

//go:generate mockgen -source=controller.go -destination=mocks/controller.go

type INotificationService interface {
	Notify(ctx context.Context, n *notification.Notification) error
}

type IUserService interface {
	NewUser(ctx context.Context, eu *user.User) error
}

type Services struct {
	UserService         IUserService
	NotificationService INotificationService
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

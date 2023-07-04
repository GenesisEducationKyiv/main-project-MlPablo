package http

import (
	"context"

	"github.com/labstack/echo/v4"

	"exchange/internal/domain/notification"
	"exchange/internal/domain/rate"
	"exchange/internal/domain/user"
)

//go:generate mockgen -source=controller.go -destination=mocks/controller.go

type INotificationService interface {
	Notify(ctx context.Context, n *notification.Notification) error
}

type ICurrencyService interface {
	GetCurrency(ctx context.Context, c *rate.Rate) (float64, error)
}

type IUserService interface {
	NewUser(ctx context.Context, eu *user.User) error
}

type Services struct {
	CurrencyService     ICurrencyService
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

	group.GET("/rate", handler.GetBtcToUahCurrency)
	group.POST("/subscribe", handler.CreateMailSubscriber)
	group.POST("/sendEmails", handler.SendEmails)
}

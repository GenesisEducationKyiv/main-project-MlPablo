package http

import (
	"github.com/labstack/echo/v4"

	"exchange/internal/services/currency"
	"exchange/internal/services/event"
	user_service "exchange/internal/services/user"
)

//go:generate mockgen -source=controller.go -destination=mocks/controller.go

type Services struct {
	CurrencyService     currency.ICurrencyService
	UserService         user_service.IUserService
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

	group.GET("/rate", handler.GetBtcToUahCurrency)
	group.POST("/subscribe", handler.CreateMailSubscriber)
	group.POST("/sendEmails", handler.SendEmails)
}

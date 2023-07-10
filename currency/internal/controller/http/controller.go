package http

import (
	"context"

	"github.com/labstack/echo/v4"

	"currency/internal/domain/rate"
)

//go:generate mockgen -source=controller.go -destination=mocks/controller.go

type ICurrencyService interface {
	GetCurrency(ctx context.Context, c *rate.Rate) (*rate.Currency, error)
}

type Services struct {
	CurrencyService ICurrencyService
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
}

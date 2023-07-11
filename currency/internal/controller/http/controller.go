package http

import (
	"github.com/labstack/echo/v4"

	"currency/internal/services/currency"
)

type Services struct {
	CurrencyService currency.ICurrencyService
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

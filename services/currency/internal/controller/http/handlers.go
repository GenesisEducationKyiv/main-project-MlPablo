package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"currency/internal/domain/rate"
)

func (e *exchangeHandler) GetBtcToUahCurrency(c echo.Context) error {
	ctx := c.Request().Context()

	cur := rate.GetBitcoinToUAH()

	resp, err := e.services.CurrencyService.GetCurrency(ctx, cur)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, resp)
}

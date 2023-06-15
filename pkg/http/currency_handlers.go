package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"exchange/pkg"
	"exchange/pkg/domain"
)

type exchangeHandler struct {
	services *pkg.Services
}

func NewCurrencyHandler(e *echo.Echo, services *pkg.Services) {
	handler := &exchangeHandler{
		services: services,
	}

	group := e.Group("/api")

	group.GET("/rate", handler.GetBtcToUahCurrency)
	group.POST("/subscribe", handler.CreateMailSubscriber)
	group.POST("/sendEmails", handler.SendEmails)
}

func (e *exchangeHandler) GetBtcToUahCurrency(c echo.Context) error {
	ctx := c.Request().Context()

	cur := domain.GetBitcoinToUAH()

	resp, err := e.services.CurrencyService.GetCurrency(ctx, cur)
	if err != nil {
		return c.JSON(getStatusCode(err), nil)
	}

	return c.JSON(http.StatusOK, resp)
}

// Due to API, we can't send an error on this response.
// goroutine here to do non-waiting operations and just log if the error had been occurred.
func (e *exchangeHandler) SendEmails(c echo.Context) error {
	go func() {
		if err := e.services.NotificationService.Notify(
			context.Background(),
			domain.DefaultNotification(),
		); err != nil {
			logrus.Errorf("error on sending emails: %v", err)
		}
	}()

	return c.JSON(http.StatusOK, nil)
}

// In API there was nothing about invalid requests,
// but I add validation to prevent invalid or empty mail requests.
func (e *exchangeHandler) CreateMailSubscriber(c echo.Context) error {
	ctx := c.Request().Context()

	email := domain.NewEmailUser(c.FormValue("email"))

	if err := email.Validate(); err != nil {
		return c.JSON(getStatusCode(err), nil)
	}

	if err := e.services.EmailUserService.NewEmailUser(ctx, email); err != nil {
		return c.JSON(getStatusCode(err), nil)
	}

	return c.JSON(http.StatusOK, nil)
}

// based on the error we define the response status code.
func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch {
	case errors.Is(err, domain.ErrInternalServer):
		return http.StatusInternalServerError
	case errors.Is(err, domain.ErrAlreadyExist):
		return http.StatusConflict
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrBadRequest) || errors.Is(err, domain.ErrInvalidStatus):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

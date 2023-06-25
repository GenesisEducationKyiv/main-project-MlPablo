package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"exchange/internal/domain/event"
	"exchange/internal/domain/rate"
	"exchange/internal/domain/user"
)

type Services struct {
	CurrencyService     rate.ICurrencyService
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

	group.GET("/rate", handler.GetBtcToUahCurrency)
	group.POST("/subscribe", handler.CreateMailSubscriber)
	group.POST("/sendEmails", handler.SendEmails)
}

func (e *exchangeHandler) GetBtcToUahCurrency(c echo.Context) error {
	ctx := c.Request().Context()

	cur := rate.GetBitcoinToUAH()

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
			event.DefaultNotification(),
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

	user := user.NewUser(c.FormValue("email"))

	if err := user.Validate(); err != nil {
		return c.JSON(getStatusCode(err), nil)
	}

	if err := e.services.UserService.NewUser(ctx, user); err != nil {
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
	case errors.Is(err, user.ErrInternalServer):
		return http.StatusInternalServerError
	case errors.Is(err, user.ErrAlreadyExist):
		return http.StatusConflict
	case errors.Is(err, user.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, user.ErrBadRequest) || errors.Is(err, user.ErrInvalidStatus):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"exchange/internal/domain/notification"
	"exchange/internal/domain/rate"
	"exchange/internal/domain/user"
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

// Due to API, we can't send an error on this response.
// goroutine here to do non-waiting operations and just log if the error had been occurred.
func (e *exchangeHandler) SendEmails(c echo.Context) error {
	go func() {
		if err := e.services.NotificationService.Notify(
			context.Background(),
			notification.DefaultNotification(),
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

	u := user.NewUser(c.FormValue("email"))

	if err := u.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := e.services.UserService.NewUser(ctx, u); err != nil {
		code := http.StatusInternalServerError
		if errors.Is(err, user.ErrAlreadyExist) {
			code = http.StatusConflict
		}

		return c.JSON(code, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"notifier/internal/domain/notification"
	"notifier/internal/domain/user"
)

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
func (e *exchangeHandler) CreateMailSubscriberSaga(c echo.Context) error {
	ctx := c.Request().Context()
	u := user.NewUser(c.FormValue("email"))

	if err := u.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	ok, err := e.services.UserService.UserExist(ctx, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if ok {
		return c.JSON(http.StatusConflict, user.ErrAlreadyExist)
	}

	globalID := dtmcli.MustGenGid(e.c.DtmCoordinatoURL)

	err = dtmcli.
		NewSaga(e.c.DtmCoordinatoURL, globalID).
		Add(e.c.NotifierServerURL+"/api/add_email", e.c.NotifierServerURL+"/api/add_email_compensate", u).
		Add(e.c.CustomerServerURL+"/create", "", nil).
		Submit() // customer service doesn't have compensate path for create user
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, globalID)
}

func (e *exchangeHandler) CreateMailSubscriber(c echo.Context) error {
	ctx := c.Request().Context()
	u := new(user.User)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

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

func (e *exchangeHandler) CreateMailCompensate(c echo.Context) error {
	ctx := c.Request().Context()

	u := new(user.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//
	if err := u.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err := e.services.UserService.DeleteUser(ctx, u); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

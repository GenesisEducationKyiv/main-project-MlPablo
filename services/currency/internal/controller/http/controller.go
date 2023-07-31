package http

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"currency/internal/domain/events"
	"currency/internal/services/currency"
)

type Eventer interface {
	Publish(topic, value string) error
}

type Services struct {
	CurrencyService currency.ICurrencyService
}

type exchangeHandler struct {
	services *Services
	event    Eventer
}

func RegisterHandlers(e *echo.Echo, services *Services, event Eventer) {
	handler := &exchangeHandler{
		services: services,
		event:    event,
	}

	group := e.Group("/api")

	group.GET("/rate", handler.EventMiddleware(handler.GetBtcToUahCurrency))
}

func (h *exchangeHandler) EventMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logrus.WithFields(logrus.Fields{
			"URI": c.Request().URL.String(),
		})
		log.Level = logrus.InfoLevel
		strLog, _ := log.String()

		if err := h.event.Publish(events.LOG_EVENT, strLog); err != nil {
			logrus.Info("error on publish event")
		}

		log.Info()

		return next(c)
	}
}

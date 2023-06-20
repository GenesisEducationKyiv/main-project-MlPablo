package app

import (
	_http "exchange/internal/controller/http"
)

func registerHandlers(server *Servers, service *Services) {
	_http.RegisterHandlers(server.HTTPServer.GetEchoServer(), &_http.Services{
		CurrencyService:     service.CurrencyService,
		UserService:         service.UserService,
		NotificationService: service.NotificationService,
	})
}

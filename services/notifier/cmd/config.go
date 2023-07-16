package main

import (
	"notifier/internal/controller/http"
	"notifier/internal/infrastructure/mail"
	"notifier/internal/infrastructure/repository/filesystem"
	echoserver "notifier/pkg/echo"
	"notifier/pkg/grpc/client"
	"notifier/utils"
)

var envGet = utils.TryGetEnv[string] //nolint: gochecknoglobals// ok here

func NewEchoConfig() *echoserver.Config {
	return &echoserver.Config{
		Address: utils.TryGetEnvDefault("HTTP_SERVER_ADDR", "8080"),
	}
}

func NewMailConfig() *mail.Config {
	return mail.NewConfig(
		envGet("EMAIL_LOGIN"),
		envGet("EMAIL_APP_PASSWORD"),
		envGet("SMTP_HOST"),
		envGet("SMTP_PORT"),
	)
}

func NewHTTPControllerConfig() *http.Config {
	return &http.Config{
		DtmCoordinatoURL:  envGet("DTM_COORDINATOR"),
		CustomerServerURL: envGet("CUSTOMERS_SERVICE_URL"),
		NotifierServerURL: envGet("NOTIFIER_SERVICE_URL"),
	}
}

func NewFileSystemConfig() *filesystem.Config {
	return &filesystem.Config{Path: utils.TryGetEnvDefault("FILE_STORE_PATH", "./file_storage.txt")}
}

func NewCurrencyGrpcConfig() *client.Config {
	return &client.Config{
		Address: envGet("GRPC_CURRENCY_ADDRESS"),
		Port:    envGet("GRPC_PORT"),
	}
}

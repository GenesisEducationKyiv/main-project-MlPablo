package main

import (
	"exchange/internal/infrastructure/currency/binance"
	"exchange/internal/infrastructure/currency/coingecko"
	"exchange/internal/infrastructure/currency/currencyapi"
	"exchange/internal/infrastructure/mail"
	"exchange/internal/infrastructure/repository/filesystem"
	echoserver "exchange/pkg/echo"
	"exchange/utils"
)

var envGet = utils.TryGetEnv[string] //nolint: gochecknoglobals// ok here

func NewEchoConfig() *echoserver.Config {
	return &echoserver.Config{
		Address: utils.TryGetEnvDefault("SERVER_ADDR", "8080"),
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

func NewCurrencyapiConfig() *currencyapi.Config {
	return currencyapi.NewConfig(
		envGet("CURR_API_KEY"),
		envGet("CURR_URL"),
	)
}

func NewCoingeckoConfig() *coingecko.Config {
	return coingecko.NewConfig(
		envGet("COINGECKO_URL"),
	)
}

func NewFileSystemConfig() *filesystem.Config {
	return &filesystem.Config{Path: utils.TryGetEnvDefault("FILE_STORE_PATH", "./file_storage.txt")}
}

func NewBinanceConfig() *binance.Config {
	return binance.NewConfig(envGet("BINANCE_URL"))
}

package main

import (
	"currency/internal/infrastructure/currency/binance"
	"currency/internal/infrastructure/currency/coingecko"
	"currency/internal/infrastructure/currency/currencyapi"
	echoserver "currency/pkg/echo"
	"currency/utils"
)

var envGet = utils.TryGetEnv[string] //nolint: gochecknoglobals// ok here

func NewEchoConfig() *echoserver.Config {
	return &echoserver.Config{
		Address: utils.TryGetEnvDefault("SERVER_ADDR", "8080"),
	}
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

func NewBinanceConfig() *binance.Config {
	return binance.NewConfig(envGet("BINANCE_URL"))
}
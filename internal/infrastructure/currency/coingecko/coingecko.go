package coingecko

import (
	"net/http"

	"exchange/internal/domain/rate"
)

type CoingeckoAPI struct {
	cfg    *Config
	cli    *http.Client
	mapper map[string]string
}

func NewCoingeckoApi(cfg *Config) *CoingeckoAPI {
	return &CoingeckoAPI{
		cfg:    cfg,
		cli:    http.DefaultClient,
		mapper: initMapper(),
	}
}

func initMapper() map[string]string {
	coins := make(map[string]string)

	coins[rate.BTC] = "bitcoin"

	return coins
}

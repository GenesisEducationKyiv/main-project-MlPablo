package coingecko

import (
	"net/http"

	"exchange/internal/domain/rate"
	"exchange/internal/infrastructure/currency"
)

type CoingeckoAPI struct {
	cfg    *Config
	cli    *http.Client
	mapper map[string]string
	next   currency.IChain
}

func NewCoingeckoApi(cfg *Config, opts ...Option) *CoingeckoAPI {
	api := &CoingeckoAPI{
		cfg:    cfg,
		cli:    &http.Client{Transport: http.DefaultTransport},
		mapper: initMapper(),
	}

	for _, opt := range opts {
		opt(api)
	}

	return api
}

func initMapper() map[string]string {
	coins := make(map[string]string)

	coins[rate.BTC] = "bitcoin"

	return coins
}

func (api *CoingeckoAPI) SetNext(chain currency.IChain) {
	api.next = chain
}

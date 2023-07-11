package binance

import (
	"net/http"

	"exchange/internal/infrastructure/currency"
)

type BinanceAPI struct {
	cfg  *Config
	cli  *http.Client
	next currency.IChain
}

func NewBinanceApi(cfg *Config, opts ...Option) *BinanceAPI {
	api := &BinanceAPI{
		cfg: cfg,
		cli: &http.Client{Transport: http.DefaultTransport},
	}

	for _, opt := range opts {
		opt(api)
	}

	return api
}

func (api *BinanceAPI) SetNext(chain currency.IChain) {
	api.next = chain
}

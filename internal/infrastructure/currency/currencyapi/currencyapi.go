package currencyapi

import (
	"net/http"

	"exchange/internal/infrastructure/currency"
)

type CurrencyAPI struct {
	cfg  *Config
	cli  *http.Client
	next currency.IChain
}

// This is the implementation of logic that can get currency.
// So service doesn't need to know about how we do this, and we can implement any currency api and interfaces we want
// I'm not sure about putting this into infrastructure folder.
func NewCurrencyAPI(cfg *Config, opts ...Option) *CurrencyAPI {
	api := &CurrencyAPI{
		cfg: cfg,
		cli: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(api)
	}

	return api
}

func (api *CurrencyAPI) SetNext(chain currency.IChain) {
	api.next = chain
}

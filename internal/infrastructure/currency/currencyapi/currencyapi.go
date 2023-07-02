package currencyapi

import (
	"context"
	"errors"
	"net/http"

	"exchange/internal/domain/rate"
)

//go:generate mockgen -source=currencyapi.go -destination=mocks/currencyapi.go

type Chain interface {
	GetCurrency(ctx context.Context, cur *rate.Rate) (float64, error)
	SetNext(any) error
}

type CurrencyAPI struct {
	cfg  *Config
	cli  *http.Client
	next Chain
}

// This is the implementation of logic that can get currency.
// So service doesn't need to know about how we do this, and we can implement any currency api and interfaces we want
// I'm not sure about putting this into infrastructure folder.
func NewCurrencyAPI(cfg *Config, opts ...Option) *CurrencyAPI {
	api := &CurrencyAPI{
		cfg: cfg,
		cli: &http.Client{Transport: http.DefaultTransport},
	}

	for _, opt := range opts {
		opt(api)
	}

	return api
}

func (api *CurrencyAPI) SetNext(chain any) error {
	v, ok := chain.(Chain)
	if !ok {
		return errors.New("unable to set next handler. Invalid type")
	}

	api.next = v

	return nil
}

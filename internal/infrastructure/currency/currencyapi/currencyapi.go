package currencyapi

import (
	"context"
	"errors"

	"exchange/internal/domain/rate"
)

type Chain interface {
	GetCurrency(ctx context.Context, cur *rate.Rate) (float64, error)
	SetNext(any) error
}

type CurrencyAPI struct {
	cfg  *Config
	next Chain
}

// This is the implementation of logic that can get currency.
// So service doesn't need to know about how we do this, and we can implement any currency api and interfaces we want
// I'm not sure about putting this into infrastructure folder.
func NewCurrencyAPI(cfg *Config) *CurrencyAPI {
	return &CurrencyAPI{
		cfg: cfg,
	}
}

func (api *CurrencyAPI) SetNext(chain any) error {
	v, ok := chain.(Chain)
	if !ok {
		return errors.New("unable to set next handler. Invalid type")
	}

	api.next = v

	return nil
}

package binance

import (
	"context"
	"errors"
	"net/http"

	"exchange/internal/domain/rate"
)

//go:generate mockgen -source=binance.go -destination=mocks/binance.go

type Chain interface {
	GetCurrency(ctx context.Context, cur *rate.Rate) (float64, error)
	SetNext(any) error
}

type BinanceAPI struct {
	cfg  *Config
	cli  *http.Client
	next Chain
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

func (api *BinanceAPI) SetNext(chain any) error {
	v, ok := chain.(Chain)
	if !ok {
		return errors.New("unable to set next handler. Invalid type")
	}

	api.next = v

	return nil
}

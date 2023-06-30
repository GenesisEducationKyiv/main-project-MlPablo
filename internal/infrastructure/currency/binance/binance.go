package binance

import (
	"context"
	"errors"
	"net/http"

	"exchange/internal/domain/rate"
)

type Chain interface {
	GetCurrency(ctx context.Context, cur *rate.Rate) (float64, error)
	SetNext(any) error
}

type BinanceAPI struct {
	cfg  *Config
	cli  *http.Client
	next Chain
}

func NewBinanceApi(cfg *Config) *BinanceAPI {
	return &BinanceAPI{
		cfg: cfg,
		cli: http.DefaultClient,
	}
}

func (api *BinanceAPI) SetNext(chain any) error {
	v, ok := chain.(Chain)
	if !ok {
		return errors.New("unable to set next handler. Invalid type")
	}

	api.next = v

	return nil
}

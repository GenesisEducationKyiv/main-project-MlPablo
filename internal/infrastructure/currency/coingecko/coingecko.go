package coingecko

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

type CoingeckoAPI struct {
	cfg    *Config
	cli    *http.Client
	mapper map[string]string
	next   Chain
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

func (api *CoingeckoAPI) SetNext(chain any) error {
	v, ok := chain.(Chain)
	if !ok {
		return errors.New("unable to set next handler. Invalid type")
	}

	api.next = v

	return nil
}

package binance

import (
	"net/http"
)

type BinanceAPI struct {
	cfg *Config
	cli *http.Client
}

func NewBinanceApi(cfg *Config) *BinanceAPI {
	return &BinanceAPI{
		cfg: cfg,
		cli: http.DefaultClient,
	}
}

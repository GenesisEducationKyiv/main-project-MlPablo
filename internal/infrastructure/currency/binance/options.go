package binance

import (
	"net/http"
)

type (
	Option func(*BinanceAPI)
)

func WithClient(client *http.Client) Option {
	return func(b *BinanceAPI) {
		b.cli = client
	}
}

package coingecko

import (
	"net/http"
)

type (
	Option func(*CoingeckoAPI)
)

func WithClient(client *http.Client) Option {
	return func(b *CoingeckoAPI) {
		b.cli = client
	}
}

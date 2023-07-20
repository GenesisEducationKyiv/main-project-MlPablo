package currencyapi

import (
	"net/http"
)

type (
	Option func(*CurrencyAPI)
)

func WithClient(client *http.Client) Option {
	return func(b *CurrencyAPI) {
		b.cli = client
	}
}

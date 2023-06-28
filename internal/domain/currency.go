package domain

import "context"

//go:generate mockgen -source=currency.go -destination=mocks/currency.go

const (
	UAH = "UAH"
	BTC = "BTC"
)

// Currency domain that response to all currency operations.
// So we can easily add any route on any currency.
type Currency struct {
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

func GetBitcoinToUAH() *Currency {
	return &Currency{
		BaseCurrency:  BTC,
		QuoteCurrency: UAH,
	}
}

type ICurrencyService interface {
	GetCurrency(ctx context.Context, c *Currency) (float64, error)
}

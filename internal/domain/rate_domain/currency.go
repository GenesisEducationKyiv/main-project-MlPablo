package rate_domain

//go:generate mockgen -source=currency.go -destination=mocks/currency.go

const (
	UAH = "UAH"
	BTC = "BTC"
)

// Rate domain that response to all currency operations.
// So we can easily add any route on any currency.
type Rate struct {
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

func GetBitcoinToUAH() *Rate {
	return &Rate{
		BaseCurrency:  BTC,
		QuoteCurrency: UAH,
	}
}

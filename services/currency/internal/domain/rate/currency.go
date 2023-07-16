package rate

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

type Currency struct {
	Value float64 `json:"value"`
}

func NewCurrency(value float64) *Currency {
	return &Currency{
		Value: value,
	}
}

func GetBitcoinToUAH() *Rate {
	return &Rate{
		BaseCurrency:  BTC,
		QuoteCurrency: UAH,
	}
}

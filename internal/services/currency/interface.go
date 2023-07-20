package currency

import (
	"context"

	"exchange/internal/domain/rate"
)

type ICurrencyService interface {
	GetCurrency(ctx context.Context, c *rate.Rate) (*rate.Currency, error)
}

package currency

import (
	"context"

	"currency/internal/domain/rate"
)

//go:generate mockgen -source=interface.go -destination=mocks/interface.go

type ICurrencyService interface {
	GetCurrency(ctx context.Context, c *rate.Rate) (*rate.Currency, error)
}

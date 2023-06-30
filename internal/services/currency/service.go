package currency

import (
	"context"

	rate_domain "exchange/internal/domain/rate"
)

//go:generate mockgen -source=service.go -destination=mocks/currency.go

type ICurrencyAPI interface {
	GetCurrency(ctx context.Context, data *rate_domain.Rate) (float64, error)
}

type Service struct {
	currencyAPI ICurrencyAPI
}

func NewCurrencyService(
	currencyAPI ICurrencyAPI,
) *Service {
	return &Service{
		currencyAPI: currencyAPI,
	}
}

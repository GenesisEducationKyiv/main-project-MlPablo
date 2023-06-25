package services

import (
	"context"

	"exchange/internal/domain/rate"
)

//go:generate mockgen -source=currency.go -destination=mocks/currency.go

type ICurrencyAPI interface {
	GetCurrency(ctx context.Context, data *rate.Rate) (float64, error)
}

type currencyService struct {
	currencyAPI ICurrencyAPI
}

func NewCurrencyService(
	currencyAPI ICurrencyAPI,
) rate.ICurrencyService {
	return &currencyService{
		currencyAPI: currencyAPI,
	}
}

func (n *currencyService) GetCurrency(ctx context.Context, data *rate.Rate) (float64, error) {
	return n.currencyAPI.GetCurrency(ctx, data)
}
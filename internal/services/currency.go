package services

import (
	"context"

	"exchange/internal/domain"
)

//go:generate mockgen -source=currency.go -destination=mocks/currency.go

type ICurrencyAPI interface {
	GetCurrency(ctx context.Context, data *domain.Currency) (float64, error)
}

type currencyService struct {
	currencyAPI ICurrencyAPI
}

func NewCurrencyService(
	currencyAPI ICurrencyAPI,
) domain.ICurrencyService {
	return &currencyService{
		currencyAPI: currencyAPI,
	}
}

func (n *currencyService) GetCurrency(ctx context.Context, data *domain.Currency) (float64, error) {
	return n.currencyAPI.GetCurrency(ctx, data)
}

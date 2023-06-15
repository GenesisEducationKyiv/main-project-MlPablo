package mock

import (
	"context"

	"exchange/pkg/domain"
)

const mockPrice = 10

type CurrencyService struct{}

func (c *CurrencyService) GetCurrency(_ context.Context, _ *domain.Currency) (float64, error) {
	return mockPrice, nil
}

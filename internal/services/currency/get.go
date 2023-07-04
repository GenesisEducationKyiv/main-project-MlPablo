package currency

import (
	"context"

	"exchange/internal/domain/rate"
)

func (s *Service) GetCurrency(ctx context.Context, data *rate.Rate) (float64, error) {
	return s.currencyAPI.GetCurrency(ctx, data)
}

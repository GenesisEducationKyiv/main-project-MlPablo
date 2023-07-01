package currency

import (
	"context"

	rate_domain "exchange/internal/domain/rate"
)

func (s *Service) GetCurrency(ctx context.Context, data *rate_domain.Rate) (float64, error) {
	return s.currencyAPI.GetCurrency(ctx, data)
}

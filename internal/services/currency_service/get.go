package currency_service

import (
	"context"

	"exchange/internal/domain/rate_domain"
)

func (s *Service) GetCurrency(ctx context.Context, data *rate_domain.Rate) (float64, error) {
	return s.currencyAPI.GetCurrency(ctx, data)
}

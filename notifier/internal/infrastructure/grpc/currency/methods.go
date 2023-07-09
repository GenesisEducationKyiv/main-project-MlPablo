package currency

import (
	"context"

	"notifier/internal/domain/rate"
)

func (s *Service) GetCurrency(ctx context.Context, r *rate.Rate) (*rate.Currency, error) {
	res, err := s.cli.GetCurrency(ctx, rateDomainConvert(r))
	if err != nil {
		return nil, err
	}

	return currencyGrpcConvert(res), nil
}

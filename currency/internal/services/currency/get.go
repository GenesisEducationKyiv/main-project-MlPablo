package currency

import (
	"context"

	"currency/internal/domain/rate"
)

func (s *Service) GetCurrency(ctx context.Context, data *rate.Rate) (*rate.Currency, error) {
	value, err := s.currencyAPI.GetCurrency(ctx, data)
	if err != nil {
		return nil, err
	}

	return &rate.Currency{Value: value}, nil
}

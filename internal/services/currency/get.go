package currency

import (
	"context"

	"exchange/internal/domain/rate"
)

func (s *Service) GetCurrency(ctx context.Context, data *rate.Rate) (*rate.Currency, error) {
	value, err := s.currencyAPI.GetCurrency(ctx, data)
	if err != nil {
		return nil, err
	}

	return &rate.Currency{Data: value}, nil
}

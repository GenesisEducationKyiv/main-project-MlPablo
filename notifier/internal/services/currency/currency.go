package currency

import (
	"context"

	"notifier/internal/domain/rate"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) GetCurrency(ctx context.Context, r *rate.Rate) (*rate.Currency, error) {
	return &rate.Currency{Value: 10000}, nil
}

package functional

import (
	"context"
	"math/rand"

	"exchange/internal/domain"
)

type thirdParyStubs struct{}

func (c *thirdParyStubs) GetCurrency(_ context.Context, _ *domain.Currency) (float64, error) {
	return rand.Float64(), nil //nolint:gosec // math/rand is ok here
}

func (c *thirdParyStubs) SendEmail(_ context.Context, _ any, _ ...string) error {
	return nil
}

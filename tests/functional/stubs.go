package functional

import (
	"context"
	"math/rand"

	"exchange/internal/domain/rate_domain"
)

type thirdParyStubs struct{}

func (c *thirdParyStubs) GetCurrency(_ context.Context, _ *rate_domain.Rate) (float64, error) {
	return rand.Float64(), nil //nolint:gosec // math/rand is ok here
}

func (c *thirdParyStubs) SendEmail(_ context.Context, _ any, _ ...string) error {
	return nil
}

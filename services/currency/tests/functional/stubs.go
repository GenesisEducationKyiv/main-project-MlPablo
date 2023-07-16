package functional

import (
	"context"
	"math/rand"

	"currency/internal/domain/rate"
)

type thirdParyStubs struct{}

func (c *thirdParyStubs) GetCurrency(_ context.Context, _ *rate.Rate) (float64, error) {
	return rand.Float64(), nil //nolint:gosec // math/rand is ok here
}

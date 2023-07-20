package currency

import (
	"context"

	"currency/internal/domain/rate"
)

//go:generate mockgen -source=chain.go -destination=mocks/chain.go

type ICryptoProvider interface {
	GetCurrency(ctx context.Context, data *rate.Rate) (float64, error)
}

type IChain interface {
	ICryptoProvider
	SetNext(IChain)
}

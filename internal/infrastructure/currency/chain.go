package currency

import (
	"exchange/internal/services/currency"
)

//go:generate mockgen -source=chain.go -destination=mocks/chain.go

type IChain interface {
	currency.ICurrencyAPI
	SetNext(IChain)
}

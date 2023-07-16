package grpc

import (
	"currency/api/proto/grpc_currency_service"
	"currency/internal/domain/rate"
)

func rateGrpcConvert(r *grpc_currency_service.Rate) *rate.Rate {
	return &rate.Rate{
		BaseCurrency:  r.BaseCurrency,
		QuoteCurrency: r.QuoteCurrency,
	}
}

func currencyDomainConvert(r *rate.Currency) *grpc_currency_service.Currency {
	return &grpc_currency_service.Currency{
		Value: r.Value,
	}
}

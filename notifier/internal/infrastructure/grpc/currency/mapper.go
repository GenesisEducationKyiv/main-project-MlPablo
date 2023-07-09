package currency

import (
	"notifier/api/proto/grpc_currency_service"
	"notifier/internal/domain/rate"
)

func rateDomainConvert(r *rate.Rate) *grpc_currency_service.Rate {
	return &grpc_currency_service.Rate{
		BaseCurrency:  r.BaseCurrency,
		QuoteCurrency: r.QuoteCurrency,
	}
}

func currencyGrpcConvert(
	r *grpc_currency_service.Currency,
) *rate.Currency {
	return &rate.Currency{
		Value: r.Value,
	}
}

package grpc

import (
	"context"

	"google.golang.org/grpc"

	"currency/api/proto/grpc_currency_service"
	"currency/internal/services/currency"
)

type Services struct {
	CurrencyService currency.ICurrencyService
}

type exchangeHandler struct {
	services *Services
}

func RegisterHandlers(g grpc.ServiceRegistrar, services *Services) {
	handler := &exchangeHandler{
		services: services,
	}
	grpc_currency_service.RegisterCurrencyServiceServer(g, handler)
}

func (h *exchangeHandler) GetCurrency(
	ctx context.Context,
	req *grpc_currency_service.Rate,
) (*grpc_currency_service.Currency, error) {
	res, err := h.services.CurrencyService.GetCurrency(ctx, rateGrpcConvert(req))
	if err != nil {
		return nil, err
	}

	return currencyDomainConvert(res), nil
}

package currency

import (
	"google.golang.org/grpc"

	"notifier/api/proto/grpc_currency_service"
)

type Service struct {
	cli grpc_currency_service.CurrencyServiceClient
}

func New(conn *grpc.ClientConn) *Service {
	cli := grpc_currency_service.NewCurrencyServiceClient(conn)
	return &Service{cli: cli}
}

package services_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain"
	"exchange/internal/services"
	mock_services "exchange/internal/services/mocks"
)

func TestGetCurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currencyAPI := mock_services.NewMockICurrencyAPI(ctrl)

	const rate = 1_000_000.0

	currencyAPI.EXPECT().GetCurrency(context.Background(), &domain.Currency{
		BaseCurrency:  domain.BTC,
		QuoteCurrency: domain.UAH,
	}).Return(rate, nil)

	currencyServiceMock := services.NewCurrencyService(currencyAPI)

	res, err := currencyServiceMock.GetCurrency(context.Background(), &domain.Currency{
		BaseCurrency:  domain.BTC,
		QuoteCurrency: domain.UAH,
	})
	require.NoError(t, err)
	require.Equal(t, rate, res)
}

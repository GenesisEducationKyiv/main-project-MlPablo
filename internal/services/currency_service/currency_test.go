package currency_service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain/rate_domain"
	"exchange/internal/services/currency_service"
	mock_currency_service "exchange/internal/services/currency_service/mocks"
)

func TestGetCurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currencyAPI := mock_currency_service.NewMockICurrencyAPI(ctrl)

	const btcUahRate = 1_000_000.0

	currencyAPI.EXPECT().GetCurrency(context.Background(), &rate_domain.Rate{
		BaseCurrency:  rate_domain.BTC,
		QuoteCurrency: rate_domain.UAH,
	}).Return(btcUahRate, nil)

	currencyServiceMock := currency_service.NewCurrencyService(currencyAPI)

	res, err := currencyServiceMock.GetCurrency(context.Background(), &rate_domain.Rate{
		BaseCurrency:  rate_domain.BTC,
		QuoteCurrency: rate_domain.UAH,
	})
	require.NoError(t, err)
	require.Equal(t, btcUahRate, res)
}

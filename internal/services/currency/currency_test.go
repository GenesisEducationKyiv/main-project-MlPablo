package currency_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	rate_domain "exchange/internal/domain/rate"
	"exchange/internal/services/currency"
	mock_currency "exchange/internal/services/currency/mocks"
)

func TestGetCurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	currencyAPI := mock_currency.NewMockICurrencyAPI(ctrl)

	const btcUahRate = 1_000_000.0

	currencyAPI.EXPECT().GetCurrency(context.Background(), &rate_domain.Rate{
		BaseCurrency:  rate_domain.BTC,
		QuoteCurrency: rate_domain.UAH,
	}).Return(btcUahRate, nil)

	currencyServiceMock := currency.NewCurrencyService(currencyAPI)

	res, err := currencyServiceMock.GetCurrency(context.Background(), &rate_domain.Rate{
		BaseCurrency:  rate_domain.BTC,
		QuoteCurrency: rate_domain.UAH,
	})
	require.NoError(t, err)
	require.Equal(t, btcUahRate, res)
}
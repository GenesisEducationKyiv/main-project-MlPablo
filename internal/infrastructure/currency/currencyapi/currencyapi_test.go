package currencyapi

import (
	"context"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain/rate"
	mock_currency "exchange/internal/infrastructure/currency/mocks"
)

const invalidURL = "invalidurl.com"

func getApi(t *testing.T) *CurrencyAPI {
	err := godotenv.Load("../../../../.env")
	require.NoError(t, err)

	url := os.Getenv("CURR_URL")
	key := os.Getenv("CURR_API_KEY")

	require.NotZero(t, url)
	require.NotZero(t, key)

	return NewCurrencyAPI(NewConfig(key, url))
}

func TestCoingecko(t *testing.T) {
	api := getApi(t)
	res, err := api.GetCurrency(context.Background(), rate.GetBitcoinToUAH())
	require.NoError(t, err)
	require.NotZero(t, res)
}

func TestSetNext(t *testing.T) {
	api := getApi(t)

	api.cfg.baseURL = invalidURL

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const mockReturn = 1_000_000.0

	bctRate := rate.GetBitcoinToUAH()

	mockApi := mock_currency.NewMockIChain(ctrl)
	mockApi.EXPECT().
		GetCurrency(context.Background(), bctRate).
		Return(mockReturn, nil)

	api.SetNext(mockApi)

	res, err := api.GetCurrency(context.Background(), bctRate)
	require.NoError(t, err)
	require.Equal(t, mockReturn, res)
}

package binance

import (
	"context"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain/rate"
	mock_binance "exchange/internal/infrastructure/currency/binance/mocks"
)

const invalidURL = "invalidurl.com"

func getApi(t *testing.T) *BinanceAPI {
	err := godotenv.Load("../../../../.env")
	require.NoError(t, err)

	url := os.Getenv("BINANCE_URL")
	require.NotZero(t, url)

	return NewBinanceApi(NewConfig(url))
}

func TestBinance(t *testing.T) {
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

	mockApi := mock_binance.NewMockChain(ctrl)
	mockApi.EXPECT().
		GetCurrency(context.Background(), bctRate).
		Return(mockReturn, nil)

	err := api.SetNext(mockApi)
	require.NoError(t, err)

	res, err := api.GetCurrency(context.Background(), bctRate)
	require.NoError(t, err)
	require.Equal(t, mockReturn, res)
}

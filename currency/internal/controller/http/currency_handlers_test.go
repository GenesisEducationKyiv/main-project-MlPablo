package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mock_http "currency/internal/controller/http/mocks"
	"currency/internal/domain/rate"
)

type mockServices struct {
	currencyService *mock_http.MockICurrencyService
}

func getMockedServices(ctrl *gomock.Controller) *mockServices {
	return &mockServices{
		currencyService: mock_http.NewMockICurrencyService(ctrl),
	}
}

func getMockedExchangeHandler(m *mockServices) *exchangeHandler {
	return &exchangeHandler{
		services: &Services{
			CurrencyService: m.currencyService,
		},
	}
}

func TestGetCurrency(t *testing.T) {
	tc := []struct {
		name                    string
		expectedErrFromCurrency error
		expectedRate            float64
		expectedStatusCode      int
	}{
		{
			name:                    "valid case",
			expectedErrFromCurrency: nil,
			expectedRate:            rand.Float64(),
			expectedStatusCode:      http.StatusOK,
		},
		{
			name:                    "currency service error",
			expectedErrFromCurrency: errors.New("dummyErr"),
			expectedRate:            0,
			expectedStatusCode:      http.StatusInternalServerError,
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockedServices := getMockedServices(ctrl)

			mockedServices.currencyService.EXPECT().
				GetCurrency(context.Background(), rate.GetBitcoinToUAH()).
				Return(rate.NewCurrency(test.expectedRate), test.expectedErrFromCurrency)

			e := echo.New()

			handlers := getMockedExchangeHandler(mockedServices)

			req := httptest.NewRequest(http.MethodGet, "/api/rate", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			assert.NoError(t, handlers.GetBtcToUahCurrency(c))
			assert.Equal(t, test.expectedStatusCode, rec.Code)
			if test.expectedErrFromCurrency != nil {
				return
			}

			respBody, err := io.ReadAll(rec.Body)
			require.NoError(t, err)

			var gotResponse rate.Currency

			err = json.Unmarshal(respBody, &gotResponse)
			require.NoError(t, err)
			assert.Equal(t, test.expectedRate, gotResponse.Value)
		})
	}
}

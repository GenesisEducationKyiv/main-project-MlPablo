package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mock_http "exchange/internal/controller/http/mocks"
	"exchange/internal/domain/notification"
	"exchange/internal/domain/rate"
	"exchange/internal/domain/user"
)

type mockServices struct {
	currencyService     *mock_http.MockICurrencyService
	userService         *mock_http.MockIUserService
	notificationService *mock_http.MockINotificationService
}

func getMockedServices(ctrl *gomock.Controller) *mockServices {
	return &mockServices{
		currencyService:     mock_http.NewMockICurrencyService(ctrl),
		userService:         mock_http.NewMockIUserService(ctrl),
		notificationService: mock_http.NewMockINotificationService(ctrl),
	}
}

func getMockedExchangeHandler(m *mockServices) *exchangeHandler {
	return &exchangeHandler{
		services: &Services{
			CurrencyService:     m.currencyService,
			UserService:         m.userService,
			NotificationService: m.notificationService,
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

func TestSendEmails(t *testing.T) {
	tc := []struct {
		name                            string
		expectedErrFromSendNotification error
		expectedStatusCode              int
	}{
		{
			name:                            "valid case",
			expectedErrFromSendNotification: nil,
			expectedStatusCode:              http.StatusOK,
		},
		{
			name:                            "error on send email",
			expectedErrFromSendNotification: errors.New("dummyErr"),
			expectedStatusCode:              http.StatusOK,
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockedServices := getMockedServices(ctrl)

			wg := sync.WaitGroup{}
			wg.Add(1)

			mockedServices.notificationService.EXPECT().
				Notify(context.Background(), notification.DefaultNotification()).
				Return(test.expectedErrFromSendNotification).
				Do(func(_, _ any) {
					defer wg.Done()
				})

			e := echo.New()

			handlers := getMockedExchangeHandler(mockedServices)

			req := httptest.NewRequest(http.MethodPost, "/sendEmails", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			assert.NoError(t, handlers.SendEmails(c))
			wg.Wait()

			assert.Equal(t, test.expectedStatusCode, rec.Code)
		})
	}
}

func TestCreateMailSubscriber(t *testing.T) {
	type args struct {
		email string
	}

	tc := []struct {
		name               string
		args               args
		serviceSetup       func(m *mockServices, a args)
		expectedStatusCode int
	}{
		{
			name: "valid case",
			serviceSetup: func(m *mockServices, a args) {
				m.userService.EXPECT().
					NewUser(context.Background(), user.NewUser(a.email)).
					Return(nil)
			},
			args: args{
				email: "some@gmail.com",
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:         "invalid email ",
			serviceSetup: func(m *mockServices, a args) {},
			args: args{
				email: "some.com",
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "email already exist",
			serviceSetup: func(m *mockServices, a args) {
				m.userService.EXPECT().
					NewUser(context.Background(), user.NewUser(a.email)).
					Return(user.ErrAlreadyExist)
			},
			args: args{
				email: "some@email.com",
			},
			expectedStatusCode: http.StatusConflict,
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockedServices := getMockedServices(ctrl)

			test.serviceSetup(mockedServices, test.args)

			e := echo.New()

			handlers := getMockedExchangeHandler(mockedServices)

			req := httptest.NewRequest(http.MethodPost, "/subscribe", nil)
			form, _ := url.ParseQuery(req.URL.RawQuery)
			form.Add("email", test.args.email)
			req.URL.RawQuery = form.Encode()

			rec := httptest.NewRecorder()

			c := e.NewContext(req, rec)

			assert.NoError(t, handlers.CreateMailSubscriber(c))

			assert.Equal(t, test.expectedStatusCode, rec.Code)
		})
	}
}

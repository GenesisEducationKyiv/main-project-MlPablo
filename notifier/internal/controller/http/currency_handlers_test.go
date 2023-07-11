package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"notifier/internal/domain/notification"
	"notifier/internal/domain/user"
	mock_event "notifier/internal/services/event/mocks"
	mock_user "notifier/internal/services/user/mocks"
)

type mockServices struct {
	notificationService *mock_event.MockINotificationService
	userService         *mock_user.MockIUserService
}

func getMockedServices(ctrl *gomock.Controller) *mockServices {
	return &mockServices{
		userService:         mock_user.NewMockIUserService(ctrl),
		notificationService: mock_event.NewMockINotificationService(ctrl),
	}
}

func getMockedExchangeHandler(m *mockServices) *exchangeHandler {
	return &exchangeHandler{
		services: &Services{
			UserService:         m.userService,
			NotificationService: m.notificationService,
		},
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

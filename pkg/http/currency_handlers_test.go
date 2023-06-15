package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"exchange/pkg"
	"exchange/pkg/domain/mock"
	"exchange/pkg/repository/mem"
	"exchange/pkg/services"
)

func TestGetCurrency(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/rate", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &exchangeHandler{
		services: &pkg.Services{
			CurrencyService: &mock.CurrencyService{},
		},
	}

	assert.NoError(t, h.GetBtcToUahCurrency(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestAddEmail(t *testing.T) {
	email := faker.Email()
	tc := []struct {
		name      string
		respCode  int
		email     string
		expectErr error
	}{
		{
			name:      "ValidCase",
			respCode:  http.StatusOK,
			email:     email,
			expectErr: nil,
		},
		{
			name:      "InValid email format",
			respCode:  http.StatusBadRequest,
			email:     "xxx.test",
			expectErr: nil,
		},
		{
			name:      "email already exist",
			respCode:  http.StatusConflict,
			email:     email,
			expectErr: nil,
		},
	}
	e := echo.New()
	h := &exchangeHandler{
		services: &pkg.Services{
			EmailUserService: services.NewEmailUserService(
				context.Background(),
				mem.NewMemoryRepository(),
			),
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/subscribe", nil)
			form, _ := url.ParseQuery(req.URL.RawQuery)
			form.Add("email", test.email)
			req.URL.RawQuery = form.Encode()

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			assert.NoError(t, h.CreateMailSubscriber(c))
			assert.Equal(t, test.respCode, rec.Code)
		})
	}
}

func TestSendEmails(t *testing.T) {
	e := echo.New()
	h := &exchangeHandler{
		services: &pkg.Services{
			NotificationService: &mock.NotificationService{},
		},
	}

	req := httptest.NewRequest(http.MethodPost, "/sendEmails", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, h.SendEmails(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

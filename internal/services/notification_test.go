package services_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain/event"
	"exchange/internal/domain/rate"
	mock_rate "exchange/internal/domain/rate/mocks"
	mock_user "exchange/internal/domain/user/mocks"
	"exchange/internal/services"
	mock_services "exchange/internal/services/mocks"
)

func TestNotify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_user.NewMockUserRepository(ctrl)
	currencyServiceMock := mock_rate.NewMockICurrencyService(ctrl)
	mailSender := mock_services.NewMockIMailService(ctrl)

	const btcUahRate = 1_000_000.0

	emails := []string{"1@email", "2@email"}

	currencyServiceMock.EXPECT().GetCurrency(context.Background(), &rate.Rate{
		BaseCurrency:  rate.BTC,
		QuoteCurrency: rate.UAH,
	}).Return(btcUahRate, nil)
	userRepoMock.EXPECT().
		GetAllEmails(context.Background()).
		Return(emails, nil)
	mailSender.EXPECT().SendEmail(context.Background(), btcUahRate, emails).Return(nil)

	notifierService := services.NewNotificationService(
		userRepoMock,
		currencyServiceMock,
		mailSender,
	)

	err := notifierService.Notify(context.Background(), event.DefaultNotification())
	require.NoError(t, err)
}

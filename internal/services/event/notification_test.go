package event_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain/notification"
	"exchange/internal/domain/rate"
	"exchange/internal/services/event"
	mock_event "exchange/internal/services/event/mocks"
)

func TestNotify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_event.NewMockUserRepository(ctrl)
	currencyServiceMock := mock_event.NewMockICurrencyService(ctrl)
	mailSender := mock_event.NewMockIMailService(ctrl)

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

	notifierService := event.NewNotificationService(
		userRepoMock,
		currencyServiceMock,
		mailSender,
	)

	err := notifierService.Notify(context.Background(), notification.DefaultNotification())
	require.NoError(t, err)
}

package event_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"notifier/internal/domain/notification"
	"notifier/internal/domain/rate"
	"notifier/internal/services/event"
	mock_event "notifier/internal/services/event/mocks"
)

func TestNotify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_event.NewMockUserRepository(ctrl)
	currencyServiceMock := mock_event.NewMockICurrencyService(ctrl)
	mailSender := mock_event.NewMockIMailService(ctrl)

	currency := rate.NewCurrency(1_000_000.0)

	emails := []string{"1@email", "2@email"}

	currencyServiceMock.EXPECT().GetCurrency(context.Background(), &rate.Rate{
		BaseCurrency:  rate.BTC,
		QuoteCurrency: rate.UAH,
	}).Return(currency, nil)
	userRepoMock.EXPECT().
		GetAllEmails(context.Background()).
		Return(emails, nil)
	mailSender.EXPECT().SendEmail(context.Background(), currency.Value, emails).Return(nil)

	notifierService := event.NewNotificationService(
		userRepoMock,
		currencyServiceMock,
		mailSender,
	)

	err := notifierService.Notify(context.Background(), notification.DefaultNotification())
	require.NoError(t, err)
}

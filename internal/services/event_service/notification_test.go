package event_service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain/event_domain"
	"exchange/internal/domain/rate_domain"
	"exchange/internal/services/event_service"
	mock_event_service "exchange/internal/services/event_service/mocks"
)

func TestNotify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_event_service.NewMockUserRepository(ctrl)
	currencyServiceMock := mock_event_service.NewMockICurrencyService(ctrl)
	mailSender := mock_event_service.NewMockIMailService(ctrl)

	const btcUahRate = 1_000_000.0

	emails := []string{"1@email", "2@email"}

	currencyServiceMock.EXPECT().GetCurrency(context.Background(), &rate_domain.Rate{
		BaseCurrency:  rate_domain.BTC,
		QuoteCurrency: rate_domain.UAH,
	}).Return(btcUahRate, nil)
	userRepoMock.EXPECT().
		GetAllEmails(context.Background()).
		Return(emails, nil)
	mailSender.EXPECT().SendEmail(context.Background(), btcUahRate, emails).Return(nil)

	notifierService := event_service.NewNotificationService(
		userRepoMock,
		currencyServiceMock,
		mailSender,
	)

	err := notifierService.Notify(context.Background(), event_domain.DefaultNotification())
	require.NoError(t, err)
}

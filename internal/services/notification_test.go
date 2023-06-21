package services_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"exchange/internal/domain"
	mock_domain "exchange/internal/domain/mocks"
	"exchange/internal/services"
	mock_services "exchange/internal/services/mocks"
)

func TestNotify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepoMock := mock_domain.NewMockUserRepository(ctrl)
	currencyServiceMock := mock_domain.NewMockICurrencyService(ctrl)
	mailSender := mock_services.NewMockIMailService(ctrl)

	const rate = 1_000_000.0
	emails := []string{"1@email", "2@email"}

	currencyServiceMock.EXPECT().GetCurrency(context.Background(), &domain.Currency{
		BaseCurrency:  domain.BTC,
		QuoteCurrency: domain.UAH,
	}).Return(rate, nil)
	userRepoMock.EXPECT().
		GetAllEmails(context.Background()).
		Return(emails, nil)
	mailSender.EXPECT().SendEmail(context.Background(), rate, emails).Return(nil)

	notifierService := services.NewNotificationService(
		userRepoMock,
		currencyServiceMock,
		mailSender,
	)

	err := notifierService.Notify(context.Background(), domain.DefaultNotification())
	require.NoError(t, err)
}

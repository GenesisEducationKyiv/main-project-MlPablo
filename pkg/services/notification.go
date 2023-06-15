package services

import (
	"context"

	"exchange/pkg/domain"
)

// Here we define mail service interface that we need for sending emails.
type IMailService interface {
	SendEmail(ctx context.Context, data any, receivers ...string) error
}

type notificationService struct {
	ctx             context.Context
	emailUserRepo   domain.EmailRepository
	currencyService domain.ICurrencyService
	mailService     IMailService
}

func NewNotificationService(
	ctx context.Context,
	emailRepo domain.EmailRepository,
	currencyService domain.ICurrencyService,
	mailService IMailService,
) domain.INotificationService {
	return &notificationService{
		ctx:             ctx,
		emailUserRepo:   emailRepo,
		mailService:     mailService,
		currencyService: currencyService,
	}
}

// Notify users via email due to our business logic.
func (n *notificationService) Notify(ctx context.Context, _ *domain.Notification) error {
	btcUah := domain.GetBitcoinToUAH()

	currency, err := n.currencyService.GetCurrency(ctx, btcUah)
	if err != nil {
		return err
	}

	emails, err := n.emailUserRepo.GetAllEmails(ctx)
	if err != nil {
		return err
	}

	return n.mailService.SendEmail(ctx, currency, emails...)
}

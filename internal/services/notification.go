package services

import (
	"context"

	"exchange/internal/domain/event"
	"exchange/internal/domain/rate"
	"exchange/internal/domain/user"
)

//go:generate mockgen -source=notification.go -destination=mocks/notification.go

// Here we define mail service interface that we need for sending emails.
type IMailService interface {
	SendEmail(ctx context.Context, data any, receivers ...string) error
}

type notificationService struct {
	userRepo        user.UserRepository
	currencyService rate.ICurrencyService
	mailService     IMailService
}

func NewNotificationService(
	userRepo user.UserRepository,
	currencyService rate.ICurrencyService,
	mailService IMailService,
) event.INotificationService {
	return &notificationService{
		userRepo:        userRepo,
		mailService:     mailService,
		currencyService: currencyService,
	}
}

// Notify users via email due to our business logic.
func (n *notificationService) Notify(ctx context.Context, _ *event.Notification) error {
	btcUah := rate.GetBitcoinToUAH()

	currency, err := n.currencyService.GetCurrency(ctx, btcUah)
	if err != nil {
		return err
	}

	emails, err := n.userRepo.GetAllEmails(ctx)
	if err != nil {
		return err
	}

	return n.mailService.SendEmail(ctx, currency, emails...)
}

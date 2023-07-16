package event

import (
	"context"

	"notifier/internal/domain/rate"
)

//go:generate mockgen -source=service.go -destination=mocks/notification.go

// Here we define mail service interface that we need for sending emails.
type IMailService interface {
	SendEmail(ctx context.Context, data any, receivers ...string) error
}

type UserRepository interface {
	GetAllEmails(ctx context.Context) ([]string, error)
}

type ICurrencyService interface {
	GetCurrency(ctx context.Context, c *rate.Rate) (*rate.Currency, error)
}

type Service struct {
	userRepo        UserRepository
	currencyService ICurrencyService
	mailService     IMailService
}

func NewNotificationService(
	userRepo UserRepository,
	currencyService ICurrencyService,
	mailService IMailService,
) *Service {
	return &Service{
		userRepo:        userRepo,
		mailService:     mailService,
		currencyService: currencyService,
	}
}

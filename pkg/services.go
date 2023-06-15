package pkg

import "exchange/pkg/domain"

// Aggregate of all services, so we can put this into the controllers layer,
// without the need to put it one by one.
type Services struct {
	CurrencyService     domain.ICurrencyService
	EmailUserService    domain.IEmailService
	NotificationService domain.INotificationService
}

func NewServices(
	currency domain.ICurrencyService,
	email domain.IEmailService,
	notification domain.INotificationService,
) *Services {
	return &Services{
		CurrencyService:     currency,
		EmailUserService:    email,
		NotificationService: notification,
	}
}

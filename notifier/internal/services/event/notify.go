package event

import (
	"context"

	"notifier/internal/domain/notification"
	"notifier/internal/domain/rate"
)

// Notify users via email due to our business logic.
func (n *Service) Notify(ctx context.Context, _ *notification.Notification) error {
	btcUah := rate.GetBitcoinToUAH()

	currency, err := n.currencyService.GetCurrency(ctx, btcUah)
	if err != nil {
		return err
	}

	emails, err := n.userRepo.GetAllEmails(ctx)
	if err != nil {
		return err
	}

	return n.mailService.SendEmail(ctx, currency.Value, emails...)
}

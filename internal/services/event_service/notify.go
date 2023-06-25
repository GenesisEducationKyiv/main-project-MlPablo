package event_service

import (
	"context"

	"exchange/internal/domain/event_domain"
	"exchange/internal/domain/rate_domain"
)

// Notify users via email due to our business logic.
func (n *Service) Notify(ctx context.Context, _ *event_domain.Notification) error {
	btcUah := rate_domain.GetBitcoinToUAH()

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

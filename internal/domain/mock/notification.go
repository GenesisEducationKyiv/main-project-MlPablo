package mock

import (
	"context"

	"exchange/internal/domain"
)

type NotificationService struct{}

func (c *NotificationService) Notify(_ context.Context, _ *domain.Notification) error {
	return nil
}

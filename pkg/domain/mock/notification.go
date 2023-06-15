package mock

import (
	"context"

	"exchange/pkg/domain"
)

type NotificationService struct{}

func (c *NotificationService) Notify(_ context.Context, _ *domain.Notification) error {
	return nil
}

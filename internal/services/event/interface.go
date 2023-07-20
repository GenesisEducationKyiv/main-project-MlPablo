package event

import (
	"context"

	"exchange/internal/domain/notification"
)

type INotificationService interface {
	Notify(ctx context.Context, n *notification.Notification) error
}

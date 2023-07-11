package event

import (
	"context"

	"notifier/internal/domain/notification"
)

//go:generate mockgen -source=interface.go -destination=mocks/interface.go

type INotificationService interface {
	Notify(ctx context.Context, n *notification.Notification) error
}

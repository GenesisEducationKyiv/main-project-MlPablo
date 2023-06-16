package domain

import "context"

// Notification domain that responses for notifying users.
// It can be easily expanded for notify specific user group.
// Here UserGroup is useless. But it was made for future purposes.
type Notification struct {
	UserGroup string `json:"user_group"`
}

func DefaultNotification() *Notification {
	return &Notification{}
}

type INotificationService interface {
	Notify(ctx context.Context, n *Notification) error
}

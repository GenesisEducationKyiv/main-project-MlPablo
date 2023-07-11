package user

import (
	"context"

	"notifier/internal/domain/user"
)

//go:generate mockgen -source=interface.go -destination=mocks/interface.go

type IUserService interface {
	NewUser(ctx context.Context, eu *user.User) error
}

package user

import (
	"context"

	"notifier/internal/domain/user"
)

//go:generate mockgen -source=interface.go -destination=mocks/interface.go

type IUserService interface {
	NewUser(ctx context.Context, eu *user.User) error
	DeleteUser(ctx context.Context, eu *user.User) error
	UserExist(ctx context.Context, eu *user.User) (bool, error)
}

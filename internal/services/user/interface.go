package user

import (
	"context"

	"exchange/internal/domain/user"
)

type IUserService interface {
	NewUser(ctx context.Context, eu *user.User) error
}

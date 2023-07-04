package user

import (
	"context"

	"exchange/internal/domain/user"
)

// Check if mail exist and than create new.
func (e *Service) NewUser(ctx context.Context, u *user.User) error {
	exist, err := e.userRepo.EmailExist(ctx, u.Email)
	if err != nil {
		return err
	}

	if exist {
		return user.ErrAlreadyExist
	}

	return e.userRepo.SaveUser(ctx, u)
}

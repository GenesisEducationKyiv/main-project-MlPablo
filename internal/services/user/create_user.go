package user

import (
	"context"

	user_domain "exchange/internal/domain/user"
)

// Check if mail exist and than create new.
func (e *Service) NewUser(ctx context.Context, u *user_domain.User) error {
	exist, err := e.userRepo.EmailExist(ctx, u.Email)
	if err != nil {
		return err
	}

	if exist {
		return user_domain.ErrAlreadyExist
	}

	return e.userRepo.SaveUser(ctx, u)
}

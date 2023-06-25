package user_service

import (
	"context"

	"exchange/internal/domain/user_domain"
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

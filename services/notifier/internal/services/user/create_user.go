package user

import (
	"context"

	"notifier/internal/domain/user"
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

	return e.userRepo.Save(ctx, u)
}

func (e *Service) DeleteUser(ctx context.Context, u *user.User) error {
	return e.userRepo.Delete(ctx, u.Email)
}

func (e *Service) UserExist(ctx context.Context, u *user.User) (bool, error) {
	return e.userRepo.EmailExist(ctx, u.Email)
}

package services

import (
	"context"

	"exchange/internal/domain/user"
)

type userService struct {
	emailUserRepo user.UserRepository
}

func NewUserService(
	emailRepo user.UserRepository,
) user.IUserService {
	return &userService{
		emailUserRepo: emailRepo,
	}
}

// Check if mail exist and than create new.
func (e *userService) NewUser(ctx context.Context, u *user.User) error {
	exist, err := e.emailUserRepo.EmailExist(ctx, u.Email)
	if err != nil {
		return err
	}

	if exist {
		return user.ErrAlreadyExist
	}

	return e.emailUserRepo.SaveUser(ctx, u)
}

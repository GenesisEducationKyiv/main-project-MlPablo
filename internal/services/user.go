package services

import (
	"context"

	"exchange/internal/domain"
)

type userService struct {
	emailUserRepo domain.UserRepository
}

func NewUserService(
	emailRepo domain.UserRepository,
) domain.IUserService {
	return &userService{
		emailUserRepo: emailRepo,
	}
}

// Check if mail exist and than create new.
func (e *userService) NewUser(ctx context.Context, user *domain.User) error {
	exist, err := e.emailUserRepo.EmailExist(ctx, user.Email)
	if err != nil {
		return err
	}

	if exist {
		return domain.ErrAlreadyExist
	}

	return e.emailUserRepo.SaveUser(ctx, user)
}

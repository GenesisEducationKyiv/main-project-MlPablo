package user

import (
	"context"

	"exchange/internal/domain/user"
)

//go:generate mockgen -source=service.go -destination=mocks/user.go

type UserRepository interface {
	Save(ctx context.Context, user *user.User) error
	EmailExist(ctx context.Context, email string) (bool, error)
}

type Service struct {
	userRepo UserRepository
}

func NewUserService(
	userRepo UserRepository,
) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

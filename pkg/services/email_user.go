package services

import (
	"context"

	"exchange/pkg/domain"
)

type emailUserService struct {
	ctx           context.Context
	emailUserRepo domain.EmailRepository
}

func NewEmailUserService(
	ctx context.Context,
	emailRepo domain.EmailRepository,
) domain.IEmailService {
	return &emailUserService{
		ctx:           ctx,
		emailUserRepo: emailRepo,
	}
}

// Check if mail exist and than create new.
func (e *emailUserService) NewEmailUser(ctx context.Context, emailUser *domain.EmailUser) error {
	exist, err := e.emailUserRepo.EmailExist(ctx, emailUser.Email)
	if err != nil {
		return err
	}

	if exist {
		return domain.ErrAlreadyExist
	}

	return e.emailUserRepo.SaveEmail(ctx, emailUser)
}

package domain

import "context"

// EmailUser domain that responds for saving info about emails.
// We can easily expand this logic for more complex users. For example user name, etc...
type EmailUser struct {
	Email string `json:"email"`
}

func NewEmailUser(email string) *EmailUser {
	return &EmailUser{
		Email: email,
	}
}

func (e *EmailUser) Validate() error {
	if !isEmailValid(e.Email) {
		return ErrBadRequest
	}

	return nil
}

type IEmailService interface {
	NewEmailUser(ctx context.Context, eu *EmailUser) error
}

type EmailRepository interface {
	SaveEmail(ctx context.Context, eu *EmailUser) error
	EmailExist(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*EmailUser, error)
	GetAllEmails(ctx context.Context) ([]string, error)
}

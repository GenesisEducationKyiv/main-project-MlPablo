package domain

import "context"

//go:generate mockgen -source=user.go -destination=mocks/user.go

// User domain that responds for saving info about emails.
// We can easily expand this logic for more complex users. For example user name, etc...
type User struct {
	Email string `json:"email"`
}

func NewUser(email string) *User {
	return &User{
		Email: email,
	}
}

func (e *User) Validate() error {
	if !isEmailValid(e.Email) {
		return ErrBadRequest
	}

	return nil
}

type IUserService interface {
	NewUser(ctx context.Context, eu *User) error
}

type UserRepository interface {
	SaveUser(ctx context.Context, user *User) error
	EmailExist(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetAllEmails(ctx context.Context) ([]string, error)
}

package user

import "net/mail"

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

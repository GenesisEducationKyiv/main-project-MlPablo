package user_domain

import "errors"

var (
	ErrNotFound     = errors.New("user: not found")
	ErrAlreadyExist = errors.New("user: already exist")
	ErrBadRequest   = errors.New("user: invalid request")
)

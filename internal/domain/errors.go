package domain

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("item not found")
	ErrAlreadyExist   = errors.New("item already exist")
	ErrBadRequest     = errors.New("invalid request")
	ErrInvalidStatus  = errors.New("invalid status")
)

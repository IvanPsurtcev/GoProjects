package auth

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exist")
	ErrUserNotFound      = errors.New("user does not exist")
)

package request

import "errors"

var (
	ErrNotFound         = errors.New("request not found")
	ErrReqAlreadyExists = errors.New("request already exists")
	ErrUserNotFound     = errors.New("user not found")
)

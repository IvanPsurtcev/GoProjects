package auth

import "context"

type Repository interface {
	CreateUser(ctx context.Context, address string) (int64, error)
	GetUsers(ctx context.Context, offset, limit int64) ([]User, error)
	GetUser(ctx context.Context, address string) (*User, error)
	GetRandUser(ctx context.Context, exceptAddress string) (*User, error)
}

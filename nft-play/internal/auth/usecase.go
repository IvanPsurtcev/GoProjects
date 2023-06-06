package auth

import "context"

type UseCase interface {
	Authorize(ctx context.Context, params *AuthorizeRequest) (*AuthorizeResponse, error)
	GetUsers(ctx context.Context, params *GetUsersRequest) (*GetUsersResponse, error)
	GetUser(ctx context.Context, params *GetUserRequest) (*User, error)
}

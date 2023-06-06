package usecase

import (
	"context"
	"go.uber.org/zap"
	"metamask-auth/internal/auth"
	"metamask-auth/pkg/jwt"
)

type authUseCase struct {
	logger    *zap.Logger
	jwtClient jwt.HmacProvider
	authRepo  auth.Repository
}

func NewUseCase(logger *zap.Logger, jwtClient jwt.HmacProvider, authRepo auth.Repository) auth.UseCase {
	return &authUseCase{
		logger:    logger,
		jwtClient: jwtClient,
		authRepo:  authRepo,
	}
}

func (u *authUseCase) Authorize(ctx context.Context, params *auth.AuthorizeRequest) (*auth.AuthorizeResponse, error) {
	_, err := u.authRepo.CreateUser(ctx, params.Address)
	// TODO: временный вариант, пока нет полноценной регистрации и входа
	if err != nil && err != auth.ErrUserAlreadyExists {
		return nil, err
	}
	token, err := u.jwtClient.CreateStandard(params.Address)
	if err != nil {
		u.logger.Error("Authorize err:", zap.Error(err))
		return nil, err
	}
	return &auth.AuthorizeResponse{
		Address: params.Address,
		Jwt:     token,
	}, nil
}

func (u *authUseCase) GetUser(ctx context.Context, params *auth.GetUserRequest) (*auth.User, error) {
	usr, err := u.authRepo.GetUser(ctx, params.Address)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *authUseCase) GetUsers(ctx context.Context, params *auth.GetUsersRequest) (*auth.GetUsersResponse, error) {
	res, err := u.authRepo.GetUsers(ctx, params.Offset, params.Limit)
	if err != nil {
		return nil, err
	}
	return &auth.GetUsersResponse{
		Users: res,
	}, nil
}

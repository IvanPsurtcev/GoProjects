package usecase

import (
	"context"
	"metamask-auth/internal/auth"
	"metamask-auth/internal/request"

	"go.uber.org/zap"
)

type requestUseCase struct {
	logger      *zap.Logger
	requestRepo request.Repository
	authRepo    auth.Repository
}

func NewUseCase(logger *zap.Logger, requestRepo request.Repository, authRepo auth.Repository) request.UseCase {
	return &requestUseCase{
		logger:      logger,
		requestRepo: requestRepo,
		authRepo:    authRepo,
	}
}

func (u *requestUseCase) CreateRequest(ctx context.Context, params *request.CreateRequest) (*request.Request, error) {
	res, err := u.requestRepo.Create(ctx, params.ApplicantAddr, params.ReceiverAddr)
	return res, err
}

func (u *requestUseCase) DeclineRequest(ctx context.Context, params *request.DeclineRequest) (*request.Request, error) {
	res, err := u.requestRepo.SetArg(ctx, ColStatus, StatusDeclined, params.ID)
	return res, err
}

func (u *requestUseCase) AcceptRequest(ctx context.Context, params *request.AcceptRequest) (*request.Request, error) {
	res, err := u.requestRepo.SetArg(ctx, ColStatus, StatusAccepted, params.ID)
	return res, err
}

func (u *requestUseCase) GetUserActiveInvites(ctx context.Context, params *request.GetInvitesRequest) (*request.GetInvitesResponse, error) {
	res, err := u.requestRepo.GetUserActiveInvites(ctx, params.Address)
	return &request.GetInvitesResponse{
		Requests: res,
	}, err
}

func (u *requestUseCase) CreateRandRequest(ctx context.Context, params *request.CreateRequest) (*request.Request, error) {
	randUser, err := u.authRepo.GetRandUser(ctx, params.ApplicantAddr)
	if err != nil {
		return nil, err
	}
	params.ReceiverAddr = randUser.Address
	return u.CreateRequest(ctx, params)
}

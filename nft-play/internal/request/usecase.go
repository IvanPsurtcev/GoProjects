package request

import "context"

type UseCase interface {
	CreateRequest(ctx context.Context, createRequest *CreateRequest) (*Request, error)
	CreateRandRequest(ctx context.Context, createRequest *CreateRequest) (*Request, error)
	GetUserActiveInvites(ctx context.Context, params *GetInvitesRequest) (*GetInvitesResponse, error)
	DeclineRequest(ctx context.Context, declineRequest *DeclineRequest) (*Request, error)
	AcceptRequest(ctx context.Context, acceptRequest *AcceptRequest) (*Request, error)
}

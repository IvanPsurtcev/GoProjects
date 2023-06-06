package game

import "context"

type UseCase interface {
	CreateGame(ctx context.Context, params *CreateGameRequest) (*CreateGameResponse, error)
	CreateRandGame(ctx context.Context, params *CreateGameRequest) (*CreateGameResponse, error)
	SetApplicantBet(ctx context.Context, params *SetBetRequest) error
	SetReceiverBet(ctx context.Context, params *SetBetRequest) error
	SetApplicantStatus(ctx context.Context, params *SetPlayerStatusRequest) error
	SetReceiverStatus(ctx context.Context, params *SetPlayerStatusRequest) error
	CancelGame(ctx context.Context, params *SetGameStatusRequest) error
	GameDone(ctx context.Context, params *SetGameStatusRequest) error
	GetGames(ctx context.Context, params *GetGamesRequest) (*GetGamesResponse, error)
	GetGame(ctx context.Context, params *BaseGameRequest) (*GetGameResponse, error)
	GetUserNfts(ctx context.Context, params *GetUserNftsRequest) (*GetUserNftsResponse, error)
	RunGame(ctx context.Context, params *BaseGameRequest) (*GetGameResponse, error)
}

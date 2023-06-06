package game

import "context"

type Repository interface {
	CreateGame(ctx context.Context, id int64, applicantAddr, receiverAddr string) error
	SetArg(ctx context.Context, id int64, colName string, colArg interface{}) error
	SetWinner(ctx context.Context, id int64, winner string) error
	GetGame(ctx context.Context, id int64) (*Game, error)
	GetGames(ctx context.Context, address string, status int64) ([]Game, error)
	GetGameNft(ctx context.Context, id int64) ([]string, []string, error)
}

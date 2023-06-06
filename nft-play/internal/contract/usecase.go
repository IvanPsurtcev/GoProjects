package contract

import "context"

type UseCase interface {
	CreateGame(ctx context.Context) (int64, error)
	SetWinner(ctx context.Context, winnerGame *WinnerGame) error
}

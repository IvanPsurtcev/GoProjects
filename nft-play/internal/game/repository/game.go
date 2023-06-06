package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"metamask-auth/internal/game"
	"strings"
)

type gameRepository struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewRepository(logger *zap.Logger, db *sqlx.DB) game.Repository {
	return &gameRepository{
		logger: logger,
		db:     db,
	}
}

func (r *gameRepository) CreateGame(ctx context.Context, id int64, applicantAddr, receiverAddr string) error {
	query := `INSERT INTO game (id, applicant_addr, receiver_addr) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING RETURNING id`
	var gameId int64
	if err := r.db.QueryRowContext(ctx, query, id, applicantAddr, receiverAddr).Scan(&gameId); err != nil {
		if err == sql.ErrNoRows {
			return game.ErrGameAlreadyExists
		}
		r.logger.Error("CreateGame err:", zap.Error(err))
		return err
	}
	return nil
}

func (r *gameRepository) SetArg(ctx context.Context, id int64, colName string, colArg interface{}) error {
	query := fmt.Sprintf(`UPDATE game SET %s = $1 WHERE id = $2 RETURNING id`, colName)
	var gameId int64
	err := r.db.QueryRowContext(ctx, query, colArg, id).Scan(&gameId)
	if err != nil {
		if err == sql.ErrNoRows {
			return game.ErrGameNotFound
		}
		r.logger.Error("SetArg err:", zap.Error(err))
		return err
	}
	return nil
}

func (r *gameRepository) GetGameNft(ctx context.Context, id int64) ([]string, []string, error) {
	query := `SELECT applicant_nft, receiver_nft FROM game WHERE id = $1`
	var applicantNft, receiverNft string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&applicantNft, &receiverNft)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, game.ErrGameNotFound
		}
		r.logger.Error("GetNft err:", zap.Error(err))
		return nil, nil, err
	}
	return strings.Split(applicantNft, " "), strings.Split(receiverNft, " "), nil
}

func (r *gameRepository) GetGame(ctx context.Context, id int64) (*game.Game, error) {
	query := `SELECT * FROM game WHERE id = $1`
	var data game.Game
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&data.Id,
		&data.ApplicantAddr,
		&data.ReceiverAddr,
		&data.WinnerAddr,
		&data.Status,
		&data.ApplicantReady,
		&data.ReceiverReady,
		&data.CreateDate,
		&data.ApplicantNft,
		&data.ReceiverNft)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, game.ErrGameNotFound
		}
		r.logger.Error("GetGame err", zap.Error(err))
		return nil, err
	}
	return &data, nil
}

func (r *gameRepository) SetWinner(ctx context.Context, id int64, winner string) error {
	query := `UPDATE game SET winner_addr = $1, status = 2 WHERE id = $2 RETURNING id`
	var gameId int64
	if err := r.db.QueryRowContext(ctx, query, winner, id).Scan(&gameId); err != nil {
		if err == sql.ErrNoRows {
			return game.ErrGameNotFound
		}
		r.logger.Error("SetWinner err", zap.Error(err))
		return err
	}
	return nil
}

func (r *gameRepository) GetGames(ctx context.Context, address string, status int64) ([]game.Game, error) {
	query := `SELECT * FROM game WHERE (receiver_addr = $1 OR applicant_addr = $1) AND status = $2`
	var data []game.Game
	if err := r.db.SelectContext(ctx, &data, query, address, status); err != nil {
		if err == sql.ErrNoRows {
			return nil, game.ErrGameNotFound
		}
		return nil, err
	}
	return data, nil
}

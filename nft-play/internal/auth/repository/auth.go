package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"metamask-auth/internal/auth"
)

type authRepository struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewRepository(logger *zap.Logger, db *sqlx.DB) auth.Repository {
	return &authRepository{
		logger: logger,
		db:     db,
	}
}

func (r *authRepository) CreateUser(ctx context.Context, address string) (int64, error) {
	query := `INSERT INTO "user" (address) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`
	var dbUserID int64
	err := r.db.QueryRowxContext(ctx, query, address).Scan(&dbUserID)
	switch err {
	case sql.ErrNoRows:
		return dbUserID, auth.ErrUserAlreadyExists
	case nil:
		return dbUserID, nil
	default:
		r.logger.Error("CreateUser err:", zap.Error(err))
		return 0, err
	}
}

func (r *authRepository) GetUsers(ctx context.Context, offset, limit int64) ([]auth.User, error) {
	var data []auth.User
	var err error
	if limit == -1 {
		query := `SELECT * from "user" OFFSET $1`
		err = r.db.SelectContext(ctx, &data, query, offset)
	} else {
		query := `SELECT * from "user" OFFSET $1 LIMIT $2`
		err = r.db.SelectContext(ctx, &data, query, offset, limit)
	}

	switch err {
	case sql.ErrNoRows:
		return []auth.User{}, auth.ErrUserAlreadyExists
	case nil:
		return data, nil
	default:
		r.logger.Error("CreateUser err:", zap.Error(err))
		return nil, err
	}
}

func (r *authRepository) GetUser(ctx context.Context, address string) (*auth.User, error) {
	query := `SELECT * from "user" WHERE address = $1`
	var data auth.User
	err := r.db.GetContext(ctx, &data, query, address)
	switch err {
	case sql.ErrNoRows:
		return nil, auth.ErrUserNotFound
	case nil:
		return &data, nil
	default:
		r.logger.Error("GetUser err:", zap.Error(err))
		return nil, err
	}
}

func (r *authRepository) GetRandUser(ctx context.Context, exceptAddress string) (*auth.User, error) {
	// TODO: сделать более оптимально. Сейчас такой запрос работает очень медленно
	query := `
		SELECT 
		    *
		FROM 
		    "user"
		WHERE 
		    address != $1
		ORDER BY random()
		LIMIT 1;`
	var data auth.User
	err := r.db.GetContext(ctx, &data, query, exceptAddress)
	switch err {
	case sql.ErrNoRows:
		return nil, auth.ErrUserNotFound
	case nil:
		return &data, nil
	default:
		r.logger.Error("GetUser err:", zap.Error(err))
		return nil, err
	}
}

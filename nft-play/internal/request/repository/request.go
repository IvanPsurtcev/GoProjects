package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"metamask-auth/internal/request"
)

type requestRepository struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewRepository(logger *zap.Logger, db *sqlx.DB) request.Repository {
	return &requestRepository{
		logger: logger,
		db:     db,
	}
}

func (r *requestRepository) Create(ctx context.Context, applicantAddr, receiverAddr string) (*request.Request, error) {
	query := "INSERT INTO request (applicant_addr, receiver_addr) VALUES ($1, $2) ON CONFLICT DO NOTHING RETURNING *"
	var data request.Request
	err := r.db.QueryRowContext(ctx, query, applicantAddr, receiverAddr).Scan(
		&data.ID,
		&data.ApplicantAddr,
		&data.ReceiverAddr,
		&data.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, request.ErrReqAlreadyExists
		}
		r.logger.Error("error creating request", zap.Error(err))
		return nil, err
	}
	return &data, nil
}

func (r *requestRepository) GetByID(ctx context.Context, id int) (*request.Request, error) {
	query := "SELECT id, applicant_addr, receiver_addr, status FROM request WHERE id = $1"
	req := &request.Request{}
	if err := r.db.GetContext(ctx, req, query, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, request.ErrNotFound
		}
		r.logger.Error("error getting request by ID", zap.Error(err))
		return nil, fmt.Errorf("error getting request by ID: %w", err)
	}
	return req, nil
}

func (r *requestRepository) SetArg(ctx context.Context, colName string, colArg interface{}, id int64) (*request.Request, error) {
	query := fmt.Sprintf("UPDATE request SET %s = $1 WHERE id = $2 RETURNING *", colName)
	var data request.Request
	err := r.db.QueryRowContext(ctx, query, colArg, id).Scan(
		&data.ID,
		&data.ApplicantAddr,
		&data.ReceiverAddr,
		&data.Status)
	switch err {
	case sql.ErrNoRows:
		return nil, request.ErrNotFound
	case nil:
		return &data, nil
	default:
		r.logger.Error("err during updating request:", zap.Error(err))
		return nil, err
	}
}

func (r *requestRepository) Delete(ctx context.Context, req *request.Request) error {
	query := "DELETE FROM request WHERE id = $1"
	if _, err := r.db.ExecContext(ctx, query, req.ID); err != nil {
		r.logger.Error("error deleting request", zap.Error(err))
		return fmt.Errorf("error deleting request: %w", err)
	}
	return nil
}

func (r *requestRepository) GetUserActiveInvites(ctx context.Context, address string) ([]request.Request, error) {
	query := `SELECT * FROM request WHERE receiver_addr = $1 AND status = 1`
	var data []request.Request
	if err := r.db.SelectContext(ctx, &data, query, address); err != nil {
		if err == sql.ErrNoRows {
			return nil, request.ErrUserNotFound
		}
		return nil, err
	}
	return data, nil
}

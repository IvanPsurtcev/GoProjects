package request

import "context"

type Repository interface {
	Create(ctx context.Context, applicantAddr, receiverAddr string) (*Request, error)
	GetUserActiveInvites(ctx context.Context, address string) ([]Request, error)
	GetByID(ctx context.Context, id int) (*Request, error)
	SetArg(ctx context.Context, colName string, colArg interface{}, id int64) (*Request, error)
	Delete(ctx context.Context, req *Request) error
}

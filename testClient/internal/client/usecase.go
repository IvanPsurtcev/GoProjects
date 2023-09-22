package client

import (
	"context"
)

type UseCase interface {
	ProcessItems(ctx context.Context, batch *BatchRequest) error
	GetLimits(ctx context.Context) (Limits, error)
}

package client

import (
	"context"
	"time"
)

// Service defines external service that can process batches of items.
type Service interface {
	GetLimits() (n uint64, p time.Duration)
	Process(ctx context.Context, batch Batch) error
}

type BatchRequest struct {
	Batch `json:"batch"`
}

type Limits struct {
	N uint64
	P time.Duration
}

// Batch is a batch of items.
type Batch []Item

// Item is some abstract item.
type Item struct{}

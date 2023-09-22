package usecase

import (
	"context"
	"sync"
	"testClient/internal/client"
	"testClient/pkg/slice"
	"time"
)

type Client struct {
	service   client.Service
	n         uint64
	p         time.Duration
	safeSlice *slice.SafeSlice
	shutdown  chan bool
	mu        sync.Mutex
}

func NewClient(service client.Service, safeSlice *slice.SafeSlice) *Client {
	n, p := service.GetLimits()
	client := &Client{
		service:   service,
		n:         n,
		p:         p,
		safeSlice: safeSlice,
		shutdown:  make(chan bool),
	}

	go client.worker()

	return client
}

func (c *Client) ProcessItems(ctx context.Context, batch *client.BatchRequest) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.safeSlice.Append(batch.Batch)
	return nil
}

func (c *Client) GetLimits(ctx context.Context) (client.Limits, error) {
	return client.Limits{
		N: c.n,
		P: c.p,
	}, nil
}

func (c *Client) worker() {
	ticker := time.NewTicker(c.p)
	defer ticker.Stop()

	for {
		select {
		case <-c.shutdown:
			c.Shutdown()
		case <-ticker.C:
			batch := c.collectBatch()
			if len(batch) > 0 {
				_ = c.service.Process(context.Background(), batch)
			}
		}
	}
}

func (c *Client) collectBatch() client.Batch {
	batchSize := int(c.n)
	deletedItems := c.safeSlice.DeleteAndReturn(batchSize)

	return deletedItems
}

// Shutdown gracefully stops the client worker.
func (c *Client) Shutdown() {
	close(c.shutdown)
}

package service

import (
	"context"
	"testClient/internal/client"
	"time"
)

// MockService is an example implementation of the Service interface for testing.
type MockService struct{}

func (m *MockService) GetLimits() (uint64, time.Duration) {
	return 10, 3 * time.Second
}

func (m *MockService) Process(ctx context.Context, batch client.Batch) error {
	if len(batch) > 10 {
		time.Sleep(100 * time.Millisecond)
		return client.ErrBlocked
	}
	return nil
}

package slice

import (
	"sync"
	"testClient/internal/client"
)

type SafeSlice struct {
	mu    sync.Mutex
	items client.Batch
}

func NewSafeSlice() *SafeSlice {
	return &SafeSlice{
		mu:    sync.Mutex{},
		items: make(client.Batch, 0, 500),
	}
}

func (s *SafeSlice) Append(batch client.Batch) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, batch...)
}

func (s *SafeSlice) DeleteAndReturn(n int) client.Batch {
	s.mu.Lock()
	defer s.mu.Unlock()

	if n > len(s.items) {
		n = len(s.items)
	}

	deletedItems := s.items[:n]
	s.items = s.items[n:]

	return deletedItems
}

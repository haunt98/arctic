package arctic

import (
	"sync"
)

type Store struct {
	mu    sync.RWMutex
	pairs map[string][]byte
}

func NewStore() *Store {
	return &Store{
		pairs: make(map[string][]byte),
	}
}

func (s *Store) get(key string) []byte {
	s.mu.RLock()
	defer s.mu.RLock()

	value, ok := s.pairs[key]
	if !ok {
		return nil
	}

	return value
}

func (s *Store) put(key string, value []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pairs[key] = value
}

func (s *Store) clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pairs = make(map[string][]byte)
}

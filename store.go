package arctic

import (
	"log"
	"sync"
)

// Store has key value pairs
type Store struct {
	sync.Mutex

	pairs map[string][]byte
}

// NewStore return Store
func NewStore() *Store {
	return &Store{
		pairs: make(map[string][]byte),
	}
}

func (s *Store) get(key string) []byte {
	s.Lock()
	defer s.Unlock()

	value, ok := s.pairs[key]
	if !ok {
		log.Printf("key %s not found\n", key)
		return nil
	}

	return value
}

func (s *Store) put(key string, value []byte) {
	s.Lock()
	defer s.Unlock()

	_, ok := s.pairs[key]
	if ok {
		log.Printf("key %s exist\n", key)
	}

	s.pairs[key] = value
}

func (s *Store) clear() {
	s.Lock()
	defer s.Unlock()

	s.pairs = make(map[string][]byte)
}

package arctic

import (
	"log"
	"sync"
)

type store struct {
	sync.Mutex

	pairs map[string][]byte
}

func NewStore() *store {
	return &store{
		pairs: make(map[string][]byte),
	}
}

func (s *store) get(key string) []byte {
	s.Lock()
	defer s.Unlock()

	value, ok := s.pairs[key]
	if !ok {
		log.Printf("key %s not found\n", key)
		return nil
	}

	return value
}

func (s *store) put(key string, value []byte) {
	s.Lock()
	defer s.Unlock()

	_, ok := s.pairs[key]
	if ok {
		log.Printf("key %s exist\n", key)
	}

	s.pairs[key] = value
}

func (s *store) clear() {
	s.Lock()
	defer s.Unlock()

	s.pairs = make(map[string][]byte)
}

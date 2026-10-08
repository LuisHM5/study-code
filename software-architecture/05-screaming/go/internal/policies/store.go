package policies

import "sync"

// Store guarda las pólizas en memoria.
type Store struct {
	mu   sync.RWMutex
	byID map[string]Policy
}

func NewStore() *Store {
	return &Store{byID: map[string]Policy{}}
}

func (s *Store) numberTaken(number string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, p := range s.byID {
		if p.Number == number {
			return true
		}
	}
	return false
}

func (s *Store) save(p Policy) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.byID[p.ID] = p
}

func (s *Store) byIDOrNil(id string) (Policy, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.byID[id]
	return p, ok
}

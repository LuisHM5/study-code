package policies

import (
	"crypto/rand"
	"encoding/hex"
	"sync"

	"arquitecturas/11-event-driven-event-sourcing/internal/platform/eventbus"
)

type Service struct {
	store *EventStore
	bus   *eventbus.Bus

	mu      sync.Mutex
	numbers map[string]bool // índice para la regla de número único
}

func NewService(store *EventStore, bus *eventbus.Bus) *Service {
	return &Service{store: store, bus: bus, numbers: map[string]bool{}}
}

func (s *Service) Issue(number string, start, end Date) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.numbers[number] {
		return "", ErrDuplicated
	}

	p, err := Issue(newID(), number, start, end)
	if err != nil {
		return "", err
	}
	if err := s.save(p); err != nil {
		return "", err
	}
	s.numbers[number] = true
	return p.ID, nil
}

func (s *Service) Renew(id string, months int) (*Policy, error) {
	p, err := s.Get(id)
	if err != nil {
		return nil, err
	}
	if err := p.Renew(months); err != nil {
		return nil, err
	}
	if err := s.save(p); err != nil {
		return nil, err
	}
	return p, nil
}

// Get reconstruye el estado actual desde los eventos.
func (s *Service) Get(id string) (*Policy, error) {
	stored := s.store.Read(id)
	if len(stored) == 0 {
		return nil, ErrNotFound
	}
	history := make([]Event, len(stored))
	for i, se := range stored {
		history[i] = se.Data
	}
	return Load(history), nil
}

func (s *Service) Events(id string) ([]StoredEvent, error) {
	stored := s.store.Read(id)
	if len(stored) == 0 {
		return nil, ErrNotFound
	}
	return stored, nil
}

// save guarda los eventos nuevos y después los publica para quien quiera reaccionar.
func (s *Service) save(p *Policy) error {
	expectedVersion := p.Version - len(p.changes)
	if err := s.store.Append(p.ID, expectedVersion, p.changes); err != nil {
		return err
	}
	for _, e := range p.changes {
		s.bus.Publish(e.EventType(), e)
	}
	p.changes = nil
	return nil
}

func newID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

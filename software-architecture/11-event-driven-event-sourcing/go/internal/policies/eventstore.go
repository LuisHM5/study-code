package policies

import (
	"errors"
	"sync"
)

var ErrConcurrency = errors.New("la póliza fue modificada por otra operación")

type StoredEvent struct {
	Version int    `json:"version"`
	Type    string `json:"type"`
	Data    Event  `json:"data"`
}

// EventStore guarda, por agregado, una lista de eventos a la que solo se le agrega.
type EventStore struct {
	mu      sync.RWMutex
	streams map[string][]StoredEvent
}

func NewEventStore() *EventStore {
	return &EventStore{streams: map[string][]StoredEvent{}}
}

// Append usa concurrencia optimista: falla si alguien escribió antes que nosotros.
func (s *EventStore) Append(streamID string, expectedVersion int, events []Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	stream := s.streams[streamID]
	if len(stream) != expectedVersion {
		return ErrConcurrency
	}
	for _, e := range events {
		stream = append(stream, StoredEvent{Version: len(stream) + 1, Type: e.EventType(), Data: e})
	}
	s.streams[streamID] = stream
	return nil
}

func (s *EventStore) Read(streamID string) []StoredEvent {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return append([]StoredEvent{}, s.streams[streamID]...)
}

// Package eventbus: bus en memoria y síncrono (nivel 1 de event-driven).
// Quien publica no sabe quién escucha.
package eventbus

import "sync"

type Handler func(event any)

type Bus struct {
	mu       sync.RWMutex
	handlers map[string][]Handler
}

func New() *Bus {
	return &Bus{handlers: map[string][]Handler{}}
}

func (b *Bus) Subscribe(topic string, h Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[topic] = append(b.handlers[topic], h)
}

func (b *Bus) Publish(topic string, event any) {
	b.mu.RLock()
	handlers := b.handlers[topic]
	b.mu.RUnlock()
	for _, h := range handlers {
		h(event)
	}
}

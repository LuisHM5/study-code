// Package memory es un adaptador de salida: implementa app.PolicyRepository en memoria.
// Cambiarlo por Postgres no toca ni el dominio ni los casos de uso.
package memory

import (
	"context"
	"sync"

	"arquitecturas/06-hexagonal/internal/policy/domain"
)

type Repository struct {
	mu   sync.RWMutex
	byID map[string]domain.Policy
}

func NewRepository() *Repository {
	return &Repository{byID: map[string]domain.Policy{}}
}

func (r *Repository) ExistsByNumber(_ context.Context, number string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, p := range r.byID {
		if p.Number == number {
			return true, nil
		}
	}
	return false, nil
}

func (r *Repository) Save(_ context.Context, p *domain.Policy) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.byID[p.ID] = *p
	return nil
}

func (r *Repository) FindByID(_ context.Context, id string) (*domain.Policy, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.byID[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return &p, nil
}

// Package memory: el módulo es dueño de sus datos; nadie más los toca.
package memory

import (
	"sync"

	"arquitecturas/09-modular-monolith/modules/policies/internal/domain"
)

type Repository struct {
	mu   sync.RWMutex
	byID map[string]domain.Policy
}

func NewRepository() *Repository {
	return &Repository{byID: map[string]domain.Policy{}}
}

func (r *Repository) ExistsByNumber(number string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, p := range r.byID {
		if p.Number == number {
			return true
		}
	}
	return false
}

func (r *Repository) Save(p *domain.Policy) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.byID[p.ID] = *p
}

func (r *Repository) FindByID(id string) (*domain.Policy, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.byID[id]
	if !ok {
		return nil, false
	}
	return &p, true
}

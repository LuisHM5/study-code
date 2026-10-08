package repositories

import (
	"sync"

	"arquitecturas/02-layered/internal/models"
)

// PolicyRepository guarda las pólizas en memoria.
type PolicyRepository struct {
	mu   sync.RWMutex
	byID map[string]models.Policy
}

func NewPolicyRepository() *PolicyRepository {
	return &PolicyRepository{byID: map[string]models.Policy{}}
}

func (r *PolicyRepository) ExistsByNumber(number string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, p := range r.byID {
		if p.Number == number {
			return true
		}
	}
	return false
}

func (r *PolicyRepository) Insert(p models.Policy) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.byID[p.ID] = p
}

func (r *PolicyRepository) FindByID(id string) (models.Policy, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.byID[id]
	return p, ok
}

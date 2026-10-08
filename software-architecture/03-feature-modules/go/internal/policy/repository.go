package policy

import "sync"

type Repository struct {
	mu   sync.RWMutex
	byID map[string]Policy
}

func NewRepository() *Repository {
	return &Repository{byID: map[string]Policy{}}
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

func (r *Repository) Insert(p Policy) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.byID[p.ID] = p
}

func (r *Repository) FindByID(id string) (Policy, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.byID[id]
	return p, ok
}

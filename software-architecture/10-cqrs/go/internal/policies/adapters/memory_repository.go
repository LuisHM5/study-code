package adapters

import (
	"context"

	"arquitecturas/10-cqrs/internal/policies/domain"
)

// MemoryRepository es el lado de ESCRITURA: trabaja con entidades de dominio.
type MemoryRepository struct {
	db *MemoryDB
}

func NewMemoryRepository(db *MemoryDB) *MemoryRepository {
	return &MemoryRepository{db: db}
}

func (r *MemoryRepository) ExistsByNumber(_ context.Context, number string) (bool, error) {
	r.db.mu.RLock()
	defer r.db.mu.RUnlock()
	for _, row := range r.db.rows {
		if row.number == number {
			return true, nil
		}
	}
	return false, nil
}

func (r *MemoryRepository) Save(_ context.Context, p *domain.Policy) error {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()
	r.db.rows[p.ID] = policyRow{id: p.ID, number: p.Number, startDate: p.StartDate, endDate: p.EndDate}
	return nil
}

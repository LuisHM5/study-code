package adapters

import (
	"context"
	"sort"
	"time"

	"arquitecturas/10-cqrs/internal/policies/app/query"
)

const dateLayout = "2006-01-02"

// MemoryReadModel es el lado de LECTURA: va directo a los datos y arma vistas,
// sin pasar por el dominio (equivale a un SELECT a la medida).
type MemoryReadModel struct {
	db *MemoryDB
}

func NewMemoryReadModel(db *MemoryDB) *MemoryReadModel {
	return &MemoryReadModel{db: db}
}

func (m *MemoryReadModel) PolicyByID(_ context.Context, id string) (*query.PolicyView, error) {
	m.db.mu.RLock()
	defer m.db.mu.RUnlock()
	row, ok := m.db.rows[id]
	if !ok {
		return nil, query.ErrNotFound
	}
	return &query.PolicyView{
		ID:        row.id,
		Number:    row.number,
		StartDate: row.startDate.Format(dateLayout),
		EndDate:   row.endDate.Format(dateLayout),
	}, nil
}

func (m *MemoryReadModel) ExpiringWithin(_ context.Context, days int) ([]query.ExpiringPolicy, error) {
	today := time.Now().UTC().Truncate(24 * time.Hour)
	limit := today.AddDate(0, 0, days)

	m.db.mu.RLock()
	defer m.db.mu.RUnlock()
	result := []query.ExpiringPolicy{}
	for _, row := range m.db.rows {
		if row.endDate.Before(today) || row.endDate.After(limit) {
			continue
		}
		result = append(result, query.ExpiringPolicy{
			ID:       row.id,
			Number:   row.number,
			EndDate:  row.endDate.Format(dateLayout),
			DaysLeft: int(row.endDate.Sub(today).Hours() / 24),
		})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].EndDate < result[j].EndDate })
	return result, nil
}

package adapters

import (
	"sync"
	"time"
)

// MemoryDB hace de base de datos. CQRS nivel 1: escritura y lectura
// comparten almacenamiento, pero cada lado tiene su propio código y modelo.
type MemoryDB struct {
	mu   sync.RWMutex
	rows map[string]policyRow
}

type policyRow struct {
	id        string
	number    string
	startDate time.Time
	endDate   time.Time
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{rows: map[string]policyRow{}}
}

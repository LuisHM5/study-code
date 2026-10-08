// Package policies: el nombre del paquete y de sus archivos habla del negocio
// (issuing, lookup), no de patrones técnicos (service, helper).
package policies

import (
	"errors"
	"time"
)

var (
	ErrInvalidTerm = errors.New("vigencia inválida")
	ErrDuplicated  = errors.New("el número de póliza ya existe")
	ErrNotFound    = errors.New("póliza no encontrada")
)

type Policy struct {
	ID        string
	Number    string
	StartDate time.Time
	EndDate   time.Time
}

type Notifier interface {
	PolicyIssued(number string)
}

// Desk es la "ventanilla" de pólizas: agrupa las operaciones de negocio.
type Desk struct {
	store    *Store
	notifier Notifier
}

func NewDesk(store *Store, notifier Notifier) *Desk {
	return &Desk{store: store, notifier: notifier}
}

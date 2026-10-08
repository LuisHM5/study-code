// Package policy contiene todo lo de la feature "pólizas".
package policy

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

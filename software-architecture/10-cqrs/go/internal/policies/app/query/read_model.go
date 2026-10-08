// Package query: solo leen. No usan el dominio; devuelven vistas hechas para la pantalla.
package query

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("póliza no encontrada")

type PolicyView struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type ExpiringPolicy struct {
	ID       string `json:"id"`
	Number   string `json:"number"`
	EndDate  string `json:"endDate"`
	DaysLeft int    `json:"daysLeft"`
}

// ReadModel es el lado de lectura; lo implementa un adaptador distinto al repositorio.
type ReadModel interface {
	PolicyByID(ctx context.Context, id string) (*PolicyView, error)
	ExpiringWithin(ctx context.Context, days int) ([]ExpiringPolicy, error)
}

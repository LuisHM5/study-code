// Package domain es el modelo de ESCRITURA: solo lo usan los comandos.
package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidTerm = errors.New("vigencia inválida")
	ErrDuplicated  = errors.New("el número de póliza ya existe")
)

type Policy struct {
	ID        string
	Number    string
	StartDate time.Time
	EndDate   time.Time
}

func Issue(id, number string, start, end time.Time) (*Policy, error) {
	if !end.After(start) {
		return nil, ErrInvalidTerm
	}
	return &Policy{ID: id, Number: number, StartDate: start, EndDate: end}, nil
}

type Repository interface {
	ExistsByNumber(ctx context.Context, number string) (bool, error)
	Save(ctx context.Context, p *Policy) error
}

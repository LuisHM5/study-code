package domain

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

func Issue(id, number string, start, end time.Time) (*Policy, error) {
	if !end.After(start) {
		return nil, ErrInvalidTerm
	}
	return &Policy{ID: id, Number: number, StartDate: start, EndDate: end}, nil
}

package domain

import "errors"

var (
	ErrInvalidTerm   = errors.New("vigencia inválida")
	ErrInvalidNumber = errors.New("el número de póliza es obligatorio")
	ErrDuplicated    = errors.New("el número de póliza ya existe")
	ErrNotFound      = errors.New("póliza no encontrada")
)

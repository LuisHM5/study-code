package domain

import "strings"

// Number es un value object: solo existe si es válido.
type Number struct {
	value string
}

func NewNumber(value string) (Number, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return Number{}, ErrInvalidNumber
	}
	return Number{value: value}, nil
}

func (n Number) String() string { return n.value }

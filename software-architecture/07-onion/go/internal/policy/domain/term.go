package domain

import "time"

// Term (vigencia) es un value object: los campos privados obligan a usar NewTerm.
type Term struct {
	start, end time.Time
}

func NewTerm(start, end time.Time) (Term, error) {
	if !end.After(start) {
		return Term{}, ErrInvalidTerm
	}
	return Term{start: start, end: end}, nil
}

func (t Term) Start() time.Time { return t.start }
func (t Term) End() time.Time   { return t.end }

func (t Term) RenewFor(months int) Term {
	return Term{start: t.end, end: t.end.AddDate(0, months, 0)}
}

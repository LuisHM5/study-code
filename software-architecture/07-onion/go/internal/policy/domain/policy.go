// Package domain es el anillo central: modelo, servicios de dominio e interfaces.
package domain

// Policy es la entidad; el comportamiento vive aquí, no en un servicio.
type Policy struct {
	id     string
	number Number
	term   Term
}

func NewPolicy(id string, number Number, term Term) *Policy {
	return &Policy{id: id, number: number, term: term}
}

func (p *Policy) ID() string     { return p.id }
func (p *Policy) Number() Number { return p.number }
func (p *Policy) Term() Term     { return p.term }

func (p *Policy) Renew(months int) {
	p.term = p.term.RenewFor(months)
}

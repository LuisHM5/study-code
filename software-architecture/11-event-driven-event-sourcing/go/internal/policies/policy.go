package policies

import "errors"

var (
	ErrInvalidTerm   = errors.New("vigencia inválida")
	ErrInvalidMonths = errors.New("months debe ser mayor a 0")
	ErrDuplicated    = errors.New("el número de póliza ya existe")
	ErrNotFound      = errors.New("póliza no encontrada")
)

// Policy es el agregado. Su estado NO se guarda: se reconstruye aplicando eventos.
type Policy struct {
	ID        string
	Number    string
	StartDate Date
	EndDate   Date
	Version   int

	changes []Event // eventos nuevos, aún sin guardar
}

// Decidir: validar y registrar un evento.

func Issue(id, number string, start, end Date) (*Policy, error) {
	if !end.After(start) {
		return nil, ErrInvalidTerm
	}
	p := &Policy{}
	p.record(Issued{PolicyID: id, Number: number, StartDate: start, EndDate: end})
	return p, nil
}

func (p *Policy) Renew(months int) error {
	if months <= 0 {
		return ErrInvalidMonths
	}
	p.record(Renewed{PolicyID: p.ID, Months: months, NewEndDate: p.EndDate.AddMonths(months)})
	return nil
}

// Aplicar: el único lugar donde cambia el estado, al ejecutar y al reproducir.
func (p *Policy) apply(e Event) {
	switch e := e.(type) {
	case Issued:
		p.ID, p.Number, p.StartDate, p.EndDate = e.PolicyID, e.Number, e.StartDate, e.EndDate
	case Renewed:
		p.EndDate = e.NewEndDate
	}
	p.Version++
}

func (p *Policy) record(e Event) {
	p.apply(e)
	p.changes = append(p.changes, e)
}

// Load reconstruye el agregado reproduciendo su historial.
func Load(history []Event) *Policy {
	p := &Policy{}
	for _, e := range history {
		p.apply(e)
	}
	return p
}

package domain

import "context"

// Uniqueness es un servicio de dominio: una regla que la entidad sola no puede verificar.
type Uniqueness struct {
	Policies Repository
}

func (u Uniqueness) Assert(ctx context.Context, number Number) error {
	exists, err := u.Policies.ExistsByNumber(ctx, number)
	if err != nil {
		return err
	}
	if exists {
		return ErrDuplicated
	}
	return nil
}

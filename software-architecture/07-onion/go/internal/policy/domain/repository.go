package domain

import "context"

// Repository lo define el dominio; lo implementa el anillo de infraestructura.
type Repository interface {
	ExistsByNumber(ctx context.Context, number Number) (bool, error)
	Save(ctx context.Context, p *Policy) error
	FindByID(ctx context.Context, id string) (*Policy, error)
}

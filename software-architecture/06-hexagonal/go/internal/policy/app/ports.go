package app

import (
	"context"

	"arquitecturas/06-hexagonal/internal/policy/domain"
)

// Puertos de salida: los define el núcleo y los implementan los adaptadores.

type PolicyRepository interface {
	ExistsByNumber(ctx context.Context, number string) (bool, error)
	Save(ctx context.Context, p *domain.Policy) error
	FindByID(ctx context.Context, id string) (*domain.Policy, error)
}

type Notifier interface {
	PolicyIssued(ctx context.Context, p *domain.Policy) error
}

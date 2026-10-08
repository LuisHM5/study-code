// Package usecase: círculo 2. Solo conoce entity y las interfaces que él mismo define.
package usecase

import (
	"context"

	"arquitecturas/08-clean/internal/policies/entity"
)

// Gateways: lo que los casos de uso necesitan de afuera.

type PolicyGateway interface {
	ExistsByNumber(ctx context.Context, number string) (bool, error)
	Save(ctx context.Context, p *entity.Policy) error
	FindByID(ctx context.Context, id string) (*entity.Policy, error)
}

type NotificationGateway interface {
	PolicyIssued(ctx context.Context, number string) error
}

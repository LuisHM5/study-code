package app

import (
	"context"

	"arquitecturas/06-hexagonal/internal/policy/domain"
)

type GetPolicy struct {
	Policies PolicyRepository
}

func (uc GetPolicy) Execute(ctx context.Context, id string) (*domain.Policy, error) {
	return uc.Policies.FindByID(ctx, id)
}

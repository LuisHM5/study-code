package application

import (
	"context"

	"arquitecturas/07-onion/internal/policy/domain"
)

type GetPolicy struct {
	policies domain.Repository
}

func NewGetPolicy(policies domain.Repository) *GetPolicy {
	return &GetPolicy{policies: policies}
}

func (s *GetPolicy) Execute(ctx context.Context, id string) (*domain.Policy, error) {
	return s.policies.FindByID(ctx, id)
}

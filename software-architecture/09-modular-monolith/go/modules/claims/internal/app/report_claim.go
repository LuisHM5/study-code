package app

import (
	"context"
	"errors"
	"sync"

	"arquitecturas/09-modular-monolith/modules/policies"
	"arquitecturas/09-modular-monolith/platform/web"
)

var ErrPolicyNotFound = errors.New("la póliza no existe")

// Claim guarda policyID como dato; no hay relación directa con la entidad de policies.
type Claim struct {
	ID          string
	PolicyID    string
	Description string
}

type Service struct {
	policies policies.API // solo la API pública del otro módulo

	mu     sync.Mutex
	claims map[string]Claim
}

func NewService(policiesAPI policies.API) *Service {
	return &Service{policies: policiesAPI, claims: map[string]Claim{}}
}

func (s *Service) Report(ctx context.Context, policyID, description string) (string, error) {
	policy, err := s.policies.FindByID(ctx, policyID)
	if err != nil {
		return "", err
	}
	if policy == nil {
		return "", ErrPolicyNotFound
	}

	claim := Claim{ID: web.NewID(), PolicyID: policy.ID, Description: description}
	s.mu.Lock()
	s.claims[claim.ID] = claim
	s.mu.Unlock()
	return claim.ID, nil
}

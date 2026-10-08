// Package application: servicios de aplicación. Orquestan, no deciden reglas.
package application

import (
	"context"
	"time"

	"arquitecturas/07-onion/internal/policy/domain"
)

type Notifier interface {
	PolicyIssued(ctx context.Context, p *domain.Policy) error
}

type IssuePolicyCommand struct {
	Number    string
	StartDate time.Time
	EndDate   time.Time
}

type IssuePolicy struct {
	policies   domain.Repository
	uniqueness domain.Uniqueness
	notifier   Notifier
	newID      func() string
}

func NewIssuePolicy(policies domain.Repository, notifier Notifier, newID func() string) *IssuePolicy {
	return &IssuePolicy{
		policies:   policies,
		uniqueness: domain.Uniqueness{Policies: policies},
		notifier:   notifier,
		newID:      newID,
	}
}

func (s *IssuePolicy) Execute(ctx context.Context, cmd IssuePolicyCommand) (string, error) {
	number, err := domain.NewNumber(cmd.Number)
	if err != nil {
		return "", err
	}
	term, err := domain.NewTerm(cmd.StartDate, cmd.EndDate)
	if err != nil {
		return "", err
	}
	if err := s.uniqueness.Assert(ctx, number); err != nil {
		return "", err
	}

	policy := domain.NewPolicy(s.newID(), number, term)
	if err := s.policies.Save(ctx, policy); err != nil {
		return "", err
	}
	return policy.ID(), s.notifier.PolicyIssued(ctx, policy)
}

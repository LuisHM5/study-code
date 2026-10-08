package app

import (
	"context"
	"time"

	"arquitecturas/06-hexagonal/internal/policy/domain"
)

// IssuePolicy es un caso de uso: orquesta dominio y puertos.
type IssuePolicy struct {
	Policies PolicyRepository
	Notifier Notifier
	NewID    func() string
}

func (uc IssuePolicy) Execute(ctx context.Context, number string, start, end time.Time) (string, error) {
	exists, err := uc.Policies.ExistsByNumber(ctx, number)
	if err != nil {
		return "", err
	}
	if exists {
		return "", domain.ErrDuplicated
	}

	p, err := domain.Issue(uc.NewID(), number, start, end)
	if err != nil {
		return "", err
	}
	if err := uc.Policies.Save(ctx, p); err != nil {
		return "", err
	}
	return p.ID, uc.Notifier.PolicyIssued(ctx, p)
}

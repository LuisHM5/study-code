// Package command: cambian estado y pasan por el dominio. Devuelven a lo mucho un id.
package command

import (
	"context"
	"time"

	"arquitecturas/10-cqrs/internal/policies/domain"
)

type Notifier interface {
	PolicyIssued(ctx context.Context, number string) error
}

type IssuePolicy struct {
	Number    string
	StartDate time.Time
	EndDate   time.Time
}

type IssuePolicyHandler struct {
	Policies domain.Repository
	Notifier Notifier
	NewID    func() string
}

func (h IssuePolicyHandler) Handle(ctx context.Context, cmd IssuePolicy) (string, error) {
	exists, err := h.Policies.ExistsByNumber(ctx, cmd.Number)
	if err != nil {
		return "", err
	}
	if exists {
		return "", domain.ErrDuplicated
	}

	p, err := domain.Issue(h.NewID(), cmd.Number, cmd.StartDate, cmd.EndDate)
	if err != nil {
		return "", err
	}
	if err := h.Policies.Save(ctx, p); err != nil {
		return "", err
	}
	return p.ID, h.Notifier.PolicyIssued(ctx, p.Number)
}

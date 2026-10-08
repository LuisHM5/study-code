package usecase

import (
	"context"
	"time"

	"arquitecturas/08-clean/internal/policies/entity"
)

// Modelos propios del caso de uso: no son el JSON de HTTP ni la entidad.
type IssuePolicyInput struct {
	Number    string
	StartDate time.Time
	EndDate   time.Time
}

type IssuePolicyOutput struct {
	ID     string
	Number string
}

// IssuePolicy es el input boundary: lo único que ve el controlador.
type IssuePolicy interface {
	Execute(ctx context.Context, in IssuePolicyInput) (IssuePolicyOutput, error)
}

type issuePolicyInteractor struct {
	policies PolicyGateway
	notifier NotificationGateway
	newID    func() string
}

func NewIssuePolicy(policies PolicyGateway, notifier NotificationGateway, newID func() string) IssuePolicy {
	return &issuePolicyInteractor{policies: policies, notifier: notifier, newID: newID}
}

func (uc *issuePolicyInteractor) Execute(ctx context.Context, in IssuePolicyInput) (IssuePolicyOutput, error) {
	exists, err := uc.policies.ExistsByNumber(ctx, in.Number)
	if err != nil {
		return IssuePolicyOutput{}, err
	}
	if exists {
		return IssuePolicyOutput{}, entity.ErrDuplicated
	}

	p, err := entity.NewPolicy(uc.newID(), in.Number, in.StartDate, in.EndDate)
	if err != nil {
		return IssuePolicyOutput{}, err
	}
	if err := uc.policies.Save(ctx, p); err != nil {
		return IssuePolicyOutput{}, err
	}
	if err := uc.notifier.PolicyIssued(ctx, p.Number); err != nil {
		return IssuePolicyOutput{}, err
	}
	return IssuePolicyOutput{ID: p.ID, Number: p.Number}, nil
}

package usecase

import (
	"context"
	"time"
)

type GetPolicyOutput struct {
	ID        string
	Number    string
	StartDate time.Time
	EndDate   time.Time
}

type GetPolicy interface {
	Execute(ctx context.Context, id string) (GetPolicyOutput, error)
}

type getPolicyInteractor struct {
	policies PolicyGateway
}

func NewGetPolicy(policies PolicyGateway) GetPolicy {
	return &getPolicyInteractor{policies: policies}
}

func (uc *getPolicyInteractor) Execute(ctx context.Context, id string) (GetPolicyOutput, error) {
	p, err := uc.policies.FindByID(ctx, id)
	if err != nil {
		return GetPolicyOutput{}, err
	}
	return GetPolicyOutput{ID: p.ID, Number: p.Number, StartDate: p.StartDate, EndDate: p.EndDate}, nil
}

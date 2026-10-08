package query

import "context"

type GetPolicy struct {
	ID string
}

type GetPolicyHandler struct {
	ReadModel ReadModel
}

func (h GetPolicyHandler) Handle(ctx context.Context, q GetPolicy) (*PolicyView, error) {
	return h.ReadModel.PolicyByID(ctx, q.ID)
}

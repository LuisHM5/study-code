package query

import "context"

type ListExpiring struct {
	Days int
}

type ListExpiringHandler struct {
	ReadModel ReadModel
}

func (h ListExpiringHandler) Handle(ctx context.Context, q ListExpiring) ([]ExpiringPolicy, error) {
	return h.ReadModel.ExpiringWithin(ctx, q.Days)
}

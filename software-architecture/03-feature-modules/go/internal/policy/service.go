package policy

import (
	"time"

	"arquitecturas/03-feature-modules/internal/platform"
)

// Notifier se declara donde se consume: policy no importa el paquete notification.
type Notifier interface {
	PolicyIssued(number string)
}

type Service struct {
	repo     *Repository
	notifier Notifier
}

func NewService(repo *Repository, notifier Notifier) *Service {
	return &Service{repo: repo, notifier: notifier}
}

func (s *Service) Issue(number string, start, end time.Time) (Policy, error) {
	if !end.After(start) {
		return Policy{}, ErrInvalidTerm
	}
	if s.repo.ExistsByNumber(number) {
		return Policy{}, ErrDuplicated
	}

	p := Policy{ID: platform.NewID(), Number: number, StartDate: start, EndDate: end}
	s.repo.Insert(p)
	s.notifier.PolicyIssued(p.Number)
	return p, nil
}

func (s *Service) GetByID(id string) (Policy, error) {
	p, ok := s.repo.FindByID(id)
	if !ok {
		return Policy{}, ErrNotFound
	}
	return p, nil
}

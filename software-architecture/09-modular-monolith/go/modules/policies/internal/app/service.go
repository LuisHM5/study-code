package app

import (
	"time"

	"arquitecturas/09-modular-monolith/modules/policies/internal/domain"
	"arquitecturas/09-modular-monolith/platform/web"
)

type Repository interface {
	ExistsByNumber(number string) bool
	Save(p *domain.Policy)
	FindByID(id string) (*domain.Policy, bool)
}

type Service struct {
	repo          Repository
	publishIssued func(policyID, number string)
}

func NewService(repo Repository, publishIssued func(policyID, number string)) *Service {
	return &Service{repo: repo, publishIssued: publishIssued}
}

func (s *Service) Issue(number string, start, end time.Time) (string, error) {
	if s.repo.ExistsByNumber(number) {
		return "", domain.ErrDuplicated
	}
	p, err := domain.Issue(web.NewID(), number, start, end)
	if err != nil {
		return "", err
	}
	s.repo.Save(p)
	// El módulo no sabe quién escucha: solo anuncia el hecho.
	s.publishIssued(p.ID, p.Number)
	return p.ID, nil
}

func (s *Service) Get(id string) (*domain.Policy, error) {
	p, ok := s.repo.FindByID(id)
	if !ok {
		return nil, domain.ErrNotFound
	}
	return p, nil
}

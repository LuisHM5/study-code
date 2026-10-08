package app

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"arquitecturas/12-microservices/services/policies/internal/domain"
)

type Repository interface {
	ExistsByNumber(number string) bool
	Save(p *domain.Policy)
	FindByID(id string) (*domain.Policy, bool)
}

// Notifier es un puerto: por dentro el servicio sigue siendo hexagonal.
// El otro microservicio es solo un adaptador más.
type Notifier interface {
	PolicyIssued(ctx context.Context, p *domain.Policy) error
}

type Service struct {
	repo     Repository
	notifier Notifier
}

func NewService(repo Repository, notifier Notifier) *Service {
	return &Service{repo: repo, notifier: notifier}
}

func (s *Service) Issue(ctx context.Context, number string, start, end time.Time) (string, error) {
	if s.repo.ExistsByNumber(number) {
		return "", domain.ErrDuplicated
	}
	p, err := domain.Issue(newID(), number, start, end)
	if err != nil {
		return "", err
	}
	s.repo.Save(p)

	// La red puede fallar: la póliza ya quedó emitida, así que solo se registra el error.
	if err := s.notifier.PolicyIssued(ctx, p); err != nil {
		log.Printf("no se pudo notificar la póliza %s: %v", p.Number, err)
	}
	return p.ID, nil
}

func (s *Service) Get(id string) (*domain.Policy, error) {
	p, ok := s.repo.FindByID(id)
	if !ok {
		return nil, domain.ErrNotFound
	}
	return p, nil
}

func newID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

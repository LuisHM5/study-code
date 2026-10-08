package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"arquitecturas/02-layered/internal/models"
	"arquitecturas/02-layered/internal/repositories"
)

var (
	ErrInvalidTerm = errors.New("vigencia inválida")
	ErrDuplicated  = errors.New("el número de póliza ya existe")
	ErrNotFound    = errors.New("póliza no encontrada")
)

// PolicyService contiene la lógica; solo conoce la capa de abajo (repositorio).
type PolicyService struct {
	repo *repositories.PolicyRepository
}

func NewPolicyService(repo *repositories.PolicyRepository) *PolicyService {
	return &PolicyService{repo: repo}
}

func (s *PolicyService) Issue(number string, start, end time.Time) (models.Policy, error) {
	if !end.After(start) {
		return models.Policy{}, ErrInvalidTerm
	}
	if s.repo.ExistsByNumber(number) {
		return models.Policy{}, ErrDuplicated
	}

	p := models.Policy{ID: newID(), Number: number, StartDate: start, EndDate: end}
	s.repo.Insert(p)
	fmt.Printf("[notificación] póliza emitida %s\n", p.Number)
	return p, nil
}

func (s *PolicyService) GetByID(id string) (models.Policy, error) {
	p, ok := s.repo.FindByID(id)
	if !ok {
		return models.Policy{}, ErrNotFound
	}
	return p, nil
}

func newID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

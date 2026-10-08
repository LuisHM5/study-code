// Package gateway: círculo 3. Implementa los gateways que piden los casos de uso.
package gateway

import (
	"context"
	"sync"

	"arquitecturas/08-clean/internal/policies/entity"
)

type MemoryPolicyGateway struct {
	mu   sync.RWMutex
	byID map[string]entity.Policy
}

func NewMemoryPolicyGateway() *MemoryPolicyGateway {
	return &MemoryPolicyGateway{byID: map[string]entity.Policy{}}
}

func (g *MemoryPolicyGateway) ExistsByNumber(_ context.Context, number string) (bool, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	for _, p := range g.byID {
		if p.Number == number {
			return true, nil
		}
	}
	return false, nil
}

func (g *MemoryPolicyGateway) Save(_ context.Context, p *entity.Policy) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.byID[p.ID] = *p
	return nil
}

func (g *MemoryPolicyGateway) FindByID(_ context.Context, id string) (*entity.Policy, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	p, ok := g.byID[id]
	if !ok {
		return nil, entity.ErrNotFound
	}
	return &p, nil
}

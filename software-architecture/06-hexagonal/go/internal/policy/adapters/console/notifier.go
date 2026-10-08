// Package console es un adaptador de salida: implementa app.Notifier imprimiendo.
package console

import (
	"context"
	"fmt"

	"arquitecturas/06-hexagonal/internal/policy/domain"
)

type Notifier struct{}

func (Notifier) PolicyIssued(_ context.Context, p *domain.Policy) error {
	fmt.Printf("[notificación] póliza emitida %s\n", p.Number)
	return nil
}

package console

import (
	"context"
	"fmt"

	"arquitecturas/07-onion/internal/policy/domain"
)

type Notifier struct{}

func (Notifier) PolicyIssued(_ context.Context, p *domain.Policy) error {
	fmt.Printf("[notificación] póliza emitida %s\n", p.Number())
	return nil
}

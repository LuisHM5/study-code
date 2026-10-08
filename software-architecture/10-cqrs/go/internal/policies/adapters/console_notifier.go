package adapters

import (
	"context"
	"fmt"
)

type ConsoleNotifier struct{}

func (ConsoleNotifier) PolicyIssued(_ context.Context, number string) error {
	fmt.Printf("[notificación] póliza emitida %s\n", number)
	return nil
}

package gateway

import (
	"context"
	"fmt"
)

type ConsoleNotificationGateway struct{}

func (ConsoleNotificationGateway) PolicyIssued(_ context.Context, number string) error {
	fmt.Printf("[notificación] póliza emitida %s\n", number)
	return nil
}

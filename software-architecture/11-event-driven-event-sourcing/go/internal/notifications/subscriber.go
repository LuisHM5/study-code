// Package notifications reacciona a eventos de pólizas sin que policies lo conozca.
package notifications

import (
	"fmt"

	"arquitecturas/11-event-driven-event-sourcing/internal/platform/eventbus"
	"arquitecturas/11-event-driven-event-sourcing/internal/policies"
)

func Subscribe(bus *eventbus.Bus) {
	bus.Subscribe(policies.Issued{}.EventType(), func(event any) {
		if issued, ok := event.(policies.Issued); ok {
			fmt.Printf("[notificación] póliza emitida %s\n", issued.Number)
		}
	})
}

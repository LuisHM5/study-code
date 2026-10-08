// Package notifications reacciona a eventos; policies no sabe que existe.
package notifications

import (
	"fmt"

	"arquitecturas/09-modular-monolith/modules/policies"
	"arquitecturas/09-modular-monolith/platform/eventbus"
)

func New(bus *eventbus.Bus) {
	bus.Subscribe(policies.TopicIssued, func(event any) {
		issued, ok := event.(policies.Issued)
		if !ok {
			return
		}
		fmt.Printf("[notificación] póliza emitida %s\n", issued.Number)
	})
}

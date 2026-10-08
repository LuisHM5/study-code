package mailer

import "fmt"

// Console simula el envío de correo imprimiendo en consola.
type Console struct{}

func (Console) PolicyIssued(number string) {
	fmt.Printf("[notificación] póliza emitida %s\n", number)
}

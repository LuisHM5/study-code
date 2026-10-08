// Package notification es otra feature; policy la usa a través de una interfaz.
package notification

import "fmt"

type Service struct{}

func NewService() *Service { return &Service{} }

func (s *Service) PolicyIssued(number string) {
	fmt.Printf("[notificación] póliza emitida %s\n", number)
}

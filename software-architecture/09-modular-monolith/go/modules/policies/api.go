// Package policies expone la API pública del módulo de pólizas.
// Todo lo demás vive en internal/ y el compilador impide importarlo desde otro módulo.
package policies

import "context"

// Summary es un contrato público: nunca se expone la entidad interna.
type Summary struct {
	ID     string
	Number string
}

type API interface {
	// FindByID devuelve nil si la póliza no existe.
	FindByID(ctx context.Context, id string) (*Summary, error)
}

// Evento público del módulo.
const TopicIssued = "policy.issued"

type Issued struct {
	PolicyID string
	Number   string
}

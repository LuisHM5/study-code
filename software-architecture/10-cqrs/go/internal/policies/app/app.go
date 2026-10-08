// Package app agrupa los handlers. CQRS sin bus: el handler se llama directo.
package app

import (
	"arquitecturas/10-cqrs/internal/policies/app/command"
	"arquitecturas/10-cqrs/internal/policies/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	IssuePolicy command.IssuePolicyHandler
}

type Queries struct {
	GetPolicy    query.GetPolicyHandler
	ListExpiring query.ListExpiringHandler
}

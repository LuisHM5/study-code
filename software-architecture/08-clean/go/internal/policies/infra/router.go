// Package infra: círculo 4. El único lugar que conoce todas las piezas y las arma.
package infra

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"arquitecturas/08-clean/internal/policies/adapter/controller"
	"arquitecturas/08-clean/internal/policies/adapter/gateway"
	"arquitecturas/08-clean/internal/policies/usecase"
)

func NewRouter() *http.ServeMux {
	policies := gateway.NewMemoryPolicyGateway()
	notifier := gateway.ConsoleNotificationGateway{}

	ctrl := controller.NewPolicyController(
		usecase.NewIssuePolicy(policies, notifier, newID),
		usecase.NewGetPolicy(policies),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /policies", ctrl.Issue)
	mux.HandleFunc("GET /policies/{id}", ctrl.Get)
	return mux
}

func newID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

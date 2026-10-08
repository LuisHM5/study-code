// Package claims es el módulo de siniestros. Depende solo de policies.API.
package claims

import (
	"net/http"

	"arquitecturas/09-modular-monolith/modules/claims/internal/app"
	"arquitecturas/09-modular-monolith/modules/claims/internal/httpapi"
	"arquitecturas/09-modular-monolith/modules/policies"
)

type Module struct {
	Routes func(mux *http.ServeMux)
}

func New(policiesAPI policies.API) *Module {
	svc := app.NewService(policiesAPI)
	return &Module{Routes: func(mux *http.ServeMux) { httpapi.Register(mux, svc) }}
}

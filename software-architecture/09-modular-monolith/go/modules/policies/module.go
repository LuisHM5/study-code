package policies

import (
	"context"
	"errors"
	"net/http"

	"arquitecturas/09-modular-monolith/modules/policies/internal/app"
	"arquitecturas/09-modular-monolith/modules/policies/internal/domain"
	"arquitecturas/09-modular-monolith/modules/policies/internal/httpapi"
	"arquitecturas/09-modular-monolith/modules/policies/internal/memory"
	"arquitecturas/09-modular-monolith/platform/eventbus"
)

type Module struct {
	API    API
	Routes func(mux *http.ServeMux)
}

func New(bus *eventbus.Bus) *Module {
	publishIssued := func(policyID, number string) {
		bus.Publish(TopicIssued, Issued{PolicyID: policyID, Number: number})
	}
	svc := app.NewService(memory.NewRepository(), publishIssued)

	return &Module{
		API:    facade{svc: svc},
		Routes: func(mux *http.ServeMux) { httpapi.Register(mux, svc) },
	}
}

// facade traduce el modelo interno al contrato público.
type facade struct {
	svc *app.Service
}

func (f facade) FindByID(_ context.Context, id string) (*Summary, error) {
	p, err := f.svc.Get(id)
	if errors.Is(err, domain.ErrNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &Summary{ID: p.ID, Number: p.Number}, nil
}

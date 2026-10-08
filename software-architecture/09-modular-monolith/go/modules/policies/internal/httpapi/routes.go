package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"arquitecturas/09-modular-monolith/modules/policies/internal/app"
	"arquitecturas/09-modular-monolith/modules/policies/internal/domain"
	"arquitecturas/09-modular-monolith/platform/web"
)

const dateLayout = "2006-01-02"

func Register(mux *http.ServeMux, svc *app.Service) {
	mux.HandleFunc("POST /policies", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Number    string `json:"number"`
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			web.WriteError(w, http.StatusBadRequest, "json inválido")
			return
		}
		start, errStart := time.Parse(dateLayout, req.StartDate)
		end, errEnd := time.Parse(dateLayout, req.EndDate)
		if req.Number == "" || errStart != nil || errEnd != nil {
			web.WriteError(w, http.StatusBadRequest, "number, startDate y endDate (YYYY-MM-DD) son obligatorios")
			return
		}

		id, err := svc.Issue(req.Number, start, end)
		switch {
		case errors.Is(err, domain.ErrInvalidTerm):
			web.WriteError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, domain.ErrDuplicated):
			web.WriteError(w, http.StatusConflict, err.Error())
		case err != nil:
			web.WriteError(w, http.StatusInternalServerError, "error interno")
		default:
			web.WriteJSON(w, http.StatusCreated, map[string]string{"id": id})
		}
	})

	mux.HandleFunc("GET /policies/{id}", func(w http.ResponseWriter, r *http.Request) {
		p, err := svc.Get(r.PathValue("id"))
		if err != nil {
			web.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		web.WriteJSON(w, http.StatusOK, map[string]string{
			"id":        p.ID,
			"number":    p.Number,
			"startDate": p.StartDate.Format(dateLayout),
			"endDate":   p.EndDate.Format(dateLayout),
		})
	})
}

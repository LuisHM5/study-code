package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"arquitecturas/12-microservices/services/policies/internal/app"
	"arquitecturas/12-microservices/services/policies/internal/domain"
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
			writeError(w, http.StatusBadRequest, "json inválido")
			return
		}
		start, errStart := time.Parse(dateLayout, req.StartDate)
		end, errEnd := time.Parse(dateLayout, req.EndDate)
		if req.Number == "" || errStart != nil || errEnd != nil {
			writeError(w, http.StatusBadRequest, "number, startDate y endDate (YYYY-MM-DD) son obligatorios")
			return
		}

		id, err := svc.Issue(r.Context(), req.Number, start, end)
		switch {
		case errors.Is(err, domain.ErrInvalidTerm):
			writeError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, domain.ErrDuplicated):
			writeError(w, http.StatusConflict, err.Error())
		case err != nil:
			writeError(w, http.StatusInternalServerError, "error interno")
		default:
			writeJSON(w, http.StatusCreated, map[string]string{"id": id})
		}
	})

	mux.HandleFunc("GET /policies/{id}", func(w http.ResponseWriter, r *http.Request) {
		p, err := svc.Get(r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{
			"id":        p.ID,
			"number":    p.Number,
			"startDate": p.StartDate.Format(dateLayout),
			"endDate":   p.EndDate.Format(dateLayout),
		})
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"message": message})
}

package policies

import (
	"encoding/json"
	"errors"
	"net/http"
)

func Register(mux *http.ServeMux, svc *Service) {
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
		start, errStart := ParseDate(req.StartDate)
		end, errEnd := ParseDate(req.EndDate)
		if req.Number == "" || errStart != nil || errEnd != nil {
			writeError(w, http.StatusBadRequest, "number, startDate y endDate (YYYY-MM-DD) son obligatorios")
			return
		}

		id, err := svc.Issue(req.Number, start, end)
		if err != nil {
			writeDomainError(w, err)
			return
		}
		writeJSON(w, http.StatusCreated, map[string]string{"id": id})
	})

	mux.HandleFunc("POST /policies/{id}/renew", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Months int `json:"months"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "json inválido")
			return
		}
		p, err := svc.Renew(r.PathValue("id"), req.Months)
		if err != nil {
			writeDomainError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, view(p))
	})

	mux.HandleFunc("GET /policies/{id}", func(w http.ResponseWriter, r *http.Request) {
		p, err := svc.Get(r.PathValue("id"))
		if err != nil {
			writeDomainError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, view(p))
	})

	mux.HandleFunc("GET /policies/{id}/events", func(w http.ResponseWriter, r *http.Request) {
		events, err := svc.Events(r.PathValue("id"))
		if err != nil {
			writeDomainError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, events)
	})
}

func view(p *Policy) map[string]string {
	return map[string]string{
		"id":        p.ID,
		"number":    p.Number,
		"startDate": p.StartDate.String(),
		"endDate":   p.EndDate.String(),
	}
}

func writeDomainError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrInvalidTerm), errors.Is(err, ErrInvalidMonths):
		writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, ErrDuplicated), errors.Is(err, ErrConcurrency):
		writeError(w, http.StatusConflict, err.Error())
	case errors.Is(err, ErrNotFound):
		writeError(w, http.StatusNotFound, err.Error())
	default:
		writeError(w, http.StatusInternalServerError, "error interno")
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"message": message})
}

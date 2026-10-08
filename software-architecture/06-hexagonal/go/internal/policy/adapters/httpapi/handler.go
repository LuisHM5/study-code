// Package httpapi es un adaptador de entrada: traduce HTTP a casos de uso.
package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"arquitecturas/06-hexagonal/internal/policy/app"
	"arquitecturas/06-hexagonal/internal/policy/domain"
)

const dateLayout = "2006-01-02"

type Handler struct {
	IssuePolicy app.IssuePolicy
	GetPolicy   app.GetPolicy
}

func (h Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /policies", h.issue)
	mux.HandleFunc("GET /policies/{id}", h.get)
}

func (h Handler) issue(w http.ResponseWriter, r *http.Request) {
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

	id, err := h.IssuePolicy.Execute(r.Context(), req.Number, start, end)
	if err != nil {
		writeDomainError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{"id": id})
}

func (h Handler) get(w http.ResponseWriter, r *http.Request) {
	p, err := h.GetPolicy.Execute(r.Context(), r.PathValue("id"))
	if err != nil {
		writeDomainError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"id":        p.ID,
		"number":    p.Number,
		"startDate": p.StartDate.Format(dateLayout),
		"endDate":   p.EndDate.Format(dateLayout),
	})
}

func writeDomainError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidTerm):
		writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, domain.ErrDuplicated):
		writeError(w, http.StatusConflict, err.Error())
	case errors.Is(err, domain.ErrNotFound):
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

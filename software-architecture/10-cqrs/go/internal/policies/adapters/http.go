package adapters

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"arquitecturas/10-cqrs/internal/policies/app"
	"arquitecturas/10-cqrs/internal/policies/app/command"
	"arquitecturas/10-cqrs/internal/policies/app/query"
	"arquitecturas/10-cqrs/internal/policies/domain"
)

type Server struct {
	app app.Application
}

func NewServer(application app.Application) *Server {
	return &Server{app: application}
}

func (s *Server) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /policies", s.issuePolicy)
	mux.HandleFunc("GET /policies/expiring", s.listExpiring)
	mux.HandleFunc("GET /policies/{id}", s.getPolicy)
}

func (s *Server) issuePolicy(w http.ResponseWriter, r *http.Request) {
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

	id, err := s.app.Commands.IssuePolicy.Handle(r.Context(), command.IssuePolicy{
		Number: req.Number, StartDate: start, EndDate: end,
	})
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
}

func (s *Server) getPolicy(w http.ResponseWriter, r *http.Request) {
	view, err := s.app.Queries.GetPolicy.Handle(r.Context(), query.GetPolicy{ID: r.PathValue("id")})
	if errors.Is(err, query.ErrNotFound) {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error interno")
		return
	}
	writeJSON(w, http.StatusOK, view)
}

func (s *Server) listExpiring(w http.ResponseWriter, r *http.Request) {
	days := 30
	if raw := r.URL.Query().Get("days"); raw != "" {
		parsed, err := strconv.Atoi(raw)
		if err != nil || parsed < 0 {
			writeError(w, http.StatusBadRequest, "days debe ser un entero mayor o igual a 0")
			return
		}
		days = parsed
	}

	list, err := s.app.Queries.ListExpiring.Handle(r.Context(), query.ListExpiring{Days: days})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error interno")
		return
	}
	writeJSON(w, http.StatusOK, list)
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"message": message})
}

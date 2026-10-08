package policy

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"arquitecturas/03-feature-modules/internal/platform"
)

const dateLayout = "2006-01-02"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Register es la única puerta de entrada HTTP de la feature.
func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /policies", h.issue)
	mux.HandleFunc("GET /policies/{id}", h.getByID)
}

type issueRequest struct {
	Number    string `json:"number"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type policyResponse struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func (h *Handler) issue(w http.ResponseWriter, r *http.Request) {
	var req issueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		platform.WriteError(w, http.StatusBadRequest, "json inválido")
		return
	}
	start, errStart := time.Parse(dateLayout, req.StartDate)
	end, errEnd := time.Parse(dateLayout, req.EndDate)
	if req.Number == "" || errStart != nil || errEnd != nil {
		platform.WriteError(w, http.StatusBadRequest, "number, startDate y endDate (YYYY-MM-DD) son obligatorios")
		return
	}

	p, err := h.service.Issue(req.Number, start, end)
	switch {
	case errors.Is(err, ErrInvalidTerm):
		platform.WriteError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, ErrDuplicated):
		platform.WriteError(w, http.StatusConflict, err.Error())
	case err != nil:
		platform.WriteError(w, http.StatusInternalServerError, "error interno")
	default:
		platform.WriteJSON(w, http.StatusCreated, map[string]string{"id": p.ID})
	}
}

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) {
	p, err := h.service.GetByID(r.PathValue("id"))
	if err != nil {
		platform.WriteError(w, http.StatusNotFound, err.Error())
		return
	}
	platform.WriteJSON(w, http.StatusOK, policyResponse{
		ID:        p.ID,
		Number:    p.Number,
		StartDate: p.StartDate.Format(dateLayout),
		EndDate:   p.EndDate.Format(dateLayout),
	})
}

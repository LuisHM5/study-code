package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"arquitecturas/02-layered/internal/services"
)

const dateLayout = "2006-01-02"

// PolicyHandler solo traduce HTTP <-> servicio. Sin lógica de negocio.
type PolicyHandler struct {
	service *services.PolicyService
}

func NewPolicyHandler(service *services.PolicyService) *PolicyHandler {
	return &PolicyHandler{service: service}
}

func (h *PolicyHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /policies", h.Issue)
	mux.HandleFunc("GET /policies/{id}", h.GetByID)
}

type issuePolicyRequest struct {
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

func (h *PolicyHandler) Issue(w http.ResponseWriter, r *http.Request) {
	var req issuePolicyRequest
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

	p, err := h.service.Issue(req.Number, start, end)
	switch {
	case errors.Is(err, services.ErrInvalidTerm):
		writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, services.ErrDuplicated):
		writeError(w, http.StatusConflict, err.Error())
	case err != nil:
		writeError(w, http.StatusInternalServerError, "error interno")
	default:
		writeJSON(w, http.StatusCreated, map[string]string{"id": p.ID})
	}
}

func (h *PolicyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	p, err := h.service.GetByID(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, policyResponse{
		ID:        p.ID,
		Number:    p.Number,
		StartDate: p.StartDate.Format(dateLayout),
		EndDate:   p.EndDate.Format(dateLayout),
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

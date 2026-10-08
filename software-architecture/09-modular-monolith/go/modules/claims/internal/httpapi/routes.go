package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"arquitecturas/09-modular-monolith/modules/claims/internal/app"
	"arquitecturas/09-modular-monolith/platform/web"
)

func Register(mux *http.ServeMux, svc *app.Service) {
	mux.HandleFunc("POST /claims", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			PolicyID    string `json:"policyId"`
			Description string `json:"description"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			web.WriteError(w, http.StatusBadRequest, "json inválido")
			return
		}
		if req.PolicyID == "" || req.Description == "" {
			web.WriteError(w, http.StatusBadRequest, "policyId y description son obligatorios")
			return
		}

		id, err := svc.Report(r.Context(), req.PolicyID, req.Description)
		switch {
		case errors.Is(err, app.ErrPolicyNotFound):
			web.WriteError(w, http.StatusUnprocessableEntity, err.Error())
		case err != nil:
			web.WriteError(w, http.StatusInternalServerError, "error interno")
		default:
			web.WriteJSON(w, http.StatusCreated, map[string]string{"id": id})
		}
	})
}

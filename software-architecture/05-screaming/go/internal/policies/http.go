package policies

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"arquitecturas/05-screaming/internal/platform/httpserver"
)

const dateLayout = "2006-01-02"

func Routes(mux *http.ServeMux, desk *Desk) {
	mux.HandleFunc("POST /policies", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Number    string `json:"number"`
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpserver.WriteError(w, http.StatusBadRequest, "json inválido")
			return
		}
		start, errStart := time.Parse(dateLayout, req.StartDate)
		end, errEnd := time.Parse(dateLayout, req.EndDate)
		if req.Number == "" || errStart != nil || errEnd != nil {
			httpserver.WriteError(w, http.StatusBadRequest, "number, startDate y endDate (YYYY-MM-DD) son obligatorios")
			return
		}

		p, err := desk.Issue(IssueInput{Number: req.Number, StartDate: start, EndDate: end})
		switch {
		case errors.Is(err, ErrInvalidTerm):
			httpserver.WriteError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, ErrDuplicated):
			httpserver.WriteError(w, http.StatusConflict, err.Error())
		case err != nil:
			httpserver.WriteError(w, http.StatusInternalServerError, "error interno")
		default:
			httpserver.WriteJSON(w, http.StatusCreated, map[string]string{"id": p.ID})
		}
	})

	mux.HandleFunc("GET /policies/{id}", func(w http.ResponseWriter, r *http.Request) {
		p, err := desk.Find(r.PathValue("id"))
		if err != nil {
			httpserver.WriteError(w, http.StatusNotFound, err.Error())
			return
		}
		httpserver.WriteJSON(w, http.StatusOK, map[string]string{
			"id":        p.ID,
			"number":    p.Number,
			"startDate": p.StartDate.Format(dateLayout),
			"endDate":   p.EndDate.Format(dateLayout),
		})
	})
}

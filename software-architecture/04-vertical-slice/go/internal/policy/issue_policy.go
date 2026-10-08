package policy

import (
	"encoding/json"
	"net/http"
	"time"
)

// Slice "emitir póliza": request, validación, lógica y persistencia juntos.
// Recibe solo las dependencias que usa.

type issuePolicyRequest struct {
	Number    string `json:"number"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func IssuePolicy(store *Store, notifier Notifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		if err := ValidateTerm(start, end); err != nil {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		store.mu.Lock()
		for _, p := range store.byID {
			if p.Number == req.Number {
				store.mu.Unlock()
				writeError(w, http.StatusConflict, "el número de póliza ya existe")
				return
			}
		}
		p := Policy{ID: newID(), Number: req.Number, StartDate: start, EndDate: end}
		store.byID[p.ID] = p
		store.mu.Unlock()

		notifier.PolicyIssued(p.Number)
		writeJSON(w, http.StatusCreated, map[string]string{"id": p.ID})
	}
}

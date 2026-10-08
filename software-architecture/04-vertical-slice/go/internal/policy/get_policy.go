package policy

import "net/http"

// Slice "consultar póliza": es simple, así que lee directo del store.

type getPolicyResponse struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

func GetPolicy(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.mu.Lock()
		p, ok := store.byID[r.PathValue("id")]
		store.mu.Unlock()
		if !ok {
			writeError(w, http.StatusNotFound, "póliza no encontrada")
			return
		}
		writeJSON(w, http.StatusOK, getPolicyResponse{
			ID:        p.ID,
			Number:    p.Number,
			StartDate: p.StartDate.Format(dateLayout),
			EndDate:   p.EndDate.Format(dateLayout),
		})
	}
}

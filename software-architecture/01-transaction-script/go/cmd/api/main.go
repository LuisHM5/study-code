// Transaction Script: cada handler hace todo, de arriba a abajo.
package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const dateLayout = "2006-01-02"

type policy struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// "Base de datos" en memoria.
var (
	mu       sync.Mutex
	policies = map[string]policy{}
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /policies", func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			Number    string `json:"number"`
			StartDate string `json:"startDate"`
			EndDate   string `json:"endDate"`
		}
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			writeError(w, http.StatusBadRequest, "json inválido")
			return
		}
		start, errStart := time.Parse(dateLayout, in.StartDate)
		end, errEnd := time.Parse(dateLayout, in.EndDate)
		if in.Number == "" || errStart != nil || errEnd != nil {
			writeError(w, http.StatusBadRequest, "number, startDate y endDate (YYYY-MM-DD) son obligatorios")
			return
		}
		if !end.After(start) {
			writeError(w, http.StatusBadRequest, "vigencia inválida")
			return
		}

		mu.Lock()
		defer mu.Unlock()
		for _, p := range policies {
			if p.Number == in.Number {
				writeError(w, http.StatusConflict, "el número de póliza ya existe")
				return
			}
		}

		p := policy{ID: newID(), Number: in.Number, StartDate: in.StartDate, EndDate: in.EndDate}
		policies[p.ID] = p
		fmt.Printf("[notificación] póliza emitida %s\n", p.Number)
		writeJSON(w, http.StatusCreated, map[string]string{"id": p.ID})
	})

	mux.HandleFunc("GET /policies/{id}", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		p, ok := policies[r.PathValue("id")]
		mu.Unlock()
		if !ok {
			writeError(w, http.StatusNotFound, "póliza no encontrada")
			return
		}
		writeJSON(w, http.StatusOK, p)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func newID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"message": message})
}

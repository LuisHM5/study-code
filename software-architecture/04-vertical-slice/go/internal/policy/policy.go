// Package policy: un archivo por caso de uso (slice).
// Este archivo tiene solo lo que los slices comparten.
package policy

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

const dateLayout = "2006-01-02"

var ErrInvalidTerm = errors.New("vigencia inválida")

type Policy struct {
	ID        string
	Number    string
	StartDate time.Time
	EndDate   time.Time
}

// ValidateTerm es una regla pura compartida entre slices.
func ValidateTerm(start, end time.Time) error {
	if !end.After(start) {
		return ErrInvalidTerm
	}
	return nil
}

type Notifier interface {
	PolicyIssued(number string)
}

// Store es la "base de datos" en memoria que usan los slices directamente.
type Store struct {
	mu   sync.Mutex
	byID map[string]Policy
}

func NewStore() *Store {
	return &Store{byID: map[string]Policy{}}
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

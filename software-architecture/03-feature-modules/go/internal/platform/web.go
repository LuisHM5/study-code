// Package platform reúne lo técnico compartido por todas las features.
package platform

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

func NewID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func WriteJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"message": message})
}

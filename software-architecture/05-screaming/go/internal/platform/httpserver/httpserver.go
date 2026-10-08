// Package httpserver: lo técnico vive en platform, explícito y fuera del negocio.
package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func WriteJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"message": message})
}

func Run(mux *http.ServeMux) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

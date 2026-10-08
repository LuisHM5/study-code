package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"os"

	"arquitecturas/06-hexagonal/internal/policy/adapters/console"
	"arquitecturas/06-hexagonal/internal/policy/adapters/httpapi"
	"arquitecturas/06-hexagonal/internal/policy/adapters/memory"
	"arquitecturas/06-hexagonal/internal/policy/app"
)

func main() {
	// Aquí se conecta cada puerto con su adaptador.
	repo := memory.NewRepository()
	handler := httpapi.Handler{
		IssuePolicy: app.IssuePolicy{Policies: repo, Notifier: console.Notifier{}, NewID: newID},
		GetPolicy:   app.GetPolicy{Policies: repo},
	}

	mux := http.NewServeMux()
	handler.Register(mux)

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

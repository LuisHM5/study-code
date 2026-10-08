package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"os"

	"arquitecturas/07-onion/internal/policy/application"
	"arquitecturas/07-onion/internal/policy/infrastructure/console"
	"arquitecturas/07-onion/internal/policy/infrastructure/memory"
	"arquitecturas/07-onion/internal/policy/presentation/httpapi"
)

func main() {
	repo := memory.NewRepository()
	handler := httpapi.Handler{
		IssuePolicy: application.NewIssuePolicy(repo, console.Notifier{}, newID),
		GetPolicy:   application.NewGetPolicy(repo),
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

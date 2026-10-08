package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"os"

	"arquitecturas/10-cqrs/internal/policies/adapters"
	"arquitecturas/10-cqrs/internal/policies/app"
	"arquitecturas/10-cqrs/internal/policies/app/command"
	"arquitecturas/10-cqrs/internal/policies/app/query"
)

func main() {
	db := adapters.NewMemoryDB()
	readModel := adapters.NewMemoryReadModel(db)

	application := app.Application{
		Commands: app.Commands{
			IssuePolicy: command.IssuePolicyHandler{
				Policies: adapters.NewMemoryRepository(db),
				Notifier: adapters.ConsoleNotifier{},
				NewID:    newID,
			},
		},
		Queries: app.Queries{
			GetPolicy:    query.GetPolicyHandler{ReadModel: readModel},
			ListExpiring: query.ListExpiringHandler{ReadModel: readModel},
		},
	}

	mux := http.NewServeMux()
	adapters.NewServer(application).Register(mux)

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

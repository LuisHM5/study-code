package main

import (
	"log"
	"net/http"
	"os"

	"arquitecturas/03-feature-modules/internal/notification"
	"arquitecturas/03-feature-modules/internal/policy"
)

func main() {
	// El "módulo" es simplemente el cableado de cada feature.
	notifier := notification.NewService()
	policyHandler := policy.NewHandler(policy.NewService(policy.NewRepository(), notifier))

	mux := http.NewServeMux()
	policyHandler.Register(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

package main

import (
	"log"
	"net/http"
	"os"

	"arquitecturas/02-layered/internal/handlers"
	"arquitecturas/02-layered/internal/repositories"
	"arquitecturas/02-layered/internal/services"
)

func main() {
	// Handler -> Service -> Repository
	repo := repositories.NewPolicyRepository()
	service := services.NewPolicyService(repo)
	handler := handlers.NewPolicyHandler(service)

	mux := http.NewServeMux()
	handler.Register(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

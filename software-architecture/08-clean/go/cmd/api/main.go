package main

import (
	"log"
	"net/http"
	"os"

	"arquitecturas/08-clean/internal/policies/infra"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, infra.NewRouter()))
}

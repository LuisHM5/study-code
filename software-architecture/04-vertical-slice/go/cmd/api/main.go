package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"arquitecturas/04-vertical-slice/internal/policy"
)

type consoleNotifier struct{}

func (consoleNotifier) PolicyIssued(number string) {
	fmt.Printf("[notificación] póliza emitida %s\n", number)
}

func main() {
	mux := http.NewServeMux()
	policy.Register(mux, policy.NewStore(), consoleNotifier{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

package main

import (
	"log"
	"net/http"
	"os"

	"arquitecturas/11-event-driven-event-sourcing/internal/notifications"
	"arquitecturas/11-event-driven-event-sourcing/internal/platform/eventbus"
	"arquitecturas/11-event-driven-event-sourcing/internal/policies"
)

func main() {
	bus := eventbus.New()
	notifications.Subscribe(bus)

	svc := policies.NewService(policies.NewEventStore(), bus)

	mux := http.NewServeMux()
	policies.Register(mux, svc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

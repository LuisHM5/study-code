package main

import (
	"log"
	"net/http"
	"os"

	"arquitecturas/09-modular-monolith/modules/claims"
	"arquitecturas/09-modular-monolith/modules/notifications"
	"arquitecturas/09-modular-monolith/modules/policies"
	"arquitecturas/09-modular-monolith/platform/eventbus"
)

func main() {
	bus := eventbus.New()

	// Un solo proceso; cada módulo recibe solo lo que necesita de los demás.
	policiesModule := policies.New(bus)
	claimsModule := claims.New(policiesModule.API)
	notifications.New(bus)

	mux := http.NewServeMux()
	policiesModule.Routes(mux)
	claimsModule.Routes(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

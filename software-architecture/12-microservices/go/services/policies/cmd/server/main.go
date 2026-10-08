// Servicio de pólizas: proceso y despliegue propios.
package main

import (
	"log"
	"net/http"
	"os"

	"arquitecturas/12-microservices/services/policies/internal/app"
	"arquitecturas/12-microservices/services/policies/internal/httpapi"
	"arquitecturas/12-microservices/services/policies/internal/memory"
	"arquitecturas/12-microservices/services/policies/internal/notificationsclient"
)

func main() {
	notificationsURL := os.Getenv("NOTIFICATIONS_URL")
	if notificationsURL == "" {
		notificationsURL = "http://localhost:3002"
	}

	svc := app.NewService(memory.NewRepository(), notificationsclient.New(notificationsURL))

	mux := http.NewServeMux()
	httpapi.Register(mux, svc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	log.Printf("policies escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

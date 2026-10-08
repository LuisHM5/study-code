// Servicio de notificaciones: proceso y despliegue propios.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"arquitecturas/12-microservices/contracts"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /notifications", func(w http.ResponseWriter, r *http.Request) {
		var msg contracts.PolicyIssuedNotification
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil || msg.Number == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "cuerpo inválido"})
			return
		}
		fmt.Printf("[notificación] póliza emitida %s\n", msg.Number)
		w.WriteHeader(http.StatusAccepted)
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	log.Printf("notifications escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

package policy

import "net/http"

// Register conecta cada slice con su ruta. Añadir una feature = un archivo + una línea aquí.
func Register(mux *http.ServeMux, store *Store, notifier Notifier) {
	mux.Handle("POST /policies", IssuePolicy(store, notifier))
	mux.Handle("GET /policies/{id}", GetPolicy(store))
}

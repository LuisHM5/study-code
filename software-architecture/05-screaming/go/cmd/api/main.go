package main

import (
	"net/http"

	"arquitecturas/05-screaming/internal/platform/httpserver"
	"arquitecturas/05-screaming/internal/platform/mailer"
	"arquitecturas/05-screaming/internal/policies"
)

func main() {
	desk := policies.NewDesk(policies.NewStore(), mailer.Console{})

	mux := http.NewServeMux()
	policies.Routes(mux, desk)
	httpserver.Run(mux)
}

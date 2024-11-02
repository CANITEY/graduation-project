package main

import "net/http"

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /events", a.stream)
	mux.HandleFunc("GET /ping", a.ping)
	mux.HandleFunc("POST /sos", a.sos)


	return mux
}

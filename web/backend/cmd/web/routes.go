package main

import "net/http"

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user/login", a.login)
	mux.HandleFunc("POST /user/challenge", a.challenge)
	mux.HandleFunc("GET /events", a.stream)
	mux.HandleFunc("GET /ping", a.ping)
	mux.HandleFunc("POST /cars/sos", a.sos)


	return mux
}

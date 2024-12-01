package main

import "net/http"

func (a *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /challenge/{challUUID}", a.challGetter)
	mux.HandleFunc("GET /events", a.stream)
	mux.HandleFunc("GET /ping", a.ping)
	mux.HandleFunc("GET /view/{carUUID}", a.viewCar)
	mux.HandleFunc("POST /user/login", a.login)
	mux.HandleFunc("POST /user/challenge", a.challenge) // this endpoint is used after the challenge was solved with the validator, to tell the front end to redirect to dashboard
	mux.HandleFunc("POST /challenge/{challUUID}", a.challengeValidator)
	mux.HandleFunc("POST /cars/sos", a.sos)


	return mux
}

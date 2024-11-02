package main

import (
	"database/sql"
	"net/http"

	"github.com/r3labs/sse/v2"
	_ "github.com/lib/pq"
)

type application struct {
	sse *sse.Server
	db *sql.DB
}

func main() {
	app := NewApplication("")

	server := http.Server{
		Handler: app.routes(),
		Addr: ":5000",
	}

	server.ListenAndServe()

}

func NewApplication(sseParameter string) *application {

	if sseParameter == "" {
		sseParameter = "messages"
	}

	// Creating the SSE server
	server := sse.New()
	server.CreateStream(sseParameter)

	return &application{
		sse: server,

	}
}

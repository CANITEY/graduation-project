package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/r3labs/sse/v2"
)

type application struct {
	sse *sse.Server
	db *sql.DB
}

func main() {
	app, err := NewApplication("")
	if err != nil {
		log.Fatalln(err.Error())
	}

	server := http.Server{
		Handler: app.routes(),
		Addr: ":5000",
	}
	log.Println("SERVER STARTED")

	server.ListenAndServe()

}


func OpenDB() (*sql.DB, error) {
	// TODO: connect through ssl (enable sslmode)
	db, err := sql.Open("postgres", "host=localhost user=emergency password=emergency dbname=cars sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	log.Println("CONNECTED TO DATABASE CARS")

	return db, nil
}


func NewApplication(sseParameter string) (*application, error) {
	if sseParameter == "" {
		sseParameter = "messages"
	}

	// Creating the SSE server
	server := sse.New()
	server.CreateStream(sseParameter)

	db, err := OpenDB()
	if err != nil {
		return nil, err
	}

	return &application{
		sse: server,
		db: db,
	}, nil
}

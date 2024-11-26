package main

import (
	"database/sql"
	"fmt"
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
	fmt.Println("Server started")

	server.ListenAndServe()

}


func OpenDB() (*sql.DB, error) {
	// BUG: 2024/11/04 00:46:14 pq: SSL is not enabled on the server
	db, err := sql.Open("postgres", "postgres://emergency:emergency@localhost/cars@sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	

	return db, nil
}


func NewApplication(sseParameter string) (*application, error) {
	if sseParameter == "" {
		sseParameter = "messages"
	}

	// Creating the SSE server
	server := sse.New()
	server.CreateStream(sseParameter)

	// BUG: Commented until I fix bug
	// open database connection
	// db, err := OpenDB()
	// if err != nil {
	// 	return nil, err
	// }

	return &application{
		sse: server,
		// db: db,
	}, nil
}

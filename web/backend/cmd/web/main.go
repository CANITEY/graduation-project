package main

import (
	"database/sql"
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	"gp-backend/database"
=======
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
	"log"
	"net/http"
	"time"

<<<<<<< Updated upstream
<<<<<<< Updated upstream
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
=======
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
	_ "github.com/lib/pq"
	"github.com/r3labs/sse/v2"
)

type application struct {
	sse *sse.Server
	db *sql.DB
	udb database.UserDB
	sessionManager *scs.SessionManager
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
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	// TODO: connect through ssl (enable sslmode)
	db, err := sql.Open("postgres", "host=localhost user=emergency password=emergency dbname=cars sslmode=disable")
=======
	// BUG: 2024/11/04 00:46:14 pq: SSL is not enabled on the server
	db, err := sql.Open("postgres", "postgres://emergency:emergency@localhost/cars@sslmode=disable")
>>>>>>> Stashed changes
=======
	// BUG: 2024/11/04 00:46:14 pq: SSL is not enabled on the server
	db, err := sql.Open("postgres", "postgres://emergency:emergency@localhost/cars@sslmode=disable")
>>>>>>> Stashed changes
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	log.Println("CONNECTED TO DATABASE CARS")
=======
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes

	return db, nil
}


func NewApplication(sseParameter string) (*application, error) {
	if sseParameter == "" {
		sseParameter = "messages"
	}

	// Creating the SSE server
	server := sse.New()
	server.CreateStream(sseParameter)

<<<<<<< Updated upstream
<<<<<<< Updated upstream
=======
	// open database connection
>>>>>>> Stashed changes
=======
	// open database connection
>>>>>>> Stashed changes
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	udb := database.UserDB{
		DB: db,
	}

	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
=======
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes

	return &application{
		sse: server,
		db: db,
<<<<<<< Updated upstream
<<<<<<< Updated upstream
		udb: udb,
		sessionManager: sessionManager,
=======
>>>>>>> Stashed changes
=======
>>>>>>> Stashed changes
	}, nil
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/r3labs/sse/v2"
)

func main() {
	server := sse.New()
	server.CreateStream("messages")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /events", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("access-control-allow-origin", "*")
		server.ServeHTTP(w, r)
	},
)
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		server.Publish("messages", &sse.Event{
			Data: []byte("ping"),
		})
		log.Println("TRIGGER SENT")
		fmt.Fprintln(w, server.Headers)
	})

	http.ListenAndServe(":5000", mux)
}



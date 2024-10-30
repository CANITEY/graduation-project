package main

import (
	"encoding/json"
	"fmt"
	"gp-backend/cmd/models"
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

	mux.HandleFunc("POST /sos", func(w http.ResponseWriter, r *http.Request) {
		var data models.CarInfo
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			log.Printf("Can't decode json body from [%v]", r.RemoteAddr)
			msg := models.Msg{
				Status: "FAIL",
				Message: "Can't parse json",
			}
			encoder := json.NewEncoder(w)
			encoder.Encode(msg)
			return
		}
	})


	http.ListenAndServe(":5000", mux)
}



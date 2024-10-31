package main

import (
	"encoding/json"
	"gp-backend/models"
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
	})

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		server.Publish("messages", &sse.Event{
			Data: []byte("ping"),
		})
		log.Println("TRIGGER SENT")
	})

	mux.HandleFunc("POST /sos", func(w http.ResponseWriter, r *http.Request) {
		var data models.CarInfo
		decoder := json.NewDecoder(r.Body)
		
		// handle prasing errors
		if err := decoder.Decode(&data); err != nil {
			log.Printf("Can't decode json body from [%v]", r.RemoteAddr)
			msg := models.Msg{
				Status: "fail",
				Message: "JSON object is malformed",
			}
			JSONwriter(w, msg)
			return
		}


		if data.UUID == "" {
			log.Printf("Can't decode json body from [%v]", r.RemoteAddr)
			msg := models.Msg{
				Status: "fail",
				Message: "provide UUID",
			}
			JSONwriter(w, msg)
			return
		} else if data.DriverStatus == "" {
			log.Printf("Can't decode json body from [%v]", r.RemoteAddr)
			msg := models.Msg{
				Status: "fail",
				Message: "provide driver_status",
			}
			JSONwriter(w, msg)
			return
		}

		msg := models.Msg{
			Status: "success",
			Message: "success",
		}
		JSONwriter(w, msg)

		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Couldn't generate cars' data")
		}
		

		server.Publish("messages", &sse.Event{
			Data: jsonData,
		})
	})


	http.ListenAndServe(":5000", mux)
}



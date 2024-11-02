package main

import (
	"encoding/json"
	"gp-backend/models"
	"log"
	"net/http"

	"github.com/r3labs/sse/v2"
)


func (a *application) stream(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("access-control-allow-origin", "*")
		a.sse.ServeHTTP(w, r)
}

func (a *application) ping(w http.ResponseWriter, r *http.Request) {
		a.sse.Publish("messages", &sse.Event{
			Data: []byte("ping"),
		})
		log.Println("TRIGGER SENT")
}

func (a *application) sos(w http.ResponseWriter, r *http.Request) {
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
			return
		}
		

		server.Publish("messages", &sse.Event{
			Data: jsonData,
		})
}

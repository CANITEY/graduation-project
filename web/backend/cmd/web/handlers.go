package main

import (
	"encoding/json"
	"gp-backend/models"
	"log"
	"net/http"

	"github.com/r3labs/sse/v2"
)

// TODO: Create a login function
func (a *application) login(w http.ResponseWriter, r *http.Request) {
	// NOTE:: SHOULD REDIRECT TO THE SECOND STAGE OF THE LOGIN, WHICH WILL BE CARD READER OR AN OTP
	// THEN AUTHORIZE THEM BOTH TO MAKE THE LOGIN

}

func (a *application) stream(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("access-control-allow-origin", "*")
		a.sse.ServeHTTP(w, r)
}

func (a *application) ping(w http.ResponseWriter, r *http.Request) {
		a.sse.Publish("messages", &sse.Event{
			Data: []byte("ping"),
		})
		writeJSON(w, http.StatusOK, models.Msg{
			Status: "success",
			Message: "success",
		}, nil)
		log.Println("TRIGGER SENT")
}

// TODO: link it with the readJSON helper and wait for the validator to validate the data
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
			writeJSON(w, http.StatusBadRequest, msg, nil)
			return
		}

		if data.UUID == "" {
			log.Printf("Can't decode json body from [%v]", r.RemoteAddr)
			msg := models.Msg{
				Status: "fail",
				Message: "provide UUID",
			}
			writeJSON(w, http.StatusBadRequest, msg, nil)
			return
		} else if data.DriverStatus == "" {
			log.Printf("Can't decode json body from [%v]", r.RemoteAddr)
			msg := models.Msg{
				Status: "fail",
				Message: "provide driver_status",
			}
			writeJSON(w, http.StatusBadRequest, msg, nil)
			return
		}

		msg := models.Msg{
			Status: "success",
			Message: "success",
		}
		writeJSON(w, http.StatusOK, msg, nil)

	jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Couldn't generate cars' data")
			return
		}
		
		a.sse.Publish("messages", &sse.Event{
			Data: jsonData,
		})
}

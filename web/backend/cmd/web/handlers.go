package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gp-backend/models"
	"gp-backend/validate"
	"log"
	"net/http"

	"github.com/r3labs/sse/v2"
)

var (
	password validate.Key = "password"
	token validate.Key = "token"
	UUID validate.Key = "UUID"
	driverStatus validate.Key = "driver_status"
	longitude validate.Key = "longitude"
	latitude validate.Key = "latitude"
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

	// error handling for json reading, json sent is read into data variable
	err := readJSON(w, r, data)
	if err != nil {
		clientError(w, http.StatusBadRequest, err)
		return
	}

	// validating json input
	validator := validate.New()
	// UUID validation
	// Checking not empty
	validator.Check(validate.NotEmpty(data.UUID), UUID, "must not be empty")
	// Checking the UUID syntax
	// TODO:

	// driver_status validation
	validator.Check(validate.NotEmpty(data.DriverStatus), driverStatus, "must not be empty")
	validator.Check(validate.In(data.DriverStatus, []string{"sleeping", "fainted", "awake"}), driverStatus, "is a value outside the intended list")

	// longitude & latitude validation
	validator.Check(data.Longitude == 0, longitude, "must not be zero")
	validator.Check(data.Latitude == 0, latitude, "must not be zero")


	if !validator.Valid() {
		buf := bytes.NewBufferString("")
		for key, value := range validator.Errors {
			format := fmt.Sprintf("%v: %v\n", key, value)
			buf.WriteString(format)
		}

		err := errors.New(buf.String())
		clientError(w, http.StatusBadRequest, err)
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

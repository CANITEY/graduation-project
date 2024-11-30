package main

import (
	"gp-backend/database"
	"net/http"
)

func (a *application) challGetter(w http.ResponseWriter, r *http.Request) {
	challUUID := r.PathValue("challUUID")
	challenge, err := database.GetChallengeStr(a.db, challUUID)
	if err != nil {
		serverError(w, err)
	}

	jsonMsg := struct{
		Status string `json:"status"`
		Challenge string `json:"challenge"`
	}{
		Status: "success", 
		Challenge: challenge,
	}

	if err := writeJSON(w, http.StatusOK, jsonMsg, nil); err != nil {
		serverError(w, err)
	}
}

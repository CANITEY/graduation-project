package main

import (
	"fmt"
	"gp-backend/crypto/validator"
	"gp-backend/database"
	"net/http"

	"github.com/google/uuid"
)

func (a *application) challGetter(w http.ResponseWriter, r *http.Request) {
	challUUID := r.PathValue("challUUID")
	if err := uuid.Validate(challUUID); err != nil {
		clientError(w, http.StatusBadRequest, err)
		return
	}
	challenge, err := database.GetChallengeStr(a.db, challUUID)
	if err != nil {
		serverError(w, err)
		return
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

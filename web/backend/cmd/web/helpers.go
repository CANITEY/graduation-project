package main

import (
	"encoding/json"
	"net/http"
)

func JSONwriter(w http.ResponseWriter, data any) {
			encoder := json.NewEncoder(w)
			encoder.Encode(data)
} 

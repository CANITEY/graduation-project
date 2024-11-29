package main

import (
	"fmt"
	"gp-backend/crypto/validator"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

// TODO: after getting the challenge follow GPT steps to hash the challenge then solve it, then send the hash again
// add 2 endpoints to give the challenge and to validate the challenge
func main() {
	var carUUID string
	_, err := fmt.Scanf("%v", &carUUID)
	if err != nil {
		log.Fatalln(err)
	}
	path, err := getUSBPath("/run/media/mohammed/*")
	if err != nil {
		log.Fatalln(err)
	}

	// loading private key from drive
	if _, err := validator.LoadPrivKey(path); err != nil {
		log.Fatalln(err)
	}
}


func getUSBPath(glob string) (string, error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return "", err
	}

	return matches[0], nil
}

func getChallenge(carUUID string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:5000/challenge/%v", carUUID))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("ERROR: STATUS RESPONSE: ", resp.StatusCode)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(response), nil
}

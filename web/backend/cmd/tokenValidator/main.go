package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gp-backend/crypto/validator"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	var challUUID string
	_, err := fmt.Scanf("%v", &challUUID)
	if err != nil {
		log.Fatalln(err)
	}
	path, err := getUSBPath("/run/media/mohammed/*")
	if err != nil {
		log.Fatalln(err)
	}

	// loading private key from drive
	pri8Key, err := validator.LoadPrivKey(path)
	if err != nil {
		log.Fatalln(err)
	}

	challenge, err := getChallenge(challUUID)
	if err != nil {
		log.Fatalln(err)
	}

	hash := validator.GenerateChallengeHash(challenge)
	solution, err := validator.SolveChallenge(pri8Key, hash)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := submitSolution(challUUID, solution)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
}


func getUSBPath(glob string) (string, error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return "", err
	}

	return matches[0], nil
}


func getChallenge(challUUID string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:5000/challenge/%v", challUUID))
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


func submitSolution(challUUID string, solution map[string]string) (string, error) {
	jsonData, err := json.Marshal(solution)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(fmt.Sprintf("http://127.0.0.1:5000/challenge/%v", challUUID), "application/json", bytes.NewBuffer(jsonData)) 

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ERROR: WRONG SOLUTION, WRONG CHALLENGE")
	}


	return "", nil
}

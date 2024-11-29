package main

import (
	"fmt"
	"gp-backend/crypto/validator"
	"log"
	"os"
	"path/filepath"
)

func main() {
	_, err := getArgs()
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

func getArgs() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("No suffecient args are providied\nusage tokenValidator <challenge code>")
	}

	return os.Args, nil
}

func getUSBPath(glob string) (string, error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return "", err
	}

	return matches[0], nil
}


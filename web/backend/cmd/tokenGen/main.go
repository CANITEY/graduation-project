package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/crypto/pbkdf2"
)

// TODO: change config to hold username instead of Salt
// make a database table holding salt, password hash, and cipher text
type Config struct {
	PasswordHash string `json:"password_hash"`
	Salt		 string `json:"salt"`
}

func main() {
	if !isRoot() {
		panic("ERROR:MUST RUN AS ROOT")
	}

	args, err := getArgs()
	if err != nil {
		panic(err)
	}


	password := getPassword(args)
	masterKey, salt, err := generateMasterKey(password)
	if err != nil {
		panic(err)
	}

	cipher, err := generateCipher(masterKey, salt)
	if err != nil {
		panic(err)
	}

	config := Config{
		PasswordHash: hashPassword(password),
		Salt: hex.EncodeToString(salt),
	}

	configData, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	// creating a dir to mount the flash drive defaulting to /tmp/usb
	if err := createMntDir(); err != nil {
		panic(err)
	}


	// mounting usb drive to write config info into it
	// getting the usb device path 
	usbPath, err := getUSBPath(args)
	if err != nil {
		panic(err)
	}

	if err := createMntDir(); err != nil {
		panic(err)
	}

	// mounting usb device
	if err := MntDrive(usbPath, "/tmp/usb"); err != nil {
		panic(err)
	}
	// defer the unmount to unmout the usb drive
	defer exec.Command("umount", "/tmp/usb")


	// writing config file into drive
	err = os.WriteFile("/tmp/usb/config.json", configData, 0644)
	if err != nil {
		panic(err)
	}

	// write encrypted data into usb drive
	err = os.WriteFile("/tmp/usb/validation.enc", cipher, 0644)
	if err != nil {
		panic(err)
	}

	println("Initialization completed successfully")
	// TODO: ADD A SAVE TO DATABASE FUNCTIONALITY BY USER CHOICE
}

func getArgs() ([]string, error) {
	if len(os.Args) < 3 {
		return nil, fmt.Errorf("No suffecient args are providied\nusage tokenGen <drive path> <user password>")
	}

	return os.Args, nil
}

func getUSBPath(args []string) (string, error) {
	matches, err := filepath.Glob(args[1])
	if err != nil {
		return "", err
	}

	return matches[0], nil
	
}

func getPassword(args []string) (string) {
	return args[2]
}

// Generate and test master key for token generation
func generateMasterKey(password string) ([]byte, []byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, nil, err
	}

	masterKey := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)


	return masterKey, salt, nil
}

func generateCipher(masterKey, salt []byte) ([]byte, error) {
	block, err := aes.NewCipher(masterKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	message := make([]byte, 1024)
	return gcm.Seal(nil, nonce, message, nil), nil
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func isRoot() (bool) {
	return os.Getuid() == 0
}

func createMntDir() (error) {
	 return os.MkdirAll("/tmp/usb", 0755)
}

func MntDrive(usbPath, mountPath string) (error) {
	cmd := exec.Command("mount", usbPath, mountPath)
	return cmd.Run()
}

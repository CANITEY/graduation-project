package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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


	// generate private key
	privateKey, err := generatePrivateKey()
	if err != nil {
		panic(err)
	}
	// write private key to memory (usb flashdrive)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	err = os.WriteFile("/tmp/usb/private_key.pem", privateKeyPEM, 0600)
	if err != nil {
		panic(err)
	}

	// Generate public key for authentication
	publicKey := &privateKey.PublicKey
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})
	fmt.Println(string(publicKeyPEM))

	fmt.Println("Initialization completed successfully")
	fmt.Println("Save token to database? [Y/n]")
	
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

// Generate RSA key pair
func generatePrivateKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
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

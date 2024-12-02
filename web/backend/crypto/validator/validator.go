package validator

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func LoadPrivKey(path string) (*rsa.PrivateKey, error) {
	privateKeyPEM, err := os.ReadFile(fmt.Sprintf("%s/private_key.pem", path))
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode private key block")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func GenerateChallengeHash(challenge []byte) ([]byte) {
	hashedChallenge := sha256.Sum256(challenge)

	return hashedChallenge[:]
}

func SolveChallenge(pri8Key *rsa.PrivateKey, challenge []byte) (map[string]string, error) {
	signature, err := rsa.SignPKCS1v15(nil, pri8Key, crypto.SHA256, challenge)
	if err != nil {
		return nil, err
	}
	
	signatureBase64Encoded := base64.StdEncoding.EncodeToString(signature)

	data := map[string]string{
		"signature": signatureBase64Encoded,
	}

	return data, nil
}

func ResolveChallenge(pubKey *rsa.PublicKey, challenge []byte, solution string) (bool, error) {
	hashedChallenge := sha256.Sum256(challenge)
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashedChallenge[:], []byte(solution))
	if err != nil {
		return false, err
	}

	return true, nil
}

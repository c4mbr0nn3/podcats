package utils

import (
	"math/rand"
	"os"
)

// get jwt secret from environment variable or generate random secret
func GetJwtSecret() []byte {
	secret := readSecretFromEnv()
	if len(secret) == 0 {
		secret = randBytes(32)
	}
	return secret
}

// read jwt secret from environment variable
func readSecretFromEnv() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// generate random hex string for jwt secret with 32 bytes length (256 bits)
func randBytes(n int) []byte {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return b
}

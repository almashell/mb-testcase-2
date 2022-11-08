package security

import (
	"crypto/rand"
	"mb_api/internal/pkg/settings"
	"net/http"
	"time"
)

// Simple function that generate random sequence of bytes for
// different aims such as security, fun, fake data
func GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// generateCSRFToken returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateCSRFToken(length int) (string, error) {
	bytes, err := GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}

	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

// NewCSRFToken set to http.ResponseWriter cookie "csrf" which contains
// just generated csrf-token using generateCSRFToken function. NewCSRFToken
// use constant CSRFTokenLen from settings structure SecureSettings which
// are set in the configs/server/settings.go
func NewCSRFToken(w http.ResponseWriter) error {
	token, err := generateCSRFToken(settings.SecureSettings.CSRFTokenLen)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "csrf",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * settings.SecureSettings.CSRFTokenTTL),
		HttpOnly: false,
		Path:     "/", // TODO: check it
	})

	return nil
}
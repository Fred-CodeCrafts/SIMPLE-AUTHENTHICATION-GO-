// utils/random.go
package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomPassword() (string, error) {
	b := make([]byte, 9) // 12 base64 chars approx
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

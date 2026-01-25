package utils

import (
	"crypto/rand"
	"encoding/hex"
)

const uidLength = 3 // 6 hex characters

func GenerateUID() (string, error) {
	b := make([]byte, uidLength)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

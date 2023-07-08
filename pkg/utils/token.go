package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken(size int) string {
	b := make([]byte, size)

	if _, err := rand.Read(b); err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}

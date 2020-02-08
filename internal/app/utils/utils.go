package utils

import (
	"crypto/rand"
	"fmt"
)

// GenerateKeySession ...
func GenerateKeySession() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

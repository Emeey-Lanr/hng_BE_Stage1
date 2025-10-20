package services

import (
	"crypto/sha256"
	"fmt"
)
func CreateSHA256Hash (value string) string {
    hash := sha256.Sum256([]byte(value))

	
	return fmt.Sprintf("%x", hash)
}
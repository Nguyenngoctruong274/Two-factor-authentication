package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashPassword(password string) string {
	textSha256 := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(textSha256[:])
}

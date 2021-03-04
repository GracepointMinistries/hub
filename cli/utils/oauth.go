package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"strings"
	"time"
)

// EncodeBase64 encodes the message into a web-safe base64 message
func EncodeBase64(msg []byte) string {
	encoded := base64.StdEncoding.EncodeToString(msg)
	encoded = strings.Replace(encoded, "+", "-", -1)
	encoded = strings.Replace(encoded, "/", "_", -1)
	encoded = strings.Replace(encoded, "=", "", -1)
	return encoded
}

// GenerateOAuthChallenge generates a challenge of the given length and returns it as a string
func GenerateOAuthChallenge(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length, length)
	for i := 0; i < length; i++ {
		b[i] = byte(r.Intn(255))
	}
	return EncodeBase64(b)
}

// HashS256 calculates the SHA256 of the message and encodes it in web-safe base64
func HashS256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return EncodeBase64(h.Sum(nil))
}

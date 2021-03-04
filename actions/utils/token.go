package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	jose "gopkg.in/square/go-jose.v2"
)

var (
	encrypter     jose.Encrypter
	encryptionKey []byte
)

func init() {
	encryptionKey = []byte(getSecret(string(GetEnvironment())))
	joseEncrypter, err := jose.NewEncrypter(
		jose.A128GCM,
		jose.Recipient{
			Algorithm: jose.A128GCMKW,
			Key:       encryptionKey,
		},
		nil,
	)
	if err != nil {
		panic("Unable to initialize jwt encryption")
	}
	encrypter = joseEncrypter
}

func webBase64Encode(msg []byte) string {
	encoded := base64.StdEncoding.EncodeToString(msg)
	encoded = strings.Replace(encoded, "+", "-", -1)
	encoded = strings.Replace(encoded, "/", "_", -1)
	encoded = strings.Replace(encoded, "=", "", -1)
	return encoded
}

func webBase64Decode(encoded string) (string, error) {
	encoded = strings.Replace(encoded, "-", "+", -1)
	encoded = strings.Replace(encoded, "_", "/", -1)
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

type tokenPayload struct {
	Admin     string `json:"admin,omitempty"`
	ID        int    `json:"id,omitempty"`
	SessionID int    `json:"session_id"`
}

func generateToken(payload *tokenPayload) (string, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	encryption, err := encrypter.Encrypt(data)
	if err != nil {
		return "", err
	}
	return webBase64Encode([]byte(encryption.FullSerialize())), nil
}

// GenerateAdminToken generates an unscoped token for an administrator
func GenerateAdminToken(admin string, session int) (string, error) {
	return generateToken(&tokenPayload{
		Admin:     admin,
		SessionID: session,
	})
}

// GenerateUserToken generates a token scoped to a user
func GenerateUserToken(user, session int) (string, error) {
	return generateToken(&tokenPayload{
		ID:        user,
		SessionID: session,
	})
}

// GenerateScopedToken generates a token for an administrator scoped to a user
func GenerateScopedToken(admin string, user, session int) (string, error) {
	return generateToken(&tokenPayload{
		Admin:     admin,
		ID:        user,
		SessionID: session,
	})
}

func parseToken(token string) (*tokenPayload, error) {
	decodedToken, err := webBase64Decode(token)
	if err != nil {
		return nil, err
	}
	encrypted, err := jose.ParseEncrypted(decodedToken)
	if err != nil {
		return nil, err
	}
	data, err := encrypted.Decrypt(encryptionKey)
	if err != nil {
		return nil, err
	}
	payload := &tokenPayload{}
	err = json.Unmarshal(data, payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

// ParseAdminToken decodes the auth token given to it
// and ensures that it has administrator information in it
func ParseAdminToken(token string) (string, int, error) {
	parsed, err := parseToken(token)
	if err != nil {
		return "", 0, err
	}
	if parsed.Admin == "" || parsed.SessionID == 0 {
		return "", 0, errors.New("not an admin token")
	}
	return parsed.Admin, parsed.SessionID, nil
}

// ParseScopedToken decodes the auth token given to it
// and ensures that it is scoped to a user
func ParseScopedToken(token string) (string, int, int, error) {
	parsed, err := parseToken(token)
	if err != nil {
		return "", 0, 0, err
	}
	if parsed.ID <= 0 || parsed.SessionID == 0 {
		return "", 0, 0, errors.New("not a user token")
	}
	return parsed.Admin, parsed.ID, parsed.SessionID, nil
}

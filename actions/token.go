package actions

import (
	"encoding/json"
	"errors"

	jose "gopkg.in/square/go-jose.v2"
)

var (
	encrypter     jose.Encrypter
	encryptionKey []byte
)

func init() {
	encryptionKey = []byte(getSecret(string(getEnvironment())))
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
	return encryption.FullSerialize(), nil
}

func generateAdminToken(admin string, session int) (string, error) {
	return generateToken(&tokenPayload{
		Admin:     admin,
		SessionID: session,
	})
}

func generateUserToken(user, session int) (string, error) {
	return generateToken(&tokenPayload{
		ID:        user,
		SessionID: session,
	})
}

func parseToken(token string) (*tokenPayload, error) {
	encrypted, err := jose.ParseEncrypted(token)
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

func parseAdminToken(token string) (string, int, error) {
	parsed, err := parseToken(token)
	if err != nil {
		return "", 0, err
	}
	if parsed.Admin == "" || parsed.SessionID == 0 {
		return "", 0, errors.New("not an admin token")
	}
	return parsed.Admin, parsed.SessionID, nil
}

func parseUserToken(token string) (int, int, error) {
	parsed, err := parseToken(token)
	if err != nil {
		return 0, 0, err
	}
	if parsed.ID <= 0 || parsed.SessionID == 0 {
		return 0, 0, errors.New("not a user token")
	}
	return parsed.ID, parsed.SessionID, nil
}

package sync

import (
	"context"

	"github.com/gobuffalo/envy"
	googleOAuth "golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	drive "google.golang.org/api/drive/v3"
	sheets "google.golang.org/api/sheets/v4"
)

var (
	syncClient  *sheets.Service
	driveClient *drive.Service
)

// SetupClient sets up the global sheets client
func SetupClient() error {
	email, err := envy.MustGet("GOOGLE_CLIENT_EMAIL")
	if err != nil {
		return err
	}
	privateKey, err := envy.MustGet("GOOGLE_CLIENT_PRIVATE_KEY")
	if err != nil {
		return err
	}
	config := &jwt.Config{
		Email:      email,
		PrivateKey: []byte(privateKey),
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets",
			"https://www.googleapis.com/auth/drive",
		},
		TokenURL: googleOAuth.JWTTokenURL,
	}
	syncClient, err = sheets.New(config.Client(context.Background()))
	if err != nil {
		return err
	}
	driveClient, err = drive.New(config.Client(context.Background()))
	return err
}

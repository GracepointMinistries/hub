package utils

import (
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gorilla/sessions"
)

// This file contains less-dumb security options for the cookie store that
// buffalo creates

func getSecret(environment string) string {
	secret := envy.Get("SESSION_SECRET", "")
	// In production a SESSION_SECRET must be set!
	if secret == "" {
		if environment == "development" || environment == "test" {
			secret = strings.Repeat("x", 32)
		} else {
			panic("Unless you set SESSION_SECRET env variable, your session storage is not protected!")
		}
	}
	return secret
}

// SetSecureStore sets cookies as HTTPOnly and with the Secure bit
// on in any environment other than development and test. Most of it
// is ripped straight from the buffalo source code;
// https://github.com/gobuffalo/buffalo/blob/9f469851d4d4b00652bf49701840ad41037e6a93/options.go#L153-L162
func SetSecureStore(opts buffalo.Options) buffalo.Options {
	store := sessions.NewCookieStore([]byte(getSecret(opts.Env)))
	if opts.Env != "development" && opts.Env != "test" {
		store.Options.Secure = true
		store.Options.HttpOnly = true
	}
	opts.SessionStore = store
	return opts
}

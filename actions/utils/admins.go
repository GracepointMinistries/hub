package utils

import (
	"strings"

	"github.com/gobuffalo/envy"
)

// GetAdmins returns a list of available admins
func GetAdmins() []string {
	return strings.Split(envy.Get("ADMINS", ""), ",")
}

// IsAdmin returns whether the given email is for an admin
func IsAdmin(email string) bool {
	for _, admin := range GetAdmins() {
		if email == admin {
			return true
		}
	}
	return false
}

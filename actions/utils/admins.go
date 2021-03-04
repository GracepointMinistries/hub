package utils

import (
	"strings"

	"github.com/gobuffalo/envy"
)

// GetAdmins returns a list of available admins
func GetAdmins() []string {
	return strings.Split(envy.Get("ADMINS", ""), ",")
}

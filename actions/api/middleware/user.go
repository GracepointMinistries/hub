package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/GracepointMinistries/hub/actions/utils"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

func getHeaderToken(c buffalo.Context) string {
	return strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
}

// RequireAdmin makes sure that the person issuing the request is an admin
func RequireAdmin(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		admin, sessionID, err := utils.ParseAdminToken(getHeaderToken(c))
		if err != nil {
			return c.Error(http.StatusUnauthorized, err)
		}
		if !utils.IsAdmin(admin) {
			// we have an old admin user, get rid of them
			if err := modelext.DeleteAdminSession(c, sessionID); err != nil {
				return c.Error(http.StatusInternalServerError, err)
			}
			return c.Error(http.StatusUnauthorized, errors.New("unauthorized"))
		}

		valid, err := modelext.ValidateAdminSession(c, admin, sessionID)
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		if !valid {
			return c.Error(http.StatusUnauthorized, errors.New("unauthorized"))
		}
		c.Set("Admin", admin)
		c.Set("SessionID", sessionID)
		return next(c)
	}
}

// RequireUser makes sure that the person issuing the request is logged in
func RequireUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		admin, id, sessionID, err := utils.ParseScopedToken(getHeaderToken(c))
		if err != nil {
			return c.Error(http.StatusUnauthorized, err)
		}
		if admin != "" {
			valid, err := modelext.ValidateAdminSession(c, admin, sessionID)
			if err != nil {
				return c.Error(http.StatusInternalServerError, err)
			}
			if !valid {
				return c.Error(http.StatusUnauthorized, errors.New("unauthorized"))
			}
		} else {
			valid, err := modelext.ValidateUserSession(c, id, sessionID)
			if err != nil {
				return c.Error(http.StatusInternalServerError, err)
			}
			if !valid {
				return c.Error(http.StatusUnauthorized, errors.New("unauthorized"))
			}
		}

		if admin != "" {
			c.Set("Admin", admin)
		}
		c.Set("ID", id)
		c.Set("SessionID", sessionID)
		return next(c)
	}
}

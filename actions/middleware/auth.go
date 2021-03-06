package middleware

import (
	"net/http"

	"github.com/GracepointMinistries/hub/actions/utils"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

//RequireAdmin ensures that the authenticated user is an admin
func RequireAdmin(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Session().Get("Admin") == nil || c.Session().Get("SessionID") == nil {
			return c.Redirect(http.StatusSeeOther, "/admin/auth/login")
		}
		admin := c.Session().Get("Admin").(string)
		sessionID := c.Session().Get("SessionID").(int)
		if !utils.IsAdmin(admin) {
			// we have an old admin user, get rid of them
			if err := modelext.DeleteAdminSession(c, sessionID); err != nil {
				return c.Error(http.StatusInternalServerError, err)
			}
			return c.Redirect(http.StatusSeeOther, "/admin/auth/login")
		}
		valid, err := modelext.ValidateAdminSession(c, admin, sessionID)
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		if !valid {
			return c.Redirect(http.StatusSeeOther, "/admin/auth/login")
		}
		c.Set("admin", c.Session().Get("Admin"))

		return next(c)
	}
}

//RequireLoggedInUser ensures that the user is authenticated
func RequireLoggedInUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Session().Get("ID") == nil || c.Session().Get("SessionID") == nil {
			return c.Redirect(http.StatusSeeOther, "/auth/login")
		}
		id := c.Session().Get("ID").(int)
		sessionID := c.Session().Get("SessionID").(int)
		if c.Session().Get("Admin") == nil {
			valid, err := modelext.ValidateUserSession(c, id, sessionID)
			if err != nil {
				return c.Error(http.StatusInternalServerError, err)
			}
			if !valid {
				return c.Redirect(http.StatusSeeOther, "/login")
			}
		} else {
			admin := c.Session().Get("Admin").(string)
			valid, err := modelext.ValidateAdminSession(c, admin, sessionID)
			if err != nil {
				return c.Error(http.StatusInternalServerError, err)
			}
			if !valid {
				return c.Redirect(http.StatusSeeOther, "/login")
			}
		}
		c.Set("id", c.Session().Get("ID"))
		c.Set("email", c.Session().Get("Email"))

		return next(c)
	}
}

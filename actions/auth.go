package actions

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/markbates/goth/gothic"
)

func getAdmins() []string {
	return strings.Split(envy.Get("ADMINS", ""), ",")
}

func adminCallback(c buffalo.Context) error {
	oauthUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(http.StatusUnauthorized, err)
	}

	for _, admin := range admins {
		if admin == oauthUser.Email {
			session, err := modelext.CreateAdminSession(c, admin, userIP(c))
			if err != nil {
				return c.Error(http.StatusInternalServerError, err)
			}
			c.Session().Set("Admin", admin)
			c.Session().Set("SessionID", session.ID)

			return c.Redirect(http.StatusSeeOther, "/admin")
		}
	}
	return c.Error(http.StatusUnauthorized, fmt.Errorf("%s is not an admin", oauthUser.Email))
}

func adminLoginPage(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin/login.html", "empty.html"))
}

func requireAdmin(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Session().Get("Admin") == nil || c.Session().Get("SessionID") == nil {
			return c.Redirect(http.StatusSeeOther, "/admin/login")
		}
		admin := c.Session().Get("Admin").(string)
		sessionID := c.Session().Get("SessionID").(int)
		valid, err := modelext.ValidateAdminSession(c, admin, sessionID)
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		if !valid {
			return c.Redirect(http.StatusSeeOther, "/admin/login")
		}
		c.Set("admin", c.Session().Get("Admin"))

		return next(c)
	}
}

func adminLogoutPage(c buffalo.Context) error {
	if err := modelext.DeleteAdminSession(c, c.Session().Get("SessionID").(int)); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Session().Clear()
	return c.Redirect(http.StatusSeeOther, "/admin/login")
}

func requireAPIAdmin(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		admin, sessionID, err := parseAdminToken(getHeaderToken(c))
		if err != nil {
			return c.Error(http.StatusUnauthorized, err)
		}
		valid, err := modelext.ValidateAdminSession(c, admin, sessionID)
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		if !valid {
			return c.Error(http.StatusUnauthorized, errors.New("unauthorized"))
		}
		c.Set("Admin", admin)
		return next(c)
	}
}

func authCallback(c buffalo.Context) error {
	oauthUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(http.StatusUnauthorized, err)
	}

	user, err := modelext.EnsureUserWithOauth(c, c.Param("provider"), oauthUser.UserID, oauthUser.Name, oauthUser.Email)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	session, err := modelext.CreateUserSession(c, user, userIP(c))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Session().Set("ID", user.ID)
	c.Session().Set("SessionID", session.ID)
	c.Session().Set("Email", user.Email)
	return c.Redirect(http.StatusSeeOther, "/")
}

func loginPage(c buffalo.Context) error {
	return c.Render(200, r.HTML("user/login.html", "empty.html"))
}

func requireLoggedInUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Session().Get("ID") == nil || c.Session().Get("SessionID") == nil {
			return c.Redirect(http.StatusSeeOther, "/login")
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

func logoutPage(c buffalo.Context) error {
	c.Session().Delete("ID")
	c.Session().Delete("Email")
	if c.Session().Get("Admin") != nil {
		return c.Redirect(http.StatusSeeOther, "/admin")
	}
	if err := modelext.DeleteUserSession(c, c.Session().Get("SessionID").(int)); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Session().Delete("SessionID")
	return c.Redirect(http.StatusSeeOther, "/")
}

func requireAPIUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		id, sessionID, err := parseUserToken(getHeaderToken(c))
		if err != nil {
			return c.Error(http.StatusUnauthorized, err)
		}
		valid, err := modelext.ValidateUserSession(c, id, sessionID)
		if err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		if !valid {
			return c.Error(http.StatusUnauthorized, errors.New("unauthorized"))
		}

		c.Set("ID", id)
		return next(c)
	}
}

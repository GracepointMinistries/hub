package actions

import (
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
			c.Session().Set("Admin", admin)
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
		if c.Session().Get("Admin") == nil {
			return c.Redirect(http.StatusSeeOther, "/admin/login")
		}
		c.Set("admin", c.Session().Get("Admin"))

		return next(c)
	}
}

func adminLogoutPage(c buffalo.Context) error {
	c.Session().Clear()
	return c.Redirect(http.StatusSeeOther, "/admin/login")
}

func requireAPIAdmin(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Session().Get("Admin") == nil {
			return c.Render(http.StatusUnauthorized, apiError("unauthorized"))
		}
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

	c.Session().Set("ID", user.ID)
	c.Session().Set("Email", user.Email)
	return c.Redirect(http.StatusSeeOther, "/")
}

func loginPage(c buffalo.Context) error {
	return c.Render(200, r.HTML("user/login.html", "empty.html"))
}

func requireLoggedInUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Session().Get("ID") == nil {
			return c.Redirect(http.StatusSeeOther, "/login")
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
	return c.Redirect(http.StatusSeeOther, "/")
}

func requireAPIUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if c.Session().Get("ID") == nil {
			return c.Render(http.StatusUnauthorized, apiError("unauthorized"))
		}
		return next(c)
	}
}

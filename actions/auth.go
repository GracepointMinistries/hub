package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth/gothic"
)

func authCallback(c buffalo.Context) error {
	oauthUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(http.StatusUnauthorized, err)
	}

	user, err := models.EnsureUserWithOAuth(c, c.Param("provider"), oauthUser.UserID, oauthUser.Email)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Session().Set("ID", user.ID)
	c.Session().Set("Email", user.Email)
	return c.Redirect(http.StatusSeeOther, "/")
}

// requireLoggedInUser checks whether or not a user is logged in with an unexpired session cookie
// if the user is not logged in, then the frontend should redirect to /auth/google
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

func logoutHandler(c buffalo.Context) error {
	c.Session().Clear()
	return c.Redirect(http.StatusSeeOther, "/")
}

package auth

import (
	"net/http"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/actions/utils"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth/gothic"
)

// Callback is the auth callback endpoint for log in
func Callback(c buffalo.Context) error {
	oauthUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(http.StatusUnauthorized, err)
	}

	user, err := modelext.EnsureUserWithOauth(c, c.Param("provider"), oauthUser.UserID, oauthUser.Name, oauthUser.Email)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	session, err := modelext.CreateUserSession(c, user, utils.UserIP(c))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	c.Session().Set("ID", user.ID)
	c.Session().Set("SessionID", session.ID)
	c.Session().Set("Email", user.Email)
	return c.Redirect(http.StatusSeeOther, "/")
}

// Login is the page for logging in
func Login(c buffalo.Context) error {
	return c.Render(200, render.HTML("user/login.html", "empty.html"))
}

// Logout is the page for logging out
func Logout(c buffalo.Context) error {
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

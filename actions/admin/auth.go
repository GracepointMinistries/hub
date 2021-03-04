package admin

import (
	"fmt"
	"net/http"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/actions/utils"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth/gothic"
)

// Callback is the auth callback endpoint for administrative log in
func Callback(c buffalo.Context) error {
	oauthUser, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(http.StatusUnauthorized, err)
	}

	for _, admin := range utils.GetAdmins() {
		if admin == oauthUser.Email {
			session, err := modelext.CreateAdminSession(c, admin, utils.UserIP(c))
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

// Login is the administrative page for logging in
func Login(c buffalo.Context) error {
	return c.Render(200, render.HTML("admin/login.html", "empty.html"))
}

// Logout is the administrative page for logging out
func Logout(c buffalo.Context) error {
	if err := modelext.DeleteAdminSession(c, c.Session().Get("SessionID").(int)); err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Session().Clear()
	return c.Redirect(http.StatusSeeOther, "/admin")
}

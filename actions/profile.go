package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/actions/render"
	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

// Profile is the page rendered when viewing the user's profile
func Profile(c buffalo.Context) error {
	user, err := modelext.FindUser(c, c.Session().Get("ID").(int))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("user", user)
	c.Set("group", user.Group)
	return c.Render(http.StatusOK, render.HTML("user/profile.html"))
}

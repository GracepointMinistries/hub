package actions

import (
	"net/http"

	"github.com/GracepointMinistries/hub/modelext"
	"github.com/gobuffalo/buffalo"
)

func profilePage(c buffalo.Context) error {
	user, err := modelext.FindUser(c, c.Session().Get("ID").(int))
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}
	c.Set("user", user)
	c.Set("zgroup", user.Zgroup)
	return c.Render(http.StatusOK, r.HTML("user/profile.html"))
}
